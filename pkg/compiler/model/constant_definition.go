package model

import (
	"math/big"
)

type constantDefinition struct {
	name     string
	constant *big.Int
}

// NewConstantDefinition creates a new named constant definition, as
// described in RFC 4506, section 4.17. It ends up generating a single
// constant definition in Go as well.
func NewConstantDefinition(name string, constant *big.Int) Definition {
	return &constantDefinition{
		name:     name,
		constant: constant,
	}
}

func (d *constantDefinition) register(r Registry, pkg string) error {
	return r.RegisterConstant(d.name, d.constant)
}

func (d *constantDefinition) emitAllDefinitions(s sink, r Resolver) error {
	s.emitString("\n")
	s.emitString("const ")
	emitMacroCase(s, d.name)
	s.emitString(" = ")
	s.emitString(d.constant.String())
	s.emitString("\n")
	return nil
}
