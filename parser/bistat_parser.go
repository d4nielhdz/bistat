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
		"", "'Program'", "';'", "'print'", "'('", "','", "')'", "'main'", "'{'",
		"'}'", "'func'", "':'", "'var'", "'='", "'['", "']'", "'#'", "'for'",
		"'in'", "'while'", "'if'", "'else'", "'read'", "'list_add'", "'list_pop'",
		"'length'", "'range'", "'plot'", "'sum'", "'min'", "'prod'", "'avg'",
		"'s_mode'", "'median'", "'sin'", "'tan'", "'cos'", "'sort'", "'sqrt'",
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
		"program", "stmt", "nested_stmt", "print", "expression", "main", "func_def",
		"var_declaration", "param_declaration", "assignment", "matrix_assignment",
		"list_assignment", "comment", "for", "while", "conditional", "input_read",
		"list_add", "list_pop", "length", "range", "plot", "sum", "min", "prod",
		"avg", "s_mode", "median", "sin", "tan", "cos", "sort", "sqrt", "floor",
		"ceil", "abs", "not", "special_function", "exp", "term", "factor", "function_call",
		"indexing",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 58, 488, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 91, 8, 0, 10, 0, 12, 0, 94, 9,
		0, 1, 0, 5, 0, 97, 8, 0, 10, 0, 12, 0, 100, 9, 0, 1, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 3, 1, 109, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 3, 1, 117, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 123, 8, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 3, 2, 130, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3,
		137, 8, 3, 10, 3, 12, 3, 140, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 3, 4,
		147, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 4, 5, 154, 8, 5, 11, 5, 12, 5,
		155, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 164, 8, 6, 10, 6, 12, 6,
		167, 9, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 173, 8, 6, 10, 6, 12, 6, 176,
		9, 6, 1, 6, 1, 6, 4, 6, 180, 8, 6, 11, 6, 12, 6, 181, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 3, 9, 201, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 207, 8,
		10, 10, 10, 12, 10, 210, 9, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11,
		5, 11, 218, 8, 11, 10, 11, 12, 11, 221, 9, 11, 1, 11, 1, 11, 1, 12, 1,
		12, 4, 12, 227, 8, 12, 11, 12, 12, 12, 228, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 4, 13, 241, 8, 13, 11, 13, 12,
		13, 242, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 4, 14,
		253, 8, 14, 11, 14, 12, 14, 254, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 4, 15, 265, 8, 15, 11, 15, 12, 15, 266, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 4, 15, 277, 8, 15, 11, 15, 12,
		15, 278, 1, 15, 1, 15, 5, 15, 283, 8, 15, 10, 15, 12, 15, 286, 9, 15, 1,
		15, 1, 15, 1, 15, 4, 15, 291, 8, 15, 11, 15, 12, 15, 292, 1, 15, 1, 15,
		3, 15, 297, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 304, 8, 16,
		10, 16, 12, 16, 307, 9, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 333, 8, 20, 1,
		20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1,
		33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37,
		1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 3, 37, 439,
		8, 37, 1, 38, 1, 38, 1, 38, 3, 38, 444, 8, 38, 1, 39, 1, 39, 1, 39, 3,
		39, 449, 8, 39, 1, 40, 3, 40, 452, 8, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1,
		40, 1, 40, 1, 40, 1, 40, 3, 40, 462, 8, 40, 1, 41, 1, 41, 1, 41, 1, 41,
		1, 41, 5, 41, 469, 8, 41, 10, 41, 12, 41, 472, 9, 41, 3, 41, 474, 8, 41,
		1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 3,
		42, 486, 8, 42, 1, 42, 0, 0, 43, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56,
		58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 0, 1, 1, 0, 16,
		16, 510, 0, 86, 1, 0, 0, 0, 2, 116, 1, 0, 0, 0, 4, 129, 1, 0, 0, 0, 6,
		131, 1, 0, 0, 0, 8, 143, 1, 0, 0, 0, 10, 148, 1, 0, 0, 0, 12, 159, 1, 0,
		0, 0, 14, 185, 1, 0, 0, 0, 16, 190, 1, 0, 0, 0, 18, 195, 1, 0, 0, 0, 20,
		202, 1, 0, 0, 0, 22, 213, 1, 0, 0, 0, 24, 224, 1, 0, 0, 0, 26, 232, 1,
		0, 0, 0, 28, 246, 1, 0, 0, 0, 30, 258, 1, 0, 0, 0, 32, 298, 1, 0, 0, 0,
		34, 310, 1, 0, 0, 0, 36, 317, 1, 0, 0, 0, 38, 322, 1, 0, 0, 0, 40, 327,
		1, 0, 0, 0, 42, 336, 1, 0, 0, 0, 44, 341, 1, 0, 0, 0, 46, 346, 1, 0, 0,
		0, 48, 351, 1, 0, 0, 0, 50, 356, 1, 0, 0, 0, 52, 361, 1, 0, 0, 0, 54, 366,
		1, 0, 0, 0, 56, 371, 1, 0, 0, 0, 58, 376, 1, 0, 0, 0, 60, 381, 1, 0, 0,
		0, 62, 386, 1, 0, 0, 0, 64, 391, 1, 0, 0, 0, 66, 396, 1, 0, 0, 0, 68, 401,
		1, 0, 0, 0, 70, 406, 1, 0, 0, 0, 72, 411, 1, 0, 0, 0, 74, 438, 1, 0, 0,
		0, 76, 440, 1, 0, 0, 0, 78, 445, 1, 0, 0, 0, 80, 451, 1, 0, 0, 0, 82, 463,
		1, 0, 0, 0, 84, 477, 1, 0, 0, 0, 86, 87, 5, 1, 0, 0, 87, 88, 5, 47, 0,
		0, 88, 92, 5, 2, 0, 0, 89, 91, 3, 14, 7, 0, 90, 89, 1, 0, 0, 0, 91, 94,
		1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 98, 1, 0, 0, 0,
		94, 92, 1, 0, 0, 0, 95, 97, 3, 12, 6, 0, 96, 95, 1, 0, 0, 0, 97, 100, 1,
		0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 101, 1, 0, 0, 0, 100,
		98, 1, 0, 0, 0, 101, 102, 3, 10, 5, 0, 102, 103, 5, 0, 0, 1, 103, 1, 1,
		0, 0, 0, 104, 109, 3, 18, 9, 0, 105, 109, 3, 6, 3, 0, 106, 109, 3, 74,
		37, 0, 107, 109, 3, 82, 41, 0, 108, 104, 1, 0, 0, 0, 108, 105, 1, 0, 0,
		0, 108, 106, 1, 0, 0, 0, 108, 107, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110,
		111, 5, 2, 0, 0, 111, 117, 1, 0, 0, 0, 112, 117, 3, 30, 15, 0, 113, 117,
		3, 28, 14, 0, 114, 117, 3, 26, 13, 0, 115, 117, 3, 24, 12, 0, 116, 108,
		1, 0, 0, 0, 116, 112, 1, 0, 0, 0, 116, 113, 1, 0, 0, 0, 116, 114, 1, 0,
		0, 0, 116, 115, 1, 0, 0, 0, 117, 3, 1, 0, 0, 0, 118, 123, 3, 18, 9, 0,
		119, 123, 3, 6, 3, 0, 120, 123, 3, 74, 37, 0, 121, 123, 3, 82, 41, 0, 122,
		118, 1, 0, 0, 0, 122, 119, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 122, 121,
		1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 125, 5, 2, 0, 0, 125, 130, 1, 0,
		0, 0, 126, 130, 3, 30, 15, 0, 127, 130, 3, 28, 14, 0, 128, 130, 3, 24,
		12, 0, 129, 122, 1, 0, 0, 0, 129, 126, 1, 0, 0, 0, 129, 127, 1, 0, 0, 0,
		129, 128, 1, 0, 0, 0, 130, 5, 1, 0, 0, 0, 131, 132, 5, 3, 0, 0, 132, 133,
		5, 4, 0, 0, 133, 138, 3, 8, 4, 0, 134, 135, 5, 5, 0, 0, 135, 137, 3, 8,
		4, 0, 136, 134, 1, 0, 0, 0, 137, 140, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0,
		138, 139, 1, 0, 0, 0, 139, 141, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 141,
		142, 5, 6, 0, 0, 142, 7, 1, 0, 0, 0, 143, 146, 3, 76, 38, 0, 144, 145,
		5, 57, 0, 0, 145, 147, 3, 76, 38, 0, 146, 144, 1, 0, 0, 0, 146, 147, 1,
		0, 0, 0, 147, 9, 1, 0, 0, 0, 148, 149, 5, 7, 0, 0, 149, 150, 5, 4, 0, 0,
		150, 151, 5, 6, 0, 0, 151, 153, 5, 8, 0, 0, 152, 154, 3, 2, 1, 0, 153,
		152, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 155, 156,
		1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 158, 5, 9, 0, 0, 158, 11, 1, 0,
		0, 0, 159, 160, 5, 10, 0, 0, 160, 161, 5, 47, 0, 0, 161, 165, 5, 4, 0,
		0, 162, 164, 3, 16, 8, 0, 163, 162, 1, 0, 0, 0, 164, 167, 1, 0, 0, 0, 165,
		163, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 168, 1, 0, 0, 0, 167, 165,
		1, 0, 0, 0, 168, 169, 5, 6, 0, 0, 169, 170, 5, 11, 0, 0, 170, 174, 5, 45,
		0, 0, 171, 173, 3, 14, 7, 0, 172, 171, 1, 0, 0, 0, 173, 176, 1, 0, 0, 0,
		174, 172, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 177, 1, 0, 0, 0, 176,
		174, 1, 0, 0, 0, 177, 179, 5, 8, 0, 0, 178, 180, 3, 2, 1, 0, 179, 178,
		1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 181, 182, 1, 0,
		0, 0, 182, 183, 1, 0, 0, 0, 183, 184, 5, 9, 0, 0, 184, 13, 1, 0, 0, 0,
		185, 186, 5, 12, 0, 0, 186, 187, 5, 44, 0, 0, 187, 188, 5, 47, 0, 0, 188,
		189, 5, 2, 0, 0, 189, 15, 1, 0, 0, 0, 190, 191, 5, 12, 0, 0, 191, 192,
		5, 46, 0, 0, 192, 193, 5, 47, 0, 0, 193, 194, 5, 2, 0, 0, 194, 17, 1, 0,
		0, 0, 195, 196, 5, 47, 0, 0, 196, 200, 5, 13, 0, 0, 197, 201, 5, 50, 0,
		0, 198, 201, 3, 22, 11, 0, 199, 201, 3, 20, 10, 0, 200, 197, 1, 0, 0, 0,
		200, 198, 1, 0, 0, 0, 200, 199, 1, 0, 0, 0, 201, 19, 1, 0, 0, 0, 202, 203,
		5, 14, 0, 0, 203, 208, 3, 22, 11, 0, 204, 205, 5, 5, 0, 0, 205, 207, 3,
		22, 11, 0, 206, 204, 1, 0, 0, 0, 207, 210, 1, 0, 0, 0, 208, 206, 1, 0,
		0, 0, 208, 209, 1, 0, 0, 0, 209, 211, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0,
		211, 212, 5, 15, 0, 0, 212, 21, 1, 0, 0, 0, 213, 214, 5, 14, 0, 0, 214,
		219, 5, 50, 0, 0, 215, 216, 5, 5, 0, 0, 216, 218, 5, 50, 0, 0, 217, 215,
		1, 0, 0, 0, 218, 221, 1, 0, 0, 0, 219, 217, 1, 0, 0, 0, 219, 220, 1, 0,
		0, 0, 220, 222, 1, 0, 0, 0, 221, 219, 1, 0, 0, 0, 222, 223, 5, 15, 0, 0,
		223, 23, 1, 0, 0, 0, 224, 226, 5, 16, 0, 0, 225, 227, 8, 0, 0, 0, 226,
		225, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 228, 229,
		1, 0, 0, 0, 229, 230, 1, 0, 0, 0, 230, 231, 5, 16, 0, 0, 231, 25, 1, 0,
		0, 0, 232, 233, 5, 17, 0, 0, 233, 234, 5, 4, 0, 0, 234, 235, 5, 47, 0,
		0, 235, 236, 5, 18, 0, 0, 236, 237, 3, 8, 4, 0, 237, 238, 5, 6, 0, 0, 238,
		240, 5, 8, 0, 0, 239, 241, 3, 4, 2, 0, 240, 239, 1, 0, 0, 0, 241, 242,
		1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 242, 243, 1, 0, 0, 0, 243, 244, 1, 0,
		0, 0, 244, 245, 5, 9, 0, 0, 245, 27, 1, 0, 0, 0, 246, 247, 5, 19, 0, 0,
		247, 248, 5, 4, 0, 0, 248, 249, 3, 8, 4, 0, 249, 250, 5, 6, 0, 0, 250,
		252, 5, 8, 0, 0, 251, 253, 3, 2, 1, 0, 252, 251, 1, 0, 0, 0, 253, 254,
		1, 0, 0, 0, 254, 252, 1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255, 256, 1, 0,
		0, 0, 256, 257, 5, 9, 0, 0, 257, 29, 1, 0, 0, 0, 258, 259, 5, 20, 0, 0,
		259, 260, 5, 4, 0, 0, 260, 261, 3, 8, 4, 0, 261, 262, 5, 6, 0, 0, 262,
		264, 5, 8, 0, 0, 263, 265, 3, 2, 1, 0, 264, 263, 1, 0, 0, 0, 265, 266,
		1, 0, 0, 0, 266, 264, 1, 0, 0, 0, 266, 267, 1, 0, 0, 0, 267, 268, 1, 0,
		0, 0, 268, 284, 5, 9, 0, 0, 269, 270, 5, 21, 0, 0, 270, 271, 5, 20, 0,
		0, 271, 272, 5, 4, 0, 0, 272, 273, 3, 8, 4, 0, 273, 274, 5, 6, 0, 0, 274,
		276, 5, 8, 0, 0, 275, 277, 3, 2, 1, 0, 276, 275, 1, 0, 0, 0, 277, 278,
		1, 0, 0, 0, 278, 276, 1, 0, 0, 0, 278, 279, 1, 0, 0, 0, 279, 280, 1, 0,
		0, 0, 280, 281, 5, 9, 0, 0, 281, 283, 1, 0, 0, 0, 282, 269, 1, 0, 0, 0,
		283, 286, 1, 0, 0, 0, 284, 282, 1, 0, 0, 0, 284, 285, 1, 0, 0, 0, 285,
		296, 1, 0, 0, 0, 286, 284, 1, 0, 0, 0, 287, 288, 5, 21, 0, 0, 288, 290,
		5, 8, 0, 0, 289, 291, 3, 2, 1, 0, 290, 289, 1, 0, 0, 0, 291, 292, 1, 0,
		0, 0, 292, 290, 1, 0, 0, 0, 292, 293, 1, 0, 0, 0, 293, 294, 1, 0, 0, 0,
		294, 295, 5, 9, 0, 0, 295, 297, 1, 0, 0, 0, 296, 287, 1, 0, 0, 0, 296,
		297, 1, 0, 0, 0, 297, 31, 1, 0, 0, 0, 298, 299, 5, 22, 0, 0, 299, 300,
		5, 4, 0, 0, 300, 305, 5, 47, 0, 0, 301, 302, 5, 5, 0, 0, 302, 304, 5, 47,
		0, 0, 303, 301, 1, 0, 0, 0, 304, 307, 1, 0, 0, 0, 305, 303, 1, 0, 0, 0,
		305, 306, 1, 0, 0, 0, 306, 308, 1, 0, 0, 0, 307, 305, 1, 0, 0, 0, 308,
		309, 5, 6, 0, 0, 309, 33, 1, 0, 0, 0, 310, 311, 5, 23, 0, 0, 311, 312,
		5, 4, 0, 0, 312, 313, 3, 8, 4, 0, 313, 314, 5, 5, 0, 0, 314, 315, 3, 8,
		4, 0, 315, 316, 5, 6, 0, 0, 316, 35, 1, 0, 0, 0, 317, 318, 5, 24, 0, 0,
		318, 319, 5, 4, 0, 0, 319, 320, 3, 8, 4, 0, 320, 321, 5, 6, 0, 0, 321,
		37, 1, 0, 0, 0, 322, 323, 5, 25, 0, 0, 323, 324, 5, 4, 0, 0, 324, 325,
		3, 8, 4, 0, 325, 326, 5, 6, 0, 0, 326, 39, 1, 0, 0, 0, 327, 328, 5, 26,
		0, 0, 328, 329, 5, 4, 0, 0, 329, 332, 3, 8, 4, 0, 330, 331, 5, 5, 0, 0,
		331, 333, 3, 8, 4, 0, 332, 330, 1, 0, 0, 0, 332, 333, 1, 0, 0, 0, 333,
		334, 1, 0, 0, 0, 334, 335, 5, 6, 0, 0, 335, 41, 1, 0, 0, 0, 336, 337, 5,
		27, 0, 0, 337, 338, 5, 4, 0, 0, 338, 339, 3, 8, 4, 0, 339, 340, 5, 6, 0,
		0, 340, 43, 1, 0, 0, 0, 341, 342, 5, 28, 0, 0, 342, 343, 5, 4, 0, 0, 343,
		344, 3, 8, 4, 0, 344, 345, 5, 6, 0, 0, 345, 45, 1, 0, 0, 0, 346, 347, 5,
		29, 0, 0, 347, 348, 5, 4, 0, 0, 348, 349, 3, 8, 4, 0, 349, 350, 5, 6, 0,
		0, 350, 47, 1, 0, 0, 0, 351, 352, 5, 30, 0, 0, 352, 353, 5, 4, 0, 0, 353,
		354, 3, 8, 4, 0, 354, 355, 5, 6, 0, 0, 355, 49, 1, 0, 0, 0, 356, 357, 5,
		31, 0, 0, 357, 358, 5, 4, 0, 0, 358, 359, 3, 8, 4, 0, 359, 360, 5, 6, 0,
		0, 360, 51, 1, 0, 0, 0, 361, 362, 5, 32, 0, 0, 362, 363, 5, 4, 0, 0, 363,
		364, 3, 8, 4, 0, 364, 365, 5, 6, 0, 0, 365, 53, 1, 0, 0, 0, 366, 367, 5,
		33, 0, 0, 367, 368, 5, 4, 0, 0, 368, 369, 3, 8, 4, 0, 369, 370, 5, 6, 0,
		0, 370, 55, 1, 0, 0, 0, 371, 372, 5, 34, 0, 0, 372, 373, 5, 4, 0, 0, 373,
		374, 3, 8, 4, 0, 374, 375, 5, 6, 0, 0, 375, 57, 1, 0, 0, 0, 376, 377, 5,
		35, 0, 0, 377, 378, 5, 4, 0, 0, 378, 379, 3, 8, 4, 0, 379, 380, 5, 6, 0,
		0, 380, 59, 1, 0, 0, 0, 381, 382, 5, 36, 0, 0, 382, 383, 5, 4, 0, 0, 383,
		384, 3, 8, 4, 0, 384, 385, 5, 6, 0, 0, 385, 61, 1, 0, 0, 0, 386, 387, 5,
		37, 0, 0, 387, 388, 5, 4, 0, 0, 388, 389, 3, 8, 4, 0, 389, 390, 5, 6, 0,
		0, 390, 63, 1, 0, 0, 0, 391, 392, 5, 38, 0, 0, 392, 393, 5, 4, 0, 0, 393,
		394, 3, 8, 4, 0, 394, 395, 5, 6, 0, 0, 395, 65, 1, 0, 0, 0, 396, 397, 5,
		39, 0, 0, 397, 398, 5, 4, 0, 0, 398, 399, 3, 8, 4, 0, 399, 400, 5, 6, 0,
		0, 400, 67, 1, 0, 0, 0, 401, 402, 5, 40, 0, 0, 402, 403, 5, 4, 0, 0, 403,
		404, 3, 8, 4, 0, 404, 405, 5, 6, 0, 0, 405, 69, 1, 0, 0, 0, 406, 407, 5,
		41, 0, 0, 407, 408, 5, 4, 0, 0, 408, 409, 3, 8, 4, 0, 409, 410, 5, 6, 0,
		0, 410, 71, 1, 0, 0, 0, 411, 412, 5, 42, 0, 0, 412, 413, 5, 4, 0, 0, 413,
		414, 3, 8, 4, 0, 414, 415, 5, 6, 0, 0, 415, 73, 1, 0, 0, 0, 416, 439, 3,
		32, 16, 0, 417, 439, 3, 6, 3, 0, 418, 439, 3, 34, 17, 0, 419, 439, 3, 36,
		18, 0, 420, 439, 3, 38, 19, 0, 421, 439, 3, 40, 20, 0, 422, 439, 3, 42,
		21, 0, 423, 439, 3, 44, 22, 0, 424, 439, 3, 46, 23, 0, 425, 439, 3, 48,
		24, 0, 426, 439, 3, 50, 25, 0, 427, 439, 3, 52, 26, 0, 428, 439, 3, 54,
		27, 0, 429, 439, 3, 56, 28, 0, 430, 439, 3, 60, 30, 0, 431, 439, 3, 58,
		29, 0, 432, 439, 3, 62, 31, 0, 433, 439, 3, 64, 32, 0, 434, 439, 3, 66,
		33, 0, 435, 439, 3, 68, 34, 0, 436, 439, 3, 70, 35, 0, 437, 439, 3, 72,
		36, 0, 438, 416, 1, 0, 0, 0, 438, 417, 1, 0, 0, 0, 438, 418, 1, 0, 0, 0,
		438, 419, 1, 0, 0, 0, 438, 420, 1, 0, 0, 0, 438, 421, 1, 0, 0, 0, 438,
		422, 1, 0, 0, 0, 438, 423, 1, 0, 0, 0, 438, 424, 1, 0, 0, 0, 438, 425,
		1, 0, 0, 0, 438, 426, 1, 0, 0, 0, 438, 427, 1, 0, 0, 0, 438, 428, 1, 0,
		0, 0, 438, 429, 1, 0, 0, 0, 438, 430, 1, 0, 0, 0, 438, 431, 1, 0, 0, 0,
		438, 432, 1, 0, 0, 0, 438, 433, 1, 0, 0, 0, 438, 434, 1, 0, 0, 0, 438,
		435, 1, 0, 0, 0, 438, 436, 1, 0, 0, 0, 438, 437, 1, 0, 0, 0, 439, 75, 1,
		0, 0, 0, 440, 443, 3, 78, 39, 0, 441, 442, 5, 55, 0, 0, 442, 444, 3, 76,
		38, 0, 443, 441, 1, 0, 0, 0, 443, 444, 1, 0, 0, 0, 444, 77, 1, 0, 0, 0,
		445, 448, 3, 80, 40, 0, 446, 447, 5, 56, 0, 0, 447, 449, 3, 78, 39, 0,
		448, 446, 1, 0, 0, 0, 448, 449, 1, 0, 0, 0, 449, 79, 1, 0, 0, 0, 450, 452,
		5, 55, 0, 0, 451, 450, 1, 0, 0, 0, 451, 452, 1, 0, 0, 0, 452, 461, 1, 0,
		0, 0, 453, 454, 5, 4, 0, 0, 454, 455, 3, 8, 4, 0, 455, 456, 5, 6, 0, 0,
		456, 462, 1, 0, 0, 0, 457, 462, 3, 84, 42, 0, 458, 462, 3, 74, 37, 0, 459,
		462, 3, 82, 41, 0, 460, 462, 5, 50, 0, 0, 461, 453, 1, 0, 0, 0, 461, 457,
		1, 0, 0, 0, 461, 458, 1, 0, 0, 0, 461, 459, 1, 0, 0, 0, 461, 460, 1, 0,
		0, 0, 462, 81, 1, 0, 0, 0, 463, 464, 5, 47, 0, 0, 464, 473, 5, 4, 0, 0,
		465, 470, 3, 8, 4, 0, 466, 467, 5, 5, 0, 0, 467, 469, 3, 8, 4, 0, 468,
		466, 1, 0, 0, 0, 469, 472, 1, 0, 0, 0, 470, 468, 1, 0, 0, 0, 470, 471,
		1, 0, 0, 0, 471, 474, 1, 0, 0, 0, 472, 470, 1, 0, 0, 0, 473, 465, 1, 0,
		0, 0, 473, 474, 1, 0, 0, 0, 474, 475, 1, 0, 0, 0, 475, 476, 5, 6, 0, 0,
		476, 83, 1, 0, 0, 0, 477, 478, 5, 47, 0, 0, 478, 479, 5, 14, 0, 0, 479,
		480, 3, 8, 4, 0, 480, 485, 5, 15, 0, 0, 481, 482, 5, 14, 0, 0, 482, 483,
		3, 8, 4, 0, 483, 484, 5, 15, 0, 0, 484, 486, 1, 0, 0, 0, 485, 481, 1, 0,
		0, 0, 485, 486, 1, 0, 0, 0, 486, 85, 1, 0, 0, 0, 33, 92, 98, 108, 116,
		122, 129, 138, 146, 155, 165, 174, 181, 200, 208, 219, 228, 242, 254, 266,
		278, 284, 292, 296, 305, 332, 438, 443, 448, 451, 461, 470, 473, 485,
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
	BistatParserRULE_program           = 0
	BistatParserRULE_stmt              = 1
	BistatParserRULE_nested_stmt       = 2
	BistatParserRULE_print             = 3
	BistatParserRULE_expression        = 4
	BistatParserRULE_main              = 5
	BistatParserRULE_func_def          = 6
	BistatParserRULE_var_declaration   = 7
	BistatParserRULE_param_declaration = 8
	BistatParserRULE_assignment        = 9
	BistatParserRULE_matrix_assignment = 10
	BistatParserRULE_list_assignment   = 11
	BistatParserRULE_comment           = 12
	BistatParserRULE_for               = 13
	BistatParserRULE_while             = 14
	BistatParserRULE_conditional       = 15
	BistatParserRULE_input_read        = 16
	BistatParserRULE_list_add          = 17
	BistatParserRULE_list_pop          = 18
	BistatParserRULE_length            = 19
	BistatParserRULE_range             = 20
	BistatParserRULE_plot              = 21
	BistatParserRULE_sum               = 22
	BistatParserRULE_min               = 23
	BistatParserRULE_prod              = 24
	BistatParserRULE_avg               = 25
	BistatParserRULE_s_mode            = 26
	BistatParserRULE_median            = 27
	BistatParserRULE_sin               = 28
	BistatParserRULE_tan               = 29
	BistatParserRULE_cos               = 30
	BistatParserRULE_sort              = 31
	BistatParserRULE_sqrt              = 32
	BistatParserRULE_floor             = 33
	BistatParserRULE_ceil              = 34
	BistatParserRULE_abs               = 35
	BistatParserRULE_not               = 36
	BistatParserRULE_special_function  = 37
	BistatParserRULE_exp               = 38
	BistatParserRULE_term              = 39
	BistatParserRULE_factor            = 40
	BistatParserRULE_function_call     = 41
	BistatParserRULE_indexing          = 42
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
	AllVar_declaration() []IVar_declarationContext
	Var_declaration(i int) IVar_declarationContext
	AllFunc_def() []IFunc_defContext
	Func_def(i int) IFunc_defContext

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

