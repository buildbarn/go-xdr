package model

import (
	"strings"
)

type rootIdentifierFactory struct{}

func (f rootIdentifierFactory) newIdentifier(name string) identifier {
	return &localIdentifier{
		name:                        name,
		hasTypeDefinitionNominal:    true,
		hasTypeDefinitionStructural: true,
		readFromWriteTo:             true,
	}
}

type localIdentifier struct {
	parent                      *localIdentifier
	name                        string
	isTainted                   bool
	hasTypeDefinitionNominal    bool
	hasTypeDefinitionStructural bool
	readFromWriteTo             bool
}

func (i *localIdentifier) emitUntainted(s sink, isPublic bool) {
	if i.parent == nil {
		emitMixedCase(s, i.name, isPublic)
	} else {
		i.parent.emitUntainted(s, isPublic)
		emitMixedCase(s, i.name, true)
	}
}

func (i *localIdentifier) emit(s sink, forcePublic bool) {
	i.emitUntainted(s, !i.isTainted && (i.parent == nil || forcePublic))
	if i.isTainted {
		s.emitString("Impl")
	}
}

func (i *localIdentifier) emitOriginalName(s sink) {
	if i.parent != nil {
		i.parent.emitOriginalName(s)
		s.emitString(".")
	}
	s.emitString(i.name)
}

func (i *localIdentifier) emitStructuralFunctionName(s sink, functionName string) {
	if !i.isTainted && i.parent == nil {
		s.emitString(functionName)
	} else {
		s.emitString(strings.ToLower(functionName))
	}
	i.emitUntainted(s, true)
}

func (i *localIdentifier) hasTypeDefinition(isNominal bool) bool {
	if isNominal {
		return i.hasTypeDefinitionNominal
	}
	return i.hasTypeDefinitionStructural
}

func (i *localIdentifier) hasReadFromWriteTo() bool {
	return i.readFromWriteTo
}

func (i *localIdentifier) inside() identifier {
	iInside := *i
	iInside.hasTypeDefinitionNominal = false
	iInside.hasTypeDefinitionStructural = false
	iInside.readFromWriteTo = false
	return &iInside
}

func (i *localIdentifier) resolve(pkg, name string) identifier {
	return newForeignIdentifier(pkg, name)
}

func (i *localIdentifier) append(typeUsedMultipleTimes, readFromWriteToUsedMultipleTimes bool) identifierFactory {
	return &localIdentifierFactory{
		parent:                           i,
		typeUsedMultipleTimes:            typeUsedMultipleTimes,
		readFromWriteToUsedMultipleTimes: readFromWriteToUsedMultipleTimes,
	}
}

func (i *localIdentifier) taint() identifier {
	return &localIdentifier{
		parent:    i.parent,
		name:      i.name,
		isTainted: true,
	}
}

type localIdentifierFactory struct {
	parent                           *localIdentifier
	typeUsedMultipleTimes            bool
	readFromWriteToUsedMultipleTimes bool
}

func (f *localIdentifierFactory) newIdentifier(name string) identifier {
	return &localIdentifier{
		parent:                   f.parent,
		name:                     name,
		hasTypeDefinitionNominal: f.typeUsedMultipleTimes,
		readFromWriteTo:          f.readFromWriteToUsedMultipleTimes,
	}
}
