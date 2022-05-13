package model

type foreignIdentifier struct {
	pkg  string
	name string
}

func newForeignIdentifier(pkg, name string) identifier {
	return &foreignIdentifier{
		pkg:  pkg,
		name: name,
	}
}

func (i *foreignIdentifier) emitPackageNamePrefix(s sink) {
	s.emitPackageNamePrefix(i.pkg)
}

func (i *foreignIdentifier) emit(s sink, forcePublic bool) {
	i.emitPackageNamePrefix(s)
	emitMixedCase(s, i.name, true)
}

func (i *foreignIdentifier) emitOriginalName(s sink) {
	s.emitString(i.name)
}

func (i *foreignIdentifier) emitStructuralFunctionName(s sink, functionName string) {
	i.emitPackageNamePrefix(s)
	s.emitString(functionName)
	emitMixedCase(s, i.name, true)
}

func (i *foreignIdentifier) hasTypeDefinition(isNominal bool) bool {
	return true
}

func (i *foreignIdentifier) hasReadFromWriteTo() bool {
	return true
}

func (i *foreignIdentifier) inside() identifier {
	panic("TODO: Do we need this?")
}

func (i *foreignIdentifier) resolve(pkg, name string) identifier {
	return &foreignIdentifier{
		pkg:  pkg,
		name: name,
	}
}

func (i *foreignIdentifier) append(typeUsedMultipleTimes, readFromWriteToUsedMultipleTimes bool) identifierFactory {
	panic("TODO: Do we need this?")
}

func (i *foreignIdentifier) taint() identifier {
	return taintedForeignIdentifier{}
}

type taintedForeignIdentifier struct{}

func (i taintedForeignIdentifier) emit(s sink, forcePublic bool) {
	panic("TODO")
}

func (i taintedForeignIdentifier) emitOriginalName(s sink) {
	panic("TODO")
}

func (i taintedForeignIdentifier) emitStructuralFunctionName(s sink, functionName string) {
	panic("TODO")
}

func (i taintedForeignIdentifier) hasTypeDefinition(isNominal bool) bool {
	panic("TODO")
}

func (i taintedForeignIdentifier) hasReadFromWriteTo() bool {
	panic("TODO")
}

func (i taintedForeignIdentifier) inside() identifier {
	panic("TODO")
}

func (i taintedForeignIdentifier) resolve(pkg, name string) identifier {
	return &foreignIdentifier{
		pkg:  pkg,
		name: name,
	}
}

func (i taintedForeignIdentifier) append(typeUsedMultipleTimes, readFromWriteToUsedMultipleTimes bool) identifierFactory {
	panic("TODO")
}

func (i taintedForeignIdentifier) taint() identifier {
	panic("TODO")
}