func (s *ProgramContext) AllVar_declaration() []IVar_declarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVar_declarationContext); ok {
			len++
		}
	}

	tst := make([]IVar_declarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVar_declarationContext); ok {
			tst[i] = t.(IVar_declarationContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Var_declaration(i int) IVar_declarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_declarationContext); ok {
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

	return t.(IVar_declarationContext)
}

func (s *ProgramContext) AllFunc_def() []IFunc_defContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunc_defContext); ok {
			len++
		}
	}

	tst := make([]IFunc_defContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunc_defContext); ok {
			tst[i] = t.(IFunc_defContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Func_def(i int) IFunc_defContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_defContext); ok {
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

	return t.(IFunc_defContext)
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

	for _la == BistatParserT__11 {
		{
			p.SetState(89)
			p.Var_declaration()
		}

		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__9 {
		{
			p.SetState(95)
			p.Func_def()
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

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assignment() IAssignmentContext
	Print_() IPrintContext
	Special_function() ISpecial_functionContext
	Function_call() IFunction_callContext
	Conditional() IConditionalContext
	While() IWhileContext
	For_() IForContext
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

func (s *StmtContext) Print_() IPrintContext {
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

func (s *StmtContext) Special_function() ISpecial_functionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpecial_functionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpecial_functionContext)
}

func (s *StmtContext) Function_call() IFunction_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
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

func (s *StmtContext) While() IWhileContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileContext)
}

func (s *StmtContext) For_() IForContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForContext)
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
	p.EnterRule(localctx, 2, BistatParserRULE_stmt)

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

	p.SetState(116)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BistatParserT__2, BistatParserT__21, BistatParserT__22, BistatParserT__23, BistatParserT__24, BistatParserT__25, BistatParserT__26, BistatParserT__27, BistatParserT__28, BistatParserT__29, BistatParserT__30, BistatParserT__31, BistatParserT__32, BistatParserT__33, BistatParserT__34, BistatParserT__35, BistatParserT__36, BistatParserT__37, BistatParserT__38, BistatParserT__39, BistatParserT__40, BistatParserT__41, BistatParserID:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(104)
				p.Assignment()
			}

		case 2:
			{
				p.SetState(105)
				p.Print_()
			}

		case 3:
			{
				p.SetState(106)
				p.Special_function()
			}

		case 4:
			{
				p.SetState(107)
				p.Function_call()
			}

		}
		{
			p.SetState(110)
			p.Match(BistatParserT__1)
		}

	case BistatParserT__19:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(112)
			p.Conditional()
		}

	case BistatParserT__18:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(113)
			p.While()
		}

	case BistatParserT__16:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(114)
			p.For_()
		}

	case BistatParserT__15:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(115)
			p.Comment()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	Special_function() ISpecial_functionContext
	Function_call() IFunction_callContext
	Conditional() IConditionalContext
	While() IWhileContext
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

