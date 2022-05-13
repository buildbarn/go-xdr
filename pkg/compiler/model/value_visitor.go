package model

import (
	"math/big"
)

type valueVisitor interface {
	visitConstant(constant *big.Int) error
	visitIdentifier(name string) error
}
