package model

import (
	"math/big"
)

type integerEmittingValueVisitor struct {
	s sink
	r Resolver
}

func newIntegerEmittingValueVisitor(s sink, r Resolver) valueVisitor {
	return &integerEmittingValueVisitor{
		s: s,
		r: r,
	}
}

func (vv *integerEmittingValueVisitor) visitConstant(constant *big.Int) error {
	vv.s.emitString(constant.String())
	return nil
}

func (vv *integerEmittingValueVisitor) visitIdentifier(name string) error {
	v, err := vv.r.ResolveConstant(name)
	if err != nil {
		return err
	}
	vv.s.emitString(v.String())
	return nil
}
