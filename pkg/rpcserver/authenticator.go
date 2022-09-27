package rpcserver

import (
	"context"

	"github.com/buildbarn/go-xdr/pkg/protocols/rpcv2"
)

// Authenticator of requests against an RPC server.
type Authenticator interface {
	Authenticate(ctx context.Context, credentials, verifier *rpcv2.OpaqueAuth) (context.Context, rpcv2.OpaqueAuth, rpcv2.AuthStat)
}
