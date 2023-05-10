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
		"'sqrt'", "'floor'", "'ceil'", "'abs'", "'not'", "'-'", "'+'", "'/'",
		"'*'", "'%'", "'<'", "'>'", "'<='", "'>='", "'=='", "'!='", "'&&'",
		"'||'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "WS", "CARDINALITY", "TYPE_PRIMITIVE", "ID",
		"INT_CONS", "NUMBER", "BOOL_CONS", "STRING_CONS", "FLOAT_CONS",
	}
	staticData.ruleNames = []string{
		"program", "varDeclaration", "var_type", "funcDef", "funcEnd", "paramDeclaration",
		"main", "stmt", "assignment", "matrixAssignment", "listAssignment",
		"comment", "forLoop", "forHeader", "forExprEnd", "whileLoop", "whileExprEnd",
		"conditional", "ifStmt", "elseIfStmt", "condExprEnd", "elseStmt", "returnStmt",
		"specialFunction", "inputRead", "print", "listAdd", "listPop", "length",
		"range", "plot", "sum", "min", "prod", "avg", "sMode", "median", "sin",
		"tan", "cos", "sort", "sqrt", "floor", "ceil", "abs", "not", "expression",
		"exp", "term", "factor", "unaryMinus", "nestedExpression", "functionCall",
		"indexing", "varCons", "opSec", "opFirst", "logicOperator",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 65, 552, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57, 7, 57,
		1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 121, 8, 0, 10, 0, 12, 0, 124, 9, 0, 1, 0,
		5, 0, 127, 8, 0, 10, 0, 12, 0, 130, 9, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 3, 2, 142, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3,
		5, 3, 148, 8, 3, 10, 3, 12, 3, 151, 9, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3,
		157, 8, 3, 10, 3, 12, 3, 160, 9, 3, 1, 3, 1, 3, 4, 3, 164, 8, 3, 11, 3,
		12, 3, 165, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 4, 6, 182, 8, 6, 11, 6, 12, 6, 183, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 192, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 3, 7, 200, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 207, 8, 8, 1,
		9, 1, 9, 1, 9, 1, 9, 5, 9, 213, 8, 9, 10, 9, 12, 9, 216, 9, 9, 1, 9, 1,
		9, 1, 10, 1, 10, 1, 10, 3, 10, 223, 8, 10, 1, 10, 1, 10, 1, 10, 3, 10,
		228, 8, 10, 5, 10, 230, 8, 10, 10, 10, 12, 10, 233, 9, 10, 1, 10, 1, 10,
		1, 11, 1, 11, 4, 11, 239, 8, 11, 11, 11, 12, 11, 240, 1, 11, 1, 11, 1,
		12, 1, 12, 1, 12, 4, 12, 248, 8, 12, 11, 12, 12, 12, 249, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 4, 15, 269, 8, 15, 11, 15, 12, 15, 270,
		1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 5, 17, 279, 8, 17, 10, 17, 12,
		17, 282, 9, 17, 1, 17, 3, 17, 285, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 4, 18, 293, 8, 18, 11, 18, 12, 18, 294, 1, 18, 1, 18, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 4, 19, 306, 8, 19, 11, 19, 12,
		19, 307, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 4, 21, 317, 8,
		21, 11, 21, 12, 21, 318, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23,
		348, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 5, 24, 355, 8, 24, 10, 24,
		12, 24, 358, 9, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 5,
		25, 367, 8, 25, 10, 25, 12, 25, 370, 9, 25, 1, 25, 1, 25, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29,
		396, 8, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33,
		1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37,
		1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1,
		39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41,
		1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1,
		43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45,
		1, 46, 1, 46, 1, 46, 1, 46, 3, 46, 484, 8, 46, 1, 47, 1, 47, 1, 47, 1,
		47, 5, 47, 490, 8, 47, 10, 47, 12, 47, 493, 9, 47, 1, 48, 1, 48, 1, 48,
		1, 48, 5, 48, 499, 8, 48, 10, 48, 12, 48, 502, 9, 48, 1, 49, 3, 49, 505,
		8, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 3, 49, 512, 8, 49, 1, 50, 1,
		50, 1, 51, 1, 51, 1, 51, 1, 51, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 5, 52,
		525, 8, 52, 10, 52, 12, 52, 528, 9, 52, 3, 52, 530, 8, 52, 1, 52, 1, 52,
		1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 3, 53, 542, 8,
		53, 1, 54, 1, 54, 1, 55, 1, 55, 1, 56, 1, 56, 1, 57, 1, 57, 1, 57, 0, 0,
		58, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70,
		72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104,
		106, 108, 110, 112, 114, 0, 5, 1, 0, 15, 15, 2, 0, 60, 61, 63, 65, 1, 0,
		44, 45, 1, 0, 46, 48, 1, 0, 49, 56, 556, 0, 116, 1, 0, 0, 0, 2, 134, 1,
		0, 0, 0, 4, 139, 1, 0, 0, 0, 6, 143, 1, 0, 0, 0, 8, 169, 1, 0, 0, 0, 10,
		171, 1, 0, 0, 0, 12, 176, 1, 0, 0, 0, 14, 199, 1, 0, 0, 0, 16, 201, 1,
		0, 0, 0, 18, 208, 1, 0, 0, 0, 20, 219, 1, 0, 0, 0, 22, 236, 1, 0, 0, 0,
		24, 244, 1, 0, 0, 0, 26, 253, 1, 0, 0, 0, 28, 260, 1, 0, 0, 0, 30, 262,
		1, 0, 0, 0, 32, 274, 1, 0, 0, 0, 34, 276, 1, 0, 0, 0, 36, 286, 1, 0, 0,
		0, 38, 298, 1, 0, 0, 0, 40, 311, 1, 0, 0, 0, 42, 313, 1, 0, 0, 0, 44, 322,
		1, 0, 0, 0, 46, 347, 1, 0, 0, 0, 48, 349, 1, 0, 0, 0, 50, 361, 1, 0, 0,
		0, 52, 373, 1, 0, 0, 0, 54, 380, 1, 0, 0, 0, 56, 385, 1, 0, 0, 0, 58, 390,
		1, 0, 0, 0, 60, 399, 1, 0, 0, 0, 62, 404, 1, 0, 0, 0, 64, 409, 1, 0, 0,
		0, 66, 414, 1, 0, 0, 0, 68, 419, 1, 0, 0, 0, 70, 424, 1, 0, 0, 0, 72, 429,
		1, 0, 0, 0, 74, 434, 1, 0, 0, 0, 76, 439, 1, 0, 0, 0, 78, 444, 1, 0, 0,
		0, 80, 449, 1, 0, 0, 0, 82, 454, 1, 0, 0, 0, 84, 459, 1, 0, 0, 0, 86, 464,
		1, 0, 0, 0, 88, 469, 1, 0, 0, 0, 90, 474, 1, 0, 0, 0, 92, 479, 1, 0, 0,
		0, 94, 485, 1, 0, 0, 0, 96, 494, 1, 0, 0, 0, 98, 504, 1, 0, 0, 0, 100,
		513, 1, 0, 0, 0, 102, 515, 1, 0, 0, 0, 104, 519, 1, 0, 0, 0, 106, 533,
		1, 0, 0, 0, 108, 543, 1, 0, 0, 0, 110, 545, 1, 0, 0, 0, 112, 547, 1, 0,
		0, 0, 114, 549, 1, 0, 0, 0, 116, 117, 5, 1, 0, 0, 117, 118, 5, 60, 0, 0,
		118, 122, 5, 2, 0, 0, 119, 121, 3, 2, 1, 0, 120, 119, 1, 0, 0, 0, 121,
		124, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 128,
		1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 125, 127, 3, 6, 3, 0, 126, 125, 1, 0,
		0, 0, 127, 130, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0,
		129, 131, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 131, 132, 3, 12, 6, 0, 132,
		133, 5, 0, 0, 1, 133, 1, 1, 0, 0, 0, 134, 135, 5, 3, 0, 0, 135, 136, 3,
		4, 2, 0, 136, 137, 5, 60, 0, 0, 137, 138, 5, 2, 0, 0, 138, 3, 1, 0, 0,
		0, 139, 141, 5, 59, 0, 0, 140, 142, 5, 58, 0, 0, 141, 140, 1, 0, 0, 0,
		141, 142, 1, 0, 0, 0, 142, 5, 1, 0, 0, 0, 143, 144, 5, 4, 0, 0, 144, 145,
		5, 60, 0, 0, 145, 149, 5, 5, 0, 0, 146, 148, 3, 10, 5, 0, 147, 146, 1,
		0, 0, 0, 148, 151, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 149, 150, 1, 0, 0,
		0, 150, 152, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 152, 153, 5, 6, 0, 0, 153,
		154, 5, 7, 0, 0, 154, 158, 3, 4, 2, 0, 155, 157, 3, 2, 1, 0, 156, 155,
		1, 0, 0, 0, 157, 160, 1, 0, 0, 0, 158, 156, 1, 0, 0, 0, 158, 159, 1, 0,
		0, 0, 159, 161, 1, 0, 0, 0, 160, 158, 1, 0, 0, 0, 161, 163, 5, 8, 0, 0,
		162, 164, 3, 14, 7, 0, 163, 162, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165,
		163, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 168,
		3, 8, 4, 0, 168, 7, 1, 0, 0, 0, 169, 170, 5, 9, 0, 0, 170, 9, 1, 0, 0,
		0, 171, 172, 5, 3, 0, 0, 172, 173, 3, 4, 2, 0, 173, 174, 5, 60, 0, 0, 174,
		175, 5, 2, 0, 0, 175, 11, 1, 0, 0, 0, 176, 177, 5, 10, 0, 0, 177, 178,
		5, 5, 0, 0, 178, 179, 5, 6, 0, 0, 179, 181, 5, 8, 0, 0, 180, 182, 3, 14,
		7, 0, 181, 180, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0,
		183, 184, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 186, 5, 9, 0, 0, 186,
		13, 1, 0, 0, 0, 187, 192, 3, 16, 8, 0, 188, 192, 3, 46, 23, 0, 189, 192,
		3, 104, 52, 0, 190, 192, 3, 44, 22, 0, 191, 187, 1, 0, 0, 0, 191, 188,
		1, 0, 0, 0, 191, 189, 1, 0, 0, 0, 191, 190, 1, 0, 0, 0, 192, 193, 1, 0,
		0, 0, 193, 194, 5, 2, 0, 0, 194, 200, 1, 0, 0, 0, 195, 200, 3, 34, 17,
		0, 196, 200, 3, 30, 15, 0, 197, 200, 3, 24, 12, 0, 198, 200, 3, 22, 11,
		0, 199, 191, 1, 0, 0, 0, 199, 195, 1, 0, 0, 0, 199, 196, 1, 0, 0, 0, 199,
		197, 1, 0, 0, 0, 199, 198, 1, 0, 0, 0, 200, 15, 1, 0, 0, 0, 201, 202, 5,
		60, 0, 0, 202, 206, 5, 11, 0, 0, 203, 207, 3, 92, 46, 0, 204, 207, 3, 20,
		10, 0, 205, 207, 3, 18, 9, 0, 206, 203, 1, 0, 0, 0, 206, 204, 1, 0, 0,
		0, 206, 205, 1, 0, 0, 0, 207, 17, 1, 0, 0, 0, 208, 209, 5, 12, 0, 0, 209,
		214, 3, 20, 10, 0, 210, 211, 5, 13, 0, 0, 211, 213, 3, 20, 10, 0, 212,
		210, 1, 0, 0, 0, 213, 216, 1, 0, 0, 0, 214, 212, 1, 0, 0, 0, 214, 215,
		1, 0, 0, 0, 215, 217, 1, 0, 0, 0, 216, 214, 1, 0, 0, 0, 217, 218, 5, 14,
		0, 0, 218, 19, 1, 0, 0, 0, 219, 222, 5, 12, 0, 0, 220, 223, 3, 108, 54,
		0, 221, 223, 3, 92, 46, 0, 222, 220, 1, 0, 0, 0, 222, 221, 1, 0, 0, 0,
		223, 231, 1, 0, 0, 0, 224, 227, 5, 13, 0, 0, 225, 228, 3, 108, 54, 0, 226,
		228, 3, 92, 46, 0, 227, 225, 1, 0, 0, 0, 227, 226, 1, 0, 0, 0, 228, 230,
		1, 0, 0, 0, 229, 224, 1, 0, 0, 0, 230, 233, 1, 0, 0, 0, 231, 229, 1, 0,
		0, 0, 231, 232, 1, 0, 0, 0, 232, 234, 1, 0, 0, 0, 233, 231, 1, 0, 0, 0,
		234, 235, 5, 14, 0, 0, 235, 21, 1, 0, 0, 0, 236, 238, 5, 15, 0, 0, 237,
		239, 8, 0, 0, 0, 238, 237, 1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240, 238,
		1, 0, 0, 0, 240, 241, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242, 243, 5, 15,
		0, 0, 243, 23, 1, 0, 0, 0, 244, 245, 3, 26, 13, 0, 245, 247, 5, 8, 0, 0,
		246, 248, 3, 14, 7, 0, 247, 246, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249,
		247, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 251, 1, 0, 0, 0, 251, 252,
		5, 9, 0, 0, 252, 25, 1, 0, 0, 0, 253, 254, 5, 16, 0, 0, 254, 255, 5, 5,
		0, 0, 255, 256, 5, 60, 0, 0, 256, 257, 5, 17, 0, 0, 257, 258, 3, 92, 46,
		0, 258, 259, 5, 6, 0, 0, 259, 27, 1, 0, 0, 0, 260, 261, 5, 6, 0, 0, 261,
		29, 1, 0, 0, 0, 262, 263, 5, 18, 0, 0, 263, 264, 5, 5, 0, 0, 264, 265,
		3, 92, 46, 0, 265, 266, 3, 32, 16, 0, 266, 268, 5, 8, 0, 0, 267, 269, 3,
		14, 7, 0, 268, 267, 1, 0, 0, 0, 269, 270, 1, 0, 0, 0, 270, 268, 1, 0, 0,
		0, 270, 271, 1, 0, 0, 0, 271, 272, 1, 0, 0, 0, 272, 273, 5, 9, 0, 0, 273,
		31, 1, 0, 0, 0, 274, 275, 5, 6, 0, 0, 275, 33, 1, 0, 0, 0, 276, 280, 3,
		36, 18, 0, 277, 279, 3, 38, 19, 0, 278, 277, 1, 0, 0, 0, 279, 282, 1, 0,
		0, 0, 280, 278, 1, 0, 0, 0, 280, 281, 1, 0, 0, 0, 281, 284, 1, 0, 0, 0,
		282, 280, 1, 0, 0, 0, 283, 285, 3, 42, 21, 0, 284, 283, 1, 0, 0, 0, 284,
		285, 1, 0, 0, 0, 285, 35, 1, 0, 0, 0, 286, 287, 5, 19, 0, 0, 287, 288,
		5, 5, 0, 0, 288, 289, 3, 92, 46, 0, 289, 290, 3, 40, 20, 0, 290, 292, 5,
		8, 0, 0, 291, 293, 3, 14, 7, 0, 292, 291, 1, 0, 0, 0, 293, 294, 1, 0, 0,
		0, 294, 292, 1, 0, 0, 0, 294, 295, 1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 296,
		297, 5, 9, 0, 0, 297, 37, 1, 0, 0, 0, 298, 299, 5, 20, 0, 0, 299, 300,
		5, 19, 0, 0, 300, 301, 5, 5, 0, 0, 301, 302, 3, 92, 46, 0, 302, 303, 3,
		40, 20, 0, 303, 305, 5, 8, 0, 0, 304, 306, 3, 14, 7, 0, 305, 304, 1, 0,
		0, 0, 306, 307, 1, 0, 0, 0, 307, 305, 1, 0, 0, 0, 307, 308, 1, 0, 0, 0,
		308, 309, 1, 0, 0, 0, 309, 310, 5, 9, 0, 0, 310, 39, 1, 0, 0, 0, 311, 312,
		5, 6, 0, 0, 312, 41, 1, 0, 0, 0, 313, 314, 5, 20, 0, 0, 314, 316, 5, 8,
		0, 0, 315, 317, 3, 14, 7, 0, 316, 315, 1, 0, 0, 0, 317, 318, 1, 0, 0, 0,
		318, 316, 1, 0, 0, 0, 318, 319, 1, 0, 0, 0, 319, 320, 1, 0, 0, 0, 320,
		321, 5, 9, 0, 0, 321, 43, 1, 0, 0, 0, 322, 323, 5, 21, 0, 0, 323, 324,
		3, 92, 46, 0, 324, 45, 1, 0, 0, 0, 325, 348, 3, 48, 24, 0, 326, 348, 3,
		50, 25, 0, 327, 348, 3, 52, 26, 0, 328, 348, 3, 54, 27, 0, 329, 348, 3,
		56, 28, 0, 330, 348, 3, 58, 29, 0, 331, 348, 3, 60, 30, 0, 332, 348, 3,
		62, 31, 0, 333, 348, 3, 64, 32, 0, 334, 348, 3, 66, 33, 0, 335, 348, 3,
		68, 34, 0, 336, 348, 3, 70, 35, 0, 337, 348, 3, 72, 36, 0, 338, 348, 3,
		74, 37, 0, 339, 348, 3, 78, 39, 0, 340, 348, 3, 76, 38, 0, 341, 348, 3,
		80, 40, 0, 342, 348, 3, 82, 41, 0, 343, 348, 3, 84, 42, 0, 344, 348, 3,
		86, 43, 0, 345, 348, 3, 88, 44, 0, 346, 348, 3, 90, 45, 0, 347, 325, 1,
		0, 0, 0, 347, 326, 1, 0, 0, 0, 347, 327, 1, 0, 0, 0, 347, 328, 1, 0, 0,
		0, 347, 329, 1, 0, 0, 0, 347, 330, 1, 0, 0, 0, 347, 331, 1, 0, 0, 0, 347,
		332, 1, 0, 0, 0, 347, 333, 1, 0, 0, 0, 347, 334, 1, 0, 0, 0, 347, 335,
		1, 0, 0, 0, 347, 336, 1, 0, 0, 0, 347, 337, 1, 0, 0, 0, 347, 338, 1, 0,
		0, 0, 347, 339, 1, 0, 0, 0, 347, 340, 1, 0, 0, 0, 347, 341, 1, 0, 0, 0,
		347, 342, 1, 0, 0, 0, 347, 343, 1, 0, 0, 0, 347, 344, 1, 0, 0, 0, 347,
		345, 1, 0, 0, 0, 347, 346, 1, 0, 0, 0, 348, 47, 1, 0, 0, 0, 349, 350, 5,
		22, 0, 0, 350, 351, 5, 5, 0, 0, 351, 356, 5, 60, 0, 0, 352, 353, 5, 13,
		0, 0, 353, 355, 5, 60, 0, 0, 354, 352, 1, 0, 0, 0, 355, 358, 1, 0, 0, 0,
		356, 354, 1, 0, 0, 0, 356, 357, 1, 0, 0, 0, 357, 359, 1, 0, 0, 0, 358,
		356, 1, 0, 0, 0, 359, 360, 5, 6, 0, 0, 360, 49, 1, 0, 0, 0, 361, 362, 5,
		23, 0, 0, 362, 363, 5, 5, 0, 0, 363, 368, 3, 92, 46, 0, 364, 365, 5, 13,
		0, 0, 365, 367, 3, 92, 46, 0, 366, 364, 1, 0, 0, 0, 367, 370, 1, 0, 0,
		0, 368, 366, 1, 0, 0, 0, 368, 369, 1, 0, 0, 0, 369, 371, 1, 0, 0, 0, 370,
		368, 1, 0, 0, 0, 371, 372, 5, 6, 0, 0, 372, 51, 1, 0, 0, 0, 373, 374, 5,
		24, 0, 0, 374, 375, 5, 5, 0, 0, 375, 376, 3, 92, 46, 0, 376, 377, 5, 13,
		0, 0, 377, 378, 3, 92, 46, 0, 378, 379, 5, 6, 0, 0, 379, 53, 1, 0, 0, 0,
		380, 381, 5, 25, 0, 0, 381, 382, 5, 5, 0, 0, 382, 383, 3, 92, 46, 0, 383,
		384, 5, 6, 0, 0, 384, 55, 1, 0, 0, 0, 385, 386, 5, 26, 0, 0, 386, 387,
		5, 5, 0, 0, 387, 388, 3, 92, 46, 0, 388, 389, 5, 6, 0, 0, 389, 57, 1, 0,
		0, 0, 390, 391, 5, 27, 0, 0, 391, 392, 5, 5, 0, 0, 392, 395, 3, 92, 46,
		0, 393, 394, 5, 13, 0, 0, 394, 396, 3, 92, 46, 0, 395, 393, 1, 0, 0, 0,
		395, 396, 1, 0, 0, 0, 396, 397, 1, 0, 0, 0, 397, 398, 5, 6, 0, 0, 398,
		59, 1, 0, 0, 0, 399, 400, 5, 28, 0, 0, 400, 401, 5, 5, 0, 0, 401, 402,
		3, 92, 46, 0, 402, 403, 5, 6, 0, 0, 403, 61, 1, 0, 0, 0, 404, 405, 5, 29,
		0, 0, 405, 406, 5, 5, 0, 0, 406, 407, 3, 92, 46, 0, 407, 408, 5, 6, 0,
		0, 408, 63, 1, 0, 0, 0, 409, 410, 5, 30, 0, 0, 410, 411, 5, 5, 0, 0, 411,
		412, 3, 92, 46, 0, 412, 413, 5, 6, 0, 0, 413, 65, 1, 0, 0, 0, 414, 415,
		5, 31, 0, 0, 415, 416, 5, 5, 0, 0, 416, 417, 3, 92, 46, 0, 417, 418, 5,
		6, 0, 0, 418, 67, 1, 0, 0, 0, 419, 420, 5, 32, 0, 0, 420, 421, 5, 5, 0,
		0, 421, 422, 3, 92, 46, 0, 422, 423, 5, 6, 0, 0, 423, 69, 1, 0, 0, 0, 424,
		425, 5, 33, 0, 0, 425, 426, 5, 5, 0, 0, 426, 427, 3, 92, 46, 0, 427, 428,
		5, 6, 0, 0, 428, 71, 1, 0, 0, 0, 429, 430, 5, 34, 0, 0, 430, 431, 5, 5,
		0, 0, 431, 432, 3, 92, 46, 0, 432, 433, 5, 6, 0, 0, 433, 73, 1, 0, 0, 0,
		434, 435, 5, 35, 0, 0, 435, 436, 5, 5, 0, 0, 436, 437, 3, 92, 46, 0, 437,
		438, 5, 6, 0, 0, 438, 75, 1, 0, 0, 0, 439, 440, 5, 36, 0, 0, 440, 441,
		5, 5, 0, 0, 441, 442, 3, 92, 46, 0, 442, 443, 5, 6, 0, 0, 443, 77, 1, 0,
		0, 0, 444, 445, 5, 37, 0, 0, 445, 446, 5, 5, 0, 0, 446, 447, 3, 92, 46,
		0, 447, 448, 5, 6, 0, 0, 448, 79, 1, 0, 0, 0, 449, 450, 5, 38, 0, 0, 450,
		451, 5, 5, 0, 0, 451, 452, 3, 92, 46, 0, 452, 453, 5, 6, 0, 0, 453, 81,
		1, 0, 0, 0, 454, 455, 5, 39, 0, 0, 455, 456, 5, 5, 0, 0, 456, 457, 3, 92,
		46, 0, 457, 458, 5, 6, 0, 0, 458, 83, 1, 0, 0, 0, 459, 460, 5, 40, 0, 0,
		460, 461, 5, 5, 0, 0, 461, 462, 3, 92, 46, 0, 462, 463, 5, 6, 0, 0, 463,
		85, 1, 0, 0, 0, 464, 465, 5, 41, 0, 0, 465, 466, 5, 5, 0, 0, 466, 467,
		3, 92, 46, 0, 467, 468, 5, 6, 0, 0, 468, 87, 1, 0, 0, 0, 469, 470, 5, 42,
		0, 0, 470, 471, 5, 5, 0, 0, 471, 472, 3, 92, 46, 0, 472, 473, 5, 6, 0,
		0, 473, 89, 1, 0, 0, 0, 474, 475, 5, 43, 0, 0, 475, 476, 5, 5, 0, 0, 476,
		477, 3, 92, 46, 0, 477, 478, 5, 6, 0, 0, 478, 91, 1, 0, 0, 0, 479, 483,
		3, 94, 47, 0, 480, 481, 3, 114, 57, 0, 481, 482, 3, 94, 47, 0, 482, 484,
		1, 0, 0, 0, 483, 480, 1, 0, 0, 0, 483, 484, 1, 0, 0, 0, 484, 93, 1, 0,
		0, 0, 485, 491, 3, 96, 48, 0, 486, 487, 3, 110, 55, 0, 487, 488, 3, 96,
		48, 0, 488, 490, 1, 0, 0, 0, 489, 486, 1, 0, 0, 0, 490, 493, 1, 0, 0, 0,
		491, 489, 1, 0, 0, 0, 491, 492, 1, 0, 0, 0, 492, 95, 1, 0, 0, 0, 493, 491,
		1, 0, 0, 0, 494, 500, 3, 98, 49, 0, 495, 496, 3, 112, 56, 0, 496, 497,
		3, 98, 49, 0, 497, 499, 1, 0, 0, 0, 498, 495, 1, 0, 0, 0, 499, 502, 1,
		0, 0, 0, 500, 498, 1, 0, 0, 0, 500, 501, 1, 0, 0, 0, 501, 97, 1, 0, 0,
		0, 502, 500, 1, 0, 0, 0, 503, 505, 3, 100, 50, 0, 504, 503, 1, 0, 0, 0,
		504, 505, 1, 0, 0, 0, 505, 511, 1, 0, 0, 0, 506, 512, 3, 102, 51, 0, 507,
		512, 3, 106, 53, 0, 508, 512, 3, 46, 23, 0, 509, 512, 3, 104, 52, 0, 510,
		512, 3, 108, 54, 0, 511, 506, 1, 0, 0, 0, 511, 507, 1, 0, 0, 0, 511, 508,
		1, 0, 0, 0, 511, 509, 1, 0, 0, 0, 511, 510, 1, 0, 0, 0, 512, 99, 1, 0,
		0, 0, 513, 514, 5, 44, 0, 0, 514, 101, 1, 0, 0, 0, 515, 516, 5, 5, 0, 0,
		516, 517, 3, 92, 46, 0, 517, 518, 5, 6, 0, 0, 518, 103, 1, 0, 0, 0, 519,
		520, 5, 60, 0, 0, 520, 529, 5, 5, 0, 0, 521, 526, 3, 92, 46, 0, 522, 523,
		5, 13, 0, 0, 523, 525, 3, 92, 46, 0, 524, 522, 1, 0, 0, 0, 525, 528, 1,
		0, 0, 0, 526, 524, 1, 0, 0, 0, 526, 527, 1, 0, 0, 0, 527, 530, 1, 0, 0,
		0, 528, 526, 1, 0, 0, 0, 529, 521, 1, 0, 0, 0, 529, 530, 1, 0, 0, 0, 530,
		531, 1, 0, 0, 0, 531, 532, 5, 6, 0, 0, 532, 105, 1, 0, 0, 0, 533, 534,
		5, 60, 0, 0, 534, 535, 5, 12, 0, 0, 535, 536, 3, 92, 46, 0, 536, 541, 5,
		14, 0, 0, 537, 538, 5, 12, 0, 0, 538, 539, 3, 92, 46, 0, 539, 540, 5, 14,
		0, 0, 540, 542, 1, 0, 0, 0, 541, 537, 1, 0, 0, 0, 541, 542, 1, 0, 0, 0,
		542, 107, 1, 0, 0, 0, 543, 544, 7, 1, 0, 0, 544, 109, 1, 0, 0, 0, 545,
		546, 7, 2, 0, 0, 546, 111, 1, 0, 0, 0, 547, 548, 7, 3, 0, 0, 548, 113,
		1, 0, 0, 0, 549, 550, 7, 4, 0, 0, 550, 115, 1, 0, 0, 0, 34, 122, 128, 141,
		149, 158, 165, 183, 191, 199, 206, 214, 222, 227, 231, 240, 249, 270, 280,
		284, 294, 307, 318, 347, 356, 368, 395, 483, 491, 500, 504, 511, 526, 529,
		541,
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
	BistatParserT__43          = 44
	BistatParserT__44          = 45
	BistatParserT__45          = 46
	BistatParserT__46          = 47
	BistatParserT__47          = 48
	BistatParserT__48          = 49
	BistatParserT__49          = 50
	BistatParserT__50          = 51
	BistatParserT__51          = 52
	BistatParserT__52          = 53
	BistatParserT__53          = 54
	BistatParserT__54          = 55
	BistatParserT__55          = 56
	BistatParserWS             = 57
	BistatParserCARDINALITY    = 58
	BistatParserTYPE_PRIMITIVE = 59
	BistatParserID             = 60
	BistatParserINT_CONS       = 61
	BistatParserNUMBER         = 62
	BistatParserBOOL_CONS      = 63
	BistatParserSTRING_CONS    = 64
	BistatParserFLOAT_CONS     = 65
)

