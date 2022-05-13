package model

import (
	"math/big"
)

// NamedDeclaration corresponds to a declaration of a type under a given
// name. Named declarations are either top-level type definitions, or
// declarations of named fields inside structs or unions.
type NamedDeclaration struct {
	name            string
	declarationType Type
}

// NewNamedDeclaration creates a new named declaration.
func NewNamedDeclaration(name string, declarationType Type) NamedDeclaration {
	return NamedDeclaration{
		name:            name,
		declarationType: declarationType,
	}
}

func (nd NamedDeclaration) emitAllDeclarations(s sink, r Resolver, f identifierFactory) error {
	return nd.declarationType.emitAllDeclarations(s, r, f.newIdentifier(nd.name))
}

func (nd NamedDeclaration) emitStructField(s sink, r Resolver, f identifierFactory) error {
	emitMixedCase(s, nd.name, true)
	s.emitString(" ")
	if err := nd.declarationType.emitDeclaration(s, r, f.newIdentifier(nd.name)); err != nil {
		return err
	}
	s.emitString("\n")
	return nil
}

func (nd NamedDeclaration) emitStructFieldReadFrom(s sink, r Resolver, f identifierFactory) error {
	tField := nd.declarationType
	isLarge, err := tField.isLarge(r)
	if err != nil {
		return err
	}
	iChild := f.newIdentifier(nd.name)

	s.emitString("{\n")
	if isLarge {
		s.emitString("m := &m.")
		emitMixedCase(s, nd.name, true)
		s.emitString("\n")
		if err := tField.emitReadFrom(s, r, iChild); err != nil {
			return err
		}
	} else {
		s.emitString("mSave := &m.")
		emitMixedCase(s, nd.name, true)
		s.emitString("\n")
		s.emitString("var m ")
		if err := tField.emitDeclaration(s, r, iChild); err != nil {
			return err
		}
		s.emitString("\n")
		if err := tField.emitReadFrom(s, r, iChild); err != nil {
			return err
		}
		s.emitString("*mSave = m\n")
	}
	s.emitString("}\n")
	return nil
}

func (nd NamedDeclaration) emitStructFieldWriteTo(s sink, r Resolver, f identifierFactory) error {
	isLarge, err := nd.declarationType.isLarge(r)
	if err != nil {
		return err
	}

	s.emitString("{\n")
	if isLarge {
		s.emitString("m := &m.")
	} else {
		s.emitString("m := m.")
	}
	emitMixedCase(s, nd.name, true)
	s.emitString("\n")
	if err := nd.declarationType.emitWriteTo(s, r, f.newIdentifier(nd.name)); err != nil {
		return err
	}
	s.emitString("}\n")
	return nil
}

func (nd NamedDeclaration) emitStructFieldGetVariableEncodedSizeBytes(s sink, r Resolver, f identifierFactory) error {
	isLarge, err := nd.declarationType.isLarge(r)
	if err != nil {
		return err
	}

	s.emitString("{\n")
	if isLarge {
		s.emitString("m := &m.")
	} else {
		s.emitString("m := m.")
	}
	emitMixedCase(s, nd.name, true)
	s.emitString("\n")
	if err := nd.declarationType.emitGetVariableEncodedSizeBytes(s, r, f.newIdentifier(nd.name)); err != nil {
		return err
	}
	s.emitString("}\n")
	return nil
}

func (nd NamedDeclaration) maybeEmitStructFieldGetVariableEncodedSizeBytes(s sink, r Resolver, f identifierFactory) error {
	size, err := nd.declarationType.getFixedEncodedSizeBytes(r)
	if err != nil {
		return err
	}
	if size == nil {
		return nd.emitStructFieldGetVariableEncodedSizeBytes(s, r, f)
	}
	s.emitString("nTotal += ")
	s.emitString(size.String())
	s.emitString("\n")
	return nil
}

func (nd NamedDeclaration) getFixedEncodedSizeBytes(r Resolver) (*big.Int, error) {
	return nd.declarationType.getFixedEncodedSizeBytes(r)
}
