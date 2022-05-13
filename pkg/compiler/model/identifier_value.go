package model

type identifierValue struct {
	identifier string
}

// NewIdentifierValue creates a value that corresponds to the name of
// another identifier. These objects are, for example, instantiated when
// a fixed length array is declared:
//
//     const LENGTH = 10;
//     typedef int mytype[LENGTH];
//                        ^^^^^^ NewIdentifierValue("LENGTH")
func NewIdentifierValue(identifier string) Value {
	return &identifierValue{
		identifier: identifier,
	}
}

func (v *identifierValue) visit(vv valueVisitor) error {
	return vv.visitIdentifier(v.identifier)
}
