package rpcserver

import (
	"bufio"
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"io"

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
//
// To write a successful return value, the service calls newReturnValue
// once with its encoded size in bytes. Replies without a return value do
// not need to call newReturnValue.
type Service func(ctx context.Context, vers, proc uint32, parameters io.ReadCloser, newReturnValue func(sizeBytes int) io.Writer) (rpcv2.AcceptedReplyData, error)

// Server of ONC RPCv2, as described in RFC 5531.
type Server struct {
	services      map[uint32]Service
	authenticator Authenticator
}

// NewServer creates a new Server that is capable of accepting incoming
// requests for a provided set of services.
func NewServer(services map[uint32]Service, authenticator Authenticator) *Server {
	return &Server{
		services:      services,
		authenticator: authenticator,
	}
}

// HandleConnection reads requests from a streaming network connection
// and provides them to one or more services, based on the program
// number in the request. After the service has processed the request
// asynchronously, a response is written back to the network connection.
func (s *Server) HandleConnection(r io.Reader, w io.Writer) error {
	group, ctx := errgroup.WithContext(context.Background())
	ch := connectionHandler{
		server: s,

		group:   group,
		context: ctx,

		reader: bufio.NewReader(r),
		writer: w,
	}
	ch.startReadingMessages()
	return group.Wait()
}

type connectionHandler struct {
	server *Server

	group   *errgroup.Group
	context context.Context

	reader io.Reader
	writer io.Writer
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
		var callRPCMessage rpcv2.RpcMsg
		if _, err := callRPCMessage.ReadFrom(&rmr); err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		// Extract call body.
		body, ok := callRPCMessage.Body.(*rpcv2.RpcMsgBody_CALL)
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
				return ch.replyWithoutReturnValue(&rpcv2.RpcMsg{
					Xid: callRPCMessage.Xid,
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
				})
			})
			continue
		}

		// Extract credentials.
		server := ch.server
		ctxWithCancel, cancelContext := context.WithCancel(ch.context)
		ctxWithAuth, replyVerifier, authStat := server.authenticator.Authenticate(ctxWithCancel, &callBody.Cred, &callBody.Verf)
		if authStat != rpcv2.AUTH_OK {
			// Authentication failed.
			cancelContext()
			if err := rmr.discardRemaining(); err != nil {
				return err
			}
			ch.group.Go(func() error {
				return ch.replyWithoutReturnValue(&rpcv2.RpcMsg{
					Xid: callRPCMessage.Xid,
					Body: &rpcv2.RpcMsgBody_REPLY{
						Rbody: &rpcv2.ReplyBody_MSG_DENIED{
							Rreply: &rpcv2.RejectedReply_AUTH_ERROR{
								Stat: authStat,
							},
						},
					},
				})
			})
			continue
		}

		service, ok := server.services[callBody.Prog]
		if !ok {
			// No service associated with this
			// program number.
			cancelContext()
			if err := rmr.discardRemaining(); err != nil {
				return err
			}
			ch.group.Go(func() error {
				return ch.replyWithoutReturnValue(&rpcv2.RpcMsg{
					Xid: callRPCMessage.Xid,
					Body: &rpcv2.RpcMsgBody_REPLY{
						Rbody: &rpcv2.ReplyBody_MSG_ACCEPTED{
							Areply: rpcv2.AcceptedReply{
								Verf: replyVerifier,
								ReplyData: &rpcv2.AcceptedReplyData_default{
									Stat: rpcv2.PROG_UNAVAIL,
								},
							},
						},
					},
				})
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

		replyBody := rpcv2.ReplyBody_MSG_ACCEPTED{
			Areply: rpcv2.AcceptedReply{
				Verf:      replyVerifier,
				ReplyData: &rpcv2.AcceptedReplyData_SUCCESS{},
			},
		}
		replyRPCMessage := rpcv2.RpcMsg{
			Xid: callRPCMessage.Xid,
			Body: &rpcv2.RpcMsgBody_REPLY{
				Rbody: &replyBody,
			},
		}
		replyReturnValueStartBytes := recordMarkerSizeBytes + replyRPCMessage.GetEncodedSizeBytes()
		var replyBuffer *bytes.Buffer
		replyData, err := service(ctxWithAuth, callBody.Vers, callBody.Proc, &rc, func(sizeBytes int) io.Writer {
			replyBuffer = bytes.NewBuffer(make([]byte, replyReturnValueStartBytes, replyReturnValueStartBytes+sizeBytes))
			return replyBuffer
		})
		cancelContext()
		rc.Close()
		if err != nil {
			return err
		}

		replyBody.Areply.ReplyData = replyData
		if replyBuffer == nil || replyData.GetStat() != rpcv2.SUCCESS {
			return ch.replyWithoutReturnValue(&replyRPCMessage)
		}
		reply := replyBuffer.Bytes()
		if _, err := replyRPCMessage.WriteTo(bytes.NewBuffer(reply[recordMarkerSizeBytes:recordMarkerSizeBytes])); err != nil {
			return err
		}
		binary.BigEndian.PutUint32(reply, uint32(len(reply)-recordMarkerSizeBytes)|0x80000000)

		_, err = ch.writer.Write(reply)
		return err
	}
}

func (ch *connectionHandler) replyWithoutReturnValue(rpcMessage *rpcv2.RpcMsg) error {
	// Serialize the reply RPC message and prepend a record marker.
	rpcMessageSizeBytes := rpcMessage.GetEncodedSizeBytes()
	replyBuffer := bytes.NewBuffer(make([]byte, recordMarkerSizeBytes, recordMarkerSizeBytes+rpcMessageSizeBytes))
	if _, err := rpcMessage.WriteTo(replyBuffer); err != nil {
		return err
	}
	reply := replyBuffer.Bytes()
	binary.BigEndian.PutUint32(reply, uint32(rpcMessageSizeBytes)|0x80000000)

	// Write the record marker and RPC message.
	_, err := ch.writer.Write(reply)
	return err
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
