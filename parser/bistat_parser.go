// Code generated from parser/Bistat.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Bistat

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type BistatParser struct {
	*antlr.BaseParser
}

var bistatParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func bistatParserInit() {
	staticData := &bistatParserStaticData
	staticData.literalNames = []string{
		"", "'Program'", "';'", "'var'", "'func'", "'('", "')'", "':'", "'{'",
		"'}'", "'main'", "'='", "'['", "','", "']'", "'#'", "'for'", "'in'",
		"'while'", "'if'", "'else'", "'return'", "'read'", "'print'", "'listAdd'",
		"'listPop'", "'length'", "'range'", "'plot'", "'sum'", "'min'", "'prod'",
		"'avg'", "'sMode'", "'median'", "'sin'", "'tan'", "'cos'", "'sort'",
		"'sqrt'", "'floor'", "'ceil'", "'abs'", "'not'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "WS", "CARDINALITY", "TYPE_PRIMITIVE",
		"ID", "VAR_CONS", "INT_CONS", "NUMBER", "BOOL_CONS", "STRING_CONS",
		"FLOAT_CONS", "OP_SEC", "OP_FIRST", "LOGIC_OPERATOR",
	}
	staticData.ruleNames = []string{
		"program", "varDeclaration", "var_type", "funcDef", "paramDeclaration",
		"main", "stmt", "assignment", "nested_stmt", "matrixAssignment", "listAssignment",
		"comment", "forLoop", "whileLoop", "conditional", "returnStmt", "specialFunction",
		"inputRead", "print", "listAdd", "listPop", "length", "range", "plot",
		"sum", "min", "prod", "avg", "sMode", "median", "sin", "tan", "cos",
		"sort", "sqrt", "floor", "ceil", "abs", "not", "expression", "exp",
		"term", "factor", "functionCall", "indexing",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 56, 500, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 95,
		8, 0, 10, 0, 12, 0, 98, 9, 0, 1, 0, 5, 0, 101, 8, 0, 10, 0, 12, 0, 104,
		9, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 3, 2,
		116, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 122, 8, 3, 10, 3, 12, 3, 125,
		9, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 131, 8, 3, 10, 3, 12, 3, 134, 9, 3,
		1, 3, 1, 3, 4, 3, 138, 8, 3, 11, 3, 12, 3, 139, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 4, 5, 154, 8, 5, 11, 5,
		12, 5, 155, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 164, 8, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 172, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 3, 7, 179, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 186, 8, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 193, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 5,
		9, 199, 8, 9, 10, 9, 12, 9, 202, 9, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10,
		1, 10, 5, 10, 210, 8, 10, 10, 10, 12, 10, 213, 9, 10, 1, 10, 1, 10, 1,
		11, 1, 11, 4, 11, 219, 8, 11, 11, 11, 12, 11, 220, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 4, 12, 233, 8, 12, 11,
		12, 12, 12, 234, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		4, 13, 245, 8, 13, 11, 13, 12, 13, 246, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 4, 14, 257, 8, 14, 11, 14, 12, 14, 258, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 4, 14, 269, 8, 14, 11,
		14, 12, 14, 270, 1, 14, 1, 14, 5, 14, 275, 8, 14, 10, 14, 12, 14, 278,
		9, 14, 1, 14, 1, 14, 1, 14, 4, 14, 283, 8, 14, 11, 14, 12, 14, 284, 1,
		14, 1, 14, 3, 14, 289, 8, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 316,
		8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 323, 8, 17, 10, 17, 12,
		17, 326, 9, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18,
		335, 8, 18, 10, 18, 12, 18, 338, 9, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 364,
		8, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1,
		32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1,
		37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39,
		1, 39, 1, 39, 3, 39, 451, 8, 39, 1, 40, 1, 40, 1, 40, 3, 40, 456, 8, 40,
		1, 41, 1, 41, 1, 41, 3, 41, 461, 8, 41, 1, 42, 3, 42, 464, 8, 42, 1, 42,
		1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 3, 42, 474, 8, 42, 1,
		43, 1, 43, 1, 43, 1, 43, 1, 43, 5, 43, 481, 8, 43, 10, 43, 12, 43, 484,
		9, 43, 3, 43, 486, 8, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1,
		44, 1, 44, 1, 44, 1, 44, 3, 44, 498, 8, 44, 1, 44, 0, 0, 45, 0, 2, 4, 6,
		8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
		44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78,
		80, 82, 84, 86, 88, 0, 1, 1, 0, 15, 15, 522, 0, 90, 1, 0, 0, 0, 2, 108,
		1, 0, 0, 0, 4, 113, 1, 0, 0, 0, 6, 117, 1, 0, 0, 0, 8, 143, 1, 0, 0, 0,
		10, 148, 1, 0, 0, 0, 12, 171, 1, 0, 0, 0, 14, 173, 1, 0, 0, 0, 16, 192,
		1, 0, 0, 0, 18, 194, 1, 0, 0, 0, 20, 205, 1, 0, 0, 0, 22, 216, 1, 0, 0,
		0, 24, 224, 1, 0, 0, 0, 26, 238, 1, 0, 0, 0, 28, 250, 1, 0, 0, 0, 30, 290,
		1, 0, 0, 0, 32, 315, 1, 0, 0, 0, 34, 317, 1, 0, 0, 0, 36, 329, 1, 0, 0,
		0, 38, 341, 1, 0, 0, 0, 40, 348, 1, 0, 0, 0, 42, 353, 1, 0, 0, 0, 44, 358,
		1, 0, 0, 0, 46, 367, 1, 0, 0, 0, 48, 372, 1, 0, 0, 0, 50, 377, 1, 0, 0,
		0, 52, 382, 1, 0, 0, 0, 54, 387, 1, 0, 0, 0, 56, 392, 1, 0, 0, 0, 58, 397,
		1, 0, 0, 0, 60, 402, 1, 0, 0, 0, 62, 407, 1, 0, 0, 0, 64, 412, 1, 0, 0,
		0, 66, 417, 1, 0, 0, 0, 68, 422, 1, 0, 0, 0, 70, 427, 1, 0, 0, 0, 72, 432,
		1, 0, 0, 0, 74, 437, 1, 0, 0, 0, 76, 442, 1, 0, 0, 0, 78, 447, 1, 0, 0,
		0, 80, 452, 1, 0, 0, 0, 82, 457, 1, 0, 0, 0, 84, 463, 1, 0, 0, 0, 86, 475,
		1, 0, 0, 0, 88, 489, 1, 0, 0, 0, 90, 91, 5, 1, 0, 0, 91, 92, 5, 47, 0,
		0, 92, 96, 5, 2, 0, 0, 93, 95, 3, 2, 1, 0, 94, 93, 1, 0, 0, 0, 95, 98,
		1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 102, 1, 0, 0, 0,
		98, 96, 1, 0, 0, 0, 99, 101, 3, 6, 3, 0, 100, 99, 1, 0, 0, 0, 101, 104,
		1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 105, 1, 0,
		0, 0, 104, 102, 1, 0, 0, 0, 105, 106, 3, 10, 5, 0, 106, 107, 5, 0, 0, 1,
		107, 1, 1, 0, 0, 0, 108, 109, 5, 3, 0, 0, 109, 110, 3, 4, 2, 0, 110, 111,
		5, 47, 0, 0, 111, 112, 5, 2, 0, 0, 112, 3, 1, 0, 0, 0, 113, 115, 5, 46,
		0, 0, 114, 116, 5, 45, 0, 0, 115, 114, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0,
		116, 5, 1, 0, 0, 0, 117, 118, 5, 4, 0, 0, 118, 119, 5, 47, 0, 0, 119, 123,
		5, 5, 0, 0, 120, 122, 3, 8, 4, 0, 121, 120, 1, 0, 0, 0, 122, 125, 1, 0,
		0, 0, 123, 121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 126, 1, 0, 0, 0,
		125, 123, 1, 0, 0, 0, 126, 127, 5, 6, 0, 0, 127, 128, 5, 7, 0, 0, 128,
		132, 3, 4, 2, 0, 129, 131, 3, 2, 1, 0, 130, 129, 1, 0, 0, 0, 131, 134,
		1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 135, 1, 0,
		0, 0, 134, 132, 1, 0, 0, 0, 135, 137, 5, 8, 0, 0, 136, 138, 3, 12, 6, 0,
		137, 136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139,
		140, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 142, 5, 9, 0, 0, 142, 7, 1,
		0, 0, 0, 143, 144, 5, 3, 0, 0, 144, 145, 3, 4, 2, 0, 145, 146, 5, 47, 0,
		0, 146, 147, 5, 2, 0, 0, 147, 9, 1, 0, 0, 0, 148, 149, 5, 10, 0, 0, 149,
		150, 5, 5, 0, 0, 150, 151, 5, 6, 0, 0, 151, 153, 5, 8, 0, 0, 152, 154,
		3, 12, 6, 0, 153, 152, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 153, 1, 0,
		0, 0, 155, 156, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 158, 5, 9, 0, 0,
		158, 11, 1, 0, 0, 0, 159, 164, 3, 14, 7, 0, 160, 164, 3, 32, 16, 0, 161,
		164, 3, 86, 43, 0, 162, 164, 3, 30, 15, 0, 163, 159, 1, 0, 0, 0, 163, 160,
		1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 163, 162, 1, 0, 0, 0, 164, 165, 1, 0,
		0, 0, 165, 166, 5, 2, 0, 0, 166, 172, 1, 0, 0, 0, 167, 172, 3, 28, 14,
		0, 168, 172, 3, 26, 13, 0, 169, 172, 3, 24, 12, 0, 170, 172, 3, 22, 11,
		0, 171, 163, 1, 0, 0, 0, 171, 167, 1, 0, 0, 0, 171, 168, 1, 0, 0, 0, 171,
		169, 1, 0, 0, 0, 171, 170, 1, 0, 0, 0, 172, 13, 1, 0, 0, 0, 173, 174, 5,
		47, 0, 0, 174, 178, 5, 11, 0, 0, 175, 179, 5, 48, 0, 0, 176, 179, 3, 20,
		10, 0, 177, 179, 3, 18, 9, 0, 178, 175, 1, 0, 0, 0, 178, 176, 1, 0, 0,
		0, 178, 177, 1, 0, 0, 0, 179, 15, 1, 0, 0, 0, 180, 186, 3, 14, 7, 0, 181,
		186, 3, 36, 18, 0, 182, 186, 3, 32, 16, 0, 183, 186, 3, 86, 43, 0, 184,
		186, 3, 30, 15, 0, 185, 180, 1, 0, 0, 0, 185, 181, 1, 0, 0, 0, 185, 182,
		1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 185, 184, 1, 0, 0, 0, 186, 187, 1, 0,
		0, 0, 187, 188, 5, 2, 0, 0, 188, 193, 1, 0, 0, 0, 189, 193, 3, 28, 14,
		0, 190, 193, 3, 26, 13, 0, 191, 193, 3, 22, 11, 0, 192, 185, 1, 0, 0, 0,
		192, 189, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 192, 191, 1, 0, 0, 0, 193,
		17, 1, 0, 0, 0, 194, 195, 5, 12, 0, 0, 195, 200, 3, 20, 10, 0, 196, 197,
		5, 13, 0, 0, 197, 199, 3, 20, 10, 0, 198, 196, 1, 0, 0, 0, 199, 202, 1,
		0, 0, 0, 200, 198, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 203, 1, 0, 0,
		0, 202, 200, 1, 0, 0, 0, 203, 204, 5, 14, 0, 0, 204, 19, 1, 0, 0, 0, 205,
		206, 5, 12, 0, 0, 206, 211, 5, 48, 0, 0, 207, 208, 5, 13, 0, 0, 208, 210,
		5, 48, 0, 0, 209, 207, 1, 0, 0, 0, 210, 213, 1, 0, 0, 0, 211, 209, 1, 0,
		0, 0, 211, 212, 1, 0, 0, 0, 212, 214, 1, 0, 0, 0, 213, 211, 1, 0, 0, 0,
		214, 215, 5, 14, 0, 0, 215, 21, 1, 0, 0, 0, 216, 218, 5, 15, 0, 0, 217,
		219, 8, 0, 0, 0, 218, 217, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220, 218,
		1, 0, 0, 0, 220, 221, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 223, 5, 15,
		0, 0, 223, 23, 1, 0, 0, 0, 224, 225, 5, 16, 0, 0, 225, 226, 5, 5, 0, 0,
		226, 227, 5, 47, 0, 0, 227, 228, 5, 17, 0, 0, 228, 229, 3, 78, 39, 0, 229,
		230, 5, 6, 0, 0, 230, 232, 5, 8, 0, 0, 231, 233, 3, 16, 8, 0, 232, 231,
		1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 232, 1, 0, 0, 0, 234, 235, 1, 0,
		0, 0, 235, 236, 1, 0, 0, 0, 236, 237, 5, 9, 0, 0, 237, 25, 1, 0, 0, 0,
		238, 239, 5, 18, 0, 0, 239, 240, 5, 5, 0, 0, 240, 241, 3, 78, 39, 0, 241,
		242, 5, 6, 0, 0, 242, 244, 5, 8, 0, 0, 243, 245, 3, 12, 6, 0, 244, 243,
		1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 244, 1, 0, 0, 0, 246, 247, 1, 0,
		0, 0, 247, 248, 1, 0, 0, 0, 248, 249, 5, 9, 0, 0, 249, 27, 1, 0, 0, 0,
		250, 251, 5, 19, 0, 0, 251, 252, 5, 5, 0, 0, 252, 253, 3, 78, 39, 0, 253,
		254, 5, 6, 0, 0, 254, 256, 5, 8, 0, 0, 255, 257, 3, 12, 6, 0, 256, 255,
		1, 0, 0, 0, 257, 258, 1, 0, 0, 0, 258, 256, 1, 0, 0, 0, 258, 259, 1, 0,
		0, 0, 259, 260, 1, 0, 0, 0, 260, 276, 5, 9, 0, 0, 261, 262, 5, 20, 0, 0,
		262, 263, 5, 19, 0, 0, 263, 264, 5, 5, 0, 0, 264, 265, 3, 78, 39, 0, 265,
		266, 5, 6, 0, 0, 266, 268, 5, 8, 0, 0, 267, 269, 3, 12, 6, 0, 268, 267,
		1, 0, 0, 0, 269, 270, 1, 0, 0, 0, 270, 268, 1, 0, 0, 0, 270, 271, 1, 0,
		0, 0, 271, 272, 1, 0, 0, 0, 272, 273, 5, 9, 0, 0, 273, 275, 1, 0, 0, 0,
		274, 261, 1, 0, 0, 0, 275, 278, 1, 0, 0, 0, 276, 274, 1, 0, 0, 0, 276,
		277, 1, 0, 0, 0, 277, 288, 1, 0, 0, 0, 278, 276, 1, 0, 0, 0, 279, 280,
		5, 20, 0, 0, 280, 282, 5, 8, 0, 0, 281, 283, 3, 12, 6, 0, 282, 281, 1,
		0, 0, 0, 283, 284, 1, 0, 0, 0, 284, 282, 1, 0, 0, 0, 284, 285, 1, 0, 0,
		0, 285, 286, 1, 0, 0, 0, 286, 287, 5, 9, 0, 0, 287, 289, 1, 0, 0, 0, 288,
		279, 1, 0, 0, 0, 288, 289, 1, 0, 0, 0, 289, 29, 1, 0, 0, 0, 290, 291, 5,
		21, 0, 0, 291, 292, 3, 78, 39, 0, 292, 31, 1, 0, 0, 0, 293, 316, 3, 34,
		17, 0, 294, 316, 3, 36, 18, 0, 295, 316, 3, 38, 19, 0, 296, 316, 3, 40,
		20, 0, 297, 316, 3, 42, 21, 0, 298, 316, 3, 44, 22, 0, 299, 316, 3, 46,
		23, 0, 300, 316, 3, 48, 24, 0, 301, 316, 3, 50, 25, 0, 302, 316, 3, 52,
		26, 0, 303, 316, 3, 54, 27, 0, 304, 316, 3, 56, 28, 0, 305, 316, 3, 58,
		29, 0, 306, 316, 3, 60, 30, 0, 307, 316, 3, 64, 32, 0, 308, 316, 3, 62,
		31, 0, 309, 316, 3, 66, 33, 0, 310, 316, 3, 68, 34, 0, 311, 316, 3, 70,
		35, 0, 312, 316, 3, 72, 36, 0, 313, 316, 3, 74, 37, 0, 314, 316, 3, 76,
		38, 0, 315, 293, 1, 0, 0, 0, 315, 294, 1, 0, 0, 0, 315, 295, 1, 0, 0, 0,
		315, 296, 1, 0, 0, 0, 315, 297, 1, 0, 0, 0, 315, 298, 1, 0, 0, 0, 315,
		299, 1, 0, 0, 0, 315, 300, 1, 0, 0, 0, 315, 301, 1, 0, 0, 0, 315, 302,
		1, 0, 0, 0, 315, 303, 1, 0, 0, 0, 315, 304, 1, 0, 0, 0, 315, 305, 1, 0,
		0, 0, 315, 306, 1, 0, 0, 0, 315, 307, 1, 0, 0, 0, 315, 308, 1, 0, 0, 0,
		315, 309, 1, 0, 0, 0, 315, 310, 1, 0, 0, 0, 315, 311, 1, 0, 0, 0, 315,
		312, 1, 0, 0, 0, 315, 313, 1, 0, 0, 0, 315, 314, 1, 0, 0, 0, 316, 33, 1,
		0, 0, 0, 317, 318, 5, 22, 0, 0, 318, 319, 5, 5, 0, 0, 319, 324, 5, 47,
		0, 0, 320, 321, 5, 13, 0, 0, 321, 323, 5, 47, 0, 0, 322, 320, 1, 0, 0,
		0, 323, 326, 1, 0, 0, 0, 324, 322, 1, 0, 0, 0, 324, 325, 1, 0, 0, 0, 325,
		327, 1, 0, 0, 0, 326, 324, 1, 0, 0, 0, 327, 328, 5, 6, 0, 0, 328, 35, 1,
		0, 0, 0, 329, 330, 5, 23, 0, 0, 330, 331, 5, 5, 0, 0, 331, 336, 3, 78,
		39, 0, 332, 333, 5, 13, 0, 0, 333, 335, 3, 78, 39, 0, 334, 332, 1, 0, 0,
		0, 335, 338, 1, 0, 0, 0, 336, 334, 1, 0, 0, 0, 336, 337, 1, 0, 0, 0, 337,
		339, 1, 0, 0, 0, 338, 336, 1, 0, 0, 0, 339, 340, 5, 6, 0, 0, 340, 37, 1,
		0, 0, 0, 341, 342, 5, 24, 0, 0, 342, 343, 5, 5, 0, 0, 343, 344, 3, 78,
		39, 0, 344, 345, 5, 13, 0, 0, 345, 346, 3, 78, 39, 0, 346, 347, 5, 6, 0,
		0, 347, 39, 1, 0, 0, 0, 348, 349, 5, 25, 0, 0, 349, 350, 5, 5, 0, 0, 350,
		351, 3, 78, 39, 0, 351, 352, 5, 6, 0, 0, 352, 41, 1, 0, 0, 0, 353, 354,
		5, 26, 0, 0, 354, 355, 5, 5, 0, 0, 355, 356, 3, 78, 39, 0, 356, 357, 5,
		6, 0, 0, 357, 43, 1, 0, 0, 0, 358, 359, 5, 27, 0, 0, 359, 360, 5, 5, 0,
		0, 360, 363, 3, 78, 39, 0, 361, 362, 5, 13, 0, 0, 362, 364, 3, 78, 39,
		0, 363, 361, 1, 0, 0, 0, 363, 364, 1, 0, 0, 0, 364, 365, 1, 0, 0, 0, 365,
		366, 5, 6, 0, 0, 366, 45, 1, 0, 0, 0, 367, 368, 5, 28, 0, 0, 368, 369,
		5, 5, 0, 0, 369, 370, 3, 78, 39, 0, 370, 371, 5, 6, 0, 0, 371, 47, 1, 0,
		0, 0, 372, 373, 5, 29, 0, 0, 373, 374, 5, 5, 0, 0, 374, 375, 3, 78, 39,
		0, 375, 376, 5, 6, 0, 0, 376, 49, 1, 0, 0, 0, 377, 378, 5, 30, 0, 0, 378,
		379, 5, 5, 0, 0, 379, 380, 3, 78, 39, 0, 380, 381, 5, 6, 0, 0, 381, 51,
		1, 0, 0, 0, 382, 383, 5, 31, 0, 0, 383, 384, 5, 5, 0, 0, 384, 385, 3, 78,
		39, 0, 385, 386, 5, 6, 0, 0, 386, 53, 1, 0, 0, 0, 387, 388, 5, 32, 0, 0,
		388, 389, 5, 5, 0, 0, 389, 390, 3, 78, 39, 0, 390, 391, 5, 6, 0, 0, 391,
		55, 1, 0, 0, 0, 392, 393, 5, 33, 0, 0, 393, 394, 5, 5, 0, 0, 394, 395,
		3, 78, 39, 0, 395, 396, 5, 6, 0, 0, 396, 57, 1, 0, 0, 0, 397, 398, 5, 34,
		0, 0, 398, 399, 5, 5, 0, 0, 399, 400, 3, 78, 39, 0, 400, 401, 5, 6, 0,
		0, 401, 59, 1, 0, 0, 0, 402, 403, 5, 35, 0, 0, 403, 404, 5, 5, 0, 0, 404,
		405, 3, 78, 39, 0, 405, 406, 5, 6, 0, 0, 406, 61, 1, 0, 0, 0, 407, 408,
		5, 36, 0, 0, 408, 409, 5, 5, 0, 0, 409, 410, 3, 78, 39, 0, 410, 411, 5,
		6, 0, 0, 411, 63, 1, 0, 0, 0, 412, 413, 5, 37, 0, 0, 413, 414, 5, 5, 0,
		0, 414, 415, 3, 78, 39, 0, 415, 416, 5, 6, 0, 0, 416, 65, 1, 0, 0, 0, 417,
		418, 5, 38, 0, 0, 418, 419, 5, 5, 0, 0, 419, 420, 3, 78, 39, 0, 420, 421,
		5, 6, 0, 0, 421, 67, 1, 0, 0, 0, 422, 423, 5, 39, 0, 0, 423, 424, 5, 5,
		0, 0, 424, 425, 3, 78, 39, 0, 425, 426, 5, 6, 0, 0, 426, 69, 1, 0, 0, 0,
		427, 428, 5, 40, 0, 0, 428, 429, 5, 5, 0, 0, 429, 430, 3, 78, 39, 0, 430,
		431, 5, 6, 0, 0, 431, 71, 1, 0, 0, 0, 432, 433, 5, 41, 0, 0, 433, 434,
		5, 5, 0, 0, 434, 435, 3, 78, 39, 0, 435, 436, 5, 6, 0, 0, 436, 73, 1, 0,
		0, 0, 437, 438, 5, 42, 0, 0, 438, 439, 5, 5, 0, 0, 439, 440, 3, 78, 39,
		0, 440, 441, 5, 6, 0, 0, 441, 75, 1, 0, 0, 0, 442, 443, 5, 43, 0, 0, 443,
		444, 5, 5, 0, 0, 444, 445, 3, 78, 39, 0, 445, 446, 5, 6, 0, 0, 446, 77,
		1, 0, 0, 0, 447, 450, 3, 80, 40, 0, 448, 449, 5, 56, 0, 0, 449, 451, 3,
		80, 40, 0, 450, 448, 1, 0, 0, 0, 450, 451, 1, 0, 0, 0, 451, 79, 1, 0, 0,
		0, 452, 455, 3, 82, 41, 0, 453, 454, 5, 54, 0, 0, 454, 456, 3, 80, 40,
		0, 455, 453, 1, 0, 0, 0, 455, 456, 1, 0, 0, 0, 456, 81, 1, 0, 0, 0, 457,
		460, 3, 84, 42, 0, 458, 459, 5, 55, 0, 0, 459, 461, 3, 82, 41, 0, 460,
		458, 1, 0, 0, 0, 460, 461, 1, 0, 0, 0, 461, 83, 1, 0, 0, 0, 462, 464, 5,
		54, 0, 0, 463, 462, 1, 0, 0, 0, 463, 464, 1, 0, 0, 0, 464, 473, 1, 0, 0,
		0, 465, 466, 5, 5, 0, 0, 466, 467, 3, 78, 39, 0, 467, 468, 5, 6, 0, 0,
		468, 474, 1, 0, 0, 0, 469, 474, 3, 88, 44, 0, 470, 474, 3, 32, 16, 0, 471,
		474, 3, 86, 43, 0, 472, 474, 5, 48, 0, 0, 473, 465, 1, 0, 0, 0, 473, 469,
		1, 0, 0, 0, 473, 470, 1, 0, 0, 0, 473, 471, 1, 0, 0, 0, 473, 472, 1, 0,
		0, 0, 474, 85, 1, 0, 0, 0, 475, 476, 5, 47, 0, 0, 476, 485, 5, 5, 0, 0,
		477, 482, 3, 78, 39, 0, 478, 479, 5, 13, 0, 0, 479, 481, 3, 78, 39, 0,
		480, 478, 1, 0, 0, 0, 481, 484, 1, 0, 0, 0, 482, 480, 1, 0, 0, 0, 482,
		483, 1, 0, 0, 0, 483, 486, 1, 0, 0, 0, 484, 482, 1, 0, 0, 0, 485, 477,
		1, 0, 0, 0, 485, 486, 1, 0, 0, 0, 486, 487, 1, 0, 0, 0, 487, 488, 5, 6,
		0, 0, 488, 87, 1, 0, 0, 0, 489, 490, 5, 47, 0, 0, 490, 491, 5, 12, 0, 0,
		491, 492, 3, 78, 39, 0, 492, 497, 5, 14, 0, 0, 493, 494, 5, 12, 0, 0, 494,
		495, 3, 78, 39, 0, 495, 496, 5, 14, 0, 0, 496, 498, 1, 0, 0, 0, 497, 493,
		1, 0, 0, 0, 497, 498, 1, 0, 0, 0, 498, 89, 1, 0, 0, 0, 34, 96, 102, 115,
		123, 132, 139, 155, 163, 171, 178, 185, 192, 200, 211, 220, 234, 246, 258,
		270, 276, 284, 288, 315, 324, 336, 363, 450, 455, 460, 463, 473, 482, 485,
		497,
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

// BistatParserInit initializes any static state used to implement BistatParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBistatParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BistatParserInit() {
	staticData := &bistatParserStaticData
	staticData.once.Do(bistatParserInit)
}

// NewBistatParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBistatParser(input antlr.TokenStream) *BistatParser {
	BistatParserInit()
	this := new(BistatParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &bistatParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Bistat.g4"

	return this
}

// BistatParser tokens.
const (
	BistatParserEOF            = antlr.TokenEOF
	BistatParserT__0           = 1
	BistatParserT__1           = 2
	BistatParserT__2           = 3
	BistatParserT__3           = 4
	BistatParserT__4           = 5
	BistatParserT__5           = 6
	BistatParserT__6           = 7
	BistatParserT__7           = 8
	BistatParserT__8           = 9
	BistatParserT__9           = 10
	BistatParserT__10          = 11
	BistatParserT__11          = 12
	BistatParserT__12          = 13
	BistatParserT__13          = 14
	BistatParserT__14          = 15
	BistatParserT__15          = 16
	BistatParserT__16          = 17
	BistatParserT__17          = 18
	BistatParserT__18          = 19
	BistatParserT__19          = 20
	BistatParserT__20          = 21
	BistatParserT__21          = 22
	BistatParserT__22          = 23
	BistatParserT__23          = 24
	BistatParserT__24          = 25
	BistatParserT__25          = 26
	BistatParserT__26          = 27
	BistatParserT__27          = 28
	BistatParserT__28          = 29
	BistatParserT__29          = 30
	BistatParserT__30          = 31
	BistatParserT__31          = 32
	BistatParserT__32          = 33
	BistatParserT__33          = 34
	BistatParserT__34          = 35
	BistatParserT__35          = 36
	BistatParserT__36          = 37
	BistatParserT__37          = 38
	BistatParserT__38          = 39
	BistatParserT__39          = 40
	BistatParserT__40          = 41
	BistatParserT__41          = 42
	BistatParserT__42          = 43
	BistatParserWS             = 44
	BistatParserCARDINALITY    = 45
	BistatParserTYPE_PRIMITIVE = 46
	BistatParserID             = 47
	BistatParserVAR_CONS       = 48
	BistatParserINT_CONS       = 49
	BistatParserNUMBER         = 50
	BistatParserBOOL_CONS      = 51
	BistatParserSTRING_CONS    = 52
	BistatParserFLOAT_CONS     = 53
	BistatParserOP_SEC         = 54
	BistatParserOP_FIRST       = 55
	BistatParserLOGIC_OPERATOR = 56
)

// BistatParser rules.
const (
	BistatParserRULE_program          = 0
	BistatParserRULE_varDeclaration   = 1
	BistatParserRULE_var_type         = 2
	BistatParserRULE_funcDef          = 3
	BistatParserRULE_paramDeclaration = 4
	BistatParserRULE_main             = 5
	BistatParserRULE_stmt             = 6
	BistatParserRULE_assignment       = 7
	BistatParserRULE_nested_stmt      = 8
	BistatParserRULE_matrixAssignment = 9
	BistatParserRULE_listAssignment   = 10
	BistatParserRULE_comment          = 11
	BistatParserRULE_forLoop          = 12
	BistatParserRULE_whileLoop        = 13
	BistatParserRULE_conditional      = 14
	BistatParserRULE_returnStmt       = 15
	BistatParserRULE_specialFunction  = 16
	BistatParserRULE_inputRead        = 17
	BistatParserRULE_print            = 18
	BistatParserRULE_listAdd          = 19
	BistatParserRULE_listPop          = 20
	BistatParserRULE_length           = 21
	BistatParserRULE_range            = 22
	BistatParserRULE_plot             = 23
	BistatParserRULE_sum              = 24
	BistatParserRULE_min              = 25
	BistatParserRULE_prod             = 26
	BistatParserRULE_avg              = 27
	BistatParserRULE_sMode            = 28
	BistatParserRULE_median           = 29
	BistatParserRULE_sin              = 30
	BistatParserRULE_tan              = 31
	BistatParserRULE_cos              = 32
	BistatParserRULE_sort             = 33
	BistatParserRULE_sqrt             = 34
	BistatParserRULE_floor            = 35
	BistatParserRULE_ceil             = 36
	BistatParserRULE_abs              = 37
	BistatParserRULE_not              = 38
	BistatParserRULE_expression       = 39
	BistatParserRULE_exp              = 40
	BistatParserRULE_term             = 41
	BistatParserRULE_factor           = 42
	BistatParserRULE_functionCall     = 43
	BistatParserRULE_indexing         = 44
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Main() IMainContext
	EOF() antlr.TerminalNode
	AllVarDeclaration() []IVarDeclarationContext
	VarDeclaration(i int) IVarDeclarationContext
	AllFuncDef() []IFuncDefContext
	FuncDef(i int) IFuncDefContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) ID() antlr.TerminalNode {
	return s.GetToken(BistatParserID, 0)
}

func (s *ProgramContext) Main() IMainContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMainContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMainContext)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(BistatParserEOF, 0)
}

