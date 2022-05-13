package model

import (
	"math/big"
)

type voidDeclaration struct{}

func (d voidDeclaration) emitAllDeclarations(s sink, r Resolver, f identifierFactory) error {
	return nil
}

func (d voidDeclaration) emitStructField(s sink, r Resolver, f identifierFactory) error {
	return nil
}

func (d voidDeclaration) emitStructFieldReadFrom(s sink, r Resolver, f identifierFactory) error {
	s.emitString("_ = m\n")
	return nil
}

func (d voidDeclaration) emitStructFieldWriteTo(s sink, r Resolver, f identifierFactory) error {
	return nil
}

func (d voidDeclaration) maybeEmitStructFieldGetVariableEncodedSizeBytes(s sink, r Resolver, f identifierFactory) error {
	return nil
}

var voidFixedEncodingSizeBytes = big.NewInt(0)

func (d voidDeclaration) getFixedEncodedSizeBytes(r Resolver) (*big.Int, error) {
	return voidFixedEncodingSizeBytes, nil
}

// VoidDeclaration is declaration of type void, as described in RFC
// 4506, section 4.16. It is typically only used inside discriminated
// unions to denote arms that do not accept a value.
var VoidDeclaration Declaration = voidDeclaration{}
