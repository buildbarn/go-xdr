package model

import (
	"math/big"
)

// Declaration of a named type, or void.
type Declaration interface {
	emitAllDeclarations(s sink, r Resolver, f identifierFactory) error
	emitStructField(s sink, r Resolver, f identifierFactory) error
	emitStructFieldReadFrom(s sink, r Resolver, f identifierFactory) error
	emitStructFieldWriteTo(s sink, r Resolver, f identifierFactory) error
	maybeEmitStructFieldGetVariableEncodedSizeBytes(s sink, r Resolver, f identifierFactory) error
	getFixedEncodedSizeBytes(r Resolver) (*big.Int, error)
}
