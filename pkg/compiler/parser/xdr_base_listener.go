// Code generated from pkg/compiler/parser/XDR.g4 by ANTLR 4.10. DO NOT EDIT.

package parser // XDR

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseXDRListener is a complete listener for a parse tree produced by XDRParser.
type BaseXDRListener struct{}

var _ XDRListener = &BaseXDRListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseXDRListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseXDRListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseXDRListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseXDRListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTopLevel is called when production topLevel is entered.
func (s *BaseXDRListener) EnterTopLevel(ctx *TopLevelContext) {}

// ExitTopLevel is called when production topLevel is exited.
func (s *BaseXDRListener) ExitTopLevel(ctx *TopLevelContext) {}

// EnterImports is called when production imports is entered.
func (s *BaseXDRListener) EnterImports(ctx *ImportsContext) {}

// ExitImports is called when production imports is exited.
func (s *BaseXDRListener) ExitImports(ctx *ImportsContext) {}

// EnterStringConstant is called when production stringConstant is entered.
func (s *BaseXDRListener) EnterStringConstant(ctx *StringConstantContext) {}

// ExitStringConstant is called when production stringConstant is exited.
func (s *BaseXDRListener) ExitStringConstant(ctx *StringConstantContext) {}

// EnterNamedDeclaration is called when production namedDeclaration is entered.
func (s *BaseXDRListener) EnterNamedDeclaration(ctx *NamedDeclarationContext) {}

// ExitNamedDeclaration is called when production namedDeclaration is exited.
func (s *BaseXDRListener) ExitNamedDeclaration(ctx *NamedDeclarationContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseXDRListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseXDRListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterValue is called when production value is entered.
func (s *BaseXDRListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseXDRListener) ExitValue(ctx *ValueContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseXDRListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseXDRListener) ExitConstant(ctx *ConstantContext) {}

// EnterTypeSpecifier is called when production typeSpecifier is entered.
func (s *BaseXDRListener) EnterTypeSpecifier(ctx *TypeSpecifierContext) {}

// ExitTypeSpecifier is called when production typeSpecifier is exited.
func (s *BaseXDRListener) ExitTypeSpecifier(ctx *TypeSpecifierContext) {}

// EnterEnumTypeSpec is called when production enumTypeSpec is entered.
func (s *BaseXDRListener) EnterEnumTypeSpec(ctx *EnumTypeSpecContext) {}

// ExitEnumTypeSpec is called when production enumTypeSpec is exited.
func (s *BaseXDRListener) ExitEnumTypeSpec(ctx *EnumTypeSpecContext) {}

// EnterEnumBody is called when production enumBody is entered.
func (s *BaseXDRListener) EnterEnumBody(ctx *EnumBodyContext) {}

// ExitEnumBody is called when production enumBody is exited.
func (s *BaseXDRListener) ExitEnumBody(ctx *EnumBodyContext) {}

// EnterStructTypeSpec is called when production structTypeSpec is entered.
func (s *BaseXDRListener) EnterStructTypeSpec(ctx *StructTypeSpecContext) {}

// ExitStructTypeSpec is called when production structTypeSpec is exited.
func (s *BaseXDRListener) ExitStructTypeSpec(ctx *StructTypeSpecContext) {}

// EnterStructBody is called when production structBody is entered.
func (s *BaseXDRListener) EnterStructBody(ctx *StructBodyContext) {}

// ExitStructBody is called when production structBody is exited.
func (s *BaseXDRListener) ExitStructBody(ctx *StructBodyContext) {}

// EnterUnionTypeSpec is called when production unionTypeSpec is entered.
func (s *BaseXDRListener) EnterUnionTypeSpec(ctx *UnionTypeSpecContext) {}

// ExitUnionTypeSpec is called when production unionTypeSpec is exited.
func (s *BaseXDRListener) ExitUnionTypeSpec(ctx *UnionTypeSpecContext) {}

// EnterUnionBody is called when production unionBody is entered.
func (s *BaseXDRListener) EnterUnionBody(ctx *UnionBodyContext) {}

// ExitUnionBody is called when production unionBody is exited.
func (s *BaseXDRListener) ExitUnionBody(ctx *UnionBodyContext) {}

// EnterCaseSpec is called when production caseSpec is entered.
func (s *BaseXDRListener) EnterCaseSpec(ctx *CaseSpecContext) {}

// ExitCaseSpec is called when production caseSpec is exited.
func (s *BaseXDRListener) ExitCaseSpec(ctx *CaseSpecContext) {}

// EnterConstantDef is called when production constantDef is entered.
func (s *BaseXDRListener) EnterConstantDef(ctx *ConstantDefContext) {}

// ExitConstantDef is called when production constantDef is exited.
func (s *BaseXDRListener) ExitConstantDef(ctx *ConstantDefContext) {}

// EnterTypeDef is called when production typeDef is entered.
func (s *BaseXDRListener) EnterTypeDef(ctx *TypeDefContext) {}

// ExitTypeDef is called when production typeDef is exited.
func (s *BaseXDRListener) ExitTypeDef(ctx *TypeDefContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BaseXDRListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BaseXDRListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterSpecification is called when production specification is entered.
func (s *BaseXDRListener) EnterSpecification(ctx *SpecificationContext) {}

// ExitSpecification is called when production specification is exited.
func (s *BaseXDRListener) ExitSpecification(ctx *SpecificationContext) {}

// EnterProgramDef is called when production programDef is entered.
func (s *BaseXDRListener) EnterProgramDef(ctx *ProgramDefContext) {}

// ExitProgramDef is called when production programDef is exited.
func (s *BaseXDRListener) ExitProgramDef(ctx *ProgramDefContext) {}

// EnterVersionDef is called when production versionDef is entered.
func (s *BaseXDRListener) EnterVersionDef(ctx *VersionDefContext) {}

// ExitVersionDef is called when production versionDef is exited.
func (s *BaseXDRListener) ExitVersionDef(ctx *VersionDefContext) {}

// EnterProcedureDef is called when production procedureDef is entered.
func (s *BaseXDRListener) EnterProcedureDef(ctx *ProcedureDefContext) {}

// ExitProcedureDef is called when production procedureDef is exited.
func (s *BaseXDRListener) ExitProcedureDef(ctx *ProcedureDefContext) {}

// EnterProcReturn is called when production procReturn is entered.
func (s *BaseXDRListener) EnterProcReturn(ctx *ProcReturnContext) {}

// ExitProcReturn is called when production procReturn is exited.
func (s *BaseXDRListener) ExitProcReturn(ctx *ProcReturnContext) {}

// EnterProcFirstArg is called when production procFirstArg is entered.
func (s *BaseXDRListener) EnterProcFirstArg(ctx *ProcFirstArgContext) {}

// ExitProcFirstArg is called when production procFirstArg is exited.
func (s *BaseXDRListener) ExitProcFirstArg(ctx *ProcFirstArgContext) {}
