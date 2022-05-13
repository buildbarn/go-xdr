// Code generated from pkg/compiler/parser/XDR.g4 by ANTLR 4.10. DO NOT EDIT.

package parser // XDR

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import (
	"math"
	"math/big"

	"github.com/buildbarn/go-xdr/pkg/compiler/model"
)

var defaultArraySize = model.NewConstantValue(big.NewInt(math.MaxUint32))

// Suppress unused import errors
var _ = fmt.Printf

var (
	_ = strconv.Itoa
	_ = sync.Once{}
)

type XDRParser struct {
	*antlr.BaseParser
}

var xdrParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func xdrParserInit() {
	staticData := &xdrParserStaticData
	staticData.literalNames = []string{
		"", "'package'", "';'", "'import'", "'['", "']'", "'<'", "'>'", "'opaque'",
		"'string'", "'utf8string'", "'*'", "'void'", "'unsigned'", "'int'",
		"'hyper'", "'float'", "'double'", "'quadruple'", "'bool'", "'enum'",
		"'{'", "'='", "','", "'}'", "'struct'", "'union'", "'switch'", "'('",
		"')'", "'default'", "':'", "'case'", "'const'", "'typedef'", "'program'",
		"'version'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "COMMENT", "WHITE_SPACE", "C_DIRECTIVE", "IDENTIFIER", "DECIMAL_CONSTANT",
		"HEXADECIMAL_CONSTANT", "OCTAL_CONSTANT", "STRING_CONSTANT",
	}
	staticData.ruleNames = []string{
		"topLevel", "imports", "stringConstant", "namedDeclaration", "declaration",
		"value", "constant", "typeSpecifier", "enumTypeSpec", "enumBody", "structTypeSpec",
		"structBody", "unionTypeSpec", "unionBody", "caseSpec", "constantDef",
		"typeDef", "definition", "specification", "programDef", "versionDef",
		"procedureDef", "procReturn", "procFirstArg",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 44, 391, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 62, 8, 1, 10, 1, 12,
		1, 65, 9, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3,
		88, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 106, 8, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 116, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 127, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 3, 4, 134, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 141, 8, 5, 1, 6,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 3, 7, 149, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		3, 7, 155, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7,
		177, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 194, 8, 9, 10, 9, 12, 9, 197, 9, 9,
		1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 214, 8, 11, 10, 11, 12, 11, 217, 9,
		11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 236, 8, 13, 10,
		13, 12, 13, 239, 9, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13,
		247, 8, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 5, 14, 261, 8, 14, 10, 14, 12, 14, 264, 9, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 3, 16, 300, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 311, 8, 17, 1, 18, 1, 18, 1, 18, 5,
		18, 316, 8, 18, 10, 18, 12, 18, 319, 9, 18, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 329, 8, 19, 10, 19, 12, 19, 332, 9,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 5, 20, 348, 8, 20, 10, 20, 12, 20, 351, 9,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 368, 8, 21, 10, 21, 12, 21, 371,
		9, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1,
		22, 3, 22, 383, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 389, 8, 23, 1,
		23, 0, 0, 24, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
		32, 34, 36, 38, 40, 42, 44, 46, 0, 1, 1, 0, 41, 43, 406, 0, 48, 1, 0, 0,
		0, 2, 63, 1, 0, 0, 0, 4, 66, 1, 0, 0, 0, 6, 126, 1, 0, 0, 0, 8, 133, 1,
		0, 0, 0, 10, 140, 1, 0, 0, 0, 12, 142, 1, 0, 0, 0, 14, 176, 1, 0, 0, 0,
		16, 178, 1, 0, 0, 0, 18, 182, 1, 0, 0, 0, 20, 201, 1, 0, 0, 0, 22, 205,
		1, 0, 0, 0, 24, 221, 1, 0, 0, 0, 26, 225, 1, 0, 0, 0, 28, 251, 1, 0, 0,
		0, 30, 269, 1, 0, 0, 0, 32, 299, 1, 0, 0, 0, 34, 310, 1, 0, 0, 0, 36, 317,
		1, 0, 0, 0, 38, 320, 1, 0, 0, 0, 40, 339, 1, 0, 0, 0, 42, 358, 1, 0, 0,
		0, 44, 382, 1, 0, 0, 0, 46, 388, 1, 0, 0, 0, 48, 49, 5, 1, 0, 0, 49, 50,
		3, 4, 2, 0, 50, 51, 5, 2, 0, 0, 51, 52, 3, 2, 1, 0, 52, 53, 3, 36, 18,
		0, 53, 54, 5, 0, 0, 1, 54, 55, 6, 0, -1, 0, 55, 1, 1, 0, 0, 0, 56, 57,
		5, 3, 0, 0, 57, 58, 3, 4, 2, 0, 58, 59, 5, 2, 0, 0, 59, 60, 6, 1, -1, 0,
		60, 62, 1, 0, 0, 0, 61, 56, 1, 0, 0, 0, 62, 65, 1, 0, 0, 0, 63, 61, 1,
		0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 3, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 66,
		67, 5, 44, 0, 0, 67, 68, 6, 2, -1, 0, 68, 5, 1, 0, 0, 0, 69, 70, 3, 14,
		7, 0, 70, 71, 5, 40, 0, 0, 71, 72, 6, 3, -1, 0, 72, 127, 1, 0, 0, 0, 73,
		74, 3, 14, 7, 0, 74, 75, 5, 40, 0, 0, 75, 76, 5, 4, 0, 0, 76, 77, 3, 10,
		5, 0, 77, 78, 5, 5, 0, 0, 78, 79, 6, 3, -1, 0, 79, 127, 1, 0, 0, 0, 80,
		81, 3, 14, 7, 0, 81, 82, 5, 40, 0, 0, 82, 83, 5, 6, 0, 0, 83, 87, 6, 3,
		-1, 0, 84, 85, 3, 10, 5, 0, 85, 86, 6, 3, -1, 0, 86, 88, 1, 0, 0, 0, 87,
		84, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 90, 5, 7, 0,
		0, 90, 127, 1, 0, 0, 0, 91, 92, 5, 8, 0, 0, 92, 93, 5, 40, 0, 0, 93, 94,
		5, 4, 0, 0, 94, 95, 3, 10, 5, 0, 95, 96, 5, 5, 0, 0, 96, 97, 6, 3, -1,
		0, 97, 127, 1, 0, 0, 0, 98, 99, 5, 8, 0, 0, 99, 100, 5, 40, 0, 0, 100,
		101, 5, 6, 0, 0, 101, 105, 6, 3, -1, 0, 102, 103, 3, 10, 5, 0, 103, 104,
		6, 3, -1, 0, 104, 106, 1, 0, 0, 0, 105, 102, 1, 0, 0, 0, 105, 106, 1, 0,
		0, 0, 106, 107, 1, 0, 0, 0, 107, 127, 5, 7, 0, 0, 108, 109, 5, 9, 0, 0,
		109, 110, 5, 40, 0, 0, 110, 111, 5, 6, 0, 0, 111, 115, 6, 3, -1, 0, 112,
		113, 3, 10, 5, 0, 113, 114, 6, 3, -1, 0, 114, 116, 1, 0, 0, 0, 115, 112,
		1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 127, 5, 7,
		0, 0, 118, 119, 5, 10, 0, 0, 119, 120, 5, 40, 0, 0, 120, 127, 6, 3, -1,
		0, 121, 122, 3, 14, 7, 0, 122, 123, 5, 11, 0, 0, 123, 124, 5, 40, 0, 0,
		124, 125, 6, 3, -1, 0, 125, 127, 1, 0, 0, 0, 126, 69, 1, 0, 0, 0, 126,
		73, 1, 0, 0, 0, 126, 80, 1, 0, 0, 0, 126, 91, 1, 0, 0, 0, 126, 98, 1, 0,
		0, 0, 126, 108, 1, 0, 0, 0, 126, 118, 1, 0, 0, 0, 126, 121, 1, 0, 0, 0,
		127, 7, 1, 0, 0, 0, 128, 129, 3, 6, 3, 0, 129, 130, 6, 4, -1, 0, 130, 134,
		1, 0, 0, 0, 131, 132, 5, 12, 0, 0, 132, 134, 6, 4, -1, 0, 133, 128, 1,
		0, 0, 0, 133, 131, 1, 0, 0, 0, 134, 9, 1, 0, 0, 0, 135, 136, 3, 12, 6,
		0, 136, 137, 6, 5, -1, 0, 137, 141, 1, 0, 0, 0, 138, 139, 5, 40, 0, 0,
		139, 141, 6, 5, -1, 0, 140, 135, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 141,
		11, 1, 0, 0, 0, 142, 143, 7, 0, 0, 0, 143, 144, 6, 6, -1, 0, 144, 13, 1,
		0, 0, 0, 145, 148, 6, 7, -1, 0, 146, 147, 5, 13, 0, 0, 147, 149, 6, 7,
		-1, 0, 148, 146, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0,
		150, 177, 5, 14, 0, 0, 151, 154, 6, 7, -1, 0, 152, 153, 5, 13, 0, 0, 153,
		155, 6, 7, -1, 0, 154, 152, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 156,
		1, 0, 0, 0, 156, 177, 5, 15, 0, 0, 157, 158, 5, 16, 0, 0, 158, 177, 6,
		7, -1, 0, 159, 160, 5, 17, 0, 0, 160, 177, 6, 7, -1, 0, 161, 162, 5, 18,
		0, 0, 162, 177, 6, 7, -1, 0, 163, 164, 5, 19, 0, 0, 164, 177, 6, 7, -1,
		0, 165, 166, 3, 16, 8, 0, 166, 167, 6, 7, -1, 0, 167, 177, 1, 0, 0, 0,
		168, 169, 3, 20, 10, 0, 169, 170, 6, 7, -1, 0, 170, 177, 1, 0, 0, 0, 171,
		172, 3, 24, 12, 0, 172, 173, 6, 7, -1, 0, 173, 177, 1, 0, 0, 0, 174, 175,
		5, 40, 0, 0, 175, 177, 6, 7, -1, 0, 176, 145, 1, 0, 0, 0, 176, 151, 1,
		0, 0, 0, 176, 157, 1, 0, 0, 0, 176, 159, 1, 0, 0, 0, 176, 161, 1, 0, 0,
		0, 176, 163, 1, 0, 0, 0, 176, 165, 1, 0, 0, 0, 176, 168, 1, 0, 0, 0, 176,
		171, 1, 0, 0, 0, 176, 174, 1, 0, 0, 0, 177, 15, 1, 0, 0, 0, 178, 179, 5,
		20, 0, 0, 179, 180, 3, 18, 9, 0, 180, 181, 6, 8, -1, 0, 181, 17, 1, 0,
		0, 0, 182, 183, 5, 21, 0, 0, 183, 184, 5, 40, 0, 0, 184, 185, 5, 22, 0,
		0, 185, 186, 3, 10, 5, 0, 186, 195, 6, 9, -1, 0, 187, 188, 5, 23, 0, 0,
		188, 189, 5, 40, 0, 0, 189, 190, 5, 22, 0, 0, 190, 191, 3, 10, 5, 0, 191,
		192, 6, 9, -1, 0, 192, 194, 1, 0, 0, 0, 193, 187, 1, 0, 0, 0, 194, 197,
		1, 0, 0, 0, 195, 193, 1, 0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 198, 1, 0,
		0, 0, 197, 195, 1, 0, 0, 0, 198, 199, 5, 24, 0, 0, 199, 200, 6, 9, -1,
		0, 200, 19, 1, 0, 0, 0, 201, 202, 5, 25, 0, 0, 202, 203, 3, 22, 11, 0,
		203, 204, 6, 10, -1, 0, 204, 21, 1, 0, 0, 0, 205, 206, 5, 21, 0, 0, 206,
		207, 3, 6, 3, 0, 207, 208, 5, 2, 0, 0, 208, 215, 6, 11, -1, 0, 209, 210,
		3, 6, 3, 0, 210, 211, 5, 2, 0, 0, 211, 212, 6, 11, -1, 0, 212, 214, 1,
		0, 0, 0, 213, 209, 1, 0, 0, 0, 214, 217, 1, 0, 0, 0, 215, 213, 1, 0, 0,
		0, 215, 216, 1, 0, 0, 0, 216, 218, 1, 0, 0, 0, 217, 215, 1, 0, 0, 0, 218,
		219, 5, 24, 0, 0, 219, 220, 6, 11, -1, 0, 220, 23, 1, 0, 0, 0, 221, 222,
		5, 26, 0, 0, 222, 223, 3, 26, 13, 0, 223, 224, 6, 12, -1, 0, 224, 25, 1,
		0, 0, 0, 225, 226, 5, 27, 0, 0, 226, 227, 5, 28, 0, 0, 227, 228, 3, 6,
		3, 0, 228, 229, 5, 29, 0, 0, 229, 230, 5, 21, 0, 0, 230, 231, 3, 28, 14,
		0, 231, 237, 6, 13, -1, 0, 232, 233, 3, 28, 14, 0, 233, 234, 6, 13, -1,
		0, 234, 236, 1, 0, 0, 0, 235, 232, 1, 0, 0, 0, 236, 239, 1, 0, 0, 0, 237,
		235, 1, 0, 0, 0, 237, 238, 1, 0, 0, 0, 238, 246, 1, 0, 0, 0, 239, 237,
		1, 0, 0, 0, 240, 241, 5, 30, 0, 0, 241, 242, 5, 31, 0, 0, 242, 243, 3,
		8, 4, 0, 243, 244, 5, 2, 0, 0, 244, 245, 6, 13, -1, 0, 245, 247, 1, 0,
		0, 0, 246, 240, 1, 0, 0, 0, 246, 247, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0,
		248, 249, 5, 24, 0, 0, 249, 250, 6, 13, -1, 0, 250, 27, 1, 0, 0, 0, 251,
		252, 5, 32, 0, 0, 252, 253, 3, 10, 5, 0, 253, 254, 5, 31, 0, 0, 254, 262,
		6, 14, -1, 0, 255, 256, 5, 32, 0, 0, 256, 257, 3, 10, 5, 0, 257, 258, 5,
		31, 0, 0, 258, 259, 6, 14, -1, 0, 259, 261, 1, 0, 0, 0, 260, 255, 1, 0,
		0, 0, 261, 264, 1, 0, 0, 0, 262, 260, 1, 0, 0, 0, 262, 263, 1, 0, 0, 0,
		263, 265, 1, 0, 0, 0, 264, 262, 1, 0, 0, 0, 265, 266, 3, 8, 4, 0, 266,
		267, 5, 2, 0, 0, 267, 268, 6, 14, -1, 0, 268, 29, 1, 0, 0, 0, 269, 270,
		5, 33, 0, 0, 270, 271, 5, 40, 0, 0, 271, 272, 5, 22, 0, 0, 272, 273, 3,
		12, 6, 0, 273, 274, 5, 2, 0, 0, 274, 275, 6, 15, -1, 0, 275, 31, 1, 0,
		0, 0, 276, 277, 5, 34, 0, 0, 277, 278, 3, 6, 3, 0, 278, 279, 5, 2, 0, 0,
		279, 280, 6, 16, -1, 0, 280, 300, 1, 0, 0, 0, 281, 282, 5, 20, 0, 0, 282,
		283, 5, 40, 0, 0, 283, 284, 3, 18, 9, 0, 284, 285, 5, 2, 0, 0, 285, 286,
		6, 16, -1, 0, 286, 300, 1, 0, 0, 0, 287, 288, 5, 25, 0, 0, 288, 289, 5,
		40, 0, 0, 289, 290, 3, 22, 11, 0, 290, 291, 5, 2, 0, 0, 291, 292, 6, 16,
		-1, 0, 292, 300, 1, 0, 0, 0, 293, 294, 5, 26, 0, 0, 294, 295, 5, 40, 0,
		0, 295, 296, 3, 26, 13, 0, 296, 297, 5, 2, 0, 0, 297, 298, 6, 16, -1, 0,
		298, 300, 1, 0, 0, 0, 299, 276, 1, 0, 0, 0, 299, 281, 1, 0, 0, 0, 299,
		287, 1, 0, 0, 0, 299, 293, 1, 0, 0, 0, 300, 33, 1, 0, 0, 0, 301, 302, 3,
		32, 16, 0, 302, 303, 6, 17, -1, 0, 303, 311, 1, 0, 0, 0, 304, 305, 3, 30,
		15, 0, 305, 306, 6, 17, -1, 0, 306, 311, 1, 0, 0, 0, 307, 308, 3, 38, 19,
		0, 308, 309, 6, 17, -1, 0, 309, 311, 1, 0, 0, 0, 310, 301, 1, 0, 0, 0,
		310, 304, 1, 0, 0, 0, 310, 307, 1, 0, 0, 0, 311, 35, 1, 0, 0, 0, 312, 313,
		3, 34, 17, 0, 313, 314, 6, 18, -1, 0, 314, 316, 1, 0, 0, 0, 315, 312, 1,
		0, 0, 0, 316, 319, 1, 0, 0, 0, 317, 315, 1, 0, 0, 0, 317, 318, 1, 0, 0,
		0, 318, 37, 1, 0, 0, 0, 319, 317, 1, 0, 0, 0, 320, 321, 5, 35, 0, 0, 321,
		322, 5, 40, 0, 0, 322, 323, 5, 21, 0, 0, 323, 324, 3, 40, 20, 0, 324, 330,
		6, 19, -1, 0, 325, 326, 3, 40, 20, 0, 326, 327, 6, 19, -1, 0, 327, 329,
		1, 0, 0, 0, 328, 325, 1, 0, 0, 0, 329, 332, 1, 0, 0, 0, 330, 328, 1, 0,
		0, 0, 330, 331, 1, 0, 0, 0, 331, 333, 1, 0, 0, 0, 332, 330, 1, 0, 0, 0,
		333, 334, 5, 24, 0, 0, 334, 335, 5, 22, 0, 0, 335, 336, 3, 12, 6, 0, 336,
		337, 5, 2, 0, 0, 337, 338, 6, 19, -1, 0, 338, 39, 1, 0, 0, 0, 339, 340,
		5, 36, 0, 0, 340, 341, 5, 40, 0, 0, 341, 342, 5, 21, 0, 0, 342, 343, 3,
		42, 21, 0, 343, 349, 6, 20, -1, 0, 344, 345, 3, 42, 21, 0, 345, 346, 6,
		20, -1, 0, 346, 348, 1, 0, 0, 0, 347, 344, 1, 0, 0, 0, 348, 351, 1, 0,
		0, 0, 349, 347, 1, 0, 0, 0, 349, 350, 1, 0, 0, 0, 350, 352, 1, 0, 0, 0,
		351, 349, 1, 0, 0, 0, 352, 353, 5, 24, 0, 0, 353, 354, 5, 22, 0, 0, 354,
		355, 3, 12, 6, 0, 355, 356, 5, 2, 0, 0, 356, 357, 6, 20, -1, 0, 357, 41,
		1, 0, 0, 0, 358, 359, 3, 44, 22, 0, 359, 360, 5, 40, 0, 0, 360, 361, 5,
		28, 0, 0, 361, 362, 3, 46, 23, 0, 362, 369, 6, 21, -1, 0, 363, 364, 5,
		23, 0, 0, 364, 365, 3, 14, 7, 0, 365, 366, 6, 21, -1, 0, 366, 368, 1, 0,
		0, 0, 367, 363, 1, 0, 0, 0, 368, 371, 1, 0, 0, 0, 369, 367, 1, 0, 0, 0,
		369, 370, 1, 0, 0, 0, 370, 372, 1, 0, 0, 0, 371, 369, 1, 0, 0, 0, 372,
		373, 5, 29, 0, 0, 373, 374, 5, 22, 0, 0, 374, 375, 3, 12, 6, 0, 375, 376,
		5, 2, 0, 0, 376, 377, 6, 21, -1, 0, 377, 43, 1, 0, 0, 0, 378, 383, 5, 12,
		0, 0, 379, 380, 3, 14, 7, 0, 380, 381, 6, 22, -1, 0, 381, 383, 1, 0, 0,
		0, 382, 378, 1, 0, 0, 0, 382, 379, 1, 0, 0, 0, 383, 45, 1, 0, 0, 0, 384,
		389, 5, 12, 0, 0, 385, 386, 3, 14, 7, 0, 386, 387, 6, 23, -1, 0, 387, 389,
		1, 0, 0, 0, 388, 384, 1, 0, 0, 0, 388, 385, 1, 0, 0, 0, 389, 47, 1, 0,
		0, 0, 23, 63, 87, 105, 115, 126, 133, 140, 148, 154, 176, 195, 215, 237,
		246, 262, 299, 310, 317, 330, 349, 369, 382, 388,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// XDRParserInit initializes any static state used to implement XDRParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewXDRParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func XDRParserInit() {
	staticData := &xdrParserStaticData
	staticData.once.Do(xdrParserInit)
}

// NewXDRParser produces a new parser instance for the optional input antlr.TokenStream.
func NewXDRParser(input antlr.TokenStream) *XDRParser {
	XDRParserInit()
	this := new(XDRParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &xdrParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "XDR.g4"

	return this
}

// XDRParser tokens.
const (
	XDRParserEOF                  = antlr.TokenEOF
	XDRParserT__0                 = 1
	XDRParserT__1                 = 2
	XDRParserT__2                 = 3
	XDRParserT__3                 = 4
	XDRParserT__4                 = 5
	XDRParserT__5                 = 6
	XDRParserT__6                 = 7
	XDRParserT__7                 = 8
	XDRParserT__8                 = 9
	XDRParserT__9                 = 10
	XDRParserT__10                = 11
	XDRParserT__11                = 12
	XDRParserT__12                = 13
	XDRParserT__13                = 14
	XDRParserT__14                = 15
	XDRParserT__15                = 16
	XDRParserT__16                = 17
	XDRParserT__17                = 18
	XDRParserT__18                = 19
	XDRParserT__19                = 20
	XDRParserT__20                = 21
	XDRParserT__21                = 22
	XDRParserT__22                = 23
	XDRParserT__23                = 24
	XDRParserT__24                = 25
	XDRParserT__25                = 26
	XDRParserT__26                = 27
	XDRParserT__27                = 28
	XDRParserT__28                = 29
	XDRParserT__29                = 30
	XDRParserT__30                = 31
	XDRParserT__31                = 32
	XDRParserT__32                = 33
	XDRParserT__33                = 34
	XDRParserT__34                = 35
	XDRParserT__35                = 36
	XDRParserCOMMENT              = 37
	XDRParserWHITE_SPACE          = 38
	XDRParserC_DIRECTIVE          = 39
	XDRParserIDENTIFIER           = 40
	XDRParserDECIMAL_CONSTANT     = 41
	XDRParserHEXADECIMAL_CONSTANT = 42
	XDRParserOCTAL_CONSTANT       = 43
	XDRParserSTRING_CONSTANT      = 44
)

// XDRParser rules.
const (
	XDRParserRULE_topLevel         = 0
	XDRParserRULE_imports          = 1
	XDRParserRULE_stringConstant   = 2
	XDRParserRULE_namedDeclaration = 3
	XDRParserRULE_declaration      = 4
	XDRParserRULE_value            = 5
	XDRParserRULE_constant         = 6
	XDRParserRULE_typeSpecifier    = 7
	XDRParserRULE_enumTypeSpec     = 8
	XDRParserRULE_enumBody         = 9
	XDRParserRULE_structTypeSpec   = 10
	XDRParserRULE_structBody       = 11
	XDRParserRULE_unionTypeSpec    = 12
	XDRParserRULE_unionBody        = 13
	XDRParserRULE_caseSpec         = 14
	XDRParserRULE_constantDef      = 15
	XDRParserRULE_typeDef          = 16
	XDRParserRULE_definition       = 17
	XDRParserRULE_specification    = 18
	XDRParserRULE_programDef       = 19
	XDRParserRULE_versionDef       = 20
	XDRParserRULE_procedureDef     = 21
	XDRParserRULE_procReturn       = 22
	XDRParserRULE_procFirstArg     = 23
)

// ITopLevelContext is an interface to support dynamic dispatch.
type ITopLevelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_stringConstant returns the _stringConstant rule contexts.
	Get_stringConstant() IStringConstantContext

	// Get_imports returns the _imports rule contexts.
	Get_imports() IImportsContext

	// Get_specification returns the _specification rule contexts.
	Get_specification() ISpecificationContext

	// Set_stringConstant sets the _stringConstant rule contexts.
	Set_stringConstant(IStringConstantContext)

	// Set_imports sets the _imports rule contexts.
	Set_imports(IImportsContext)

	// Set_specification sets the _specification rule contexts.
	Set_specification(ISpecificationContext)

	// GetV returns the v attribute.
	GetV() model.TopLevel

	// SetV sets the v attribute.
	SetV(model.TopLevel)

	// IsTopLevelContext differentiates from other interfaces.
	IsTopLevelContext()
}

type TopLevelContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	v               model.TopLevel
	_stringConstant IStringConstantContext
	_imports        IImportsContext
	_specification  ISpecificationContext
}

