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
		"'while'", "'if'", "'else'", "'read'", "'print'", "'listAdd'", "'listPop'",
		"'length'", "'range'", "'plot'", "'sum'", "'min'", "'prod'", "'avg'",
		"'sMode'", "'median'", "'sin'", "'tan'", "'cos'", "'sort'", "'sqrt'",
		"'floor'", "'ceil'", "'abs'", "'not'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "WS", "NON_VOID_TYPE", "RETURN_TYPE",
		"PARAM_TYPE", "ID", "PARAM_CARDINALITY", "CARDINALITY", "VAR_CONS",
		"BOOL_CONS", "STRING_CONS", "FLOAT_CONS", "INT_CONS", "OP_SEC", "OP_FIRST",
		"LOGIC_OPERATOR", "NUMBER",
	}
	staticData.ruleNames = []string{
		"program", "varDeclaration", "funcDef", "main", "stmt", "paramDeclaration",
		"assignment", "nested_stmt", "matrixAssignment", "listAssignment", "comment",
		"forLoop", "whileLoop", "conditional", "specialFunction", "inputRead",
		"print", "listAdd", "listPop", "length", "range", "plot", "sum", "min",
		"prod", "avg", "sMode", "median", "sin", "tan", "cos", "sort", "sqrt",
		"floor", "ceil", "abs", "not", "expression", "exp", "term", "factor",
		"functionCall", "indexing",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 58, 487, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 91, 8, 0, 10, 0, 12, 0, 94, 9,
		0, 1, 0, 5, 0, 97, 8, 0, 10, 0, 12, 0, 100, 9, 0, 1, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 114, 8, 2, 10,
		2, 12, 2, 117, 9, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 123, 8, 2, 10, 2, 12,
		2, 126, 9, 2, 1, 2, 1, 2, 4, 2, 130, 8, 2, 11, 2, 12, 2, 131, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 4, 3, 141, 8, 3, 11, 3, 12, 3, 142, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 4, 3, 4, 150, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 3, 4, 158, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 3, 6, 170, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 176, 8, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 183, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8,
		5, 8, 189, 8, 8, 10, 8, 12, 8, 192, 9, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9,
		1, 9, 5, 9, 200, 8, 9, 10, 9, 12, 9, 203, 9, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		4, 10, 209, 8, 10, 11, 10, 12, 10, 210, 1, 10, 1, 10, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 4, 11, 223, 8, 11, 11, 11, 12, 11,
		224, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 4, 12, 235,
		8, 12, 11, 12, 12, 12, 236, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 4, 13, 247, 8, 13, 11, 13, 12, 13, 248, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 4, 13, 259, 8, 13, 11, 13, 12, 13, 260,
		1, 13, 1, 13, 5, 13, 265, 8, 13, 10, 13, 12, 13, 268, 9, 13, 1, 13, 1,
		13, 1, 13, 4, 13, 273, 8, 13, 11, 13, 12, 13, 274, 1, 13, 1, 13, 3, 13,
		279, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 3, 14, 303, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 5, 15, 310, 8, 15, 10, 15, 12, 15, 313, 9, 15, 1, 15, 1, 15, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 322, 8, 16, 10, 16, 12, 16, 325, 9,
		16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 3, 20, 351, 8, 20, 1, 20, 1, 20, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29,
		1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 3, 37, 438, 8, 37,
		1, 38, 1, 38, 1, 38, 3, 38, 443, 8, 38, 1, 39, 1, 39, 1, 39, 3, 39, 448,
		8, 39, 1, 40, 3, 40, 451, 8, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1,
		40, 1, 40, 1, 40, 3, 40, 461, 8, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41,
		5, 41, 468, 8, 41, 10, 41, 12, 41, 471, 9, 41, 3, 41, 473, 8, 41, 1, 41,
		1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 3, 42, 485,
		8, 42, 1, 42, 0, 0, 43, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
		26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60,
		62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 0, 1, 1, 0, 15, 15, 508,
		0, 86, 1, 0, 0, 0, 2, 104, 1, 0, 0, 0, 4, 109, 1, 0, 0, 0, 6, 135, 1, 0,
		0, 0, 8, 157, 1, 0, 0, 0, 10, 159, 1, 0, 0, 0, 12, 164, 1, 0, 0, 0, 14,
		182, 1, 0, 0, 0, 16, 184, 1, 0, 0, 0, 18, 195, 1, 0, 0, 0, 20, 206, 1,
		0, 0, 0, 22, 214, 1, 0, 0, 0, 24, 228, 1, 0, 0, 0, 26, 240, 1, 0, 0, 0,
		28, 302, 1, 0, 0, 0, 30, 304, 1, 0, 0, 0, 32, 316, 1, 0, 0, 0, 34, 328,
		1, 0, 0, 0, 36, 335, 1, 0, 0, 0, 38, 340, 1, 0, 0, 0, 40, 345, 1, 0, 0,
		0, 42, 354, 1, 0, 0, 0, 44, 359, 1, 0, 0, 0, 46, 364, 1, 0, 0, 0, 48, 369,
		1, 0, 0, 0, 50, 374, 1, 0, 0, 0, 52, 379, 1, 0, 0, 0, 54, 384, 1, 0, 0,
		0, 56, 389, 1, 0, 0, 0, 58, 394, 1, 0, 0, 0, 60, 399, 1, 0, 0, 0, 62, 404,
		1, 0, 0, 0, 64, 409, 1, 0, 0, 0, 66, 414, 1, 0, 0, 0, 68, 419, 1, 0, 0,
		0, 70, 424, 1, 0, 0, 0, 72, 429, 1, 0, 0, 0, 74, 434, 1, 0, 0, 0, 76, 439,
		1, 0, 0, 0, 78, 444, 1, 0, 0, 0, 80, 450, 1, 0, 0, 0, 82, 462, 1, 0, 0,
		0, 84, 476, 1, 0, 0, 0, 86, 87, 5, 1, 0, 0, 87, 88, 5, 47, 0, 0, 88, 92,
		5, 2, 0, 0, 89, 91, 3, 2, 1, 0, 90, 89, 1, 0, 0, 0, 91, 94, 1, 0, 0, 0,
		92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 98, 1, 0, 0, 0, 94, 92, 1,
		0, 0, 0, 95, 97, 3, 4, 2, 0, 96, 95, 1, 0, 0, 0, 97, 100, 1, 0, 0, 0, 98,
		96, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 101, 1, 0, 0, 0, 100, 98, 1, 0,
		0, 0, 101, 102, 3, 6, 3, 0, 102, 103, 5, 0, 0, 1, 103, 1, 1, 0, 0, 0, 104,
		105, 5, 3, 0, 0, 105, 106, 5, 44, 0, 0, 106, 107, 5, 47, 0, 0, 107, 108,
		5, 2, 0, 0, 108, 3, 1, 0, 0, 0, 109, 110, 5, 4, 0, 0, 110, 111, 5, 47,
		0, 0, 111, 115, 5, 5, 0, 0, 112, 114, 3, 10, 5, 0, 113, 112, 1, 0, 0, 0,
		114, 117, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116,
		118, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 118, 119, 5, 6, 0, 0, 119, 120,
		5, 7, 0, 0, 120, 124, 5, 45, 0, 0, 121, 123, 3, 2, 1, 0, 122, 121, 1, 0,
		0, 0, 123, 126, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0,
		125, 127, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 127, 129, 5, 8, 0, 0, 128,
		130, 3, 8, 4, 0, 129, 128, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 129,
		1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 134, 5, 9,
		0, 0, 134, 5, 1, 0, 0, 0, 135, 136, 5, 10, 0, 0, 136, 137, 5, 5, 0, 0,
		137, 138, 5, 6, 0, 0, 138, 140, 5, 8, 0, 0, 139, 141, 3, 8, 4, 0, 140,
		139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 142, 143,
		1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 145, 5, 9, 0, 0, 145, 7, 1, 0, 0,
		0, 146, 150, 3, 12, 6, 0, 147, 150, 3, 28, 14, 0, 148, 150, 3, 82, 41,
		0, 149, 146, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 149, 148, 1, 0, 0, 0, 150,
		151, 1, 0, 0, 0, 151, 152, 5, 2, 0, 0, 152, 158, 1, 0, 0, 0, 153, 158,
		3, 26, 13, 0, 154, 158, 3, 24, 12, 0, 155, 158, 3, 22, 11, 0, 156, 158,
		3, 20, 10, 0, 157, 149, 1, 0, 0, 0, 157, 153, 1, 0, 0, 0, 157, 154, 1,
		0, 0, 0, 157, 155, 1, 0, 0, 0, 157, 156, 1, 0, 0, 0, 158, 9, 1, 0, 0, 0,
		159, 160, 5, 3, 0, 0, 160, 161, 5, 46, 0, 0, 161, 162, 5, 47, 0, 0, 162,
		163, 5, 2, 0, 0, 163, 11, 1, 0, 0, 0, 164, 165, 5, 47, 0, 0, 165, 169,
		5, 11, 0, 0, 166, 170, 5, 50, 0, 0, 167, 170, 3, 18, 9, 0, 168, 170, 3,
		16, 8, 0, 169, 166, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 169, 168, 1, 0, 0,
		0, 170, 13, 1, 0, 0, 0, 171, 176, 3, 12, 6, 0, 172, 176, 3, 32, 16, 0,
		173, 176, 3, 28, 14, 0, 174, 176, 3, 82, 41, 0, 175, 171, 1, 0, 0, 0, 175,
		172, 1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 175, 174, 1, 0, 0, 0, 176, 177,
		1, 0, 0, 0, 177, 178, 5, 2, 0, 0, 178, 183, 1, 0, 0, 0, 179, 183, 3, 26,
		13, 0, 180, 183, 3, 24, 12, 0, 181, 183, 3, 20, 10, 0, 182, 175, 1, 0,
		0, 0, 182, 179, 1, 0, 0, 0, 182, 180, 1, 0, 0, 0, 182, 181, 1, 0, 0, 0,
		183, 15, 1, 0, 0, 0, 184, 185, 5, 12, 0, 0, 185, 190, 3, 18, 9, 0, 186,
		187, 5, 13, 0, 0, 187, 189, 3, 18, 9, 0, 188, 186, 1, 0, 0, 0, 189, 192,
		1, 0, 0, 0, 190, 188, 1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191, 193, 1, 0,
		0, 0, 192, 190, 1, 0, 0, 0, 193, 194, 5, 14, 0, 0, 194, 17, 1, 0, 0, 0,
		195, 196, 5, 12, 0, 0, 196, 201, 5, 50, 0, 0, 197, 198, 5, 13, 0, 0, 198,
		200, 5, 50, 0, 0, 199, 197, 1, 0, 0, 0, 200, 203, 1, 0, 0, 0, 201, 199,
		1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 204, 1, 0, 0, 0, 203, 201, 1, 0,
		0, 0, 204, 205, 5, 14, 0, 0, 205, 19, 1, 0, 0, 0, 206, 208, 5, 15, 0, 0,
		207, 209, 8, 0, 0, 0, 208, 207, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210,
		208, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 212, 1, 0, 0, 0, 212, 213,
		5, 15, 0, 0, 213, 21, 1, 0, 0, 0, 214, 215, 5, 16, 0, 0, 215, 216, 5, 5,
		0, 0, 216, 217, 5, 47, 0, 0, 217, 218, 5, 17, 0, 0, 218, 219, 3, 74, 37,
		0, 219, 220, 5, 6, 0, 0, 220, 222, 5, 8, 0, 0, 221, 223, 3, 14, 7, 0, 222,
		221, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 222, 1, 0, 0, 0, 224, 225,
		1, 0, 0, 0, 225, 226, 1, 0, 0, 0, 226, 227, 5, 9, 0, 0, 227, 23, 1, 0,
		0, 0, 228, 229, 5, 18, 0, 0, 229, 230, 5, 5, 0, 0, 230, 231, 3, 74, 37,
		0, 231, 232, 5, 6, 0, 0, 232, 234, 5, 8, 0, 0, 233, 235, 3, 8, 4, 0, 234,
		233, 1, 0, 0, 0, 235, 236, 1, 0, 0, 0, 236, 234, 1, 0, 0, 0, 236, 237,
		1, 0, 0, 0, 237, 238, 1, 0, 0, 0, 238, 239, 5, 9, 0, 0, 239, 25, 1, 0,
		0, 0, 240, 241, 5, 19, 0, 0, 241, 242, 5, 5, 0, 0, 242, 243, 3, 74, 37,
		0, 243, 244, 5, 6, 0, 0, 244, 246, 5, 8, 0, 0, 245, 247, 3, 8, 4, 0, 246,
		245, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248, 246, 1, 0, 0, 0, 248, 249,
		1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 266, 5, 9, 0, 0, 251, 252, 5, 20,
		0, 0, 252, 253, 5, 19, 0, 0, 253, 254, 5, 5, 0, 0, 254, 255, 3, 74, 37,
		0, 255, 256, 5, 6, 0, 0, 256, 258, 5, 8, 0, 0, 257, 259, 3, 8, 4, 0, 258,
		257, 1, 0, 0, 0, 259, 260, 1, 0, 0, 0, 260, 258, 1, 0, 0, 0, 260, 261,
		1, 0, 0, 0, 261, 262, 1, 0, 0, 0, 262, 263, 5, 9, 0, 0, 263, 265, 1, 0,
		0, 0, 264, 251, 1, 0, 0, 0, 265, 268, 1, 0, 0, 0, 266, 264, 1, 0, 0, 0,
		266, 267, 1, 0, 0, 0, 267, 278, 1, 0, 0, 0, 268, 266, 1, 0, 0, 0, 269,
		270, 5, 20, 0, 0, 270, 272, 5, 8, 0, 0, 271, 273, 3, 8, 4, 0, 272, 271,
		1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 272, 1, 0, 0, 0, 274, 275, 1, 0,
		0, 0, 275, 276, 1, 0, 0, 0, 276, 277, 5, 9, 0, 0, 277, 279, 1, 0, 0, 0,
		278, 269, 1, 0, 0, 0, 278, 279, 1, 0, 0, 0, 279, 27, 1, 0, 0, 0, 280, 303,
		3, 30, 15, 0, 281, 303, 3, 32, 16, 0, 282, 303, 3, 34, 17, 0, 283, 303,
		3, 36, 18, 0, 284, 303, 3, 38, 19, 0, 285, 303, 3, 40, 20, 0, 286, 303,
		3, 42, 21, 0, 287, 303, 3, 44, 22, 0, 288, 303, 3, 46, 23, 0, 289, 303,
		3, 48, 24, 0, 290, 303, 3, 50, 25, 0, 291, 303, 3, 52, 26, 0, 292, 303,
		3, 54, 27, 0, 293, 303, 3, 56, 28, 0, 294, 303, 3, 60, 30, 0, 295, 303,
		3, 58, 29, 0, 296, 303, 3, 62, 31, 0, 297, 303, 3, 64, 32, 0, 298, 303,
		3, 66, 33, 0, 299, 303, 3, 68, 34, 0, 300, 303, 3, 70, 35, 0, 301, 303,
		3, 72, 36, 0, 302, 280, 1, 0, 0, 0, 302, 281, 1, 0, 0, 0, 302, 282, 1,
		0, 0, 0, 302, 283, 1, 0, 0, 0, 302, 284, 1, 0, 0, 0, 302, 285, 1, 0, 0,
		0, 302, 286, 1, 0, 0, 0, 302, 287, 1, 0, 0, 0, 302, 288, 1, 0, 0, 0, 302,
		289, 1, 0, 0, 0, 302, 290, 1, 0, 0, 0, 302, 291, 1, 0, 0, 0, 302, 292,
		1, 0, 0, 0, 302, 293, 1, 0, 0, 0, 302, 294, 1, 0, 0, 0, 302, 295, 1, 0,
		0, 0, 302, 296, 1, 0, 0, 0, 302, 297, 1, 0, 0, 0, 302, 298, 1, 0, 0, 0,
		302, 299, 1, 0, 0, 0, 302, 300, 1, 0, 0, 0, 302, 301, 1, 0, 0, 0, 303,
		29, 1, 0, 0, 0, 304, 305, 5, 21, 0, 0, 305, 306, 5, 5, 0, 0, 306, 311,
		5, 47, 0, 0, 307, 308, 5, 13, 0, 0, 308, 310, 5, 47, 0, 0, 309, 307, 1,
		0, 0, 0, 310, 313, 1, 0, 0, 0, 311, 309, 1, 0, 0, 0, 311, 312, 1, 0, 0,
		0, 312, 314, 1, 0, 0, 0, 313, 311, 1, 0, 0, 0, 314, 315, 5, 6, 0, 0, 315,
		31, 1, 0, 0, 0, 316, 317, 5, 22, 0, 0, 317, 318, 5, 5, 0, 0, 318, 323,
		3, 74, 37, 0, 319, 320, 5, 13, 0, 0, 320, 322, 3, 74, 37, 0, 321, 319,
		1, 0, 0, 0, 322, 325, 1, 0, 0, 0, 323, 321, 1, 0, 0, 0, 323, 324, 1, 0,
		0, 0, 324, 326, 1, 0, 0, 0, 325, 323, 1, 0, 0, 0, 326, 327, 5, 6, 0, 0,
		327, 33, 1, 0, 0, 0, 328, 329, 5, 23, 0, 0, 329, 330, 5, 5, 0, 0, 330,
		331, 3, 74, 37, 0, 331, 332, 5, 13, 0, 0, 332, 333, 3, 74, 37, 0, 333,
		334, 5, 6, 0, 0, 334, 35, 1, 0, 0, 0, 335, 336, 5, 24, 0, 0, 336, 337,
		5, 5, 0, 0, 337, 338, 3, 74, 37, 0, 338, 339, 5, 6, 0, 0, 339, 37, 1, 0,
		0, 0, 340, 341, 5, 25, 0, 0, 341, 342, 5, 5, 0, 0, 342, 343, 3, 74, 37,
		0, 343, 344, 5, 6, 0, 0, 344, 39, 1, 0, 0, 0, 345, 346, 5, 26, 0, 0, 346,
		347, 5, 5, 0, 0, 347, 350, 3, 74, 37, 0, 348, 349, 5, 13, 0, 0, 349, 351,
		3, 74, 37, 0, 350, 348, 1, 0, 0, 0, 350, 351, 1, 0, 0, 0, 351, 352, 1,
		0, 0, 0, 352, 353, 5, 6, 0, 0, 353, 41, 1, 0, 0, 0, 354, 355, 5, 27, 0,
		0, 355, 356, 5, 5, 0, 0, 356, 357, 3, 74, 37, 0, 357, 358, 5, 6, 0, 0,
		358, 43, 1, 0, 0, 0, 359, 360, 5, 28, 0, 0, 360, 361, 5, 5, 0, 0, 361,
		362, 3, 74, 37, 0, 362, 363, 5, 6, 0, 0, 363, 45, 1, 0, 0, 0, 364, 365,
		5, 29, 0, 0, 365, 366, 5, 5, 0, 0, 366, 367, 3, 74, 37, 0, 367, 368, 5,
		6, 0, 0, 368, 47, 1, 0, 0, 0, 369, 370, 5, 30, 0, 0, 370, 371, 5, 5, 0,
		0, 371, 372, 3, 74, 37, 0, 372, 373, 5, 6, 0, 0, 373, 49, 1, 0, 0, 0, 374,
		375, 5, 31, 0, 0, 375, 376, 5, 5, 0, 0, 376, 377, 3, 74, 37, 0, 377, 378,
		5, 6, 0, 0, 378, 51, 1, 0, 0, 0, 379, 380, 5, 32, 0, 0, 380, 381, 5, 5,
		0, 0, 381, 382, 3, 74, 37, 0, 382, 383, 5, 6, 0, 0, 383, 53, 1, 0, 0, 0,
		384, 385, 5, 33, 0, 0, 385, 386, 5, 5, 0, 0, 386, 387, 3, 74, 37, 0, 387,
		388, 5, 6, 0, 0, 388, 55, 1, 0, 0, 0, 389, 390, 5, 34, 0, 0, 390, 391,
		5, 5, 0, 0, 391, 392, 3, 74, 37, 0, 392, 393, 5, 6, 0, 0, 393, 57, 1, 0,
		0, 0, 394, 395, 5, 35, 0, 0, 395, 396, 5, 5, 0, 0, 396, 397, 3, 74, 37,
		0, 397, 398, 5, 6, 0, 0, 398, 59, 1, 0, 0, 0, 399, 400, 5, 36, 0, 0, 400,
		401, 5, 5, 0, 0, 401, 402, 3, 74, 37, 0, 402, 403, 5, 6, 0, 0, 403, 61,
		1, 0, 0, 0, 404, 405, 5, 37, 0, 0, 405, 406, 5, 5, 0, 0, 406, 407, 3, 74,
		37, 0, 407, 408, 5, 6, 0, 0, 408, 63, 1, 0, 0, 0, 409, 410, 5, 38, 0, 0,
		410, 411, 5, 5, 0, 0, 411, 412, 3, 74, 37, 0, 412, 413, 5, 6, 0, 0, 413,
		65, 1, 0, 0, 0, 414, 415, 5, 39, 0, 0, 415, 416, 5, 5, 0, 0, 416, 417,
		3, 74, 37, 0, 417, 418, 5, 6, 0, 0, 418, 67, 1, 0, 0, 0, 419, 420, 5, 40,
		0, 0, 420, 421, 5, 5, 0, 0, 421, 422, 3, 74, 37, 0, 422, 423, 5, 6, 0,
		0, 423, 69, 1, 0, 0, 0, 424, 425, 5, 41, 0, 0, 425, 426, 5, 5, 0, 0, 426,
		427, 3, 74, 37, 0, 427, 428, 5, 6, 0, 0, 428, 71, 1, 0, 0, 0, 429, 430,
		5, 42, 0, 0, 430, 431, 5, 5, 0, 0, 431, 432, 3, 74, 37, 0, 432, 433, 5,
		6, 0, 0, 433, 73, 1, 0, 0, 0, 434, 437, 3, 76, 38, 0, 435, 436, 5, 57,
		0, 0, 436, 438, 3, 76, 38, 0, 437, 435, 1, 0, 0, 0, 437, 438, 1, 0, 0,
		0, 438, 75, 1, 0, 0, 0, 439, 442, 3, 78, 39, 0, 440, 441, 5, 55, 0, 0,
		441, 443, 3, 76, 38, 0, 442, 440, 1, 0, 0, 0, 442, 443, 1, 0, 0, 0, 443,
		77, 1, 0, 0, 0, 444, 447, 3, 80, 40, 0, 445, 446, 5, 56, 0, 0, 446, 448,
		3, 78, 39, 0, 447, 445, 1, 0, 0, 0, 447, 448, 1, 0, 0, 0, 448, 79, 1, 0,
		0, 0, 449, 451, 5, 55, 0, 0, 450, 449, 1, 0, 0, 0, 450, 451, 1, 0, 0, 0,
		451, 460, 1, 0, 0, 0, 452, 453, 5, 5, 0, 0, 453, 454, 3, 74, 37, 0, 454,
		455, 5, 6, 0, 0, 455, 461, 1, 0, 0, 0, 456, 461, 3, 84, 42, 0, 457, 461,
		3, 28, 14, 0, 458, 461, 3, 82, 41, 0, 459, 461, 5, 50, 0, 0, 460, 452,
		1, 0, 0, 0, 460, 456, 1, 0, 0, 0, 460, 457, 1, 0, 0, 0, 460, 458, 1, 0,
		0, 0, 460, 459, 1, 0, 0, 0, 461, 81, 1, 0, 0, 0, 462, 463, 5, 47, 0, 0,
		463, 472, 5, 5, 0, 0, 464, 469, 3, 74, 37, 0, 465, 466, 5, 13, 0, 0, 466,
		468, 3, 74, 37, 0, 467, 465, 1, 0, 0, 0, 468, 471, 1, 0, 0, 0, 469, 467,
		1, 0, 0, 0, 469, 470, 1, 0, 0, 0, 470, 473, 1, 0, 0, 0, 471, 469, 1, 0,
		0, 0, 472, 464, 1, 0, 0, 0, 472, 473, 1, 0, 0, 0, 473, 474, 1, 0, 0, 0,
		474, 475, 5, 6, 0, 0, 475, 83, 1, 0, 0, 0, 476, 477, 5, 47, 0, 0, 477,
		478, 5, 12, 0, 0, 478, 479, 3, 74, 37, 0, 479, 484, 5, 14, 0, 0, 480, 481,
		5, 12, 0, 0, 481, 482, 3, 74, 37, 0, 482, 483, 5, 14, 0, 0, 483, 485, 1,
		0, 0, 0, 484, 480, 1, 0, 0, 0, 484, 485, 1, 0, 0, 0, 485, 85, 1, 0, 0,
		0, 33, 92, 98, 115, 124, 131, 142, 149, 157, 169, 175, 182, 190, 201, 210,
		224, 236, 248, 260, 266, 274, 278, 302, 311, 323, 350, 437, 442, 447, 450,
		460, 469, 472, 484,
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
	BistatParserEOF               = antlr.TokenEOF
	BistatParserT__0              = 1
	BistatParserT__1              = 2
	BistatParserT__2              = 3
	BistatParserT__3              = 4
	BistatParserT__4              = 5
	BistatParserT__5              = 6
	BistatParserT__6              = 7
	BistatParserT__7              = 8
	BistatParserT__8              = 9
	BistatParserT__9              = 10
	BistatParserT__10             = 11
	BistatParserT__11             = 12
	BistatParserT__12             = 13
	BistatParserT__13             = 14
	BistatParserT__14             = 15
	BistatParserT__15             = 16
	BistatParserT__16             = 17
	BistatParserT__17             = 18
	BistatParserT__18             = 19
	BistatParserT__19             = 20
	BistatParserT__20             = 21
	BistatParserT__21             = 22
	BistatParserT__22             = 23
	BistatParserT__23             = 24
	BistatParserT__24             = 25
	BistatParserT__25             = 26
	BistatParserT__26             = 27
	BistatParserT__27             = 28
	BistatParserT__28             = 29
	BistatParserT__29             = 30
	BistatParserT__30             = 31
	BistatParserT__31             = 32
	BistatParserT__32             = 33
	BistatParserT__33             = 34
	BistatParserT__34             = 35
	BistatParserT__35             = 36
	BistatParserT__36             = 37
	BistatParserT__37             = 38
	BistatParserT__38             = 39
	BistatParserT__39             = 40
	BistatParserT__40             = 41
	BistatParserT__41             = 42
	BistatParserWS                = 43
	BistatParserNON_VOID_TYPE     = 44
	BistatParserRETURN_TYPE       = 45
	BistatParserPARAM_TYPE        = 46
	BistatParserID                = 47
	BistatParserPARAM_CARDINALITY = 48
	BistatParserCARDINALITY       = 49
	BistatParserVAR_CONS          = 50
	BistatParserBOOL_CONS         = 51
	BistatParserSTRING_CONS       = 52
	BistatParserFLOAT_CONS        = 53
	BistatParserINT_CONS          = 54
	BistatParserOP_SEC            = 55
	BistatParserOP_FIRST          = 56
	BistatParserLOGIC_OPERATOR    = 57
	BistatParserNUMBER            = 58
)