// BistatParser rules.
const (
	BistatParserRULE_program          = 0
	BistatParserRULE_varDeclaration   = 1
	BistatParserRULE_var_type         = 2
	BistatParserRULE_funcDef          = 3
	BistatParserRULE_funcEnd          = 4
	BistatParserRULE_paramDeclaration = 5
	BistatParserRULE_main             = 6
	BistatParserRULE_stmt             = 7
	BistatParserRULE_assignment       = 8
	BistatParserRULE_matrixAssignment = 9
	BistatParserRULE_listAssignment   = 10
	BistatParserRULE_comment          = 11
	BistatParserRULE_forLoop          = 12
	BistatParserRULE_forHeader        = 13
	BistatParserRULE_forExprEnd       = 14
	BistatParserRULE_whileLoop        = 15
	BistatParserRULE_whileExprEnd     = 16
	BistatParserRULE_conditional      = 17
	BistatParserRULE_ifStmt           = 18
	BistatParserRULE_elseIfStmt       = 19
	BistatParserRULE_condExprEnd      = 20
	BistatParserRULE_elseStmt         = 21
	BistatParserRULE_returnStmt       = 22
	BistatParserRULE_specialFunction  = 23
	BistatParserRULE_inputRead        = 24
	BistatParserRULE_print            = 25
	BistatParserRULE_listAdd          = 26
	BistatParserRULE_listPop          = 27
	BistatParserRULE_length           = 28
	BistatParserRULE_range            = 29
	BistatParserRULE_plot             = 30
	BistatParserRULE_sum              = 31
	BistatParserRULE_min              = 32
	BistatParserRULE_prod             = 33
	BistatParserRULE_avg              = 34
	BistatParserRULE_sMode            = 35
	BistatParserRULE_median           = 36
	BistatParserRULE_sin              = 37
	BistatParserRULE_tan              = 38
	BistatParserRULE_cos              = 39
	BistatParserRULE_sort             = 40
	BistatParserRULE_sqrt             = 41
	BistatParserRULE_floor            = 42
	BistatParserRULE_ceil             = 43
	BistatParserRULE_abs              = 44
	BistatParserRULE_not              = 45
	BistatParserRULE_expression       = 46
	BistatParserRULE_exp              = 47
	BistatParserRULE_term             = 48
	BistatParserRULE_factor           = 49
	BistatParserRULE_unaryMinus       = 50
	BistatParserRULE_nestedExpression = 51
	BistatParserRULE_functionCall     = 52
	BistatParserRULE_indexing         = 53
	BistatParserRULE_varCons          = 54
	BistatParserRULE_opSec            = 55
	BistatParserRULE_opFirst          = 56
	BistatParserRULE_logicOperator    = 57
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
		p.SetState(116)
		p.Match(BistatParserT__0)
	}
	{
		p.SetState(117)
		p.Match(BistatParserID)
	}
	{
		p.SetState(118)
		p.Match(BistatParserT__1)
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__2 {
		{
			p.SetState(119)
			p.VarDeclaration()
		}

		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__3 {
		{
			p.SetState(125)
			p.FuncDef()
		}

		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(131)
		p.Main()
	}
	{
		p.SetState(132)
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
		p.SetState(134)
		p.Match(BistatParserT__2)
	}
	{
		p.SetState(135)
		p.Var_type()
	}
	{
		p.SetState(136)
		p.Match(BistatParserID)
	}
	{
		p.SetState(137)
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
		p.SetState(139)
		p.Match(BistatParserTYPE_PRIMITIVE)
	}

	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BistatParserCARDINALITY {
		{
			p.SetState(140)
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
	FuncEnd() IFuncEndContext
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

func (s *FuncDefContext) FuncEnd() IFuncEndContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncEndContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncEndContext)
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
		p.SetState(143)
		p.Match(BistatParserT__3)
	}
	{
		p.SetState(144)
		p.Match(BistatParserID)
	}
	{
		p.SetState(145)
		p.Match(BistatParserT__4)
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__2 {
		{
			p.SetState(146)
			p.ParamDeclaration()
		}

		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(152)
		p.Match(BistatParserT__5)
	}
	{
		p.SetState(153)
		p.Match(BistatParserT__6)
	}
	{
		p.SetState(154)
		p.Var_type()
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__2 {
		{
			p.SetState(155)
			p.VarDeclaration()
		}

		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(161)
		p.Match(BistatParserT__7)
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1152939096791678976) != 0) {
		{
			p.SetState(162)
			p.Stmt()
		}

		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(167)
		p.FuncEnd()
	}

	return localctx
}

// IFuncEndContext is an interface to support dynamic dispatch.
type IFuncEndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFuncEndContext differentiates from other interfaces.
	IsFuncEndContext()
}

type FuncEndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncEndContext() *FuncEndContext {
	var p = new(FuncEndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_funcEnd
	return p
}

func (*FuncEndContext) IsFuncEndContext() {}

func NewFuncEndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncEndContext {
	var p = new(FuncEndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_funcEnd

	return p
}

func (s *FuncEndContext) GetParser() antlr.Parser { return s.parser }
func (s *FuncEndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncEndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncEndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterFuncEnd(s)
	}
}

func (s *FuncEndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitFuncEnd(s)
	}
}

func (p *BistatParser) FuncEnd() (localctx IFuncEndContext) {
	this := p
	_ = this

	localctx = NewFuncEndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BistatParserRULE_funcEnd)

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
		p.SetState(169)
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
		p.SetState(171)
		p.Match(BistatParserT__2)
	}
	{
		p.SetState(172)
		p.Var_type()
	}
	{
		p.SetState(173)
		p.Match(BistatParserID)
	}
	{
		p.SetState(174)
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
	p.EnterRule(localctx, 12, BistatParserRULE_main)
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
		p.SetState(176)
		p.Match(BistatParserT__9)
	}
	{
		p.SetState(177)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(178)
		p.Match(BistatParserT__5)
	}
	{
		p.SetState(179)
		p.Match(BistatParserT__7)
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1152939096791678976) != 0) {
		{
			p.SetState(180)
			p.Stmt()
		}

		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(185)
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
	p.EnterRule(localctx, 14, BistatParserRULE_stmt)

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

	p.SetState(199)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BistatParserT__20, BistatParserT__21, BistatParserT__22, BistatParserT__23, BistatParserT__24, BistatParserT__25, BistatParserT__26, BistatParserT__27, BistatParserT__28, BistatParserT__29, BistatParserT__30, BistatParserT__31, BistatParserT__32, BistatParserT__33, BistatParserT__34, BistatParserT__35, BistatParserT__36, BistatParserT__37, BistatParserT__38, BistatParserT__39, BistatParserT__40, BistatParserT__41, BistatParserT__42, BistatParserID:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(187)
				p.Assignment()
			}

		case 2:
			{
				p.SetState(188)
				p.SpecialFunction()
			}

		case 3:
			{
				p.SetState(189)
				p.FunctionCall()
			}

		case 4:
			{
				p.SetState(190)
				p.ReturnStmt()
			}

		}
		{
			p.SetState(193)
			p.Match(BistatParserT__1)
		}

	case BistatParserT__18:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(195)
			p.Conditional()
		}

	case BistatParserT__17:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(196)
			p.WhileLoop()
		}

	case BistatParserT__15:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(197)
			p.ForLoop()
		}

	case BistatParserT__14:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(198)
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
	Expression() IExpressionContext
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