func NewEmptyTopLevelContext() *TopLevelContext {
	p := new(TopLevelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XDRParserRULE_topLevel
	return p
}

func (*TopLevelContext) IsTopLevelContext() {}

func NewTopLevelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopLevelContext {
	p := new(TopLevelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XDRParserRULE_topLevel

	return p
}

func (s *TopLevelContext) GetParser() antlr.Parser { return s.parser }

func (s *TopLevelContext) Get_stringConstant() IStringConstantContext { return s._stringConstant }

func (s *TopLevelContext) Get_imports() IImportsContext { return s._imports }

func (s *TopLevelContext) Get_specification() ISpecificationContext { return s._specification }

func (s *TopLevelContext) Set_stringConstant(v IStringConstantContext) { s._stringConstant = v }

func (s *TopLevelContext) Set_imports(v IImportsContext) { s._imports = v }

func (s *TopLevelContext) Set_specification(v ISpecificationContext) { s._specification = v }

func (s *TopLevelContext) GetV() model.TopLevel { return s.v }

func (s *TopLevelContext) SetV(v model.TopLevel) { s.v = v }

func (s *TopLevelContext) StringConstant() IStringConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringConstantContext)
}

func (s *TopLevelContext) Imports() IImportsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportsContext)
}

func (s *TopLevelContext) Specification() ISpecificationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpecificationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpecificationContext)
}

func (s *TopLevelContext) EOF() antlr.TerminalNode {
	return s.GetToken(XDRParserEOF, 0)
}

func (s *TopLevelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopLevelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.EnterTopLevel(s)
	}
}

func (s *TopLevelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.ExitTopLevel(s)
	}
}

func (p *XDRParser) TopLevel() (localctx ITopLevelContext) {
	this := p
	_ = this

	localctx = NewTopLevelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, XDRParserRULE_topLevel)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Match(XDRParserT__0)
	}
	{
		p.SetState(49)

		_x := p.StringConstant()

		localctx.(*TopLevelContext)._stringConstant = _x
	}
	{
		p.SetState(50)
		p.Match(XDRParserT__1)
	}
	{
		p.SetState(51)

		_x := p.Imports()

		localctx.(*TopLevelContext)._imports = _x
	}
	{
		p.SetState(52)

		_x := p.Specification()

		localctx.(*TopLevelContext)._specification = _x
	}
	{
		p.SetState(53)
		p.Match(XDRParserEOF)
	}

	localctx.(*TopLevelContext).v = model.NewTopLevel(localctx.(*TopLevelContext).Get_stringConstant().GetV(), localctx.(*TopLevelContext).Get_imports().GetV(), localctx.(*TopLevelContext).Get_specification().GetV())

	return localctx
}

// IImportsContext is an interface to support dynamic dispatch.
type IImportsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_stringConstant returns the _stringConstant rule contexts.
	Get_stringConstant() IStringConstantContext

	// Set_stringConstant sets the _stringConstant rule contexts.
	Set_stringConstant(IStringConstantContext)

	// GetV returns the v attribute.
	GetV() []string

	// SetV sets the v attribute.
	SetV([]string)

	// IsImportsContext differentiates from other interfaces.
	IsImportsContext()
}

type ImportsContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	v               []string
	_stringConstant IStringConstantContext
}

func NewEmptyImportsContext() *ImportsContext {
	p := new(ImportsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XDRParserRULE_imports
	return p
}

func (*ImportsContext) IsImportsContext() {}

func NewImportsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportsContext {
	p := new(ImportsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XDRParserRULE_imports

	return p
}

func (s *ImportsContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportsContext) Get_stringConstant() IStringConstantContext { return s._stringConstant }

func (s *ImportsContext) Set_stringConstant(v IStringConstantContext) { s._stringConstant = v }

func (s *ImportsContext) GetV() []string { return s.v }

func (s *ImportsContext) SetV(v []string) { s.v = v }

func (s *ImportsContext) AllStringConstant() []IStringConstantContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStringConstantContext); ok {
			len++
		}
	}

	tst := make([]IStringConstantContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStringConstantContext); ok {
			tst[i] = t.(IStringConstantContext)
			i++
		}
	}

	return tst
}

func (s *ImportsContext) StringConstant(i int) IStringConstantContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringConstantContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringConstantContext)
}

func (s *ImportsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.EnterImports(s)
	}
}

func (s *ImportsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.ExitImports(s)
	}
}