func (s *ProgramContext) AllVarDeclaration() []IVarDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IVarDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarDeclarationContext); ok {
			tst[i] = t.(IVarDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) VarDeclaration(i int) IVarDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclarationContext); ok {
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

	return t.(IVarDeclarationContext)
}

func (s *ProgramContext) AllFuncDef() []IFuncDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncDefContext); ok {
			len++
		}
	}

	tst := make([]IFuncDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncDefContext); ok {
			tst[i] = t.(IFuncDefContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) FuncDef(i int) IFuncDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDefContext); ok {
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

	return t.(IFuncDefContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *BistatParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BistatParserRULE_program)
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
		p.SetState(90)
		p.Match(BistatParserT__0)
	}
	{
		p.SetState(91)
		p.Match(BistatParserID)
	}
	{
		p.SetState(92)
		p.Match(BistatParserT__1)
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__2 {
		{
			p.SetState(93)
			p.VarDeclaration()
		}

		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__3 {
		{
			p.SetState(99)
			p.FuncDef()
		}

		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(105)
		p.Main()
	}
	{
		p.SetState(106)
		p.Match(BistatParserEOF)
	}

	return localctx
}

// IVarDeclarationContext is an interface to support dynamic dispatch.
type IVarDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var_type() IVar_typeContext
	ID() antlr.TerminalNode

	// IsVarDeclarationContext differentiates from other interfaces.
	IsVarDeclarationContext()
}

type VarDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclarationContext() *VarDeclarationContext {
	var p = new(VarDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_varDeclaration
	return p
}

func (*VarDeclarationContext) IsVarDeclarationContext() {}

func NewVarDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclarationContext {
	var p = new(VarDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_varDeclaration

	return p
}

func (s *VarDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclarationContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *VarDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(BistatParserID, 0)
}

func (s *VarDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterVarDeclaration(s)
	}
}

func (s *VarDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitVarDeclaration(s)
	}
}

func (p *BistatParser) VarDeclaration() (localctx IVarDeclarationContext) {
	this := p
	_ = this

	localctx = NewVarDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BistatParserRULE_varDeclaration)

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
		p.SetState(108)
		p.Match(BistatParserT__2)
	}
	{
		p.SetState(109)
		p.Var_type()
	}
	{
		p.SetState(110)
		p.Match(BistatParserID)
	}
	{
		p.SetState(111)
		p.Match(BistatParserT__1)
	}

	return localctx
}

// IVar_typeContext is an interface to support dynamic dispatch.
type IVar_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TYPE_PRIMITIVE() antlr.TerminalNode
	CARDINALITY() antlr.TerminalNode

	// IsVar_typeContext differentiates from other interfaces.
	IsVar_typeContext()
}

type Var_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_typeContext() *Var_typeContext {
	var p = new(Var_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_var_type
	return p
}

func (*Var_typeContext) IsVar_typeContext() {}

func NewVar_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_typeContext {
	var p = new(Var_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_var_type

	return p
}

func (s *Var_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_typeContext) TYPE_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(BistatParserTYPE_PRIMITIVE, 0)
}

func (s *Var_typeContext) CARDINALITY() antlr.TerminalNode {
	return s.GetToken(BistatParserCARDINALITY, 0)
}

func (s *Var_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterVar_type(s)
	}
}

