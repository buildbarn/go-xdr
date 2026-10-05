package rpcserver_test

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"io"
	"testing"

	"github.com/buildbarn/go-xdr/internal/mock"
	"github.com/buildbarn/go-xdr/pkg/protocols/nfsv4"
	"github.com/buildbarn/go-xdr/pkg/protocols/rpcv2"
	"github.com/buildbarn/go-xdr/pkg/rpcserver"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestServer(t *testing.T) {
	ctrl := gomock.NewController(t)

	service := mock.NewMockService(ctrl)
	authenticator := mock.NewMockAuthenticator(ctrl)
	s := rpcserver.NewServer(map[uint32]rpcserver.Service{
		123: service.Call,
	}, authenticator)

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
		}).Return(28, nil)

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

	t.Run("Unauthenticated", func(t *testing.T) {
		// Let the authenticator reject the RPC.
		authenticator.EXPECT().Authenticate(gomock.Any(), &rpcv2.OpaqueAuth{Flavor: rpcv2.AUTH_NONE}, &rpcv2.OpaqueAuth{Flavor: rpcv2.AUTH_NONE}).
			DoAndReturn(func(ctx context.Context, credentials, verifier *rpcv2.OpaqueAuth) (context.Context, rpcv2.OpaqueAuth, rpcv2.AuthStat) {
				return nil, rpcv2.OpaqueAuth{}, rpcv2.AUTH_BADCRED
			})
		w := mock.NewMockWriter(ctrl)
		w.EXPECT().Write([]byte{
			// Record marker.
			0x80, 0x00, 0x00, 0x14,
			// xid.
			0xc2, 0xa7, 0xfb, 0xc6,
			// body.msg_type == REPLY.
			0x00, 0x00, 0x00, 0x01,
			// body.rbody.reply_stat == MSG_DENIED.
			0x00, 0x00, 0x00, 0x01,
			// body.rbody.rreply.reject_stat == AUTH_ERROR.
			0x00, 0x00, 0x00, 0x01,
			// body.rbody.rreply.stat == AUTH_BADCRED.
			0x00, 0x00, 0x00, 0x01,
		}).Return(24, nil)

		require.Equal(
			t,
			nil,
			s.HandleConnection(bytes.NewBuffer([]byte{
				// Record marker.
				0x80, 0x00, 0x00, 0x30,
				// xid.
				0xc2, 0xa7, 0xfb, 0xc6,
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

	t.Run("ProgramUnavailable", func(t *testing.T) {
		// Sending an RPC for an unknown program number should
		// cause us to return PROG_UNAVAIL.
		authenticator.EXPECT().Authenticate(gomock.Any(), &rpcv2.OpaqueAuth{Flavor: rpcv2.AUTH_NONE}, &rpcv2.OpaqueAuth{Flavor: rpcv2.AUTH_NONE}).
			DoAndReturn(func(ctx context.Context, credentials, verifier *rpcv2.OpaqueAuth) (context.Context, rpcv2.OpaqueAuth, rpcv2.AuthStat) {
				return ctx, rpcv2.OpaqueAuth{Flavor: rpcv2.AUTH_NONE}, rpcv2.AUTH_OK
			})
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
		mockContext1 := mock.NewMockContext(ctrl)
		authenticator.EXPECT().Authenticate(gomock.Any(), &rpcv2.OpaqueAuth{
			Flavor: rpcv2.AUTH_SHORT,
			Body:   []byte{0xe6, 0xda, 0x1d, 0x8d, 0x27, 0x35, 0xa1, 0xf2},
		}, &rpcv2.OpaqueAuth{Flavor: rpcv2.AUTH_NONE}).
			DoAndReturn(func(ctx context.Context, credentials, verifier *rpcv2.OpaqueAuth) (context.Context, rpcv2.OpaqueAuth, rpcv2.AuthStat) {
				return mockContext1, rpcv2.OpaqueAuth{
					Flavor: rpcv2.AUTH_SHORT,
					Body:   []byte{0x24, 0x3e, 0xa5, 0xeb, 0xa1, 0x91, 0x78, 0x7d},
				}, rpcv2.AUTH_OK
			})
		service.EXPECT().Call(mockContext1, uint32(7), uint32(4), gomock.Any(), gomock.Any()).
			DoAndReturn(func(ctx context.Context, vers, proc uint32, parameters io.ReadCloser, newReturnValue func(int) io.Writer) (rpcv2.AcceptedReplyData, error) {
				var in [8]byte
				n, err := parameters.Read(in[:])
				require.Equal(t, 8, n)
				require.NoError(t, err)
				require.Equal(t, [...]byte{0xb3, 0x83, 0x19, 0x90, 0x0a, 0xe0, 0xf1, 0x2a}, in)
				require.NoError(t, parameters.Close())

				n, err = newReturnValue(8).Write([]byte{0x44, 0xe5, 0x33, 0x30, 0xa5, 0xf4, 0x75, 0xbb})
				require.Equal(t, 8, n)
				require.NoError(t, err)

				return &rpcv2.AcceptedReplyData_SUCCESS{}, nil
			})
		mockContext2 := mock.NewMockContext(ctrl)
		authenticator.EXPECT().Authenticate(gomock.Any(), &rpcv2.OpaqueAuth{Flavor: rpcv2.AUTH_NONE}, &rpcv2.OpaqueAuth{Flavor: rpcv2.AUTH_NONE}).
			DoAndReturn(func(ctx context.Context, credentials, verifier *rpcv2.OpaqueAuth) (context.Context, rpcv2.OpaqueAuth, rpcv2.AuthStat) {
				return mockContext2, rpcv2.OpaqueAuth{Flavor: rpcv2.AUTH_NONE}, rpcv2.AUTH_OK
			})
		service.EXPECT().Call(mockContext2, uint32(3), uint32(9), gomock.Any(), gomock.Any()).
			DoAndReturn(func(ctx context.Context, vers, proc uint32, parameters io.ReadCloser, newReturnValue func(int) io.Writer) (rpcv2.AcceptedReplyData, error) {
				var in [8]byte
				n, err := parameters.Read(in[:])
				require.Equal(t, 8, n)
				require.NoError(t, err)
				require.Equal(t, [...]byte{0xa9, 0x73, 0x1c, 0xfa, 0xfe, 0x16, 0xe0, 0x81}, in)
				require.NoError(t, parameters.Close())

				n, err = newReturnValue(8).Write([]byte{0x26, 0xb5, 0x37, 0xb0, 0xe4, 0xf4, 0x6a, 0x84})
				require.Equal(t, 8, n)
				require.NoError(t, err)

				return &rpcv2.AcceptedReplyData_SUCCESS{}, nil
			})
		w := mock.NewMockWriter(ctrl)
		w.EXPECT().Write([]byte{
			// Record marker.
			0x80, 0x00, 0x00, 0x28,
			// xid.
			0x44, 0xe5, 0x33, 0x30,
			// body.msg_type == REPLY.
			0x00, 0x00, 0x00, 0x01,
			// body.rbody.reply_stat == MSG_ACCEPTED.
			0x00, 0x00, 0x00, 0x00,
			// body.rbody.areply.verf.flavor == AUTH_SHORT.
			0x00, 0x00, 0x00, 0x02,
			// body.rbody.areply.verf.body.
			0x00, 0x00, 0x00, 0x08,
			0x24, 0x3e, 0xa5, 0xeb, 0xa1, 0x91, 0x78, 0x7d,
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
				0x80, 0x00, 0x00, 0x38,
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
				// body.cbody.cred.flavor == AUTH_SHORT,
				0x00, 0x00, 0x00, 0x02,
				// body.cbody.cred.body.
				0x00, 0x00, 0x00, 0x08,
				0xe6, 0xda, 0x1d, 0x8d, 0x27, 0x35, 0xa1, 0xf2,
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

	t.Run("WithoutReturnValue", func(t *testing.T) {
		mismatch := rpcv2.AcceptedReplyData_PROG_MISMATCH{}
		mismatch.MismatchInfo.Low = 3
		mismatch.MismatchInfo.High = 4
		for _, replyData := range []rpcv2.AcceptedReplyData{
			&rpcv2.AcceptedReplyData_SUCCESS{},
			&rpcv2.AcceptedReplyData_default{Stat: rpcv2.PROC_UNAVAIL},
			&mismatch,
		} {
			authenticator.EXPECT().Authenticate(gomock.Any(), &rpcv2.OpaqueAuth{}, &rpcv2.OpaqueAuth{}).
				DoAndReturn(func(ctx context.Context, credentials, verifier *rpcv2.OpaqueAuth) (context.Context, rpcv2.OpaqueAuth, rpcv2.AuthStat) {
					return ctx, rpcv2.OpaqueAuth{}, rpcv2.AUTH_OK
				})
			service.EXPECT().Call(gomock.Any(), uint32(1), uint32(1), gomock.Any(), gomock.Any()).Return(replyData, nil)
			w := mock.NewMockWriter(ctrl)
			w.EXPECT().Write(gomock.Any()).DoAndReturn(func(p []byte) (int, error) {
				require.Equal(t, uint32(len(p)-4)|0x80000000, binary.BigEndian.Uint32(p))
				require.Equal(t, len(p), cap(p))
				var reply rpcv2.RpcMsg
				n, err := reply.ReadFrom(bytes.NewReader(p[4:]))
				require.NoError(t, err)
				require.Equal(t, int64(len(p)-4), n)
				require.Equal(t, replyData, reply.Body.(*rpcv2.RpcMsgBody_REPLY).Rbody.(*rpcv2.ReplyBody_MSG_ACCEPTED).Areply.ReplyData)
				return len(p), nil
			})
			require.NoError(t, s.HandleConnection(bytes.NewReader([]byte{
				0x80, 0x00, 0x00, 0x28, // Record marker.
				0x00, 0x00, 0x00, 0x01, // XID.
				0x00, 0x00, 0x00, 0x00, // CALL.
				0x00, 0x00, 0x00, 0x02, // RPC version.
				0x00, 0x00, 0x00, 0x7b, // Program.
				0x00, 0x00, 0x00, 0x01, // Program version.
				0x00, 0x00, 0x00, 0x01, // Procedure.
				0x00, 0x00, 0x00, 0x00, // AUTH_NONE credentials.
				0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, // AUTH_NONE verifier.
				0x00, 0x00, 0x00, 0x00,
			}), w))
		}
	})

	t.Run("ReplyBodySizing", func(t *testing.T) {
		r := mock.NewMockReader(ctrl)
		w := mock.NewMockWriter(ctrl)
		requestReady := make(chan struct{})
		close(requestReady)
		var expectedCapacities, actualCapacities []int
		sizes := []int{0, 32, 1 << 20, 32, 4095, 4096, 4097, 0}
		authenticator.EXPECT().Authenticate(gomock.Any(), &rpcv2.OpaqueAuth{}, &rpcv2.OpaqueAuth{}).
			DoAndReturn(func(ctx context.Context, credentials, verifier *rpcv2.OpaqueAuth) (context.Context, rpcv2.OpaqueAuth, rpcv2.AuthStat) {
				return ctx, rpcv2.OpaqueAuth{}, rpcv2.AUTH_OK
			}).Times(len(sizes))

		for i, size := range sizes {
			request := []byte{
				0x80, 0x00, 0x00, 0x28, // Record marker.
				0x00, 0x00, 0x00, 0x00, // XID, set below.
				0x00, 0x00, 0x00, 0x00, // CALL.
				0x00, 0x00, 0x00, 0x02, // RPC version.
				0x00, 0x00, 0x00, 0x7b, // Program.
				0x00, 0x00, 0x00, 0x01, // Program version.
				0x00, 0x00, 0x00, 0x01, // Procedure.
				0x00, 0x00, 0x00, 0x00, // AUTH_NONE credentials.
				0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, // AUTH_NONE verifier.
				0x00, 0x00, 0x00, 0x00,
			}
			binary.BigEndian.PutUint32(request[4:], uint32(i+1))
			ready := requestReady
			r.EXPECT().Read(gomock.Any()).DoAndReturn(func(p []byte) (int, error) {
				// Send the next request after the previous reply.
				<-ready
				return copy(p, request), nil
			})

			expectedCapacities = append(expectedCapacities, size)
			var replyStart *byte
			data := bytes.Repeat([]byte{0xa5}, size)
			service.EXPECT().Call(gomock.Any(), uint32(1), uint32(1), gomock.Any(), gomock.Any()).
				DoAndReturn(func(ctx context.Context, vers, proc uint32, parameters io.ReadCloser, newReturnValue func(int) io.Writer) (rpcv2.AcceptedReplyData, error) {
					returnValue := newReturnValue(len(data))
					replyStart = &returnValue.(*bytes.Buffer).Bytes()[0]
					actualCapacities = append(actualCapacities, returnValue.(*bytes.Buffer).Available())
					n, err := returnValue.Write(data)
					require.Equal(t, len(data), n)
					require.NoError(t, err)
					return &rpcv2.AcceptedReplyData_SUCCESS{}, nil
				})

			reply := []byte{
				0x00, 0x00, 0x00, 0x00, // Record marker, set below.
				0x00, 0x00, 0x00, 0x00, // XID, set below.
				0x00, 0x00, 0x00, 0x01, // REPLY.
				0x00, 0x00, 0x00, 0x00, // MSG_ACCEPTED.
				0x00, 0x00, 0x00, 0x00, // AUTH_NONE verifier.
				0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, // SUCCESS.
			}
			binary.BigEndian.PutUint32(reply, uint32(24+len(data))|0x80000000)
			binary.BigEndian.PutUint32(reply[4:], uint32(i+1))
			reply = append(reply, data...)
			nextRequestReady := make(chan struct{})
			w.EXPECT().Write(gomock.Any()).DoAndReturn(func(p []byte) (int, error) {
				defer close(nextRequestReady)
				require.Equal(t, reply, p)
				require.Equal(t, len(p), cap(p))
				require.Same(t, replyStart, &p[0])
				return len(p), nil
			})
			requestReady = nextRequestReady
		}
		r.EXPECT().Read(gomock.Any()).Return(0, io.EOF)

		require.NoError(t, s.HandleConnection(r, w))
		require.Equal(t, expectedCapacities, actualCapacities)
	})
}

func TestNFSv4Service(t *testing.T) {
	ctrl := gomock.NewController(t)
	program := mock.NewMockNfs4Program(ctrl)
	service := nfsv4.NewNfs4ProgramService(program)
	ctx := context.Background()

	t.Run("Compound", func(t *testing.T) {
		for _, size := range []int{0, 1, 1024, 1 << 20} {
			response := nfsv4.Compound4res{
				Tag: "read",
				Resarray: []nfsv4.NfsResop4{
					&nfsv4.NfsResop4_OP_READ{
						Opread: &nfsv4.Read4res_NFS4_OK{
							Resok4: nfsv4.Read4resok{Data: bytes.Repeat([]byte{0xa5}, size)},
						},
					},
				},
			}
			program.EXPECT().NfsV4Nfsproc4Compound(ctx, &nfsv4.Compound4args{}).Return(&response, nil)
			expected := bytes.NewBuffer(nil)
			_, err := response.WriteTo(expected)
			require.NoError(t, err)

			var actual *bytes.Buffer
			replyData, err := service(ctx, 4, 1, io.NopCloser(bytes.NewReader(make([]byte, 12))), func(sizeBytes int) io.Writer {
				require.Nil(t, actual)
				require.Equal(t, expected.Len(), sizeBytes)
				actual = bytes.NewBuffer(make([]byte, 0, sizeBytes))
				return actual
			})
			require.NoError(t, err)
			require.Equal(t, &rpcv2.AcceptedReplyData_SUCCESS{}, replyData)
			require.Equal(t, expected.Bytes(), actual.Bytes())
			require.Equal(t, expected.Len(), actual.Cap())
		}
	})

	t.Run("Void", func(t *testing.T) {
		program.EXPECT().NfsV4Nfsproc4Null(ctx).Return(nil)
		replyData, err := service(ctx, 4, 0, io.NopCloser(bytes.NewReader(nil)), nil)
		require.NoError(t, err)
		require.Equal(t, &rpcv2.AcceptedReplyData_SUCCESS{}, replyData)
	})

	t.Run("ProcedureUnavailable", func(t *testing.T) {
		replyData, err := service(ctx, 4, 2, io.NopCloser(bytes.NewReader(nil)), nil)
		require.NoError(t, err)
		require.Equal(t, &rpcv2.AcceptedReplyData_default{Stat: rpcv2.PROC_UNAVAIL}, replyData)
	})

	t.Run("VersionMismatch", func(t *testing.T) {
		replyData, err := service(ctx, 3, 0, io.NopCloser(bytes.NewReader(nil)), nil)
		require.NoError(t, err)
		expected := rpcv2.AcceptedReplyData_PROG_MISMATCH{}
		expected.MismatchInfo.Low = 4
		expected.MismatchInfo.High = 4
		require.Equal(t, &expected, replyData)
	})

	t.Run("ReadError", func(t *testing.T) {
		replyData, err := service(ctx, 4, 1, io.NopCloser(bytes.NewReader(nil)), nil)
		require.Equal(t, io.EOF, err)
		require.Nil(t, replyData)
	})

	t.Run("ProcedureError", func(t *testing.T) {
		program.EXPECT().NfsV4Nfsproc4Compound(ctx, &nfsv4.Compound4args{}).Return(nil, io.ErrUnexpectedEOF)
		replyData, err := service(ctx, 4, 1, io.NopCloser(bytes.NewReader(make([]byte, 12))), nil)
		require.Equal(t, io.ErrUnexpectedEOF, err)
		require.Nil(t, replyData)
	})

	t.Run("WriteError", func(t *testing.T) {
		program.EXPECT().NfsV4Nfsproc4Compound(ctx, &nfsv4.Compound4args{}).Return(&nfsv4.Compound4res{}, nil)
		w := mock.NewMockWriter(ctrl)
		w.EXPECT().Write(gomock.Any()).Return(0, io.ErrClosedPipe)
		replyData, err := service(ctx, 4, 1, io.NopCloser(bytes.NewReader(make([]byte, 12))), func(sizeBytes int) io.Writer {
			require.Equal(t, 12, sizeBytes)
			return w
		})
		require.Equal(t, io.ErrClosedPipe, err)
		require.Nil(t, replyData)
	})
}