func (s *Nested_stmtContext) Special_function() ISpecial_functionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpecial_functionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpecial_functionContext)
}

func (s *Nested_stmtContext) Function_call() IFunction_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
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

func (s *Nested_stmtContext) While() IWhileContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileContext)
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
	p.EnterRule(localctx, 4, BistatParserRULE_nested_stmt)

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

	p.SetState(129)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BistatParserT__2, BistatParserT__21, BistatParserT__22, BistatParserT__23, BistatParserT__24, BistatParserT__25, BistatParserT__26, BistatParserT__27, BistatParserT__28, BistatParserT__29, BistatParserT__30, BistatParserT__31, BistatParserT__32, BistatParserT__33, BistatParserT__34, BistatParserT__35, BistatParserT__36, BistatParserT__37, BistatParserT__38, BistatParserT__39, BistatParserT__40, BistatParserT__41, BistatParserID:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(118)
				p.Assignment()
			}

		case 2:
			{
				p.SetState(119)
				p.Print_()
			}

		case 3:
			{
				p.SetState(120)
				p.Special_function()
			}

		case 4:
			{
				p.SetState(121)
				p.Function_call()
			}

		}
		{
			p.SetState(124)
			p.Match(BistatParserT__1)
		}

	case BistatParserT__19:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(126)
			p.Conditional()
		}

	case BistatParserT__18:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(127)
			p.While()
		}

	case BistatParserT__15:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(128)
			p.Comment()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.EnterRule(localctx, 6, BistatParserRULE_print)
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
		p.SetState(131)
		p.Match(BistatParserT__2)
	}
	{
		p.SetState(132)
		p.Match(BistatParserT__3)
	}
	{
		p.SetState(133)
		p.Expression()
	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__4 {
		{
			p.SetState(134)
			p.Match(BistatParserT__4)
		}
		{
			p.SetState(135)
			p.Expression()
		}

		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(141)
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
	p.EnterRule(localctx, 8, BistatParserRULE_expression)
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
		p.Exp()
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BistatParserLOGIC_OPERATOR {
		{
			p.SetState(144)
			p.Match(BistatParserLOGIC_OPERATOR)
		}
		{
			p.SetState(145)
			p.Exp()
		}

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
		p.Match(BistatParserT__6)
	}
	{
		p.SetState(149)
		p.Match(BistatParserT__3)
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

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&149533578952712) != 0) {
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

// IFunc_defContext is an interface to support dynamic dispatch.
type IFunc_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	RETURN_TYPE() antlr.TerminalNode
	AllParam_declaration() []IParam_declarationContext
	Param_declaration(i int) IParam_declarationContext
	AllVar_declaration() []IVar_declarationContext
	Var_declaration(i int) IVar_declarationContext
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsFunc_defContext differentiates from other interfaces.
	IsFunc_defContext()
}

type Func_defContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_defContext() *Func_defContext {
	var p = new(Func_defContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_func_def
	return p
}

func (*Func_defContext) IsFunc_defContext() {}

func NewFunc_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_defContext {
	var p = new(Func_defContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_func_def

	return p
}

func (s *Func_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_defContext) ID() antlr.TerminalNode {
	return s.GetToken(BistatParserID, 0)
}

func (s *Func_defContext) RETURN_TYPE() antlr.TerminalNode {
	return s.GetToken(BistatParserRETURN_TYPE, 0)
}

func (s *Func_defContext) AllParam_declaration() []IParam_declarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParam_declarationContext); ok {
			len++
		}
	}

	tst := make([]IParam_declarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParam_declarationContext); ok {
			tst[i] = t.(IParam_declarationContext)
			i++
		}
	}

	return tst
}