func (s *AssignmentContext) Expression() IExpressionContext {
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
	p.EnterRule(localctx, 16, BistatParserRULE_assignment)

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
		p.Match(BistatParserID)
	}
	{
		p.SetState(202)
		p.Match(BistatParserT__10)
	}
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(203)
			p.Expression()
		}

	case 2:
		{
			p.SetState(204)
			p.ListAssignment()
		}

	case 3:
		{
			p.SetState(205)
			p.MatrixAssignment()
		}

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
		p.SetState(208)
		p.Match(BistatParserT__11)
	}
	{
		p.SetState(209)
		p.ListAssignment()
	}
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__12 {
		{
			p.SetState(210)
			p.Match(BistatParserT__12)
		}
		{
			p.SetState(211)
			p.ListAssignment()
		}

		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(217)
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
	AllVarCons() []IVarConsContext
	VarCons(i int) IVarConsContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

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

func (s *ListAssignmentContext) AllVarCons() []IVarConsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarConsContext); ok {
			len++
		}
	}

	tst := make([]IVarConsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarConsContext); ok {
			tst[i] = t.(IVarConsContext)
			i++
		}
	}

	return tst
}

func (s *ListAssignmentContext) VarCons(i int) IVarConsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarConsContext); ok {
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

	return t.(IVarConsContext)
}

