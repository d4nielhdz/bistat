// Code generated from parser/Bistat.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Bistat

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type BistatParser struct {
	*antlr.BaseParser
}

var BistatParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func bistatParserInit() {
	staticData := &BistatParserStaticData
	staticData.LiteralNames = []string{
		"", "'Program'", "';'", "'var'", "'func'", "'('", "')'", "':'", "'{'",
		"'}'", "'main'", "'='", "'['", "','", "']'", "'#'", "'for'", "'in'",
		"'while'", "'if'", "'else'", "'return'", "'read'", "'print'", "'plot'",
		"'sum'", "'min'", "'max'", "'prod'", "'avg'", "'sMode'", "'median'",
		"'sin'", "'tan'", "'cos'", "'sqrt'", "'floor'", "'ceil'", "'abs'", "'not'",
		"'-'", "'+'", "'/'", "'*'", "'%'", "'<'", "'>'", "'<='", "'>='", "'=='",
		"'!='", "'&&'", "'||'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "WS", "CARDINALITY", "TYPE_PRIMITIVE", "BOOL_CONS", "ID", "INT_CONS",
		"NUMBER", "STRING_CONS", "FLOAT_CONS",
	}
	staticData.RuleNames = []string{
		"program", "varDeclaration", "var_type", "funcDef", "funcBlockStart",
		"funcBlockEnd", "paramDeclaration", "main", "stmt", "assignment", "indexing",
		"matrixAssignment", "listAssignment", "comment", "forLoop", "forHeader",
		"forExprEnd", "whileLoop", "whileExprEnd", "conditional", "ifStmt",
		"elseIfStmt", "condExprEnd", "elseStmt", "returnStmt", "specialFunction",
		"inputRead", "print", "plot", "sum", "min", "max", "prod", "avg", "sMode",
		"median", "sin", "tan", "cos", "sqrt", "floor", "ceil", "abs", "not",
		"expression", "exp", "term", "factor", "unaryMinus", "nestedExpression",
		"functionCall", "varCons", "opSec", "opFirst", "logicOperator",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 61, 519, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 115, 8, 0,
		10, 0, 12, 0, 118, 9, 0, 1, 0, 5, 0, 121, 8, 0, 10, 0, 12, 0, 124, 9, 0,
		1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 3, 2, 136,
		8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 142, 8, 3, 10, 3, 12, 3, 145, 9, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 151, 8, 3, 10, 3, 12, 3, 154, 9, 3, 1, 3,
		1, 3, 4, 3, 158, 8, 3, 11, 3, 12, 3, 159, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 4, 7,
		178, 8, 7, 11, 7, 12, 7, 179, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8,
		188, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 196, 8, 8, 1, 9, 1,
		9, 3, 9, 200, 8, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 206, 8, 9, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 217, 8, 10,
		1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 223, 8, 11, 10, 11, 12, 11, 226, 9,
		11, 1, 11, 3, 11, 229, 8, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12,
		5, 12, 237, 8, 12, 10, 12, 12, 12, 240, 9, 12, 1, 12, 1, 12, 1, 13, 1,
		13, 4, 13, 246, 8, 13, 11, 13, 12, 13, 247, 1, 13, 1, 13, 1, 14, 1, 14,
		1, 14, 4, 14, 255, 8, 14, 11, 14, 12, 14, 256, 1, 14, 1, 14, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 4, 17, 276, 8, 17, 11, 17, 12, 17, 277, 1, 17, 1,
		17, 1, 18, 1, 18, 1, 19, 1, 19, 5, 19, 286, 8, 19, 10, 19, 12, 19, 289,
		9, 19, 1, 19, 3, 19, 292, 8, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 4, 20, 300, 8, 20, 11, 20, 12, 20, 301, 1, 20, 1, 20, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 4, 21, 313, 8, 21, 11, 21, 12, 21, 314,
		1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 4, 23, 324, 8, 23, 11,
		23, 12, 23, 325, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 351, 8, 25, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 5, 26, 358, 8, 26, 10, 26, 12, 26, 361, 9, 26, 1, 26, 1,
		26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 5, 27, 370, 8, 27, 10, 27, 12, 27,
		373, 9, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1,
		33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35,
		1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39,
		1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1,
		41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43,
		1, 44, 1, 44, 1, 44, 1, 44, 3, 44, 461, 8, 44, 1, 45, 1, 45, 1, 45, 1,
		45, 5, 45, 467, 8, 45, 10, 45, 12, 45, 470, 9, 45, 1, 46, 1, 46, 1, 46,
		1, 46, 5, 46, 476, 8, 46, 10, 46, 12, 46, 479, 9, 46, 1, 47, 3, 47, 482,
		8, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 3, 47, 489, 8, 47, 1, 48, 1,
		48, 1, 49, 1, 49, 1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 5, 50,
		502, 8, 50, 10, 50, 12, 50, 505, 9, 50, 3, 50, 507, 8, 50, 1, 50, 1, 50,
		1, 51, 1, 51, 1, 52, 1, 52, 1, 53, 1, 53, 1, 54, 1, 54, 1, 54, 0, 0, 55,
		0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
		38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72,
		74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106,
		108, 0, 5, 1, 0, 15, 15, 2, 0, 56, 58, 60, 61, 1, 0, 40, 41, 1, 0, 42,
		44, 1, 0, 45, 52, 521, 0, 110, 1, 0, 0, 0, 2, 128, 1, 0, 0, 0, 4, 133,
		1, 0, 0, 0, 6, 137, 1, 0, 0, 0, 8, 163, 1, 0, 0, 0, 10, 165, 1, 0, 0, 0,
		12, 167, 1, 0, 0, 0, 14, 172, 1, 0, 0, 0, 16, 195, 1, 0, 0, 0, 18, 199,
		1, 0, 0, 0, 20, 207, 1, 0, 0, 0, 22, 218, 1, 0, 0, 0, 24, 232, 1, 0, 0,
		0, 26, 243, 1, 0, 0, 0, 28, 251, 1, 0, 0, 0, 30, 260, 1, 0, 0, 0, 32, 267,
		1, 0, 0, 0, 34, 269, 1, 0, 0, 0, 36, 281, 1, 0, 0, 0, 38, 283, 1, 0, 0,
		0, 40, 293, 1, 0, 0, 0, 42, 305, 1, 0, 0, 0, 44, 318, 1, 0, 0, 0, 46, 320,
		1, 0, 0, 0, 48, 329, 1, 0, 0, 0, 50, 350, 1, 0, 0, 0, 52, 352, 1, 0, 0,
		0, 54, 364, 1, 0, 0, 0, 56, 376, 1, 0, 0, 0, 58, 381, 1, 0, 0, 0, 60, 386,
		1, 0, 0, 0, 62, 391, 1, 0, 0, 0, 64, 396, 1, 0, 0, 0, 66, 401, 1, 0, 0,
		0, 68, 406, 1, 0, 0, 0, 70, 411, 1, 0, 0, 0, 72, 416, 1, 0, 0, 0, 74, 421,
		1, 0, 0, 0, 76, 426, 1, 0, 0, 0, 78, 431, 1, 0, 0, 0, 80, 436, 1, 0, 0,
		0, 82, 441, 1, 0, 0, 0, 84, 446, 1, 0, 0, 0, 86, 451, 1, 0, 0, 0, 88, 456,
		1, 0, 0, 0, 90, 462, 1, 0, 0, 0, 92, 471, 1, 0, 0, 0, 94, 481, 1, 0, 0,
		0, 96, 490, 1, 0, 0, 0, 98, 492, 1, 0, 0, 0, 100, 496, 1, 0, 0, 0, 102,
		510, 1, 0, 0, 0, 104, 512, 1, 0, 0, 0, 106, 514, 1, 0, 0, 0, 108, 516,
		1, 0, 0, 0, 110, 111, 5, 1, 0, 0, 111, 112, 5, 57, 0, 0, 112, 116, 5, 2,
		0, 0, 113, 115, 3, 2, 1, 0, 114, 113, 1, 0, 0, 0, 115, 118, 1, 0, 0, 0,
		116, 114, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 122, 1, 0, 0, 0, 118,
		116, 1, 0, 0, 0, 119, 121, 3, 6, 3, 0, 120, 119, 1, 0, 0, 0, 121, 124,
		1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 125, 1, 0,
		0, 0, 124, 122, 1, 0, 0, 0, 125, 126, 3, 14, 7, 0, 126, 127, 5, 0, 0, 1,
		127, 1, 1, 0, 0, 0, 128, 129, 5, 3, 0, 0, 129, 130, 3, 4, 2, 0, 130, 131,
		5, 57, 0, 0, 131, 132, 5, 2, 0, 0, 132, 3, 1, 0, 0, 0, 133, 135, 5, 55,
		0, 0, 134, 136, 5, 54, 0, 0, 135, 134, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0,
		136, 5, 1, 0, 0, 0, 137, 138, 5, 4, 0, 0, 138, 139, 5, 57, 0, 0, 139, 143,
		5, 5, 0, 0, 140, 142, 3, 12, 6, 0, 141, 140, 1, 0, 0, 0, 142, 145, 1, 0,
		0, 0, 143, 141, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 146, 1, 0, 0, 0,
		145, 143, 1, 0, 0, 0, 146, 147, 5, 6, 0, 0, 147, 148, 5, 7, 0, 0, 148,
		152, 5, 55, 0, 0, 149, 151, 3, 2, 1, 0, 150, 149, 1, 0, 0, 0, 151, 154,
		1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 155, 1, 0,
		0, 0, 154, 152, 1, 0, 0, 0, 155, 157, 3, 8, 4, 0, 156, 158, 3, 16, 8, 0,
		157, 156, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 159,
		160, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 162, 3, 10, 5, 0, 162, 7, 1,
		0, 0, 0, 163, 164, 5, 8, 0, 0, 164, 9, 1, 0, 0, 0, 165, 166, 5, 9, 0, 0,
		166, 11, 1, 0, 0, 0, 167, 168, 5, 3, 0, 0, 168, 169, 5, 55, 0, 0, 169,
		170, 5, 57, 0, 0, 170, 171, 5, 2, 0, 0, 171, 13, 1, 0, 0, 0, 172, 173,
		5, 10, 0, 0, 173, 174, 5, 5, 0, 0, 174, 175, 5, 6, 0, 0, 175, 177, 5, 8,
		0, 0, 176, 178, 3, 16, 8, 0, 177, 176, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0,
		179, 177, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181,
		182, 5, 9, 0, 0, 182, 15, 1, 0, 0, 0, 183, 188, 3, 18, 9, 0, 184, 188,
		3, 50, 25, 0, 185, 188, 3, 100, 50, 0, 186, 188, 3, 48, 24, 0, 187, 183,
		1, 0, 0, 0, 187, 184, 1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 187, 186, 1, 0,
		0, 0, 188, 189, 1, 0, 0, 0, 189, 190, 5, 2, 0, 0, 190, 196, 1, 0, 0, 0,
		191, 196, 3, 38, 19, 0, 192, 196, 3, 34, 17, 0, 193, 196, 3, 28, 14, 0,
		194, 196, 3, 26, 13, 0, 195, 187, 1, 0, 0, 0, 195, 191, 1, 0, 0, 0, 195,
		192, 1, 0, 0, 0, 195, 193, 1, 0, 0, 0, 195, 194, 1, 0, 0, 0, 196, 17, 1,
		0, 0, 0, 197, 200, 3, 20, 10, 0, 198, 200, 5, 57, 0, 0, 199, 197, 1, 0,
		0, 0, 199, 198, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 205, 5, 11, 0, 0,
		202, 206, 3, 88, 44, 0, 203, 206, 3, 24, 12, 0, 204, 206, 3, 22, 11, 0,
		205, 202, 1, 0, 0, 0, 205, 203, 1, 0, 0, 0, 205, 204, 1, 0, 0, 0, 206,
		19, 1, 0, 0, 0, 207, 208, 5, 57, 0, 0, 208, 209, 5, 8, 0, 0, 209, 210,
		3, 88, 44, 0, 210, 211, 5, 9, 0, 0, 211, 216, 1, 0, 0, 0, 212, 213, 5,
		8, 0, 0, 213, 214, 3, 88, 44, 0, 214, 215, 5, 9, 0, 0, 215, 217, 1, 0,
		0, 0, 216, 212, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217, 21, 1, 0, 0, 0,
		218, 219, 5, 12, 0, 0, 219, 224, 3, 24, 12, 0, 220, 221, 5, 13, 0, 0, 221,
		223, 3, 24, 12, 0, 222, 220, 1, 0, 0, 0, 223, 226, 1, 0, 0, 0, 224, 222,
		1, 0, 0, 0, 224, 225, 1, 0, 0, 0, 225, 228, 1, 0, 0, 0, 226, 224, 1, 0,
		0, 0, 227, 229, 5, 13, 0, 0, 228, 227, 1, 0, 0, 0, 228, 229, 1, 0, 0, 0,
		229, 230, 1, 0, 0, 0, 230, 231, 5, 14, 0, 0, 231, 23, 1, 0, 0, 0, 232,
		233, 5, 12, 0, 0, 233, 238, 3, 88, 44, 0, 234, 235, 5, 13, 0, 0, 235, 237,
		3, 88, 44, 0, 236, 234, 1, 0, 0, 0, 237, 240, 1, 0, 0, 0, 238, 236, 1,
		0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 241, 1, 0, 0, 0, 240, 238, 1, 0, 0,
		0, 241, 242, 5, 14, 0, 0, 242, 25, 1, 0, 0, 0, 243, 245, 5, 15, 0, 0, 244,
		246, 8, 0, 0, 0, 245, 244, 1, 0, 0, 0, 246, 247, 1, 0, 0, 0, 247, 245,
		1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 250, 5, 15,
		0, 0, 250, 27, 1, 0, 0, 0, 251, 252, 3, 30, 15, 0, 252, 254, 5, 8, 0, 0,
		253, 255, 3, 16, 8, 0, 254, 253, 1, 0, 0, 0, 255, 256, 1, 0, 0, 0, 256,
		254, 1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257, 258, 1, 0, 0, 0, 258, 259,
		5, 9, 0, 0, 259, 29, 1, 0, 0, 0, 260, 261, 5, 16, 0, 0, 261, 262, 5, 5,
		0, 0, 262, 263, 5, 57, 0, 0, 263, 264, 5, 17, 0, 0, 264, 265, 3, 88, 44,
		0, 265, 266, 5, 6, 0, 0, 266, 31, 1, 0, 0, 0, 267, 268, 5, 6, 0, 0, 268,
		33, 1, 0, 0, 0, 269, 270, 5, 18, 0, 0, 270, 271, 5, 5, 0, 0, 271, 272,
		3, 88, 44, 0, 272, 273, 3, 36, 18, 0, 273, 275, 5, 8, 0, 0, 274, 276, 3,
		16, 8, 0, 275, 274, 1, 0, 0, 0, 276, 277, 1, 0, 0, 0, 277, 275, 1, 0, 0,
		0, 277, 278, 1, 0, 0, 0, 278, 279, 1, 0, 0, 0, 279, 280, 5, 9, 0, 0, 280,
		35, 1, 0, 0, 0, 281, 282, 5, 6, 0, 0, 282, 37, 1, 0, 0, 0, 283, 287, 3,
		40, 20, 0, 284, 286, 3, 42, 21, 0, 285, 284, 1, 0, 0, 0, 286, 289, 1, 0,
		0, 0, 287, 285, 1, 0, 0, 0, 287, 288, 1, 0, 0, 0, 288, 291, 1, 0, 0, 0,
		289, 287, 1, 0, 0, 0, 290, 292, 3, 46, 23, 0, 291, 290, 1, 0, 0, 0, 291,
		292, 1, 0, 0, 0, 292, 39, 1, 0, 0, 0, 293, 294, 5, 19, 0, 0, 294, 295,
		5, 5, 0, 0, 295, 296, 3, 88, 44, 0, 296, 297, 3, 44, 22, 0, 297, 299, 5,
		8, 0, 0, 298, 300, 3, 16, 8, 0, 299, 298, 1, 0, 0, 0, 300, 301, 1, 0, 0,
		0, 301, 299, 1, 0, 0, 0, 301, 302, 1, 0, 0, 0, 302, 303, 1, 0, 0, 0, 303,
		304, 5, 9, 0, 0, 304, 41, 1, 0, 0, 0, 305, 306, 5, 20, 0, 0, 306, 307,
		5, 19, 0, 0, 307, 308, 5, 5, 0, 0, 308, 309, 3, 88, 44, 0, 309, 310, 3,
		44, 22, 0, 310, 312, 5, 8, 0, 0, 311, 313, 3, 16, 8, 0, 312, 311, 1, 0,
		0, 0, 313, 314, 1, 0, 0, 0, 314, 312, 1, 0, 0, 0, 314, 315, 1, 0, 0, 0,
		315, 316, 1, 0, 0, 0, 316, 317, 5, 9, 0, 0, 317, 43, 1, 0, 0, 0, 318, 319,
		5, 6, 0, 0, 319, 45, 1, 0, 0, 0, 320, 321, 5, 20, 0, 0, 321, 323, 5, 8,
		0, 0, 322, 324, 3, 16, 8, 0, 323, 322, 1, 0, 0, 0, 324, 325, 1, 0, 0, 0,
		325, 323, 1, 0, 0, 0, 325, 326, 1, 0, 0, 0, 326, 327, 1, 0, 0, 0, 327,
		328, 5, 9, 0, 0, 328, 47, 1, 0, 0, 0, 329, 330, 5, 21, 0, 0, 330, 331,
		3, 88, 44, 0, 331, 49, 1, 0, 0, 0, 332, 351, 3, 52, 26, 0, 333, 351, 3,
		54, 27, 0, 334, 351, 3, 56, 28, 0, 335, 351, 3, 58, 29, 0, 336, 351, 3,
		62, 31, 0, 337, 351, 3, 60, 30, 0, 338, 351, 3, 64, 32, 0, 339, 351, 3,
		66, 33, 0, 340, 351, 3, 68, 34, 0, 341, 351, 3, 70, 35, 0, 342, 351, 3,
		72, 36, 0, 343, 351, 3, 76, 38, 0, 344, 351, 3, 74, 37, 0, 345, 351, 3,
		78, 39, 0, 346, 351, 3, 80, 40, 0, 347, 351, 3, 82, 41, 0, 348, 351, 3,
		84, 42, 0, 349, 351, 3, 86, 43, 0, 350, 332, 1, 0, 0, 0, 350, 333, 1, 0,
		0, 0, 350, 334, 1, 0, 0, 0, 350, 335, 1, 0, 0, 0, 350, 336, 1, 0, 0, 0,
		350, 337, 1, 0, 0, 0, 350, 338, 1, 0, 0, 0, 350, 339, 1, 0, 0, 0, 350,
		340, 1, 0, 0, 0, 350, 341, 1, 0, 0, 0, 350, 342, 1, 0, 0, 0, 350, 343,
		1, 0, 0, 0, 350, 344, 1, 0, 0, 0, 350, 345, 1, 0, 0, 0, 350, 346, 1, 0,
		0, 0, 350, 347, 1, 0, 0, 0, 350, 348, 1, 0, 0, 0, 350, 349, 1, 0, 0, 0,
		351, 51, 1, 0, 0, 0, 352, 353, 5, 22, 0, 0, 353, 354, 5, 5, 0, 0, 354,
		359, 3, 88, 44, 0, 355, 356, 5, 13, 0, 0, 356, 358, 3, 88, 44, 0, 357,
		355, 1, 0, 0, 0, 358, 361, 1, 0, 0, 0, 359, 357, 1, 0, 0, 0, 359, 360,
		1, 0, 0, 0, 360, 362, 1, 0, 0, 0, 361, 359, 1, 0, 0, 0, 362, 363, 5, 6,
		0, 0, 363, 53, 1, 0, 0, 0, 364, 365, 5, 23, 0, 0, 365, 366, 5, 5, 0, 0,
		366, 371, 3, 88, 44, 0, 367, 368, 5, 13, 0, 0, 368, 370, 3, 88, 44, 0,
		369, 367, 1, 0, 0, 0, 370, 373, 1, 0, 0, 0, 371, 369, 1, 0, 0, 0, 371,
		372, 1, 0, 0, 0, 372, 374, 1, 0, 0, 0, 373, 371, 1, 0, 0, 0, 374, 375,
		5, 6, 0, 0, 375, 55, 1, 0, 0, 0, 376, 377, 5, 24, 0, 0, 377, 378, 5, 5,
		0, 0, 378, 379, 3, 88, 44, 0, 379, 380, 5, 6, 0, 0, 380, 57, 1, 0, 0, 0,
		381, 382, 5, 25, 0, 0, 382, 383, 5, 5, 0, 0, 383, 384, 3, 88, 44, 0, 384,
		385, 5, 6, 0, 0, 385, 59, 1, 0, 0, 0, 386, 387, 5, 26, 0, 0, 387, 388,
		5, 5, 0, 0, 388, 389, 3, 88, 44, 0, 389, 390, 5, 6, 0, 0, 390, 61, 1, 0,
		0, 0, 391, 392, 5, 27, 0, 0, 392, 393, 5, 5, 0, 0, 393, 394, 3, 88, 44,
		0, 394, 395, 5, 6, 0, 0, 395, 63, 1, 0, 0, 0, 396, 397, 5, 28, 0, 0, 397,
		398, 5, 5, 0, 0, 398, 399, 3, 88, 44, 0, 399, 400, 5, 6, 0, 0, 400, 65,
		1, 0, 0, 0, 401, 402, 5, 29, 0, 0, 402, 403, 5, 5, 0, 0, 403, 404, 3, 88,
		44, 0, 404, 405, 5, 6, 0, 0, 405, 67, 1, 0, 0, 0, 406, 407, 5, 30, 0, 0,
		407, 408, 5, 5, 0, 0, 408, 409, 3, 88, 44, 0, 409, 410, 5, 6, 0, 0, 410,
		69, 1, 0, 0, 0, 411, 412, 5, 31, 0, 0, 412, 413, 5, 5, 0, 0, 413, 414,
		3, 88, 44, 0, 414, 415, 5, 6, 0, 0, 415, 71, 1, 0, 0, 0, 416, 417, 5, 32,
		0, 0, 417, 418, 5, 5, 0, 0, 418, 419, 3, 88, 44, 0, 419, 420, 5, 6, 0,
		0, 420, 73, 1, 0, 0, 0, 421, 422, 5, 33, 0, 0, 422, 423, 5, 5, 0, 0, 423,
		424, 3, 88, 44, 0, 424, 425, 5, 6, 0, 0, 425, 75, 1, 0, 0, 0, 426, 427,
		5, 34, 0, 0, 427, 428, 5, 5, 0, 0, 428, 429, 3, 88, 44, 0, 429, 430, 5,
		6, 0, 0, 430, 77, 1, 0, 0, 0, 431, 432, 5, 35, 0, 0, 432, 433, 5, 5, 0,
		0, 433, 434, 3, 88, 44, 0, 434, 435, 5, 6, 0, 0, 435, 79, 1, 0, 0, 0, 436,
		437, 5, 36, 0, 0, 437, 438, 5, 5, 0, 0, 438, 439, 3, 88, 44, 0, 439, 440,
		5, 6, 0, 0, 440, 81, 1, 0, 0, 0, 441, 442, 5, 37, 0, 0, 442, 443, 5, 5,
		0, 0, 443, 444, 3, 88, 44, 0, 444, 445, 5, 6, 0, 0, 445, 83, 1, 0, 0, 0,
		446, 447, 5, 38, 0, 0, 447, 448, 5, 5, 0, 0, 448, 449, 3, 88, 44, 0, 449,
		450, 5, 6, 0, 0, 450, 85, 1, 0, 0, 0, 451, 452, 5, 39, 0, 0, 452, 453,
		5, 5, 0, 0, 453, 454, 3, 88, 44, 0, 454, 455, 5, 6, 0, 0, 455, 87, 1, 0,
		0, 0, 456, 460, 3, 90, 45, 0, 457, 458, 3, 108, 54, 0, 458, 459, 3, 90,
		45, 0, 459, 461, 1, 0, 0, 0, 460, 457, 1, 0, 0, 0, 460, 461, 1, 0, 0, 0,
		461, 89, 1, 0, 0, 0, 462, 468, 3, 92, 46, 0, 463, 464, 3, 104, 52, 0, 464,
		465, 3, 92, 46, 0, 465, 467, 1, 0, 0, 0, 466, 463, 1, 0, 0, 0, 467, 470,
		1, 0, 0, 0, 468, 466, 1, 0, 0, 0, 468, 469, 1, 0, 0, 0, 469, 91, 1, 0,
		0, 0, 470, 468, 1, 0, 0, 0, 471, 477, 3, 94, 47, 0, 472, 473, 3, 106, 53,
		0, 473, 474, 3, 94, 47, 0, 474, 476, 1, 0, 0, 0, 475, 472, 1, 0, 0, 0,
		476, 479, 1, 0, 0, 0, 477, 475, 1, 0, 0, 0, 477, 478, 1, 0, 0, 0, 478,
		93, 1, 0, 0, 0, 479, 477, 1, 0, 0, 0, 480, 482, 3, 96, 48, 0, 481, 480,
		1, 0, 0, 0, 481, 482, 1, 0, 0, 0, 482, 488, 1, 0, 0, 0, 483, 489, 3, 98,
		49, 0, 484, 489, 3, 50, 25, 0, 485, 489, 3, 100, 50, 0, 486, 489, 3, 20,
		10, 0, 487, 489, 3, 102, 51, 0, 488, 483, 1, 0, 0, 0, 488, 484, 1, 0, 0,
		0, 488, 485, 1, 0, 0, 0, 488, 486, 1, 0, 0, 0, 488, 487, 1, 0, 0, 0, 489,
		95, 1, 0, 0, 0, 490, 491, 5, 40, 0, 0, 491, 97, 1, 0, 0, 0, 492, 493, 5,
		5, 0, 0, 493, 494, 3, 88, 44, 0, 494, 495, 5, 6, 0, 0, 495, 99, 1, 0, 0,
		0, 496, 497, 5, 57, 0, 0, 497, 506, 5, 5, 0, 0, 498, 503, 3, 88, 44, 0,
		499, 500, 5, 13, 0, 0, 500, 502, 3, 88, 44, 0, 501, 499, 1, 0, 0, 0, 502,
		505, 1, 0, 0, 0, 503, 501, 1, 0, 0, 0, 503, 504, 1, 0, 0, 0, 504, 507,
		1, 0, 0, 0, 505, 503, 1, 0, 0, 0, 506, 498, 1, 0, 0, 0, 506, 507, 1, 0,
		0, 0, 507, 508, 1, 0, 0, 0, 508, 509, 5, 6, 0, 0, 509, 101, 1, 0, 0, 0,
		510, 511, 7, 1, 0, 0, 511, 103, 1, 0, 0, 0, 512, 513, 7, 2, 0, 0, 513,
		105, 1, 0, 0, 0, 514, 515, 7, 3, 0, 0, 515, 107, 1, 0, 0, 0, 516, 517,
		7, 4, 0, 0, 517, 109, 1, 0, 0, 0, 33, 116, 122, 135, 143, 152, 159, 179,
		187, 195, 199, 205, 216, 224, 228, 238, 247, 256, 277, 287, 291, 301, 314,
		325, 350, 359, 371, 460, 468, 477, 481, 488, 503, 506,
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
	staticData := &BistatParserStaticData
	staticData.once.Do(bistatParserInit)
}

// NewBistatParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBistatParser(input antlr.TokenStream) *BistatParser {
	BistatParserInit()
	this := new(BistatParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &BistatParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
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
	BistatParserWS             = 53
	BistatParserCARDINALITY    = 54
	BistatParserTYPE_PRIMITIVE = 55
	BistatParserBOOL_CONS      = 56
	BistatParserID             = 57
	BistatParserINT_CONS       = 58
	BistatParserNUMBER         = 59
	BistatParserSTRING_CONS    = 60
	BistatParserFLOAT_CONS     = 61
)

// BistatParser rules.
const (
	BistatParserRULE_program          = 0
	BistatParserRULE_varDeclaration   = 1
	BistatParserRULE_var_type         = 2
	BistatParserRULE_funcDef          = 3
	BistatParserRULE_funcBlockStart   = 4
	BistatParserRULE_funcBlockEnd     = 5
	BistatParserRULE_paramDeclaration = 6
	BistatParserRULE_main             = 7
	BistatParserRULE_stmt             = 8
	BistatParserRULE_assignment       = 9
	BistatParserRULE_indexing         = 10
	BistatParserRULE_matrixAssignment = 11
	BistatParserRULE_listAssignment   = 12
	BistatParserRULE_comment          = 13
	BistatParserRULE_forLoop          = 14
	BistatParserRULE_forHeader        = 15
	BistatParserRULE_forExprEnd       = 16
	BistatParserRULE_whileLoop        = 17
	BistatParserRULE_whileExprEnd     = 18
	BistatParserRULE_conditional      = 19
	BistatParserRULE_ifStmt           = 20
	BistatParserRULE_elseIfStmt       = 21
	BistatParserRULE_condExprEnd      = 22
	BistatParserRULE_elseStmt         = 23
	BistatParserRULE_returnStmt       = 24
	BistatParserRULE_specialFunction  = 25
	BistatParserRULE_inputRead        = 26
	BistatParserRULE_print            = 27
	BistatParserRULE_plot             = 28
	BistatParserRULE_sum              = 29
	BistatParserRULE_min              = 30
	BistatParserRULE_max              = 31
	BistatParserRULE_prod             = 32
	BistatParserRULE_avg              = 33
	BistatParserRULE_sMode            = 34
	BistatParserRULE_median           = 35
	BistatParserRULE_sin              = 36
	BistatParserRULE_tan              = 37
	BistatParserRULE_cos              = 38
	BistatParserRULE_sqrt             = 39
	BistatParserRULE_floor            = 40
	BistatParserRULE_ceil             = 41
	BistatParserRULE_abs              = 42
	BistatParserRULE_not              = 43
	BistatParserRULE_expression       = 44
	BistatParserRULE_exp              = 45
	BistatParserRULE_term             = 46
	BistatParserRULE_factor           = 47
	BistatParserRULE_unaryMinus       = 48
	BistatParserRULE_nestedExpression = 49
	BistatParserRULE_functionCall     = 50
	BistatParserRULE_varCons          = 51
	BistatParserRULE_opSec            = 52
	BistatParserRULE_opFirst          = 53
	BistatParserRULE_logicOperator    = 54
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BistatParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		p.Match(BistatParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(111)
		p.Match(BistatParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(112)
		p.Match(BistatParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__2 {
		{
			p.SetState(113)
			p.VarDeclaration()
		}

		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__3 {
		{
			p.SetState(119)
			p.FuncDef()
		}

		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(125)
		p.Main()
	}
	{
		p.SetState(126)
		p.Match(BistatParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclarationContext() *VarDeclarationContext {
	var p = new(VarDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_varDeclaration
	return p
}

func InitEmptyVarDeclarationContext(p *VarDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_varDeclaration
}

func (*VarDeclarationContext) IsVarDeclarationContext() {}

func NewVarDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclarationContext {
	var p = new(VarDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewVarDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BistatParserRULE_varDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Match(BistatParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)
		p.Var_type()
	}
	{
		p.SetState(130)
		p.Match(BistatParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(131)
		p.Match(BistatParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_typeContext() *Var_typeContext {
	var p = new(Var_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_var_type
	return p
}

func InitEmptyVar_typeContext(p *Var_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_var_type
}

func (*Var_typeContext) IsVar_typeContext() {}

func NewVar_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_typeContext {
	var p = new(Var_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewVar_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BistatParserRULE_var_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		p.Match(BistatParserTYPE_PRIMITIVE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BistatParserCARDINALITY {
		{
			p.SetState(134)
			p.Match(BistatParserCARDINALITY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncDefContext is an interface to support dynamic dispatch.
type IFuncDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	TYPE_PRIMITIVE() antlr.TerminalNode
	FuncBlockStart() IFuncBlockStartContext
	FuncBlockEnd() IFuncBlockEndContext
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDefContext() *FuncDefContext {
	var p = new(FuncDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_funcDef
	return p
}

func InitEmptyFuncDefContext(p *FuncDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_funcDef
}

func (*FuncDefContext) IsFuncDefContext() {}

func NewFuncDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDefContext {
	var p = new(FuncDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_funcDef

	return p
}

func (s *FuncDefContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDefContext) ID() antlr.TerminalNode {
	return s.GetToken(BistatParserID, 0)
}

func (s *FuncDefContext) TYPE_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(BistatParserTYPE_PRIMITIVE, 0)
}

func (s *FuncDefContext) FuncBlockStart() IFuncBlockStartContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncBlockStartContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncBlockStartContext)
}

func (s *FuncDefContext) FuncBlockEnd() IFuncBlockEndContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncBlockEndContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncBlockEndContext)
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
	localctx = NewFuncDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BistatParserRULE_funcDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(137)
		p.Match(BistatParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(138)
		p.Match(BistatParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)
		p.Match(BistatParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__2 {
		{
			p.SetState(140)
			p.ParamDeclaration()
		}

		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(146)
		p.Match(BistatParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(147)
		p.Match(BistatParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(148)
		p.Match(BistatParserTYPE_PRIMITIVE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__2 {
		{
			p.SetState(149)
			p.VarDeclaration()
		}

		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(155)
		p.FuncBlockStart()
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&144116287586271232) != 0) {
		{
			p.SetState(156)
			p.Stmt()
		}

		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(161)
		p.FuncBlockEnd()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncBlockStartContext is an interface to support dynamic dispatch.
type IFuncBlockStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFuncBlockStartContext differentiates from other interfaces.
	IsFuncBlockStartContext()
}

type FuncBlockStartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncBlockStartContext() *FuncBlockStartContext {
	var p = new(FuncBlockStartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_funcBlockStart
	return p
}

func InitEmptyFuncBlockStartContext(p *FuncBlockStartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_funcBlockStart
}

func (*FuncBlockStartContext) IsFuncBlockStartContext() {}

func NewFuncBlockStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncBlockStartContext {
	var p = new(FuncBlockStartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_funcBlockStart

	return p
}

func (s *FuncBlockStartContext) GetParser() antlr.Parser { return s.parser }
func (s *FuncBlockStartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncBlockStartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncBlockStartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterFuncBlockStart(s)
	}
}

func (s *FuncBlockStartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitFuncBlockStart(s)
	}
}

func (p *BistatParser) FuncBlockStart() (localctx IFuncBlockStartContext) {
	localctx = NewFuncBlockStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BistatParserRULE_funcBlockStart)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		p.Match(BistatParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncBlockEndContext is an interface to support dynamic dispatch.
type IFuncBlockEndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFuncBlockEndContext differentiates from other interfaces.
	IsFuncBlockEndContext()
}

type FuncBlockEndContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncBlockEndContext() *FuncBlockEndContext {
	var p = new(FuncBlockEndContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_funcBlockEnd
	return p
}

func InitEmptyFuncBlockEndContext(p *FuncBlockEndContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_funcBlockEnd
}

func (*FuncBlockEndContext) IsFuncBlockEndContext() {}

func NewFuncBlockEndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncBlockEndContext {
	var p = new(FuncBlockEndContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_funcBlockEnd

	return p
}

func (s *FuncBlockEndContext) GetParser() antlr.Parser { return s.parser }
func (s *FuncBlockEndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncBlockEndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncBlockEndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterFuncBlockEnd(s)
	}
}

func (s *FuncBlockEndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitFuncBlockEnd(s)
	}
}

func (p *BistatParser) FuncBlockEnd() (localctx IFuncBlockEndContext) {
	localctx = NewFuncBlockEndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BistatParserRULE_funcBlockEnd)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		p.Match(BistatParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamDeclarationContext is an interface to support dynamic dispatch.
type IParamDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TYPE_PRIMITIVE() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsParamDeclarationContext differentiates from other interfaces.
	IsParamDeclarationContext()
}

type ParamDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamDeclarationContext() *ParamDeclarationContext {
	var p = new(ParamDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_paramDeclaration
	return p
}

func InitEmptyParamDeclarationContext(p *ParamDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_paramDeclaration
}

func (*ParamDeclarationContext) IsParamDeclarationContext() {}

func NewParamDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamDeclarationContext {
	var p = new(ParamDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_paramDeclaration

	return p
}

func (s *ParamDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamDeclarationContext) TYPE_PRIMITIVE() antlr.TerminalNode {
	return s.GetToken(BistatParserTYPE_PRIMITIVE, 0)
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
	localctx = NewParamDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BistatParserRULE_paramDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.Match(BistatParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(168)
		p.Match(BistatParserTYPE_PRIMITIVE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
		p.Match(BistatParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
		p.Match(BistatParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMainContext() *MainContext {
	var p = new(MainContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_main
	return p
}

func InitEmptyMainContext(p *MainContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_main
}

func (*MainContext) IsMainContext() {}

func NewMainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainContext {
	var p = new(MainContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewMainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BistatParserRULE_main)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
		p.Match(BistatParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(173)
		p.Match(BistatParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(174)
		p.Match(BistatParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(175)
		p.Match(BistatParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&144116287586271232) != 0) {
		{
			p.SetState(176)
			p.Stmt()
		}

		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(181)
		p.Match(BistatParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, BistatParserRULE_stmt)
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BistatParserT__20, BistatParserT__21, BistatParserT__22, BistatParserT__23, BistatParserT__24, BistatParserT__25, BistatParserT__26, BistatParserT__27, BistatParserT__28, BistatParserT__29, BistatParserT__30, BistatParserT__31, BistatParserT__32, BistatParserT__33, BistatParserT__34, BistatParserT__35, BistatParserT__36, BistatParserT__37, BistatParserT__38, BistatParserID:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(183)
				p.Assignment()
			}

		case 2:
			{
				p.SetState(184)
				p.SpecialFunction()
			}

		case 3:
			{
				p.SetState(185)
				p.FunctionCall()
			}

		case 4:
			{
				p.SetState(186)
				p.ReturnStmt()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}
		{
			p.SetState(189)
			p.Match(BistatParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case BistatParserT__18:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(191)
			p.Conditional()
		}

	case BistatParserT__17:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(192)
			p.WhileLoop()
		}

	case BistatParserT__15:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(193)
			p.ForLoop()
		}

	case BistatParserT__14:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(194)
			p.Comment()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Indexing() IIndexingContext
	ID() antlr.TerminalNode
	Expression() IExpressionContext
	ListAssignment() IListAssignmentContext
	MatrixAssignment() IMatrixAssignmentContext

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Indexing() IIndexingContext {
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
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, BistatParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(197)
			p.Indexing()
		}

	case 2:
		{
			p.SetState(198)
			p.Match(BistatParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(201)
		p.Match(BistatParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(202)
			p.Expression()
		}

	case 2:
		{
			p.SetState(203)
			p.ListAssignment()
		}

	case 3:
		{
			p.SetState(204)
			p.MatrixAssignment()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexingContext() *IndexingContext {
	var p = new(IndexingContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_indexing
	return p
}

func InitEmptyIndexingContext(p *IndexingContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_indexing
}

func (*IndexingContext) IsIndexingContext() {}

func NewIndexingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexingContext {
	var p = new(IndexingContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewIndexingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, BistatParserRULE_indexing)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(207)
		p.Match(BistatParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(208)
		p.Match(BistatParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.Expression()
	}
	{
		p.SetState(210)
		p.Match(BistatParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BistatParserT__7 {
		{
			p.SetState(212)
			p.Match(BistatParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(213)
			p.Expression()
		}
		{
			p.SetState(214)
			p.Match(BistatParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatrixAssignmentContext() *MatrixAssignmentContext {
	var p = new(MatrixAssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_matrixAssignment
	return p
}

func InitEmptyMatrixAssignmentContext(p *MatrixAssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_matrixAssignment
}

func (*MatrixAssignmentContext) IsMatrixAssignmentContext() {}

func NewMatrixAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatrixAssignmentContext {
	var p = new(MatrixAssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewMatrixAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, BistatParserRULE_matrixAssignment)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(218)
		p.Match(BistatParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(219)
		p.ListAssignment()
	}
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(220)
				p.Match(BistatParserT__12)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(221)
				p.ListAssignment()
			}

		}
		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BistatParserT__12 {
		{
			p.SetState(227)
			p.Match(BistatParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(230)
		p.Match(BistatParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListAssignmentContext is an interface to support dynamic dispatch.
type IListAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsListAssignmentContext differentiates from other interfaces.
	IsListAssignmentContext()
}

type ListAssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListAssignmentContext() *ListAssignmentContext {
	var p = new(ListAssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_listAssignment
	return p
}

func InitEmptyListAssignmentContext(p *ListAssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_listAssignment
}

func (*ListAssignmentContext) IsListAssignmentContext() {}

func NewListAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListAssignmentContext {
	var p = new(ListAssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_listAssignment

	return p
}

func (s *ListAssignmentContext) GetParser() antlr.Parser { return s.parser }

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
	localctx = NewListAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, BistatParserRULE_listAssignment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(232)
		p.Match(BistatParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(233)
		p.Expression()
	}

	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__12 {
		{
			p.SetState(234)
			p.Match(BistatParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(235)
			p.Expression()
		}

		p.SetState(240)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(241)
		p.Match(BistatParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_comment
	return p
}

func InitEmptyCommentContext(p *CommentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_comment
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, BistatParserRULE_comment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.Match(BistatParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4611686018427355134) != 0) {
		{
			p.SetState(244)
			_la = p.GetTokenStream().LA(1)

			if _la <= 0 || _la == BistatParserT__14 {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(247)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(249)
		p.Match(BistatParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForLoopContext() *ForLoopContext {
	var p = new(ForLoopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_forLoop
	return p
}

func InitEmptyForLoopContext(p *ForLoopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_forLoop
}

func (*ForLoopContext) IsForLoopContext() {}

func NewForLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForLoopContext {
	var p = new(ForLoopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewForLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, BistatParserRULE_forLoop)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(251)
		p.ForHeader()
	}
	{
		p.SetState(252)
		p.Match(BistatParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&144116287586271232) != 0) {
		{
			p.SetState(253)
			p.Stmt()
		}

		p.SetState(256)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(258)
		p.Match(BistatParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForHeaderContext() *ForHeaderContext {
	var p = new(ForHeaderContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_forHeader
	return p
}

func InitEmptyForHeaderContext(p *ForHeaderContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_forHeader
}

func (*ForHeaderContext) IsForHeaderContext() {}

func NewForHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForHeaderContext {
	var p = new(ForHeaderContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewForHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, BistatParserRULE_forHeader)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(260)
		p.Match(BistatParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(261)
		p.Match(BistatParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(262)
		p.Match(BistatParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(263)
		p.Match(BistatParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(264)
		p.Expression()
	}
	{
		p.SetState(265)
		p.Match(BistatParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForExprEndContext() *ForExprEndContext {
	var p = new(ForExprEndContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_forExprEnd
	return p
}

func InitEmptyForExprEndContext(p *ForExprEndContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_forExprEnd
}

func (*ForExprEndContext) IsForExprEndContext() {}

func NewForExprEndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForExprEndContext {
	var p = new(ForExprEndContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewForExprEndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, BistatParserRULE_forExprEnd)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(267)
		p.Match(BistatParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileLoopContext() *WhileLoopContext {
	var p = new(WhileLoopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_whileLoop
	return p
}

func InitEmptyWhileLoopContext(p *WhileLoopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_whileLoop
}

func (*WhileLoopContext) IsWhileLoopContext() {}

func NewWhileLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileLoopContext {
	var p = new(WhileLoopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewWhileLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, BistatParserRULE_whileLoop)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(269)
		p.Match(BistatParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(270)
		p.Match(BistatParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(271)
		p.Expression()
	}
	{
		p.SetState(272)
		p.WhileExprEnd()
	}
	{
		p.SetState(273)
		p.Match(BistatParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&144116287586271232) != 0) {
		{
			p.SetState(274)
			p.Stmt()
		}

		p.SetState(277)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(279)
		p.Match(BistatParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileExprEndContext() *WhileExprEndContext {
	var p = new(WhileExprEndContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_whileExprEnd
	return p
}

func InitEmptyWhileExprEndContext(p *WhileExprEndContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_whileExprEnd
}

func (*WhileExprEndContext) IsWhileExprEndContext() {}

func NewWhileExprEndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileExprEndContext {
	var p = new(WhileExprEndContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewWhileExprEndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, BistatParserRULE_whileExprEnd)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(281)
		p.Match(BistatParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalContext() *ConditionalContext {
	var p = new(ConditionalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_conditional
	return p
}

func InitEmptyConditionalContext(p *ConditionalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_conditional
}

func (*ConditionalContext) IsConditionalContext() {}

func NewConditionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalContext {
	var p = new(ConditionalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewConditionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, BistatParserRULE_conditional)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(283)
		p.IfStmt()
	}
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(284)
				p.ElseIfStmt()
			}

		}
		p.SetState(289)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BistatParserT__19 {
		{
			p.SetState(290)
			p.ElseStmt()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_ifStmt
	return p
}

func InitEmptyIfStmtContext(p *IfStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_ifStmt
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, BistatParserRULE_ifStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(293)
		p.Match(BistatParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(294)
		p.Match(BistatParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(295)
		p.Expression()
	}
	{
		p.SetState(296)
		p.CondExprEnd()
	}
	{
		p.SetState(297)
		p.Match(BistatParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&144116287586271232) != 0) {
		{
			p.SetState(298)
			p.Stmt()
		}

		p.SetState(301)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(303)
		p.Match(BistatParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseIfStmtContext() *ElseIfStmtContext {
	var p = new(ElseIfStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_elseIfStmt
	return p
}

func InitEmptyElseIfStmtContext(p *ElseIfStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_elseIfStmt
}

func (*ElseIfStmtContext) IsElseIfStmtContext() {}

func NewElseIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseIfStmtContext {
	var p = new(ElseIfStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewElseIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, BistatParserRULE_elseIfStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(305)
		p.Match(BistatParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(306)
		p.Match(BistatParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(307)
		p.Match(BistatParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(308)
		p.Expression()
	}
	{
		p.SetState(309)
		p.CondExprEnd()
	}
	{
		p.SetState(310)
		p.Match(BistatParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(312)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&144116287586271232) != 0) {
		{
			p.SetState(311)
			p.Stmt()
		}

		p.SetState(314)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(316)
		p.Match(BistatParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCondExprEndContext() *CondExprEndContext {
	var p = new(CondExprEndContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_condExprEnd
	return p
}

func InitEmptyCondExprEndContext(p *CondExprEndContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_condExprEnd
}

func (*CondExprEndContext) IsCondExprEndContext() {}

func NewCondExprEndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CondExprEndContext {
	var p = new(CondExprEndContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewCondExprEndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, BistatParserRULE_condExprEnd)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(318)
		p.Match(BistatParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseStmtContext() *ElseStmtContext {
	var p = new(ElseStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_elseStmt
	return p
}

func InitEmptyElseStmtContext(p *ElseStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_elseStmt
}

func (*ElseStmtContext) IsElseStmtContext() {}

func NewElseStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseStmtContext {
	var p = new(ElseStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewElseStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, BistatParserRULE_elseStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(320)
		p.Match(BistatParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(321)
		p.Match(BistatParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(323)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&144116287586271232) != 0) {
		{
			p.SetState(322)
			p.Stmt()
		}

		p.SetState(325)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(327)
		p.Match(BistatParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStmtContext() *ReturnStmtContext {
	var p = new(ReturnStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_returnStmt
	return p
}

func InitEmptyReturnStmtContext(p *ReturnStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_returnStmt
}

func (*ReturnStmtContext) IsReturnStmtContext() {}

func NewReturnStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewReturnStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, BistatParserRULE_returnStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(329)
		p.Match(BistatParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(330)
		p.Expression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISpecialFunctionContext is an interface to support dynamic dispatch.
type ISpecialFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	InputRead() IInputReadContext
	Print_() IPrintContext
	Plot() IPlotContext
	Sum() ISumContext
	Max() IMaxContext
	Min() IMinContext
	Prod() IProdContext
	Avg() IAvgContext
	SMode() ISModeContext
	Median() IMedianContext
	Sin() ISinContext
	Cos() ICosContext
	Tan() ITanContext
	Sqrt() ISqrtContext
	Floor() IFloorContext
	Ceil() ICeilContext
	Abs() IAbsContext
	Not() INotContext

	// IsSpecialFunctionContext differentiates from other interfaces.
	IsSpecialFunctionContext()
}

type SpecialFunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpecialFunctionContext() *SpecialFunctionContext {
	var p = new(SpecialFunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_specialFunction
	return p
}

func InitEmptySpecialFunctionContext(p *SpecialFunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_specialFunction
}

func (*SpecialFunctionContext) IsSpecialFunctionContext() {}

func NewSpecialFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecialFunctionContext {
	var p = new(SpecialFunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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

func (s *SpecialFunctionContext) Max() IMaxContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMaxContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMaxContext)
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
	localctx = NewSpecialFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, BistatParserRULE_specialFunction)
	p.SetState(350)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BistatParserT__21:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(332)
			p.InputRead()
		}

	case BistatParserT__22:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(333)
			p.Print_()
		}

	case BistatParserT__23:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(334)
			p.Plot()
		}

	case BistatParserT__24:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(335)
			p.Sum()
		}

	case BistatParserT__26:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(336)
			p.Max()
		}

	case BistatParserT__25:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(337)
			p.Min()
		}

	case BistatParserT__27:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(338)
			p.Prod()
		}

	case BistatParserT__28:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(339)
			p.Avg()
		}

	case BistatParserT__29:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(340)
			p.SMode()
		}

	case BistatParserT__30:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(341)
			p.Median()
		}

	case BistatParserT__31:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(342)
			p.Sin()
		}

	case BistatParserT__33:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(343)
			p.Cos()
		}

	case BistatParserT__32:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(344)
			p.Tan()
		}

	case BistatParserT__34:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(345)
			p.Sqrt()
		}

	case BistatParserT__35:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(346)
			p.Floor()
		}

	case BistatParserT__36:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(347)
			p.Ceil()
		}

	case BistatParserT__37:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(348)
			p.Abs()
		}

	case BistatParserT__38:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(349)
			p.Not()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInputReadContext is an interface to support dynamic dispatch.
type IInputReadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsInputReadContext differentiates from other interfaces.
	IsInputReadContext()
}

type InputReadContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputReadContext() *InputReadContext {
	var p = new(InputReadContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_inputRead
	return p
}

func InitEmptyInputReadContext(p *InputReadContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_inputRead
}

func (*InputReadContext) IsInputReadContext() {}

func NewInputReadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputReadContext {
	var p = new(InputReadContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_inputRead

	return p
}

func (s *InputReadContext) GetParser() antlr.Parser { return s.parser }

func (s *InputReadContext) AllExpression() []IExpressionContext {
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

func (s *InputReadContext) Expression(i int) IExpressionContext {
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
	localctx = NewInputReadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, BistatParserRULE_inputRead)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(352)
		p.Match(BistatParserT__21)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(353)
		p.Match(BistatParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(354)
		p.Expression()
	}
	p.SetState(359)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__12 {
		{
			p.SetState(355)
			p.Match(BistatParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(356)
			p.Expression()
		}

		p.SetState(361)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(362)
		p.Match(BistatParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintContext() *PrintContext {
	var p = new(PrintContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_print
	return p
}

func InitEmptyPrintContext(p *PrintContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_print
}

func (*PrintContext) IsPrintContext() {}

func NewPrintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintContext {
	var p = new(PrintContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewPrintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, BistatParserRULE_print)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(364)
		p.Match(BistatParserT__22)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(365)
		p.Match(BistatParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(366)
		p.Expression()
	}
	p.SetState(371)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__12 {
		{
			p.SetState(367)
			p.Match(BistatParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(368)
			p.Expression()
		}

		p.SetState(373)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(374)
		p.Match(BistatParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPlotContext() *PlotContext {
	var p = new(PlotContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_plot
	return p
}

func InitEmptyPlotContext(p *PlotContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_plot
}

func (*PlotContext) IsPlotContext() {}

func NewPlotContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PlotContext {
	var p = new(PlotContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewPlotContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, BistatParserRULE_plot)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(376)
		p.Match(BistatParserT__23)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(377)
		p.Match(BistatParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(378)
		p.Expression()
	}
	{
		p.SetState(379)
		p.Match(BistatParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySumContext() *SumContext {
	var p = new(SumContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_sum
	return p
}

func InitEmptySumContext(p *SumContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_sum
}

func (*SumContext) IsSumContext() {}

func NewSumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SumContext {
	var p = new(SumContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewSumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, BistatParserRULE_sum)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(381)
		p.Match(BistatParserT__24)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(382)
		p.Match(BistatParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(383)
		p.Expression()
	}
	{
		p.SetState(384)
		p.Match(BistatParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMinContext() *MinContext {
	var p = new(MinContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_min
	return p
}

func InitEmptyMinContext(p *MinContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_min
}

func (*MinContext) IsMinContext() {}

func NewMinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MinContext {
	var p = new(MinContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewMinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, BistatParserRULE_min)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(386)
		p.Match(BistatParserT__25)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(387)
		p.Match(BistatParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(388)
		p.Expression()
	}
	{
		p.SetState(389)
		p.Match(BistatParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMaxContext is an interface to support dynamic dispatch.
type IMaxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsMaxContext differentiates from other interfaces.
	IsMaxContext()
}

type MaxContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMaxContext() *MaxContext {
	var p = new(MaxContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_max
	return p
}

func InitEmptyMaxContext(p *MaxContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_max
}

func (*MaxContext) IsMaxContext() {}

func NewMaxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MaxContext {
	var p = new(MaxContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_max

	return p
}

func (s *MaxContext) GetParser() antlr.Parser { return s.parser }

func (s *MaxContext) Expression() IExpressionContext {
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

func (s *MaxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MaxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MaxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterMax(s)
	}
}

func (s *MaxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitMax(s)
	}
}

func (p *BistatParser) Max() (localctx IMaxContext) {
	localctx = NewMaxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, BistatParserRULE_max)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(391)
		p.Match(BistatParserT__26)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(392)
		p.Match(BistatParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(393)
		p.Expression()
	}
	{
		p.SetState(394)
		p.Match(BistatParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProdContext() *ProdContext {
	var p = new(ProdContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_prod
	return p
}

func InitEmptyProdContext(p *ProdContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_prod
}

func (*ProdContext) IsProdContext() {}

func NewProdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProdContext {
	var p = new(ProdContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewProdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, BistatParserRULE_prod)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(396)
		p.Match(BistatParserT__27)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(397)
		p.Match(BistatParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(398)
		p.Expression()
	}
	{
		p.SetState(399)
		p.Match(BistatParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAvgContext() *AvgContext {
	var p = new(AvgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_avg
	return p
}

func InitEmptyAvgContext(p *AvgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_avg
}

func (*AvgContext) IsAvgContext() {}

func NewAvgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AvgContext {
	var p = new(AvgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewAvgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, BistatParserRULE_avg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(401)
		p.Match(BistatParserT__28)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(402)
		p.Match(BistatParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(403)
		p.Expression()
	}
	{
		p.SetState(404)
		p.Match(BistatParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySModeContext() *SModeContext {
	var p = new(SModeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_sMode
	return p
}

func InitEmptySModeContext(p *SModeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_sMode
}

func (*SModeContext) IsSModeContext() {}

func NewSModeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SModeContext {
	var p = new(SModeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewSModeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, BistatParserRULE_sMode)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(406)
		p.Match(BistatParserT__29)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(407)
		p.Match(BistatParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(408)
		p.Expression()
	}
	{
		p.SetState(409)
		p.Match(BistatParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMedianContext() *MedianContext {
	var p = new(MedianContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_median
	return p
}

func InitEmptyMedianContext(p *MedianContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_median
}

func (*MedianContext) IsMedianContext() {}

func NewMedianContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MedianContext {
	var p = new(MedianContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewMedianContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, BistatParserRULE_median)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(411)
		p.Match(BistatParserT__30)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(412)
		p.Match(BistatParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(413)
		p.Expression()
	}
	{
		p.SetState(414)
		p.Match(BistatParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySinContext() *SinContext {
	var p = new(SinContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_sin
	return p
}

func InitEmptySinContext(p *SinContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_sin
}

func (*SinContext) IsSinContext() {}

func NewSinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SinContext {
	var p = new(SinContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewSinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, BistatParserRULE_sin)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(416)
		p.Match(BistatParserT__31)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(417)
		p.Match(BistatParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(418)
		p.Expression()
	}
	{
		p.SetState(419)
		p.Match(BistatParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTanContext() *TanContext {
	var p = new(TanContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_tan
	return p
}

func InitEmptyTanContext(p *TanContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_tan
}

func (*TanContext) IsTanContext() {}

func NewTanContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TanContext {
	var p = new(TanContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewTanContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, BistatParserRULE_tan)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(421)
		p.Match(BistatParserT__32)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(422)
		p.Match(BistatParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(423)
		p.Expression()
	}
	{
		p.SetState(424)
		p.Match(BistatParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCosContext() *CosContext {
	var p = new(CosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_cos
	return p
}

func InitEmptyCosContext(p *CosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_cos
}

func (*CosContext) IsCosContext() {}

func NewCosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CosContext {
	var p = new(CosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewCosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, BistatParserRULE_cos)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(426)
		p.Match(BistatParserT__33)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(427)
		p.Match(BistatParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(428)
		p.Expression()
	}
	{
		p.SetState(429)
		p.Match(BistatParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySqrtContext() *SqrtContext {
	var p = new(SqrtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_sqrt
	return p
}

func InitEmptySqrtContext(p *SqrtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_sqrt
}

func (*SqrtContext) IsSqrtContext() {}

func NewSqrtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SqrtContext {
	var p = new(SqrtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewSqrtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, BistatParserRULE_sqrt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(431)
		p.Match(BistatParserT__34)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(432)
		p.Match(BistatParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(433)
		p.Expression()
	}
	{
		p.SetState(434)
		p.Match(BistatParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloorContext() *FloorContext {
	var p = new(FloorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_floor
	return p
}

func InitEmptyFloorContext(p *FloorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_floor
}

func (*FloorContext) IsFloorContext() {}

func NewFloorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloorContext {
	var p = new(FloorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewFloorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, BistatParserRULE_floor)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(436)
		p.Match(BistatParserT__35)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(437)
		p.Match(BistatParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(438)
		p.Expression()
	}
	{
		p.SetState(439)
		p.Match(BistatParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCeilContext() *CeilContext {
	var p = new(CeilContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_ceil
	return p
}

func InitEmptyCeilContext(p *CeilContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_ceil
}

func (*CeilContext) IsCeilContext() {}

func NewCeilContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CeilContext {
	var p = new(CeilContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewCeilContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, BistatParserRULE_ceil)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(441)
		p.Match(BistatParserT__36)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(442)
		p.Match(BistatParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(443)
		p.Expression()
	}
	{
		p.SetState(444)
		p.Match(BistatParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAbsContext() *AbsContext {
	var p = new(AbsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_abs
	return p
}

func InitEmptyAbsContext(p *AbsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_abs
}

func (*AbsContext) IsAbsContext() {}

func NewAbsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AbsContext {
	var p = new(AbsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewAbsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, BistatParserRULE_abs)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(446)
		p.Match(BistatParserT__37)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(447)
		p.Match(BistatParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(448)
		p.Expression()
	}
	{
		p.SetState(449)
		p.Match(BistatParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotContext() *NotContext {
	var p = new(NotContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_not
	return p
}

func InitEmptyNotContext(p *NotContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_not
}

func (*NotContext) IsNotContext() {}

func NewNotContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotContext {
	var p = new(NotContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewNotContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, BistatParserRULE_not)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(451)
		p.Match(BistatParserT__38)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(452)
		p.Match(BistatParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(453)
		p.Expression()
	}
	{
		p.SetState(454)
		p.Match(BistatParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, BistatParserRULE_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(456)
		p.Exp()
	}
	p.SetState(460)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8972014882652160) != 0 {
		{
			p.SetState(457)
			p.LogicOperator()
		}
		{
			p.SetState(458)
			p.Exp()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_exp
	return p
}

func InitEmptyExpContext(p *ExpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_exp
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, BistatParserRULE_exp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(462)
		p.Term()
	}
	p.SetState(468)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__39 || _la == BistatParserT__40 {
		{
			p.SetState(463)
			p.OpSec()
		}
		{
			p.SetState(464)
			p.Term()
		}

		p.SetState(470)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_term
	return p
}

func InitEmptyTermContext(p *TermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_term
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, BistatParserRULE_term)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(471)
		p.Factor()
	}
	p.SetState(477)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30786325577728) != 0 {
		{
			p.SetState(472)
			p.OpFirst()
		}
		{
			p.SetState(473)
			p.Factor()
		}

		p.SetState(479)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NestedExpression() INestedExpressionContext
	SpecialFunction() ISpecialFunctionContext
	FunctionCall() IFunctionCallContext
	Indexing() IIndexingContext
	VarCons() IVarConsContext
	UnaryMinus() IUnaryMinusContext

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_factor
	return p
}

func InitEmptyFactorContext(p *FactorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_factor
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, BistatParserRULE_factor)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(481)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BistatParserT__39 {
		{
			p.SetState(480)
			p.UnaryMinus()
		}

	}
	p.SetState(488)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(483)
			p.NestedExpression()
		}

	case 2:
		{
			p.SetState(484)
			p.SpecialFunction()
		}

	case 3:
		{
			p.SetState(485)
			p.FunctionCall()
		}

	case 4:
		{
			p.SetState(486)
			p.Indexing()
		}

	case 5:
		{
			p.SetState(487)
			p.VarCons()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryMinusContext() *UnaryMinusContext {
	var p = new(UnaryMinusContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_unaryMinus
	return p
}

func InitEmptyUnaryMinusContext(p *UnaryMinusContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_unaryMinus
}

func (*UnaryMinusContext) IsUnaryMinusContext() {}

func NewUnaryMinusContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryMinusContext {
	var p = new(UnaryMinusContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewUnaryMinusContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, BistatParserRULE_unaryMinus)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(490)
		p.Match(BistatParserT__39)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNestedExpressionContext() *NestedExpressionContext {
	var p = new(NestedExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_nestedExpression
	return p
}

func InitEmptyNestedExpressionContext(p *NestedExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_nestedExpression
}

func (*NestedExpressionContext) IsNestedExpressionContext() {}

func NewNestedExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NestedExpressionContext {
	var p = new(NestedExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewNestedExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, BistatParserRULE_nestedExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(492)
		p.Match(BistatParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(493)
		p.Expression()
	}
	{
		p.SetState(494)
		p.Match(BistatParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_functionCall
	return p
}

func InitEmptyFunctionCallContext(p *FunctionCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_functionCall
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, BistatParserRULE_functionCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(496)
		p.Match(BistatParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(497)
		p.Match(BistatParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(506)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3963169871105097760) != 0 {
		{
			p.SetState(498)
			p.Expression()
		}
		p.SetState(503)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == BistatParserT__12 {
			{
				p.SetState(499)
				p.Match(BistatParserT__12)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(500)
				p.Expression()
			}

			p.SetState(505)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(508)
		p.Match(BistatParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarConsContext() *VarConsContext {
	var p = new(VarConsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_varCons
	return p
}

func InitEmptyVarConsContext(p *VarConsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_varCons
}

func (*VarConsContext) IsVarConsContext() {}

func NewVarConsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarConsContext {
	var p = new(VarConsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewVarConsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, BistatParserRULE_varCons)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(510)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3963167672086036480) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpSecContext() *OpSecContext {
	var p = new(OpSecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_opSec
	return p
}

func InitEmptyOpSecContext(p *OpSecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_opSec
}

func (*OpSecContext) IsOpSecContext() {}

func NewOpSecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpSecContext {
	var p = new(OpSecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewOpSecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, BistatParserRULE_opSec)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(512)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BistatParserT__39 || _la == BistatParserT__40) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpFirstContext() *OpFirstContext {
	var p = new(OpFirstContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_opFirst
	return p
}

func InitEmptyOpFirstContext(p *OpFirstContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_opFirst
}

func (*OpFirstContext) IsOpFirstContext() {}

func NewOpFirstContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpFirstContext {
	var p = new(OpFirstContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewOpFirstContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, BistatParserRULE_opFirst)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(514)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30786325577728) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicOperatorContext() *LogicOperatorContext {
	var p = new(LogicOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_logicOperator
	return p
}

func InitEmptyLogicOperatorContext(p *LogicOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BistatParserRULE_logicOperator
}

func (*LogicOperatorContext) IsLogicOperatorContext() {}

func NewLogicOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicOperatorContext {
	var p = new(LogicOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	localctx = NewLogicOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, BistatParserRULE_logicOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(516)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8972014882652160) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