func (s *Func_defContext) Param_declaration(i int) IParam_declarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParam_declarationContext); ok {
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

	return t.(IParam_declarationContext)
}

func (s *Func_defContext) AllVar_declaration() []IVar_declarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVar_declarationContext); ok {
			len++
		}
	}

	tst := make([]IVar_declarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVar_declarationContext); ok {
			tst[i] = t.(IVar_declarationContext)
			i++
		}
	}

	return tst
}

func (s *Func_defContext) Var_declaration(i int) IVar_declarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_declarationContext); ok {
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

	return t.(IVar_declarationContext)
}

func (s *Func_defContext) AllStmt() []IStmtContext {
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

func (s *Func_defContext) Stmt(i int) IStmtContext {
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

func (s *Func_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Func_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterFunc_def(s)
	}
}

func (s *Func_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitFunc_def(s)
	}
}

func (p *BistatParser) Func_def() (localctx IFunc_defContext) {
	this := p
	_ = this

	localctx = NewFunc_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BistatParserRULE_func_def)
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
		p.SetState(159)
		p.Match(BistatParserT__9)
	}
	{
		p.SetState(160)
		p.Match(BistatParserID)
	}
	{
		p.SetState(161)
		p.Match(BistatParserT__3)
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__11 {
		{
			p.SetState(162)
			p.Param_declaration()
		}

		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(168)
		p.Match(BistatParserT__5)
	}
	{
		p.SetState(169)
		p.Match(BistatParserT__10)
	}
	{
		p.SetState(170)
		p.Match(BistatParserRETURN_TYPE)
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__11 {
		{
			p.SetState(171)
			p.Var_declaration()
		}

		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(177)
		p.Match(BistatParserT__7)
	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&149533578952712) != 0) {
		{
			p.SetState(178)
			p.Stmt()
		}

		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(183)
		p.Match(BistatParserT__8)
	}

	return localctx
}