func (p *XDRParser) Imports() (localctx IImportsContext) {
	this := p
	_ = this

	localctx = NewImportsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, XDRParserRULE_imports)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == XDRParserT__2 {
		{
			p.SetState(56)
			p.Match(XDRParserT__2)
		}
		{
			p.SetState(57)

			_x := p.StringConstant()

			localctx.(*ImportsContext)._stringConstant = _x
		}
		{
			p.SetState(58)
			p.Match(XDRParserT__1)
		}

		localctx.(*ImportsContext).v = append(localctx.(*ImportsContext).v, localctx.(*ImportsContext).Get_stringConstant().GetV())

		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStringConstantContext is an interface to support dynamic dispatch.
type IStringConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_STRING_CONSTANT returns the _STRING_CONSTANT token.
	Get_STRING_CONSTANT() antlr.Token

	// Set_STRING_CONSTANT sets the _STRING_CONSTANT token.
	Set_STRING_CONSTANT(antlr.Token)

	// GetV returns the v attribute.
	GetV() string

	// SetV sets the v attribute.
	SetV(string)

	// IsStringConstantContext differentiates from other interfaces.
	IsStringConstantContext()
}

type StringConstantContext struct {
	*antlr.BaseParserRuleContext
	parser           antlr.Parser
	v                string
	_STRING_CONSTANT antlr.Token
}

func NewEmptyStringConstantContext() *StringConstantContext {
	p := new(StringConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XDRParserRULE_stringConstant
	return p
}

func (*StringConstantContext) IsStringConstantContext() {}

func NewStringConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringConstantContext {
	p := new(StringConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XDRParserRULE_stringConstant

	return p
}

func (s *StringConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *StringConstantContext) Get_STRING_CONSTANT() antlr.Token { return s._STRING_CONSTANT }

func (s *StringConstantContext) Set_STRING_CONSTANT(v antlr.Token) { s._STRING_CONSTANT = v }

func (s *StringConstantContext) GetV() string { return s.v }

func (s *StringConstantContext) SetV(v string) { s.v = v }

func (s *StringConstantContext) STRING_CONSTANT() antlr.TerminalNode {
	return s.GetToken(XDRParserSTRING_CONSTANT, 0)
}

func (s *StringConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.EnterStringConstant(s)
	}
}

func (s *StringConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.ExitStringConstant(s)
	}
}

func (p *XDRParser) StringConstant() (localctx IStringConstantContext) {
	this := p
	_ = this

	localctx = NewStringConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, XDRParserRULE_stringConstant)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)

		_m := p.Match(XDRParserSTRING_CONSTANT)

		localctx.(*StringConstantContext)._STRING_CONSTANT = _m
	}

	s := (func() string {
		if localctx.(*StringConstantContext).Get_STRING_CONSTANT() == nil {
			return ""
		} else {
			return localctx.(*StringConstantContext).Get_STRING_CONSTANT().GetText()
		}
	}())
	localctx.(*StringConstantContext).v = s[1 : len(s)-1]

	return localctx
}

// INamedDeclarationContext is an interface to support dynamic dispatch.
type INamedDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_IDENTIFIER returns the _IDENTIFIER token.
	Get_IDENTIFIER() antlr.Token

	// Set_IDENTIFIER sets the _IDENTIFIER token.
	Set_IDENTIFIER(antlr.Token)

	// Get_typeSpecifier returns the _typeSpecifier rule contexts.
	Get_typeSpecifier() ITypeSpecifierContext

	// Get_value returns the _value rule contexts.
	Get_value() IValueContext

	// Set_typeSpecifier sets the _typeSpecifier rule contexts.
	Set_typeSpecifier(ITypeSpecifierContext)

	// Set_value sets the _value rule contexts.
	Set_value(IValueContext)

	// GetV returns the v attribute.
	GetV() model.NamedDeclaration

	// SetV sets the v attribute.
	SetV(model.NamedDeclaration)

	// IsNamedDeclarationContext differentiates from other interfaces.
	IsNamedDeclarationContext()
}

type NamedDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	v              model.NamedDeclaration
	_typeSpecifier ITypeSpecifierContext
	_IDENTIFIER    antlr.Token
	_value         IValueContext
}

func NewEmptyNamedDeclarationContext() *NamedDeclarationContext {
	p := new(NamedDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XDRParserRULE_namedDeclaration
	return p
}

func (*NamedDeclarationContext) IsNamedDeclarationContext() {}

func NewNamedDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamedDeclarationContext {
	p := new(NamedDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XDRParserRULE_namedDeclaration

	return p
}

func (s *NamedDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *NamedDeclarationContext) Get_IDENTIFIER() antlr.Token { return s._IDENTIFIER }

func (s *NamedDeclarationContext) Set_IDENTIFIER(v antlr.Token) { s._IDENTIFIER = v }

func (s *NamedDeclarationContext) Get_typeSpecifier() ITypeSpecifierContext { return s._typeSpecifier }

func (s *NamedDeclarationContext) Get_value() IValueContext { return s._value }

func (s *NamedDeclarationContext) Set_typeSpecifier(v ITypeSpecifierContext) { s._typeSpecifier = v }

func (s *NamedDeclarationContext) Set_value(v IValueContext) { s._value = v }

func (s *NamedDeclarationContext) GetV() model.NamedDeclaration { return s.v }

func (s *NamedDeclarationContext) SetV(v model.NamedDeclaration) { s.v = v }

func (s *NamedDeclarationContext) TypeSpecifier() ITypeSpecifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeSpecifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *NamedDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(XDRParserIDENTIFIER, 0)
}

func (s *NamedDeclarationContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *NamedDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamedDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.EnterNamedDeclaration(s)
	}
}

func (s *NamedDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.ExitNamedDeclaration(s)
	}
}

func (p *XDRParser) NamedDeclaration() (localctx INamedDeclarationContext) {
	this := p
	_ = this

	localctx = NewNamedDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, XDRParserRULE_namedDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(69)

			_x := p.TypeSpecifier()

			localctx.(*NamedDeclarationContext)._typeSpecifier = _x
		}
		{
			p.SetState(70)

			_m := p.Match(XDRParserIDENTIFIER)

			localctx.(*NamedDeclarationContext)._IDENTIFIER = _m
		}

		localctx.(*NamedDeclarationContext).v = model.NewNamedDeclaration((func() string {
			if localctx.(*NamedDeclarationContext).Get_IDENTIFIER() == nil {
				return ""
			} else {
				return localctx.(*NamedDeclarationContext).Get_IDENTIFIER().GetText()
			}
		}()), localctx.(*NamedDeclarationContext).Get_typeSpecifier().GetV())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(73)

			_x := p.TypeSpecifier()

			localctx.(*NamedDeclarationContext)._typeSpecifier = _x
		}
		{
			p.SetState(74)

			_m := p.Match(XDRParserIDENTIFIER)

			localctx.(*NamedDeclarationContext)._IDENTIFIER = _m
		}
		{
			p.SetState(75)
			p.Match(XDRParserT__3)
		}
		{
			p.SetState(76)

			_x := p.Value()

			localctx.(*NamedDeclarationContext)._value = _x
		}
		{
			p.SetState(77)
			p.Match(XDRParserT__4)
		}

		localctx.(*NamedDeclarationContext).v = model.NewNamedDeclaration(
			(func() string {
				if localctx.(*NamedDeclarationContext).Get_IDENTIFIER() == nil {
					return ""
				} else {
					return localctx.(*NamedDeclarationContext).Get_IDENTIFIER().GetText()
				}
			}()),
			model.NewFixedLengthArrayType(localctx.(*NamedDeclarationContext).Get_typeSpecifier().GetV(), localctx.(*NamedDeclarationContext).Get_value().GetV()))

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(80)

			_x := p.TypeSpecifier()

			localctx.(*NamedDeclarationContext)._typeSpecifier = _x
		}
		{
			p.SetState(81)

			_m := p.Match(XDRParserIDENTIFIER)

			localctx.(*NamedDeclarationContext)._IDENTIFIER = _m
		}
		{
			p.SetState(82)
			p.Match(XDRParserT__5)
		}

		localctx.(*NamedDeclarationContext).v = model.NewNamedDeclaration(
			(func() string {
				if localctx.(*NamedDeclarationContext).Get_IDENTIFIER() == nil {
					return ""
				} else {
					return localctx.(*NamedDeclarationContext).Get_IDENTIFIER().GetText()
				}
			}()),
			model.NewVariableLengthArrayType(localctx.(*NamedDeclarationContext).Get_typeSpecifier().GetV(), defaultArraySize))

		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(XDRParserIDENTIFIER-40))|(1<<(XDRParserDECIMAL_CONSTANT-40))|(1<<(XDRParserHEXADECIMAL_CONSTANT-40))|(1<<(XDRParserOCTAL_CONSTANT-40)))) != 0 {
			{
				p.SetState(84)

				_x := p.Value()

				localctx.(*NamedDeclarationContext)._value = _x
			}

			localctx.(*NamedDeclarationContext).v = model.NewNamedDeclaration(
				(func() string {
					if localctx.(*NamedDeclarationContext).Get_IDENTIFIER() == nil {
						return ""
					} else {
						return localctx.(*NamedDeclarationContext).Get_IDENTIFIER().GetText()
					}
				}()),
				model.NewVariableLengthArrayType(localctx.(*NamedDeclarationContext).Get_typeSpecifier().GetV(), localctx.(*NamedDeclarationContext).Get_value().GetV()))

		}
		{
			p.SetState(89)
			p.Match(XDRParserT__6)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(91)
			p.Match(XDRParserT__7)
		}
		{
			p.SetState(92)

			_m := p.Match(XDRParserIDENTIFIER)

			localctx.(*NamedDeclarationContext)._IDENTIFIER = _m
		}
		{
			p.SetState(93)
			p.Match(XDRParserT__3)
		}
		{
			p.SetState(94)

			_x := p.Value()

			localctx.(*NamedDeclarationContext)._value = _x
		}
		{
			p.SetState(95)
			p.Match(XDRParserT__4)
		}

		localctx.(*NamedDeclarationContext).v = model.NewNamedDeclaration(
			(func() string {
				if localctx.(*NamedDeclarationContext).Get_IDENTIFIER() == nil {
					return ""
				} else {
					return localctx.(*NamedDeclarationContext).Get_IDENTIFIER().GetText()
				}
			}()),
			model.NewFixedLengthOpaqueType(localctx.(*NamedDeclarationContext).Get_value().GetV()))

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(98)
			p.Match(XDRParserT__7)
		}
		{
			p.SetState(99)

			_m := p.Match(XDRParserIDENTIFIER)

			localctx.(*NamedDeclarationContext)._IDENTIFIER = _m
		}
		{
			p.SetState(100)
			p.Match(XDRParserT__5)
		}

		localctx.(*NamedDeclarationContext).v = model.NewNamedDeclaration(
			(func() string {
				if localctx.(*NamedDeclarationContext).Get_IDENTIFIER() == nil {
					return ""
				} else {
					return localctx.(*NamedDeclarationContext).Get_IDENTIFIER().GetText()
				}
			}()),
			model.NewVariableLengthOpaqueType(defaultArraySize))

		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(XDRParserIDENTIFIER-40))|(1<<(XDRParserDECIMAL_CONSTANT-40))|(1<<(XDRParserHEXADECIMAL_CONSTANT-40))|(1<<(XDRParserOCTAL_CONSTANT-40)))) != 0 {
			{
				p.SetState(102)

				_x := p.Value()

				localctx.(*NamedDeclarationContext)._value = _x
			}

			localctx.(*NamedDeclarationContext).v = model.NewNamedDeclaration(
				(func() string {
					if localctx.(*NamedDeclarationContext).Get_IDENTIFIER() == nil {
						return ""
					} else {
						return localctx.(*NamedDeclarationContext).Get_IDENTIFIER().GetText()
					}
				}()),
				model.NewVariableLengthOpaqueType(localctx.(*NamedDeclarationContext).Get_value().GetV()))

		}
		{
			p.SetState(107)
			p.Match(XDRParserT__6)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(108)
			p.Match(XDRParserT__8)
		}
		{
			p.SetState(109)

			_m := p.Match(XDRParserIDENTIFIER)

			localctx.(*NamedDeclarationContext)._IDENTIFIER = _m
		}
		{
			p.SetState(110)
			p.Match(XDRParserT__5)
		}

		localctx.(*NamedDeclarationContext).v = model.NewNamedDeclaration(
			(func() string {
				if localctx.(*NamedDeclarationContext).Get_IDENTIFIER() == nil {
					return ""
				} else {
					return localctx.(*NamedDeclarationContext).Get_IDENTIFIER().GetText()
				}
			}()),
			model.NewASCIIStringType(defaultArraySize))

		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(XDRParserIDENTIFIER-40))|(1<<(XDRParserDECIMAL_CONSTANT-40))|(1<<(XDRParserHEXADECIMAL_CONSTANT-40))|(1<<(XDRParserOCTAL_CONSTANT-40)))) != 0 {
			{
				p.SetState(112)

				_x := p.Value()

				localctx.(*NamedDeclarationContext)._value = _x
			}

			localctx.(*NamedDeclarationContext).v = model.NewNamedDeclaration(
				(func() string {
					if localctx.(*NamedDeclarationContext).Get_IDENTIFIER() == nil {
						return ""
					} else {
						return localctx.(*NamedDeclarationContext).Get_IDENTIFIER().GetText()
					}
				}()),
				model.NewASCIIStringType(localctx.(*NamedDeclarationContext).Get_value().GetV()))

		}
		{
			p.SetState(117)
			p.Match(XDRParserT__6)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(118)
			p.Match(XDRParserT__9)
		}
		{
			p.SetState(119)

			_m := p.Match(XDRParserIDENTIFIER)

			localctx.(*NamedDeclarationContext)._IDENTIFIER = _m
		}

		localctx.(*NamedDeclarationContext).v = model.NewNamedDeclaration(
			(func() string {
				if localctx.(*NamedDeclarationContext).Get_IDENTIFIER() == nil {
					return ""
				} else {
					return localctx.(*NamedDeclarationContext).Get_IDENTIFIER().GetText()
				}
			}()),
			model.NewUTF8StringType(defaultArraySize))

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(121)

			_x := p.TypeSpecifier()

			localctx.(*NamedDeclarationContext)._typeSpecifier = _x
		}
		{
			p.SetState(122)
			p.Match(XDRParserT__10)
		}
		{
			p.SetState(123)

			_m := p.Match(XDRParserIDENTIFIER)

			localctx.(*NamedDeclarationContext)._IDENTIFIER = _m
		}

		localctx.(*NamedDeclarationContext).v = model.NewNamedDeclaration(
			(func() string {
				if localctx.(*NamedDeclarationContext).Get_IDENTIFIER() == nil {
					return ""
				} else {
					return localctx.(*NamedDeclarationContext).Get_IDENTIFIER().GetText()
				}
			}()),
			model.NewOptionalType(localctx.(*NamedDeclarationContext).Get_typeSpecifier().GetV()))

	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_namedDeclaration returns the _namedDeclaration rule contexts.
	Get_namedDeclaration() INamedDeclarationContext

	// Set_namedDeclaration sets the _namedDeclaration rule contexts.
	Set_namedDeclaration(INamedDeclarationContext)

	// GetV returns the v attribute.
	GetV() model.Declaration

	// SetV sets the v attribute.
	SetV(model.Declaration)

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	v                 model.Declaration
	_namedDeclaration INamedDeclarationContext
}

