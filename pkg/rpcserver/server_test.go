package rpcserver_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"testing"

	"github.com/buildbarn/go-xdr/internal/mock"
	"github.com/buildbarn/go-xdr/pkg/protocols/rpcv2"
	"github.com/buildbarn/go-xdr/pkg/rpcserver"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestServer(t *testing.T) {
	ctrl := gomock.NewController(t)

	service := mock.NewMockService(ctrl)
	s := rpcserver.NewServer(map[uint32]rpcserver.Service{
		123: service.Call,
	})

	t.Run("EOF", func(t *testing.T) {
		// It's valid for a client to close the connection
		// without sending any requests.
		w := mock.NewMockWriter(ctrl)
		require.NoError(t, s.HandleConnection(bytes.NewBuffer(nil), w))
	})

	t.Run("TruncatedRecordMarker", func(t *testing.T) {
		// The transmission should start with four bytes
		// containing the record marker. Sending fewer bytes
		// than this is incorrect.
		w := mock.NewMockWriter(ctrl)
		require.Equal(
			t,
			io.ErrUnexpectedEOF,
			s.HandleConnection(bytes.NewBuffer([]byte{
				0x12,
			}), w))
	})

	t.Run("EOFAfterRecordMarker", func(t *testing.T) {
		// The record marker indicates that 256 bytes of data
		// should be provided. In practice, no data was
		// returned.
		w := mock.NewMockWriter(ctrl)
		require.Equal(
			t,
			io.ErrUnexpectedEOF,
			s.HandleConnection(bytes.NewBuffer([]byte{
				0x80, 0x00, 0x01, 0x00,
			}), w))
	})

	t.Run("EmptyRecord", func(t *testing.T) {
		// The payload should be an rpc_msg, so we can't have
		// zero sized records.
		w := mock.NewMockWriter(ctrl)
		require.Equal(
			t,
			errors.New("attempted to read beyond end of record"),
			s.HandleConnection(bytes.NewBuffer([]byte{
				0x80, 0x00, 0x00, 0x00,
			}), w))
	})

	t.Run("InvalidMessageType", func(t *testing.T) {
		// Payload with an unknown message type.
		w := mock.NewMockWriter(ctrl)
		require.Equal(
			t,
			errors.New("discriminant rpc_msg.body.mtype has unknown value 3"),
			s.HandleConnection(bytes.NewBuffer([]byte{
				// Record marker.
				0x80, 0x00, 0x01, 0x00,
				// XID.
				0x97, 0x76, 0x78, 0x21,
				// Unknown msg_type value.
				0x00, 0x00, 0x00, 0x03,
			}), w))
	})

	t.Run("MessageTypeReply", func(t *testing.T) {
		// Attempted to send a REPLY message to a server. This
		// is invalid, as we should see CALL messages instead.
		w := mock.NewMockWriter(ctrl)
		require.Equal(
			t,
			errors.New("RPC message is not of type CALL"),
			s.HandleConnection(bytes.NewBuffer([]byte{
				// Record marker.
				0x80, 0x00, 0x00, 0x18,
				// xid.
				0x97, 0x76, 0x78, 0x21,
				// body.msg_type == REPLY.
				0x00, 0x00, 0x00, 0x01,
				// body.rbody.reply_stat == MSG_DENIED.
				0x00, 0x00, 0x00, 0x01,
				// body.rbody.rreply.reject_stat == RPC_MISMATCH.
				0x00, 0x00, 0x00, 0x00,
				// body.rbody.rreply.mismatch_info.low == 2.
				0x00, 0x00, 0x00, 0x02,
				// body.rbody.rreply.mismatch_info.high == 2.
				0x00, 0x00, 0x00, 0x02,
			}), w))
	})

	t.Run("InvalidRPCVersion", func(t *testing.T) {
		// Send a request that uses RPC version 3. This should
		// cause us to return an RPC_MISMATCH response.
		w := mock.NewMockWriter(ctrl)
		w.EXPECT().Write([]byte{
			// Record marker.
			0x80, 0x00, 0x00, 0x18,
			// xid.
			0xa3, 0x26, 0x71, 0xfc,
			// body.msg_type == REPLY.
			0x00, 0x00, 0x00, 0x01,
			// body.rbody.reply_stat == MSG_DENIED.
			0x00, 0x00, 0x00, 0x01,
			// body.rbody.rreply.reject_stat == RPC_MISMATCH.
			0x00, 0x00, 0x00, 0x00,
			// body.rbody.rreply.mismatch_info.low == 2.
			0x00, 0x00, 0x00, 0x02,
			// body.rbody.rreply.mismatch_info.high == 2.
			0x00, 0x00, 0x00, 0x02,
		}).Return(24, nil)

		require.Equal(
			t,
			nil,
			s.HandleConnection(bytes.NewBuffer([]byte{
				// Record marker.
				0x80, 0x00, 0x00, 0x30,
				// xid.
				0xa3, 0x26, 0x71, 0xfc,
				// body.msg_type == CALL.
				0x00, 0x00, 0x00, 0x00,
				// body.cbody.rpcvers == 3.
				0x00, 0x00, 0x00, 0x03,
				// body.cbody.prog == 10.
				0x00, 0x00, 0x00, 0x0a,
				// body.cbody.vers == 7.
				0x00, 0x00, 0x00, 0x07,
				// body.cbody.proc == 4.
				0x00, 0x00, 0x00, 0x04,
				// body.cbody.cred.flavor == AUTH_NONE,
				0x00, 0x00, 0x00, 0x00,
				// body.cbody.cred.body.
				0x00, 0x00, 0x00, 0x00,
				// body.cbody.verf.flavor == AUTH_NONE,
				0x00, 0x00, 0x00, 0x00,
				// body.cbody.verf.body.
				0x00, 0x00, 0x00, 0x00,
				// Some payload that follows.
				0xb3, 0x83, 0x19, 0x90, 0x0a, 0xe0, 0xf1, 0x2a,
			}), w))
	})

	t.Run("ProgramUnavailable", func(t *testing.T) {
		// Sending an RPC for an unknown program number should
		// cause us to return PROG_UNAVAIL.
		w := mock.NewMockWriter(ctrl)
		w.EXPECT().Write([]byte{
			// Record marker.
			0x80, 0x00, 0x00, 0x18,
			// xid.
			0x99, 0x12, 0x0f, 0xdc,
			// body.msg_type == REPLY.
			0x00, 0x00, 0x00, 0x01,
			// body.rbody.reply_stat == MSG_ACCEPTED.
			0x00, 0x00, 0x00, 0x00,
			// body.rbody.areply.verf.flavor == AUTH_NONE.
			0x00, 0x00, 0x00, 0x00,
			// body.rbody.areply.verf.body.
			0x00, 0x00, 0x00, 0x00,
			// body.rbody.areply.stat == PROG_UNAVAIL.
			0x00, 0x00, 0x00, 0x01,
		}).Return(28, nil)

		require.Equal(
			t,
			nil,
			s.HandleConnection(bytes.NewBuffer([]byte{
				// Record marker.
				0x80, 0x00, 0x00, 0x30,
				// xid.
				0x99, 0x12, 0x0f, 0xdc,
				// body.msg_type == CALL.
				0x00, 0x00, 0x00, 0x00,
				// body.cbody.rpcvers == 2.
				0x00, 0x00, 0x00, 0x02,
				// body.cbody.prog == 10.
				0x00, 0x00, 0x00, 0x0a,
				// body.cbody.vers == 7.
				0x00, 0x00, 0x00, 0x07,
				// body.cbody.proc == 4.
				0x00, 0x00, 0x00, 0x04,
				// body.cbody.cred.flavor == AUTH_NONE,
				0x00, 0x00, 0x00, 0x00,
				// body.cbody.cred.body.
				0x00, 0x00, 0x00, 0x00,
				// body.cbody.verf.flavor == AUTH_NONE,
				0x00, 0x00, 0x00, 0x00,
				// body.cbody.verf.body.
				0x00, 0x00, 0x00, 0x00,
				// Some payload that follows.
				0xb3, 0x83, 0x19, 0x90, 0x0a, 0xe0, 0xf1, 0x2a,
			}), w))
	})

	t.Run("Success", func(t *testing.T) {
		// Simulate the case where two valid requests are
		// transmitted. Both these requests should trigger a
		// call into the Service. Because the requests may be
		// handled in parallel, responses can be written in any
		// order.
		service.EXPECT().Call(gomock.Any(), uint32(7), uint32(4), gomock.Any(), gomock.Any()).
			DoAndReturn(func(ctx context.Context, vers, proc uint32, parameters io.ReadCloser, returnValue io.Writer) (*rpcv2.AcceptedReply, error) {
				var in [8]byte
				n, err := parameters.Read(in[:])
				require.Equal(t, 8, n)
				require.NoError(t, err)
				require.Equal(t, [...]byte{0xb3, 0x83, 0x19, 0x90, 0x0a, 0xe0, 0xf1, 0x2a}, in)
				require.NoError(t, parameters.Close())

				n, err = returnValue.Write([]byte{0x44, 0xe5, 0x33, 0x30, 0xa5, 0xf4, 0x75, 0xbb})
				require.Equal(t, 8, n)
				require.NoError(t, err)

				return &rpcv2.AcceptedReply{
					ReplyData: &rpcv2.AcceptedReplyReplyData_SUCCESS{},
				}, nil
			})
		service.EXPECT().Call(gomock.Any(), uint32(3), uint32(9), gomock.Any(), gomock.Any()).
			DoAndReturn(func(ctx context.Context, vers, proc uint32, parameters io.ReadCloser, returnValue io.Writer) (*rpcv2.AcceptedReply, error) {
				var in [8]byte
				n, err := parameters.Read(in[:])
				require.Equal(t, 8, n)
				require.NoError(t, err)
				require.Equal(t, [...]byte{0xa9, 0x73, 0x1c, 0xfa, 0xfe, 0x16, 0xe0, 0x81}, in)
				require.NoError(t, parameters.Close())

				n, err = returnValue.Write([]byte{0x26, 0xb5, 0x37, 0xb0, 0xe4, 0xf4, 0x6a, 0x84})
				require.Equal(t, 8, n)
				require.NoError(t, err)

				return &rpcv2.AcceptedReply{
					ReplyData: &rpcv2.AcceptedReplyReplyData_SUCCESS{},
				}, nil
			})
		w := mock.NewMockWriter(ctrl)
		w.EXPECT().Write([]byte{
			// Record marker.
			0x80, 0x00, 0x00, 0x20,
			// xid.
			0x44, 0xe5, 0x33, 0x30,
			// body.msg_type == REPLY.
			0x00, 0x00, 0x00, 0x01,
			// body.rbody.reply_stat == MSG_ACCEPTED.
			0x00, 0x00, 0x00, 0x00,
			// body.rbody.areply.verf.flavor == AUTH_NONE.
			0x00, 0x00, 0x00, 0x00,
			// body.rbody.areply.verf.body.
			0x00, 0x00, 0x00, 0x00,
			// body.rbody.areply.stat == PROG_UNAVAIL.
			0x00, 0x00, 0x00, 0x00,
			// Return value.
			0x44, 0xe5, 0x33, 0x30, 0xa5, 0xf4, 0x75, 0xbb,
		}).Return(36, nil)
		w.EXPECT().Write([]byte{
			// Record marker.
			0x80, 0x00, 0x00, 0x20,
			// xid.
			0x35, 0x91, 0xaa, 0x5a,
			// body.msg_type == REPLY.
			0x00, 0x00, 0x00, 0x01,
			// body.rbody.reply_stat == MSG_ACCEPTED.
			0x00, 0x00, 0x00, 0x00,
			// body.rbody.areply.verf.flavor == AUTH_NONE.
			0x00, 0x00, 0x00, 0x00,
			// body.rbody.areply.verf.body.
			0x00, 0x00, 0x00, 0x00,
			// body.rbody.areply.stat == PROG_UNAVAIL.
			0x00, 0x00, 0x00, 0x00,
			// Return value.
			0x26, 0xb5, 0x37, 0xb0, 0xe4, 0xf4, 0x6a, 0x84,
		}).Return(36, nil)

		require.Equal(
			t,
			nil,
			s.HandleConnection(bytes.NewBuffer([]byte{
				// Record marker.
				0x80, 0x00, 0x00, 0x30,
				// xid.
				0x44, 0xe5, 0x33, 0x30,
				// body.msg_type == CALL.
				0x00, 0x00, 0x00, 0x00,
				// body.cbody.rpcvers == 2.
				0x00, 0x00, 0x00, 0x02,
				// body.cbody.prog == 123.
				0x00, 0x00, 0x00, 0x7b,
				// body.cbody.vers == 7.
				0x00, 0x00, 0x00, 0x07,
				// body.cbody.proc == 4.
				0x00, 0x00, 0x00, 0x04,
				// body.cbody.cred.flavor == AUTH_NONE,
				0x00, 0x00, 0x00, 0x00,
				// body.cbody.cred.body.
				0x00, 0x00, 0x00, 0x00,
				// body.cbody.verf.flavor == AUTH_NONE,
				0x00, 0x00, 0x00, 0x00,
				// body.cbody.verf.body.
				0x00, 0x00, 0x00, 0x00,
				// Some payload that follows.
				0xb3, 0x83, 0x19, 0x90, 0x0a, 0xe0, 0xf1, 0x2a,

				// Record marker.
				0x80, 0x00, 0x00, 0x30,
				// xid.
				0x35, 0x91, 0xaa, 0x5a,
				// body.msg_type == CALL.
				0x00, 0x00, 0x00, 0x00,
				// body.cbody.rpcvers == 2.
				0x00, 0x00, 0x00, 0x02,
				// body.cbody.prog == 123.
				0x00, 0x00, 0x00, 0x7b,
				// body.cbody.vers == 3.
				0x00, 0x00, 0x00, 0x03,
				// body.cbody.proc == 9.
				0x00, 0x00, 0x00, 0x09,
				// body.cbody.cred.flavor == AUTH_NONE,
				0x00, 0x00, 0x00, 0x00,
				// body.cbody.cred.body.
				0x00, 0x00, 0x00, 0x00,
				// body.cbody.verf.flavor == AUTH_NONE,
				0x00, 0x00, 0x00, 0x00,
				// body.cbody.verf.body.
				0x00, 0x00, 0x00, 0x00,
				// Some payload that follows.
				0xa9, 0x73, 0x1c, 0xfa, 0xfe, 0x16, 0xe0, 0x81,
			}), w))
	})
}
