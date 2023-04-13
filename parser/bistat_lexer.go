// Code generated from Bistat.g4 by ANTLR 4.12.0. DO NOT EDIT.

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
		"", "'Program'", "';'",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 2, 15, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 0, 0, 2, 1, 1, 3, 2, 1, 0, 0, 14, 0, 1,
		1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 1, 5, 1, 0, 0, 0, 3, 13, 1, 0, 0, 0, 5, 6,
		5, 80, 0, 0, 6, 7, 5, 114, 0, 0, 7, 8, 5, 111, 0, 0, 8, 9, 5, 103, 0, 0,
		9, 10, 5, 114, 0, 0, 10, 11, 5, 97, 0, 0, 11, 12, 5, 109, 0, 0, 12, 2,
		1, 0, 0, 0, 13, 14, 5, 59, 0, 0, 14, 4, 1, 0, 0, 0, 1, 0, 0,
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
	BistatLexerT__0 = 1
	BistatLexerT__1 = 2
)