func NewEmptyDeclarationContext() *DeclarationContext {
	p := new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XDRParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	p := new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XDRParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) Get_namedDeclaration() INamedDeclarationContext {
	return s._namedDeclaration
}

func (s *DeclarationContext) Set_namedDeclaration(v INamedDeclarationContext) {
	s._namedDeclaration = v
}

func (s *DeclarationContext) GetV() model.Declaration { return s.v }

func (s *DeclarationContext) SetV(v model.Declaration) { s.v = v }

func (s *DeclarationContext) NamedDeclaration() INamedDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamedDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamedDeclarationContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (p *XDRParser) Declaration() (localctx IDeclarationContext) {
	this := p
	_ = this

	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, XDRParserRULE_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(133)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case XDRParserT__7, XDRParserT__8, XDRParserT__9, XDRParserT__12, XDRParserT__13, XDRParserT__14, XDRParserT__15, XDRParserT__16, XDRParserT__17, XDRParserT__18, XDRParserT__19, XDRParserT__24, XDRParserT__25, XDRParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(128)

			_x := p.NamedDeclaration()

			localctx.(*DeclarationContext)._namedDeclaration = _x
		}
		localctx.(*DeclarationContext).v = localctx.(*DeclarationContext).Get_namedDeclaration().GetV()

	case XDRParserT__11:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(131)
			p.Match(XDRParserT__11)
		}
		localctx.(*DeclarationContext).v = model.VoidDeclaration

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_IDENTIFIER returns the _IDENTIFIER token.
	Get_IDENTIFIER() antlr.Token

	// Set_IDENTIFIER sets the _IDENTIFIER token.
	Set_IDENTIFIER(antlr.Token)

	// Get_constant returns the _constant rule contexts.
	Get_constant() IConstantContext

	// Set_constant sets the _constant rule contexts.
	Set_constant(IConstantContext)

	// GetV returns the v attribute.
	GetV() model.Value

	// SetV sets the v attribute.
	SetV(model.Value)

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	v           model.Value
	_constant   IConstantContext
	_IDENTIFIER antlr.Token
}

func NewEmptyValueContext() *ValueContext {
	p := new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XDRParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	p := new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XDRParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) Get_IDENTIFIER() antlr.Token { return s._IDENTIFIER }

func (s *ValueContext) Set_IDENTIFIER(v antlr.Token) { s._IDENTIFIER = v }

func (s *ValueContext) Get_constant() IConstantContext { return s._constant }

func (s *ValueContext) Set_constant(v IConstantContext) { s._constant = v }

func (s *ValueContext) GetV() model.Value { return s.v }

func (s *ValueContext) SetV(v model.Value) { s.v = v }

func (s *ValueContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ValueContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(XDRParserIDENTIFIER, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *XDRParser) Value() (localctx IValueContext) {
	this := p
	_ = this

	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, XDRParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(140)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case XDRParserDECIMAL_CONSTANT, XDRParserHEXADECIMAL_CONSTANT, XDRParserOCTAL_CONSTANT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(135)

			_x := p.Constant()

			localctx.(*ValueContext)._constant = _x
		}
		localctx.(*ValueContext).v = model.NewConstantValue(localctx.(*ValueContext).Get_constant().GetV())

	case XDRParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(138)

			_m := p.Match(XDRParserIDENTIFIER)

			localctx.(*ValueContext)._IDENTIFIER = _m
		}
		localctx.(*ValueContext).v = model.NewIdentifierValue((func() string {
			if localctx.(*ValueContext).Get_IDENTIFIER() == nil {
				return ""
			} else {
				return localctx.(*ValueContext).Get_IDENTIFIER().GetText()
			}
		}()))

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetS returns the s token.
	GetS() antlr.Token

	// SetS sets the s token.
	SetS(antlr.Token)

	// GetV returns the v attribute.
	GetV() *big.Int

	// SetV sets the v attribute.
	SetV(*big.Int)

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	v      *big.Int
	s      antlr.Token
}

func NewEmptyConstantContext() *ConstantContext {
	p := new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XDRParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	p := new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XDRParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) GetS() antlr.Token { return s.s }

func (s *ConstantContext) SetS(v antlr.Token) { s.s = v }

func (s *ConstantContext) GetV() *big.Int { return s.v }

func (s *ConstantContext) SetV(v *big.Int) { s.v = v }

func (s *ConstantContext) DECIMAL_CONSTANT() antlr.TerminalNode {
	return s.GetToken(XDRParserDECIMAL_CONSTANT, 0)
}

func (s *ConstantContext) HEXADECIMAL_CONSTANT() antlr.TerminalNode {
	return s.GetToken(XDRParserHEXADECIMAL_CONSTANT, 0)
}

func (s *ConstantContext) OCTAL_CONSTANT() antlr.TerminalNode {
	return s.GetToken(XDRParserOCTAL_CONSTANT, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (p *XDRParser) Constant() (localctx IConstantContext) {
	this := p
	_ = this

	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, XDRParserRULE_constant)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)

		_lt := p.GetTokenStream().LT(1)

		localctx.(*ConstantContext).s = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(XDRParserDECIMAL_CONSTANT-41))|(1<<(XDRParserHEXADECIMAL_CONSTANT-41))|(1<<(XDRParserOCTAL_CONSTANT-41)))) != 0) {
			_ri := p.GetErrorHandler().RecoverInline(p)

			localctx.(*ConstantContext).s = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	i := big.NewInt(0)
	if _, ok := i.SetString((func() string {
		if localctx.(*ConstantContext).GetS() == nil {
			return ""
		} else {
			return localctx.(*ConstantContext).GetS().GetText()
		}
	}()), 0); !ok {
		panic("Grammar guarantees integers are well-formed")
	}
	localctx.(*ConstantContext).v = i

	return localctx
}

// ITypeSpecifierContext is an interface to support dynamic dispatch.
type ITypeSpecifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_IDENTIFIER returns the _IDENTIFIER token.
	Get_IDENTIFIER() antlr.Token

	// Set_IDENTIFIER sets the _IDENTIFIER token.
	Set_IDENTIFIER(antlr.Token)

	// Get_enumTypeSpec returns the _enumTypeSpec rule contexts.
	Get_enumTypeSpec() IEnumTypeSpecContext

	// Get_structTypeSpec returns the _structTypeSpec rule contexts.
	Get_structTypeSpec() IStructTypeSpecContext

	// Get_unionTypeSpec returns the _unionTypeSpec rule contexts.
	Get_unionTypeSpec() IUnionTypeSpecContext

	// Set_enumTypeSpec sets the _enumTypeSpec rule contexts.
	Set_enumTypeSpec(IEnumTypeSpecContext)

	// Set_structTypeSpec sets the _structTypeSpec rule contexts.
	Set_structTypeSpec(IStructTypeSpecContext)

	// Set_unionTypeSpec sets the _unionTypeSpec rule contexts.
	Set_unionTypeSpec(IUnionTypeSpecContext)

	// GetV returns the v attribute.
	GetV() model.Type

	// SetV sets the v attribute.
	SetV(model.Type)

	// IsTypeSpecifierContext differentiates from other interfaces.
	IsTypeSpecifierContext()
}

type TypeSpecifierContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	v               model.Type
	_enumTypeSpec   IEnumTypeSpecContext
	_structTypeSpec IStructTypeSpecContext
	_unionTypeSpec  IUnionTypeSpecContext
	_IDENTIFIER     antlr.Token
}

func NewEmptyTypeSpecifierContext() *TypeSpecifierContext {
	p := new(TypeSpecifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XDRParserRULE_typeSpecifier
	return p
}

func (*TypeSpecifierContext) IsTypeSpecifierContext() {}

func NewTypeSpecifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSpecifierContext {
	p := new(TypeSpecifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XDRParserRULE_typeSpecifier

	return p
}

func (s *TypeSpecifierContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSpecifierContext) Get_IDENTIFIER() antlr.Token { return s._IDENTIFIER }

func (s *TypeSpecifierContext) Set_IDENTIFIER(v antlr.Token) { s._IDENTIFIER = v }

func (s *TypeSpecifierContext) Get_enumTypeSpec() IEnumTypeSpecContext { return s._enumTypeSpec }

func (s *TypeSpecifierContext) Get_structTypeSpec() IStructTypeSpecContext { return s._structTypeSpec }

func (s *TypeSpecifierContext) Get_unionTypeSpec() IUnionTypeSpecContext { return s._unionTypeSpec }

func (s *TypeSpecifierContext) Set_enumTypeSpec(v IEnumTypeSpecContext) { s._enumTypeSpec = v }

func (s *TypeSpecifierContext) Set_structTypeSpec(v IStructTypeSpecContext) { s._structTypeSpec = v }

func (s *TypeSpecifierContext) Set_unionTypeSpec(v IUnionTypeSpecContext) { s._unionTypeSpec = v }

func (s *TypeSpecifierContext) GetV() model.Type { return s.v }

func (s *TypeSpecifierContext) SetV(v model.Type) { s.v = v }

func (s *TypeSpecifierContext) EnumTypeSpec() IEnumTypeSpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumTypeSpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumTypeSpecContext)
}

func (s *TypeSpecifierContext) StructTypeSpec() IStructTypeSpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructTypeSpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructTypeSpecContext)
}

func (s *TypeSpecifierContext) UnionTypeSpec() IUnionTypeSpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnionTypeSpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnionTypeSpecContext)
}

func (s *TypeSpecifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(XDRParserIDENTIFIER, 0)
}

func (s *TypeSpecifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSpecifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.EnterTypeSpecifier(s)
	}
}

func (s *TypeSpecifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.ExitTypeSpecifier(s)
	}
}

func (p *XDRParser) TypeSpecifier() (localctx ITypeSpecifierContext) {
	this := p
	_ = this

	localctx = NewTypeSpecifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, XDRParserRULE_typeSpecifier)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		localctx.(*TypeSpecifierContext).v = model.IntType
		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == XDRParserT__12 {
			{
				p.SetState(146)
				p.Match(XDRParserT__12)
			}
			localctx.(*TypeSpecifierContext).v = model.UnsignedIntType

		}
		{
			p.SetState(150)
			p.Match(XDRParserT__13)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		localctx.(*TypeSpecifierContext).v = model.HyperType
		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == XDRParserT__12 {
			{
				p.SetState(152)
				p.Match(XDRParserT__12)
			}
			localctx.(*TypeSpecifierContext).v = model.UnsignedHyperType

		}
		{
			p.SetState(156)
			p.Match(XDRParserT__14)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(157)
			p.Match(XDRParserT__15)
		}
		localctx.(*TypeSpecifierContext).v = model.FloatType

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(159)
			p.Match(XDRParserT__16)
		}
		localctx.(*TypeSpecifierContext).v = model.DoubleType

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(161)
			p.Match(XDRParserT__17)
		}
		localctx.(*TypeSpecifierContext).v = model.QuadrupleType

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(163)
			p.Match(XDRParserT__18)
		}
		localctx.(*TypeSpecifierContext).v = model.BoolType

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(165)

			_x := p.EnumTypeSpec()

			localctx.(*TypeSpecifierContext)._enumTypeSpec = _x
		}
		localctx.(*TypeSpecifierContext).v = localctx.(*TypeSpecifierContext).Get_enumTypeSpec().GetV()

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(168)

			_x := p.StructTypeSpec()

			localctx.(*TypeSpecifierContext)._structTypeSpec = _x
		}
		localctx.(*TypeSpecifierContext).v = localctx.(*TypeSpecifierContext).Get_structTypeSpec().GetV()

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(171)

			_x := p.UnionTypeSpec()

			localctx.(*TypeSpecifierContext)._unionTypeSpec = _x
		}
		localctx.(*TypeSpecifierContext).v = localctx.(*TypeSpecifierContext).Get_unionTypeSpec().GetV()

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(174)

			_m := p.Match(XDRParserIDENTIFIER)

			localctx.(*TypeSpecifierContext)._IDENTIFIER = _m
		}
		localctx.(*TypeSpecifierContext).v = model.NewIdentifierType((func() string {
			if localctx.(*TypeSpecifierContext).Get_IDENTIFIER() == nil {
				return ""
			} else {
				return localctx.(*TypeSpecifierContext).Get_IDENTIFIER().GetText()
			}
		}()))

	}

	return localctx
}

// IEnumTypeSpecContext is an interface to support dynamic dispatch.
type IEnumTypeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_enumBody returns the _enumBody rule contexts.
	Get_enumBody() IEnumBodyContext

	// Set_enumBody sets the _enumBody rule contexts.
	Set_enumBody(IEnumBodyContext)

	// GetV returns the v attribute.
	GetV() model.Type

	// SetV sets the v attribute.
	SetV(model.Type)

	// IsEnumTypeSpecContext differentiates from other interfaces.
	IsEnumTypeSpecContext()
}

type EnumTypeSpecContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	v         model.Type
	_enumBody IEnumBodyContext
}

func NewEmptyEnumTypeSpecContext() *EnumTypeSpecContext {
	p := new(EnumTypeSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XDRParserRULE_enumTypeSpec
	return p
}

func (*EnumTypeSpecContext) IsEnumTypeSpecContext() {}

func NewEnumTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumTypeSpecContext {
	p := new(EnumTypeSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XDRParserRULE_enumTypeSpec

	return p
}

func (s *EnumTypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumTypeSpecContext) Get_enumBody() IEnumBodyContext { return s._enumBody }

func (s *EnumTypeSpecContext) Set_enumBody(v IEnumBodyContext) { s._enumBody = v }

func (s *EnumTypeSpecContext) GetV() model.Type { return s.v }

func (s *EnumTypeSpecContext) SetV(v model.Type) { s.v = v }

func (s *EnumTypeSpecContext) EnumBody() IEnumBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumBodyContext)
}

func (s *EnumTypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumTypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumTypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.EnterEnumTypeSpec(s)
	}
}

func (s *EnumTypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.ExitEnumTypeSpec(s)
	}
}

func (p *XDRParser) EnumTypeSpec() (localctx IEnumTypeSpecContext) {
	this := p
	_ = this

	localctx = NewEnumTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, XDRParserRULE_enumTypeSpec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(178)
		p.Match(XDRParserT__19)
	}
	{
		p.SetState(179)

		_x := p.EnumBody()

		localctx.(*EnumTypeSpecContext)._enumBody = _x
	}
	localctx.(*EnumTypeSpecContext).v = localctx.(*EnumTypeSpecContext).Get_enumBody().GetV()

	return localctx
}