// IVar_declarationContext is an interface to support dynamic dispatch.
type IVar_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NON_VOID_TYPE() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsVar_declarationContext differentiates from other interfaces.
	IsVar_declarationContext()
}

type Var_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_declarationContext() *Var_declarationContext {
	var p = new(Var_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_var_declaration
	return p
}

func (*Var_declarationContext) IsVar_declarationContext() {}

func NewVar_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_declarationContext {
	var p = new(Var_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_var_declaration

	return p
}

func (s *Var_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_declarationContext) NON_VOID_TYPE() antlr.TerminalNode {
	return s.GetToken(BistatParserNON_VOID_TYPE, 0)
}

func (s *Var_declarationContext) ID() antlr.TerminalNode {
	return s.GetToken(BistatParserID, 0)
}

func (s *Var_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterVar_declaration(s)
	}
}

func (s *Var_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitVar_declaration(s)
	}
}

func (p *BistatParser) Var_declaration() (localctx IVar_declarationContext) {
	this := p
	_ = this

	localctx = NewVar_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BistatParserRULE_var_declaration)

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
		p.SetState(185)
		p.Match(BistatParserT__11)
	}
	{
		p.SetState(186)
		p.Match(BistatParserNON_VOID_TYPE)
	}
	{
		p.SetState(187)
		p.Match(BistatParserID)
	}
	{
		p.SetState(188)
		p.Match(BistatParserT__1)
	}

	return localctx
}

// IParam_declarationContext is an interface to support dynamic dispatch.
type IParam_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PARAM_TYPE() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsParam_declarationContext differentiates from other interfaces.
	IsParam_declarationContext()
}

type Param_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParam_declarationContext() *Param_declarationContext {
	var p = new(Param_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_param_declaration
	return p
}

func (*Param_declarationContext) IsParam_declarationContext() {}

func NewParam_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Param_declarationContext {
	var p = new(Param_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_param_declaration

	return p
}

func (s *Param_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Param_declarationContext) PARAM_TYPE() antlr.TerminalNode {
	return s.GetToken(BistatParserPARAM_TYPE, 0)
}

func (s *Param_declarationContext) ID() antlr.TerminalNode {
	return s.GetToken(BistatParserID, 0)
}

func (s *Param_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Param_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Param_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterParam_declaration(s)
	}
}

func (s *Param_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitParam_declaration(s)
	}
}

func (p *BistatParser) Param_declaration() (localctx IParam_declarationContext) {
	this := p
	_ = this

	localctx = NewParam_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, BistatParserRULE_param_declaration)

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
		p.SetState(190)
		p.Match(BistatParserT__11)
	}
	{
		p.SetState(191)
		p.Match(BistatParserPARAM_TYPE)
	}
	{
		p.SetState(192)
		p.Match(BistatParserID)
	}
	{
		p.SetState(193)
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
	List_assignment() IList_assignmentContext
	Matrix_assignment() IMatrix_assignmentContext

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

func (s *AssignmentContext) List_assignment() IList_assignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_assignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_assignmentContext)
}

func (s *AssignmentContext) Matrix_assignment() IMatrix_assignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrix_assignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrix_assignmentContext)
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
	p.EnterRule(localctx, 18, BistatParserRULE_assignment)

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
		p.Match(BistatParserID)
	}
	{
		p.SetState(196)
		p.Match(BistatParserT__12)
	}
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(197)
			p.Match(BistatParserVAR_CONS)
		}

	case 2:
		{
			p.SetState(198)
			p.List_assignment()
		}

	case 3:
		{
			p.SetState(199)
			p.Matrix_assignment()
		}

	}

	return localctx
}

// IMatrix_assignmentContext is an interface to support dynamic dispatch.
type IMatrix_assignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllList_assignment() []IList_assignmentContext
	List_assignment(i int) IList_assignmentContext

	// IsMatrix_assignmentContext differentiates from other interfaces.
	IsMatrix_assignmentContext()
}

type Matrix_assignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatrix_assignmentContext() *Matrix_assignmentContext {
	var p = new(Matrix_assignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_matrix_assignment
	return p
}

func (*Matrix_assignmentContext) IsMatrix_assignmentContext() {}

func NewMatrix_assignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Matrix_assignmentContext {
	var p = new(Matrix_assignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_matrix_assignment

	return p
}

func (s *Matrix_assignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *Matrix_assignmentContext) AllList_assignment() []IList_assignmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IList_assignmentContext); ok {
			len++
		}
	}

	tst := make([]IList_assignmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IList_assignmentContext); ok {
			tst[i] = t.(IList_assignmentContext)
			i++
		}
	}

	return tst
}

func (s *Matrix_assignmentContext) List_assignment(i int) IList_assignmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_assignmentContext); ok {
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

	return t.(IList_assignmentContext)
}

func (s *Matrix_assignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Matrix_assignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Matrix_assignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterMatrix_assignment(s)
	}
}

func (s *Matrix_assignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitMatrix_assignment(s)
	}
}

