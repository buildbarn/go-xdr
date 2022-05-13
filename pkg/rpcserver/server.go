package rpcserver

import (
	"bufio"
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"io"
	"sync"

	"github.com/buildbarn/go-xdr/pkg/protocols/rpcv2"

	"golang.org/x/sync/errgroup"
)

// recordMarkerSizeBytes is the size of record markers that are
// prepended to RPC messages that are transmitted across a streaming
// protocol, such as TCP.
const recordMarkerSizeBytes = 4

// Service that is capable of processing incoming requests on the RPC
// server.
//
// The service is responsible for reading parameters from a ReadCloser.
// It should close it immediately after it has finished reading the
// parameters, so that the RPC server can continue to read the next
// incoming request.
type Service func(ctx context.Context, vers, proc uint32, parameters io.ReadCloser, returnValue io.Writer) (*rpcv2.AcceptedReply, error)

// Server of ONC RPCv2, as described in RFC 5531.
type Server struct {
	services map[uint32]Service
}

// NewServer creates a new Server that is capable of accepting incoming
// requests for a provided set of services.
func NewServer(services map[uint32]Service) *Server {
	return &Server{
		services: services,
	}
}

// HandleConnection reads requests from a streaming network connection
// and provides them to one or more services, based on the program
// number in the request. After the service has processed the request
// asynchronously, a response is written back to the network connection.
func (s *Server) HandleConnection(r io.Reader, w io.Writer) error {
	group, ctx := errgroup.WithContext(context.Background())
	ch := connectionHandler{
		services: s.services,

		group:   group,
		context: ctx,

		reader: bufio.NewReader(r),
		writer: w,
	}
	ch.startReadingMessages()
	return group.Wait()
}

type connectionHandler struct {
	services map[uint32]Service

	group   *errgroup.Group
	context context.Context

	reader     io.Reader
	writerLock sync.Mutex
	writer     io.Writer
}

func (ch *connectionHandler) startReadingMessages() {
	ch.group.Go(ch.readMessages)
}

func (ch *connectionHandler) readMessages() error {
	for {
		rmr := recordMarkingReader{
			reader: ch.reader,
		}

		// Extract RPC message.
		var rpcMessage rpcv2.RpcMsg
		if _, err := rpcMessage.ReadFrom(&rmr); err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		// Extract call body.
		body, ok := rpcMessage.Body.(*rpcv2.RpcMsgBody_CALL)
		if !ok {
			return errors.New("RPC message is not of type CALL")
		}
		callBody := &body.Cbody

		if callBody.Rpcvers != 2 {
			// We only support RPCv2.
			if err := rmr.discardRemaining(); err != nil {
				return err
			}
			ch.group.Go(func() error {
				return ch.reply(&rpcv2.RpcMsg{
					Xid: rpcMessage.Xid,
					Body: &rpcv2.RpcMsgBody_REPLY{
						Rbody: &rpcv2.ReplyBody_MSG_DENIED{
							Rreply: &rpcv2.RejectedReply_RPC_MISMATCH{
								MismatchInfo: struct {
									Low  uint32
									High uint32
								}{
									Low:  2,
									High: 2,
								},
							},
						},
					},
				}, nil)
			})
			continue
		}

		service, ok := ch.services[callBody.Prog]
		if !ok {
			// No service associated with this
			// program number.
			if err := rmr.discardRemaining(); err != nil {
				return err
			}
			ch.group.Go(func() error {
				return ch.reply(&rpcv2.RpcMsg{
					Xid: rpcMessage.Xid,
					Body: &rpcv2.RpcMsgBody_REPLY{
						Rbody: &rpcv2.ReplyBody_MSG_ACCEPTED{
							Areply: rpcv2.AcceptedReply{
								ReplyData: &rpcv2.AcceptedReplyReplyData_default{
									Stat: rpcv2.PROG_UNAVAIL,
								},
							},
						},
					},
				}, nil)
			})
			continue
		}

		// Request against a known program. Forward the request
		// to the Service. When the Service has finished reading
		// the request, we create a new goroutine to read the
		// next incoming request.
		rc := successiveReadTriggeringReader{
			recordMarkingReader: &rmr,
			connectionHandler:   ch,
		}
		returnValueBuffer := bytes.NewBuffer(nil)
		acceptedReply, err := service(ch.context, callBody.Vers, callBody.Proc, &rc, returnValueBuffer)
		rc.Close()
		if err != nil {
			return err
		}

		return ch.reply(&rpcv2.RpcMsg{
			Xid: rpcMessage.Xid,
			Body: &rpcv2.RpcMsgBody_REPLY{
				Rbody: &rpcv2.ReplyBody_MSG_ACCEPTED{
					Areply: *acceptedReply,
				},
			},
		}, returnValueBuffer.Bytes())
	}
}