// IEnumBodyContext is an interface to support dynamic dispatch.
type IEnumBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_IDENTIFIER returns the _IDENTIFIER token.
	Get_IDENTIFIER() antlr.Token

	// Set_IDENTIFIER sets the _IDENTIFIER token.
	Set_IDENTIFIER(antlr.Token)

	// Get_value returns the _value rule contexts.
	Get_value() IValueContext

	// Set_value sets the _value rule contexts.
	Set_value(IValueContext)

	// GetV returns the v attribute.
	GetV() model.Type

	// GetElements returns the elements attribute.
	GetElements() map[string]model.Value

	// SetV sets the v attribute.
	SetV(model.Type)

	// SetElements sets the elements attribute.
	SetElements(map[string]model.Value)

	// IsEnumBodyContext differentiates from other interfaces.
	IsEnumBodyContext()
}

type EnumBodyContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	v           model.Type
	elements    map[string]model.Value
	_IDENTIFIER antlr.Token
	_value      IValueContext
}

func NewEmptyEnumBodyContext() *EnumBodyContext {
	p := new(EnumBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XDRParserRULE_enumBody
	return p
}

func (*EnumBodyContext) IsEnumBodyContext() {}

func NewEnumBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumBodyContext {
	p := new(EnumBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XDRParserRULE_enumBody

	return p
}

func (s *EnumBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumBodyContext) Get_IDENTIFIER() antlr.Token { return s._IDENTIFIER }

func (s *EnumBodyContext) Set_IDENTIFIER(v antlr.Token) { s._IDENTIFIER = v }

func (s *EnumBodyContext) Get_value() IValueContext { return s._value }

func (s *EnumBodyContext) Set_value(v IValueContext) { s._value = v }

func (s *EnumBodyContext) GetV() model.Type { return s.v }

func (s *EnumBodyContext) GetElements() map[string]model.Value { return s.elements }

func (s *EnumBodyContext) SetV(v model.Type) { s.v = v }

func (s *EnumBodyContext) SetElements(v map[string]model.Value) { s.elements = v }

func (s *EnumBodyContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(XDRParserIDENTIFIER)
}

func (s *EnumBodyContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(XDRParserIDENTIFIER, i)
}

func (s *EnumBodyContext) AllValue() []IValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *EnumBodyContext) Value(i int) IValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *EnumBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.EnterEnumBody(s)
	}
}

func (s *EnumBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.ExitEnumBody(s)
	}
}

func (p *XDRParser) EnumBody() (localctx IEnumBodyContext) {
	this := p
	_ = this

	localctx = NewEnumBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, XDRParserRULE_enumBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		p.Match(XDRParserT__20)
	}
	{
		p.SetState(183)

		_m := p.Match(XDRParserIDENTIFIER)

		localctx.(*EnumBodyContext)._IDENTIFIER = _m
	}
	{
		p.SetState(184)
		p.Match(XDRParserT__21)
	}
	{
		p.SetState(185)

		_x := p.Value()

		localctx.(*EnumBodyContext)._value = _x
	}

	localctx.(*EnumBodyContext).elements = map[string]model.Value{
		(func() string {
			if localctx.(*EnumBodyContext).Get_IDENTIFIER() == nil {
				return ""
			} else {
				return localctx.(*EnumBodyContext).Get_IDENTIFIER().GetText()
			}
		}()): localctx.(*EnumBodyContext).Get_value().GetV(),
	}

	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == XDRParserT__22 {
		{
			p.SetState(187)
			p.Match(XDRParserT__22)
		}
		{
			p.SetState(188)

			_m := p.Match(XDRParserIDENTIFIER)

			localctx.(*EnumBodyContext)._IDENTIFIER = _m
		}
		{
			p.SetState(189)
			p.Match(XDRParserT__21)
		}
		{
			p.SetState(190)

			_x := p.Value()

			localctx.(*EnumBodyContext)._value = _x
		}

		// TODO: Throw an error when there are duplicates!
		localctx.(*EnumBodyContext).elements[(func() string {
			if localctx.(*EnumBodyContext).Get_IDENTIFIER() == nil {
				return ""
			} else {
				return localctx.(*EnumBodyContext).Get_IDENTIFIER().GetText()
			}
		}())] = localctx.(*EnumBodyContext).Get_value().GetV()

		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(198)
		p.Match(XDRParserT__23)
	}
	localctx.(*EnumBodyContext).v = model.NewEnumType(localctx.(*EnumBodyContext).elements)

	return localctx
}

// IStructTypeSpecContext is an interface to support dynamic dispatch.
type IStructTypeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_structBody returns the _structBody rule contexts.
	Get_structBody() IStructBodyContext

	// Set_structBody sets the _structBody rule contexts.
	Set_structBody(IStructBodyContext)

	// GetV returns the v attribute.
	GetV() model.Type

	// SetV sets the v attribute.
	SetV(model.Type)

	// IsStructTypeSpecContext differentiates from other interfaces.
	IsStructTypeSpecContext()
}

type StructTypeSpecContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	v           model.Type
	_structBody IStructBodyContext
}

func NewEmptyStructTypeSpecContext() *StructTypeSpecContext {
	p := new(StructTypeSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XDRParserRULE_structTypeSpec
	return p
}

func (*StructTypeSpecContext) IsStructTypeSpecContext() {}

func NewStructTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructTypeSpecContext {
	p := new(StructTypeSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XDRParserRULE_structTypeSpec

	return p
}

func (s *StructTypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *StructTypeSpecContext) Get_structBody() IStructBodyContext { return s._structBody }

func (s *StructTypeSpecContext) Set_structBody(v IStructBodyContext) { s._structBody = v }

func (s *StructTypeSpecContext) GetV() model.Type { return s.v }

func (s *StructTypeSpecContext) SetV(v model.Type) { s.v = v }

func (s *StructTypeSpecContext) StructBody() IStructBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructBodyContext)
}

func (s *StructTypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructTypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructTypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.EnterStructTypeSpec(s)
	}
}

func (s *StructTypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.ExitStructTypeSpec(s)
	}
}

func (p *XDRParser) StructTypeSpec() (localctx IStructTypeSpecContext) {
	this := p
	_ = this

	localctx = NewStructTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, XDRParserRULE_structTypeSpec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.Match(XDRParserT__24)
	}
	{
		p.SetState(202)

		_x := p.StructBody()

		localctx.(*StructTypeSpecContext)._structBody = _x
	}
	localctx.(*StructTypeSpecContext).v = localctx.(*StructTypeSpecContext).Get_structBody().GetV()

	return localctx
}

// IStructBodyContext is an interface to support dynamic dispatch.
type IStructBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_namedDeclaration returns the _namedDeclaration rule contexts.
	Get_namedDeclaration() INamedDeclarationContext

	// Set_namedDeclaration sets the _namedDeclaration rule contexts.
	Set_namedDeclaration(INamedDeclarationContext)

	// GetV returns the v attribute.
	GetV() model.Type

	// GetDeclarations returns the declarations attribute.
	GetDeclarations() []model.NamedDeclaration

	// SetV sets the v attribute.
	SetV(model.Type)

	// SetDeclarations sets the declarations attribute.
	SetDeclarations([]model.NamedDeclaration)

	// IsStructBodyContext differentiates from other interfaces.
	IsStructBodyContext()
}

type StructBodyContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	v                 model.Type
	declarations      []model.NamedDeclaration
	_namedDeclaration INamedDeclarationContext
}

func NewEmptyStructBodyContext() *StructBodyContext {
	p := new(StructBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XDRParserRULE_structBody
	return p
}

func (*StructBodyContext) IsStructBodyContext() {}

func NewStructBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructBodyContext {
	p := new(StructBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XDRParserRULE_structBody

	return p
}

func (s *StructBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *StructBodyContext) Get_namedDeclaration() INamedDeclarationContext {
	return s._namedDeclaration
}

func (s *StructBodyContext) Set_namedDeclaration(v INamedDeclarationContext) { s._namedDeclaration = v }

func (s *StructBodyContext) GetV() model.Type { return s.v }

func (s *StructBodyContext) GetDeclarations() []model.NamedDeclaration { return s.declarations }

func (s *StructBodyContext) SetV(v model.Type) { s.v = v }

func (s *StructBodyContext) SetDeclarations(v []model.NamedDeclaration) { s.declarations = v }

func (s *StructBodyContext) AllNamedDeclaration() []INamedDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INamedDeclarationContext); ok {
			len++
		}
	}

	tst := make([]INamedDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INamedDeclarationContext); ok {
			tst[i] = t.(INamedDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *StructBodyContext) NamedDeclaration(i int) INamedDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamedDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamedDeclarationContext)
}

func (s *StructBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.EnterStructBody(s)
	}
}

func (s *StructBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.ExitStructBody(s)
	}
}

func (p *XDRParser) StructBody() (localctx IStructBodyContext) {
	this := p
	_ = this

	localctx = NewStructBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, XDRParserRULE_structBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Match(XDRParserT__20)
	}
	{
		p.SetState(206)

		_x := p.NamedDeclaration()

		localctx.(*StructBodyContext)._namedDeclaration = _x
	}
	{
		p.SetState(207)
		p.Match(XDRParserT__1)
	}

	localctx.(*StructBodyContext).declarations = append(localctx.(*StructBodyContext).declarations, localctx.(*StructBodyContext).Get_namedDeclaration().GetV())

	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<XDRParserT__7)|(1<<XDRParserT__8)|(1<<XDRParserT__9)|(1<<XDRParserT__12)|(1<<XDRParserT__13)|(1<<XDRParserT__14)|(1<<XDRParserT__15)|(1<<XDRParserT__16)|(1<<XDRParserT__17)|(1<<XDRParserT__18)|(1<<XDRParserT__19)|(1<<XDRParserT__24)|(1<<XDRParserT__25))) != 0) || _la == XDRParserIDENTIFIER {
		{
			p.SetState(209)

			_x := p.NamedDeclaration()

			localctx.(*StructBodyContext)._namedDeclaration = _x
		}
		{
			p.SetState(210)
			p.Match(XDRParserT__1)
		}

		localctx.(*StructBodyContext).declarations = append(localctx.(*StructBodyContext).declarations, localctx.(*StructBodyContext).Get_namedDeclaration().GetV())

		p.SetState(217)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(218)
		p.Match(XDRParserT__23)
	}
	localctx.(*StructBodyContext).v = model.NewStructType(localctx.(*StructBodyContext).declarations)

	return localctx
}

// IUnionTypeSpecContext is an interface to support dynamic dispatch.
type IUnionTypeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_unionBody returns the _unionBody rule contexts.
	Get_unionBody() IUnionBodyContext

	// Set_unionBody sets the _unionBody rule contexts.
	Set_unionBody(IUnionBodyContext)

	// GetV returns the v attribute.
	GetV() model.Type

	// SetV sets the v attribute.
	SetV(model.Type)

	// IsUnionTypeSpecContext differentiates from other interfaces.
	IsUnionTypeSpecContext()
}

type UnionTypeSpecContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	v          model.Type
	_unionBody IUnionBodyContext
}

func NewEmptyUnionTypeSpecContext() *UnionTypeSpecContext {
	p := new(UnionTypeSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XDRParserRULE_unionTypeSpec
	return p
}

func (*UnionTypeSpecContext) IsUnionTypeSpecContext() {}

func NewUnionTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionTypeSpecContext {
	p := new(UnionTypeSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XDRParserRULE_unionTypeSpec

	return p
}

func (s *UnionTypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionTypeSpecContext) Get_unionBody() IUnionBodyContext { return s._unionBody }

func (s *UnionTypeSpecContext) Set_unionBody(v IUnionBodyContext) { s._unionBody = v }

func (s *UnionTypeSpecContext) GetV() model.Type { return s.v }

func (s *UnionTypeSpecContext) SetV(v model.Type) { s.v = v }

func (s *UnionTypeSpecContext) UnionBody() IUnionBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnionBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnionBodyContext)
}

func (s *UnionTypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionTypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionTypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.EnterUnionTypeSpec(s)
	}
}

func (s *UnionTypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.ExitUnionTypeSpec(s)
	}
}

func (p *XDRParser) UnionTypeSpec() (localctx IUnionTypeSpecContext) {
	this := p
	_ = this

	localctx = NewUnionTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, XDRParserRULE_unionTypeSpec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(221)
		p.Match(XDRParserT__25)
	}
	{
		p.SetState(222)

		_x := p.UnionBody()

		localctx.(*UnionTypeSpecContext)._unionBody = _x
	}
	localctx.(*UnionTypeSpecContext).v = localctx.(*UnionTypeSpecContext).Get_unionBody().GetV()

	return localctx
}

// IUnionBodyContext is an interface to support dynamic dispatch.
type IUnionBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_namedDeclaration returns the _namedDeclaration rule contexts.
	Get_namedDeclaration() INamedDeclarationContext

	// Get_caseSpec returns the _caseSpec rule contexts.
	Get_caseSpec() ICaseSpecContext

	// Get_declaration returns the _declaration rule contexts.
	Get_declaration() IDeclarationContext

	// Set_namedDeclaration sets the _namedDeclaration rule contexts.
	Set_namedDeclaration(INamedDeclarationContext)

	// Set_caseSpec sets the _caseSpec rule contexts.
	Set_caseSpec(ICaseSpecContext)

	// Set_declaration sets the _declaration rule contexts.
	Set_declaration(IDeclarationContext)

	// GetV returns the v attribute.
	GetV() model.Type

	// GetCases returns the cases attribute.
	GetCases() []model.CaseSpec

	// GetDefaultCase returns the defaultCase attribute.
	GetDefaultCase() model.Declaration

	// SetV sets the v attribute.
	SetV(model.Type)

	// SetCases sets the cases attribute.
	SetCases([]model.CaseSpec)

	// SetDefaultCase sets the defaultCase attribute.
	SetDefaultCase(model.Declaration)

	// IsUnionBodyContext differentiates from other interfaces.
	IsUnionBodyContext()
}

type UnionBodyContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	v                 model.Type
	cases             []model.CaseSpec
	defaultCase       model.Declaration
	_namedDeclaration INamedDeclarationContext
	_caseSpec         ICaseSpecContext
	_declaration      IDeclarationContext
}

func NewEmptyUnionBodyContext() *UnionBodyContext {
	p := new(UnionBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XDRParserRULE_unionBody
	return p
}

func (*UnionBodyContext) IsUnionBodyContext() {}

func NewUnionBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionBodyContext {
	p := new(UnionBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XDRParserRULE_unionBody

	return p
}

func (s *UnionBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionBodyContext) Get_namedDeclaration() INamedDeclarationContext {
	return s._namedDeclaration
}

func (s *UnionBodyContext) Get_caseSpec() ICaseSpecContext { return s._caseSpec }

func (s *UnionBodyContext) Get_declaration() IDeclarationContext { return s._declaration }

func (s *UnionBodyContext) Set_namedDeclaration(v INamedDeclarationContext) { s._namedDeclaration = v }

func (s *UnionBodyContext) Set_caseSpec(v ICaseSpecContext) { s._caseSpec = v }

func (s *UnionBodyContext) Set_declaration(v IDeclarationContext) { s._declaration = v }

func (s *UnionBodyContext) GetV() model.Type { return s.v }

func (s *UnionBodyContext) GetCases() []model.CaseSpec { return s.cases }

func (s *UnionBodyContext) GetDefaultCase() model.Declaration { return s.defaultCase }

func (s *UnionBodyContext) SetV(v model.Type) { s.v = v }

func (s *UnionBodyContext) SetCases(v []model.CaseSpec) { s.cases = v }

func (s *UnionBodyContext) SetDefaultCase(v model.Declaration) { s.defaultCase = v }

func (s *UnionBodyContext) NamedDeclaration() INamedDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamedDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamedDeclarationContext)
}