func (p *BistatParser) Matrix_assignment() (localctx IMatrix_assignmentContext) {
	this := p
	_ = this

	localctx = NewMatrix_assignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, BistatParserRULE_matrix_assignment)
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
		p.SetState(202)
		p.Match(BistatParserT__13)
	}
	{
		p.SetState(203)
		p.List_assignment()
	}
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__4 {
		{
			p.SetState(204)
			p.Match(BistatParserT__4)
		}
		{
			p.SetState(205)
			p.List_assignment()
		}

		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(211)
		p.Match(BistatParserT__14)
	}

	return localctx
}

// IList_assignmentContext is an interface to support dynamic dispatch.
type IList_assignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVAR_CONS() []antlr.TerminalNode
	VAR_CONS(i int) antlr.TerminalNode

	// IsList_assignmentContext differentiates from other interfaces.
	IsList_assignmentContext()
}

type List_assignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyList_assignmentContext() *List_assignmentContext {
	var p = new(List_assignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_list_assignment
	return p
}

func (*List_assignmentContext) IsList_assignmentContext() {}

func NewList_assignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_assignmentContext {
	var p = new(List_assignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_list_assignment

	return p
}

func (s *List_assignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *List_assignmentContext) AllVAR_CONS() []antlr.TerminalNode {
	return s.GetTokens(BistatParserVAR_CONS)
}

func (s *List_assignmentContext) VAR_CONS(i int) antlr.TerminalNode {
	return s.GetToken(BistatParserVAR_CONS, i)
}

func (s *List_assignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_assignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_assignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterList_assignment(s)
	}
}

func (s *List_assignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitList_assignment(s)
	}
}

func (p *BistatParser) List_assignment() (localctx IList_assignmentContext) {
	this := p
	_ = this

	localctx = NewList_assignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, BistatParserRULE_list_assignment)
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
		p.SetState(213)
		p.Match(BistatParserT__13)
	}
	{
		p.SetState(214)
		p.Match(BistatParserVAR_CONS)
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__4 {
		{
			p.SetState(215)
			p.Match(BistatParserT__4)
		}
		{
			p.SetState(216)
			p.Match(BistatParserVAR_CONS)
		}

		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(222)
		p.Match(BistatParserT__14)
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
	p.EnterRule(localctx, 24, BistatParserRULE_comment)
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
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&576460752303357950) != 0) {
		{
			p.SetState(225)
			_la = p.GetTokenStream().LA(1)

			if _la <= 0 || _la == BistatParserT__15 {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(230)
		p.Match(BistatParserT__15)
	}

	return localctx
}

// IForContext is an interface to support dynamic dispatch.
type IForContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Expression() IExpressionContext
	AllNested_stmt() []INested_stmtContext
	Nested_stmt(i int) INested_stmtContext

	// IsForContext differentiates from other interfaces.
	IsForContext()
}

type ForContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForContext() *ForContext {
	var p = new(ForContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_for
	return p
}

func (*ForContext) IsForContext() {}

func NewForContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForContext {
	var p = new(ForContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_for

	return p
}

func (s *ForContext) GetParser() antlr.Parser { return s.parser }

func (s *ForContext) ID() antlr.TerminalNode {
	return s.GetToken(BistatParserID, 0)
}

func (s *ForContext) Expression() IExpressionContext {
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

func (s *ForContext) AllNested_stmt() []INested_stmtContext {
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

func (s *ForContext) Nested_stmt(i int) INested_stmtContext {
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

func (s *ForContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterFor(s)
	}
}

func (s *ForContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitFor(s)
	}
}

func (p *BistatParser) For_() (localctx IForContext) {
	this := p
	_ = this

	localctx = NewForContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, BistatParserRULE_for)
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
		p.SetState(232)
		p.Match(BistatParserT__16)
	}
	{
		p.SetState(233)
		p.Match(BistatParserT__3)
	}
	{
		p.SetState(234)
		p.Match(BistatParserID)
	}
	{
		p.SetState(235)
		p.Match(BistatParserT__17)
	}
	{
		p.SetState(236)
		p.Expression()
	}
	{
		p.SetState(237)
		p.Match(BistatParserT__5)
	}
	{
		p.SetState(238)
		p.Match(BistatParserT__7)
	}
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&149533578821640) != 0) {
		{
			p.SetState(239)
			p.Nested_stmt()
		}

		p.SetState(242)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(244)
		p.Match(BistatParserT__8)
	}

	return localctx
}

// IWhileContext is an interface to support dynamic dispatch.
type IWhileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsWhileContext differentiates from other interfaces.
	IsWhileContext()
}

type WhileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileContext() *WhileContext {
	var p = new(WhileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_while
	return p
}

func (*WhileContext) IsWhileContext() {}

func NewWhileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileContext {
	var p = new(WhileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_while

	return p
}

func (s *WhileContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileContext) Expression() IExpressionContext {
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

func (s *WhileContext) AllStmt() []IStmtContext {
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

func (s *WhileContext) Stmt(i int) IStmtContext {
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

func (s *WhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterWhile(s)
	}
}

func (s *WhileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitWhile(s)
	}
}

func (p *BistatParser) While() (localctx IWhileContext) {
	this := p
	_ = this

	localctx = NewWhileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, BistatParserRULE_while)
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
		p.SetState(246)
		p.Match(BistatParserT__18)
	}
	{
		p.SetState(247)
		p.Match(BistatParserT__3)
	}
	{
		p.SetState(248)
		p.Expression()
	}
	{
		p.SetState(249)
		p.Match(BistatParserT__5)
	}
	{
		p.SetState(250)
		p.Match(BistatParserT__7)
	}
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&149533578952712) != 0) {
		{
			p.SetState(251)
			p.Stmt()
		}

		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(256)
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
	p.EnterRule(localctx, 30, BistatParserRULE_conditional)
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
		p.SetState(258)
		p.Match(BistatParserT__19)
	}
	{
		p.SetState(259)
		p.Match(BistatParserT__3)
	}
	{
		p.SetState(260)
		p.Expression()
	}
	{
		p.SetState(261)
		p.Match(BistatParserT__5)
	}
	{
		p.SetState(262)
		p.Match(BistatParserT__7)
	}
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&149533578952712) != 0) {
		{
			p.SetState(263)
			p.Stmt()
		}

		p.SetState(266)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(268)
		p.Match(BistatParserT__8)
	}
	p.SetState(284)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(269)
				p.Match(BistatParserT__20)
			}
			{
				p.SetState(270)
				p.Match(BistatParserT__19)
			}
			{
				p.SetState(271)
				p.Match(BistatParserT__3)
			}
			{
				p.SetState(272)
				p.Expression()
			}
			{
				p.SetState(273)
				p.Match(BistatParserT__5)
			}
			{
				p.SetState(274)
				p.Match(BistatParserT__7)
			}
			p.SetState(276)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&149533578952712) != 0) {
				{
					p.SetState(275)
					p.Stmt()
				}

				p.SetState(278)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(280)
				p.Match(BistatParserT__8)
			}

		}
		p.SetState(286)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BistatParserT__20 {
		{
			p.SetState(287)
			p.Match(BistatParserT__20)
		}
		{
			p.SetState(288)
			p.Match(BistatParserT__7)
		}
		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&149533578952712) != 0) {
			{
				p.SetState(289)
				p.Stmt()
			}

			p.SetState(292)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(294)
			p.Match(BistatParserT__8)
		}

	}

	return localctx
}

// IInput_readContext is an interface to support dynamic dispatch.
type IInput_readContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode

	// IsInput_readContext differentiates from other interfaces.
	IsInput_readContext()
}

type Input_readContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInput_readContext() *Input_readContext {
	var p = new(Input_readContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_input_read
	return p
}

func (*Input_readContext) IsInput_readContext() {}

func NewInput_readContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Input_readContext {
	var p = new(Input_readContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_input_read

	return p
}

func (s *Input_readContext) GetParser() antlr.Parser { return s.parser }

func (s *Input_readContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(BistatParserID)
}

func (s *Input_readContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(BistatParserID, i)
}

func (s *Input_readContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Input_readContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Input_readContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterInput_read(s)
	}
}

func (s *Input_readContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitInput_read(s)
	}
}

