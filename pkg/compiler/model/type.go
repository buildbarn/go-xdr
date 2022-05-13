package model

import (
	"math/big"
)

// Type that is used as part of a type or program definition.
type Type interface {
	getEmittingValueVisitor(s sink, r Resolver) (valueVisitor, error)
	emitDeclaration(s sink, r Resolver, i identifier) error
	emitAllDeclarations(s sink, r Resolver, i identifier) error
	emitReadFrom(s sink, r Resolver, i identifier) error
	emitWriteTo(s sink, r Resolver, i identifier) error
	emitGetVariableEncodedSizeBytes(s sink, r Resolver, i identifier) error
	isLarge(r Resolver) (bool, error)
	getFixedEncodedSizeBytes(r Resolver) (*big.Int, error)
}
