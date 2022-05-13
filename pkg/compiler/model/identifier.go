package model

type identifier interface {
	emit(s sink, forcePublic bool)
	emitOriginalName(s sink)
	emitStructuralFunctionName(s sink, functionName string)

	hasTypeDefinition(isNominal bool) bool
	hasReadFromWriteTo() bool

	resolve(pkg, name string) identifier
	append(typeUsedMultipleTimes, readFromWriteToUsedMultipleTimes bool) identifierFactory
	inside() identifier
	taint() identifier
}
