package model

type identifierFactory interface {
	newIdentifier(name string) identifier
}
