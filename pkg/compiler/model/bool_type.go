package model

import (
	"errors"
	"math/big"
)

type emittingBoolValueVisitor struct {
	s sink
}

func (vv *emittingBoolValueVisitor) visitConstant(constant *big.Int) error {
	return errors.New("Booleans can only be compared against identifiers FALSE and TRUE")
}

func (vv *emittingBoolValueVisitor) visitIdentifier(name string) error {
	switch name {
	case "FALSE":
		vv.s.emitString("false")
		return nil
	case "TRUE":
		vv.s.emitString("true")
		return nil
	default:
		return errors.New("Booleans can only be compared against identifiers FALSE and TRUE")
	}
}

type boolType struct {
	sizeBits int
}

func (t *boolType) getEmittingValueVisitor(s sink, r Resolver) (valueVisitor, error) {
	return &emittingBoolValueVisitor{s: s}, nil
}

func (t *boolType) emitDeclaration(s sink, r Resolver, i identifier) error {
	s.emitString("bool")
	return nil
}

func (t *boolType) emitAllDeclarations(s sink, r Resolver, i identifier) error {
	return emitDeclarationsStructural(s, r, i, t, true, true)
}

func (t *boolType) emitReadFrom(s sink, r Resolver, i identifier) error {
	s.emitString("m, nField, err = ")
	s.emitPackageNamePrefix(xdrRuntimePackage)
	s.emitString("ReadBool(r)\n")
	emitReadWriteErrorHandling(s)
	return nil
}

func (t *boolType) emitWriteTo(s sink, r Resolver, i identifier) error {
	s.emitString("nField, err = ")
	s.emitPackageNamePrefix(xdrRuntimePackage)
	s.emitString("WriteBool(w, m)\n")
	emitReadWriteErrorHandling(s)
	return nil
}

func (t *boolType) emitGetVariableEncodedSizeBytes(s sink, r Resolver, i identifier) error {
	panic("Booleans have a fixed size")
}

func (t *boolType) isLarge(r Resolver) (bool, error) {
	return false, nil
}

var boolFixedEncodingSizeBytes = big.NewInt(4)

func (t *boolType) getFixedEncodedSizeBytes(r Resolver) (*big.Int, error) {
	return boolFixedEncodingSizeBytes, nil
}

// BoolType corresponds to the boolean type as described in RFC 4506,
// section 4.4. It gets translated to Go's 'bool' type.
var BoolType Type = &boolType{}
