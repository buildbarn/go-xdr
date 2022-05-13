package model

import (
	"errors"
	"math/big"
)

type rawEmittingValueVisitor struct {
	s sink
}

func (vv *rawEmittingValueVisitor) visitConstant(constant *big.Int) error {
	vv.s.emitString(constant.String())
	return nil
}

func (vv *rawEmittingValueVisitor) visitIdentifier(name string) error {
	vv.s.emitString(name)
	return nil
}

func newRawEmittingValueVisitor(s sink) valueVisitor {
	return &rawEmittingValueVisitor{s: s}
}

// CaseSpec corresponds to an arm of a discriminated union.
type CaseSpec struct {
	values      []Value
	declaration Declaration
}

// NewCaseSpec creates a new arm of a discriminated union.
func NewCaseSpec(values []Value, declaration Declaration) CaseSpec {
	return CaseSpec{
		values:      values,
		declaration: declaration,
	}
}

type unionType struct {
	discriminant NamedDeclaration
	cases        []CaseSpec
	defaultCase  Declaration
}

// NewUnionType creates a discriminated union type, as described in RFC
// 4506, section 4.15. It translates to an interface type in Go. Each
// arm of the discriminant union translates to a struct that satisfies
// the interface type.
func NewUnionType(discriminant NamedDeclaration, cases []CaseSpec, defaultCase Declaration) Type {
	return &unionType{
		discriminant: discriminant,
		cases:        cases,
		defaultCase:  defaultCase,
	}
}

func (t *unionType) getEmittingValueVisitor(s sink, r Resolver) (valueVisitor, error) {
	return nil, errors.New("Union constants are not supported")
}

func (t *unionType) emitDeclaration(s sink, r Resolver, i identifier) error {
	if i.hasTypeDefinition(true) {
		i.emit(s, false)
	} else {
		fDiscriminant := i.append(true, true)
		iDiscriminant := fDiscriminant.newIdentifier(t.discriminant.name)
		tDiscriminant := t.discriminant.declarationType

		s.emitString("interface {\n")

		s.emitString("is")
		i.emit(s, true)
		s.emitString("()\n")

		s.emitString("Get")
		emitMixedCase(s, t.discriminant.name, true)
		s.emitString("() ")
		if err := tDiscriminant.emitDeclaration(s, r, iDiscriminant); err != nil {
			return err
		}
		s.emitString("\n")

		s.emitPackageNamePrefix("io")
		s.emitString("WriterTo\n")

		if size, err := t.getFixedEncodedSizeBytes(r); err != nil {
			return err
		} else if size == nil {
			s.emitString("GetEncodedSizeBytes() int\n")
		}

		s.emitString("}")
	}
	return nil
}

