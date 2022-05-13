package model

import (
	"math/big"
)

type constantValue struct {
	constant *big.Int
}

// NewConstantValue creates a value that corresponds to a constant.
// These objects are, for example, instantiated when a fixed length
// array is declared:
//
//     typedef int mytype[10];
//                        ^^ NewConstantValue(10)
func NewConstantValue(constant *big.Int) Value {
	return &constantValue{
		constant: constant,
	}
}

func (v *constantValue) visit(vv valueVisitor) error {
	return vv.visitConstant(v.constant)
}
