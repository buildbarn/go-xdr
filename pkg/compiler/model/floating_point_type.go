package model

import (
	"errors"
	"math/big"
	"strconv"
)

type floatingPointType struct {
	sizeBits              int
	xdrType               string
	fixedEncodedSizeBytes *big.Int
}

func (t *floatingPointType) getEmittingValueVisitor(s sink, r Resolver) (valueVisitor, error) {
	return nil, errors.New("Floating point constants are not supported")
}

func (t *floatingPointType) emitDeclaration(s sink, r Resolver, i identifier) error {
	s.emitString("float")
	s.emitString(strconv.FormatInt(int64(t.sizeBits), 10))
	return nil
}

func (t *floatingPointType) emitAllDeclarations(s sink, r Resolver, i identifier) error {
	return emitDeclarationsStructural(s, r, i, t, true, true)
}

func (t *floatingPointType) emitReadFrom(s sink, r Resolver, i identifier) error {
	s.emitString("m, nField, err = ")
	s.emitPackageNamePrefix(xdrRuntimePackage)
	s.emitString("Read")
	s.emitString(t.xdrType)
	s.emitString("(r)\n")
	emitReadWriteErrorHandling(s)
	return nil
}

func (t *floatingPointType) emitWriteTo(s sink, r Resolver, i identifier) error {
	s.emitString("nField, err = ")
	s.emitPackageNamePrefix(xdrRuntimePackage)
	s.emitString("Write")
	s.emitString(t.xdrType)
	s.emitString("(w, m)\n")
	emitReadWriteErrorHandling(s)
	return nil
}

func (t *floatingPointType) emitGetVariableEncodedSizeBytes(s sink, r Resolver, i identifier) error {
	panic("Floating points have a fixed size")
}

func (t *floatingPointType) isLarge(r Resolver) (bool, error) {
	return false, nil
}

func (t *floatingPointType) getFixedEncodedSizeBytes(r Resolver) (*big.Int, error) {
	return t.fixedEncodedSizeBytes, nil
}

// FloatType corresponds to a single-precision floating-point type as
// described in RFC 4506, section 4.6.
var FloatType Type = &floatingPointType{
	sizeBits:              32,
	xdrType:               "Float",
	fixedEncodedSizeBytes: big.NewInt(4),
}

// DoubleType corresponds to a double-precision floating-point type as
// described in RFC 4506, section 4.7.
var DoubleType Type = &floatingPointType{
	sizeBits:              64,
	xdrType:               "Double",
	fixedEncodedSizeBytes: big.NewInt(8),
}

// QuadrupleType corresponds to a quadruple-precision floating-point
// type as described in RFC 4506, section 4.8.
var QuadrupleType Type = &floatingPointType{
	sizeBits:              128,
	xdrType:               "Quadruple",
	fixedEncodedSizeBytes: big.NewInt(16),
}