func (s *UnionBodyContext) AllCaseSpec() []ICaseSpecContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICaseSpecContext); ok {
			len++
		}
	}

	tst := make([]ICaseSpecContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICaseSpecContext); ok {
			tst[i] = t.(ICaseSpecContext)
			i++
		}
	}

	return tst
}

func (s *UnionBodyContext) CaseSpec(i int) ICaseSpecContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseSpecContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICaseSpecContext)
}

func (s *UnionBodyContext) Declaration() IDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *UnionBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.EnterUnionBody(s)
	}
}

func (s *UnionBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.ExitUnionBody(s)
	}
}

func (p *XDRParser) UnionBody() (localctx IUnionBodyContext) {
	this := p
	_ = this

	localctx = NewUnionBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, XDRParserRULE_unionBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(225)
		p.Match(XDRParserT__26)
	}
	{
		p.SetState(226)
		p.Match(XDRParserT__27)
	}
	{
		p.SetState(227)

		_x := p.NamedDeclaration()

		localctx.(*UnionBodyContext)._namedDeclaration = _x
	}
	{
		p.SetState(228)
		p.Match(XDRParserT__28)
	}
	{
		p.SetState(229)
		p.Match(XDRParserT__20)
	}
	{
		p.SetState(230)

		_x := p.CaseSpec()

		localctx.(*UnionBodyContext)._caseSpec = _x
	}
	localctx.(*UnionBodyContext).cases = append(localctx.(*UnionBodyContext).cases, localctx.(*UnionBodyContext).Get_caseSpec().GetV())
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == XDRParserT__31 {
		{
			p.SetState(232)

			_x := p.CaseSpec()

			localctx.(*UnionBodyContext)._caseSpec = _x
		}
		localctx.(*UnionBodyContext).cases = append(localctx.(*UnionBodyContext).cases, localctx.(*UnionBodyContext).Get_caseSpec().GetV())

		p.SetState(239)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == XDRParserT__29 {
		{
			p.SetState(240)
			p.Match(XDRParserT__29)
		}
		{
			p.SetState(241)
			p.Match(XDRParserT__30)
		}
		{
			p.SetState(242)

			_x := p.Declaration()

			localctx.(*UnionBodyContext)._declaration = _x
		}
		{
			p.SetState(243)
			p.Match(XDRParserT__1)
		}

		localctx.(*UnionBodyContext).defaultCase = localctx.(*UnionBodyContext).Get_declaration().GetV()

	}
	{
		p.SetState(248)
		p.Match(XDRParserT__23)
	}
	localctx.(*UnionBodyContext).v = model.NewUnionType(localctx.(*UnionBodyContext).Get_namedDeclaration().GetV(), localctx.(*UnionBodyContext).cases, localctx.(*UnionBodyContext).defaultCase)

	return localctx
}

// ICaseSpecContext is an interface to support dynamic dispatch.
type ICaseSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_value returns the _value rule contexts.
	Get_value() IValueContext

	// Get_declaration returns the _declaration rule contexts.
	Get_declaration() IDeclarationContext

	// Set_value sets the _value rule contexts.
	Set_value(IValueContext)

	// Set_declaration sets the _declaration rule contexts.
	Set_declaration(IDeclarationContext)

	// GetV returns the v attribute.
	GetV() model.CaseSpec

	// GetValues returns the values attribute.
	GetValues() []model.Value

	// SetV sets the v attribute.
	SetV(model.CaseSpec)

	// SetValues sets the values attribute.
	SetValues([]model.Value)

	// IsCaseSpecContext differentiates from other interfaces.
	IsCaseSpecContext()
}

type CaseSpecContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	v            model.CaseSpec
	values       []model.Value
	_value       IValueContext
	_declaration IDeclarationContext
}

func NewEmptyCaseSpecContext() *CaseSpecContext {
	p := new(CaseSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XDRParserRULE_caseSpec
	return p
}

func (*CaseSpecContext) IsCaseSpecContext() {}

func NewCaseSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseSpecContext {
	p := new(CaseSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XDRParserRULE_caseSpec

	return p
}

func (s *CaseSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseSpecContext) Get_value() IValueContext { return s._value }

func (s *CaseSpecContext) Get_declaration() IDeclarationContext { return s._declaration }

func (s *CaseSpecContext) Set_value(v IValueContext) { s._value = v }

func (s *CaseSpecContext) Set_declaration(v IDeclarationContext) { s._declaration = v }

func (s *CaseSpecContext) GetV() model.CaseSpec { return s.v }

func (s *CaseSpecContext) GetValues() []model.Value { return s.values }

func (s *CaseSpecContext) SetV(v model.CaseSpec) { s.v = v }

func (s *CaseSpecContext) SetValues(v []model.Value) { s.values = v }

func (s *CaseSpecContext) AllValue() []IValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *CaseSpecContext) Value(i int) IValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *CaseSpecContext) Declaration() IDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *CaseSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.EnterCaseSpec(s)
	}
}

func (s *CaseSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.ExitCaseSpec(s)
	}
}

func (p *XDRParser) CaseSpec() (localctx ICaseSpecContext) {
	this := p
	_ = this

	localctx = NewCaseSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, XDRParserRULE_caseSpec)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(251)
		p.Match(XDRParserT__31)
	}
	{
		p.SetState(252)

		_x := p.Value()

		localctx.(*CaseSpecContext)._value = _x
	}
	{
		p.SetState(253)
		p.Match(XDRParserT__30)
	}
	localctx.(*CaseSpecContext).values = append(localctx.(*CaseSpecContext).values, localctx.(*CaseSpecContext).Get_value().GetV())
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == XDRParserT__31 {
		{
			p.SetState(255)
			p.Match(XDRParserT__31)
		}
		{
			p.SetState(256)

			_x := p.Value()

			localctx.(*CaseSpecContext)._value = _x
		}
		{
			p.SetState(257)
			p.Match(XDRParserT__30)
		}
		localctx.(*CaseSpecContext).values = append(localctx.(*CaseSpecContext).values, localctx.(*CaseSpecContext).Get_value().GetV())

		p.SetState(264)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(265)

		_x := p.Declaration()

		localctx.(*CaseSpecContext)._declaration = _x
	}
	{
		p.SetState(266)
		p.Match(XDRParserT__1)
	}
	localctx.(*CaseSpecContext).v = model.NewCaseSpec(localctx.(*CaseSpecContext).values, localctx.(*CaseSpecContext).Get_declaration().GetV())

	return localctx
}

// IConstantDefContext is an interface to support dynamic dispatch.
type IConstantDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_IDENTIFIER returns the _IDENTIFIER token.
	Get_IDENTIFIER() antlr.Token

	// Set_IDENTIFIER sets the _IDENTIFIER token.
	Set_IDENTIFIER(antlr.Token)

	// Get_constant returns the _constant rule contexts.
	Get_constant() IConstantContext

	// Set_constant sets the _constant rule contexts.
	Set_constant(IConstantContext)

	// GetV returns the v attribute.
	GetV() model.Definition

	// SetV sets the v attribute.
	SetV(model.Definition)

	// IsConstantDefContext differentiates from other interfaces.
	IsConstantDefContext()
}

type ConstantDefContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	v           model.Definition
	_IDENTIFIER antlr.Token
	_constant   IConstantContext
}

func NewEmptyConstantDefContext() *ConstantDefContext {
	p := new(ConstantDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XDRParserRULE_constantDef
	return p
}

func (*ConstantDefContext) IsConstantDefContext() {}

func NewConstantDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantDefContext {
	p := new(ConstantDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XDRParserRULE_constantDef

	return p
}

func (s *ConstantDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantDefContext) Get_IDENTIFIER() antlr.Token { return s._IDENTIFIER }

func (s *ConstantDefContext) Set_IDENTIFIER(v antlr.Token) { s._IDENTIFIER = v }

func (s *ConstantDefContext) Get_constant() IConstantContext { return s._constant }

func (s *ConstantDefContext) Set_constant(v IConstantContext) { s._constant = v }

func (s *ConstantDefContext) GetV() model.Definition { return s.v }

func (s *ConstantDefContext) SetV(v model.Definition) { s.v = v }

func (s *ConstantDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(XDRParserIDENTIFIER, 0)
}

func (s *ConstantDefContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ConstantDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.EnterConstantDef(s)
	}
}

func (s *ConstantDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.ExitConstantDef(s)
	}
}

func (p *XDRParser) ConstantDef() (localctx IConstantDefContext) {
	this := p
	_ = this

	localctx = NewConstantDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, XDRParserRULE_constantDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(269)
		p.Match(XDRParserT__32)
	}
	{
		p.SetState(270)

		_m := p.Match(XDRParserIDENTIFIER)

		localctx.(*ConstantDefContext)._IDENTIFIER = _m
	}
	{
		p.SetState(271)
		p.Match(XDRParserT__21)
	}
	{
		p.SetState(272)

		_x := p.Constant()

		localctx.(*ConstantDefContext)._constant = _x
	}
	{
		p.SetState(273)
		p.Match(XDRParserT__1)
	}

	localctx.(*ConstantDefContext).v = model.NewConstantDefinition((func() string {
		if localctx.(*ConstantDefContext).Get_IDENTIFIER() == nil {
			return ""
		} else {
			return localctx.(*ConstantDefContext).Get_IDENTIFIER().GetText()
		}
	}()), localctx.(*ConstantDefContext).Get_constant().GetV())

	return localctx
}

// ITypeDefContext is an interface to support dynamic dispatch.
type ITypeDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_IDENTIFIER returns the _IDENTIFIER token.
	Get_IDENTIFIER() antlr.Token

	// Set_IDENTIFIER sets the _IDENTIFIER token.
	Set_IDENTIFIER(antlr.Token)

	// Get_namedDeclaration returns the _namedDeclaration rule contexts.
	Get_namedDeclaration() INamedDeclarationContext

	// Get_enumBody returns the _enumBody rule contexts.
	Get_enumBody() IEnumBodyContext

	// Get_structBody returns the _structBody rule contexts.
	Get_structBody() IStructBodyContext

	// Get_unionBody returns the _unionBody rule contexts.
	Get_unionBody() IUnionBodyContext

	// Set_namedDeclaration sets the _namedDeclaration rule contexts.
	Set_namedDeclaration(INamedDeclarationContext)

	// Set_enumBody sets the _enumBody rule contexts.
	Set_enumBody(IEnumBodyContext)

	// Set_structBody sets the _structBody rule contexts.
	Set_structBody(IStructBodyContext)

	// Set_unionBody sets the _unionBody rule contexts.
	Set_unionBody(IUnionBodyContext)

	// GetV returns the v attribute.
	GetV() model.NamedDeclaration

	// SetV sets the v attribute.
	SetV(model.NamedDeclaration)

	// IsTypeDefContext differentiates from other interfaces.
	IsTypeDefContext()
}

type TypeDefContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	v                 model.NamedDeclaration
	_namedDeclaration INamedDeclarationContext
	_IDENTIFIER       antlr.Token
	_enumBody         IEnumBodyContext
	_structBody       IStructBodyContext
	_unionBody        IUnionBodyContext
}

func NewEmptyTypeDefContext() *TypeDefContext {
	p := new(TypeDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XDRParserRULE_typeDef
	return p
}

func (*TypeDefContext) IsTypeDefContext() {}

func NewTypeDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDefContext {
	p := new(TypeDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XDRParserRULE_typeDef

	return p
}

func (s *TypeDefContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDefContext) Get_IDENTIFIER() antlr.Token { return s._IDENTIFIER }

func (s *TypeDefContext) Set_IDENTIFIER(v antlr.Token) { s._IDENTIFIER = v }

func (s *TypeDefContext) Get_namedDeclaration() INamedDeclarationContext { return s._namedDeclaration }

func (s *TypeDefContext) Get_enumBody() IEnumBodyContext { return s._enumBody }

func (s *TypeDefContext) Get_structBody() IStructBodyContext { return s._structBody }

func (s *TypeDefContext) Get_unionBody() IUnionBodyContext { return s._unionBody }

func (s *TypeDefContext) Set_namedDeclaration(v INamedDeclarationContext) { s._namedDeclaration = v }

func (s *TypeDefContext) Set_enumBody(v IEnumBodyContext) { s._enumBody = v }

func (s *TypeDefContext) Set_structBody(v IStructBodyContext) { s._structBody = v }

func (s *TypeDefContext) Set_unionBody(v IUnionBodyContext) { s._unionBody = v }

func (s *TypeDefContext) GetV() model.NamedDeclaration { return s.v }

func (s *TypeDefContext) SetV(v model.NamedDeclaration) { s.v = v }

func (s *TypeDefContext) NamedDeclaration() INamedDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamedDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamedDeclarationContext)
}

func (s *TypeDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(XDRParserIDENTIFIER, 0)
}

func (s *TypeDefContext) EnumBody() IEnumBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumBodyContext)
}

func (s *TypeDefContext) StructBody() IStructBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructBodyContext)
}

func (s *TypeDefContext) UnionBody() IUnionBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnionBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnionBodyContext)
}

func (s *TypeDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.EnterTypeDef(s)
	}
}

func (s *TypeDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.ExitTypeDef(s)
	}
}

