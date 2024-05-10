package model

import (
	"math/big"
	"strconv"
)

const rpcv2Package = "github.com/buildbarn/go-xdr/pkg/protocols/rpcv2"

type programDefinition struct {
	name          string
	versions      []Version
	programNumber *big.Int
}

// NewProgramDefinition creates a definition of an RPC program, as
// described in RFC 5531, chapter 12.
func NewProgramDefinition(name string, versions []Version, programNumber *big.Int) Definition {
	return programDefinition{
		name:          name,
		versions:      versions,
		programNumber: programNumber,
	}
}

func (d programDefinition) register(r Registry, pkg string) error {
	// TODO: This should iterate over versions!
	return nil
}

func (d programDefinition) emitAllDefinitions(s sink, r Resolver) error {
	s.emitString("\n")
	s.emitString("const ")
	emitMacroCase(s, d.name)
	s.emitString("_PROGRAM_NUMBER uint32 = ")
	s.emitString(d.programNumber.String())
	s.emitString("\n")

	i := rootIdentifierFactory{}.newIdentifier("TODO")

	s.emitString("\n")
	s.emitString("type ")
	emitMixedCase(s, d.name, true)
	s.emitString(" interface {\n")
	for _, version := range d.versions {
		for _, procedure := range version.procedures {
			emitMixedCase(s, version.name, true)
			emitMixedCase(s, procedure.name, true)
			s.emitString("(")
			s.emitPackageNamePrefix("context")
			s.emitString("Context")
			for _, t := range procedure.argumentTypes {
				s.emitString(", ")
				if large, err := t.isLarge(r); err != nil {
					return err
				} else if large {
					s.emitString("*")
				}
				if err := t.emitDeclaration(s, r, i); err != nil {
					return err
				}
			}
			s.emitString(")")
			if t := procedure.returnType; t == nil {
				s.emitString(" error\n")
			} else {
				s.emitString(" (")
				if large, err := t.isLarge(r); err != nil {
					return err
				} else if large {
					s.emitString("*")
				}
				if err := t.emitDeclaration(s, r, i); err != nil {
					return err
				}
				s.emitString(", error)\n")
			}
		}
	}
	s.emitString("}\n")

	s.emitString("\n")
	s.emitString("func New")
	emitMixedCase(s, d.name, true)
	s.emitString("Service(p ")
	emitMixedCase(s, d.name, true)
	s.emitString(") func(")
	s.emitPackageNamePrefix("context")
	s.emitString("Context, uint32, uint32, ")
	s.emitPackageNamePrefix("io")
	s.emitString("ReadCloser, ")
	s.emitPackageNamePrefix("io")
	s.emitString("Writer) (")
	s.emitPackageNamePrefix(rpcv2Package)
	s.emitString("AcceptedReplyData, error) {\n")
	s.emitString("return func(ctx ")
	s.emitPackageNamePrefix("context")
	s.emitString("Context, vers uint32, proc uint32, r ")
	s.emitPackageNamePrefix("io")
	s.emitString("ReadCloser, w ")
	s.emitPackageNamePrefix("io")
	s.emitString("Writer) (")
	s.emitPackageNamePrefix(rpcv2Package)
	s.emitString("AcceptedReplyData, error) {\nvar err error\nswitch vers {\n")
	needsErrorHandling := false
	for _, version := range d.versions {
		s.emitString("case ")
		s.emitString(version.versionNumber.String())
		s.emitString(":\nswitch proc {\n")
		for _, procedure := range version.procedures {
			s.emitString("case ")
			s.emitString(procedure.procedureNumber.String())
			s.emitString(":\n")

			for index, argumentType := range procedure.argumentTypes {
				indexStr := strconv.FormatInt(int64(index), 10)
				s.emitString("var a")
				s.emitString(indexStr)
				s.emitString(" ")
				if err := argumentType.emitDeclaration(s, r, i); err != nil {
					return err
				}
				s.emitString("\n{\n")
				large, err := argumentType.isLarge(r)
				if err != nil {
					return err
				}
				if large {
					s.emitString("m := &a")
					s.emitString(indexStr)
				} else {
					s.emitString("var m ")
					if err := argumentType.emitDeclaration(s, r, i); err != nil {
						return err
					}
				}
				s.emitString("\nvar nField, nTotal int64\n")
				if err := argumentType.emitReadFrom(s, r, i); err != nil {
					return err
				}
				if !large {
					s.emitString("a")
					s.emitString(indexStr)
					s.emitString(" = m\n")
				}
				s.emitString("}\n")

				needsErrorHandling = true
			}
			s.emitString("r.Close()\nr = nil\n")
			if procedure.returnType != nil {
				s.emitString("m, ")
			}
			s.emitString("err := p.")
			emitMixedCase(s, version.name, true)
			emitMixedCase(s, procedure.name, true)
			s.emitString("(ctx")
			for index, argumentType := range procedure.argumentTypes {
				s.emitString(", ")
				if large, err := argumentType.isLarge(r); err != nil {
					return err
				} else if large {
					s.emitString("&")
				}
				s.emitString("a")
				s.emitString(strconv.FormatInt(int64(index), 10))
			}
			s.emitString(")\nif err != nil {\nreturn nil, err\n}\n")
			if returnType := procedure.returnType; returnType != nil {
				s.emitString("{\nvar nField, nTotal int64\n")
				if err := returnType.emitWriteTo(s, r, i); err != nil {
					return err
				}
				s.emitString("}\n")
			}
		}
		s.emitString("default:\nr.Close()\nreturn &")
		s.emitPackageNamePrefix(rpcv2Package)
		s.emitString("AcceptedReplyData_default{Stat: ")
		s.emitPackageNamePrefix(rpcv2Package)
		s.emitString("PROC_UNAVAIL}, nil\n}\n")
	}
	s.emitString("default:\nr.Close()\nvar replyData ")
	s.emitPackageNamePrefix(rpcv2Package)
	minVersion, maxVersion := d.versions[0].versionNumber, d.versions[0].versionNumber
	for _, version := range d.versions[1:] {
		if minVersion.Cmp(version.versionNumber) > 0 {
			minVersion = version.versionNumber
		}
		if maxVersion.Cmp(version.versionNumber) < 0 {
			maxVersion = version.versionNumber
		}
	}
	s.emitString("AcceptedReplyData_PROG_MISMATCH\nreplyData.MismatchInfo.Low = ")
	s.emitString(minVersion.String())
	s.emitString("\nreplyData.MismatchInfo.High = ")
	s.emitString(maxVersion.String())
	s.emitString("\nreturn &replyData, nil\n")
	s.emitString("}\nreturn &")
	s.emitPackageNamePrefix(rpcv2Package)
	s.emitString("AcceptedReplyData_SUCCESS{}, nil\n")
	if needsErrorHandling {
		s.emitString("done:\nif r != nil {\nr.Close()\n}\nreturn nil, err\n")
	}
	s.emitString("}\n")
	s.emitString("}\n")

	return nil
}