func (p *BistatParser) Input_read() (localctx IInput_readContext) {
	this := p
	_ = this

	localctx = NewInput_readContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, BistatParserRULE_input_read)
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
		p.Match(BistatParserT__21)
	}
	{
		p.SetState(299)
		p.Match(BistatParserT__3)
	}
	{
		p.SetState(300)
		p.Match(BistatParserID)
	}
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BistatParserT__4 {
		{
			p.SetState(301)
			p.Match(BistatParserT__4)
		}
		{
			p.SetState(302)
			p.Match(BistatParserID)
		}

		p.SetState(307)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(308)
		p.Match(BistatParserT__5)
	}

	return localctx
}

// IList_addContext is an interface to support dynamic dispatch.
type IList_addContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsList_addContext differentiates from other interfaces.
	IsList_addContext()
}

type List_addContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyList_addContext() *List_addContext {
	var p = new(List_addContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_list_add
	return p
}

func (*List_addContext) IsList_addContext() {}

func NewList_addContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_addContext {
	var p = new(List_addContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_list_add

	return p
}

func (s *List_addContext) GetParser() antlr.Parser { return s.parser }

func (s *List_addContext) AllExpression() []IExpressionContext {
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

func (s *List_addContext) Expression(i int) IExpressionContext {
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

func (s *List_addContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_addContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_addContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterList_add(s)
	}
}

func (s *List_addContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitList_add(s)
	}
}

func (p *BistatParser) List_add() (localctx IList_addContext) {
	this := p
	_ = this

	localctx = NewList_addContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, BistatParserRULE_list_add)

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
		p.SetState(310)
		p.Match(BistatParserT__22)
	}
	{
		p.SetState(311)
		p.Match(BistatParserT__3)
	}
	{
		p.SetState(312)
		p.Expression()
	}
	{
		p.SetState(313)
		p.Match(BistatParserT__4)
	}
	{
		p.SetState(314)
		p.Expression()
	}
	{
		p.SetState(315)
		p.Match(BistatParserT__5)
	}

	return localctx
}

// IList_popContext is an interface to support dynamic dispatch.
type IList_popContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsList_popContext differentiates from other interfaces.
	IsList_popContext()
}

type List_popContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyList_popContext() *List_popContext {
	var p = new(List_popContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_list_pop
	return p
}

func (*List_popContext) IsList_popContext() {}

func NewList_popContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_popContext {
	var p = new(List_popContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_list_pop

	return p
}

func (s *List_popContext) GetParser() antlr.Parser { return s.parser }

func (s *List_popContext) Expression() IExpressionContext {
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

func (s *List_popContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_popContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_popContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterList_pop(s)
	}
}

func (s *List_popContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitList_pop(s)
	}
}

func (p *BistatParser) List_pop() (localctx IList_popContext) {
	this := p
	_ = this

	localctx = NewList_popContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, BistatParserRULE_list_pop)

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
		p.Match(BistatParserT__23)
	}
	{
		p.SetState(318)
		p.Match(BistatParserT__3)
	}
	{
		p.SetState(319)
		p.Expression()
	}
	{
		p.SetState(320)
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
		p.SetState(322)
		p.Match(BistatParserT__24)
	}
	{
		p.SetState(323)
		p.Match(BistatParserT__3)
	}
	{
		p.SetState(324)
		p.Expression()
	}
	{
		p.SetState(325)
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
		p.SetState(327)
		p.Match(BistatParserT__25)
	}
	{
		p.SetState(328)
		p.Match(BistatParserT__3)
	}
	{
		p.SetState(329)
		p.Expression()
	}
	p.SetState(332)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BistatParserT__4 {
		{
			p.SetState(330)
			p.Match(BistatParserT__4)
		}
		{
			p.SetState(331)
			p.Expression()
		}

	}
	{
		p.SetState(334)
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
		p.SetState(336)
		p.Match(BistatParserT__26)
	}
	{
		p.SetState(337)
		p.Match(BistatParserT__3)
	}
	{
		p.SetState(338)
		p.Expression()
	}
	{
		p.SetState(339)
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
		p.SetState(341)
		p.Match(BistatParserT__27)
	}
	{
		p.SetState(342)
		p.Match(BistatParserT__3)
	}
	{
		p.SetState(343)
		p.Expression()
	}
	{
		p.SetState(344)
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
		p.SetState(346)
		p.Match(BistatParserT__28)
	}
	{
		p.SetState(347)
		p.Match(BistatParserT__3)
	}
	{
		p.SetState(348)
		p.Expression()
	}
	{
		p.SetState(349)
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
		p.SetState(351)
		p.Match(BistatParserT__29)
	}
	{
		p.SetState(352)
		p.Match(BistatParserT__3)
	}
	{
		p.SetState(353)
		p.Expression()
	}
	{
		p.SetState(354)
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
		p.SetState(356)
		p.Match(BistatParserT__30)
	}
	{
		p.SetState(357)
		p.Match(BistatParserT__3)
	}
	{
		p.SetState(358)
		p.Expression()
	}
	{
		p.SetState(359)
		p.Match(BistatParserT__5)
	}

	return localctx
}

// IS_modeContext is an interface to support dynamic dispatch.
type IS_modeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsS_modeContext differentiates from other interfaces.
	IsS_modeContext()
}

type S_modeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyS_modeContext() *S_modeContext {
	var p = new(S_modeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_s_mode
	return p
}

func (*S_modeContext) IsS_modeContext() {}

func NewS_modeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *S_modeContext {
	var p = new(S_modeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_s_mode

	return p
}

func (s *S_modeContext) GetParser() antlr.Parser { return s.parser }

func (s *S_modeContext) Expression() IExpressionContext {
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

func (s *S_modeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_modeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *S_modeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterS_mode(s)
	}
}

func (s *S_modeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitS_mode(s)
	}
}

func (p *BistatParser) S_mode() (localctx IS_modeContext) {
	this := p
	_ = this

	localctx = NewS_modeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, BistatParserRULE_s_mode)

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
		p.Match(BistatParserT__31)
	}
	{
		p.SetState(362)
		p.Match(BistatParserT__3)
	}
	{
		p.SetState(363)
		p.Expression()
	}
	{
		p.SetState(364)
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
		p.SetState(366)
		p.Match(BistatParserT__32)
	}
	{
		p.SetState(367)
		p.Match(BistatParserT__3)
	}
	{
		p.SetState(368)
		p.Expression()
	}
	{
		p.SetState(369)
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
		p.SetState(371)
		p.Match(BistatParserT__33)
	}
	{
		p.SetState(372)
		p.Match(BistatParserT__3)
	}
	{
		p.SetState(373)
		p.Expression()
	}
	{
		p.SetState(374)
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
		p.SetState(376)
		p.Match(BistatParserT__34)
	}
	{
		p.SetState(377)
		p.Match(BistatParserT__3)
	}
	{
		p.SetState(378)
		p.Expression()
	}
	{
		p.SetState(379)
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
		p.SetState(381)
		p.Match(BistatParserT__35)
	}
	{
		p.SetState(382)
		p.Match(BistatParserT__3)
	}
	{
		p.SetState(383)
		p.Expression()
	}
	{
		p.SetState(384)
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
		p.SetState(386)
		p.Match(BistatParserT__36)
	}
	{
		p.SetState(387)
		p.Match(BistatParserT__3)
	}
	{
		p.SetState(388)
		p.Expression()
	}
	{
		p.SetState(389)
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
		p.SetState(391)
		p.Match(BistatParserT__37)
	}
	{
		p.SetState(392)
		p.Match(BistatParserT__3)
	}
	{
		p.SetState(393)
		p.Expression()
	}
	{
		p.SetState(394)
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
		p.SetState(396)
		p.Match(BistatParserT__38)
	}
	{
		p.SetState(397)
		p.Match(BistatParserT__3)
	}
	{
		p.SetState(398)
		p.Expression()
	}
	{
		p.SetState(399)
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
		p.SetState(401)
		p.Match(BistatParserT__39)
	}
	{
		p.SetState(402)
		p.Match(BistatParserT__3)
	}
	{
		p.SetState(403)
		p.Expression()
	}
	{
		p.SetState(404)
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
		p.SetState(406)
		p.Match(BistatParserT__40)
	}
	{
		p.SetState(407)
		p.Match(BistatParserT__3)
	}
	{
		p.SetState(408)
		p.Expression()
	}
	{
		p.SetState(409)
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
		p.SetState(411)
		p.Match(BistatParserT__41)
	}
	{
		p.SetState(412)
		p.Match(BistatParserT__3)
	}
	{
		p.SetState(413)
		p.Expression()
	}
	{
		p.SetState(414)
		p.Match(BistatParserT__5)
	}

	return localctx
}

