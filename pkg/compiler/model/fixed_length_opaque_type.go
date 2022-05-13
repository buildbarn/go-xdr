package model

import (
	"errors"
	"math/big"
)

type fixedLengthOpaqueType struct {
	sizeBytes Value
}

// NewFixedLengthOpaqueType creates a fixed-length opaque data type, as
// described in RFC 4506, section 4.9. It translates to a fixed-length byte
// array in Go.
func NewFixedLengthOpaqueType(sizeBytes Value) Type {
	return &fixedLengthOpaqueType{
		sizeBytes: sizeBytes,
	}
}

func (t *fixedLengthOpaqueType) getEmittingValueVisitor(s sink, r Resolver) (valueVisitor, error) {
	return nil, errors.New("Opaque constants are not supported")
}

func (t *fixedLengthOpaqueType) emitDeclaration(s sink, r Resolver, i identifier) error {
	s.emitString("[")
	if err := t.sizeBytes.visit(newIntegerEmittingValueVisitor(s, r)); err != nil {
		return err
	}
	s.emitString("]byte")
	return nil
}

func (t *fixedLengthOpaqueType) emitAllDeclarations(s sink, r Resolver, i identifier) error {
	return emitDeclarationsStructural(s, r, i, t, true, true)
}

func (t *fixedLengthOpaqueType) emitReadFrom(s sink, r Resolver, i identifier) error {
	s.emitString("nField, err = ")
	s.emitPackageNamePrefix(xdrRuntimePackage)
	s.emitString("ReadFixedLengthOpaque(r, m[:])\n")
	emitReadWriteErrorHandling(s)
	return nil
}

func (t *fixedLengthOpaqueType) emitWriteTo(s sink, r Resolver, i identifier) error {
	s.emitString("nField, err = ")
	s.emitPackageNamePrefix(xdrRuntimePackage)
	s.emitString("WriteFixedLengthOpaque(w, m[:])\n")
	emitReadWriteErrorHandling(s)
	return nil
}

func (t *fixedLengthOpaqueType) emitGetVariableEncodedSizeBytes(s sink, r Resolver, i identifier) error {
	panic("Fixed length opaque types have a fixed size")
}

func (t *fixedLengthOpaqueType) isLarge(r Resolver) (bool, error) {
	return true, nil
}

var paddingMask = big.NewInt(0x3)

func (t *fixedLengthOpaqueType) getFixedEncodedSizeBytes(r Resolver) (*big.Int, error) {
	vv := newConstantResolvingValueVisitor(r)
	if err := t.sizeBytes.visit(vv); err != nil {
		return nil, err
	}
	var withPadding big.Int
	withPadding.Add(vv.constant, paddingMask)
	return withPadding.AndNot(vv.constant, paddingMask), nil
}
