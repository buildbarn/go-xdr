package model

import (
	"math/big"
)

// Registry can be implemented by consumers of this package to extract
// constants, types and imports from a TopLevel.
type Registry interface {
	RegisterConstant(name string, c *big.Int) error
	RegisterType(name, pkg string, t Type) error
	RegisterImport(path string) error
}
