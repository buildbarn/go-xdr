package model

import (
	"errors"
	"math/big"
)

type structType struct {
	declarations []NamedDeclaration
}

// NewStructType creates a structure type, as described in RFC 4506,
// section 4.14. It translates to a struct in Go.
func NewStructType(declarations []NamedDeclaration) Type {
	return &structType{
		declarations: declarations,
	}
}

func (t *structType) getEmittingValueVisitor(s sink, r Resolver) (valueVisitor, error) {
	return nil, errors.New("Struct constants are not supported")
}

func (t *structType) emitDeclaration(s sink, r Resolver, i identifier) error {
	if i.hasTypeDefinition(true) {
		i.emit(s, false)
	} else {
		s.emitString("struct {\n")
		for _, nd := range t.declarations {
			if err := nd.emitStructField(s, r, i.append(false, false)); err != nil {
				return err
			}
		}
		s.emitString("}")
	}
	return nil
}

func (t *structType) emitAllDeclarations(s sink, r Resolver, i identifier) error {
	if i.hasTypeDefinition(true) {
		s.emitString("\n")
		s.emitString("type ")
		i.emit(s, false)
		s.emitString(" struct {\n")
		for _, nd := range t.declarations {
			if err := nd.emitStructField(s, r, i.append(false, false)); err != nil {
				return err
			}
		}
		s.emitString("}\n")
	}

	if err := emitAdditionalDeclarationsNominal(s, r, i, t); err != nil {
		return err
	}

	for _, nd := range t.declarations {
		if err := nd.emitAllDeclarations(s, r, i.append(false, false)); err != nil {
			return err
		}
	}
	return nil
}

func (t *structType) emitReadFrom(s sink, r Resolver, i identifier) error {
	if i.hasReadFromWriteTo() {
		emitForwardReadFromNominal(s)
	} else {
		for _, nd := range t.declarations {
			if err := nd.emitStructFieldReadFrom(s, r, i.append(false, false)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (t *structType) emitWriteTo(s sink, r Resolver, i identifier) error {
	if i.hasReadFromWriteTo() {
		emitForwardWriteToNominal(s)
	} else {
		for _, nd := range t.declarations {
			if err := nd.emitStructFieldWriteTo(s, r, i.append(false, false)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (t *structType) emitGetVariableEncodedSizeBytes(s sink, r Resolver, i identifier) error {
	if i.hasReadFromWriteTo() {
		emitForwardGetVariableEncodedSizeBytesNominal(s)
	} else {
		for _, nd := range t.declarations {
			if err := nd.maybeEmitStructFieldGetVariableEncodedSizeBytes(s, r, i.append(false, false)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (t *structType) isLarge(r Resolver) (bool, error) {
	return true, nil
}

func (t *structType) getFixedEncodedSizeBytes(r Resolver) (*big.Int, error) {
	var total big.Int
	for _, nd := range t.declarations {
		size, err := nd.declarationType.getFixedEncodedSizeBytes(r)
		if err != nil || size == nil {
			return nil, err
		}
		total.Add(&total, size)
	}
	return &total, nil
}