func (ch *connectionHandler) reply(rpcMessage *rpcv2.RpcMsg, returnValue []byte) error {
	// Serialize the reply RPC message and prepend a record marker.
	rpcMessageSizeBytes := rpcMessage.GetEncodedSizeBytes()
	replyBuffer := bytes.NewBuffer(make([]byte, recordMarkerSizeBytes, recordMarkerSizeBytes+rpcMessageSizeBytes))
	if _, err := rpcMessage.WriteTo(replyBuffer); err != nil {
		return err
	}
	reply := replyBuffer.Bytes()
	binary.BigEndian.PutUint32(reply, uint32(rpcMessageSizeBytes+len(returnValue))|0x80000000)

	ch.writerLock.Lock()
	defer ch.writerLock.Unlock()

	// Write the record marker, RPC message and return value.
	if _, err := ch.writer.Write(reply); err != nil {
		return err
	}
	if len(returnValue) > 0 {
		if _, err := ch.writer.Write(returnValue); err != nil {
			return err
		}
	}
	return nil
}

// recordMarkingReader is an io.Reader that reads a single record from a
// stream according to the ONC RPCv2 record marking standard. A single
// record may consist of one or more record fragments.
//
// More details: RFC 5531, chapter 11.
type recordMarkingReader struct {
	reader             io.Reader
	bytesInFragment    int
	successiveFragment bool
	lastFragment       bool
}

func (r *recordMarkingReader) done() bool {
	return r.bytesInFragment == 0 && r.lastFragment
}

func (r *recordMarkingReader) maybeReadRecordMarker() error {
	if r.bytesInFragment == 0 {
		if r.lastFragment {
			return errors.New("attempted to read beyond end of record")
		}
		var headerBytes [recordMarkerSizeBytes]byte
		if _, err := io.ReadFull(r.reader, headerBytes[:]); err != nil {
			if err == io.EOF && r.successiveFragment {
				return io.ErrUnexpectedEOF
			}
			return err
		}
		header := binary.BigEndian.Uint32(headerBytes[:])
		r.bytesInFragment = int(header &^ 0x80000000)
		r.successiveFragment = true
		r.lastFragment = header&0x80000000 != 0
	}
	return nil
}

func (r *recordMarkingReader) Read(p []byte) (int, error) {
	nTotal := 0
	for len(p) > 0 {
		if err := r.maybeReadRecordMarker(); err != nil {
			return nTotal, err
		}
		pRead := p
		if len(pRead) > r.bytesInFragment {
			pRead = pRead[:r.bytesInFragment]
		}
		n, err := r.reader.Read(pRead)
		nTotal += n
		p = p[n:]
		r.bytesInFragment -= n
		if err != nil {
			if err == io.EOF && !r.done() {
				return nTotal, io.ErrUnexpectedEOF
			}
			return nTotal, err
		}
	}
	return nTotal, nil
}

func (r *recordMarkingReader) discardRemaining() error {
	for !r.done() {
		if err := r.maybeReadRecordMarker(); err != nil {
			return err
		}
		nCopied, err := io.Copy(io.Discard, io.LimitReader(r.reader, int64(r.bytesInFragment)))
		if err != nil {
			return err
		}
		if int(nCopied) != r.bytesInFragment {
			return io.ErrUnexpectedEOF
		}
		r.bytesInFragment = 0
	}
	return nil
}

// successiveReadTriggeringReader is used to detect when an
// implementation of Service has finished reading the arguments of the
// incoming request from the socket. When finished, it spawns a new
// goroutine that will read the next incoming request from the socket.
type successiveReadTriggeringReader struct {
	*recordMarkingReader
	connectionHandler *connectionHandler
}

func (rc *successiveReadTriggeringReader) Close() error {
	if err := rc.recordMarkingReader.discardRemaining(); err != nil {
		return err
	}
	if rc.connectionHandler != nil {
		rc.connectionHandler.startReadingMessages()
		rc.connectionHandler = nil
	}
	return nil
}
