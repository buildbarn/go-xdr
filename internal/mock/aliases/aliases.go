package aliases

import (
	"context"
	"io"
)

// This file contains aliases for some of the interfaces provided by the
// Go standard library. The only reason this file exists is to allow the
// gomock() Bazel rule to emit mocks for them, as that rule is only
// capable of emitting mocks for interfaces built through a
// go_library().

// Context is an alias of context.Context.
type Context = context.Context

// Reader is an alias of io.Reader.
type Reader = io.Reader

// Writer is an alias of io.Writer.
type Writer = io.Writer