func (p *XDRParser) TypeDef() (localctx ITypeDefContext) {
	this := p
	_ = this

	localctx = NewTypeDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, XDRParserRULE_typeDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(299)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case XDRParserT__33:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(276)
			p.Match(XDRParserT__33)
		}
		{
			p.SetState(277)

			_x := p.NamedDeclaration()

			localctx.(*TypeDefContext)._namedDeclaration = _x
		}
		{
			p.SetState(278)
			p.Match(XDRParserT__1)
		}

		localctx.(*TypeDefContext).v = localctx.(*TypeDefContext).Get_namedDeclaration().GetV()

	case XDRParserT__19:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(281)
			p.Match(XDRParserT__19)
		}
		{
			p.SetState(282)

			_m := p.Match(XDRParserIDENTIFIER)

			localctx.(*TypeDefContext)._IDENTIFIER = _m
		}
		{
			p.SetState(283)

			_x := p.EnumBody()

			localctx.(*TypeDefContext)._enumBody = _x
		}
		{
			p.SetState(284)
			p.Match(XDRParserT__1)
		}

		localctx.(*TypeDefContext).v = model.NewNamedDeclaration((func() string {
			if localctx.(*TypeDefContext).Get_IDENTIFIER() == nil {
				return ""
			} else {
				return localctx.(*TypeDefContext).Get_IDENTIFIER().GetText()
			}
		}()), localctx.(*TypeDefContext).Get_enumBody().GetV())

	case XDRParserT__24:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(287)
			p.Match(XDRParserT__24)
		}
		{
			p.SetState(288)

			_m := p.Match(XDRParserIDENTIFIER)

			localctx.(*TypeDefContext)._IDENTIFIER = _m
		}
		{
			p.SetState(289)

			_x := p.StructBody()

			localctx.(*TypeDefContext)._structBody = _x
		}
		{
			p.SetState(290)
			p.Match(XDRParserT__1)
		}

		localctx.(*TypeDefContext).v = model.NewNamedDeclaration((func() string {
			if localctx.(*TypeDefContext).Get_IDENTIFIER() == nil {
				return ""
			} else {
				return localctx.(*TypeDefContext).Get_IDENTIFIER().GetText()
			}
		}()), localctx.(*TypeDefContext).Get_structBody().GetV())

	case XDRParserT__25:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(293)
			p.Match(XDRParserT__25)
		}
		{
			p.SetState(294)

			_m := p.Match(XDRParserIDENTIFIER)

			localctx.(*TypeDefContext)._IDENTIFIER = _m
		}
		{
			p.SetState(295)

			_x := p.UnionBody()

			localctx.(*TypeDefContext)._unionBody = _x
		}
		{
			p.SetState(296)
			p.Match(XDRParserT__1)
		}

		localctx.(*TypeDefContext).v = model.NewNamedDeclaration((func() string {
			if localctx.(*TypeDefContext).Get_IDENTIFIER() == nil {
				return ""
			} else {
				return localctx.(*TypeDefContext).Get_IDENTIFIER().GetText()
			}
		}()), localctx.(*TypeDefContext).Get_unionBody().GetV())

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDefinitionContext is an interface to support dynamic dispatch.
type IDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_typeDef returns the _typeDef rule contexts.
	Get_typeDef() ITypeDefContext

	// Get_constantDef returns the _constantDef rule contexts.
	Get_constantDef() IConstantDefContext

	// Get_programDef returns the _programDef rule contexts.
	Get_programDef() IProgramDefContext

	// Set_typeDef sets the _typeDef rule contexts.
	Set_typeDef(ITypeDefContext)

	// Set_constantDef sets the _constantDef rule contexts.
	Set_constantDef(IConstantDefContext)

	// Set_programDef sets the _programDef rule contexts.
	Set_programDef(IProgramDefContext)

	// GetV returns the v attribute.
	GetV() model.Definition

	// SetV sets the v attribute.
	SetV(model.Definition)

	// IsDefinitionContext differentiates from other interfaces.
	IsDefinitionContext()
}

type DefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	v            model.Definition
	_typeDef     ITypeDefContext
	_constantDef IConstantDefContext
	_programDef  IProgramDefContext
}

func NewEmptyDefinitionContext() *DefinitionContext {
	p := new(DefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XDRParserRULE_definition
	return p
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	p := new(DefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XDRParserRULE_definition

	return p
}

func (s *DefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionContext) Get_typeDef() ITypeDefContext { return s._typeDef }

func (s *DefinitionContext) Get_constantDef() IConstantDefContext { return s._constantDef }

func (s *DefinitionContext) Get_programDef() IProgramDefContext { return s._programDef }

func (s *DefinitionContext) Set_typeDef(v ITypeDefContext) { s._typeDef = v }

func (s *DefinitionContext) Set_constantDef(v IConstantDefContext) { s._constantDef = v }

func (s *DefinitionContext) Set_programDef(v IProgramDefContext) { s._programDef = v }

func (s *DefinitionContext) GetV() model.Definition { return s.v }

func (s *DefinitionContext) SetV(v model.Definition) { s.v = v }

func (s *DefinitionContext) TypeDef() ITypeDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDefContext)
}

func (s *DefinitionContext) ConstantDef() IConstantDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantDefContext)
}

func (s *DefinitionContext) ProgramDef() IProgramDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProgramDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProgramDefContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.EnterDefinition(s)
	}
}

func (s *DefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.ExitDefinition(s)
	}
}

func (p *XDRParser) Definition() (localctx IDefinitionContext) {
	this := p
	_ = this

	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, XDRParserRULE_definition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(310)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case XDRParserT__19, XDRParserT__24, XDRParserT__25, XDRParserT__33:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(301)

			_x := p.TypeDef()

			localctx.(*DefinitionContext)._typeDef = _x
		}

		localctx.(*DefinitionContext).v = model.NewTypeDefinition(localctx.(*DefinitionContext).Get_typeDef().GetV())

	case XDRParserT__32:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(304)

			_x := p.ConstantDef()

			localctx.(*DefinitionContext)._constantDef = _x
		}
		localctx.(*DefinitionContext).v = localctx.(*DefinitionContext).Get_constantDef().GetV()

	case XDRParserT__34:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(307)

			_x := p.ProgramDef()

			localctx.(*DefinitionContext)._programDef = _x
		}
		localctx.(*DefinitionContext).v = localctx.(*DefinitionContext).Get_programDef().GetV()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISpecificationContext is an interface to support dynamic dispatch.
type ISpecificationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_definition returns the _definition rule contexts.
	Get_definition() IDefinitionContext

	// Set_definition sets the _definition rule contexts.
	Set_definition(IDefinitionContext)

	// GetV returns the v attribute.
	GetV() []model.Definition

	// SetV sets the v attribute.
	SetV([]model.Definition)

	// IsSpecificationContext differentiates from other interfaces.
	IsSpecificationContext()
}

type SpecificationContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	v           []model.Definition
	_definition IDefinitionContext
}

func NewEmptySpecificationContext() *SpecificationContext {
	p := new(SpecificationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XDRParserRULE_specification
	return p
}

func (*SpecificationContext) IsSpecificationContext() {}

func NewSpecificationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecificationContext {
	p := new(SpecificationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XDRParserRULE_specification

	return p
}

func (s *SpecificationContext) GetParser() antlr.Parser { return s.parser }

func (s *SpecificationContext) Get_definition() IDefinitionContext { return s._definition }

func (s *SpecificationContext) Set_definition(v IDefinitionContext) { s._definition = v }

func (s *SpecificationContext) GetV() []model.Definition { return s.v }

func (s *SpecificationContext) SetV(v []model.Definition) { s.v = v }

func (s *SpecificationContext) AllDefinition() []IDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDefinitionContext); ok {
			tst[i] = t.(IDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *SpecificationContext) Definition(i int) IDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefinitionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *SpecificationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecificationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpecificationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.EnterSpecification(s)
	}
}

func (s *SpecificationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.ExitSpecification(s)
	}
}

func (p *XDRParser) Specification() (localctx ISpecificationContext) {
	this := p
	_ = this

	localctx = NewSpecificationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, XDRParserRULE_specification)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-20)&-(0x1f+1)) == 0 && ((1<<uint((_la-20)))&((1<<(XDRParserT__19-20))|(1<<(XDRParserT__24-20))|(1<<(XDRParserT__25-20))|(1<<(XDRParserT__32-20))|(1<<(XDRParserT__33-20))|(1<<(XDRParserT__34-20)))) != 0 {
		{
			p.SetState(312)

			_x := p.Definition()

			localctx.(*SpecificationContext)._definition = _x
		}

		localctx.(*SpecificationContext).v = append(localctx.(*SpecificationContext).v, localctx.(*SpecificationContext).Get_definition().GetV())

		p.SetState(319)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IProgramDefContext is an interface to support dynamic dispatch.
type IProgramDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_IDENTIFIER returns the _IDENTIFIER token.
	Get_IDENTIFIER() antlr.Token

	// Set_IDENTIFIER sets the _IDENTIFIER token.
	Set_IDENTIFIER(antlr.Token)

	// Get_versionDef returns the _versionDef rule contexts.
	Get_versionDef() IVersionDefContext

	// Get_constant returns the _constant rule contexts.
	Get_constant() IConstantContext

	// Set_versionDef sets the _versionDef rule contexts.
	Set_versionDef(IVersionDefContext)

	// Set_constant sets the _constant rule contexts.
	Set_constant(IConstantContext)

	// GetV returns the v attribute.
	GetV() model.Definition

	// GetVersions returns the versions attribute.
	GetVersions() []model.Version

	// SetV sets the v attribute.
	SetV(model.Definition)

	// SetVersions sets the versions attribute.
	SetVersions([]model.Version)

	// IsProgramDefContext differentiates from other interfaces.
	IsProgramDefContext()
}

type ProgramDefContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	v           model.Definition
	versions    []model.Version
	_IDENTIFIER antlr.Token
	_versionDef IVersionDefContext
	_constant   IConstantContext
}

func NewEmptyProgramDefContext() *ProgramDefContext {
	p := new(ProgramDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XDRParserRULE_programDef
	return p
}

func (*ProgramDefContext) IsProgramDefContext() {}

func NewProgramDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramDefContext {
	p := new(ProgramDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XDRParserRULE_programDef

	return p
}

func (s *ProgramDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramDefContext) Get_IDENTIFIER() antlr.Token { return s._IDENTIFIER }

func (s *ProgramDefContext) Set_IDENTIFIER(v antlr.Token) { s._IDENTIFIER = v }

func (s *ProgramDefContext) Get_versionDef() IVersionDefContext { return s._versionDef }

func (s *ProgramDefContext) Get_constant() IConstantContext { return s._constant }

func (s *ProgramDefContext) Set_versionDef(v IVersionDefContext) { s._versionDef = v }

func (s *ProgramDefContext) Set_constant(v IConstantContext) { s._constant = v }

func (s *ProgramDefContext) GetV() model.Definition { return s.v }

func (s *ProgramDefContext) GetVersions() []model.Version { return s.versions }

func (s *ProgramDefContext) SetV(v model.Definition) { s.v = v }

func (s *ProgramDefContext) SetVersions(v []model.Version) { s.versions = v }

func (s *ProgramDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(XDRParserIDENTIFIER, 0)
}

func (s *ProgramDefContext) AllVersionDef() []IVersionDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVersionDefContext); ok {
			len++
		}
	}

	tst := make([]IVersionDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVersionDefContext); ok {
			tst[i] = t.(IVersionDefContext)
			i++
		}
	}

	return tst
}

func (s *ProgramDefContext) VersionDef(i int) IVersionDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVersionDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVersionDefContext)
}

func (s *ProgramDefContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ProgramDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.EnterProgramDef(s)
	}
}

func (s *ProgramDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.ExitProgramDef(s)
	}
}

func (p *XDRParser) ProgramDef() (localctx IProgramDefContext) {
	this := p
	_ = this

	localctx = NewProgramDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, XDRParserRULE_programDef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(320)
		p.Match(XDRParserT__34)
	}
	{
		p.SetState(321)

		_m := p.Match(XDRParserIDENTIFIER)

		localctx.(*ProgramDefContext)._IDENTIFIER = _m
	}
	{
		p.SetState(322)
		p.Match(XDRParserT__20)
	}
	{
		p.SetState(323)

		_x := p.VersionDef()

		localctx.(*ProgramDefContext)._versionDef = _x
	}
	localctx.(*ProgramDefContext).versions = append(localctx.(*ProgramDefContext).versions, localctx.(*ProgramDefContext).Get_versionDef().GetV())
	p.SetState(330)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == XDRParserT__35 {
		{
			p.SetState(325)

			_x := p.VersionDef()

			localctx.(*ProgramDefContext)._versionDef = _x
		}
		localctx.(*ProgramDefContext).versions = append(localctx.(*ProgramDefContext).versions, localctx.(*ProgramDefContext).Get_versionDef().GetV())

		p.SetState(332)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(333)
		p.Match(XDRParserT__23)
	}
	{
		p.SetState(334)
		p.Match(XDRParserT__21)
	}
	{
		p.SetState(335)

		_x := p.Constant()

		localctx.(*ProgramDefContext)._constant = _x
	}
	{
		p.SetState(336)
		p.Match(XDRParserT__1)
	}

	localctx.(*ProgramDefContext).v = model.NewProgramDefinition((func() string {
		if localctx.(*ProgramDefContext).Get_IDENTIFIER() == nil {
			return ""
		} else {
			return localctx.(*ProgramDefContext).Get_IDENTIFIER().GetText()
		}
	}()), localctx.(*ProgramDefContext).versions, localctx.(*ProgramDefContext).Get_constant().GetV())

	return localctx
}

// IVersionDefContext is an interface to support dynamic dispatch.
type IVersionDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_IDENTIFIER returns the _IDENTIFIER token.
	Get_IDENTIFIER() antlr.Token

	// Set_IDENTIFIER sets the _IDENTIFIER token.
	Set_IDENTIFIER(antlr.Token)

	// Get_procedureDef returns the _procedureDef rule contexts.
	Get_procedureDef() IProcedureDefContext

	// Get_constant returns the _constant rule contexts.
	Get_constant() IConstantContext

	// Set_procedureDef sets the _procedureDef rule contexts.
	Set_procedureDef(IProcedureDefContext)

	// Set_constant sets the _constant rule contexts.
	Set_constant(IConstantContext)

	// GetV returns the v attribute.
	GetV() model.Version

	// GetProcedures returns the procedures attribute.
	GetProcedures() []model.Procedure

	// SetV sets the v attribute.
	SetV(model.Version)

	// SetProcedures sets the procedures attribute.
	SetProcedures([]model.Procedure)

	// IsVersionDefContext differentiates from other interfaces.
	IsVersionDefContext()
}

type VersionDefContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	v             model.Version
	procedures    []model.Procedure
	_IDENTIFIER   antlr.Token
	_procedureDef IProcedureDefContext
	_constant     IConstantContext
}

func NewEmptyVersionDefContext() *VersionDefContext {
	p := new(VersionDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XDRParserRULE_versionDef
	return p
}

func (*VersionDefContext) IsVersionDefContext() {}

func NewVersionDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VersionDefContext {
	p := new(VersionDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XDRParserRULE_versionDef

	return p
}

func (s *VersionDefContext) GetParser() antlr.Parser { return s.parser }

func (s *VersionDefContext) Get_IDENTIFIER() antlr.Token { return s._IDENTIFIER }

func (s *VersionDefContext) Set_IDENTIFIER(v antlr.Token) { s._IDENTIFIER = v }

func (s *VersionDefContext) Get_procedureDef() IProcedureDefContext { return s._procedureDef }

func (s *VersionDefContext) Get_constant() IConstantContext { return s._constant }

func (s *VersionDefContext) Set_procedureDef(v IProcedureDefContext) { s._procedureDef = v }

func (s *VersionDefContext) Set_constant(v IConstantContext) { s._constant = v }

func (s *VersionDefContext) GetV() model.Version { return s.v }

func (s *VersionDefContext) GetProcedures() []model.Procedure { return s.procedures }

func (s *VersionDefContext) SetV(v model.Version) { s.v = v }

func (s *VersionDefContext) SetProcedures(v []model.Procedure) { s.procedures = v }

func (s *VersionDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(XDRParserIDENTIFIER, 0)
}

func (s *VersionDefContext) AllProcedureDef() []IProcedureDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IProcedureDefContext); ok {
			len++
		}
	}

	tst := make([]IProcedureDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IProcedureDefContext); ok {
			tst[i] = t.(IProcedureDefContext)
			i++
		}
	}

	return tst
}

