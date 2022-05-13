package model

import (
	"math/big"
)

type identifierType struct {
	identifier string
}

// NewIdentifierType creates a type that corresponds to an identifier.
// It merely forwards all methods to the underlying type of that
// identifier after resolving it. This can be used to simply alias
// types, or to compose them:
//
//     typedef int t1;
//             ^^^ IntType
//     typedef t1 t2;
//             ^^ NewIdentifierType("t1")
func NewIdentifierType(identifier string) Type {
	return &identifierType{
		identifier: identifier,
	}
}

func (t *identifierType) getEmittingValueVisitor(s sink, r Resolver) (valueVisitor, error) {
	_, tNext, err := r.ResolveType(t.identifier)
	if err != nil {
		return nil, err
	}
	return tNext.getEmittingValueVisitor(s, r)
}

func (t *identifierType) emitDeclaration(s sink, r Resolver, i identifier) error {
	pkg, tNext, err := r.ResolveType(t.identifier)
	if err != nil {
		return err
	}
	return tNext.emitDeclaration(s, r, i.resolve(pkg, t.identifier))
}

func (t *identifierType) emitAllDeclarations(s sink, r Resolver, i identifier) error {
	return emitDeclarationsStructural(s, r, i, t, true, true)
}

func (t *identifierType) emitReadFrom(s sink, r Resolver, i identifier) error {
	pkg, tNext, err := r.ResolveType(t.identifier)
	if err != nil {
		return err
	}
	return tNext.emitReadFrom(s, r, i.resolve(pkg, t.identifier))
}

func (t *identifierType) emitWriteTo(s sink, r Resolver, i identifier) error {
	pkg, tNext, err := r.ResolveType(t.identifier)
	if err != nil {
		return err
	}
	return tNext.emitWriteTo(s, r, i.resolve(pkg, t.identifier))
}

func (t *identifierType) emitGetVariableEncodedSizeBytes(s sink, r Resolver, i identifier) error {
	pkg, tNext, err := r.ResolveType(t.identifier)
	if err != nil {
		return err
	}
	return tNext.emitGetVariableEncodedSizeBytes(s, r, i.resolve(pkg, t.identifier))
}

func (t *identifierType) isLarge(r Resolver) (bool, error) {
	_, tNext, err := r.ResolveType(t.identifier)
	if err != nil {
		return false, err
	}
	return tNext.isLarge(r)
}

func (t *identifierType) getFixedEncodedSizeBytes(r Resolver) (*big.Int, error) {
	_, tNext, err := r.ResolveType(t.identifier)
	if err != nil {
		return nil, err
	}
	return tNext.getFixedEncodedSizeBytes(r)
}
