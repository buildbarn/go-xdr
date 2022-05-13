package model

import (
	"math/big"
)

type constantResolvingValueVisitor struct {
	r        Resolver
	constant *big.Int
}

func newConstantResolvingValueVisitor(r Resolver) *constantResolvingValueVisitor {
	return &constantResolvingValueVisitor{r: r}
}

func (vv *constantResolvingValueVisitor) visitConstant(constant *big.Int) error {
	vv.constant = constant
	return nil
}

func (vv *constantResolvingValueVisitor) visitIdentifier(name string) error {
	v, err := vv.r.ResolveConstant(name)
	if err != nil {
		return err
	}
	vv.constant = v
	return nil
}