func (s *ListAssignmentContext) AllExpression() []IExpressionContext {
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

func (s *ListAssignmentContext) Expression(i int) IExpressionContext {
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
		p.SetState(219)
		p.Match(BistatParserT__11)
	}
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(220)
			p.VarCons()
		}

	case 2:
		{
			p.SetState(221)
			p.Expression()
		}

	}
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__12 {
		{
			p.SetState(224)
			p.Match(BistatParserT__12)
		}
		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(225)
				p.VarCons()
			}

		case 2:
			{
				p.SetState(226)
				p.Expression()
			}

		}

		p.SetState(233)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(234)
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
		p.SetState(236)
		p.Match(BistatParserT__14)
	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-32770) != 0) || _la == BistatParserSTRING_CONS || _la == BistatParserFLOAT_CONS {
		{
			p.SetState(237)
			_la = p.GetTokenStream().LA(1)

			if _la <= 0 || _la == BistatParserT__14 {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(240)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(242)
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
	ForHeader() IForHeaderContext
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

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

func (s *ForLoopContext) ForHeader() IForHeaderContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForHeaderContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForHeaderContext)
}

func (s *ForLoopContext) AllStmt() []IStmtContext {
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

func (s *ForLoopContext) Stmt(i int) IStmtContext {
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
		p.SetState(244)
		p.ForHeader()
	}
	{
		p.SetState(245)
		p.Match(BistatParserT__7)
	}
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1152939096791678976) != 0) {
		{
			p.SetState(246)
			p.Stmt()
		}

		p.SetState(249)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(251)
		p.Match(BistatParserT__8)
	}

	return localctx
}

