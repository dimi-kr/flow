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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 17, 105,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 3, 2, 3,
	3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 58,
	10, 10, 12, 10, 14, 10, 61, 11, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3,
	13, 3, 14, 3, 14, 7, 14, 81, 10, 14, 12, 14, 14, 14, 84, 11, 14, 3, 15,
	3, 15, 3, 15, 3, 15, 6, 15, 90, 10, 15, 13, 15, 14, 15, 91, 3, 15, 5, 15,
	95, 10, 15, 3, 15, 3, 15, 3, 16, 6, 16, 100, 10, 16, 13, 16, 14, 16, 101,
	3, 16, 3, 16, 4, 59, 91, 2, 17, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15,
	9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 3, 2,
	6, 4, 2, 67, 92, 99, 124, 5, 2, 50, 59, 67, 92, 99, 124, 3, 3, 12, 12,
	5, 2, 11, 12, 15, 15, 34, 34, 2, 109, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2,
	2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2,
	2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3,
	2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29,
	3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 3, 33, 3, 2, 2, 2, 5, 35, 3, 2, 2, 2, 7,
	37, 3, 2, 2, 2, 9, 39, 3, 2, 2, 2, 11, 41, 3, 2, 2, 2, 13, 44, 3, 2, 2,
	2, 15, 47, 3, 2, 2, 2, 17, 51, 3, 2, 2, 2, 19, 53, 3, 2, 2, 2, 21, 64,
	3, 2, 2, 2, 23, 70, 3, 2, 2, 2, 25, 76, 3, 2, 2, 2, 27, 78, 3, 2, 2, 2,
	29, 85, 3, 2, 2, 2, 31, 99, 3, 2, 2, 2, 33, 34, 7, 62, 2, 2, 34, 4, 3,
	2, 2, 2, 35, 36, 7, 64, 2, 2, 36, 6, 3, 2, 2, 2, 37, 38, 7, 93, 2, 2, 38,
	8, 3, 2, 2, 2, 39, 40, 7, 95, 2, 2, 40, 10, 3, 2, 2, 2, 41, 42, 7, 47,
	2, 2, 42, 43, 7, 64, 2, 2, 43, 12, 3, 2, 2, 2, 44, 45, 7, 47, 2, 2, 45,
	46, 7, 47, 2, 2, 46, 14, 3, 2, 2, 2, 47, 48, 7, 36, 2, 2, 48, 49, 5, 19,
	10, 2, 49, 50, 7, 36, 2, 2, 50, 16, 3, 2, 2, 2, 51, 52, 7, 46, 2, 2, 52,
	18, 3, 2, 2, 2, 53, 59, 7, 36, 2, 2, 54, 55, 7, 94, 2, 2, 55, 58, 7, 36,
	2, 2, 56, 58, 11, 2, 2, 2, 57, 54, 3, 2, 2, 2, 57, 56, 3, 2, 2, 2, 58,
	61, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 60, 62, 3, 2, 2,
	2, 61, 59, 3, 2, 2, 2, 62, 63, 7, 36, 2, 2, 63, 20, 3, 2, 2, 2, 64, 65,
	7, 38, 2, 2, 65, 66, 7, 114, 2, 2, 66, 67, 7, 116, 2, 2, 67, 68, 7, 103,
	2, 2, 68, 69, 7, 120, 2, 2, 69, 22, 3, 2, 2, 2, 70, 71, 7, 38, 2, 2, 71,
	72, 7, 112, 2, 2, 72, 73, 7, 103, 2, 2, 73, 74, 7, 122, 2, 2, 74, 75, 7,
	118, 2, 2, 75, 24, 3, 2, 2, 2, 76, 77, 7, 61, 2, 2, 77, 26, 3, 2, 2, 2,
	78, 82, 9, 2, 2, 2, 79, 81, 9, 3, 2, 2, 80, 79, 3, 2, 2, 2, 81, 84, 3,
	2, 2, 2, 82, 80, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 28, 3, 2, 2, 2, 84,
	82, 3, 2, 2, 2, 85, 86, 7, 49, 2, 2, 86, 87, 7, 49, 2, 2, 87, 89, 3, 2,
	2, 2, 88, 90, 11, 2, 2, 2, 89, 88, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2, 91,
	92, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 92, 94, 3, 2, 2, 2, 93, 95, 9, 4, 2,
	2, 94, 93, 3, 2, 2, 2, 95, 96, 3, 2, 2, 2, 96, 97, 8, 15, 2, 2, 97, 30,
	3, 2, 2, 2, 98, 100, 9, 5, 2, 2, 99, 98, 3, 2, 2, 2, 100, 101, 3, 2, 2,
	2, 101, 99, 3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 103, 3, 2, 2, 2, 103,
	104, 8, 16, 2, 2, 104, 32, 3, 2, 2, 2, 9, 2, 57, 59, 82, 91, 94, 101, 3,
	8, 2, 2,
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
	"", "'<'", "'>'", "'['", "']'", "'->'", "'--'", "", "','", "", "'$prev'",
	"'$next'", "';'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "TEXT", "AND", "STRING", "LINK_PREV", "LINK_NEXT",
	"END", "ID", "COMMENT", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "TEXT", "AND", "STRING",
	"LINK_PREV", "LINK_NEXT", "END", "ID", "COMMENT", "WS",
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
	FlowLexerTEXT      = 7
	FlowLexerAND       = 8
	FlowLexerSTRING    = 9
	FlowLexerLINK_PREV = 10
	FlowLexerLINK_NEXT = 11
	FlowLexerEND       = 12
	FlowLexerID        = 13
	FlowLexerCOMMENT   = 14
	FlowLexerWS        = 15
)
