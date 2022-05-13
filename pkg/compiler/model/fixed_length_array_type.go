package model

import (
	"errors"
	"math/big"
)

type fixedLengthArrayType struct {
	specifier    Type
	sizeElements Value
}

// NewFixedLengthArrayType creates a type that corresponds to a
// fixed-length array, as described in RFC 4506, section 4.12. It
// translates to a fixed-length array in Go.
func NewFixedLengthArrayType(specifier Type, sizeElements Value) Type {
	return &fixedLengthArrayType{
		specifier:    specifier,
		sizeElements: sizeElements,
	}
}

func (t *fixedLengthArrayType) getEmittingValueVisitor(s sink, r Resolver) (valueVisitor, error) {
	return nil, errors.New("Array constants are not supported")
}

func (t *fixedLengthArrayType) emitDeclaration(s sink, r Resolver, i identifier) error {
	s.emitString("[")
	if err := t.sizeElements.visit(newIntegerEmittingValueVisitor(s, r)); err != nil {
		return err
	}
	s.emitString("]")
	return t.specifier.emitDeclaration(s, r, i.taint())
}

func (t *fixedLengthArrayType) emitAllDeclarations(s sink, r Resolver, i identifier) error {
	if err := emitDeclarationsStructural(s, r, i, t, true, true); err != nil {
		return err
	}
	return t.specifier.emitAllDeclarations(s, r, i.taint())
}

func (t *fixedLengthArrayType) emitReadFrom(s sink, r Resolver, i identifier) error {
	if i.hasReadFromWriteTo() {
		return emitForwardReadFromStructural(s, r, i, true)
	}

	iChild := i.taint()
	s.emitString("mParent := m\n")
	s.emitString("for i := 0; i < len(m); i++ {\n")
	s.emitString("var m ")
	if err := t.specifier.emitDeclaration(s, r, iChild); err != nil {
	}
	s.emitString("\n")
	if err := t.specifier.emitReadFrom(s, r, iChild); err != nil {
		return err
	}
	s.emitString("mParent[i] = m\n")
	s.emitString("}\n")
	return nil
}

func (t *fixedLengthArrayType) emitWriteTo(s sink, r Resolver, i identifier) error {
	if i.hasReadFromWriteTo() {
		emitForwardWriteToStructural(s, i)
	} else {
		isLarge, err := t.specifier.isLarge(r)
		if err != nil {
			return err
		}

		if isLarge {
			s.emitString("for i := 0; i < len(m); i++ {\n")
			s.emitString("m := &m[i]\n")
		} else {
			s.emitString("for _, m := range m {\n")
		}
		if err := t.specifier.emitWriteTo(s, r, i.taint()); err != nil {
			return err
		}
		s.emitString("}\n")
	}
	return nil
}

func (t *fixedLengthArrayType) emitGetVariableEncodedSizeBytes(s sink, r Resolver, i identifier) error {
	panic("TODO")
}

func (t *fixedLengthArrayType) isLarge(r Resolver) (bool, error) {
	return true, nil
}

func (t *fixedLengthArrayType) getFixedEncodedSizeBytes(r Resolver) (*big.Int, error) {
	elementSizeBytes, err := t.specifier.getFixedEncodedSizeBytes(r)
	if err != nil || elementSizeBytes == nil {
		return elementSizeBytes, err
	}
	vv := newConstantResolvingValueVisitor(r)
	if err := t.sizeElements.visit(vv); err != nil {
		return nil, err
	}
	var size big.Int
	return size.Mul(elementSizeBytes, vv.constant), nil
}
