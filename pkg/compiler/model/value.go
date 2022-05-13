package model

// Value that is either a constant or an identifier.
type Value interface {
	visit(ve valueVisitor) error
}