// ISpecial_functionContext is an interface to support dynamic dispatch.
type ISpecial_functionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Input_read() IInput_readContext
	Print_() IPrintContext
	List_add() IList_addContext
	List_pop() IList_popContext
	Length() ILengthContext
	Range_() IRangeContext
	Plot() IPlotContext
	Sum() ISumContext
	Min() IMinContext
	Prod() IProdContext
	Avg() IAvgContext
	S_mode() IS_modeContext
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

	// IsSpecial_functionContext differentiates from other interfaces.
	IsSpecial_functionContext()
}

type Special_functionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpecial_functionContext() *Special_functionContext {
	var p = new(Special_functionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_special_function
	return p
}

func (*Special_functionContext) IsSpecial_functionContext() {}

func NewSpecial_functionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Special_functionContext {
	var p = new(Special_functionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_special_function

	return p
}

func (s *Special_functionContext) GetParser() antlr.Parser { return s.parser }

func (s *Special_functionContext) Input_read() IInput_readContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInput_readContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInput_readContext)
}

func (s *Special_functionContext) Print_() IPrintContext {
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

func (s *Special_functionContext) List_add() IList_addContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_addContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_addContext)
}

func (s *Special_functionContext) List_pop() IList_popContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_popContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_popContext)
}

func (s *Special_functionContext) Length() ILengthContext {
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

func (s *Special_functionContext) Range_() IRangeContext {
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

func (s *Special_functionContext) Plot() IPlotContext {
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

func (s *Special_functionContext) Sum() ISumContext {
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

func (s *Special_functionContext) Min() IMinContext {
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

func (s *Special_functionContext) Prod() IProdContext {
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

func (s *Special_functionContext) Avg() IAvgContext {
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

func (s *Special_functionContext) S_mode() IS_modeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IS_modeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IS_modeContext)
}

func (s *Special_functionContext) Median() IMedianContext {
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

func (s *Special_functionContext) Sin() ISinContext {
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

func (s *Special_functionContext) Cos() ICosContext {
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

func (s *Special_functionContext) Tan() ITanContext {
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

func (s *Special_functionContext) Sort() ISortContext {
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

func (s *Special_functionContext) Sqrt() ISqrtContext {
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

func (s *Special_functionContext) Floor() IFloorContext {
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

func (s *Special_functionContext) Ceil() ICeilContext {
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

func (s *Special_functionContext) Abs() IAbsContext {
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

func (s *Special_functionContext) Not() INotContext {
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

func (s *Special_functionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Special_functionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Special_functionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterSpecial_function(s)
	}
}

func (s *Special_functionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitSpecial_function(s)
	}
}

func (p *BistatParser) Special_function() (localctx ISpecial_functionContext) {
	this := p
	_ = this

	localctx = NewSpecial_functionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, BistatParserRULE_special_function)

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

	p.SetState(438)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BistatParserT__21:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(416)
			p.Input_read()
		}

	case BistatParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(417)
			p.Print_()
		}

	case BistatParserT__22:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(418)
			p.List_add()
		}

	case BistatParserT__23:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(419)
			p.List_pop()
		}

	case BistatParserT__24:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(420)
			p.Length()
		}

	case BistatParserT__25:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(421)
			p.Range_()
		}

	case BistatParserT__26:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(422)
			p.Plot()
		}

	case BistatParserT__27:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(423)
			p.Sum()
		}

	case BistatParserT__28:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(424)
			p.Min()
		}

	case BistatParserT__29:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(425)
			p.Prod()
		}

	case BistatParserT__30:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(426)
			p.Avg()
		}

	case BistatParserT__31:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(427)
			p.S_mode()
		}

	case BistatParserT__32:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(428)
			p.Median()
		}

	case BistatParserT__33:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(429)
			p.Sin()
		}

	case BistatParserT__35:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(430)
			p.Cos()
		}

	case BistatParserT__34:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(431)
			p.Tan()
		}

	case BistatParserT__36:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(432)
			p.Sort()
		}

	case BistatParserT__37:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(433)
			p.Sqrt()
		}

	case BistatParserT__38:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(434)
			p.Floor()
		}

	case BistatParserT__39:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(435)
			p.Ceil()
		}

	case BistatParserT__40:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(436)
			p.Abs()
		}

	case BistatParserT__41:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(437)
			p.Not()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
		p.SetState(440)
		p.Term()
	}
	p.SetState(443)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BistatParserOP_SEC {
		{
			p.SetState(441)
			p.Match(BistatParserOP_SEC)
		}
		{
			p.SetState(442)
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
		p.SetState(445)
		p.Factor()
	}
	p.SetState(448)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BistatParserOP_FIRST {
		{
			p.SetState(446)
			p.Match(BistatParserOP_FIRST)
		}
		{
			p.SetState(447)
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
	Special_function() ISpecial_functionContext
	Function_call() IFunction_callContext
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

func (s *FactorContext) Special_function() ISpecial_functionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpecial_functionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpecial_functionContext)
}

func (s *FactorContext) Function_call() IFunction_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
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
	p.SetState(451)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BistatParserOP_SEC {
		{
			p.SetState(450)
			p.Match(BistatParserOP_SEC)
		}

	}
	p.SetState(461)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(453)
			p.Match(BistatParserT__3)
		}
		{
			p.SetState(454)
			p.Expression()
		}
		{
			p.SetState(455)
			p.Match(BistatParserT__5)
		}

	case 2:
		{
			p.SetState(457)
			p.Indexing()
		}

	case 3:
		{
			p.SetState(458)
			p.Special_function()
		}

	case 4:
		{
			p.SetState(459)
			p.Function_call()
		}

	case 5:
		{
			p.SetState(460)
			p.Match(BistatParserVAR_CONS)
		}

	}

	return localctx
}

// IFunction_callContext is an interface to support dynamic dispatch.
type IFunction_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsFunction_callContext differentiates from other interfaces.
	IsFunction_callContext()
}

type Function_callContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_callContext() *Function_callContext {
	var p = new(Function_callContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BistatParserRULE_function_call
	return p
}

func (*Function_callContext) IsFunction_callContext() {}

func NewFunction_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_callContext {
	var p = new(Function_callContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BistatParserRULE_function_call

	return p
}

func (s *Function_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_callContext) ID() antlr.TerminalNode {
	return s.GetToken(BistatParserID, 0)
}

func (s *Function_callContext) AllExpression() []IExpressionContext {
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

func (s *Function_callContext) Expression(i int) IExpressionContext {
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

func (s *Function_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.EnterFunction_call(s)
	}
}

func (s *Function_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BistatListener); ok {
		listenerT.ExitFunction_call(s)
	}
}

func (p *BistatParser) Function_call() (localctx IFunction_callContext) {
	this := p
	_ = this

	localctx = NewFunction_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, BistatParserRULE_function_call)
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
		p.SetState(463)
		p.Match(BistatParserID)
	}
	{
		p.SetState(464)
		p.Match(BistatParserT__3)
	}
	p.SetState(473)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&37304230502989848) != 0 {
		{
			p.SetState(465)
			p.Expression()
		}
		p.SetState(470)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == BistatParserT__4 {
			{
				p.SetState(466)
				p.Match(BistatParserT__4)
			}
			{
				p.SetState(467)
				p.Expression()
			}

			p.SetState(472)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(475)
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
		p.SetState(477)
		p.Match(BistatParserID)
	}
	{
		p.SetState(478)
		p.Match(BistatParserT__13)
	}
	{
		p.SetState(479)
		p.Expression()
	}
	{
		p.SetState(480)
		p.Match(BistatParserT__14)
	}
	p.SetState(485)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BistatParserT__13 {
		{
			p.SetState(481)
			p.Match(BistatParserT__13)
		}
		{
			p.SetState(482)
			p.Expression()
		}
		{
			p.SetState(483)
			p.Match(BistatParserT__14)
		}

	}

	return localctx
}
