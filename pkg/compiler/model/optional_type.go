package model

import (
	"errors"
	"math/big"
)

type optionalType struct {
	base Type
}

// NewOptionalType creates an optional-data type as described in RFC
// 4506, section 4.19. It translates to a pointer type in Go.
func NewOptionalType(base Type) Type {
	return &optionalType{
		base: base,
	}
}

func (t *optionalType) getEmittingValueVisitor(s sink, r Resolver) (valueVisitor, error) {
	return nil, errors.New("Optional constants are not supported")
}

func (t *optionalType) emitDeclaration(s sink, r Resolver, i identifier) error {
	s.emitString("*")
	return t.base.emitDeclaration(s, r, i.taint())
}

func (t *optionalType) emitAllDeclarations(s sink, r Resolver, i identifier) error {
	if err := emitDeclarationsStructural(s, r, i, t, true, true); err != nil {
		return err
	}
	return t.base.emitAllDeclarations(s, r, i.taint())
}

func (t *optionalType) emitReadFrom(s sink, r Resolver, i identifier) error {
	childNP := i.taint()

	s.emitString("var isSet bool\n")
	s.emitString("isSet, nField, err = ")
	s.emitPackageNamePrefix(xdrRuntimePackage)
	s.emitString("ReadBool(r)\n")
	emitReadWriteErrorHandling(s)
	s.emitString("if isSet {\n")
	s.emitString("mParent := &m\n")
	s.emitString("var m ")
	if err := t.base.emitDeclaration(s, r, childNP); err != nil {
		return err
	}
	s.emitString("\n")
	if err := t.base.emitReadFrom(s, r, childNP); err != nil {
		return err
	}
	s.emitString("*mParent = &m\n")
	s.emitString("}\n")
	return nil
}

func (t *optionalType) emitWriteTo(s sink, r Resolver, i identifier) error {
	isLarge, err := t.base.isLarge(r)
	if err != nil {
		return err
	}

	s.emitString("if m == nil {\n")
	s.emitString("nField, err = ")
	s.emitPackageNamePrefix(xdrRuntimePackage)
	s.emitString("WriteBool(w, false)\n")
	emitReadWriteErrorHandling(s)
	s.emitString("} else {\n")
	s.emitString("nField, err = ")
	s.emitPackageNamePrefix(xdrRuntimePackage)
	s.emitString("WriteBool(w, true)\n")
	emitReadWriteErrorHandling(s)
	s.emitString("{\n")
	if !isLarge {
		s.emitString("m := *m\n")
	}
	if err := t.base.emitWriteTo(s, r, i.taint()); err != nil {
		return err
	}
	s.emitString("}\n")
	s.emitString("}\n")
	return nil
}

func (t *optionalType) emitGetVariableEncodedSizeBytes(s sink, r Resolver, i identifier) error {
	s.emitString("nTotal += 4\n")
	s.emitString("if m != nil {\n")
	size, err := t.base.getFixedEncodedSizeBytes(r)
	if err != nil {
		return err
	} else if size != nil {
		s.emitString("nTotal += ")
		s.emitString(size.String())
		s.emitString("\n")
	} else {
		isLarge, err := t.base.isLarge(r)
		if err != nil {
			return err
		}
		if !isLarge {
			s.emitString("m := *m\n")
		}
		if err := t.base.emitGetVariableEncodedSizeBytes(s, r, i.taint()); err != nil {
			return err
		}
	}
	s.emitString("}\n")
	return nil
}

func (t *optionalType) isLarge(r Resolver) (bool, error) {
	return false, nil
}

func (t *optionalType) getFixedEncodedSizeBytes(r Resolver) (*big.Int, error) {
	return nil, nil
}