func (s *Var_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitVar_type(s)
	}
}

func (p *BistatParser) Var_type() (localctx IVar_typeContext) {
	this := p
	_ = this

	localctx = NewVar_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BistatParserRULE_var_type)
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
		p.SetState(113)
		p.Match(BistatParserTYPE_PRIMITIVE)
	}

	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BistatParserCARDINALITY {
		{
			p.SetState(114)
			p.Match(BistatParserCARDINALITY)
		}

	}

	return localctx
}

// IFuncDefContext is an interface to support dynamic dispatch.
type IFuncDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Var_type() IVar_typeContext
	AllParamDeclaration() []IParamDeclarationContext
	ParamDeclaration(i int) IParamDeclarationContext
	AllVarDeclaration() []IVarDeclarationContext
	VarDeclaration(i int) IVarDeclarationContext
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsFuncDefContext differentiates from other interfaces.
	IsFuncDefContext()
}

type FuncDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDefContext() *FuncDefContext {
	var p = new(FuncDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_funcDef
	return p
}

func (*FuncDefContext) IsFuncDefContext() {}

func NewFuncDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDefContext {
	var p = new(FuncDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_funcDef

	return p
}

func (s *FuncDefContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDefContext) ID() antlr.TerminalNode {
	return s.GetToken(BistatParserID, 0)
}

func (s *FuncDefContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *FuncDefContext) AllParamDeclaration() []IParamDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IParamDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamDeclarationContext); ok {
			tst[i] = t.(IParamDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *FuncDefContext) ParamDeclaration(i int) IParamDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamDeclarationContext); ok {
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

	return t.(IParamDeclarationContext)
}

func (s *FuncDefContext) AllVarDeclaration() []IVarDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IVarDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarDeclarationContext); ok {
			tst[i] = t.(IVarDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *FuncDefContext) VarDeclaration(i int) IVarDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclarationContext); ok {
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

	return t.(IVarDeclarationContext)
}

func (s *FuncDefContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *FuncDefContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
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

	return t.(IStmtContext)
}

func (s *FuncDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterFuncDef(s)
	}
}

func (s *FuncDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitFuncDef(s)
	}
}

func (p *BistatParser) FuncDef() (localctx IFuncDefContext) {
	this := p
	_ = this

	localctx = NewFuncDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BistatParserRULE_funcDef)
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
		p.SetState(117)
		p.Match(BistatParserT__3)
	}
	{
		p.SetState(118)
		p.Match(BistatParserID)
	}
	{
		p.SetState(119)
		p.Match(BistatParserT__4)
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__2 {
		{
			p.SetState(120)
			p.ParamDeclaration()
		}

		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(126)
		p.Match(BistatParserT__5)
	}
	{
		p.SetState(127)
		p.Match(BistatParserT__6)
	}
	{
		p.SetState(128)
		p.Var_type()
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__2 {
		{
			p.SetState(129)
			p.VarDeclaration()
		}

		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(135)
		p.Match(BistatParserT__7)
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&158329673187328) != 0) {
		{
			p.SetState(136)
			p.Stmt()
		}

		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(141)
		p.Match(BistatParserT__8)
	}

	return localctx
}

// IParamDeclarationContext is an interface to support dynamic dispatch.
type IParamDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var_type() IVar_typeContext
	ID() antlr.TerminalNode

	// IsParamDeclarationContext differentiates from other interfaces.
	IsParamDeclarationContext()
}

type ParamDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamDeclarationContext() *ParamDeclarationContext {
	var p = new(ParamDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_paramDeclaration
	return p
}

func (*ParamDeclarationContext) IsParamDeclarationContext() {}

func NewParamDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamDeclarationContext {
	var p = new(ParamDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_paramDeclaration

	return p
}

func (s *ParamDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamDeclarationContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *ParamDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(BistatParserID, 0)
}

func (s *ParamDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterParamDeclaration(s)
	}
}

func (s *ParamDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitParamDeclaration(s)
	}
}

func (p *BistatParser) ParamDeclaration() (localctx IParamDeclarationContext) {
	this := p
	_ = this

	localctx = NewParamDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BistatParserRULE_paramDeclaration)

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
		p.SetState(143)
		p.Match(BistatParserT__2)
	}
	{
		p.SetState(144)
		p.Var_type()
	}
	{
		p.SetState(145)
		p.Match(BistatParserID)
	}
	{
		p.SetState(146)
		p.Match(BistatParserT__1)
	}

	return localctx
}

// IMainContext is an interface to support dynamic dispatch.
type IMainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsMainContext differentiates from other interfaces.
	IsMainContext()
}

type MainContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMainContext() *MainContext {
	var p = new(MainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_main
	return p
}

func (*MainContext) IsMainContext() {}

func NewMainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainContext {
	var p = new(MainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_main

	return p
}

func (s *MainContext) GetParser() antlr.Parser { return s.parser }

func (s *MainContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *MainContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
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

	return t.(IStmtContext)
}

func (s *MainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterMain(s)
	}
}

func (s *MainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitMain(s)
	}
}

func (p *BistatParser) Main() (localctx IMainContext) {
	this := p
	_ = this

	localctx = NewMainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BistatParserRULE_main)
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
		p.SetState(148)
		p.Match(BistatParserT__9)
	}
	{
		p.SetState(149)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(150)
		p.Match(BistatParserT__5)
	}
	{
		p.SetState(151)
		p.Match(BistatParserT__7)
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&158329673187328) != 0) {
		{
			p.SetState(152)
			p.Stmt()
		}

		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(157)
		p.Match(BistatParserT__8)
	}

	return localctx
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assignment() IAssignmentContext
	SpecialFunction() ISpecialFunctionContext
	FunctionCall() IFunctionCallContext
	ReturnStmt() IReturnStmtContext
	Conditional() IConditionalContext
	WhileLoop() IWhileLoopContext
	ForLoop() IForLoopContext
	Comment() ICommentContext

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_stmt
	return p
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StmtContext) SpecialFunction() ISpecialFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpecialFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpecialFunctionContext)
}

func (s *StmtContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *StmtContext) ReturnStmt() IReturnStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStmtContext)
}

func (s *StmtContext) Conditional() IConditionalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalContext)
}

func (s *StmtContext) WhileLoop() IWhileLoopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileLoopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileLoopContext)
}

func (s *StmtContext) ForLoop() IForLoopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForLoopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForLoopContext)
}

func (s *StmtContext) Comment() ICommentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (p *BistatParser) Stmt() (localctx IStmtContext) {
	this := p
	_ = this

	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BistatParserRULE_stmt)

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

	p.SetState(171)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BistatParserT__20, BistatParserT__21, BistatParserT__22, BistatParserT__23, BistatParserT__24, BistatParserT__25, BistatParserT__26, BistatParserT__27, BistatParserT__28, BistatParserT__29, BistatParserT__30, BistatParserT__31, BistatParserT__32, BistatParserT__33, BistatParserT__34, BistatParserT__35, BistatParserT__36, BistatParserT__37, BistatParserT__38, BistatParserT__39, BistatParserT__40, BistatParserT__41, BistatParserT__42, BistatParserID:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(159)
				p.Assignment()
			}

		case 2:
			{
				p.SetState(160)
				p.SpecialFunction()
			}

		case 3:
			{
				p.SetState(161)
				p.FunctionCall()
			}

		case 4:
			{
				p.SetState(162)
				p.ReturnStmt()
			}

		}
		{
			p.SetState(165)
			p.Match(BistatParserT__1)
		}

	case BistatParserT__18:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(167)
			p.Conditional()
		}

	case BistatParserT__17:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(168)
			p.WhileLoop()
		}

	case BistatParserT__15:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(169)
			p.ForLoop()
		}

	case BistatParserT__14:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(170)
			p.Comment()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	VAR_CONS() antlr.TerminalNode
	ListAssignment() IListAssignmentContext
	MatrixAssignment() IMatrixAssignmentContext

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) ID() antlr.TerminalNode {
	return s.GetToken(BistatParserID, 0)
}

func (s *AssignmentContext) VAR_CONS() antlr.TerminalNode {
	return s.GetToken(BistatParserVAR_CONS, 0)
}

func (s *AssignmentContext) ListAssignment() IListAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListAssignmentContext)
}

func (s *AssignmentContext) MatrixAssignment() IMatrixAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrixAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrixAssignmentContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *BistatParser) Assignment() (localctx IAssignmentContext) {
	this := p
	_ = this

	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BistatParserRULE_assignment)

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
		p.SetState(173)
		p.Match(BistatParserID)
	}
	{
		p.SetState(174)
		p.Match(BistatParserT__10)
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(175)
			p.Match(BistatParserVAR_CONS)
		}

	case 2:
		{
			p.SetState(176)
			p.ListAssignment()
		}

	case 3:
		{
			p.SetState(177)
			p.MatrixAssignment()
		}

	}

	return localctx
}

// INested_stmtContext is an interface to support dynamic dispatch.
type INested_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assignment() IAssignmentContext
	Print_() IPrintContext
	SpecialFunction() ISpecialFunctionContext
	FunctionCall() IFunctionCallContext
	ReturnStmt() IReturnStmtContext
	Conditional() IConditionalContext
	WhileLoop() IWhileLoopContext
	Comment() ICommentContext

	// IsNested_stmtContext differentiates from other interfaces.
	IsNested_stmtContext()
}

type Nested_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNested_stmtContext() *Nested_stmtContext {
	var p = new(Nested_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_nested_stmt
	return p
}

func (*Nested_stmtContext) IsNested_stmtContext() {}

func NewNested_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Nested_stmtContext {
	var p = new(Nested_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_nested_stmt

	return p
}

func (s *Nested_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Nested_stmtContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *Nested_stmtContext) Print_() IPrintContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintContext)
}

func (s *Nested_stmtContext) SpecialFunction() ISpecialFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpecialFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpecialFunctionContext)
}

func (s *Nested_stmtContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *Nested_stmtContext) ReturnStmt() IReturnStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStmtContext)
}

func (s *Nested_stmtContext) Conditional() IConditionalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalContext)
}

func (s *Nested_stmtContext) WhileLoop() IWhileLoopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileLoopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileLoopContext)
}

func (s *Nested_stmtContext) Comment() ICommentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *Nested_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Nested_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Nested_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterNested_stmt(s)
	}
}

func (s *Nested_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitNested_stmt(s)
	}
}

