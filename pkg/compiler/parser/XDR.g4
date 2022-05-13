grammar XDR;

@parser::header {
import (
    "math"
    "math/big"

    "github.com/buildbarn/go-xdr/pkg/compiler/model"
)

var defaultArraySize = model.NewConstantValue(big.NewInt(math.MaxUint32))
}

// Top-level structure
//
// Below is custom syntax that is inspired by Protobuf that can be used
// to let XDR files include each other.

topLevel returns [model.TopLevel v]
    : 'package' stringConstant ';' imports specification EOF {
          $v = model.NewTopLevel($stringConstant.v, $imports.v, $specification.v)
      }
    ;

imports returns [[]string v]
    : (
          'import' stringConstant ';' {
              $v = append($v, $stringConstant.v)
          }
      )*
    ;

stringConstant returns [string v]
    : STRING_CONSTANT {
          s := $STRING_CONSTANT.text
          $v = s[1:len(s)-1]
      }
    ;

// RFC 4506: XDR: External Data Representation Standard
// https://tools.ietf.org/html/rfc4506

// Section 6.3: Syntax Information

// Deviation from the RFC: declaration has been decomposed into
// namedDeclaration. The RFC allows definitions as follows:
//
// typedef void;
// struct empty { void; };
//
// These definitions are meaningless and are also rejected by tools such
// as rpcgen.
//
// An 'utf8string' type has been added. This type is normally declared as part
// of NFSv4 as follows:
//
// typedef opaque  utf8string<>;
//
// We declare it as a native type, so that we don't need to work with
// byte slices.

namedDeclaration returns [model.NamedDeclaration v]
    : typeSpecifier IDENTIFIER {
          $v = model.NewNamedDeclaration($IDENTIFIER.text, $typeSpecifier.v)
      }
    | typeSpecifier IDENTIFIER '[' value ']' {
          $v = model.NewNamedDeclaration(
              $IDENTIFIER.text,
              model.NewFixedLengthArrayType($typeSpecifier.v, $value.v))
      }
    | typeSpecifier IDENTIFIER '<' {
          $v = model.NewNamedDeclaration(
              $IDENTIFIER.text,
              model.NewVariableLengthArrayType($typeSpecifier.v, defaultArraySize))
      }
      (
          value {
              $v = model.NewNamedDeclaration(
                  $IDENTIFIER.text,
                  model.NewVariableLengthArrayType($typeSpecifier.v, $value.v))
          }
      )? '>'
    | 'opaque' IDENTIFIER '[' value ']' {
          $v = model.NewNamedDeclaration(
              $IDENTIFIER.text,
              model.NewFixedLengthOpaqueType($value.v))
      }
    | 'opaque' IDENTIFIER '<' {
          $v = model.NewNamedDeclaration(
              $IDENTIFIER.text,
              model.NewVariableLengthOpaqueType(defaultArraySize))
      }
      (
          value {
              $v = model.NewNamedDeclaration(
                  $IDENTIFIER.text,
                  model.NewVariableLengthOpaqueType($value.v))
          }
      )? '>'
    | 'string' IDENTIFIER '<' {
          $v = model.NewNamedDeclaration(
              $IDENTIFIER.text,
              model.NewASCIIStringType(defaultArraySize))
      }
      (
          value {
              $v = model.NewNamedDeclaration(
                  $IDENTIFIER.text,
                  model.NewASCIIStringType($value.v))
          }
      )? '>'
    | 'utf8string' IDENTIFIER {
          $v = model.NewNamedDeclaration(
              $IDENTIFIER.text,
              model.NewUTF8StringType(defaultArraySize))
      }
    | typeSpecifier '*' IDENTIFIER {
          $v = model.NewNamedDeclaration(
              $IDENTIFIER.text,
              model.NewOptionalType($typeSpecifier.v))
      }
    ;

declaration returns [model.Declaration v]
    : namedDeclaration { $v = $namedDeclaration.v }
    | 'void' { $v = model.VoidDeclaration }
    ;

value returns [model.Value v]
    : constant { $v = model.NewConstantValue($constant.v) }
    | IDENTIFIER { $v = model.NewIdentifierValue($IDENTIFIER.text) }
    ;

constant returns [*big.Int v]
    : s=(DECIMAL_CONSTANT | HEXADECIMAL_CONSTANT | OCTAL_CONSTANT) {
          i := big.NewInt(0)
          if _, ok := i.SetString($s.text, 0); !ok {
              panic("Grammar guarantees integers are well-formed")
          }
          $v = i
      }
    ;

typeSpecifier returns [model.Type v]
    : { $v = model.IntType }
      (
          'unsigned' { $v = model.UnsignedIntType }
      )?
      'int'
    | { $v = model.HyperType }
      (
          'unsigned' { $v = model.UnsignedHyperType }
      )?
      'hyper'
    | 'float' { $v = model.FloatType }
    | 'double' { $v = model.DoubleType }
    | 'quadruple' { $v = model.QuadrupleType }
    | 'bool' { $v = model.BoolType }
    | enumTypeSpec { $v = $enumTypeSpec.v }
    | structTypeSpec { $v = $structTypeSpec.v }
    | unionTypeSpec { $v = $unionTypeSpec.v }
    | IDENTIFIER { $v = model.NewIdentifierType($IDENTIFIER.text) }
    ;

enumTypeSpec returns [model.Type v]
    : 'enum' enumBody { $v = $enumBody.v }
    ;

