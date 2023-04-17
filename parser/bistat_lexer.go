// Code generated from parser/Bistat.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type BistatLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var bistatlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func bistatlexerLexerInit() {
	staticData := &bistatlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
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
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
		"T__33", "T__34", "T__35", "T__36", "T__37", "T__38", "T__39", "T__40",
		"T__41", "WS", "NON_VOID_TYPE", "RETURN_TYPE", "PARAM_TYPE", "Type_Primitive",
		"ID", "PARAM_CARDINALITY", "CARDINALITY", "VAR_CONS", "BOOL_CONS", "STRING_CONS",
		"FLOAT_CONS", "INT_CONS", "OP_SEC", "OP_FIRST", "LOGIC_OPERATOR", "Alpha",
		"NUMBER",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 58, 445, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38,
		1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1,
		41, 1, 41, 1, 41, 1, 41, 1, 42, 4, 42, 304, 8, 42, 11, 42, 12, 42, 305,
		1, 42, 1, 42, 1, 43, 1, 43, 3, 43, 312, 8, 43, 1, 44, 1, 44, 1, 44, 1,
		44, 1, 44, 3, 44, 319, 8, 44, 1, 45, 1, 45, 3, 45, 323, 8, 45, 1, 46, 1,
		46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46,
		1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 3, 46, 343, 8, 46, 1, 47, 1,
		47, 4, 47, 347, 8, 47, 11, 47, 12, 47, 348, 1, 47, 1, 47, 1, 47, 5, 47,
		354, 8, 47, 10, 47, 12, 47, 357, 9, 47, 1, 48, 1, 48, 1, 48, 1, 48, 1,
		48, 1, 48, 3, 48, 365, 8, 48, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49,
		1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 3, 49, 378, 8, 49, 1, 50, 1, 50, 1,
		50, 1, 50, 1, 50, 3, 50, 385, 8, 50, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51,
		1, 51, 1, 51, 1, 51, 1, 51, 3, 51, 396, 8, 51, 1, 52, 1, 52, 5, 52, 400,
		8, 52, 10, 52, 12, 52, 403, 9, 52, 1, 52, 1, 52, 1, 53, 4, 53, 408, 8,
		53, 11, 53, 12, 53, 409, 1, 53, 1, 53, 4, 53, 414, 8, 53, 11, 53, 12, 53,
		415, 1, 54, 4, 54, 419, 8, 54, 11, 54, 12, 54, 420, 1, 55, 1, 55, 1, 56,
		1, 56, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1,
		57, 1, 57, 1, 57, 1, 57, 3, 57, 440, 8, 57, 1, 58, 1, 58, 1, 59, 1, 59,
		0, 0, 60, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19,
		10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37,
		19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55,
		28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73,
		37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91,
		46, 93, 0, 95, 47, 97, 48, 99, 49, 101, 50, 103, 51, 105, 52, 107, 53,
		109, 54, 111, 55, 113, 56, 115, 57, 117, 0, 119, 58, 1, 0, 6, 3, 0, 9,
		10, 13, 13, 32, 32, 1, 0, 34, 34, 2, 0, 43, 43, 45, 45, 3, 0, 37, 37, 42,
		42, 47, 47, 2, 0, 60, 60, 62, 62, 2, 0, 65, 90, 97, 122, 471, 0, 1, 1,
		0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1,
		0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17,
		1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0,
		25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0,
		0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0,
		0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0,
		0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1,
		0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63,
		1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0,
		71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0,
		0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0,
		0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 95, 1, 0,
		0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103,
		1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0,
		0, 111, 1, 0, 0, 0, 0, 113, 1, 0, 0, 0, 0, 115, 1, 0, 0, 0, 0, 119, 1,
		0, 0, 0, 1, 121, 1, 0, 0, 0, 3, 129, 1, 0, 0, 0, 5, 131, 1, 0, 0, 0, 7,
		135, 1, 0, 0, 0, 9, 140, 1, 0, 0, 0, 11, 142, 1, 0, 0, 0, 13, 144, 1, 0,
		0, 0, 15, 146, 1, 0, 0, 0, 17, 148, 1, 0, 0, 0, 19, 150, 1, 0, 0, 0, 21,
		155, 1, 0, 0, 0, 23, 157, 1, 0, 0, 0, 25, 159, 1, 0, 0, 0, 27, 161, 1,
		0, 0, 0, 29, 163, 1, 0, 0, 0, 31, 165, 1, 0, 0, 0, 33, 169, 1, 0, 0, 0,
		35, 172, 1, 0, 0, 0, 37, 178, 1, 0, 0, 0, 39, 181, 1, 0, 0, 0, 41, 186,
		1, 0, 0, 0, 43, 191, 1, 0, 0, 0, 45, 197, 1, 0, 0, 0, 47, 205, 1, 0, 0,
		0, 49, 213, 1, 0, 0, 0, 51, 220, 1, 0, 0, 0, 53, 226, 1, 0, 0, 0, 55, 231,
		1, 0, 0, 0, 57, 235, 1, 0, 0, 0, 59, 239, 1, 0, 0, 0, 61, 244, 1, 0, 0,
		0, 63, 248, 1, 0, 0, 0, 65, 254, 1, 0, 0, 0, 67, 261, 1, 0, 0, 0, 69, 265,
		1, 0, 0, 0, 71, 269, 1, 0, 0, 0, 73, 273, 1, 0, 0, 0, 75, 278, 1, 0, 0,
		0, 77, 283, 1, 0, 0, 0, 79, 289, 1, 0, 0, 0, 81, 294, 1, 0, 0, 0, 83, 298,
		1, 0, 0, 0, 85, 303, 1, 0, 0, 0, 87, 309, 1, 0, 0, 0, 89, 318, 1, 0, 0,
		0, 91, 320, 1, 0, 0, 0, 93, 342, 1, 0, 0, 0, 95, 346, 1, 0, 0, 0, 97, 364,
		1, 0, 0, 0, 99, 377, 1, 0, 0, 0, 101, 384, 1, 0, 0, 0, 103, 395, 1, 0,
		0, 0, 105, 397, 1, 0, 0, 0, 107, 407, 1, 0, 0, 0, 109, 418, 1, 0, 0, 0,
		111, 422, 1, 0, 0, 0, 113, 424, 1, 0, 0, 0, 115, 439, 1, 0, 0, 0, 117,
		441, 1, 0, 0, 0, 119, 443, 1, 0, 0, 0, 121, 122, 5, 80, 0, 0, 122, 123,
		5, 114, 0, 0, 123, 124, 5, 111, 0, 0, 124, 125, 5, 103, 0, 0, 125, 126,
		5, 114, 0, 0, 126, 127, 5, 97, 0, 0, 127, 128, 5, 109, 0, 0, 128, 2, 1,
		0, 0, 0, 129, 130, 5, 59, 0, 0, 130, 4, 1, 0, 0, 0, 131, 132, 5, 118, 0,
		0, 132, 133, 5, 97, 0, 0, 133, 134, 5, 114, 0, 0, 134, 6, 1, 0, 0, 0, 135,
		136, 5, 102, 0, 0, 136, 137, 5, 117, 0, 0, 137, 138, 5, 110, 0, 0, 138,
		139, 5, 99, 0, 0, 139, 8, 1, 0, 0, 0, 140, 141, 5, 40, 0, 0, 141, 10, 1,
		0, 0, 0, 142, 143, 5, 41, 0, 0, 143, 12, 1, 0, 0, 0, 144, 145, 5, 58, 0,
		0, 145, 14, 1, 0, 0, 0, 146, 147, 5, 123, 0, 0, 147, 16, 1, 0, 0, 0, 148,
		149, 5, 125, 0, 0, 149, 18, 1, 0, 0, 0, 150, 151, 5, 109, 0, 0, 151, 152,
		5, 97, 0, 0, 152, 153, 5, 105, 0, 0, 153, 154, 5, 110, 0, 0, 154, 20, 1,
		0, 0, 0, 155, 156, 5, 61, 0, 0, 156, 22, 1, 0, 0, 0, 157, 158, 5, 91, 0,
		0, 158, 24, 1, 0, 0, 0, 159, 160, 5, 44, 0, 0, 160, 26, 1, 0, 0, 0, 161,
		162, 5, 93, 0, 0, 162, 28, 1, 0, 0, 0, 163, 164, 5, 35, 0, 0, 164, 30,
		1, 0, 0, 0, 165, 166, 5, 102, 0, 0, 166, 167, 5, 111, 0, 0, 167, 168, 5,
		114, 0, 0, 168, 32, 1, 0, 0, 0, 169, 170, 5, 105, 0, 0, 170, 171, 5, 110,
		0, 0, 171, 34, 1, 0, 0, 0, 172, 173, 5, 119, 0, 0, 173, 174, 5, 104, 0,
		0, 174, 175, 5, 105, 0, 0, 175, 176, 5, 108, 0, 0, 176, 177, 5, 101, 0,
		0, 177, 36, 1, 0, 0, 0, 178, 179, 5, 105, 0, 0, 179, 180, 5, 102, 0, 0,
		180, 38, 1, 0, 0, 0, 181, 182, 5, 101, 0, 0, 182, 183, 5, 108, 0, 0, 183,
		184, 5, 115, 0, 0, 184, 185, 5, 101, 0, 0, 185, 40, 1, 0, 0, 0, 186, 187,
		5, 114, 0, 0, 187, 188, 5, 101, 0, 0, 188, 189, 5, 97, 0, 0, 189, 190,
		5, 100, 0, 0, 190, 42, 1, 0, 0, 0, 191, 192, 5, 112, 0, 0, 192, 193, 5,
		114, 0, 0, 193, 194, 5, 105, 0, 0, 194, 195, 5, 110, 0, 0, 195, 196, 5,
		116, 0, 0, 196, 44, 1, 0, 0, 0, 197, 198, 5, 108, 0, 0, 198, 199, 5, 105,
		0, 0, 199, 200, 5, 115, 0, 0, 200, 201, 5, 116, 0, 0, 201, 202, 5, 65,
		0, 0, 202, 203, 5, 100, 0, 0, 203, 204, 5, 100, 0, 0, 204, 46, 1, 0, 0,
		0, 205, 206, 5, 108, 0, 0, 206, 207, 5, 105, 0, 0, 207, 208, 5, 115, 0,
		0, 208, 209, 5, 116, 0, 0, 209, 210, 5, 80, 0, 0, 210, 211, 5, 111, 0,
		0, 211, 212, 5, 112, 0, 0, 212, 48, 1, 0, 0, 0, 213, 214, 5, 108, 0, 0,
		214, 215, 5, 101, 0, 0, 215, 216, 5, 110, 0, 0, 216, 217, 5, 103, 0, 0,
		217, 218, 5, 116, 0, 0, 218, 219, 5, 104, 0, 0, 219, 50, 1, 0, 0, 0, 220,
		221, 5, 114, 0, 0, 221, 222, 5, 97, 0, 0, 222, 223, 5, 110, 0, 0, 223,
		224, 5, 103, 0, 0, 224, 225, 5, 101, 0, 0, 225, 52, 1, 0, 0, 0, 226, 227,
		5, 112, 0, 0, 227, 228, 5, 108, 0, 0, 228, 229, 5, 111, 0, 0, 229, 230,
		5, 116, 0, 0, 230, 54, 1, 0, 0, 0, 231, 232, 5, 115, 0, 0, 232, 233, 5,
		117, 0, 0, 233, 234, 5, 109, 0, 0, 234, 56, 1, 0, 0, 0, 235, 236, 5, 109,
		0, 0, 236, 237, 5, 105, 0, 0, 237, 238, 5, 110, 0, 0, 238, 58, 1, 0, 0,
		0, 239, 240, 5, 112, 0, 0, 240, 241, 5, 114, 0, 0, 241, 242, 5, 111, 0,
		0, 242, 243, 5, 100, 0, 0, 243, 60, 1, 0, 0, 0, 244, 245, 5, 97, 0, 0,
		245, 246, 5, 118, 0, 0, 246, 247, 5, 103, 0, 0, 247, 62, 1, 0, 0, 0, 248,
		249, 5, 115, 0, 0, 249, 250, 5, 77, 0, 0, 250, 251, 5, 111, 0, 0, 251,
		252, 5, 100, 0, 0, 252, 253, 5, 101, 0, 0, 253, 64, 1, 0, 0, 0, 254, 255,
		5, 109, 0, 0, 255, 256, 5, 101, 0, 0, 256, 257, 5, 100, 0, 0, 257, 258,
		5, 105, 0, 0, 258, 259, 5, 97, 0, 0, 259, 260, 5, 110, 0, 0, 260, 66, 1,
		0, 0, 0, 261, 262, 5, 115, 0, 0, 262, 263, 5, 105, 0, 0, 263, 264, 5, 110,
		0, 0, 264, 68, 1, 0, 0, 0, 265, 266, 5, 116, 0, 0, 266, 267, 5, 97, 0,
		0, 267, 268, 5, 110, 0, 0, 268, 70, 1, 0, 0, 0, 269, 270, 5, 99, 0, 0,
		270, 271, 5, 111, 0, 0, 271, 272, 5, 115, 0, 0, 272, 72, 1, 0, 0, 0, 273,
		274, 5, 115, 0, 0, 274, 275, 5, 111, 0, 0, 275, 276, 5, 114, 0, 0, 276,
		277, 5, 116, 0, 0, 277, 74, 1, 0, 0, 0, 278, 279, 5, 115, 0, 0, 279, 280,
		5, 113, 0, 0, 280, 281, 5, 114, 0, 0, 281, 282, 5, 116, 0, 0, 282, 76,
		1, 0, 0, 0, 283, 284, 5, 102, 0, 0, 284, 285, 5, 108, 0, 0, 285, 286, 5,
		111, 0, 0, 286, 287, 5, 111, 0, 0, 287, 288, 5, 114, 0, 0, 288, 78, 1,
		0, 0, 0, 289, 290, 5, 99, 0, 0, 290, 291, 5, 101, 0, 0, 291, 292, 5, 105,
		0, 0, 292, 293, 5, 108, 0, 0, 293, 80, 1, 0, 0, 0, 294, 295, 5, 97, 0,
		0, 295, 296, 5, 98, 0, 0, 296, 297, 5, 115, 0, 0, 297, 82, 1, 0, 0, 0,
		298, 299, 5, 110, 0, 0, 299, 300, 5, 111, 0, 0, 300, 301, 5, 116, 0, 0,
		301, 84, 1, 0, 0, 0, 302, 304, 7, 0, 0, 0, 303, 302, 1, 0, 0, 0, 304, 305,
		1, 0, 0, 0, 305, 303, 1, 0, 0, 0, 305, 306, 1, 0, 0, 0, 306, 307, 1, 0,
		0, 0, 307, 308, 6, 42, 0, 0, 308, 86, 1, 0, 0, 0, 309, 311, 3, 93, 46,
		0, 310, 312, 3, 99, 49, 0, 311, 310, 1, 0, 0, 0, 311, 312, 1, 0, 0, 0,
		312, 88, 1, 0, 0, 0, 313, 319, 3, 87, 43, 0, 314, 315, 5, 118, 0, 0, 315,
		316, 5, 111, 0, 0, 316, 317, 5, 105, 0, 0, 317, 319, 5, 100, 0, 0, 318,
		313, 1, 0, 0, 0, 318, 314, 1, 0, 0, 0, 319, 90, 1, 0, 0, 0, 320, 322, 3,
		93, 46, 0, 321, 323, 3, 97, 48, 0, 322, 321, 1, 0, 0, 0, 322, 323, 1, 0,
		0, 0, 323, 92, 1, 0, 0, 0, 324, 325, 5, 105, 0, 0, 325, 326, 5, 110, 0,
		0, 326, 343, 5, 116, 0, 0, 327, 328, 5, 102, 0, 0, 328, 329, 5, 108, 0,
		0, 329, 330, 5, 111, 0, 0, 330, 331, 5, 97, 0, 0, 331, 343, 5, 116, 0,
		0, 332, 333, 5, 115, 0, 0, 333, 334, 5, 116, 0, 0, 334, 335, 5, 114, 0,
		0, 335, 336, 5, 105, 0, 0, 336, 337, 5, 110, 0, 0, 337, 343, 5, 103, 0,
		0, 338, 339, 5, 98, 0, 0, 339, 340, 5, 111, 0, 0, 340, 341, 5, 111, 0,
		0, 341, 343, 5, 108, 0, 0, 342, 324, 1, 0, 0, 0, 342, 327, 1, 0, 0, 0,
		342, 332, 1, 0, 0, 0, 342, 338, 1, 0, 0, 0, 343, 94, 1, 0, 0, 0, 344, 347,
		5, 95, 0, 0, 345, 347, 3, 117, 58, 0, 346, 344, 1, 0, 0, 0, 346, 345, 1,
		0, 0, 0, 347, 348, 1, 0, 0, 0, 348, 346, 1, 0, 0, 0, 348, 349, 1, 0, 0,
		0, 349, 355, 1, 0, 0, 0, 350, 354, 3, 117, 58, 0, 351, 354, 3, 119, 59,
		0, 352, 354, 5, 95, 0, 0, 353, 350, 1, 0, 0, 0, 353, 351, 1, 0, 0, 0, 353,
		352, 1, 0, 0, 0, 354, 357, 1, 0, 0, 0, 355, 353, 1, 0, 0, 0, 355, 356,
		1, 0, 0, 0, 356, 96, 1, 0, 0, 0, 357, 355, 1, 0, 0, 0, 358, 359, 5, 91,
		0, 0, 359, 365, 5, 93, 0, 0, 360, 361, 5, 91, 0, 0, 361, 362, 5, 93, 0,
		0, 362, 363, 5, 91, 0, 0, 363, 365, 5, 93, 0, 0, 364, 358, 1, 0, 0, 0,
		364, 360, 1, 0, 0, 0, 365, 98, 1, 0, 0, 0, 366, 367, 5, 91, 0, 0, 367,
		368, 3, 109, 54, 0, 368, 369, 5, 93, 0, 0, 369, 378, 1, 0, 0, 0, 370, 371,
		5, 91, 0, 0, 371, 372, 3, 109, 54, 0, 372, 373, 5, 93, 0, 0, 373, 374,
		5, 91, 0, 0, 374, 375, 3, 109, 54, 0, 375, 376, 5, 93, 0, 0, 376, 378,
		1, 0, 0, 0, 377, 366, 1, 0, 0, 0, 377, 370, 1, 0, 0, 0, 378, 100, 1, 0,
		0, 0, 379, 385, 3, 105, 52, 0, 380, 385, 3, 107, 53, 0, 381, 385, 3, 109,
		54, 0, 382, 385, 3, 103, 51, 0, 383, 385, 3, 95, 47, 0, 384, 379, 1, 0,
		0, 0, 384, 380, 1, 0, 0, 0, 384, 381, 1, 0, 0, 0, 384, 382, 1, 0, 0, 0,
		384, 383, 1, 0, 0, 0, 385, 102, 1, 0, 0, 0, 386, 387, 5, 116, 0, 0, 387,
		388, 5, 114, 0, 0, 388, 389, 5, 117, 0, 0, 389, 396, 5, 101, 0, 0, 390,
		391, 5, 102, 0, 0, 391, 392, 5, 97, 0, 0, 392, 393, 5, 108, 0, 0, 393,
		394, 5, 115, 0, 0, 394, 396, 5, 101, 0, 0, 395, 386, 1, 0, 0, 0, 395, 390,
		1, 0, 0, 0, 396, 104, 1, 0, 0, 0, 397, 401, 5, 34, 0, 0, 398, 400, 8, 1,
		0, 0, 399, 398, 1, 0, 0, 0, 400, 403, 1, 0, 0, 0, 401, 399, 1, 0, 0, 0,
		401, 402, 1, 0, 0, 0, 402, 404, 1, 0, 0, 0, 403, 401, 1, 0, 0, 0, 404,
		405, 5, 34, 0, 0, 405, 106, 1, 0, 0, 0, 406, 408, 3, 119, 59, 0, 407, 406,
		1, 0, 0, 0, 408, 409, 1, 0, 0, 0, 409, 407, 1, 0, 0, 0, 409, 410, 1, 0,
		0, 0, 410, 411, 1, 0, 0, 0, 411, 413, 5, 46, 0, 0, 412, 414, 3, 119, 59,
		0, 413, 412, 1, 0, 0, 0, 414, 415, 1, 0, 0, 0, 415, 413, 1, 0, 0, 0, 415,
		416, 1, 0, 0, 0, 416, 108, 1, 0, 0, 0, 417, 419, 3, 119, 59, 0, 418, 417,
		1, 0, 0, 0, 419, 420, 1, 0, 0, 0, 420, 418, 1, 0, 0, 0, 420, 421, 1, 0,
		0, 0, 421, 110, 1, 0, 0, 0, 422, 423, 7, 2, 0, 0, 423, 112, 1, 0, 0, 0,
		424, 425, 7, 3, 0, 0, 425, 114, 1, 0, 0, 0, 426, 440, 7, 4, 0, 0, 427,
		428, 5, 60, 0, 0, 428, 440, 5, 61, 0, 0, 429, 430, 5, 62, 0, 0, 430, 440,
		5, 61, 0, 0, 431, 432, 5, 61, 0, 0, 432, 440, 5, 61, 0, 0, 433, 434, 5,
		33, 0, 0, 434, 440, 5, 61, 0, 0, 435, 436, 5, 38, 0, 0, 436, 440, 5, 38,
		0, 0, 437, 438, 5, 124, 0, 0, 438, 440, 5, 124, 0, 0, 439, 426, 1, 0, 0,
		0, 439, 427, 1, 0, 0, 0, 439, 429, 1, 0, 0, 0, 439, 431, 1, 0, 0, 0, 439,
		433, 1, 0, 0, 0, 439, 435, 1, 0, 0, 0, 439, 437, 1, 0, 0, 0, 440, 116,
		1, 0, 0, 0, 441, 442, 7, 5, 0, 0, 442, 118, 1, 0, 0, 0, 443, 444, 2, 48,
		57, 0, 444, 120, 1, 0, 0, 0, 19, 0, 305, 311, 318, 322, 342, 346, 348,
		353, 355, 364, 377, 384, 395, 401, 409, 415, 420, 439, 1, 0, 1, 0,
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

