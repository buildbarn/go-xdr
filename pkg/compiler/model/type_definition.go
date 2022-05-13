package model

type typeDefinition struct {
	declaration NamedDeclaration
}

// NewTypeDefinition creates a new named type definition, as described
// in RFC 4506, section 4.18.
func NewTypeDefinition(declaration NamedDeclaration) Definition {
	return &typeDefinition{
		declaration: declaration,
	}
}

func (d *typeDefinition) register(r Registry, pkg string) error {
	nd := d.declaration
	return r.RegisterType(nd.name, pkg, nd.declarationType)
}

func (d *typeDefinition) emitAllDefinitions(s sink, r Resolver) error {
	s.emitString("\n")
	s.emitString("// Type definition \"")
	s.emitString(d.declaration.name)
	s.emitString("\".\n")
	return d.declaration.emitAllDeclarations(s, r, rootIdentifierFactory{})
}
