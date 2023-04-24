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
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
		"T__33", "T__34", "T__35", "T__36", "T__37", "T__38", "T__39", "T__40",
		"T__41", "T__42", "WS", "CARDINALITY", "TYPE_PRIMITIVE", "ID", "VAR_CONS",
		"INT_CONS", "NUMBER", "BOOL_CONS", "STRING_CONS", "FLOAT_CONS", "OP_SEC",
		"OP_FIRST", "LOGIC_OPERATOR", "Alpha",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 56, 427, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1,
		11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1,
		33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39,
		1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1,
		41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 4, 43, 305, 8, 43, 11, 43, 12, 43,
		306, 1, 43, 1, 43, 1, 44, 1, 44, 3, 44, 313, 8, 44, 1, 44, 1, 44, 1, 44,
		1, 44, 3, 44, 319, 8, 44, 1, 44, 3, 44, 322, 8, 44, 1, 45, 1, 45, 1, 45,
		1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1,
		45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 3, 45, 346,
		8, 45, 1, 46, 1, 46, 4, 46, 350, 8, 46, 11, 46, 12, 46, 351, 1, 46, 1,
		46, 1, 46, 5, 46, 357, 8, 46, 10, 46, 12, 46, 360, 9, 46, 1, 47, 1, 47,
		1, 47, 1, 47, 1, 47, 3, 47, 367, 8, 47, 1, 48, 4, 48, 370, 8, 48, 11, 48,
		12, 48, 371, 1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1,
		50, 1, 50, 1, 50, 3, 50, 385, 8, 50, 1, 51, 1, 51, 5, 51, 389, 8, 51, 10,
		51, 12, 51, 392, 9, 51, 1, 51, 1, 51, 1, 52, 4, 52, 397, 8, 52, 11, 52,
		12, 52, 398, 1, 52, 1, 52, 4, 52, 403, 8, 52, 11, 52, 12, 52, 404, 1, 53,
		1, 53, 1, 54, 1, 54, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1,
		55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 3, 55, 424, 8, 55, 1, 56, 1, 56,
		0, 0, 57, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19,
		10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37,
		19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55,
		28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73,
		37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91,
		46, 93, 47, 95, 48, 97, 49, 99, 50, 101, 51, 103, 52, 105, 53, 107, 54,
		109, 55, 111, 56, 113, 0, 1, 0, 6, 3, 0, 9, 10, 13, 13, 32, 32, 1, 0, 34,
		34, 2, 0, 43, 43, 45, 45, 3, 0, 37, 37, 42, 42, 47, 47, 2, 0, 60, 60, 62,
		62, 2, 0, 65, 90, 97, 122, 453, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0,
		5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0,
		13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0,
		0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0,
		0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0,
		0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1,
		0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51,
		1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0,
		59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0,
		0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0,
		0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0,
		0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1,
		0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97,
		1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0,
		0, 105, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 111, 1,
		0, 0, 0, 1, 115, 1, 0, 0, 0, 3, 123, 1, 0, 0, 0, 5, 125, 1, 0, 0, 0, 7,
		129, 1, 0, 0, 0, 9, 134, 1, 0, 0, 0, 11, 136, 1, 0, 0, 0, 13, 138, 1, 0,
		0, 0, 15, 140, 1, 0, 0, 0, 17, 142, 1, 0, 0, 0, 19, 144, 1, 0, 0, 0, 21,
		149, 1, 0, 0, 0, 23, 151, 1, 0, 0, 0, 25, 153, 1, 0, 0, 0, 27, 155, 1,
		0, 0, 0, 29, 157, 1, 0, 0, 0, 31, 159, 1, 0, 0, 0, 33, 163, 1, 0, 0, 0,
		35, 166, 1, 0, 0, 0, 37, 172, 1, 0, 0, 0, 39, 175, 1, 0, 0, 0, 41, 180,
		1, 0, 0, 0, 43, 187, 1, 0, 0, 0, 45, 192, 1, 0, 0, 0, 47, 198, 1, 0, 0,
		0, 49, 206, 1, 0, 0, 0, 51, 214, 1, 0, 0, 0, 53, 221, 1, 0, 0, 0, 55, 227,
		1, 0, 0, 0, 57, 232, 1, 0, 0, 0, 59, 236, 1, 0, 0, 0, 61, 240, 1, 0, 0,
		0, 63, 245, 1, 0, 0, 0, 65, 249, 1, 0, 0, 0, 67, 255, 1, 0, 0, 0, 69, 262,
		1, 0, 0, 0, 71, 266, 1, 0, 0, 0, 73, 270, 1, 0, 0, 0, 75, 274, 1, 0, 0,
		0, 77, 279, 1, 0, 0, 0, 79, 284, 1, 0, 0, 0, 81, 290, 1, 0, 0, 0, 83, 295,
		1, 0, 0, 0, 85, 299, 1, 0, 0, 0, 87, 304, 1, 0, 0, 0, 89, 310, 1, 0, 0,
		0, 91, 345, 1, 0, 0, 0, 93, 349, 1, 0, 0, 0, 95, 366, 1, 0, 0, 0, 97, 369,
		1, 0, 0, 0, 99, 373, 1, 0, 0, 0, 101, 384, 1, 0, 0, 0, 103, 386, 1, 0,
		0, 0, 105, 396, 1, 0, 0, 0, 107, 406, 1, 0, 0, 0, 109, 408, 1, 0, 0, 0,
		111, 423, 1, 0, 0, 0, 113, 425, 1, 0, 0, 0, 115, 116, 5, 80, 0, 0, 116,
		117, 5, 114, 0, 0, 117, 118, 5, 111, 0, 0, 118, 119, 5, 103, 0, 0, 119,
		120, 5, 114, 0, 0, 120, 121, 5, 97, 0, 0, 121, 122, 5, 109, 0, 0, 122,
		2, 1, 0, 0, 0, 123, 124, 5, 59, 0, 0, 124, 4, 1, 0, 0, 0, 125, 126, 5,
		118, 0, 0, 126, 127, 5, 97, 0, 0, 127, 128, 5, 114, 0, 0, 128, 6, 1, 0,
		0, 0, 129, 130, 5, 102, 0, 0, 130, 131, 5, 117, 0, 0, 131, 132, 5, 110,
		0, 0, 132, 133, 5, 99, 0, 0, 133, 8, 1, 0, 0, 0, 134, 135, 5, 40, 0, 0,
		135, 10, 1, 0, 0, 0, 136, 137, 5, 41, 0, 0, 137, 12, 1, 0, 0, 0, 138, 139,
		5, 58, 0, 0, 139, 14, 1, 0, 0, 0, 140, 141, 5, 123, 0, 0, 141, 16, 1, 0,
		0, 0, 142, 143, 5, 125, 0, 0, 143, 18, 1, 0, 0, 0, 144, 145, 5, 109, 0,
		0, 145, 146, 5, 97, 0, 0, 146, 147, 5, 105, 0, 0, 147, 148, 5, 110, 0,
		0, 148, 20, 1, 0, 0, 0, 149, 150, 5, 61, 0, 0, 150, 22, 1, 0, 0, 0, 151,
		152, 5, 91, 0, 0, 152, 24, 1, 0, 0, 0, 153, 154, 5, 44, 0, 0, 154, 26,
		1, 0, 0, 0, 155, 156, 5, 93, 0, 0, 156, 28, 1, 0, 0, 0, 157, 158, 5, 35,
		0, 0, 158, 30, 1, 0, 0, 0, 159, 160, 5, 102, 0, 0, 160, 161, 5, 111, 0,
		0, 161, 162, 5, 114, 0, 0, 162, 32, 1, 0, 0, 0, 163, 164, 5, 105, 0, 0,
		164, 165, 5, 110, 0, 0, 165, 34, 1, 0, 0, 0, 166, 167, 5, 119, 0, 0, 167,
		168, 5, 104, 0, 0, 168, 169, 5, 105, 0, 0, 169, 170, 5, 108, 0, 0, 170,
		171, 5, 101, 0, 0, 171, 36, 1, 0, 0, 0, 172, 173, 5, 105, 0, 0, 173, 174,
		5, 102, 0, 0, 174, 38, 1, 0, 0, 0, 175, 176, 5, 101, 0, 0, 176, 177, 5,
		108, 0, 0, 177, 178, 5, 115, 0, 0, 178, 179, 5, 101, 0, 0, 179, 40, 1,
		0, 0, 0, 180, 181, 5, 114, 0, 0, 181, 182, 5, 101, 0, 0, 182, 183, 5, 116,
		0, 0, 183, 184, 5, 117, 0, 0, 184, 185, 5, 114, 0, 0, 185, 186, 5, 110,
		0, 0, 186, 42, 1, 0, 0, 0, 187, 188, 5, 114, 0, 0, 188, 189, 5, 101, 0,
		0, 189, 190, 5, 97, 0, 0, 190, 191, 5, 100, 0, 0, 191, 44, 1, 0, 0, 0,
		192, 193, 5, 112, 0, 0, 193, 194, 5, 114, 0, 0, 194, 195, 5, 105, 0, 0,
		195, 196, 5, 110, 0, 0, 196, 197, 5, 116, 0, 0, 197, 46, 1, 0, 0, 0, 198,
		199, 5, 108, 0, 0, 199, 200, 5, 105, 0, 0, 200, 201, 5, 115, 0, 0, 201,
		202, 5, 116, 0, 0, 202, 203, 5, 65, 0, 0, 203, 204, 5, 100, 0, 0, 204,
		205, 5, 100, 0, 0, 205, 48, 1, 0, 0, 0, 206, 207, 5, 108, 0, 0, 207, 208,
		5, 105, 0, 0, 208, 209, 5, 115, 0, 0, 209, 210, 5, 116, 0, 0, 210, 211,
		5, 80, 0, 0, 211, 212, 5, 111, 0, 0, 212, 213, 5, 112, 0, 0, 213, 50, 1,
		0, 0, 0, 214, 215, 5, 108, 0, 0, 215, 216, 5, 101, 0, 0, 216, 217, 5, 110,
		0, 0, 217, 218, 5, 103, 0, 0, 218, 219, 5, 116, 0, 0, 219, 220, 5, 104,
		0, 0, 220, 52, 1, 0, 0, 0, 221, 222, 5, 114, 0, 0, 222, 223, 5, 97, 0,
		0, 223, 224, 5, 110, 0, 0, 224, 225, 5, 103, 0, 0, 225, 226, 5, 101, 0,
		0, 226, 54, 1, 0, 0, 0, 227, 228, 5, 112, 0, 0, 228, 229, 5, 108, 0, 0,
		229, 230, 5, 111, 0, 0, 230, 231, 5, 116, 0, 0, 231, 56, 1, 0, 0, 0, 232,
		233, 5, 115, 0, 0, 233, 234, 5, 117, 0, 0, 234, 235, 5, 109, 0, 0, 235,
		58, 1, 0, 0, 0, 236, 237, 5, 109, 0, 0, 237, 238, 5, 105, 0, 0, 238, 239,
		5, 110, 0, 0, 239, 60, 1, 0, 0, 0, 240, 241, 5, 112, 0, 0, 241, 242, 5,
		114, 0, 0, 242, 243, 5, 111, 0, 0, 243, 244, 5, 100, 0, 0, 244, 62, 1,
		0, 0, 0, 245, 246, 5, 97, 0, 0, 246, 247, 5, 118, 0, 0, 247, 248, 5, 103,
		0, 0, 248, 64, 1, 0, 0, 0, 249, 250, 5, 115, 0, 0, 250, 251, 5, 77, 0,
		0, 251, 252, 5, 111, 0, 0, 252, 253, 5, 100, 0, 0, 253, 254, 5, 101, 0,
		0, 254, 66, 1, 0, 0, 0, 255, 256, 5, 109, 0, 0, 256, 257, 5, 101, 0, 0,
		257, 258, 5, 100, 0, 0, 258, 259, 5, 105, 0, 0, 259, 260, 5, 97, 0, 0,
		260, 261, 5, 110, 0, 0, 261, 68, 1, 0, 0, 0, 262, 263, 5, 115, 0, 0, 263,
		264, 5, 105, 0, 0, 264, 265, 5, 110, 0, 0, 265, 70, 1, 0, 0, 0, 266, 267,
		5, 116, 0, 0, 267, 268, 5, 97, 0, 0, 268, 269, 5, 110, 0, 0, 269, 72, 1,
		0, 0, 0, 270, 271, 5, 99, 0, 0, 271, 272, 5, 111, 0, 0, 272, 273, 5, 115,
		0, 0, 273, 74, 1, 0, 0, 0, 274, 275, 5, 115, 0, 0, 275, 276, 5, 111, 0,
		0, 276, 277, 5, 114, 0, 0, 277, 278, 5, 116, 0, 0, 278, 76, 1, 0, 0, 0,
		279, 280, 5, 115, 0, 0, 280, 281, 5, 113, 0, 0, 281, 282, 5, 114, 0, 0,
		282, 283, 5, 116, 0, 0, 283, 78, 1, 0, 0, 0, 284, 285, 5, 102, 0, 0, 285,
		286, 5, 108, 0, 0, 286, 287, 5, 111, 0, 0, 287, 288, 5, 111, 0, 0, 288,
		289, 5, 114, 0, 0, 289, 80, 1, 0, 0, 0, 290, 291, 5, 99, 0, 0, 291, 292,
		5, 101, 0, 0, 292, 293, 5, 105, 0, 0, 293, 294, 5, 108, 0, 0, 294, 82,
		1, 0, 0, 0, 295, 296, 5, 97, 0, 0, 296, 297, 5, 98, 0, 0, 297, 298, 5,
		115, 0, 0, 298, 84, 1, 0, 0, 0, 299, 300, 5, 110, 0, 0, 300, 301, 5, 111,
		0, 0, 301, 302, 5, 116, 0, 0, 302, 86, 1, 0, 0, 0, 303, 305, 7, 0, 0, 0,
		304, 303, 1, 0, 0, 0, 305, 306, 1, 0, 0, 0, 306, 304, 1, 0, 0, 0, 306,
		307, 1, 0, 0, 0, 307, 308, 1, 0, 0, 0, 308, 309, 6, 43, 0, 0, 309, 88,
		1, 0, 0, 0, 310, 312, 5, 91, 0, 0, 311, 313, 3, 97, 48, 0, 312, 311, 1,
		0, 0, 0, 312, 313, 1, 0, 0, 0, 313, 314, 1, 0, 0, 0, 314, 315, 5, 93, 0,
		0, 315, 321, 1, 0, 0, 0, 316, 318, 5, 91, 0, 0, 317, 319, 3, 97, 48, 0,
		318, 317, 1, 0, 0, 0, 318, 319, 1, 0, 0, 0, 319, 320, 1, 0, 0, 0, 320,
		322, 5, 93, 0, 0, 321, 316, 1, 0, 0, 0, 321, 322, 1, 0, 0, 0, 322, 90,
		1, 0, 0, 0, 323, 324, 5, 105, 0, 0, 324, 325, 5, 110, 0, 0, 325, 346, 5,
		116, 0, 0, 326, 327, 5, 102, 0, 0, 327, 328, 5, 108, 0, 0, 328, 329, 5,
		111, 0, 0, 329, 330, 5, 97, 0, 0, 330, 346, 5, 116, 0, 0, 331, 332, 5,
		115, 0, 0, 332, 333, 5, 116, 0, 0, 333, 334, 5, 114, 0, 0, 334, 335, 5,
		105, 0, 0, 335, 336, 5, 110, 0, 0, 336, 346, 5, 103, 0, 0, 337, 338, 5,
		98, 0, 0, 338, 339, 5, 111, 0, 0, 339, 340, 5, 111, 0, 0, 340, 346, 5,
		108, 0, 0, 341, 342, 5, 118, 0, 0, 342, 343, 5, 111, 0, 0, 343, 344, 5,
		105, 0, 0, 344, 346, 5, 100, 0, 0, 345, 323, 1, 0, 0, 0, 345, 326, 1, 0,
		0, 0, 345, 331, 1, 0, 0, 0, 345, 337, 1, 0, 0, 0, 345, 341, 1, 0, 0, 0,
		346, 92, 1, 0, 0, 0, 347, 350, 5, 95, 0, 0, 348, 350, 3, 113, 56, 0, 349,
		347, 1, 0, 0, 0, 349, 348, 1, 0, 0, 0, 350, 351, 1, 0, 0, 0, 351, 349,
		1, 0, 0, 0, 351, 352, 1, 0, 0, 0, 352, 358, 1, 0, 0, 0, 353, 357, 3, 113,
		56, 0, 354, 357, 3, 99, 49, 0, 355, 357, 5, 95, 0, 0, 356, 353, 1, 0, 0,
		0, 356, 354, 1, 0, 0, 0, 356, 355, 1, 0, 0, 0, 357, 360, 1, 0, 0, 0, 358,
		356, 1, 0, 0, 0, 358, 359, 1, 0, 0, 0, 359, 94, 1, 0, 0, 0, 360, 358, 1,
		0, 0, 0, 361, 367, 3, 103, 51, 0, 362, 367, 3, 105, 52, 0, 363, 367, 3,
		97, 48, 0, 364, 367, 3, 101, 50, 0, 365, 367, 3, 93, 46, 0, 366, 361, 1,
		0, 0, 0, 366, 362, 1, 0, 0, 0, 366, 363, 1, 0, 0, 0, 366, 364, 1, 0, 0,
		0, 366, 365, 1, 0, 0, 0, 367, 96, 1, 0, 0, 0, 368, 370, 3, 99, 49, 0, 369,
		368, 1, 0, 0, 0, 370, 371, 1, 0, 0, 0, 371, 369, 1, 0, 0, 0, 371, 372,
		1, 0, 0, 0, 372, 98, 1, 0, 0, 0, 373, 374, 2, 48, 57, 0, 374, 100, 1, 0,
		0, 0, 375, 376, 5, 116, 0, 0, 376, 377, 5, 114, 0, 0, 377, 378, 5, 117,
		0, 0, 378, 385, 5, 101, 0, 0, 379, 380, 5, 102, 0, 0, 380, 381, 5, 97,
		0, 0, 381, 382, 5, 108, 0, 0, 382, 383, 5, 115, 0, 0, 383, 385, 5, 101,
		0, 0, 384, 375, 1, 0, 0, 0, 384, 379, 1, 0, 0, 0, 385, 102, 1, 0, 0, 0,
		386, 390, 5, 34, 0, 0, 387, 389, 8, 1, 0, 0, 388, 387, 1, 0, 0, 0, 389,
		392, 1, 0, 0, 0, 390, 388, 1, 0, 0, 0, 390, 391, 1, 0, 0, 0, 391, 393,
		1, 0, 0, 0, 392, 390, 1, 0, 0, 0, 393, 394, 5, 34, 0, 0, 394, 104, 1, 0,
		0, 0, 395, 397, 3, 99, 49, 0, 396, 395, 1, 0, 0, 0, 397, 398, 1, 0, 0,
		0, 398, 396, 1, 0, 0, 0, 398, 399, 1, 0, 0, 0, 399, 400, 1, 0, 0, 0, 400,
		402, 5, 46, 0, 0, 401, 403, 3, 99, 49, 0, 402, 401, 1, 0, 0, 0, 403, 404,
		1, 0, 0, 0, 404, 402, 1, 0, 0, 0, 404, 405, 1, 0, 0, 0, 405, 106, 1, 0,
		0, 0, 406, 407, 7, 2, 0, 0, 407, 108, 1, 0, 0, 0, 408, 409, 7, 3, 0, 0,
		409, 110, 1, 0, 0, 0, 410, 424, 7, 4, 0, 0, 411, 412, 5, 60, 0, 0, 412,
		424, 5, 61, 0, 0, 413, 414, 5, 62, 0, 0, 414, 424, 5, 61, 0, 0, 415, 416,
		5, 61, 0, 0, 416, 424, 5, 61, 0, 0, 417, 418, 5, 33, 0, 0, 418, 424, 5,
		61, 0, 0, 419, 420, 5, 38, 0, 0, 420, 424, 5, 38, 0, 0, 421, 422, 5, 124,
		0, 0, 422, 424, 5, 124, 0, 0, 423, 410, 1, 0, 0, 0, 423, 411, 1, 0, 0,
		0, 423, 413, 1, 0, 0, 0, 423, 415, 1, 0, 0, 0, 423, 417, 1, 0, 0, 0, 423,
		419, 1, 0, 0, 0, 423, 421, 1, 0, 0, 0, 424, 112, 1, 0, 0, 0, 425, 426,
		7, 5, 0, 0, 426, 114, 1, 0, 0, 0, 17, 0, 306, 312, 318, 321, 345, 349,
		351, 356, 358, 366, 371, 384, 390, 398, 404, 423, 1, 0, 1, 0,
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
	BistatLexerT__0           = 1
	BistatLexerT__1           = 2
	BistatLexerT__2           = 3
	BistatLexerT__3           = 4
	BistatLexerT__4           = 5
	BistatLexerT__5           = 6
	BistatLexerT__6           = 7
	BistatLexerT__7           = 8
	BistatLexerT__8           = 9
	BistatLexerT__9           = 10
	BistatLexerT__10          = 11
	BistatLexerT__11          = 12
	BistatLexerT__12          = 13
	BistatLexerT__13          = 14
	BistatLexerT__14          = 15
	BistatLexerT__15          = 16
	BistatLexerT__16          = 17
	BistatLexerT__17          = 18
	BistatLexerT__18          = 19
	BistatLexerT__19          = 20
	BistatLexerT__20          = 21
	BistatLexerT__21          = 22
	BistatLexerT__22          = 23
	BistatLexerT__23          = 24
	BistatLexerT__24          = 25
	BistatLexerT__25          = 26
	BistatLexerT__26          = 27
	BistatLexerT__27          = 28
	BistatLexerT__28          = 29
	BistatLexerT__29          = 30
	BistatLexerT__30          = 31
	BistatLexerT__31          = 32
	BistatLexerT__32          = 33
	BistatLexerT__33          = 34
	BistatLexerT__34          = 35
	BistatLexerT__35          = 36
	BistatLexerT__36          = 37
	BistatLexerT__37          = 38
	BistatLexerT__38          = 39
	BistatLexerT__39          = 40
	BistatLexerT__40          = 41
	BistatLexerT__41          = 42
	BistatLexerT__42          = 43
	BistatLexerWS             = 44
	BistatLexerCARDINALITY    = 45
	BistatLexerTYPE_PRIMITIVE = 46
	BistatLexerID             = 47
	BistatLexerVAR_CONS       = 48
	BistatLexerINT_CONS       = 49
	BistatLexerNUMBER         = 50
	BistatLexerBOOL_CONS      = 51
	BistatLexerSTRING_CONS    = 52
	BistatLexerFLOAT_CONS     = 53
	BistatLexerOP_SEC         = 54
	BistatLexerOP_FIRST       = 55
	BistatLexerLOGIC_OPERATOR = 56
)
