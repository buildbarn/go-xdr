package model

import (
	"math/big"
)

// Version of an RPC program, as described in RFC 5531, chapter 12.
type Version struct {
	name          string
	procedures    []Procedure
	versionNumber *big.Int
}

// NewVersion creates a new version of an RPC program.
func NewVersion(name string, procedures []Procedure, versionNumber *big.Int) Version {
	return Version{
		name:          name,
		procedures:    procedures,
		versionNumber: versionNumber,
	}
}
