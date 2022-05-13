package model

import (
	"math/big"
	"strings"
	"unicode"
)

type sink interface {
	emitString(v string)
	emitPackageNamePrefix(pkg string)
}

const xdrRuntimePackage = "github.com/buildbarn/go-xdr/pkg/runtime"

func emitMacroCase(s sink, name string) {
	s.emitString(strings.ToUpper(name))
}

func emitMixedCase(s sink, name string, isPublic bool) {
	conversion := unicode.ToLower
	if isPublic {
		conversion = unicode.ToUpper
	}

	for _, field := range strings.FieldsFunc(name, func(r rune) bool { return r == '_' }) {
		isFirst := true
		s.emitString(strings.Map(func(r rune) rune {
			if isFirst {
				isFirst = false
				return conversion(r)
			}
			return unicode.ToLower(r)
		}, field))
		conversion = unicode.ToUpper
	}
}

func emitConstantFixedEncodedSizeBytes(s sink, i identifier, size *big.Int) {
	s.emitString("const ")
	i.emit(s, false)
	s.emitString("EncodedSizeBytes = ")
	s.emitString(size.String())
	s.emitString("\n")
}

func emitDeclarationsStructural(s sink, r Resolver, i identifier, t Type, useTypeEquals, allMethods bool) error {
	if i.hasTypeDefinition(false) {
		s.emitString("\n")
		s.emitString("type ")
		i.emit(s, false)
		if useTypeEquals {
			s.emitString(" = ")
		} else {
			s.emitString(" ")
		}
		if err := t.emitDeclaration(s, r, i.inside()); err != nil {
			return err
		}
		s.emitString("\n")
	}

	if i.hasReadFromWriteTo() {
		isLarge, err := t.isLarge(r)
		if err != nil {
			return err
		}

		s.emitString("\n")
		if isLarge {
			s.emitString("func ")
			i.emitStructuralFunctionName(s, "Read")
			s.emitString("(r ")
			s.emitPackageNamePrefix("io")
			s.emitString("Reader, m *")
			if err := t.emitDeclaration(s, r, i); err != nil {
				return err
			}
			s.emitString(") (nTotal int64, err error) {\n")
		} else {
			s.emitString("func ")
			i.emitStructuralFunctionName(s, "Read")
			s.emitString("(r ")
			s.emitPackageNamePrefix("io")
			s.emitString("Reader) (m ")
			if err := t.emitDeclaration(s, r, i); err != nil {
				return err
			}
			s.emitString(", nTotal int64, err error) {\n")
		}
		s.emitString("var nField int64\n")
		if err := t.emitReadFrom(s, r, i.inside()); err != nil {
			return err
		}
		s.emitString("done:\n")
		s.emitString("return\n")
		s.emitString("}\n")

		if allMethods {
			s.emitString("\n")
			s.emitString("func ")
			i.emitStructuralFunctionName(s, "Write")
			s.emitString("(w ")
			s.emitPackageNamePrefix("io")
			s.emitString("Writer, m ")
			if isLarge {
				s.emitString("*")
			}
			if err := t.emitDeclaration(s, r, i); err != nil {
				return err
			}
			s.emitString(") (nTotal int64, err error) {\n")
			s.emitString("var nField int64\n")
			if err := t.emitWriteTo(s, r, i.inside()); err != nil {
				return err
			}
			s.emitString("done:\n")
			s.emitString("return\n")
			s.emitString("}\n")

			s.emitString("\n")
			if size, err := t.getFixedEncodedSizeBytes(r); err != nil {
				return err
			} else if size != nil {
				emitConstantFixedEncodedSizeBytes(s, i, size)
			} else {
				s.emitString("func Get")
				i.emit(s, true)
				s.emitString("EncodedSizeBytes(m ")
				if isLarge {
					s.emitString("*")
				}
				if err := t.emitDeclaration(s, r, i); err != nil {
					return err
				}
				s.emitString(") (nTotal int) {\n")
				if err := t.emitGetVariableEncodedSizeBytes(s, r, i.inside()); err != nil {
					return err
				}
				s.emitString("return\n")
				s.emitString("}\n")
			}
		}
	}
	return nil
}

