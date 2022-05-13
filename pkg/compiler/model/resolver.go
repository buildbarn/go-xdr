package model

import (
	"math/big"
)

// Resolver needs to be provided to TopLevel.Emit() to resolve constants
// and types.
type Resolver interface {
	ResolveConstant(name string) (*big.Int, error)
	ResolveType(name string) (string, Type, error)
}