func (t *unionType) emitAllDeclarations(s sink, r Resolver, i identifier) error {
	if err := emitDeclarationsStructural(s, r, i, t, false, false); err != nil {
		return err
	}
	fDiscriminant := i.append(true, true)
	if err := t.discriminant.emitAllDeclarations(s, r, fDiscriminant); err != nil {
		return err
	}
	for _, caseSpec := range t.cases {
		// TODO
		fArm := i.append(false, false)
		if err := caseSpec.declaration.emitAllDeclarations(s, r, fArm); err != nil {
			return err
		}
	}

	iDiscriminant := fDiscriminant.newIdentifier(t.discriminant.name)
	tDiscriminant := t.discriminant.declarationType
	discriminantEmittingValueVisitor, err := tDiscriminant.getEmittingValueVisitor(s, r)
	if err != nil {
		return err
	}

	size, err := t.getFixedEncodedSizeBytes(r)
	if err != nil {
		return err
	}
	emitGetEncodedSizeBytes := size == nil

	for _, caseSpec := range t.cases {
		// TODO
		fArm := i.append(false, false)
		for _, value := range caseSpec.values {
			s.emitString("\n")
			s.emitString("type ")
			i.emit(s, true)
			s.emitString("_")
			if err := value.visit(newRawEmittingValueVisitor(s)); err != nil {
				return err
			}
			s.emitString(" struct {\n")
			if err := caseSpec.declaration.emitStructField(s, r, fArm); err != nil {
				return err
			}
			s.emitString("}\n")

			s.emitString("\n")
			s.emitString("func (m *")
			i.emit(s, true)
			s.emitString("_")
			if err := value.visit(newRawEmittingValueVisitor(s)); err != nil {
				return err
			}
			s.emitString(") is")
			i.emit(s, true)
			s.emitString("() {}\n")

			// Emit a method for getting the discriminant's value.
			s.emitString("\n")
			s.emitString("func (m *")
			i.emit(s, true)
			s.emitString("_")
			if err := value.visit(newRawEmittingValueVisitor(s)); err != nil {
				return err
			}
			s.emitString(") Get")
			emitMixedCase(s, t.discriminant.name, true)
			s.emitString("() ")
			if err := tDiscriminant.emitDeclaration(s, r, iDiscriminant); err != nil {
				return err
			}
			s.emitString(" {\nreturn ")
			if err := value.visit(discriminantEmittingValueVisitor); err != nil {
				return err
			}
			s.emitString("\n}\n")

			s.emitString("\n")
			s.emitString("func (m *")
			i.emit(s, true)
			s.emitString("_")
			if err := value.visit(newRawEmittingValueVisitor(s)); err != nil {
				return err
			}
			s.emitString(") WriteTo(w ")
			s.emitPackageNamePrefix("io")
			s.emitString("Writer) (nTotal int64, err error) {\n")
			s.emitString("var nField int64\n")

			// Write the discriminant.
			s.emitString("{\n")
			s.emitString("var m ")
			if err := tDiscriminant.emitDeclaration(s, r, iDiscriminant); err != nil {
				return err
			}
			s.emitString(" = ")
			if err := value.visit(discriminantEmittingValueVisitor); err != nil {
				return err
			}
			s.emitString("\n")
			if err := tDiscriminant.emitWriteTo(s, r, iDiscriminant); err != nil {
				return err
			}
			s.emitString("}\n")

			// Write the arm.
			// TODO: Would it be possible to let all cases
			// of the same type use a single write function?
			if err := caseSpec.declaration.emitStructFieldWriteTo(s, r, fArm); err != nil {
				return err
			}
			s.emitString("done:\n")
			s.emitString("return\n")
			s.emitString("}\n")

			if emitGetEncodedSizeBytes {
				s.emitString("\n")
				s.emitString("func (m *")
				i.emit(s, true)
				s.emitString("_")
				if err := value.visit(newRawEmittingValueVisitor(s)); err != nil {
					return err
				}
				s.emitString(") GetEncodedSizeBytes() (nTotal int) {\n")
				if err := t.discriminant.maybeEmitStructFieldGetVariableEncodedSizeBytes(s, r, fDiscriminant); err != nil {
					return err
				}
				if err := caseSpec.declaration.maybeEmitStructFieldGetVariableEncodedSizeBytes(s, r, fArm); err != nil {
					return err
				}
				s.emitString("return\n")
				s.emitString("}\n")
			}
		}
	}
	if t.defaultCase != nil {
		// TODO
		fArm := i.append(false, false)

		s.emitString("\n")
		s.emitString("type ")
		i.emit(s, true)
		s.emitString("_default struct {\n")
		if err := t.discriminant.emitStructField(s, r, fDiscriminant); err != nil {
			return err
		}
		if err := t.defaultCase.emitStructField(s, r, fArm); err != nil {
			return err
		}
		s.emitString("}\n")

		s.emitString("\n")
		s.emitString("func (m *")
		i.emit(s, true)
		s.emitString("_default) is")
		i.emit(s, true)
		s.emitString("() {}\n")

		// Emit a method for getting the discriminant's value.
		s.emitString("\n")
		s.emitString("func (m *")
		i.emit(s, true)
		s.emitString("_default) Get")
		emitMixedCase(s, t.discriminant.name, true)
		s.emitString("() ")
		if err := tDiscriminant.emitDeclaration(s, r, iDiscriminant); err != nil {
			return err
		}
		s.emitString(" {\nreturn m.")
		emitMixedCase(s, t.discriminant.name, true)
		s.emitString("\n}\n")

		s.emitString("\n")
		s.emitString("func (m *")
		i.emit(s, true)
		s.emitString("_default) WriteTo(w ")
		s.emitPackageNamePrefix("io")
		s.emitString("Writer) (nTotal int64, err error) {\n")
		s.emitString("var nField int64\n")

		// Write the discriminant.
		s.emitString("{\n")
		s.emitString("m := m.")
		emitMixedCase(s, t.discriminant.name, true)
		s.emitString("\n")
		if err := tDiscriminant.emitWriteTo(s, r, iDiscriminant); err != nil {
			return err
		}
		s.emitString("}\n")

		// Write the arm.
		// TODO: Would it be possible to let all cases
		// of the same type use a single write function?
		if err := t.defaultCase.emitStructFieldWriteTo(s, r, fArm); err != nil {
			return err
		}
		s.emitString("done:\n")
		s.emitString("return\n")
		s.emitString("}\n")

		if emitGetEncodedSizeBytes {
			s.emitString("\n")
			s.emitString("func (m *")
			i.emit(s, true)
			s.emitString("_default) GetEncodedSizeBytes() (nTotal int) {\n")
			if err := t.discriminant.maybeEmitStructFieldGetVariableEncodedSizeBytes(s, r, fDiscriminant); err != nil {
				return err
			}
			if err := t.defaultCase.maybeEmitStructFieldGetVariableEncodedSizeBytes(s, r, fArm); err != nil {
				return err
			}
			s.emitString("return\n")
			s.emitString("}\n")
		}
	}
	return nil
}

