package model

import (
	"math/big"
)

// Procedure of an RPC program, as described in RFC 5531, chapter 12.
type Procedure struct {
	name            string
	argumentTypes   []Type
	returnType      Type
	procedureNumber *big.Int
}

// NewProcedure creates a new procedure of an RPC program.
func NewProcedure(name string, argumentTypes []Type, returnType Type, procedureNumber *big.Int) Procedure {
	return Procedure{
		name:            name,
		argumentTypes:   argumentTypes,
		returnType:      returnType,
		procedureNumber: procedureNumber,
	}
}