func (p *BistatParser) Nested_stmt() (localctx INested_stmtContext) {
	this := p
	_ = this

	localctx = NewNested_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, BistatParserRULE_nested_stmt)

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

	p.SetState(192)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BistatParserT__20, BistatParserT__21, BistatParserT__22, BistatParserT__23, BistatParserT__24, BistatParserT__25, BistatParserT__26, BistatParserT__27, BistatParserT__28, BistatParserT__29, BistatParserT__30, BistatParserT__31, BistatParserT__32, BistatParserT__33, BistatParserT__34, BistatParserT__35, BistatParserT__36, BistatParserT__37, BistatParserT__38, BistatParserT__39, BistatParserT__40, BistatParserT__41, BistatParserT__42, BistatParserID:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(180)
				p.Assignment()
			}

		case 2:
			{
				p.SetState(181)
				p.Print_()
			}

		case 3:
			{
				p.SetState(182)
				p.SpecialFunction()
			}

		case 4:
			{
				p.SetState(183)
				p.FunctionCall()
			}

		case 5:
			{
				p.SetState(184)
				p.ReturnStmt()
			}

		}
		{
			p.SetState(187)
			p.Match(BistatParserT__1)
		}

	case BistatParserT__18:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(189)
			p.Conditional()
		}

	case BistatParserT__17:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(190)
			p.WhileLoop()
		}

	case BistatParserT__14:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(191)
			p.Comment()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMatrixAssignmentContext is an interface to support dynamic dispatch.
type IMatrixAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllListAssignment() []IListAssignmentContext
	ListAssignment(i int) IListAssignmentContext

	// IsMatrixAssignmentContext differentiates from other interfaces.
	IsMatrixAssignmentContext()
}

type MatrixAssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatrixAssignmentContext() *MatrixAssignmentContext {
	var p = new(MatrixAssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_matrixAssignment
	return p
}

func (*MatrixAssignmentContext) IsMatrixAssignmentContext() {}

func NewMatrixAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatrixAssignmentContext {
	var p = new(MatrixAssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_matrixAssignment

	return p
}

func (s *MatrixAssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *MatrixAssignmentContext) AllListAssignment() []IListAssignmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IListAssignmentContext); ok {
			len++
		}
	}

	tst := make([]IListAssignmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IListAssignmentContext); ok {
			tst[i] = t.(IListAssignmentContext)
			i++
		}
	}

	return tst
}

func (s *MatrixAssignmentContext) ListAssignment(i int) IListAssignmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListAssignmentContext); ok {
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

	return t.(IListAssignmentContext)
}

func (s *MatrixAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatrixAssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatrixAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterMatrixAssignment(s)
	}
}

func (s *MatrixAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitMatrixAssignment(s)
	}
}

func (p *BistatParser) MatrixAssignment() (localctx IMatrixAssignmentContext) {
	this := p
	_ = this

	localctx = NewMatrixAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, BistatParserRULE_matrixAssignment)
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
		p.SetState(194)
		p.Match(BistatParserT__11)
	}
	{
		p.SetState(195)
		p.ListAssignment()
	}
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__12 {
		{
			p.SetState(196)
			p.Match(BistatParserT__12)
		}
		{
			p.SetState(197)
			p.ListAssignment()
		}

		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(203)
		p.Match(BistatParserT__13)
	}

	return localctx
}

// IListAssignmentContext is an interface to support dynamic dispatch.
type IListAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVAR_CONS() []antlr.TerminalNode
	VAR_CONS(i int) antlr.TerminalNode

	// IsListAssignmentContext differentiates from other interfaces.
	IsListAssignmentContext()
}

type ListAssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListAssignmentContext() *ListAssignmentContext {
	var p = new(ListAssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_listAssignment
	return p
}

func (*ListAssignmentContext) IsListAssignmentContext() {}

func NewListAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListAssignmentContext {
	var p = new(ListAssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_listAssignment

	return p
}

func (s *ListAssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *ListAssignmentContext) AllVAR_CONS() []antlr.TerminalNode {
	return s.GetTokens(BistatParserVAR_CONS)
}

func (s *ListAssignmentContext) VAR_CONS(i int) antlr.TerminalNode {
	return s.GetToken(BistatParserVAR_CONS, i)
}

func (s *ListAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListAssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterListAssignment(s)
	}
}

func (s *ListAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitListAssignment(s)
	}
}

func (p *BistatParser) ListAssignment() (localctx IListAssignmentContext) {
	this := p
	_ = this

	localctx = NewListAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, BistatParserRULE_listAssignment)
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
		p.Match(BistatParserT__11)
	}
	{
		p.SetState(206)
		p.Match(BistatParserVAR_CONS)
	}
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__12 {
		{
			p.SetState(207)
			p.Match(BistatParserT__12)
		}
		{
			p.SetState(208)
			p.Match(BistatParserVAR_CONS)
		}

		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(214)
		p.Match(BistatParserT__13)
	}

	return localctx
}

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_comment
	return p
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }
func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitComment(s)
	}
}

func (p *BistatParser) Comment() (localctx ICommentContext) {
	this := p
	_ = this

	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, BistatParserRULE_comment)
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
		p.SetState(216)
		p.Match(BistatParserT__14)
	}
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&144115188075823102) != 0) {
		{
			p.SetState(217)
			_la = p.GetTokenStream().LA(1)

			if _la <= 0 || _la == BistatParserT__14 {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(220)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(222)
		p.Match(BistatParserT__14)
	}

	return localctx
}

// IForLoopContext is an interface to support dynamic dispatch.
type IForLoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Expression() IExpressionContext
	AllNested_stmt() []INested_stmtContext
	Nested_stmt(i int) INested_stmtContext

	// IsForLoopContext differentiates from other interfaces.
	IsForLoopContext()
}

type ForLoopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForLoopContext() *ForLoopContext {
	var p = new(ForLoopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_forLoop
	return p
}

func (*ForLoopContext) IsForLoopContext() {}

func NewForLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForLoopContext {
	var p = new(ForLoopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_forLoop

	return p
}

func (s *ForLoopContext) GetParser() antlr.Parser { return s.parser }

func (s *ForLoopContext) ID() antlr.TerminalNode {
	return s.GetToken(BistatParserID, 0)
}

func (s *ForLoopContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ForLoopContext) AllNested_stmt() []INested_stmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INested_stmtContext); ok {
			len++
		}
	}

	tst := make([]INested_stmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INested_stmtContext); ok {
			tst[i] = t.(INested_stmtContext)
			i++
		}
	}

	return tst
}

func (s *ForLoopContext) Nested_stmt(i int) INested_stmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INested_stmtContext); ok {
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

	return t.(INested_stmtContext)
}

func (s *ForLoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForLoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForLoopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterForLoop(s)
	}
}

func (s *ForLoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitForLoop(s)
	}
}

func (p *BistatParser) ForLoop() (localctx IForLoopContext) {
	this := p
	_ = this

	localctx = NewForLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, BistatParserRULE_forLoop)
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
		p.SetState(224)
		p.Match(BistatParserT__15)
	}
	{
		p.SetState(225)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(226)
		p.Match(BistatParserID)
	}
	{
		p.SetState(227)
		p.Match(BistatParserT__16)
	}
	{
		p.SetState(228)
		p.Expression()
	}
	{
		p.SetState(229)
		p.Match(BistatParserT__5)
	}
	{
		p.SetState(230)
		p.Match(BistatParserT__7)
	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&158329673121792) != 0) {
		{
			p.SetState(231)
			p.Nested_stmt()
		}

		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(236)
		p.Match(BistatParserT__8)
	}

	return localctx
}

// IWhileLoopContext is an interface to support dynamic dispatch.
type IWhileLoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsWhileLoopContext differentiates from other interfaces.
	IsWhileLoopContext()
}

type WhileLoopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileLoopContext() *WhileLoopContext {
	var p = new(WhileLoopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_whileLoop
	return p
}

func (*WhileLoopContext) IsWhileLoopContext() {}

func NewWhileLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileLoopContext {
	var p = new(WhileLoopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_whileLoop

	return p
}

func (s *WhileLoopContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileLoopContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhileLoopContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *WhileLoopContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
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

	return t.(IStmtContext)
}

func (s *WhileLoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileLoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileLoopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterWhileLoop(s)
	}
}

func (s *WhileLoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitWhileLoop(s)
	}
}

func (p *BistatParser) WhileLoop() (localctx IWhileLoopContext) {
	this := p
	_ = this

	localctx = NewWhileLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, BistatParserRULE_whileLoop)
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
		p.SetState(238)
		p.Match(BistatParserT__17)
	}
	{
		p.SetState(239)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(240)
		p.Expression()
	}
	{
		p.SetState(241)
		p.Match(BistatParserT__5)
	}
	{
		p.SetState(242)
		p.Match(BistatParserT__7)
	}
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&158329673187328) != 0) {
		{
			p.SetState(243)
			p.Stmt()
		}

		p.SetState(246)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(248)
		p.Match(BistatParserT__8)
	}

	return localctx
}

// IConditionalContext is an interface to support dynamic dispatch.
type IConditionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsConditionalContext differentiates from other interfaces.
	IsConditionalContext()
}

type ConditionalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalContext() *ConditionalContext {
	var p = new(ConditionalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_conditional
	return p
}

func (*ConditionalContext) IsConditionalContext() {}

func NewConditionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalContext {
	var p = new(ConditionalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_conditional

	return p
}

func (s *ConditionalContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ConditionalContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ConditionalContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ConditionalContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
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

	return t.(IStmtContext)
}

func (s *ConditionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterConditional(s)
	}
}

func (s *ConditionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitConditional(s)
	}
}

func (p *BistatParser) Conditional() (localctx IConditionalContext) {
	this := p
	_ = this

	localctx = NewConditionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, BistatParserRULE_conditional)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(250)
		p.Match(BistatParserT__18)
	}
	{
		p.SetState(251)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(252)
		p.Expression()
	}
	{
		p.SetState(253)
		p.Match(BistatParserT__5)
	}
	{
		p.SetState(254)
		p.Match(BistatParserT__7)
	}
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&158329673187328) != 0) {
		{
			p.SetState(255)
			p.Stmt()
		}

		p.SetState(258)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(260)
		p.Match(BistatParserT__8)
	}
	p.SetState(276)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(261)
				p.Match(BistatParserT__19)
			}
			{
				p.SetState(262)
				p.Match(BistatParserT__18)
			}
			{
				p.SetState(263)
				p.Match(BistatParserT__4)
			}
			{
				p.SetState(264)
				p.Expression()
			}
			{
				p.SetState(265)
				p.Match(BistatParserT__5)
			}
			{
				p.SetState(266)
				p.Match(BistatParserT__7)
			}
			p.SetState(268)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&158329673187328) != 0) {
				{
					p.SetState(267)
					p.Stmt()
				}

				p.SetState(270)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(272)
				p.Match(BistatParserT__8)
			}

		}
		p.SetState(278)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
	}
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BistatParserT__19 {
		{
			p.SetState(279)
			p.Match(BistatParserT__19)
		}
		{
			p.SetState(280)
			p.Match(BistatParserT__7)
		}
		p.SetState(282)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&158329673187328) != 0) {
			{
				p.SetState(281)
				p.Stmt()
			}

			p.SetState(284)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(286)
			p.Match(BistatParserT__8)
		}

	}

	return localctx
}

// IReturnStmtContext is an interface to support dynamic dispatch.
type IReturnStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsReturnStmtContext differentiates from other interfaces.
	IsReturnStmtContext()
}

type ReturnStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStmtContext() *ReturnStmtContext {
	var p = new(ReturnStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_returnStmt
	return p
}

func (*ReturnStmtContext) IsReturnStmtContext() {}

func NewReturnStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_returnStmt

	return p
}

func (s *ReturnStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStmtContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterReturnStmt(s)
	}
}

func (s *ReturnStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitReturnStmt(s)
	}
}

func (p *BistatParser) ReturnStmt() (localctx IReturnStmtContext) {
	this := p
	_ = this

	localctx = NewReturnStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, BistatParserRULE_returnStmt)

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
		p.SetState(290)
		p.Match(BistatParserT__20)
	}
	{
		p.SetState(291)
		p.Expression()
	}

	return localctx
}

// ISpecialFunctionContext is an interface to support dynamic dispatch.
type ISpecialFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	InputRead() IInputReadContext
	Print_() IPrintContext
	ListAdd() IListAddContext
	ListPop() IListPopContext
	Length() ILengthContext
	Range_() IRangeContext
	Plot() IPlotContext
	Sum() ISumContext
	Min() IMinContext
	Prod() IProdContext
	Avg() IAvgContext
	SMode() ISModeContext
	Median() IMedianContext
	Sin() ISinContext
	Cos() ICosContext
	Tan() ITanContext
	Sort() ISortContext
	Sqrt() ISqrtContext
	Floor() IFloorContext
	Ceil() ICeilContext
	Abs() IAbsContext
	Not() INotContext

	// IsSpecialFunctionContext differentiates from other interfaces.
	IsSpecialFunctionContext()
}

type SpecialFunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpecialFunctionContext() *SpecialFunctionContext {
	var p = new(SpecialFunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_specialFunction
	return p
}

func (*SpecialFunctionContext) IsSpecialFunctionContext() {}

func NewSpecialFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecialFunctionContext {
	var p = new(SpecialFunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_specialFunction

	return p
}

func (s *SpecialFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *SpecialFunctionContext) InputRead() IInputReadContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInputReadContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInputReadContext)
}

func (s *SpecialFunctionContext) Print_() IPrintContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintContext)
}

func (s *SpecialFunctionContext) ListAdd() IListAddContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListAddContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListAddContext)
}

func (s *SpecialFunctionContext) ListPop() IListPopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListPopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListPopContext)
}

func (s *SpecialFunctionContext) Length() ILengthContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILengthContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILengthContext)
}

func (s *SpecialFunctionContext) Range_() IRangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRangeContext)
}

func (s *SpecialFunctionContext) Plot() IPlotContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPlotContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPlotContext)
}

func (s *SpecialFunctionContext) Sum() ISumContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISumContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISumContext)
}

func (s *SpecialFunctionContext) Min() IMinContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMinContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMinContext)
}

func (s *SpecialFunctionContext) Prod() IProdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProdContext)
}

func (s *SpecialFunctionContext) Avg() IAvgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAvgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAvgContext)
}

func (s *SpecialFunctionContext) SMode() ISModeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISModeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISModeContext)
}

func (s *SpecialFunctionContext) Median() IMedianContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMedianContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMedianContext)
}

func (s *SpecialFunctionContext) Sin() ISinContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISinContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISinContext)
}

func (s *SpecialFunctionContext) Cos() ICosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICosContext)
}

func (s *SpecialFunctionContext) Tan() ITanContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITanContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITanContext)
}

func (s *SpecialFunctionContext) Sort() ISortContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISortContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISortContext)
}

func (s *SpecialFunctionContext) Sqrt() ISqrtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISqrtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISqrtContext)
}

func (s *SpecialFunctionContext) Floor() IFloorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloorContext)
}

func (s *SpecialFunctionContext) Ceil() ICeilContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICeilContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICeilContext)
}

func (s *SpecialFunctionContext) Abs() IAbsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAbsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAbsContext)
}

func (s *SpecialFunctionContext) Not() INotContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INotContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INotContext)
}

func (s *SpecialFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecialFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpecialFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterSpecialFunction(s)
	}
}

func (s *SpecialFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitSpecialFunction(s)
	}
}

func (p *BistatParser) SpecialFunction() (localctx ISpecialFunctionContext) {
	this := p
	_ = this

	localctx = NewSpecialFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, BistatParserRULE_specialFunction)

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

	p.SetState(315)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BistatParserT__21:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(293)
			p.InputRead()
		}

	case BistatParserT__22:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(294)
			p.Print_()
		}

	case BistatParserT__23:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(295)
			p.ListAdd()
		}

	case BistatParserT__24:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(296)
			p.ListPop()
		}

	case BistatParserT__25:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(297)
			p.Length()
		}

	case BistatParserT__26:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(298)
			p.Range_()
		}

	case BistatParserT__27:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(299)
			p.Plot()
		}

	case BistatParserT__28:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(300)
			p.Sum()
		}

	case BistatParserT__29:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(301)
			p.Min()
		}

	case BistatParserT__30:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(302)
			p.Prod()
		}

	case BistatParserT__31:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(303)
			p.Avg()
		}

	case BistatParserT__32:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(304)
			p.SMode()
		}

	case BistatParserT__33:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(305)
			p.Median()
		}

	case BistatParserT__34:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(306)
			p.Sin()
		}

	case BistatParserT__36:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(307)
			p.Cos()
		}

	case BistatParserT__35:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(308)
			p.Tan()
		}

	case BistatParserT__37:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(309)
			p.Sort()
		}

	case BistatParserT__38:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(310)
			p.Sqrt()
		}

	case BistatParserT__39:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(311)
			p.Floor()
		}

	case BistatParserT__40:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(312)
			p.Ceil()
		}

	case BistatParserT__41:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(313)
			p.Abs()
		}

	case BistatParserT__42:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(314)
			p.Not()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IInputReadContext is an interface to support dynamic dispatch.
type IInputReadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode

	// IsInputReadContext differentiates from other interfaces.
	IsInputReadContext()
}

type InputReadContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputReadContext() *InputReadContext {
	var p = new(InputReadContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_inputRead
	return p
}

func (*InputReadContext) IsInputReadContext() {}

func NewInputReadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputReadContext {
	var p = new(InputReadContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_inputRead

	return p
}

func (s *InputReadContext) GetParser() antlr.Parser { return s.parser }

func (s *InputReadContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(BistatParserID)
}

func (s *InputReadContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(BistatParserID, i)
}

func (s *InputReadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputReadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputReadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterInputRead(s)
	}
}

func (s *InputReadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitInputRead(s)
	}
}

func (p *BistatParser) InputRead() (localctx IInputReadContext) {
	this := p
	_ = this

	localctx = NewInputReadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, BistatParserRULE_inputRead)
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
		p.SetState(317)
		p.Match(BistatParserT__21)
	}
	{
		p.SetState(318)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(319)
		p.Match(BistatParserID)
	}
	p.SetState(324)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__12 {
		{
			p.SetState(320)
			p.Match(BistatParserT__12)
		}
		{
			p.SetState(321)
			p.Match(BistatParserID)
		}

		p.SetState(326)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(327)
		p.Match(BistatParserT__5)
	}

	return localctx
}

// IPrintContext is an interface to support dynamic dispatch.
type IPrintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsPrintContext differentiates from other interfaces.
	IsPrintContext()
}

type PrintContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintContext() *PrintContext {
	var p = new(PrintContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_print
	return p
}

func (*PrintContext) IsPrintContext() {}

func NewPrintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintContext {
	var p = new(PrintContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_print

	return p
}

func (s *PrintContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *PrintContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *PrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterPrint(s)
	}
}

func (s *PrintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitPrint(s)
	}
}

func (p *BistatParser) Print_() (localctx IPrintContext) {
	this := p
	_ = this

	localctx = NewPrintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, BistatParserRULE_print)
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
		p.SetState(329)
		p.Match(BistatParserT__22)
	}
	{
		p.SetState(330)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(331)
		p.Expression()
	}
	p.SetState(336)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__12 {
		{
			p.SetState(332)
			p.Match(BistatParserT__12)
		}
		{
			p.SetState(333)
			p.Expression()
		}

		p.SetState(338)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(339)
		p.Match(BistatParserT__5)
	}

	return localctx
}

// IListAddContext is an interface to support dynamic dispatch.
type IListAddContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsListAddContext differentiates from other interfaces.
	IsListAddContext()
}

type ListAddContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListAddContext() *ListAddContext {
	var p = new(ListAddContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_listAdd
	return p
}

func (*ListAddContext) IsListAddContext() {}

func NewListAddContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListAddContext {
	var p = new(ListAddContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_listAdd

	return p
}

func (s *ListAddContext) GetParser() antlr.Parser { return s.parser }

func (s *ListAddContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ListAddContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ListAddContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListAddContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListAddContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterListAdd(s)
	}
}

func (s *ListAddContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitListAdd(s)
	}
}

func (p *BistatParser) ListAdd() (localctx IListAddContext) {
	this := p
	_ = this

	localctx = NewListAddContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, BistatParserRULE_listAdd)

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
		p.SetState(341)
		p.Match(BistatParserT__23)
	}
	{
		p.SetState(342)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(343)
		p.Expression()
	}
	{
		p.SetState(344)
		p.Match(BistatParserT__12)
	}
	{
		p.SetState(345)
		p.Expression()
	}
	{
		p.SetState(346)
		p.Match(BistatParserT__5)
	}

	return localctx
}

// IListPopContext is an interface to support dynamic dispatch.
type IListPopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsListPopContext differentiates from other interfaces.
	IsListPopContext()
}

type ListPopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListPopContext() *ListPopContext {
	var p = new(ListPopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_listPop
	return p
}

func (*ListPopContext) IsListPopContext() {}

func NewListPopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListPopContext {
	var p = new(ListPopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_listPop

	return p
}

func (s *ListPopContext) GetParser() antlr.Parser { return s.parser }

func (s *ListPopContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ListPopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListPopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListPopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterListPop(s)
	}
}

func (s *ListPopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitListPop(s)
	}
}

func (p *BistatParser) ListPop() (localctx IListPopContext) {
	this := p
	_ = this

	localctx = NewListPopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, BistatParserRULE_listPop)

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
		p.SetState(348)
		p.Match(BistatParserT__24)
	}
	{
		p.SetState(349)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(350)
		p.Expression()
	}
	{
		p.SetState(351)
		p.Match(BistatParserT__5)
	}

	return localctx
}

// ILengthContext is an interface to support dynamic dispatch.
type ILengthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsLengthContext differentiates from other interfaces.
	IsLengthContext()
}

type LengthContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLengthContext() *LengthContext {
	var p = new(LengthContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_length
	return p
}

func (*LengthContext) IsLengthContext() {}

func NewLengthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LengthContext {
	var p = new(LengthContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_length

	return p
}

func (s *LengthContext) GetParser() antlr.Parser { return s.parser }

func (s *LengthContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LengthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LengthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterLength(s)
	}
}

func (s *LengthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitLength(s)
	}
}

func (p *BistatParser) Length() (localctx ILengthContext) {
	this := p
	_ = this

	localctx = NewLengthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, BistatParserRULE_length)

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
		p.SetState(353)
		p.Match(BistatParserT__25)
	}
	{
		p.SetState(354)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(355)
		p.Expression()
	}
	{
		p.SetState(356)
		p.Match(BistatParserT__5)
	}

	return localctx
}