// BistatLexerInit initializes any static state used to implement BistatLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewBistatLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func BistatLexerInit() {
	staticData := &bistatlexerLexerStaticData
	staticData.once.Do(bistatlexerLexerInit)
}

// NewBistatLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewBistatLexer(input antlr.CharStream) *BistatLexer {
	BistatLexerInit()
	l := new(BistatLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &bistatlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Bistat.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// BistatLexer tokens.
const (
	BistatLexerT__0              = 1
	BistatLexerT__1              = 2
	BistatLexerT__2              = 3
	BistatLexerT__3              = 4
	BistatLexerT__4              = 5
	BistatLexerT__5              = 6
	BistatLexerT__6              = 7
	BistatLexerT__7              = 8
	BistatLexerT__8              = 9
	BistatLexerT__9              = 10
	BistatLexerT__10             = 11
	BistatLexerT__11             = 12
	BistatLexerT__12             = 13
	BistatLexerT__13             = 14
	BistatLexerT__14             = 15
	BistatLexerT__15             = 16
	BistatLexerT__16             = 17
	BistatLexerT__17             = 18
	BistatLexerT__18             = 19
	BistatLexerT__19             = 20
	BistatLexerT__20             = 21
	BistatLexerT__21             = 22
	BistatLexerT__22             = 23
	BistatLexerT__23             = 24
	BistatLexerT__24             = 25
	BistatLexerT__25             = 26
	BistatLexerT__26             = 27
	BistatLexerT__27             = 28
	BistatLexerT__28             = 29
	BistatLexerT__29             = 30
	BistatLexerT__30             = 31
	BistatLexerT__31             = 32
	BistatLexerT__32             = 33
	BistatLexerT__33             = 34
	BistatLexerT__34             = 35
	BistatLexerT__35             = 36
	BistatLexerT__36             = 37
	BistatLexerT__37             = 38
	BistatLexerT__38             = 39
	BistatLexerT__39             = 40
	BistatLexerT__40             = 41
	BistatLexerT__41             = 42
	BistatLexerWS                = 43
	BistatLexerNON_VOID_TYPE     = 44
	BistatLexerRETURN_TYPE       = 45
	BistatLexerPARAM_TYPE        = 46
	BistatLexerID                = 47
	BistatLexerPARAM_CARDINALITY = 48
	BistatLexerCARDINALITY       = 49
	BistatLexerVAR_CONS          = 50
	BistatLexerBOOL_CONS         = 51
	BistatLexerSTRING_CONS       = 52
	BistatLexerFLOAT_CONS        = 53
	BistatLexerINT_CONS          = 54
	BistatLexerOP_SEC            = 55
	BistatLexerOP_FIRST          = 56
	BistatLexerLOGIC_OPERATOR    = 57
	BistatLexerNUMBER            = 58
)
