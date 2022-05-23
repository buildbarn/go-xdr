package model

import (
	"errors"
	"math/big"
)

type variableLengthArrayType struct {
	base                Type
	maximumSizeElements Value
}

// NewVariableLengthArrayType creates a type that corresponds to a
// variable-length array, as described in RFC 4506, section 4.13. It
// translates to a slice in Go.
func NewVariableLengthArrayType(base Type, maximumSizeElements Value) Type {
	return &variableLengthArrayType{
		base:                base,
		maximumSizeElements: maximumSizeElements,
	}
}

func (t *variableLengthArrayType) getEmittingValueVisitor(s sink, r Resolver) (valueVisitor, error) {
	return nil, errors.New("Array constants are not supported")
}

func (t *variableLengthArrayType) emitDeclaration(s sink, r Resolver, i identifier) error {
	s.emitString("[]")
	return t.base.emitDeclaration(s, r, i.taint())
}

func (t *variableLengthArrayType) emitAllDeclarations(s sink, r Resolver, i identifier) error {
	if err := emitDeclarationsStructural(s, r, i, t, true, true); err != nil {
		return err
	}
	return t.base.emitAllDeclarations(s, r, i.taint())
}

func (t *variableLengthArrayType) emitReadFrom(s sink, r Resolver, i identifier) error {
	childNP := i.taint()

	// Read the number of elements in the array.
	s.emitString("var nElements uint32\n")
	s.emitString("nElements, nField, err = ")
	s.emitPackageNamePrefix(xdrRuntimePackage)
	s.emitString("ReadUnsignedInt(r)\n")
	emitReadWriteErrorHandling(s)
	s.emitString("if nElements > ")
	if err := t.maximumSizeElements.visit(newIntegerEmittingValueVisitor(s, r)); err != nil {
		return err
	}
	s.emitString(" {\nerr = ")
	s.emitPackageNamePrefix("fmt")
	s.emitString("Errorf(\"size of %d elements exceeds ")
	i.emitOriginalName(s)
	s.emitString("'s maximum of ")
	if err := t.maximumSizeElements.visit(newIntegerEmittingValueVisitor(s, r)); err != nil {
		return err
	}
	s.emitString(" elements\", nElements)\ngoto done\n}\n")

	isLarge, err := t.base.isLarge(r)
	if err != nil {
		return err
	}

	// Read the elements in the array. Ideally we would call
	// make([]T, nElements) here, but this has the downside that the
	// caller of ReadFrom() cannot limit memory usage through
	// io.LimitReader.
	s.emitString("for nElements > 0 {\n")
	s.emitString("nElements--\n")
	if isLarge {
		s.emitString("m = append(m, ")
		if err := t.base.emitDeclaration(s, r, childNP); err != nil {
			return err
		}
		s.emitString("{})\n")
		s.emitString("m := &m[len(m)-1]\n")
		if err := t.base.emitReadFrom(s, r, childNP); err != nil {
			return err
		}
	} else {
		s.emitString("mParent := &m\n")
		s.emitString("var m ")
		if err := t.base.emitDeclaration(s, r, childNP); err != nil {
			return err
		}
		s.emitString("\n")
		if err := t.base.emitReadFrom(s, r, childNP); err != nil {
			return err
		}
		s.emitString("*mParent = append(*mParent, m)\n")
	}
	s.emitString("}\n")
	return nil
}

func (t *variableLengthArrayType) emitWriteTo(s sink, r Resolver, i identifier) error {
	s.emitString("if uint(len(m)) > ")
	if err := t.maximumSizeElements.visit(newIntegerEmittingValueVisitor(s, r)); err != nil {
		return err
	}
	s.emitString(" {\nerr = ")
	s.emitPackageNamePrefix("fmt")
	s.emitString("Errorf(\"size of %d elements exceeds ")
	i.emitOriginalName(s)
	s.emitString("'s maximum of ")
	if err := t.maximumSizeElements.visit(newIntegerEmittingValueVisitor(s, r)); err != nil {
		return err
	}
	s.emitString(" elements\", len(m))\ngoto done\n}\n")

	// Write the number of elements in the array.
	s.emitString("nField, err = ")
	s.emitPackageNamePrefix(xdrRuntimePackage)
	s.emitString("WriteUnsignedInt(w, uint32(len(m)))\n")
	emitReadWriteErrorHandling(s)

	// Write the elements in the array.
	s.emitString("for _, m := range m {\n")
	if err := t.base.emitWriteTo(s, r, i.taint()); err != nil {
		return err
	}
	s.emitString("}\n")
	return nil
}

func (t *variableLengthArrayType) emitGetVariableEncodedSizeBytes(s sink, r Resolver, i identifier) error {
	elementSizeBytes, err := t.base.getFixedEncodedSizeBytes(r)
	if err != nil {
		return err
	}
	if elementSizeBytes != nil {
		s.emitString("nTotal += 4 + ")
		s.emitString(elementSizeBytes.String())
		s.emitString(" * len(m)\n")
	} else {
		s.emitString("nTotal += 4\n")
		s.emitString("for _, m := range m {\n")
		if err := t.base.emitGetVariableEncodedSizeBytes(s, r, i.taint()); err != nil {
			return err
		}
		s.emitString("}\n")
	}
	return nil
}

func (t *variableLengthArrayType) isComplexToReadOrWrite(r Resolver) (bool, error) {
	return true, nil
}

func (t *variableLengthArrayType) isLarge(r Resolver) (bool, error) {
	return false, nil
}

func (t *variableLengthArrayType) getFixedEncodedSizeBytes(r Resolver) (*big.Int, error) {
	return nil, nil
}