func emitForwardReadFromStructural(s sink, r Resolver, i identifier, isLarge bool) error {
	if isLarge {
		s.emitString("nField, err = ")
		i.emitStructuralFunctionName(s, "Read")
		s.emitString("(r, m)\n")
	} else {
		s.emitString("m, nField, err = ")
		i.emitStructuralFunctionName(s, "Read")
		s.emitString("(r)\n")
	}
	emitReadWriteErrorHandling(s)
	return nil
}

func emitForwardWriteToStructural(s sink, i identifier) {
	s.emitString("nField, err = ")
	i.emitStructuralFunctionName(s, "Write")
	s.emitString("(w, m)\n")
	emitReadWriteErrorHandling(s)
}

func emitForwardGetVariableEncodedSizeBytesStructural(s sink, i identifier) {
	s.emitString("nTotal += ")
	i.emitStructuralFunctionName(s, "Get")
	s.emitString("EncodedSizeBytes(m)\n")
}

func emitAdditionalDeclarationsNominal(s sink, r Resolver, i identifier, t Type) error {
	if i.hasReadFromWriteTo() {
		isLarge, err := t.isLarge(r)
		if err != nil {
			return err
		}

		s.emitString("\n")
		if isLarge {
			s.emitString("func (m *")
		} else {
			s.emitString("func (mParent *")
		}
		i.emit(s, false)
		s.emitString(") ReadFrom(r ")
		s.emitPackageNamePrefix("io")
		s.emitString("Reader) (nTotal int64, err error) {\n")
		s.emitString("var nField int64\n")
		if isLarge {
			if err := t.emitReadFrom(s, r, i.inside()); err != nil {
				return err
			}
		} else {
			s.emitString("var m ")
			i.emit(s, false)
			s.emitString("\n")
			if err := t.emitReadFrom(s, r, i.inside()); err != nil {
				return err
			}
			s.emitString("*mParent = m\n")
		}
		s.emitString("done:\n")
		s.emitString("return\n")
		s.emitString("}\n")

		s.emitString("\n")
		s.emitString("func (m ")
		if isLarge {
			s.emitString("*")
		}
		i.emit(s, false)
		s.emitString(") WriteTo(w ")
		s.emitPackageNamePrefix("io")
		s.emitString("Writer) (nTotal int64, err error) {\n")
		s.emitString("var nField int64\n")
		if err := t.emitWriteTo(s, r, i.inside()); err != nil {
			return err
		}
		s.emitString("done:\n")
		s.emitString("return\n")
		s.emitString("}\n")

		s.emitString("\n")
		if size, err := t.getFixedEncodedSizeBytes(r); err != nil {
			return err
		} else if size != nil {
			emitConstantFixedEncodedSizeBytes(s, i, size)
		} else {
			s.emitString("func (m ")
			if isLarge {
				s.emitString("*")
			}
			i.emit(s, true)
			s.emitString(") GetEncodedSizeBytes() (nTotal int) {\n")
			if err := t.emitGetVariableEncodedSizeBytes(s, r, i.inside()); err != nil {
				return err
			}
			s.emitString("return\n")
			s.emitString("}\n")
		}
	}

	return nil
}

func emitForwardReadFromNominal(s sink) {
	s.emitString("nField, err = m.ReadFrom(r)\n")
	emitReadWriteErrorHandling(s)
}

func emitForwardWriteToNominal(s sink) {
	s.emitString("nField, err = m.WriteTo(w)\n")
	emitReadWriteErrorHandling(s)
}

func emitForwardGetVariableEncodedSizeBytesNominal(s sink) {
	s.emitString("nTotal += m.GetEncodedSizeBytes()\n")
}

func emitReadWriteErrorHandling(s sink) {
	s.emitString("nTotal += nField\n")
	s.emitString("if err != nil {\n")
	s.emitString("goto done\n")
	s.emitString("}\n")
}
