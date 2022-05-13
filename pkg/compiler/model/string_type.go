package model

import (
	"errors"
	"math/big"
)

type stringType struct {
	maximumSizeBytes Value
	stringType       string
}

// NewASCIIStringType creates a new ASCII string type as described in
// RFC 4506, section 4.11.
func NewASCIIStringType(maximumSizeBytes Value) Type {
	return &stringType{
		maximumSizeBytes: maximumSizeBytes,
		stringType:       "ASCII",
	}
}

// NewUTF8StringType creates a new UTF-8 string type. This is an
// extension to RFC 4506, which only supports ASCII strings.
func NewUTF8StringType(maximumSizeBytes Value) Type {
	return &stringType{
		maximumSizeBytes: maximumSizeBytes,
		stringType:       "UTF8",
	}
}

func (t *stringType) getEmittingValueVisitor(s sink, r Resolver) (valueVisitor, error) {
	return nil, errors.New("String constants are not supported")
}

func (t *stringType) emitDeclaration(s sink, r Resolver, i identifier) error {
	s.emitString("string")
	return nil
}

func (t *stringType) emitAllDeclarations(s sink, r Resolver, i identifier) error {
	return emitDeclarationsStructural(s, r, i, t, true, true)
}

func (t *stringType) emitReadFrom(s sink, r Resolver, i identifier) error {
	s.emitString("m, nField, err = ")
	s.emitPackageNamePrefix(xdrRuntimePackage)
	s.emitString("Read")
	s.emitString(t.stringType)
	s.emitString("String(r, ")
	if err := t.maximumSizeBytes.visit(newIntegerEmittingValueVisitor(s, r)); err != nil {
		return err
	}
	s.emitString(")\n")
	emitReadWriteErrorHandling(s)
	return nil
}

func (t *stringType) emitWriteTo(s sink, r Resolver, i identifier) error {
	s.emitString("nField, err = ")
	s.emitPackageNamePrefix(xdrRuntimePackage)
	s.emitString("Write")
	s.emitString(t.stringType)
	s.emitString("String(w, ")
	if err := t.maximumSizeBytes.visit(newIntegerEmittingValueVisitor(s, r)); err != nil {
		return err
	}
	s.emitString(", m)\n")
	emitReadWriteErrorHandling(s)
	return nil
}

func (t *stringType) emitGetVariableEncodedSizeBytes(s sink, r Resolver, i identifier) error {
	s.emitString("nTotal += (len(m) + 7) &^ 3\n")
	return nil
}

func (t *stringType) isLarge(r Resolver) (bool, error) {
	return false, nil
}

func (t *stringType) getFixedEncodedSizeBytes(r Resolver) (*big.Int, error) {
	return nil, nil
}