// IRangeContext is an interface to support dynamic dispatch.
type IRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsRangeContext differentiates from other interfaces.
	IsRangeContext()
}

type RangeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeContext() *RangeContext {
	var p = new(RangeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_range
	return p
}

func (*RangeContext) IsRangeContext() {}

func NewRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeContext {
	var p = new(RangeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_range

	return p
}

func (s *RangeContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *RangeContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *RangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterRange(s)
	}
}

func (s *RangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitRange(s)
	}
}

func (p *BistatParser) Range_() (localctx IRangeContext) {
	this := p
	_ = this

	localctx = NewRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, BistatParserRULE_range)
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
		p.Match(BistatParserT__26)
	}
	{
		p.SetState(359)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(360)
		p.Expression()
	}
	p.SetState(363)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BistatParserT__12 {
		{
			p.SetState(361)
			p.Match(BistatParserT__12)
		}
		{
			p.SetState(362)
			p.Expression()
		}

	}
	{
		p.SetState(365)
		p.Match(BistatParserT__5)
	}

	return localctx
}

// IPlotContext is an interface to support dynamic dispatch.
type IPlotContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsPlotContext differentiates from other interfaces.
	IsPlotContext()
}

type PlotContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPlotContext() *PlotContext {
	var p = new(PlotContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_plot
	return p
}

func (*PlotContext) IsPlotContext() {}

func NewPlotContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PlotContext {
	var p = new(PlotContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_plot

	return p
}

func (s *PlotContext) GetParser() antlr.Parser { return s.parser }

func (s *PlotContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PlotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlotContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PlotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterPlot(s)
	}
}

func (s *PlotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitPlot(s)
	}
}

func (p *BistatParser) Plot() (localctx IPlotContext) {
	this := p
	_ = this

	localctx = NewPlotContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, BistatParserRULE_plot)

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
		p.SetState(367)
		p.Match(BistatParserT__27)
	}
	{
		p.SetState(368)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(369)
		p.Expression()
	}
	{
		p.SetState(370)
		p.Match(BistatParserT__5)
	}

	return localctx
}

// ISumContext is an interface to support dynamic dispatch.
type ISumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsSumContext differentiates from other interfaces.
	IsSumContext()
}

type SumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySumContext() *SumContext {
	var p = new(SumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_sum
	return p
}

func (*SumContext) IsSumContext() {}

func NewSumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SumContext {
	var p = new(SumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_sum

	return p
}

func (s *SumContext) GetParser() antlr.Parser { return s.parser }

func (s *SumContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterSum(s)
	}
}

func (s *SumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitSum(s)
	}
}

func (p *BistatParser) Sum() (localctx ISumContext) {
	this := p
	_ = this

	localctx = NewSumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, BistatParserRULE_sum)

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
		p.SetState(372)
		p.Match(BistatParserT__28)
	}
	{
		p.SetState(373)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(374)
		p.Expression()
	}
	{
		p.SetState(375)
		p.Match(BistatParserT__5)
	}

	return localctx
}

// IMinContext is an interface to support dynamic dispatch.
type IMinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsMinContext differentiates from other interfaces.
	IsMinContext()
}

type MinContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMinContext() *MinContext {
	var p = new(MinContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_min
	return p
}

func (*MinContext) IsMinContext() {}

func NewMinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MinContext {
	var p = new(MinContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_min

	return p
}

func (s *MinContext) GetParser() antlr.Parser { return s.parser }

func (s *MinContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterMin(s)
	}
}

func (s *MinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitMin(s)
	}
}

func (p *BistatParser) Min() (localctx IMinContext) {
	this := p
	_ = this

	localctx = NewMinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, BistatParserRULE_min)

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
		p.SetState(377)
		p.Match(BistatParserT__29)
	}
	{
		p.SetState(378)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(379)
		p.Expression()
	}
	{
		p.SetState(380)
		p.Match(BistatParserT__5)
	}

	return localctx
}

// IProdContext is an interface to support dynamic dispatch.
type IProdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsProdContext differentiates from other interfaces.
	IsProdContext()
}

type ProdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProdContext() *ProdContext {
	var p = new(ProdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_prod
	return p
}

func (*ProdContext) IsProdContext() {}

func NewProdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProdContext {
	var p = new(ProdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_prod

	return p
}

func (s *ProdContext) GetParser() antlr.Parser { return s.parser }

func (s *ProdContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ProdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterProd(s)
	}
}

func (s *ProdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitProd(s)
	}
}

func (p *BistatParser) Prod() (localctx IProdContext) {
	this := p
	_ = this

	localctx = NewProdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, BistatParserRULE_prod)

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
		p.SetState(382)
		p.Match(BistatParserT__30)
	}
	{
		p.SetState(383)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(384)
		p.Expression()
	}
	{
		p.SetState(385)
		p.Match(BistatParserT__5)
	}

	return localctx
}

// IAvgContext is an interface to support dynamic dispatch.
type IAvgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsAvgContext differentiates from other interfaces.
	IsAvgContext()
}

type AvgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAvgContext() *AvgContext {
	var p = new(AvgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_avg
	return p
}

func (*AvgContext) IsAvgContext() {}

func NewAvgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AvgContext {
	var p = new(AvgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_avg

	return p
}

func (s *AvgContext) GetParser() antlr.Parser { return s.parser }

func (s *AvgContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AvgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AvgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AvgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterAvg(s)
	}
}

func (s *AvgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitAvg(s)
	}
}

func (p *BistatParser) Avg() (localctx IAvgContext) {
	this := p
	_ = this

	localctx = NewAvgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, BistatParserRULE_avg)

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
		p.SetState(387)
		p.Match(BistatParserT__31)
	}
	{
		p.SetState(388)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(389)
		p.Expression()
	}
	{
		p.SetState(390)
		p.Match(BistatParserT__5)
	}

	return localctx
}

// ISModeContext is an interface to support dynamic dispatch.
type ISModeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsSModeContext differentiates from other interfaces.
	IsSModeContext()
}

type SModeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySModeContext() *SModeContext {
	var p = new(SModeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_sMode
	return p
}

func (*SModeContext) IsSModeContext() {}

func NewSModeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SModeContext {
	var p = new(SModeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_sMode

	return p
}

func (s *SModeContext) GetParser() antlr.Parser { return s.parser }

func (s *SModeContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SModeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SModeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SModeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterSMode(s)
	}
}

func (s *SModeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitSMode(s)
	}
}

func (p *BistatParser) SMode() (localctx ISModeContext) {
	this := p
	_ = this

	localctx = NewSModeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, BistatParserRULE_sMode)

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
		p.SetState(392)
		p.Match(BistatParserT__32)
	}
	{
		p.SetState(393)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(394)
		p.Expression()
	}
	{
		p.SetState(395)
		p.Match(BistatParserT__5)
	}

	return localctx
}

// IMedianContext is an interface to support dynamic dispatch.
type IMedianContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsMedianContext differentiates from other interfaces.
	IsMedianContext()
}

type MedianContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMedianContext() *MedianContext {
	var p = new(MedianContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_median
	return p
}

func (*MedianContext) IsMedianContext() {}

func NewMedianContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MedianContext {
	var p = new(MedianContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_median

	return p
}

func (s *MedianContext) GetParser() antlr.Parser { return s.parser }

func (s *MedianContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MedianContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MedianContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MedianContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterMedian(s)
	}
}

func (s *MedianContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitMedian(s)
	}
}

func (p *BistatParser) Median() (localctx IMedianContext) {
	this := p
	_ = this

	localctx = NewMedianContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, BistatParserRULE_median)

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
		p.SetState(397)
		p.Match(BistatParserT__33)
	}
	{
		p.SetState(398)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(399)
		p.Expression()
	}
	{
		p.SetState(400)
		p.Match(BistatParserT__5)
	}

	return localctx
}

// ISinContext is an interface to support dynamic dispatch.
type ISinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsSinContext differentiates from other interfaces.
	IsSinContext()
}

type SinContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySinContext() *SinContext {
	var p = new(SinContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_sin
	return p
}

func (*SinContext) IsSinContext() {}

func NewSinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SinContext {
	var p = new(SinContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_sin

	return p
}

func (s *SinContext) GetParser() antlr.Parser { return s.parser }

func (s *SinContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterSin(s)
	}
}

func (s *SinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitSin(s)
	}
}

func (p *BistatParser) Sin() (localctx ISinContext) {
	this := p
	_ = this

	localctx = NewSinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, BistatParserRULE_sin)

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
		p.SetState(402)
		p.Match(BistatParserT__34)
	}
	{
		p.SetState(403)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(404)
		p.Expression()
	}
	{
		p.SetState(405)
		p.Match(BistatParserT__5)
	}

	return localctx
}

// ITanContext is an interface to support dynamic dispatch.
type ITanContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsTanContext differentiates from other interfaces.
	IsTanContext()
}

type TanContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTanContext() *TanContext {
	var p = new(TanContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_tan
	return p
}

func (*TanContext) IsTanContext() {}

func NewTanContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TanContext {
	var p = new(TanContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_tan

	return p
}

func (s *TanContext) GetParser() antlr.Parser { return s.parser }

func (s *TanContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TanContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterTan(s)
	}
}

func (s *TanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitTan(s)
	}
}

func (p *BistatParser) Tan() (localctx ITanContext) {
	this := p
	_ = this

	localctx = NewTanContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, BistatParserRULE_tan)

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
		p.SetState(407)
		p.Match(BistatParserT__35)
	}
	{
		p.SetState(408)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(409)
		p.Expression()
	}
	{
		p.SetState(410)
		p.Match(BistatParserT__5)
	}

	return localctx
}

// ICosContext is an interface to support dynamic dispatch.
type ICosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsCosContext differentiates from other interfaces.
	IsCosContext()
}

type CosContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCosContext() *CosContext {
	var p = new(CosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_cos
	return p
}

func (*CosContext) IsCosContext() {}

func NewCosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CosContext {
	var p = new(CosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_cos

	return p
}

func (s *CosContext) GetParser() antlr.Parser { return s.parser }

func (s *CosContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterCos(s)
	}
}

func (s *CosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitCos(s)
	}
}

func (p *BistatParser) Cos() (localctx ICosContext) {
	this := p
	_ = this

	localctx = NewCosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, BistatParserRULE_cos)

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
		p.SetState(412)
		p.Match(BistatParserT__36)
	}
	{
		p.SetState(413)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(414)
		p.Expression()
	}
	{
		p.SetState(415)
		p.Match(BistatParserT__5)
	}

	return localctx
}

// ISortContext is an interface to support dynamic dispatch.
type ISortContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsSortContext differentiates from other interfaces.
	IsSortContext()
}

type SortContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySortContext() *SortContext {
	var p = new(SortContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_sort
	return p
}

func (*SortContext) IsSortContext() {}

func NewSortContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SortContext {
	var p = new(SortContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_sort

	return p
}

func (s *SortContext) GetParser() antlr.Parser { return s.parser }

func (s *SortContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SortContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SortContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SortContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterSort(s)
	}
}

func (s *SortContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitSort(s)
	}
}