// IForHeaderContext is an interface to support dynamic dispatch.
type IForHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Expression() IExpressionContext

	// IsForHeaderContext differentiates from other interfaces.
	IsForHeaderContext()
}

type ForHeaderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForHeaderContext() *ForHeaderContext {
	var p = new(ForHeaderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_forHeader
	return p
}

func (*ForHeaderContext) IsForHeaderContext() {}

func NewForHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForHeaderContext {
	var p = new(ForHeaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_forHeader

	return p
}

func (s *ForHeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *ForHeaderContext) ID() antlr.TerminalNode {
	return s.GetToken(BistatParserID, 0)
}

func (s *ForHeaderContext) Expression() IExpressionContext {
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

func (s *ForHeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForHeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForHeaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterForHeader(s)
	}
}

func (s *ForHeaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitForHeader(s)
	}
}

func (p *BistatParser) ForHeader() (localctx IForHeaderContext) {
	this := p
	_ = this

	localctx = NewForHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, BistatParserRULE_forHeader)

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
		p.SetState(253)
		p.Match(BistatParserT__15)
	}
	{
		p.SetState(254)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(255)
		p.Match(BistatParserID)
	}
	{
		p.SetState(256)
		p.Match(BistatParserT__16)
	}
	{
		p.SetState(257)
		p.Expression()
	}
	{
		p.SetState(258)
		p.Match(BistatParserT__5)
	}

	return localctx
}

// IForExprEndContext is an interface to support dynamic dispatch.
type IForExprEndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsForExprEndContext differentiates from other interfaces.
	IsForExprEndContext()
}

type ForExprEndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForExprEndContext() *ForExprEndContext {
	var p = new(ForExprEndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_forExprEnd
	return p
}

func (*ForExprEndContext) IsForExprEndContext() {}

func NewForExprEndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForExprEndContext {
	var p = new(ForExprEndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_forExprEnd

	return p
}

func (s *ForExprEndContext) GetParser() antlr.Parser { return s.parser }
func (s *ForExprEndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForExprEndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForExprEndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterForExprEnd(s)
	}
}

func (s *ForExprEndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitForExprEnd(s)
	}
}

func (p *BistatParser) ForExprEnd() (localctx IForExprEndContext) {
	this := p
	_ = this

	localctx = NewForExprEndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, BistatParserRULE_forExprEnd)

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
		p.SetState(260)
		p.Match(BistatParserT__5)
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
	WhileExprEnd() IWhileExprEndContext
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

func (s *WhileLoopContext) WhileExprEnd() IWhileExprEndContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileExprEndContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileExprEndContext)
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
	p.EnterRule(localctx, 30, BistatParserRULE_whileLoop)
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
		p.SetState(262)
		p.Match(BistatParserT__17)
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
		p.WhileExprEnd()
	}
	{
		p.SetState(266)
		p.Match(BistatParserT__7)
	}
	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1152939096791678976) != 0) {
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

	return localctx
}

// IWhileExprEndContext is an interface to support dynamic dispatch.
type IWhileExprEndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsWhileExprEndContext differentiates from other interfaces.
	IsWhileExprEndContext()
}

type WhileExprEndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileExprEndContext() *WhileExprEndContext {
	var p = new(WhileExprEndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_whileExprEnd
	return p
}

func (*WhileExprEndContext) IsWhileExprEndContext() {}

func NewWhileExprEndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileExprEndContext {
	var p = new(WhileExprEndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_whileExprEnd

	return p
}

func (s *WhileExprEndContext) GetParser() antlr.Parser { return s.parser }
func (s *WhileExprEndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileExprEndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileExprEndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterWhileExprEnd(s)
	}
}

func (s *WhileExprEndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitWhileExprEnd(s)
	}
}