func (t *unionType) emitReadFrom(s sink, r Resolver, i identifier) error {
	if i.hasReadFromWriteTo() {
		return emitForwardReadFromStructural(s, r, i, false)
	}

	fDiscriminant := i.append(true, true)
	iDiscriminant := fDiscriminant.newIdentifier(t.discriminant.name)
	tDiscriminant := t.discriminant.declarationType
	s.emitString("var discriminant ")
	if err := tDiscriminant.emitDeclaration(s, r, iDiscriminant); err != nil {
		return err
	}
	s.emitString("\n")
	s.emitString("{\n")
	s.emitString("var m ")
	if err := tDiscriminant.emitDeclaration(s, r, iDiscriminant); err != nil {
		return err
	}
	s.emitString("\n")
	if err := tDiscriminant.emitReadFrom(s, r, iDiscriminant); err != nil {
		return err
	}
	s.emitString("discriminant = m\n")
	s.emitString("}\n")

	discriminantEmittingValueVisitor, err := tDiscriminant.getEmittingValueVisitor(s, r)
	if err != nil {
		return err
	}
	s.emitString("switch discriminant {\n")
	for _, caseSpec := range t.cases {
		// TODO
		fArm := i.append(false, false)
		for _, value := range caseSpec.values {
			s.emitString("case ")

			if err := value.visit(discriminantEmittingValueVisitor); err != nil {
				return err
			}
			s.emitString(":\n")

			s.emitString("var mArm ")
			i.emit(s, true)
			s.emitString("_")
			value.visit(newRawEmittingValueVisitor(s))
			s.emitString("\n")
			s.emitString("{\n")
			s.emitString("m := &mArm\n")
			if err := caseSpec.declaration.emitStructFieldReadFrom(s, r, fArm); err != nil {
				return err
			}
			s.emitString("}\n")
			s.emitString("m = &mArm\n")
		}
	}
	s.emitString("default:\n")
	if t.defaultCase == nil {
		s.emitString("err = ")
		s.emitPackageNamePrefix("fmt")
		s.emitString("Errorf(\"discriminant ")
		iDiscriminant.emitOriginalName(s)
		s.emitString(" has unknown value %d\", discriminant)\ngoto done\n")
	} else {
		// TODO
		fArm := i.append(false, false)
		s.emitString("var mArm ")
		i.emit(s, true)
		s.emitString("_default\n{\nm := &mArm\nm.")
		emitMixedCase(s, t.discriminant.name, true)
		s.emitString(" = discriminant\n")
		if err := t.defaultCase.emitStructFieldReadFrom(s, r, fArm); err != nil {
			return err
		}
		s.emitString("}\n")
		s.emitString("m = &mArm\n")
	}
	s.emitString("}\n")
	return nil
}

func (t *unionType) emitWriteTo(s sink, r Resolver, i identifier) error {
	emitForwardWriteToNominal(s)
	return nil
}

func (t *unionType) emitGetVariableEncodedSizeBytes(s sink, r Resolver, i identifier) error {
	emitForwardGetVariableEncodedSizeBytesNominal(s)
	return nil
}

func (t *unionType) isComplexToReadOrWrite(r Resolver) (bool, error) {
	return true, nil
}

func (t *unionType) isLarge(r Resolver) (bool, error) {
	return false, nil
}

func (t *unionType) getFixedEncodedSizeBytes(r Resolver) (*big.Int, error) {
	discriminantSize, err := t.discriminant.declarationType.getFixedEncodedSizeBytes(r)
	if err != nil || discriminantSize == nil {
		return nil, err
	}
	armSize, err := t.cases[0].declaration.getFixedEncodedSizeBytes(r)
	if err != nil || armSize == nil {
		return nil, err
	}
	for _, caseSpec := range t.cases[1:] {
		otherArmSize, err := caseSpec.declaration.getFixedEncodedSizeBytes(r)
		if err != nil || otherArmSize == nil {
			return nil, err
		}
		if armSize.Cmp(otherArmSize) != 0 {
			return nil, nil
		}
	}
	if t.defaultCase != nil {
		otherArmSize, err := t.defaultCase.getFixedEncodedSizeBytes(r)
		if err != nil || otherArmSize == nil {
			return nil, err
		}
		if armSize.Cmp(otherArmSize) != 0 {
			return nil, nil
		}
	}
	var total big.Int
	return total.Add(discriminantSize, armSize), nil
}