func (p *BistatParser) Sort() (localctx ISortContext) {
	this := p
	_ = this

	localctx = NewSortContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, BistatParserRULE_sort)

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
		p.SetState(417)
		p.Match(BistatParserT__37)
	}
	{
		p.SetState(418)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(419)
		p.Expression()
	}
	{
		p.SetState(420)
		p.Match(BistatParserT__5)
	}

	return localctx
}

// ISqrtContext is an interface to support dynamic dispatch.
type ISqrtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsSqrtContext differentiates from other interfaces.
	IsSqrtContext()
}

type SqrtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySqrtContext() *SqrtContext {
	var p = new(SqrtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_sqrt
	return p
}

func (*SqrtContext) IsSqrtContext() {}

func NewSqrtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SqrtContext {
	var p = new(SqrtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_sqrt

	return p
}

func (s *SqrtContext) GetParser() antlr.Parser { return s.parser }

func (s *SqrtContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SqrtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SqrtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SqrtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterSqrt(s)
	}
}

func (s *SqrtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitSqrt(s)
	}
}

func (p *BistatParser) Sqrt() (localctx ISqrtContext) {
	this := p
	_ = this

	localctx = NewSqrtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, BistatParserRULE_sqrt)

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
		p.SetState(422)
		p.Match(BistatParserT__38)
	}
	{
		p.SetState(423)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(424)
		p.Expression()
	}
	{
		p.SetState(425)
		p.Match(BistatParserT__5)
	}

	return localctx
}

// IFloorContext is an interface to support dynamic dispatch.
type IFloorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsFloorContext differentiates from other interfaces.
	IsFloorContext()
}

type FloorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloorContext() *FloorContext {
	var p = new(FloorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_floor
	return p
}

func (*FloorContext) IsFloorContext() {}

func NewFloorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloorContext {
	var p = new(FloorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_floor

	return p
}

func (s *FloorContext) GetParser() antlr.Parser { return s.parser }

func (s *FloorContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FloorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterFloor(s)
	}
}

func (s *FloorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitFloor(s)
	}
}

func (p *BistatParser) Floor() (localctx IFloorContext) {
	this := p
	_ = this

	localctx = NewFloorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, BistatParserRULE_floor)

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
		p.SetState(427)
		p.Match(BistatParserT__39)
	}
	{
		p.SetState(428)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(429)
		p.Expression()
	}
	{
		p.SetState(430)
		p.Match(BistatParserT__5)
	}

	return localctx
}

// ICeilContext is an interface to support dynamic dispatch.
type ICeilContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsCeilContext differentiates from other interfaces.
	IsCeilContext()
}

type CeilContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCeilContext() *CeilContext {
	var p = new(CeilContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_ceil
	return p
}

func (*CeilContext) IsCeilContext() {}

func NewCeilContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CeilContext {
	var p = new(CeilContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_ceil

	return p
}

func (s *CeilContext) GetParser() antlr.Parser { return s.parser }

func (s *CeilContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CeilContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CeilContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CeilContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterCeil(s)
	}
}

func (s *CeilContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitCeil(s)
	}
}

func (p *BistatParser) Ceil() (localctx ICeilContext) {
	this := p
	_ = this

	localctx = NewCeilContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, BistatParserRULE_ceil)

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
		p.SetState(432)
		p.Match(BistatParserT__40)
	}
	{
		p.SetState(433)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(434)
		p.Expression()
	}
	{
		p.SetState(435)
		p.Match(BistatParserT__5)
	}

	return localctx
}

// IAbsContext is an interface to support dynamic dispatch.
type IAbsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsAbsContext differentiates from other interfaces.
	IsAbsContext()
}

type AbsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAbsContext() *AbsContext {
	var p = new(AbsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_abs
	return p
}

func (*AbsContext) IsAbsContext() {}

func NewAbsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AbsContext {
	var p = new(AbsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_abs

	return p
}

func (s *AbsContext) GetParser() antlr.Parser { return s.parser }

func (s *AbsContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AbsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AbsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AbsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterAbs(s)
	}
}

func (s *AbsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitAbs(s)
	}
}

func (p *BistatParser) Abs() (localctx IAbsContext) {
	this := p
	_ = this

	localctx = NewAbsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, BistatParserRULE_abs)

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
		p.SetState(437)
		p.Match(BistatParserT__41)
	}
	{
		p.SetState(438)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(439)
		p.Expression()
	}
	{
		p.SetState(440)
		p.Match(BistatParserT__5)
	}

	return localctx
}

// INotContext is an interface to support dynamic dispatch.
type INotContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsNotContext differentiates from other interfaces.
	IsNotContext()
}

type NotContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotContext() *NotContext {
	var p = new(NotContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_not
	return p
}

func (*NotContext) IsNotContext() {}

func NewNotContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotContext {
	var p = new(NotContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_not

	return p
}

func (s *NotContext) GetParser() antlr.Parser { return s.parser }

func (s *NotContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterNot(s)
	}
}

func (s *NotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitNot(s)
	}
}

func (p *BistatParser) Not() (localctx INotContext) {
	this := p
	_ = this

	localctx = NewNotContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, BistatParserRULE_not)

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
		p.SetState(442)
		p.Match(BistatParserT__42)
	}
	{
		p.SetState(443)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(444)
		p.Expression()
	}
	{
		p.SetState(445)
		p.Match(BistatParserT__5)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExp() []IExpContext
	Exp(i int) IExpContext
	LOGIC_OPERATOR() antlr.TerminalNode

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
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

	return t.(IExpContext)
}

func (s *ExpressionContext) LOGIC_OPERATOR() antlr.TerminalNode {
	return s.GetToken(BistatParserLOGIC_OPERATOR, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *BistatParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, BistatParserRULE_expression)
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
		p.SetState(447)
		p.Exp()
	}
	p.SetState(450)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BistatParserLOGIC_OPERATOR {
		{
			p.SetState(448)
			p.Match(BistatParserLOGIC_OPERATOR)
		}
		{
			p.SetState(449)
			p.Exp()
		}

	}

	return localctx
}

// IExpContext is an interface to support dynamic dispatch.
type IExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Term() ITermContext
	OP_SEC() antlr.TerminalNode
	Exp() IExpContext

	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_exp
	return p
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) Term() ITermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ExpContext) OP_SEC() antlr.TerminalNode {
	return s.GetToken(BistatParserOP_SEC, 0)
}

func (s *ExpContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterExp(s)
	}
}

func (s *ExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitExp(s)
	}
}

func (p *BistatParser) Exp() (localctx IExpContext) {
	this := p
	_ = this

	localctx = NewExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, BistatParserRULE_exp)
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
		p.SetState(452)
		p.Term()
	}
	p.SetState(455)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BistatParserOP_SEC {
		{
			p.SetState(453)
			p.Match(BistatParserOP_SEC)
		}
		{
			p.SetState(454)
			p.Exp()
		}

	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Factor() IFactorContext
	OP_FIRST() antlr.TerminalNode
	Term() ITermContext

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) Factor() IFactorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *TermContext) OP_FIRST() antlr.TerminalNode {
	return s.GetToken(BistatParserOP_FIRST, 0)
}

func (s *TermContext) Term() ITermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *BistatParser) Term() (localctx ITermContext) {
	this := p
	_ = this

	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, BistatParserRULE_term)
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
		p.SetState(457)
		p.Factor()
	}
	p.SetState(460)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BistatParserOP_FIRST {
		{
			p.SetState(458)
			p.Match(BistatParserOP_FIRST)
		}
		{
			p.SetState(459)
			p.Term()
		}

	}

	return localctx
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Indexing() IIndexingContext
	SpecialFunction() ISpecialFunctionContext
	FunctionCall() IFunctionCallContext
	VAR_CONS() antlr.TerminalNode
	OP_SEC() antlr.TerminalNode
	Expression() IExpressionContext

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_factor
	return p
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) Indexing() IIndexingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexingContext)
}

func (s *FactorContext) SpecialFunction() ISpecialFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpecialFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpecialFunctionContext)
}

func (s *FactorContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *FactorContext) VAR_CONS() antlr.TerminalNode {
	return s.GetToken(BistatParserVAR_CONS, 0)
}

func (s *FactorContext) OP_SEC() antlr.TerminalNode {
	return s.GetToken(BistatParserOP_SEC, 0)
}

func (s *FactorContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (p *BistatParser) Factor() (localctx IFactorContext) {
	this := p
	_ = this

	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, BistatParserRULE_factor)
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
	p.SetState(463)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BistatParserOP_SEC {
		{
			p.SetState(462)
			p.Match(BistatParserOP_SEC)
		}

	}
	p.SetState(473)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(465)
			p.Match(BistatParserT__4)
		}
		{
			p.SetState(466)
			p.Expression()
		}
		{
			p.SetState(467)
			p.Match(BistatParserT__5)
		}

	case 2:
		{
			p.SetState(469)
			p.Indexing()
		}

	case 3:
		{
			p.SetState(470)
			p.SpecialFunction()
		}

	case 4:
		{
			p.SetState(471)
			p.FunctionCall()
		}

	case 5:
		{
			p.SetState(472)
			p.Match(BistatParserVAR_CONS)
		}

	}

	return localctx
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) ID() antlr.TerminalNode {
	return s.GetToken(BistatParserID, 0)
}

func (s *FunctionCallContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *FunctionCallContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (p *BistatParser) FunctionCall() (localctx IFunctionCallContext) {
	this := p
	_ = this

	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, BistatParserRULE_functionCall)
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
		p.SetState(475)
		p.Match(BistatParserID)
	}
	{
		p.SetState(476)
		p.Match(BistatParserT__4)
	}
	p.SetState(485)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&18454203156398112) != 0 {
		{
			p.SetState(477)
			p.Expression()
		}
		p.SetState(482)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == BistatParserT__12 {
			{
				p.SetState(478)
				p.Match(BistatParserT__12)
			}
			{
				p.SetState(479)
				p.Expression()
			}

			p.SetState(484)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(487)
		p.Match(BistatParserT__5)
	}

	return localctx
}

// IIndexingContext is an interface to support dynamic dispatch.
type IIndexingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsIndexingContext differentiates from other interfaces.
	IsIndexingContext()
}

type IndexingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexingContext() *IndexingContext {
	var p = new(IndexingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_indexing
	return p
}

func (*IndexingContext) IsIndexingContext() {}

func NewIndexingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexingContext {
	var p = new(IndexingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_indexing

	return p
}

func (s *IndexingContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexingContext) ID() antlr.TerminalNode {
	return s.GetToken(BistatParserID, 0)
}

func (s *IndexingContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *IndexingContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *IndexingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterIndexing(s)
	}
}

func (s *IndexingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitIndexing(s)
	}
}

func (p *BistatParser) Indexing() (localctx IIndexingContext) {
	this := p
	_ = this

	localctx = NewIndexingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, BistatParserRULE_indexing)
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
		p.SetState(489)
		p.Match(BistatParserID)
	}
	{
		p.SetState(490)
		p.Match(BistatParserT__11)
	}
	{
		p.SetState(491)
		p.Expression()
	}
	{
		p.SetState(492)
		p.Match(BistatParserT__13)
	}
	p.SetState(497)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BistatParserT__11 {
		{
			p.SetState(493)
			p.Match(BistatParserT__11)
		}
		{
			p.SetState(494)
			p.Expression()
		}
		{
			p.SetState(495)
			p.Match(BistatParserT__13)
		}

	}

	return localctx
}