enumBody returns [model.Type v] locals [map[string]model.Value elements]
    : '{'
          IDENTIFIER '=' value {
              $elements = map[string]model.Value{
                  $IDENTIFIER.text: $value.v,
              }
          }
          (
              ',' IDENTIFIER '=' value {
                  // TODO: Throw an error when there are duplicates!
                  $elements[$IDENTIFIER.text] = $value.v
              }
          )*
      '}' { $v = model.NewEnumType($elements) }
    ;

structTypeSpec returns [model.Type v]
    : 'struct' structBody { $v = $structBody.v }
    ;

structBody returns [model.Type v] locals [[]model.NamedDeclaration declarations]
    : '{'
          namedDeclaration ';' {
              $declarations = append($declarations, $namedDeclaration.v)
          }
          (
              namedDeclaration ';' {
                  $declarations = append($declarations, $namedDeclaration.v)
              }
          )*
      '}' { $v = model.NewStructType($declarations) }
    ;

unionTypeSpec returns [model.Type v]
    : 'union' unionBody { $v = $unionBody.v }
    ;

unionBody returns [model.Type v] locals [[]model.CaseSpec cases, model.Declaration defaultCase]
    : 'switch' '(' namedDeclaration ')' '{'
          caseSpec { $cases = append($cases, $caseSpec.v) }
          (
              caseSpec { $cases = append($cases, $caseSpec.v) }
          )*
          (
              'default' ':' declaration ';' {
                  $defaultCase = $declaration.v
              }
          )?
      '}' { $v = model.NewUnionType($namedDeclaration.v, $cases, $defaultCase) }
    ;

caseSpec returns [model.CaseSpec v] locals [[]model.Value values]
    : 'case' value ':' { $values = append($values, $value.v) }
      (
          'case' value ':' { $values = append($values, $value.v) }
      )*
      declaration ';' { $v = model.NewCaseSpec($values, $declaration.v) }
    ;

constantDef returns [model.Definition v]
    : 'const' IDENTIFIER '=' constant ';' {
          $v = model.NewConstantDefinition($IDENTIFIER.text, $constant.v)
      }
    ;

typeDef returns [model.NamedDeclaration v]
    : 'typedef' namedDeclaration ';' {
          $v = $namedDeclaration.v
      }
    | 'enum' IDENTIFIER enumBody ';' {
          $v = model.NewNamedDeclaration($IDENTIFIER.text, $enumBody.v)
      }
    | 'struct' IDENTIFIER structBody ';' {
          $v = model.NewNamedDeclaration($IDENTIFIER.text, $structBody.v)
      }
    | 'union' IDENTIFIER unionBody ';' {
          $v = model.NewNamedDeclaration($IDENTIFIER.text, $unionBody.v)
      }
    ;

definition returns [model.Definition v]
    : typeDef {
          $v = model.NewTypeDefinition($typeDef.v)
      }
    | constantDef { $v = $constantDef.v }
    | programDef { $v = $programDef.v }
    ;

specification returns [[]model.Definition v]
    : (
          definition {
              $v = append($v, $definition.v)
          }
      )*
    ;

// RFC 5531: RPC: Remote Procedure Call Protocol specification Version 2
// https://tools.ietf.org/html/rfc5531

// Section 12.2: The RPC Language specification

programDef returns [model.Definition v] locals [[]model.Version versions]
    : 'program' IDENTIFIER '{'
      versionDef { $versions = append($versions, $versionDef.v) }
      (
          versionDef { $versions = append($versions, $versionDef.v) }
      )*
      '}' '=' constant ';' {
          $v = model.NewProgramDefinition($IDENTIFIER.text, $versions, $constant.v)
      }
    ;

versionDef returns [model.Version v] locals [[]model.Procedure procedures]
    : 'version' IDENTIFIER '{'
      procedureDef { $procedures = append($procedures, $procedureDef.v) }
      (
          procedureDef { $procedures = append($procedures, $procedureDef.v) }
      )*
      '}' '=' constant ';' {
          $v = model.NewVersion($IDENTIFIER.text, $procedures, $constant.v)
      }
    ;

procedureDef returns [model.Procedure v] locals [[]model.Type arguments]
    : procReturn IDENTIFIER '('
      procFirstArg {
         if $procFirstArg.v != nil {
             $arguments = append($arguments, $procFirstArg.v)
         }
      }
      (
          ',' typeSpecifier {
              $arguments = append($arguments, $typeSpecifier.v)
          }
      )* ')' '=' constant ';' {
          $v = model.NewProcedure($IDENTIFIER.text, $arguments, $procReturn.v, $constant.v)
      }
    ;

procReturn returns [model.Type v]
    : 'void'
    | typeSpecifier
    { $v = $typeSpecifier.v }
    ;

procFirstArg returns [model.Type v]
    : 'void'
    | typeSpecifier
    { $v = $typeSpecifier.v }
    ;

// RFC 4506: XDR: External Data Representation Standard

// Section 6.2: Lexical Notes

COMMENT : '/*' .*? '*/' -> skip ;

WHITE_SPACE : [ \t\n\r] + -> skip ;

C_DIRECTIVE : '%' ~[\r\n]* -> skip ;

IDENTIFIER : [a-zA-Z] [a-zA-Z0-9_]* ;

DECIMAL_CONSTANT : '-'? [1-9] [0-9]* ;

HEXADECIMAL_CONSTANT : '0x' [A-Fa-f0-9]+ ;

OCTAL_CONSTANT : '0' [0-7]* ;

STRING_CONSTANT : '"' ~["\\\u0000-\u001F]* '"' ;
