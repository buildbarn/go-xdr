package rpcserver

import (
	"context"

	"github.com/buildbarn/go-xdr/pkg/protocols/rpcv2"
)

type allowAuthenticator struct{}

func (allowAuthenticator) Authenticate(ctx context.Context, credentials, verifier *rpcv2.OpaqueAuth) (context.Context, rpcv2.OpaqueAuth, rpcv2.AuthStat) {
	return ctx, rpcv2.OpaqueAuth{Flavor: rpcv2.AUTH_NONE}, rpcv2.AUTH_OK
}

// AllowAuthenticator is an implementation of Authenticator that permits
// the execution of all requests, regardless of the credentials and
// verifier they provide.
var AllowAuthenticator Authenticator = allowAuthenticator{}
