package model

import (
	"errors"
	"math/big"
)

type variableLengthOpaqueType struct {
	maximumSizeBytes Value
}

// NewVariableLengthOpaqueType creates a variable-length opaque data
// type, as described in RFC 4506, section 4.10. It translates to a byte
// slice in Go.
func NewVariableLengthOpaqueType(maximumSizeBytes Value) Type {
	return &variableLengthOpaqueType{
		maximumSizeBytes: maximumSizeBytes,
	}
}

func (t *variableLengthOpaqueType) getEmittingValueVisitor(s sink, r Resolver) (valueVisitor, error) {
	return nil, errors.New("Opaque constants are not supported")
}

func (t *variableLengthOpaqueType) emitDeclaration(s sink, r Resolver, i identifier) error {
	s.emitString("[]byte")
	return nil
}

func (t *variableLengthOpaqueType) emitAllDeclarations(s sink, r Resolver, i identifier) error {
	return emitDeclarationsStructural(s, r, i, t, true, true)
}

func (t *variableLengthOpaqueType) emitReadFrom(s sink, r Resolver, i identifier) error {
	s.emitString("m, nField, err = ")
	s.emitPackageNamePrefix(xdrRuntimePackage)
	s.emitString("ReadVariableLengthOpaque(r, ")
	if err := t.maximumSizeBytes.visit(newIntegerEmittingValueVisitor(s, r)); err != nil {
		return err
	}
	s.emitString(")\n")
	emitReadWriteErrorHandling(s)
	return nil
}

func (t *variableLengthOpaqueType) emitWriteTo(s sink, r Resolver, i identifier) error {
	s.emitString("nField, err = ")
	s.emitPackageNamePrefix(xdrRuntimePackage)
	s.emitString("WriteVariableLengthOpaque(w, ")
	if err := t.maximumSizeBytes.visit(newIntegerEmittingValueVisitor(s, r)); err != nil {
		return err
	}
	s.emitString(", m)\n")
	emitReadWriteErrorHandling(s)
	return nil
}

func (t *variableLengthOpaqueType) emitGetVariableEncodedSizeBytes(s sink, r Resolver, i identifier) error {
	s.emitString("nTotal += (len(m) + 7) &^ 3\n")
	return nil
}

func (t *variableLengthOpaqueType) isComplexToReadOrWrite(r Resolver) (bool, error) {
	return false, nil
}

func (t *variableLengthOpaqueType) isLarge(r Resolver) (bool, error) {
	return false, nil
}

func (t *variableLengthOpaqueType) getFixedEncodedSizeBytes(r Resolver) (*big.Int, error) {
	return nil, nil
}
