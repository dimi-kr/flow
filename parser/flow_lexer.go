// Code generated from ./parser/Flow.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 20, 125,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13, 78, 10, 13, 12,
	13, 14, 13, 81, 11, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3,
	17, 7, 17, 101, 10, 17, 12, 17, 14, 17, 104, 11, 17, 3, 18, 3, 18, 3, 18,
	3, 18, 6, 18, 110, 10, 18, 13, 18, 14, 18, 111, 3, 18, 5, 18, 115, 10,
	18, 3, 18, 3, 18, 3, 19, 6, 19, 120, 10, 19, 13, 19, 14, 19, 121, 3, 19,
	3, 19, 4, 79, 111, 2, 20, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9,
	17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18,
	35, 19, 37, 20, 3, 2, 6, 4, 2, 67, 92, 99, 124, 5, 2, 50, 59, 67, 92, 99,
	124, 3, 3, 12, 12, 5, 2, 11, 12, 15, 15, 34, 34, 2, 129, 2, 3, 3, 2, 2,
	2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2,
	2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2,
	2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3,
	2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35,
	3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 3, 39, 3, 2, 2, 2, 5, 41, 3, 2, 2, 2, 7,
	43, 3, 2, 2, 2, 9, 45, 3, 2, 2, 2, 11, 47, 3, 2, 2, 2, 13, 52, 3, 2, 2,
	2, 15, 54, 3, 2, 2, 2, 17, 61, 3, 2, 2, 2, 19, 64, 3, 2, 2, 2, 21, 67,
	3, 2, 2, 2, 23, 71, 3, 2, 2, 2, 25, 73, 3, 2, 2, 2, 27, 84, 3, 2, 2, 2,
	29, 90, 3, 2, 2, 2, 31, 96, 3, 2, 2, 2, 33, 98, 3, 2, 2, 2, 35, 105, 3,
	2, 2, 2, 37, 119, 3, 2, 2, 2, 39, 40, 7, 62, 2, 2, 40, 4, 3, 2, 2, 2, 41,
	42, 7, 64, 2, 2, 42, 6, 3, 2, 2, 2, 43, 44, 7, 93, 2, 2, 44, 8, 3, 2, 2,
	2, 45, 46, 7, 95, 2, 2, 46, 10, 3, 2, 2, 2, 47, 48, 7, 34, 2, 2, 48, 49,
	7, 99, 2, 2, 49, 50, 7, 117, 2, 2, 50, 51, 7, 34, 2, 2, 51, 12, 3, 2, 2,
	2, 52, 53, 7, 60, 2, 2, 53, 14, 3, 2, 2, 2, 54, 55, 7, 34, 2, 2, 55, 56,
	7, 121, 2, 2, 56, 57, 7, 107, 2, 2, 57, 58, 7, 118, 2, 2, 58, 59, 7, 106,
	2, 2, 59, 60, 7, 34, 2, 2, 60, 16, 3, 2, 2, 2, 61, 62, 7, 47, 2, 2, 62,
	63, 7, 64, 2, 2, 63, 18, 3, 2, 2, 2, 64, 65, 7, 47, 2, 2, 65, 66, 7, 47,
	2, 2, 66, 20, 3, 2, 2, 2, 67, 68, 7, 36, 2, 2, 68, 69, 5, 25, 13, 2, 69,
	70, 7, 36, 2, 2, 70, 22, 3, 2, 2, 2, 71, 72, 7, 46, 2, 2, 72, 24, 3, 2,
	2, 2, 73, 79, 7, 36, 2, 2, 74, 75, 7, 94, 2, 2, 75, 78, 7, 36, 2, 2, 76,
	78, 11, 2, 2, 2, 77, 74, 3, 2, 2, 2, 77, 76, 3, 2, 2, 2, 78, 81, 3, 2,
	2, 2, 79, 80, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 80, 82, 3, 2, 2, 2, 81, 79,
	3, 2, 2, 2, 82, 83, 7, 36, 2, 2, 83, 26, 3, 2, 2, 2, 84, 85, 7, 38, 2,
	2, 85, 86, 7, 114, 2, 2, 86, 87, 7, 116, 2, 2, 87, 88, 7, 103, 2, 2, 88,
	89, 7, 120, 2, 2, 89, 28, 3, 2, 2, 2, 90, 91, 7, 38, 2, 2, 91, 92, 7, 112,
	2, 2, 92, 93, 7, 103, 2, 2, 93, 94, 7, 122, 2, 2, 94, 95, 7, 118, 2, 2,
	95, 30, 3, 2, 2, 2, 96, 97, 7, 61, 2, 2, 97, 32, 3, 2, 2, 2, 98, 102, 9,
	2, 2, 2, 99, 101, 9, 3, 2, 2, 100, 99, 3, 2, 2, 2, 101, 104, 3, 2, 2, 2,
	102, 100, 3, 2, 2, 2, 102, 103, 3, 2, 2, 2, 103, 34, 3, 2, 2, 2, 104, 102,
	3, 2, 2, 2, 105, 106, 7, 49, 2, 2, 106, 107, 7, 49, 2, 2, 107, 109, 3,
	2, 2, 2, 108, 110, 11, 2, 2, 2, 109, 108, 3, 2, 2, 2, 110, 111, 3, 2, 2,
	2, 111, 112, 3, 2, 2, 2, 111, 109, 3, 2, 2, 2, 112, 114, 3, 2, 2, 2, 113,
	115, 9, 4, 2, 2, 114, 113, 3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 117,
	8, 18, 2, 2, 117, 36, 3, 2, 2, 2, 118, 120, 9, 5, 2, 2, 119, 118, 3, 2,
	2, 2, 120, 121, 3, 2, 2, 2, 121, 119, 3, 2, 2, 2, 121, 122, 3, 2, 2, 2,
	122, 123, 3, 2, 2, 2, 123, 124, 8, 19, 2, 2, 124, 38, 3, 2, 2, 2, 9, 2,
	77, 79, 102, 111, 114, 121, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'<'", "'>'", "'['", "']'", "' as '", "':'", "' with '", "'->'", "'--'",
	"", "','", "", "'$prev'", "'$next'", "';'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "TEXT", "AND", "STRING", "LINK_PREV",
	"LINK_NEXT", "END", "ID", "COMMENT", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"TEXT", "AND", "STRING", "LINK_PREV", "LINK_NEXT", "END", "ID", "COMMENT",
	"WS",
}

type FlowLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewFlowLexer(input antlr.CharStream) *FlowLexer {

	l := new(FlowLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Flow.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// FlowLexer tokens.
const (
	FlowLexerT__0      = 1
	FlowLexerT__1      = 2
	FlowLexerT__2      = 3
	FlowLexerT__3      = 4
	FlowLexerT__4      = 5
	FlowLexerT__5      = 6
	FlowLexerT__6      = 7
	FlowLexerT__7      = 8
	FlowLexerT__8      = 9
	FlowLexerTEXT      = 10
	FlowLexerAND       = 11
	FlowLexerSTRING    = 12
	FlowLexerLINK_PREV = 13
	FlowLexerLINK_NEXT = 14
	FlowLexerEND       = 15
	FlowLexerID        = 16
	FlowLexerCOMMENT   = 17
	FlowLexerWS        = 18
)
