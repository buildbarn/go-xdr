package model

import (
	"math/big"
)

type integerType struct {
	nativeType            string
	xdrType               string
	fixedEncodedSizeBytes *big.Int
}

func (t *integerType) getEmittingValueVisitor(s sink, r Resolver) (valueVisitor, error) {
	return newIntegerEmittingValueVisitor(s, r), nil
}

func (t *integerType) emitDeclaration(s sink, r Resolver, i identifier) error {
	s.emitString(t.nativeType)
	return nil
}

func (t *integerType) emitAllDeclarations(s sink, r Resolver, i identifier) error {
	return emitDeclarationsStructural(s, r, i, t, true, true)
}

func (t *integerType) emitReadFrom(s sink, r Resolver, i identifier) error {
	s.emitString("m, nField, err = ")
	s.emitPackageNamePrefix(xdrRuntimePackage)
	s.emitString("Read")
	s.emitString(t.xdrType)
	s.emitString("(r)\n")
	emitReadWriteErrorHandling(s)
	return nil
}

func (t *integerType) emitWriteTo(s sink, r Resolver, i identifier) error {
	s.emitString("nField, err = ")
	s.emitPackageNamePrefix(xdrRuntimePackage)
	s.emitString("Write")
	s.emitString(t.xdrType)
	s.emitString("(w, m)\n")
	emitReadWriteErrorHandling(s)
	return nil
}

func (t *integerType) emitGetVariableEncodedSizeBytes(s sink, r Resolver, i identifier) error {
	panic("Integers have a fixed size")
}

func (t *integerType) isLarge(r Resolver) (bool, error) {
	return false, nil
}

func (t *integerType) getFixedEncodedSizeBytes(r Resolver) (*big.Int, error) {
	return t.fixedEncodedSizeBytes, nil
}

// IntType corresponds to a 32-bit signed integer type as described in
// RFC 4506, section 4.1.
var IntType Type = &integerType{
	nativeType:            "int32",
	xdrType:               "Int",
	fixedEncodedSizeBytes: big.NewInt(4),
}

// UnsignedIntType corresponds to a 32-bit unsigned integer type as
// described in RFC 4506, section 4.2.
var UnsignedIntType Type = &integerType{
	nativeType:            "uint32",
	xdrType:               "UnsignedInt",
	fixedEncodedSizeBytes: big.NewInt(4),
}

// HyperType corresponds to a 64-bit signed integer type as described in
// RFC 4506, section 4.5.
var HyperType Type = &integerType{
	nativeType:            "int64",
	xdrType:               "Hyper",
	fixedEncodedSizeBytes: big.NewInt(8),
}

// UnsignedHyperType corresponds to a 64-bit unsigned integer type as
// described in RFC 4506, section 4.5.
var UnsignedHyperType Type = &integerType{
	nativeType:            "uint64",
	xdrType:               "UnsignedHyper",
	fixedEncodedSizeBytes: big.NewInt(8),
}