func (p *BistatParser) WhileExprEnd() (localctx IWhileExprEndContext) {
	this := p
	_ = this

	localctx = NewWhileExprEndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, BistatParserRULE_whileExprEnd)

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
		p.SetState(274)
		p.Match(BistatParserT__5)
	}

	return localctx
}

// IConditionalContext is an interface to support dynamic dispatch.
type IConditionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IfStmt() IIfStmtContext
	AllElseIfStmt() []IElseIfStmtContext
	ElseIfStmt(i int) IElseIfStmtContext
	ElseStmt() IElseStmtContext

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

func (s *ConditionalContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *ConditionalContext) AllElseIfStmt() []IElseIfStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElseIfStmtContext); ok {
			len++
		}
	}

	tst := make([]IElseIfStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElseIfStmtContext); ok {
			tst[i] = t.(IElseIfStmtContext)
			i++
		}
	}

	return tst
}

func (s *ConditionalContext) ElseIfStmt(i int) IElseIfStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseIfStmtContext); ok {
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

	return t.(IElseIfStmtContext)
}

func (s *ConditionalContext) ElseStmt() IElseStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseStmtContext)
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
	p.EnterRule(localctx, 34, BistatParserRULE_conditional)
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
		p.SetState(276)
		p.IfStmt()
	}
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(277)
				p.ElseIfStmt()
			}

		}
		p.SetState(282)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
	}
	p.SetState(284)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BistatParserT__19 {
		{
			p.SetState(283)
			p.ElseStmt()
		}

	}

	return localctx
}

// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	CondExprEnd() ICondExprEndContext
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_ifStmt
	return p
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_ifStmt

	return p
}

func (s *IfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtContext) Expression() IExpressionContext {
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

func (s *IfStmtContext) CondExprEnd() ICondExprEndContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICondExprEndContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICondExprEndContext)
}

func (s *IfStmtContext) AllStmt() []IStmtContext {
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

func (s *IfStmtContext) Stmt(i int) IStmtContext {
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

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (p *BistatParser) IfStmt() (localctx IIfStmtContext) {
	this := p
	_ = this

	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, BistatParserRULE_ifStmt)
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
		p.SetState(286)
		p.Match(BistatParserT__18)
	}
	{
		p.SetState(287)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(288)
		p.Expression()
	}
	{
		p.SetState(289)
		p.CondExprEnd()
	}
	{
		p.SetState(290)
		p.Match(BistatParserT__7)
	}
	p.SetState(292)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1152939096791678976) != 0) {
		{
			p.SetState(291)
			p.Stmt()
		}

		p.SetState(294)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(296)
		p.Match(BistatParserT__8)
	}

	return localctx
}

// IElseIfStmtContext is an interface to support dynamic dispatch.
type IElseIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	CondExprEnd() ICondExprEndContext
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsElseIfStmtContext differentiates from other interfaces.
	IsElseIfStmtContext()
}

type ElseIfStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseIfStmtContext() *ElseIfStmtContext {
	var p = new(ElseIfStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_elseIfStmt
	return p
}

func (*ElseIfStmtContext) IsElseIfStmtContext() {}

func NewElseIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseIfStmtContext {
	var p = new(ElseIfStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_elseIfStmt

	return p
}

func (s *ElseIfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseIfStmtContext) Expression() IExpressionContext {
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

func (s *ElseIfStmtContext) CondExprEnd() ICondExprEndContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICondExprEndContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICondExprEndContext)
}

func (s *ElseIfStmtContext) AllStmt() []IStmtContext {
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

func (s *ElseIfStmtContext) Stmt(i int) IStmtContext {
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

func (s *ElseIfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseIfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseIfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterElseIfStmt(s)
	}
}

func (s *ElseIfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitElseIfStmt(s)
	}
}

func (p *BistatParser) ElseIfStmt() (localctx IElseIfStmtContext) {
	this := p
	_ = this

	localctx = NewElseIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, BistatParserRULE_elseIfStmt)
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
		p.SetState(298)
		p.Match(BistatParserT__19)
	}
	{
		p.SetState(299)
		p.Match(BistatParserT__18)
	}
	{
		p.SetState(300)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(301)
		p.Expression()
	}
	{
		p.SetState(302)
		p.CondExprEnd()
	}
	{
		p.SetState(303)
		p.Match(BistatParserT__7)
	}
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1152939096791678976) != 0) {
		{
			p.SetState(304)
			p.Stmt()
		}

		p.SetState(307)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(309)
		p.Match(BistatParserT__8)
	}

	return localctx
}

// ICondExprEndContext is an interface to support dynamic dispatch.
type ICondExprEndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsCondExprEndContext differentiates from other interfaces.
	IsCondExprEndContext()
}

type CondExprEndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCondExprEndContext() *CondExprEndContext {
	var p = new(CondExprEndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_condExprEnd
	return p
}

func (*CondExprEndContext) IsCondExprEndContext() {}

func NewCondExprEndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CondExprEndContext {
	var p = new(CondExprEndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_condExprEnd

	return p
}

func (s *CondExprEndContext) GetParser() antlr.Parser { return s.parser }
func (s *CondExprEndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondExprEndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CondExprEndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterCondExprEnd(s)
	}
}

func (s *CondExprEndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitCondExprEnd(s)
	}
}

func (p *BistatParser) CondExprEnd() (localctx ICondExprEndContext) {
	this := p
	_ = this

	localctx = NewCondExprEndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, BistatParserRULE_condExprEnd)

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
		p.SetState(311)
		p.Match(BistatParserT__5)
	}

	return localctx
}

// IElseStmtContext is an interface to support dynamic dispatch.
type IElseStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsElseStmtContext differentiates from other interfaces.
	IsElseStmtContext()
}

type ElseStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseStmtContext() *ElseStmtContext {
	var p = new(ElseStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_elseStmt
	return p
}

func (*ElseStmtContext) IsElseStmtContext() {}

func NewElseStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseStmtContext {
	var p = new(ElseStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_elseStmt

	return p
}

func (s *ElseStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseStmtContext) AllStmt() []IStmtContext {
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

func (s *ElseStmtContext) Stmt(i int) IStmtContext {
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

func (s *ElseStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterElseStmt(s)
	}
}

func (s *ElseStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitElseStmt(s)
	}
}

func (p *BistatParser) ElseStmt() (localctx IElseStmtContext) {
	this := p
	_ = this

	localctx = NewElseStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, BistatParserRULE_elseStmt)
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
		p.SetState(313)
		p.Match(BistatParserT__19)
	}
	{
		p.SetState(314)
		p.Match(BistatParserT__7)
	}
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1152939096791678976) != 0) {
		{
			p.SetState(315)
			p.Stmt()
		}

		p.SetState(318)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(320)
		p.Match(BistatParserT__8)
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
	p.EnterRule(localctx, 44, BistatParserRULE_returnStmt)

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
		p.SetState(322)
		p.Match(BistatParserT__20)
	}
	{
		p.SetState(323)
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
	p.EnterRule(localctx, 46, BistatParserRULE_specialFunction)

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

	p.SetState(347)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BistatParserT__21:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(325)
			p.InputRead()
		}

	case BistatParserT__22:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(326)
			p.Print_()
		}

	case BistatParserT__23:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(327)
			p.ListAdd()
		}

	case BistatParserT__24:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(328)
			p.ListPop()
		}

	case BistatParserT__25:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(329)
			p.Length()
		}

	case BistatParserT__26:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(330)
			p.Range_()
		}

	case BistatParserT__27:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(331)
			p.Plot()
		}

	case BistatParserT__28:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(332)
			p.Sum()
		}

	case BistatParserT__29:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(333)
			p.Min()
		}

	case BistatParserT__30:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(334)
			p.Prod()
		}

	case BistatParserT__31:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(335)
			p.Avg()
		}

	case BistatParserT__32:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(336)
			p.SMode()
		}

	case BistatParserT__33:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(337)
			p.Median()
		}

	case BistatParserT__34:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(338)
			p.Sin()
		}

	case BistatParserT__36:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(339)
			p.Cos()
		}

	case BistatParserT__35:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(340)
			p.Tan()
		}

	case BistatParserT__37:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(341)
			p.Sort()
		}

	case BistatParserT__38:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(342)
			p.Sqrt()
		}

	case BistatParserT__39:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(343)
			p.Floor()
		}

	case BistatParserT__40:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(344)
			p.Ceil()
		}

	case BistatParserT__41:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(345)
			p.Abs()
		}

	case BistatParserT__42:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(346)
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
	p.EnterRule(localctx, 48, BistatParserRULE_inputRead)
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
		p.SetState(349)
		p.Match(BistatParserT__21)
	}
	{
		p.SetState(350)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(351)
		p.Match(BistatParserID)
	}
	p.SetState(356)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__12 {
		{
			p.SetState(352)
			p.Match(BistatParserT__12)
		}
		{
			p.SetState(353)
			p.Match(BistatParserID)
		}

		p.SetState(358)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(359)
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
	p.EnterRule(localctx, 50, BistatParserRULE_print)
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
		p.SetState(361)
		p.Match(BistatParserT__22)
	}
	{
		p.SetState(362)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(363)
		p.Expression()
	}
	p.SetState(368)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__12 {
		{
			p.SetState(364)
			p.Match(BistatParserT__12)
		}
		{
			p.SetState(365)
			p.Expression()
		}

		p.SetState(370)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(371)
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
	p.EnterRule(localctx, 52, BistatParserRULE_listAdd)

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
		p.SetState(373)
		p.Match(BistatParserT__23)
	}
	{
		p.SetState(374)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(375)
		p.Expression()
	}
	{
		p.SetState(376)
		p.Match(BistatParserT__12)
	}
	{
		p.SetState(377)
		p.Expression()
	}
	{
		p.SetState(378)
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
	p.EnterRule(localctx, 54, BistatParserRULE_listPop)

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
		p.SetState(380)
		p.Match(BistatParserT__24)
	}
	{
		p.SetState(381)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(382)
		p.Expression()
	}
	{
		p.SetState(383)
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
	p.EnterRule(localctx, 56, BistatParserRULE_length)

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
		p.SetState(385)
		p.Match(BistatParserT__25)
	}
	{
		p.SetState(386)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(387)
		p.Expression()
	}
	{
		p.SetState(388)
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
	p.EnterRule(localctx, 58, BistatParserRULE_range)
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
		p.SetState(390)
		p.Match(BistatParserT__26)
	}
	{
		p.SetState(391)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(392)
		p.Expression()
	}
	p.SetState(395)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BistatParserT__12 {
		{
			p.SetState(393)
			p.Match(BistatParserT__12)
		}
		{
			p.SetState(394)
			p.Expression()
		}

	}
	{
		p.SetState(397)
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
	p.EnterRule(localctx, 60, BistatParserRULE_plot)

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
		p.Match(BistatParserT__27)
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
	p.EnterRule(localctx, 62, BistatParserRULE_sum)

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
		p.Match(BistatParserT__28)
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
	p.EnterRule(localctx, 64, BistatParserRULE_min)

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
		p.Match(BistatParserT__29)
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
	p.EnterRule(localctx, 66, BistatParserRULE_prod)

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
		p.Match(BistatParserT__30)
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
	p.EnterRule(localctx, 68, BistatParserRULE_avg)

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
		p.Match(BistatParserT__31)
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
	p.EnterRule(localctx, 70, BistatParserRULE_sMode)

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
		p.Match(BistatParserT__32)
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
	p.EnterRule(localctx, 72, BistatParserRULE_median)

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
		p.Match(BistatParserT__33)
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
	p.EnterRule(localctx, 74, BistatParserRULE_sin)

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
		p.Match(BistatParserT__34)
	}
	{
		p.SetState(435)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(436)
		p.Expression()
	}
	{
		p.SetState(437)
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
	p.EnterRule(localctx, 76, BistatParserRULE_tan)

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
		p.Match(BistatParserT__35)
	}
	{
		p.SetState(440)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(441)
		p.Expression()
	}
	{
		p.SetState(442)
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
	p.EnterRule(localctx, 78, BistatParserRULE_cos)

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
		p.Match(BistatParserT__36)
	}
	{
		p.SetState(445)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(446)
		p.Expression()
	}
	{
		p.SetState(447)
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
	p.EnterRule(localctx, 80, BistatParserRULE_sort)

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
		p.SetState(449)
		p.Match(BistatParserT__37)
	}
	{
		p.SetState(450)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(451)
		p.Expression()
	}
	{
		p.SetState(452)
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
	p.EnterRule(localctx, 82, BistatParserRULE_sqrt)

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
		p.SetState(454)
		p.Match(BistatParserT__38)
	}
	{
		p.SetState(455)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(456)
		p.Expression()
	}
	{
		p.SetState(457)
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
	p.EnterRule(localctx, 84, BistatParserRULE_floor)

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
		p.SetState(459)
		p.Match(BistatParserT__39)
	}
	{
		p.SetState(460)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(461)
		p.Expression()
	}
	{
		p.SetState(462)
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
	p.EnterRule(localctx, 86, BistatParserRULE_ceil)

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
		p.SetState(464)
		p.Match(BistatParserT__40)
	}
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
	p.EnterRule(localctx, 88, BistatParserRULE_abs)

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
		p.SetState(469)
		p.Match(BistatParserT__41)
	}
	{
		p.SetState(470)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(471)
		p.Expression()
	}
	{
		p.SetState(472)
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
	p.EnterRule(localctx, 90, BistatParserRULE_not)

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
		p.SetState(474)
		p.Match(BistatParserT__42)
	}
	{
		p.SetState(475)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(476)
		p.Expression()
	}
	{
		p.SetState(477)
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
	LogicOperator() ILogicOperatorContext

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

func (s *ExpressionContext) LogicOperator() ILogicOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicOperatorContext)
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
	p.EnterRule(localctx, 92, BistatParserRULE_expression)
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
		p.SetState(479)
		p.Exp()
	}
	p.SetState(483)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&143552238122434560) != 0 {
		{
			p.SetState(480)
			p.LogicOperator()
		}
		{
			p.SetState(481)
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
	AllTerm() []ITermContext
	Term(i int) ITermContext
	AllOpSec() []IOpSecContext
	OpSec(i int) IOpSecContext

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

func (s *ExpContext) AllTerm() []ITermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContext); ok {
			len++
		}
	}

	tst := make([]ITermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContext); ok {
			tst[i] = t.(ITermContext)
			i++
		}
	}

	return tst
}

func (s *ExpContext) Term(i int) ITermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
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

	return t.(ITermContext)
}

func (s *ExpContext) AllOpSec() []IOpSecContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOpSecContext); ok {
			len++
		}
	}

	tst := make([]IOpSecContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOpSecContext); ok {
			tst[i] = t.(IOpSecContext)
			i++
		}
	}

	return tst
}

func (s *ExpContext) OpSec(i int) IOpSecContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpSecContext); ok {
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

	return t.(IOpSecContext)
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
	p.EnterRule(localctx, 94, BistatParserRULE_exp)
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
		p.SetState(485)
		p.Term()
	}
	p.SetState(491)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__43 || _la == BistatParserT__44 {
		{
			p.SetState(486)
			p.OpSec()
		}
		{
			p.SetState(487)
			p.Term()
		}

		p.SetState(493)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFactor() []IFactorContext
	Factor(i int) IFactorContext
	AllOpFirst() []IOpFirstContext
	OpFirst(i int) IOpFirstContext

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

func (s *TermContext) AllFactor() []IFactorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFactorContext); ok {
			len++
		}
	}

	tst := make([]IFactorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFactorContext); ok {
			tst[i] = t.(IFactorContext)
			i++
		}
	}

	return tst
}

func (s *TermContext) Factor(i int) IFactorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
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

	return t.(IFactorContext)
}

func (s *TermContext) AllOpFirst() []IOpFirstContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOpFirstContext); ok {
			len++
		}
	}

	tst := make([]IOpFirstContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOpFirstContext); ok {
			tst[i] = t.(IOpFirstContext)
			i++
		}
	}

	return tst
}

func (s *TermContext) OpFirst(i int) IOpFirstContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpFirstContext); ok {
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

	return t.(IOpFirstContext)
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
	p.EnterRule(localctx, 96, BistatParserRULE_term)
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
		p.SetState(494)
		p.Factor()
	}
	p.SetState(500)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&492581209243648) != 0 {
		{
			p.SetState(495)
			p.OpFirst()
		}
		{
			p.SetState(496)
			p.Factor()
		}

		p.SetState(502)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NestedExpression() INestedExpressionContext
	Indexing() IIndexingContext
	SpecialFunction() ISpecialFunctionContext
	FunctionCall() IFunctionCallContext
	VarCons() IVarConsContext
	UnaryMinus() IUnaryMinusContext

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

func (s *FactorContext) NestedExpression() INestedExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INestedExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INestedExpressionContext)
}

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

func (s *FactorContext) VarCons() IVarConsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarConsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarConsContext)
}