// BistatParser rules.
const (
	BistatParserRULE_program          = 0
	BistatParserRULE_varDeclaration   = 1
	BistatParserRULE_funcDef          = 2
	BistatParserRULE_main             = 3
	BistatParserRULE_stmt             = 4
	BistatParserRULE_paramDeclaration = 5
	BistatParserRULE_assignment       = 6
	BistatParserRULE_nested_stmt      = 7
	BistatParserRULE_matrixAssignment = 8
	BistatParserRULE_listAssignment   = 9
	BistatParserRULE_comment          = 10
	BistatParserRULE_forLoop          = 11
	BistatParserRULE_whileLoop        = 12
	BistatParserRULE_conditional      = 13
	BistatParserRULE_specialFunction  = 14
	BistatParserRULE_inputRead        = 15
	BistatParserRULE_print            = 16
	BistatParserRULE_listAdd          = 17
	BistatParserRULE_listPop          = 18
	BistatParserRULE_length           = 19
	BistatParserRULE_range            = 20
	BistatParserRULE_plot             = 21
	BistatParserRULE_sum              = 22
	BistatParserRULE_min              = 23
	BistatParserRULE_prod             = 24
	BistatParserRULE_avg              = 25
	BistatParserRULE_sMode            = 26
	BistatParserRULE_median           = 27
	BistatParserRULE_sin              = 28
	BistatParserRULE_tan              = 29
	BistatParserRULE_cos              = 30
	BistatParserRULE_sort             = 31
	BistatParserRULE_sqrt             = 32
	BistatParserRULE_floor            = 33
	BistatParserRULE_ceil             = 34
	BistatParserRULE_abs              = 35
	BistatParserRULE_not              = 36
	BistatParserRULE_expression       = 37
	BistatParserRULE_exp              = 38
	BistatParserRULE_term             = 39
	BistatParserRULE_factor           = 40
	BistatParserRULE_functionCall     = 41
	BistatParserRULE_indexing         = 42
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
		p.SetState(86)
		p.Match(BistatParserT__0)
	}
	{
		p.SetState(87)
		p.Match(BistatParserID)
	}
	{
		p.SetState(88)
		p.Match(BistatParserT__1)
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__2 {
		{
			p.SetState(89)
			p.VarDeclaration()
		}

		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__3 {
		{
			p.SetState(95)
			p.FuncDef()
		}

		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(101)
		p.Main()
	}
	{
		p.SetState(102)
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
	NON_VOID_TYPE() antlr.TerminalNode
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

func (s *VarDeclarationContext) NON_VOID_TYPE() antlr.TerminalNode {
	return s.GetToken(BistatParserNON_VOID_TYPE, 0)
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
		p.SetState(104)
		p.Match(BistatParserT__2)
	}
	{
		p.SetState(105)
		p.Match(BistatParserNON_VOID_TYPE)
	}
	{
		p.SetState(106)
		p.Match(BistatParserID)
	}
	{
		p.SetState(107)
		p.Match(BistatParserT__1)
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
	RETURN_TYPE() antlr.TerminalNode
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

func (s *FuncDefContext) RETURN_TYPE() antlr.TerminalNode {
	return s.GetToken(BistatParserRETURN_TYPE, 0)
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
	p.EnterRule(localctx, 4, BistatParserRULE_funcDef)
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
		p.SetState(109)
		p.Match(BistatParserT__3)
	}
	{
		p.SetState(110)
		p.Match(BistatParserID)
	}
	{
		p.SetState(111)
		p.Match(BistatParserT__4)
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__2 {
		{
			p.SetState(112)
			p.ParamDeclaration()
		}

		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(118)
		p.Match(BistatParserT__5)
	}
	{
		p.SetState(119)
		p.Match(BistatParserT__6)
	}
	{
		p.SetState(120)
		p.Match(BistatParserRETURN_TYPE)
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__2 {
		{
			p.SetState(121)
			p.VarDeclaration()
		}

		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(127)
		p.Match(BistatParserT__7)
	}
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&149533580165120) != 0) {
		{
			p.SetState(128)
			p.Stmt()
		}

		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(133)
		p.Match(BistatParserT__8)
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
	p.EnterRule(localctx, 6, BistatParserRULE_main)
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
		p.SetState(135)
		p.Match(BistatParserT__9)
	}
	{
		p.SetState(136)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(137)
		p.Match(BistatParserT__5)
	}
	{
		p.SetState(138)
		p.Match(BistatParserT__7)
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&149533580165120) != 0) {
		{
			p.SetState(139)
			p.Stmt()
		}

		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(144)
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
	p.EnterRule(localctx, 8, BistatParserRULE_stmt)

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

	p.SetState(157)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BistatParserT__20, BistatParserT__21, BistatParserT__22, BistatParserT__23, BistatParserT__24, BistatParserT__25, BistatParserT__26, BistatParserT__27, BistatParserT__28, BistatParserT__29, BistatParserT__30, BistatParserT__31, BistatParserT__32, BistatParserT__33, BistatParserT__34, BistatParserT__35, BistatParserT__36, BistatParserT__37, BistatParserT__38, BistatParserT__39, BistatParserT__40, BistatParserT__41, BistatParserID:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(146)
				p.Assignment()
			}

		case 2:
			{
				p.SetState(147)
				p.SpecialFunction()
			}

		case 3:
			{
				p.SetState(148)
				p.FunctionCall()
			}

		}
		{
			p.SetState(151)
			p.Match(BistatParserT__1)
		}

	case BistatParserT__18:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(153)
			p.Conditional()
		}

	case BistatParserT__17:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(154)
			p.WhileLoop()
		}

	case BistatParserT__15:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(155)
			p.ForLoop()
		}

	case BistatParserT__14:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(156)
			p.Comment()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IParamDeclarationContext is an interface to support dynamic dispatch.
type IParamDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PARAM_TYPE() antlr.TerminalNode
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

func (s *ParamDeclarationContext) PARAM_TYPE() antlr.TerminalNode {
	return s.GetToken(BistatParserPARAM_TYPE, 0)
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
	p.EnterRule(localctx, 10, BistatParserRULE_paramDeclaration)

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
		p.SetState(159)
		p.Match(BistatParserT__2)
	}
	{
		p.SetState(160)
		p.Match(BistatParserPARAM_TYPE)
	}
	{
		p.SetState(161)
		p.Match(BistatParserID)
	}
	{
		p.SetState(162)
		p.Match(BistatParserT__1)
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
	p.EnterRule(localctx, 12, BistatParserRULE_assignment)

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
		p.SetState(164)
		p.Match(BistatParserID)
	}
	{
		p.SetState(165)
		p.Match(BistatParserT__10)
	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(166)
			p.Match(BistatParserVAR_CONS)
		}

	case 2:
		{
			p.SetState(167)
			p.ListAssignment()
		}

	case 3:
		{
			p.SetState(168)
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
	p.EnterRule(localctx, 14, BistatParserRULE_nested_stmt)

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

	p.SetState(182)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BistatParserT__20, BistatParserT__21, BistatParserT__22, BistatParserT__23, BistatParserT__24, BistatParserT__25, BistatParserT__26, BistatParserT__27, BistatParserT__28, BistatParserT__29, BistatParserT__30, BistatParserT__31, BistatParserT__32, BistatParserT__33, BistatParserT__34, BistatParserT__35, BistatParserT__36, BistatParserT__37, BistatParserT__38, BistatParserT__39, BistatParserT__40, BistatParserT__41, BistatParserID:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(171)
				p.Assignment()
			}

		case 2:
			{
				p.SetState(172)
				p.Print_()
			}

		case 3:
			{
				p.SetState(173)
				p.SpecialFunction()
			}

		case 4:
			{
				p.SetState(174)
				p.FunctionCall()
			}

		}
		{
			p.SetState(177)
			p.Match(BistatParserT__1)
		}

	case BistatParserT__18:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(179)
			p.Conditional()
		}

	case BistatParserT__17:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(180)
			p.WhileLoop()
		}

	case BistatParserT__14:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(181)
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
	p.EnterRule(localctx, 16, BistatParserRULE_matrixAssignment)
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
		p.SetState(184)
		p.Match(BistatParserT__11)
	}
	{
		p.SetState(185)
		p.ListAssignment()
	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__12 {
		{
			p.SetState(186)
			p.Match(BistatParserT__12)
		}
		{
			p.SetState(187)
			p.ListAssignment()
		}

		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(193)
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
	p.EnterRule(localctx, 18, BistatParserRULE_listAssignment)
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
		p.SetState(195)
		p.Match(BistatParserT__11)
	}
	{
		p.SetState(196)
		p.Match(BistatParserVAR_CONS)
	}
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__12 {
		{
			p.SetState(197)
			p.Match(BistatParserT__12)
		}
		{
			p.SetState(198)
			p.Match(BistatParserVAR_CONS)
		}

		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(204)
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
	p.EnterRule(localctx, 20, BistatParserRULE_comment)
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
		p.SetState(206)
		p.Match(BistatParserT__14)
	}
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&576460752303390718) != 0) {
		{
			p.SetState(207)
			_la = p.GetTokenStream().LA(1)

			if _la <= 0 || _la == BistatParserT__14 {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(212)
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
	p.EnterRule(localctx, 22, BistatParserRULE_forLoop)
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
		p.SetState(214)
		p.Match(BistatParserT__15)
	}
	{
		p.SetState(215)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(216)
		p.Match(BistatParserID)
	}
	{
		p.SetState(217)
		p.Match(BistatParserT__16)
	}
	{
		p.SetState(218)
		p.Expression()
	}
	{
		p.SetState(219)
		p.Match(BistatParserT__5)
	}
	{
		p.SetState(220)
		p.Match(BistatParserT__7)
	}
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&149533580099584) != 0) {
		{
			p.SetState(221)
			p.Nested_stmt()
		}

		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(226)
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
	p.EnterRule(localctx, 24, BistatParserRULE_whileLoop)
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
		p.SetState(228)
		p.Match(BistatParserT__17)
	}
	{
		p.SetState(229)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(230)
		p.Expression()
	}
	{
		p.SetState(231)
		p.Match(BistatParserT__5)
	}
	{
		p.SetState(232)
		p.Match(BistatParserT__7)
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&149533580165120) != 0) {
		{
			p.SetState(233)
			p.Stmt()
		}

		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(238)
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
	p.EnterRule(localctx, 26, BistatParserRULE_conditional)
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
		p.SetState(240)
		p.Match(BistatParserT__18)
	}
	{
		p.SetState(241)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(242)
		p.Expression()
	}
	{
		p.SetState(243)
		p.Match(BistatParserT__5)
	}
	{
		p.SetState(244)
		p.Match(BistatParserT__7)
	}
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&149533580165120) != 0) {
		{
			p.SetState(245)
			p.Stmt()
		}

		p.SetState(248)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(250)
		p.Match(BistatParserT__8)
	}
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(251)
				p.Match(BistatParserT__19)
			}
			{
				p.SetState(252)
				p.Match(BistatParserT__18)
			}
			{
				p.SetState(253)
				p.Match(BistatParserT__4)
			}
			{
				p.SetState(254)
				p.Expression()
			}
			{
				p.SetState(255)
				p.Match(BistatParserT__5)
			}
			{
				p.SetState(256)
				p.Match(BistatParserT__7)
			}
			p.SetState(258)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&149533580165120) != 0) {
				{
					p.SetState(257)
					p.Stmt()
				}

				p.SetState(260)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(262)
				p.Match(BistatParserT__8)
			}

		}
		p.SetState(268)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
	}
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BistatParserT__19 {
		{
			p.SetState(269)
			p.Match(BistatParserT__19)
		}
		{
			p.SetState(270)
			p.Match(BistatParserT__7)
		}
		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&149533580165120) != 0) {
			{
				p.SetState(271)
				p.Stmt()
			}

			p.SetState(274)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(276)
			p.Match(BistatParserT__8)
		}

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
	p.EnterRule(localctx, 28, BistatParserRULE_specialFunction)

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

	p.SetState(302)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BistatParserT__20:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(280)
			p.InputRead()
		}

	case BistatParserT__21:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(281)
			p.Print_()
		}

	case BistatParserT__22:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(282)
			p.ListAdd()
		}

	case BistatParserT__23:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(283)
			p.ListPop()
		}

	case BistatParserT__24:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(284)
			p.Length()
		}

	case BistatParserT__25:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(285)
			p.Range_()
		}

	case BistatParserT__26:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(286)
			p.Plot()
		}

	case BistatParserT__27:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(287)
			p.Sum()
		}

	case BistatParserT__28:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(288)
			p.Min()
		}

	case BistatParserT__29:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(289)
			p.Prod()
		}

	case BistatParserT__30:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(290)
			p.Avg()
		}

	case BistatParserT__31:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(291)
			p.SMode()
		}

	case BistatParserT__32:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(292)
			p.Median()
		}

	case BistatParserT__33:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(293)
			p.Sin()
		}

	case BistatParserT__35:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(294)
			p.Cos()
		}

	case BistatParserT__34:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(295)
			p.Tan()
		}

	case BistatParserT__36:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(296)
			p.Sort()
		}

	case BistatParserT__37:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(297)
			p.Sqrt()
		}

	case BistatParserT__38:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(298)
			p.Floor()
		}

	case BistatParserT__39:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(299)
			p.Ceil()
		}

	case BistatParserT__40:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(300)
			p.Abs()
		}

	case BistatParserT__41:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(301)
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
	p.EnterRule(localctx, 30, BistatParserRULE_inputRead)
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
		p.SetState(304)
		p.Match(BistatParserT__20)
	}
	{
		p.SetState(305)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(306)
		p.Match(BistatParserID)
	}
	p.SetState(311)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__12 {
		{
			p.SetState(307)
			p.Match(BistatParserT__12)
		}
		{
			p.SetState(308)
			p.Match(BistatParserID)
		}

		p.SetState(313)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(314)
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
	p.EnterRule(localctx, 32, BistatParserRULE_print)
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
		p.SetState(316)
		p.Match(BistatParserT__21)
	}
	{
		p.SetState(317)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(318)
		p.Expression()
	}
	p.SetState(323)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__12 {
		{
			p.SetState(319)
			p.Match(BistatParserT__12)
		}
		{
			p.SetState(320)
			p.Expression()
		}

		p.SetState(325)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(326)
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
	p.EnterRule(localctx, 34, BistatParserRULE_listAdd)

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
		p.SetState(328)
		p.Match(BistatParserT__22)
	}
	{
		p.SetState(329)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(330)
		p.Expression()
	}
	{
		p.SetState(331)
		p.Match(BistatParserT__12)
	}
	{
		p.SetState(332)
		p.Expression()
	}
	{
		p.SetState(333)
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
	p.EnterRule(localctx, 36, BistatParserRULE_listPop)

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
		p.SetState(335)
		p.Match(BistatParserT__23)
	}
	{
		p.SetState(336)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(337)
		p.Expression()
	}
	{
		p.SetState(338)
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
	p.EnterRule(localctx, 38, BistatParserRULE_length)

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
		p.SetState(340)
		p.Match(BistatParserT__24)
	}
	{
		p.SetState(341)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(342)
		p.Expression()
	}
	{
		p.SetState(343)
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
	p.EnterRule(localctx, 40, BistatParserRULE_range)
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
		p.SetState(345)
		p.Match(BistatParserT__25)
	}
	{
		p.SetState(346)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(347)
		p.Expression()
	}
	p.SetState(350)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BistatParserT__12 {
		{
			p.SetState(348)
			p.Match(BistatParserT__12)
		}
		{
			p.SetState(349)
			p.Expression()
		}

	}
	{
		p.SetState(352)
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
	p.EnterRule(localctx, 42, BistatParserRULE_plot)

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
		p.SetState(354)
		p.Match(BistatParserT__26)
	}
	{
		p.SetState(355)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(356)
		p.Expression()
	}
	{
		p.SetState(357)
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
	p.EnterRule(localctx, 44, BistatParserRULE_sum)

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
		p.SetState(359)
		p.Match(BistatParserT__27)
	}
	{
		p.SetState(360)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(361)
		p.Expression()
	}
	{
		p.SetState(362)
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
	p.EnterRule(localctx, 46, BistatParserRULE_min)

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
		p.SetState(364)
		p.Match(BistatParserT__28)
	}
	{
		p.SetState(365)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(366)
		p.Expression()
	}
	{
		p.SetState(367)
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
	p.EnterRule(localctx, 48, BistatParserRULE_prod)

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
		p.SetState(369)
		p.Match(BistatParserT__29)
	}
	{
		p.SetState(370)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(371)
		p.Expression()
	}
	{
		p.SetState(372)
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
	p.EnterRule(localctx, 50, BistatParserRULE_avg)

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
		p.SetState(374)
		p.Match(BistatParserT__30)
	}
	{
		p.SetState(375)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(376)
		p.Expression()
	}
	{
		p.SetState(377)
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
	p.EnterRule(localctx, 52, BistatParserRULE_sMode)

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
		p.SetState(379)
		p.Match(BistatParserT__31)
	}
	{
		p.SetState(380)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(381)
		p.Expression()
	}
	{
		p.SetState(382)
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
	p.EnterRule(localctx, 54, BistatParserRULE_median)

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
		p.SetState(384)
		p.Match(BistatParserT__32)
	}
	{
		p.SetState(385)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(386)
		p.Expression()
	}
	{
		p.SetState(387)
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
	p.EnterRule(localctx, 56, BistatParserRULE_sin)

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
		p.SetState(389)
		p.Match(BistatParserT__33)
	}
	{
		p.SetState(390)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(391)
		p.Expression()
	}
	{
		p.SetState(392)
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
	p.EnterRule(localctx, 58, BistatParserRULE_tan)

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
		p.SetState(394)
		p.Match(BistatParserT__34)
	}
	{
		p.SetState(395)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(396)
		p.Expression()
	}
	{
		p.SetState(397)
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
	p.EnterRule(localctx, 60, BistatParserRULE_cos)

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
		p.SetState(399)
		p.Match(BistatParserT__35)
	}
	{
		p.SetState(400)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(401)
		p.Expression()
	}
	{
		p.SetState(402)
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
	p.EnterRule(localctx, 62, BistatParserRULE_sort)

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
		p.SetState(404)
		p.Match(BistatParserT__36)
	}
	{
		p.SetState(405)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(406)
		p.Expression()
	}
	{
		p.SetState(407)
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
	p.EnterRule(localctx, 64, BistatParserRULE_sqrt)

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
		p.SetState(409)
		p.Match(BistatParserT__37)
	}
	{
		p.SetState(410)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(411)
		p.Expression()
	}
	{
		p.SetState(412)
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
	p.EnterRule(localctx, 66, BistatParserRULE_floor)

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
		p.SetState(414)
		p.Match(BistatParserT__38)
	}
	{
		p.SetState(415)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(416)
		p.Expression()
	}
	{
		p.SetState(417)
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
	p.EnterRule(localctx, 68, BistatParserRULE_ceil)

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
		p.SetState(419)
		p.Match(BistatParserT__39)
	}
	{
		p.SetState(420)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(421)
		p.Expression()
	}
	{
		p.SetState(422)
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
	p.EnterRule(localctx, 70, BistatParserRULE_abs)

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
		p.SetState(424)
		p.Match(BistatParserT__40)
	}
	{
		p.SetState(425)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(426)
		p.Expression()
	}
	{
		p.SetState(427)
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
	p.EnterRule(localctx, 72, BistatParserRULE_not)

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
		p.SetState(429)
		p.Match(BistatParserT__41)
	}
	{
		p.SetState(430)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(431)
		p.Expression()
	}
	{
		p.SetState(432)
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
	p.EnterRule(localctx, 74, BistatParserRULE_expression)
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
		p.SetState(434)
		p.Exp()
	}
	p.SetState(437)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BistatParserLOGIC_OPERATOR {
		{
			p.SetState(435)
			p.Match(BistatParserLOGIC_OPERATOR)
		}
		{
			p.SetState(436)
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
	p.EnterRule(localctx, 76, BistatParserRULE_exp)
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
		p.SetState(439)
		p.Term()
	}
	p.SetState(442)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BistatParserOP_SEC {
		{
			p.SetState(440)
			p.Match(BistatParserOP_SEC)
		}
		{
			p.SetState(441)
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
	p.EnterRule(localctx, 78, BistatParserRULE_term)
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
		p.SetState(444)
		p.Factor()
	}
	p.SetState(447)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BistatParserOP_FIRST {
		{
			p.SetState(445)
			p.Match(BistatParserOP_FIRST)
		}
		{
			p.SetState(446)
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
	p.EnterRule(localctx, 80, BistatParserRULE_factor)
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
	p.SetState(450)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BistatParserOP_SEC {
		{
			p.SetState(449)
			p.Match(BistatParserOP_SEC)
		}

	}
	p.SetState(460)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(452)
			p.Match(BistatParserT__4)
		}
		{
			p.SetState(453)
			p.Expression()
		}
		{
			p.SetState(454)
			p.Match(BistatParserT__5)
		}

	case 2:
		{
			p.SetState(456)
			p.Indexing()
		}

	case 3:
		{
			p.SetState(457)
			p.SpecialFunction()
		}

	case 4:
		{
			p.SetState(458)
			p.FunctionCall()
		}

	case 5:
		{
			p.SetState(459)
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
	p.EnterRule(localctx, 82, BistatParserRULE_functionCall)
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
		p.SetState(462)
		p.Match(BistatParserID)
	}
	{
		p.SetState(463)
		p.Match(BistatParserT__4)
	}
	p.SetState(472)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&37304230505087008) != 0 {
		{
			p.SetState(464)
			p.Expression()
		}
		p.SetState(469)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == BistatParserT__12 {
			{
				p.SetState(465)
				p.Match(BistatParserT__12)
			}
			{
				p.SetState(466)
				p.Expression()
			}

			p.SetState(471)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(474)
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
	p.EnterRule(localctx, 84, BistatParserRULE_indexing)
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
		p.SetState(476)
		p.Match(BistatParserID)
	}
	{
		p.SetState(477)
		p.Match(BistatParserT__11)
	}
	{
		p.SetState(478)
		p.Expression()
	}
	{
		p.SetState(479)
		p.Match(BistatParserT__13)
	}
	p.SetState(484)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BistatParserT__11 {
		{
			p.SetState(480)
			p.Match(BistatParserT__11)
		}
		{
			p.SetState(481)
			p.Expression()
		}
		{
			p.SetState(482)
			p.Match(BistatParserT__13)
		}

	}

	return localctx
}
