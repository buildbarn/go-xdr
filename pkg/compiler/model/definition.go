package model

// Definition of a type, constant or program.
type Definition interface {
	register(r Registry, pkg string) error
	emitAllDefinitions(s sink, r Resolver) error
}
