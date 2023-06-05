// Code generated from parser/Bistat.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
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

var BistatLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func bistatlexerLexerInit() {
	staticData := &BistatLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
		"T__33", "T__34", "T__35", "T__36", "T__37", "T__38", "T__39", "T__40",
		"T__41", "T__42", "T__43", "T__44", "T__45", "T__46", "T__47", "T__48",
		"T__49", "T__50", "T__51", "WS", "CARDINALITY", "TYPE_PRIMITIVE", "BOOL_CONS",
		"ID", "INT_CONS", "NUMBER", "STRING_CONS", "FLOAT_CONS", "Alpha",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 61, 409, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
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
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1,
		7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11,
		1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18,
		1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1,
		33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41,
		1, 41, 1, 42, 1, 42, 1, 43, 1, 43, 1, 44, 1, 44, 1, 45, 1, 45, 1, 46, 1,
		46, 1, 46, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1, 48, 1, 49, 1, 49, 1, 49,
		1, 50, 1, 50, 1, 50, 1, 51, 1, 51, 1, 51, 1, 52, 4, 52, 317, 8, 52, 11,
		52, 12, 52, 318, 1, 52, 1, 52, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53,
		1, 53, 3, 53, 330, 8, 53, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1,
		54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54,
		1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 3, 54, 354, 8, 54, 1, 55, 1, 55, 1,
		55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 3, 55, 365, 8, 55, 1, 56,
		1, 56, 4, 56, 369, 8, 56, 11, 56, 12, 56, 370, 1, 56, 1, 56, 1, 56, 5,
		56, 376, 8, 56, 10, 56, 12, 56, 379, 9, 56, 1, 57, 4, 57, 382, 8, 57, 11,
		57, 12, 57, 383, 1, 58, 1, 58, 1, 59, 1, 59, 5, 59, 390, 8, 59, 10, 59,
		12, 59, 393, 9, 59, 1, 59, 1, 59, 1, 60, 4, 60, 398, 8, 60, 11, 60, 12,
		60, 399, 1, 60, 1, 60, 4, 60, 404, 8, 60, 11, 60, 12, 60, 405, 1, 61, 1,
		61, 0, 0, 62, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9,
		19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18,
		37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27,
		55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36,
		73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45,
		91, 46, 93, 47, 95, 48, 97, 49, 99, 50, 101, 51, 103, 52, 105, 53, 107,
		54, 109, 55, 111, 56, 113, 57, 115, 58, 117, 59, 119, 60, 121, 61, 123,
		0, 1, 0, 3, 3, 0, 9, 10, 13, 13, 32, 32, 1, 0, 34, 34, 2, 0, 65, 90, 97,
		122, 423, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0,
		23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0,
		0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0,
		0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0,
		0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1,
		0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61,
		1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0,
		69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0,
		0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0,
		0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0,
		0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1,
		0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0,
		107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 111, 1, 0, 0, 0, 0, 113, 1, 0,
		0, 0, 0, 115, 1, 0, 0, 0, 0, 117, 1, 0, 0, 0, 0, 119, 1, 0, 0, 0, 0, 121,
		1, 0, 0, 0, 1, 125, 1, 0, 0, 0, 3, 133, 1, 0, 0, 0, 5, 135, 1, 0, 0, 0,
		7, 139, 1, 0, 0, 0, 9, 144, 1, 0, 0, 0, 11, 146, 1, 0, 0, 0, 13, 148, 1,
		0, 0, 0, 15, 150, 1, 0, 0, 0, 17, 152, 1, 0, 0, 0, 19, 154, 1, 0, 0, 0,
		21, 159, 1, 0, 0, 0, 23, 161, 1, 0, 0, 0, 25, 163, 1, 0, 0, 0, 27, 165,
		1, 0, 0, 0, 29, 167, 1, 0, 0, 0, 31, 169, 1, 0, 0, 0, 33, 173, 1, 0, 0,
		0, 35, 176, 1, 0, 0, 0, 37, 182, 1, 0, 0, 0, 39, 185, 1, 0, 0, 0, 41, 190,
		1, 0, 0, 0, 43, 197, 1, 0, 0, 0, 45, 202, 1, 0, 0, 0, 47, 208, 1, 0, 0,
		0, 49, 213, 1, 0, 0, 0, 51, 217, 1, 0, 0, 0, 53, 221, 1, 0, 0, 0, 55, 225,
		1, 0, 0, 0, 57, 230, 1, 0, 0, 0, 59, 234, 1, 0, 0, 0, 61, 240, 1, 0, 0,
		0, 63, 247, 1, 0, 0, 0, 65, 251, 1, 0, 0, 0, 67, 255, 1, 0, 0, 0, 69, 259,
		1, 0, 0, 0, 71, 264, 1, 0, 0, 0, 73, 270, 1, 0, 0, 0, 75, 275, 1, 0, 0,
		0, 77, 279, 1, 0, 0, 0, 79, 283, 1, 0, 0, 0, 81, 285, 1, 0, 0, 0, 83, 287,
		1, 0, 0, 0, 85, 289, 1, 0, 0, 0, 87, 291, 1, 0, 0, 0, 89, 293, 1, 0, 0,
		0, 91, 295, 1, 0, 0, 0, 93, 297, 1, 0, 0, 0, 95, 300, 1, 0, 0, 0, 97, 303,
		1, 0, 0, 0, 99, 306, 1, 0, 0, 0, 101, 309, 1, 0, 0, 0, 103, 312, 1, 0,
		0, 0, 105, 316, 1, 0, 0, 0, 107, 322, 1, 0, 0, 0, 109, 353, 1, 0, 0, 0,
		111, 364, 1, 0, 0, 0, 113, 368, 1, 0, 0, 0, 115, 381, 1, 0, 0, 0, 117,
		385, 1, 0, 0, 0, 119, 387, 1, 0, 0, 0, 121, 397, 1, 0, 0, 0, 123, 407,
		1, 0, 0, 0, 125, 126, 5, 80, 0, 0, 126, 127, 5, 114, 0, 0, 127, 128, 5,
		111, 0, 0, 128, 129, 5, 103, 0, 0, 129, 130, 5, 114, 0, 0, 130, 131, 5,
		97, 0, 0, 131, 132, 5, 109, 0, 0, 132, 2, 1, 0, 0, 0, 133, 134, 5, 59,
		0, 0, 134, 4, 1, 0, 0, 0, 135, 136, 5, 118, 0, 0, 136, 137, 5, 97, 0, 0,
		137, 138, 5, 114, 0, 0, 138, 6, 1, 0, 0, 0, 139, 140, 5, 102, 0, 0, 140,
		141, 5, 117, 0, 0, 141, 142, 5, 110, 0, 0, 142, 143, 5, 99, 0, 0, 143,
		8, 1, 0, 0, 0, 144, 145, 5, 40, 0, 0, 145, 10, 1, 0, 0, 0, 146, 147, 5,
		41, 0, 0, 147, 12, 1, 0, 0, 0, 148, 149, 5, 58, 0, 0, 149, 14, 1, 0, 0,
		0, 150, 151, 5, 123, 0, 0, 151, 16, 1, 0, 0, 0, 152, 153, 5, 125, 0, 0,
		153, 18, 1, 0, 0, 0, 154, 155, 5, 109, 0, 0, 155, 156, 5, 97, 0, 0, 156,
		157, 5, 105, 0, 0, 157, 158, 5, 110, 0, 0, 158, 20, 1, 0, 0, 0, 159, 160,
		5, 61, 0, 0, 160, 22, 1, 0, 0, 0, 161, 162, 5, 91, 0, 0, 162, 24, 1, 0,
		0, 0, 163, 164, 5, 44, 0, 0, 164, 26, 1, 0, 0, 0, 165, 166, 5, 93, 0, 0,
		166, 28, 1, 0, 0, 0, 167, 168, 5, 35, 0, 0, 168, 30, 1, 0, 0, 0, 169, 170,
		5, 102, 0, 0, 170, 171, 5, 111, 0, 0, 171, 172, 5, 114, 0, 0, 172, 32,
		1, 0, 0, 0, 173, 174, 5, 105, 0, 0, 174, 175, 5, 110, 0, 0, 175, 34, 1,
		0, 0, 0, 176, 177, 5, 119, 0, 0, 177, 178, 5, 104, 0, 0, 178, 179, 5, 105,
		0, 0, 179, 180, 5, 108, 0, 0, 180, 181, 5, 101, 0, 0, 181, 36, 1, 0, 0,
		0, 182, 183, 5, 105, 0, 0, 183, 184, 5, 102, 0, 0, 184, 38, 1, 0, 0, 0,
		185, 186, 5, 101, 0, 0, 186, 187, 5, 108, 0, 0, 187, 188, 5, 115, 0, 0,
		188, 189, 5, 101, 0, 0, 189, 40, 1, 0, 0, 0, 190, 191, 5, 114, 0, 0, 191,
		192, 5, 101, 0, 0, 192, 193, 5, 116, 0, 0, 193, 194, 5, 117, 0, 0, 194,
		195, 5, 114, 0, 0, 195, 196, 5, 110, 0, 0, 196, 42, 1, 0, 0, 0, 197, 198,
		5, 114, 0, 0, 198, 199, 5, 101, 0, 0, 199, 200, 5, 97, 0, 0, 200, 201,
		5, 100, 0, 0, 201, 44, 1, 0, 0, 0, 202, 203, 5, 112, 0, 0, 203, 204, 5,
		114, 0, 0, 204, 205, 5, 105, 0, 0, 205, 206, 5, 110, 0, 0, 206, 207, 5,
		116, 0, 0, 207, 46, 1, 0, 0, 0, 208, 209, 5, 112, 0, 0, 209, 210, 5, 108,
		0, 0, 210, 211, 5, 111, 0, 0, 211, 212, 5, 116, 0, 0, 212, 48, 1, 0, 0,
		0, 213, 214, 5, 115, 0, 0, 214, 215, 5, 117, 0, 0, 215, 216, 5, 109, 0,
		0, 216, 50, 1, 0, 0, 0, 217, 218, 5, 109, 0, 0, 218, 219, 5, 105, 0, 0,
		219, 220, 5, 110, 0, 0, 220, 52, 1, 0, 0, 0, 221, 222, 5, 109, 0, 0, 222,
		223, 5, 97, 0, 0, 223, 224, 5, 120, 0, 0, 224, 54, 1, 0, 0, 0, 225, 226,
		5, 112, 0, 0, 226, 227, 5, 114, 0, 0, 227, 228, 5, 111, 0, 0, 228, 229,
		5, 100, 0, 0, 229, 56, 1, 0, 0, 0, 230, 231, 5, 97, 0, 0, 231, 232, 5,
		118, 0, 0, 232, 233, 5, 103, 0, 0, 233, 58, 1, 0, 0, 0, 234, 235, 5, 115,
		0, 0, 235, 236, 5, 77, 0, 0, 236, 237, 5, 111, 0, 0, 237, 238, 5, 100,
		0, 0, 238, 239, 5, 101, 0, 0, 239, 60, 1, 0, 0, 0, 240, 241, 5, 109, 0,
		0, 241, 242, 5, 101, 0, 0, 242, 243, 5, 100, 0, 0, 243, 244, 5, 105, 0,
		0, 244, 245, 5, 97, 0, 0, 245, 246, 5, 110, 0, 0, 246, 62, 1, 0, 0, 0,
		247, 248, 5, 115, 0, 0, 248, 249, 5, 105, 0, 0, 249, 250, 5, 110, 0, 0,
		250, 64, 1, 0, 0, 0, 251, 252, 5, 116, 0, 0, 252, 253, 5, 97, 0, 0, 253,
		254, 5, 110, 0, 0, 254, 66, 1, 0, 0, 0, 255, 256, 5, 99, 0, 0, 256, 257,
		5, 111, 0, 0, 257, 258, 5, 115, 0, 0, 258, 68, 1, 0, 0, 0, 259, 260, 5,
		115, 0, 0, 260, 261, 5, 113, 0, 0, 261, 262, 5, 114, 0, 0, 262, 263, 5,
		116, 0, 0, 263, 70, 1, 0, 0, 0, 264, 265, 5, 102, 0, 0, 265, 266, 5, 108,
		0, 0, 266, 267, 5, 111, 0, 0, 267, 268, 5, 111, 0, 0, 268, 269, 5, 114,
		0, 0, 269, 72, 1, 0, 0, 0, 270, 271, 5, 99, 0, 0, 271, 272, 5, 101, 0,
		0, 272, 273, 5, 105, 0, 0, 273, 274, 5, 108, 0, 0, 274, 74, 1, 0, 0, 0,
		275, 276, 5, 97, 0, 0, 276, 277, 5, 98, 0, 0, 277, 278, 5, 115, 0, 0, 278,
		76, 1, 0, 0, 0, 279, 280, 5, 110, 0, 0, 280, 281, 5, 111, 0, 0, 281, 282,
		5, 116, 0, 0, 282, 78, 1, 0, 0, 0, 283, 284, 5, 45, 0, 0, 284, 80, 1, 0,
		0, 0, 285, 286, 5, 43, 0, 0, 286, 82, 1, 0, 0, 0, 287, 288, 5, 47, 0, 0,
		288, 84, 1, 0, 0, 0, 289, 290, 5, 42, 0, 0, 290, 86, 1, 0, 0, 0, 291, 292,
		5, 37, 0, 0, 292, 88, 1, 0, 0, 0, 293, 294, 5, 60, 0, 0, 294, 90, 1, 0,
		0, 0, 295, 296, 5, 62, 0, 0, 296, 92, 1, 0, 0, 0, 297, 298, 5, 60, 0, 0,
		298, 299, 5, 61, 0, 0, 299, 94, 1, 0, 0, 0, 300, 301, 5, 62, 0, 0, 301,
		302, 5, 61, 0, 0, 302, 96, 1, 0, 0, 0, 303, 304, 5, 61, 0, 0, 304, 305,
		5, 61, 0, 0, 305, 98, 1, 0, 0, 0, 306, 307, 5, 33, 0, 0, 307, 308, 5, 61,
		0, 0, 308, 100, 1, 0, 0, 0, 309, 310, 5, 38, 0, 0, 310, 311, 5, 38, 0,
		0, 311, 102, 1, 0, 0, 0, 312, 313, 5, 124, 0, 0, 313, 314, 5, 124, 0, 0,
		314, 104, 1, 0, 0, 0, 315, 317, 7, 0, 0, 0, 316, 315, 1, 0, 0, 0, 317,
		318, 1, 0, 0, 0, 318, 316, 1, 0, 0, 0, 318, 319, 1, 0, 0, 0, 319, 320,
		1, 0, 0, 0, 320, 321, 6, 52, 0, 0, 321, 106, 1, 0, 0, 0, 322, 323, 5, 91,
		0, 0, 323, 324, 3, 115, 57, 0, 324, 329, 5, 93, 0, 0, 325, 326, 5, 91,
		0, 0, 326, 327, 3, 115, 57, 0, 327, 328, 5, 93, 0, 0, 328, 330, 1, 0, 0,
		0, 329, 325, 1, 0, 0, 0, 329, 330, 1, 0, 0, 0, 330, 108, 1, 0, 0, 0, 331,
		332, 5, 105, 0, 0, 332, 333, 5, 110, 0, 0, 333, 354, 5, 116, 0, 0, 334,
		335, 5, 102, 0, 0, 335, 336, 5, 108, 0, 0, 336, 337, 5, 111, 0, 0, 337,
		338, 5, 97, 0, 0, 338, 354, 5, 116, 0, 0, 339, 340, 5, 115, 0, 0, 340,
		341, 5, 116, 0, 0, 341, 342, 5, 114, 0, 0, 342, 343, 5, 105, 0, 0, 343,
		344, 5, 110, 0, 0, 344, 354, 5, 103, 0, 0, 345, 346, 5, 98, 0, 0, 346,
		347, 5, 111, 0, 0, 347, 348, 5, 111, 0, 0, 348, 354, 5, 108, 0, 0, 349,
		350, 5, 118, 0, 0, 350, 351, 5, 111, 0, 0, 351, 352, 5, 105, 0, 0, 352,
		354, 5, 100, 0, 0, 353, 331, 1, 0, 0, 0, 353, 334, 1, 0, 0, 0, 353, 339,
		1, 0, 0, 0, 353, 345, 1, 0, 0, 0, 353, 349, 1, 0, 0, 0, 354, 110, 1, 0,
		0, 0, 355, 356, 5, 116, 0, 0, 356, 357, 5, 114, 0, 0, 357, 358, 5, 117,
		0, 0, 358, 365, 5, 101, 0, 0, 359, 360, 5, 102, 0, 0, 360, 361, 5, 97,
		0, 0, 361, 362, 5, 108, 0, 0, 362, 363, 5, 115, 0, 0, 363, 365, 5, 101,
		0, 0, 364, 355, 1, 0, 0, 0, 364, 359, 1, 0, 0, 0, 365, 112, 1, 0, 0, 0,
		366, 369, 5, 95, 0, 0, 367, 369, 3, 123, 61, 0, 368, 366, 1, 0, 0, 0, 368,
		367, 1, 0, 0, 0, 369, 370, 1, 0, 0, 0, 370, 368, 1, 0, 0, 0, 370, 371,
		1, 0, 0, 0, 371, 377, 1, 0, 0, 0, 372, 376, 3, 123, 61, 0, 373, 376, 3,
		117, 58, 0, 374, 376, 5, 95, 0, 0, 375, 372, 1, 0, 0, 0, 375, 373, 1, 0,
		0, 0, 375, 374, 1, 0, 0, 0, 376, 379, 1, 0, 0, 0, 377, 375, 1, 0, 0, 0,
		377, 378, 1, 0, 0, 0, 378, 114, 1, 0, 0, 0, 379, 377, 1, 0, 0, 0, 380,
		382, 3, 117, 58, 0, 381, 380, 1, 0, 0, 0, 382, 383, 1, 0, 0, 0, 383, 381,
		1, 0, 0, 0, 383, 384, 1, 0, 0, 0, 384, 116, 1, 0, 0, 0, 385, 386, 2, 48,
		57, 0, 386, 118, 1, 0, 0, 0, 387, 391, 5, 34, 0, 0, 388, 390, 8, 1, 0,
		0, 389, 388, 1, 0, 0, 0, 390, 393, 1, 0, 0, 0, 391, 389, 1, 0, 0, 0, 391,
		392, 1, 0, 0, 0, 392, 394, 1, 0, 0, 0, 393, 391, 1, 0, 0, 0, 394, 395,
		5, 34, 0, 0, 395, 120, 1, 0, 0, 0, 396, 398, 3, 117, 58, 0, 397, 396, 1,
		0, 0, 0, 398, 399, 1, 0, 0, 0, 399, 397, 1, 0, 0, 0, 399, 400, 1, 0, 0,
		0, 400, 401, 1, 0, 0, 0, 401, 403, 5, 46, 0, 0, 402, 404, 3, 117, 58, 0,
		403, 402, 1, 0, 0, 0, 404, 405, 1, 0, 0, 0, 405, 403, 1, 0, 0, 0, 405,
		406, 1, 0, 0, 0, 406, 122, 1, 0, 0, 0, 407, 408, 7, 2, 0, 0, 408, 124,
		1, 0, 0, 0, 13, 0, 318, 329, 353, 364, 368, 370, 375, 377, 383, 391, 399,
		405, 1, 0, 1, 0,
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
	staticData := &BistatLexerLexerStaticData
	staticData.once.Do(bistatlexerLexerInit)
}

// NewBistatLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewBistatLexer(input antlr.CharStream) *BistatLexer {
	BistatLexerInit()
	l := new(BistatLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &BistatLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
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
	BistatLexerT__43          = 44
	BistatLexerT__44          = 45
	BistatLexerT__45          = 46
	BistatLexerT__46          = 47
	BistatLexerT__47          = 48
	BistatLexerT__48          = 49
	BistatLexerT__49          = 50
	BistatLexerT__50          = 51
	BistatLexerT__51          = 52
	BistatLexerWS             = 53
	BistatLexerCARDINALITY    = 54
	BistatLexerTYPE_PRIMITIVE = 55
	BistatLexerBOOL_CONS      = 56
	BistatLexerID             = 57
	BistatLexerINT_CONS       = 58
	BistatLexerNUMBER         = 59
	BistatLexerSTRING_CONS    = 60
	BistatLexerFLOAT_CONS     = 61
)