func (s *VersionDefContext) ProcedureDef(i int) IProcedureDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProcedureDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProcedureDefContext)
}

func (s *VersionDefContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *VersionDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VersionDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VersionDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.EnterVersionDef(s)
	}
}

func (s *VersionDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.ExitVersionDef(s)
	}
}

func (p *XDRParser) VersionDef() (localctx IVersionDefContext) {
	this := p
	_ = this

	localctx = NewVersionDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, XDRParserRULE_versionDef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(339)
		p.Match(XDRParserT__35)
	}
	{
		p.SetState(340)

		_m := p.Match(XDRParserIDENTIFIER)

		localctx.(*VersionDefContext)._IDENTIFIER = _m
	}
	{
		p.SetState(341)
		p.Match(XDRParserT__20)
	}
	{
		p.SetState(342)

		_x := p.ProcedureDef()

		localctx.(*VersionDefContext)._procedureDef = _x
	}
	localctx.(*VersionDefContext).procedures = append(localctx.(*VersionDefContext).procedures, localctx.(*VersionDefContext).Get_procedureDef().GetV())
	p.SetState(349)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-12)&-(0x1f+1)) == 0 && ((1<<uint((_la-12)))&((1<<(XDRParserT__11-12))|(1<<(XDRParserT__12-12))|(1<<(XDRParserT__13-12))|(1<<(XDRParserT__14-12))|(1<<(XDRParserT__15-12))|(1<<(XDRParserT__16-12))|(1<<(XDRParserT__17-12))|(1<<(XDRParserT__18-12))|(1<<(XDRParserT__19-12))|(1<<(XDRParserT__24-12))|(1<<(XDRParserT__25-12))|(1<<(XDRParserIDENTIFIER-12)))) != 0 {
		{
			p.SetState(344)

			_x := p.ProcedureDef()

			localctx.(*VersionDefContext)._procedureDef = _x
		}
		localctx.(*VersionDefContext).procedures = append(localctx.(*VersionDefContext).procedures, localctx.(*VersionDefContext).Get_procedureDef().GetV())

		p.SetState(351)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(352)
		p.Match(XDRParserT__23)
	}
	{
		p.SetState(353)
		p.Match(XDRParserT__21)
	}
	{
		p.SetState(354)

		_x := p.Constant()

		localctx.(*VersionDefContext)._constant = _x
	}
	{
		p.SetState(355)
		p.Match(XDRParserT__1)
	}

	localctx.(*VersionDefContext).v = model.NewVersion((func() string {
		if localctx.(*VersionDefContext).Get_IDENTIFIER() == nil {
			return ""
		} else {
			return localctx.(*VersionDefContext).Get_IDENTIFIER().GetText()
		}
	}()), localctx.(*VersionDefContext).procedures, localctx.(*VersionDefContext).Get_constant().GetV())

	return localctx
}

// IProcedureDefContext is an interface to support dynamic dispatch.
type IProcedureDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_IDENTIFIER returns the _IDENTIFIER token.
	Get_IDENTIFIER() antlr.Token

	// Set_IDENTIFIER sets the _IDENTIFIER token.
	Set_IDENTIFIER(antlr.Token)

	// Get_procReturn returns the _procReturn rule contexts.
	Get_procReturn() IProcReturnContext

	// Get_procFirstArg returns the _procFirstArg rule contexts.
	Get_procFirstArg() IProcFirstArgContext

	// Get_typeSpecifier returns the _typeSpecifier rule contexts.
	Get_typeSpecifier() ITypeSpecifierContext

	// Get_constant returns the _constant rule contexts.
	Get_constant() IConstantContext

	// Set_procReturn sets the _procReturn rule contexts.
	Set_procReturn(IProcReturnContext)

	// Set_procFirstArg sets the _procFirstArg rule contexts.
	Set_procFirstArg(IProcFirstArgContext)

	// Set_typeSpecifier sets the _typeSpecifier rule contexts.
	Set_typeSpecifier(ITypeSpecifierContext)

	// Set_constant sets the _constant rule contexts.
	Set_constant(IConstantContext)

	// GetV returns the v attribute.
	GetV() model.Procedure

	// GetArguments returns the arguments attribute.
	GetArguments() []model.Type

	// SetV sets the v attribute.
	SetV(model.Procedure)

	// SetArguments sets the arguments attribute.
	SetArguments([]model.Type)

	// IsProcedureDefContext differentiates from other interfaces.
	IsProcedureDefContext()
}

type ProcedureDefContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	v              model.Procedure
	arguments      []model.Type
	_procReturn    IProcReturnContext
	_IDENTIFIER    antlr.Token
	_procFirstArg  IProcFirstArgContext
	_typeSpecifier ITypeSpecifierContext
	_constant      IConstantContext
}

func NewEmptyProcedureDefContext() *ProcedureDefContext {
	p := new(ProcedureDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XDRParserRULE_procedureDef
	return p
}

func (*ProcedureDefContext) IsProcedureDefContext() {}

func NewProcedureDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcedureDefContext {
	p := new(ProcedureDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XDRParserRULE_procedureDef

	return p
}

func (s *ProcedureDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcedureDefContext) Get_IDENTIFIER() antlr.Token { return s._IDENTIFIER }

func (s *ProcedureDefContext) Set_IDENTIFIER(v antlr.Token) { s._IDENTIFIER = v }

func (s *ProcedureDefContext) Get_procReturn() IProcReturnContext { return s._procReturn }

func (s *ProcedureDefContext) Get_procFirstArg() IProcFirstArgContext { return s._procFirstArg }

func (s *ProcedureDefContext) Get_typeSpecifier() ITypeSpecifierContext { return s._typeSpecifier }

func (s *ProcedureDefContext) Get_constant() IConstantContext { return s._constant }

func (s *ProcedureDefContext) Set_procReturn(v IProcReturnContext) { s._procReturn = v }

func (s *ProcedureDefContext) Set_procFirstArg(v IProcFirstArgContext) { s._procFirstArg = v }

func (s *ProcedureDefContext) Set_typeSpecifier(v ITypeSpecifierContext) { s._typeSpecifier = v }

func (s *ProcedureDefContext) Set_constant(v IConstantContext) { s._constant = v }

func (s *ProcedureDefContext) GetV() model.Procedure { return s.v }

func (s *ProcedureDefContext) GetArguments() []model.Type { return s.arguments }

func (s *ProcedureDefContext) SetV(v model.Procedure) { s.v = v }

func (s *ProcedureDefContext) SetArguments(v []model.Type) { s.arguments = v }

func (s *ProcedureDefContext) ProcReturn() IProcReturnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProcReturnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProcReturnContext)
}

func (s *ProcedureDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(XDRParserIDENTIFIER, 0)
}

func (s *ProcedureDefContext) ProcFirstArg() IProcFirstArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProcFirstArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProcFirstArgContext)
}

func (s *ProcedureDefContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ProcedureDefContext) AllTypeSpecifier() []ITypeSpecifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeSpecifierContext); ok {
			len++
		}
	}

	tst := make([]ITypeSpecifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeSpecifierContext); ok {
			tst[i] = t.(ITypeSpecifierContext)
			i++
		}
	}

	return tst
}

func (s *ProcedureDefContext) TypeSpecifier(i int) ITypeSpecifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeSpecifierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *ProcedureDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcedureDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcedureDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.EnterProcedureDef(s)
	}
}

func (s *ProcedureDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.ExitProcedureDef(s)
	}
}

func (p *XDRParser) ProcedureDef() (localctx IProcedureDefContext) {
	this := p
	_ = this

	localctx = NewProcedureDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, XDRParserRULE_procedureDef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(358)

		_x := p.ProcReturn()

		localctx.(*ProcedureDefContext)._procReturn = _x
	}
	{
		p.SetState(359)

		_m := p.Match(XDRParserIDENTIFIER)

		localctx.(*ProcedureDefContext)._IDENTIFIER = _m
	}
	{
		p.SetState(360)
		p.Match(XDRParserT__27)
	}
	{
		p.SetState(361)

		_x := p.ProcFirstArg()

		localctx.(*ProcedureDefContext)._procFirstArg = _x
	}

	if localctx.(*ProcedureDefContext).Get_procFirstArg().GetV() != nil {
		localctx.(*ProcedureDefContext).arguments = append(localctx.(*ProcedureDefContext).arguments, localctx.(*ProcedureDefContext).Get_procFirstArg().GetV())
	}

	p.SetState(369)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == XDRParserT__22 {
		{
			p.SetState(363)
			p.Match(XDRParserT__22)
		}
		{
			p.SetState(364)

			_x := p.TypeSpecifier()

			localctx.(*ProcedureDefContext)._typeSpecifier = _x
		}

		localctx.(*ProcedureDefContext).arguments = append(localctx.(*ProcedureDefContext).arguments, localctx.(*ProcedureDefContext).Get_typeSpecifier().GetV())

		p.SetState(371)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(372)
		p.Match(XDRParserT__28)
	}
	{
		p.SetState(373)
		p.Match(XDRParserT__21)
	}
	{
		p.SetState(374)

		_x := p.Constant()

		localctx.(*ProcedureDefContext)._constant = _x
	}
	{
		p.SetState(375)
		p.Match(XDRParserT__1)
	}

	localctx.(*ProcedureDefContext).v = model.NewProcedure((func() string {
		if localctx.(*ProcedureDefContext).Get_IDENTIFIER() == nil {
			return ""
		} else {
			return localctx.(*ProcedureDefContext).Get_IDENTIFIER().GetText()
		}
	}()), localctx.(*ProcedureDefContext).arguments, localctx.(*ProcedureDefContext).Get_procReturn().GetV(), localctx.(*ProcedureDefContext).Get_constant().GetV())

	return localctx
}

// IProcReturnContext is an interface to support dynamic dispatch.
type IProcReturnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_typeSpecifier returns the _typeSpecifier rule contexts.
	Get_typeSpecifier() ITypeSpecifierContext

	// Set_typeSpecifier sets the _typeSpecifier rule contexts.
	Set_typeSpecifier(ITypeSpecifierContext)

	// GetV returns the v attribute.
	GetV() model.Type

	// SetV sets the v attribute.
	SetV(model.Type)

	// IsProcReturnContext differentiates from other interfaces.
	IsProcReturnContext()
}

type ProcReturnContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	v              model.Type
	_typeSpecifier ITypeSpecifierContext
}

func NewEmptyProcReturnContext() *ProcReturnContext {
	p := new(ProcReturnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XDRParserRULE_procReturn
	return p
}

func (*ProcReturnContext) IsProcReturnContext() {}

func NewProcReturnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcReturnContext {
	p := new(ProcReturnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XDRParserRULE_procReturn

	return p
}

func (s *ProcReturnContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcReturnContext) Get_typeSpecifier() ITypeSpecifierContext { return s._typeSpecifier }

func (s *ProcReturnContext) Set_typeSpecifier(v ITypeSpecifierContext) { s._typeSpecifier = v }

func (s *ProcReturnContext) GetV() model.Type { return s.v }

func (s *ProcReturnContext) SetV(v model.Type) { s.v = v }

func (s *ProcReturnContext) TypeSpecifier() ITypeSpecifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeSpecifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *ProcReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcReturnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.EnterProcReturn(s)
	}
}

func (s *ProcReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.ExitProcReturn(s)
	}
}

func (p *XDRParser) ProcReturn() (localctx IProcReturnContext) {
	this := p
	_ = this

	localctx = NewProcReturnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, XDRParserRULE_procReturn)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(382)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case XDRParserT__11:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(378)
			p.Match(XDRParserT__11)
		}

	case XDRParserT__12, XDRParserT__13, XDRParserT__14, XDRParserT__15, XDRParserT__16, XDRParserT__17, XDRParserT__18, XDRParserT__19, XDRParserT__24, XDRParserT__25, XDRParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(379)

			_x := p.TypeSpecifier()

			localctx.(*ProcReturnContext)._typeSpecifier = _x
		}
		localctx.(*ProcReturnContext).v = localctx.(*ProcReturnContext).Get_typeSpecifier().GetV()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IProcFirstArgContext is an interface to support dynamic dispatch.
type IProcFirstArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_typeSpecifier returns the _typeSpecifier rule contexts.
	Get_typeSpecifier() ITypeSpecifierContext

	// Set_typeSpecifier sets the _typeSpecifier rule contexts.
	Set_typeSpecifier(ITypeSpecifierContext)

	// GetV returns the v attribute.
	GetV() model.Type

	// SetV sets the v attribute.
	SetV(model.Type)

	// IsProcFirstArgContext differentiates from other interfaces.
	IsProcFirstArgContext()
}

type ProcFirstArgContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	v              model.Type
	_typeSpecifier ITypeSpecifierContext
}

func NewEmptyProcFirstArgContext() *ProcFirstArgContext {
	p := new(ProcFirstArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XDRParserRULE_procFirstArg
	return p
}

func (*ProcFirstArgContext) IsProcFirstArgContext() {}

func NewProcFirstArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcFirstArgContext {
	p := new(ProcFirstArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XDRParserRULE_procFirstArg

	return p
}

func (s *ProcFirstArgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcFirstArgContext) Get_typeSpecifier() ITypeSpecifierContext { return s._typeSpecifier }

func (s *ProcFirstArgContext) Set_typeSpecifier(v ITypeSpecifierContext) { s._typeSpecifier = v }

func (s *ProcFirstArgContext) GetV() model.Type { return s.v }

func (s *ProcFirstArgContext) SetV(v model.Type) { s.v = v }

func (s *ProcFirstArgContext) TypeSpecifier() ITypeSpecifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeSpecifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *ProcFirstArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcFirstArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcFirstArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.EnterProcFirstArg(s)
	}
}

func (s *ProcFirstArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XDRListener); ok {
		listenerT.ExitProcFirstArg(s)
	}
}

func (p *XDRParser) ProcFirstArg() (localctx IProcFirstArgContext) {
	this := p
	_ = this

	localctx = NewProcFirstArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, XDRParserRULE_procFirstArg)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(388)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case XDRParserT__11:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(384)
			p.Match(XDRParserT__11)
		}

	case XDRParserT__12, XDRParserT__13, XDRParserT__14, XDRParserT__15, XDRParserT__16, XDRParserT__17, XDRParserT__18, XDRParserT__19, XDRParserT__24, XDRParserT__25, XDRParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(385)

			_x := p.TypeSpecifier()

			localctx.(*ProcFirstArgContext)._typeSpecifier = _x
		}
		localctx.(*ProcFirstArgContext).v = localctx.(*ProcFirstArgContext).Get_typeSpecifier().GetV()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