func (s *FactorContext) UnaryMinus() IUnaryMinusContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryMinusContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryMinusContext)
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
	p.EnterRule(localctx, 98, BistatParserRULE_factor)
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
	p.SetState(504)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BistatParserT__43 {
		{
			p.SetState(503)
			p.UnaryMinus()
		}

	}
	p.SetState(511)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(506)
			p.NestedExpression()
		}

	case 2:
		{
			p.SetState(507)
			p.Indexing()
		}

	case 3:
		{
			p.SetState(508)
			p.SpecialFunction()
		}

	case 4:
		{
			p.SetState(509)
			p.FunctionCall()
		}

	case 5:
		{
			p.SetState(510)
			p.VarCons()
		}

	}

	return localctx
}

// IUnaryMinusContext is an interface to support dynamic dispatch.
type IUnaryMinusContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsUnaryMinusContext differentiates from other interfaces.
	IsUnaryMinusContext()
}

type UnaryMinusContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryMinusContext() *UnaryMinusContext {
	var p = new(UnaryMinusContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_unaryMinus
	return p
}

func (*UnaryMinusContext) IsUnaryMinusContext() {}

func NewUnaryMinusContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryMinusContext {
	var p = new(UnaryMinusContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_unaryMinus

	return p
}

func (s *UnaryMinusContext) GetParser() antlr.Parser { return s.parser }
func (s *UnaryMinusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryMinusContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryMinusContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterUnaryMinus(s)
	}
}

func (s *UnaryMinusContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitUnaryMinus(s)
	}
}

func (p *BistatParser) UnaryMinus() (localctx IUnaryMinusContext) {
	this := p
	_ = this

	localctx = NewUnaryMinusContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, BistatParserRULE_unaryMinus)

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
		p.SetState(513)
		p.Match(BistatParserT__43)
	}

	return localctx
}

// INestedExpressionContext is an interface to support dynamic dispatch.
type INestedExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsNestedExpressionContext differentiates from other interfaces.
	IsNestedExpressionContext()
}

type NestedExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNestedExpressionContext() *NestedExpressionContext {
	var p = new(NestedExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_nestedExpression
	return p
}

func (*NestedExpressionContext) IsNestedExpressionContext() {}

func NewNestedExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NestedExpressionContext {
	var p = new(NestedExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_nestedExpression

	return p
}

func (s *NestedExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *NestedExpressionContext) Expression() IExpressionContext {
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

func (s *NestedExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestedExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NestedExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterNestedExpression(s)
	}
}

func (s *NestedExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitNestedExpression(s)
	}
}

func (p *BistatParser) NestedExpression() (localctx INestedExpressionContext) {
	this := p
	_ = this

	localctx = NewNestedExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, BistatParserRULE_nestedExpression)

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
		p.SetState(515)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(516)
		p.Expression()
	}
	{
		p.SetState(517)
		p.Match(BistatParserT__5)
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
	p.EnterRule(localctx, 104, BistatParserRULE_functionCall)
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
		p.SetState(519)
		p.Match(BistatParserID)
	}
	{
		p.SetState(520)
		p.Match(BistatParserT__4)
	}
	p.SetState(529)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-5)) & ^0x3f) == 0 && ((int64(1)<<(_la-5))&2125700123630370817) != 0 {
		{
			p.SetState(521)
			p.Expression()
		}
		p.SetState(526)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == BistatParserT__12 {
			{
				p.SetState(522)
				p.Match(BistatParserT__12)
			}
			{
				p.SetState(523)
				p.Expression()
			}

			p.SetState(528)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(531)
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
	p.EnterRule(localctx, 106, BistatParserRULE_indexing)
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
		p.SetState(533)
		p.Match(BistatParserID)
	}
	{
		p.SetState(534)
		p.Match(BistatParserT__11)
	}
	{
		p.SetState(535)
		p.Expression()
	}
	{
		p.SetState(536)
		p.Match(BistatParserT__13)
	}
	p.SetState(541)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BistatParserT__11 {
		{
			p.SetState(537)
			p.Match(BistatParserT__11)
		}
		{
			p.SetState(538)
			p.Expression()
		}
		{
			p.SetState(539)
			p.Match(BistatParserT__13)
		}

	}

	return localctx
}

// IVarConsContext is an interface to support dynamic dispatch.
type IVarConsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING_CONS() antlr.TerminalNode
	FLOAT_CONS() antlr.TerminalNode
	INT_CONS() antlr.TerminalNode
	BOOL_CONS() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsVarConsContext differentiates from other interfaces.
	IsVarConsContext()
}

type VarConsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarConsContext() *VarConsContext {
	var p = new(VarConsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_varCons
	return p
}

func (*VarConsContext) IsVarConsContext() {}

func NewVarConsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarConsContext {
	var p = new(VarConsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_varCons

	return p
}

func (s *VarConsContext) GetParser() antlr.Parser { return s.parser }

func (s *VarConsContext) STRING_CONS() antlr.TerminalNode {
	return s.GetToken(BistatParserSTRING_CONS, 0)
}

func (s *VarConsContext) FLOAT_CONS() antlr.TerminalNode {
	return s.GetToken(BistatParserFLOAT_CONS, 0)
}

func (s *VarConsContext) INT_CONS() antlr.TerminalNode {
	return s.GetToken(BistatParserINT_CONS, 0)
}

func (s *VarConsContext) BOOL_CONS() antlr.TerminalNode {
	return s.GetToken(BistatParserBOOL_CONS, 0)
}

func (s *VarConsContext) ID() antlr.TerminalNode {
	return s.GetToken(BistatParserID, 0)
}

func (s *VarConsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarConsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarConsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterVarCons(s)
	}
}

func (s *VarConsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitVarCons(s)
	}
}

func (p *BistatParser) VarCons() (localctx IVarConsContext) {
	this := p
	_ = this

	localctx = NewVarConsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, BistatParserRULE_varCons)
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
		p.SetState(543)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-60)) & ^0x3f) == 0 && ((int64(1)<<(_la-60))&59) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOpSecContext is an interface to support dynamic dispatch.
type IOpSecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOpSecContext differentiates from other interfaces.
	IsOpSecContext()
}

type OpSecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpSecContext() *OpSecContext {
	var p = new(OpSecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_opSec
	return p
}

func (*OpSecContext) IsOpSecContext() {}

func NewOpSecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpSecContext {
	var p = new(OpSecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_opSec

	return p
}

func (s *OpSecContext) GetParser() antlr.Parser { return s.parser }
func (s *OpSecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpSecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpSecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterOpSec(s)
	}
}

func (s *OpSecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitOpSec(s)
	}
}

func (p *BistatParser) OpSec() (localctx IOpSecContext) {
	this := p
	_ = this

	localctx = NewOpSecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, BistatParserRULE_opSec)
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
		p.SetState(545)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BistatParserT__43 || _la == BistatParserT__44) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOpFirstContext is an interface to support dynamic dispatch.
type IOpFirstContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOpFirstContext differentiates from other interfaces.
	IsOpFirstContext()
}

type OpFirstContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpFirstContext() *OpFirstContext {
	var p = new(OpFirstContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_opFirst
	return p
}

func (*OpFirstContext) IsOpFirstContext() {}

func NewOpFirstContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpFirstContext {
	var p = new(OpFirstContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_opFirst

	return p
}

func (s *OpFirstContext) GetParser() antlr.Parser { return s.parser }
func (s *OpFirstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpFirstContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpFirstContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterOpFirst(s)
	}
}

func (s *OpFirstContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitOpFirst(s)
	}
}

func (p *BistatParser) OpFirst() (localctx IOpFirstContext) {
	this := p
	_ = this

	localctx = NewOpFirstContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, BistatParserRULE_opFirst)
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
		p.SetState(547)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&492581209243648) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILogicOperatorContext is an interface to support dynamic dispatch.
type ILogicOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLogicOperatorContext differentiates from other interfaces.
	IsLogicOperatorContext()
}

type LogicOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicOperatorContext() *LogicOperatorContext {
	var p = new(LogicOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_logicOperator
	return p
}

func (*LogicOperatorContext) IsLogicOperatorContext() {}

func NewLogicOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicOperatorContext {
	var p = new(LogicOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_logicOperator

	return p
}

func (s *LogicOperatorContext) GetParser() antlr.Parser { return s.parser }
func (s *LogicOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterLogicOperator(s)
	}
}

func (s *LogicOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitLogicOperator(s)
	}
}

func (p *BistatParser) LogicOperator() (localctx ILogicOperatorContext) {
	this := p
	_ = this

	localctx = NewLogicOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, BistatParserRULE_logicOperator)
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
		p.SetState(549)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&143552238122434560) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
