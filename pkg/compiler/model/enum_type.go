package model

import (
	"errors"
	"fmt"
	"math/big"
	"sort"
)

type enumEmittingValueVisitor struct {
	s        sink
	r        Resolver
	elements map[string]Value
}

func (vv *enumEmittingValueVisitor) visitConstant(constant *big.Int) error {
	return errors.New("Enumerations can only be compared against enumeration values")
}

func (vv *enumEmittingValueVisitor) visitIdentifier(name string) error {
	v, ok := vv.elements[name]
	if !ok {
		return fmt.Errorf("%#v is not a valid enumeration value", name)
	}
	return v.visit(newIntegerEmittingValueVisitor(vv.s, vv.r))
}

type enumType struct {
	elements map[string]Value
	names    []string
}

// NewEnumType creates an enumeration type, as described in RFC 4506,
// section 4.3. It translates to a type definition of an integer in Go.
// A constant is declared for every enumeration value.
func NewEnumType(elements map[string]Value) Type {
	names := make([]string, 0, len(elements))
	for name := range elements {
		names = append(names, name)
	}
	sort.Strings(names)

	return &enumType{
		elements: elements,
		names:    names,
	}
}

func (t *enumType) getEmittingValueVisitor(s sink, r Resolver) (valueVisitor, error) {
	return &enumEmittingValueVisitor{
		s:        s,
		r:        r,
		elements: t.elements,
	}, nil
}

func (t *enumType) emitDeclaration(s sink, r Resolver, i identifier) error {
	i.emit(s, false)
	return nil
}

func (t *enumType) emitAllDeclarations(s sink, r Resolver, i identifier) error {
	s.emitString("\n")
	s.emitString("type ")
	i.emit(s, false)
	s.emitString(" int32\n")

	if err := emitAdditionalDeclarationsNominal(s, r, i, t); err != nil {
		return err
	}

	vv := newIntegerEmittingValueVisitor(s, r)
	for _, name := range t.names {
		s.emitString("\n")
		s.emitString("const ")
		emitMacroCase(s, name)
		s.emitString(" ")
		i.emit(s, false)
		s.emitString(" = ")
		if err := t.elements[name].visit(vv); err != nil {
			return err
		}
		s.emitString("\n")
	}

	s.emitString("\nvar ")
	i.emit(s, false)
	s.emitString("_name = map[")
	i.emit(s, false)
	s.emitString("]string{\n")
	for _, name := range t.names {
		if err := t.elements[name].visit(vv); err != nil {
			return err
		}
		s.emitString(": \"")
		s.emitString(name)
		s.emitString("\",\n")
	}
	s.emitString("}\n")
	return nil
}

func (t *enumType) emitReadFrom(s sink, r Resolver, i identifier) error {
	s.emitString("*(*int32)(&m), nField, err = ")
	s.emitPackageNamePrefix(xdrRuntimePackage)
	s.emitString("ReadInt(r)\n")
	emitReadWriteErrorHandling(s)
	return nil
}

func (t *enumType) emitWriteTo(s sink, r Resolver, i identifier) error {
	s.emitString("nField, err = ")
	s.emitPackageNamePrefix(xdrRuntimePackage)
	s.emitString("WriteInt(w, int32(m))\n")
	emitReadWriteErrorHandling(s)
	return nil
}

func (t *enumType) emitGetVariableEncodedSizeBytes(s sink, r Resolver, i identifier) error {
	panic("Enums have a fixed size")
}

func (t *enumType) isLarge(r Resolver) (bool, error) {
	return false, nil
}

var enumFixedEncodingSizeBytes = big.NewInt(4)

func (t *enumType) getFixedEncodedSizeBytes(r Resolver) (*big.Int, error) {
	return enumFixedEncodingSizeBytes, nil
}
