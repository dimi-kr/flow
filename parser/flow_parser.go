// Code generated from ./parser/Flow.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Flow

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 20, 104,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 3, 2, 6, 2, 30, 10, 2, 13, 2, 14, 2, 31, 3, 2, 3,
	2, 3, 3, 6, 3, 37, 10, 3, 13, 3, 14, 3, 38, 3, 3, 3, 3, 3, 4, 3, 4, 5,
	4, 45, 10, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 6, 6, 53, 10, 6, 13,
	6, 14, 6, 54, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 6, 6, 62, 10, 6, 13, 6, 14,
	6, 63, 5, 6, 66, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 74, 10,
	7, 3, 8, 3, 8, 3, 8, 5, 8, 79, 10, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3,
	10, 3, 10, 3, 10, 5, 10, 89, 10, 10, 3, 10, 3, 10, 5, 10, 93, 10, 10, 3,
	11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 2, 2,
	15, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 2, 4, 3, 2, 7, 8, 3,
	2, 10, 11, 2, 104, 2, 29, 3, 2, 2, 2, 4, 36, 3, 2, 2, 2, 6, 44, 3, 2, 2,
	2, 8, 46, 3, 2, 2, 2, 10, 65, 3, 2, 2, 2, 12, 73, 3, 2, 2, 2, 14, 75, 3,
	2, 2, 2, 16, 82, 3, 2, 2, 2, 18, 85, 3, 2, 2, 2, 20, 94, 3, 2, 2, 2, 22,
	96, 3, 2, 2, 2, 24, 98, 3, 2, 2, 2, 26, 101, 3, 2, 2, 2, 28, 30, 5, 4,
	3, 2, 29, 28, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 29, 3, 2, 2, 2, 31, 32,
	3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 34, 7, 2, 2, 3, 34, 3, 3, 2, 2, 2,
	35, 37, 5, 6, 4, 2, 36, 35, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 36, 3,
	2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 41, 7, 17, 2, 2, 41,
	5, 3, 2, 2, 2, 42, 45, 5, 12, 7, 2, 43, 45, 5, 8, 5, 2, 44, 42, 3, 2, 2,
	2, 44, 43, 3, 2, 2, 2, 45, 7, 3, 2, 2, 2, 46, 47, 5, 12, 7, 2, 47, 48,
	5, 10, 6, 2, 48, 9, 3, 2, 2, 2, 49, 50, 5, 26, 14, 2, 50, 51, 5, 12, 7,
	2, 51, 53, 3, 2, 2, 2, 52, 49, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 52,
	3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 66, 3, 2, 2, 2, 56, 57, 5, 26, 14,
	2, 57, 58, 5, 12, 7, 2, 58, 59, 7, 13, 2, 2, 59, 60, 5, 12, 7, 2, 60, 62,
	3, 2, 2, 2, 61, 56, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 61, 3, 2, 2, 2,
	63, 64, 3, 2, 2, 2, 64, 66, 3, 2, 2, 2, 65, 52, 3, 2, 2, 2, 65, 61, 3,
	2, 2, 2, 66, 11, 3, 2, 2, 2, 67, 74, 7, 15, 2, 2, 68, 74, 7, 16, 2, 2,
	69, 74, 7, 18, 2, 2, 70, 74, 5, 14, 8, 2, 71, 74, 5, 18, 10, 2, 72, 74,
	7, 12, 2, 2, 73, 67, 3, 2, 2, 2, 73, 68, 3, 2, 2, 2, 73, 69, 3, 2, 2, 2,
	73, 70, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 73, 72, 3, 2, 2, 2, 74, 13, 3,
	2, 2, 2, 75, 76, 7, 3, 2, 2, 76, 78, 7, 14, 2, 2, 77, 79, 5, 16, 9, 2,
	78, 77, 3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 81, 7,
	4, 2, 2, 81, 15, 3, 2, 2, 2, 82, 83, 5, 20, 11, 2, 83, 84, 7, 18, 2, 2,
	84, 17, 3, 2, 2, 2, 85, 86, 7, 5, 2, 2, 86, 88, 7, 14, 2, 2, 87, 89, 5,
	16, 9, 2, 88, 87, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90,
	92, 7, 6, 2, 2, 91, 93, 5, 24, 13, 2, 92, 91, 3, 2, 2, 2, 92, 93, 3, 2,
	2, 2, 93, 19, 3, 2, 2, 2, 94, 95, 9, 2, 2, 2, 95, 21, 3, 2, 2, 2, 96, 97,
	7, 9, 2, 2, 97, 23, 3, 2, 2, 2, 98, 99, 5, 22, 12, 2, 99, 100, 7, 12, 2,
	2, 100, 25, 3, 2, 2, 2, 101, 102, 9, 3, 2, 2, 102, 27, 3, 2, 2, 2, 12,
	31, 38, 44, 54, 63, 65, 73, 78, 88, 92,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'<'", "'>'", "'['", "']'", "' as '", "':'", "' with '", "'->'", "'--'",
	"", "','", "", "'$prev'", "'$next'", "';'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "TEXT", "AND", "STRING", "LINK_PREV",
	"LINK_NEXT", "END", "ID", "COMMENT", "WS",
}

var ruleNames = []string{
	"flows", "flow", "stmt", "edge", "edge_rhs", "node", "question", "as_id",
	"block", "as", "with", "caption", "op",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type FlowParser struct {
	*antlr.BaseParser
}

func NewFlowParser(input antlr.TokenStream) *FlowParser {
	this := new(FlowParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Flow.g4"

	return this
}

// FlowParser tokens.
const (
	FlowParserEOF       = antlr.TokenEOF
	FlowParserT__0      = 1
	FlowParserT__1      = 2
	FlowParserT__2      = 3
	FlowParserT__3      = 4
	FlowParserT__4      = 5
	FlowParserT__5      = 6
	FlowParserT__6      = 7
	FlowParserT__7      = 8
	FlowParserT__8      = 9
	FlowParserTEXT      = 10
	FlowParserAND       = 11
	FlowParserSTRING    = 12
	FlowParserLINK_PREV = 13
	FlowParserLINK_NEXT = 14
	FlowParserEND       = 15
	FlowParserID        = 16
	FlowParserCOMMENT   = 17
	FlowParserWS        = 18
)

// FlowParser rules.
const (
	FlowParserRULE_flows    = 0
	FlowParserRULE_flow     = 1
	FlowParserRULE_stmt     = 2
	FlowParserRULE_edge     = 3
	FlowParserRULE_edge_rhs = 4
	FlowParserRULE_node     = 5
	FlowParserRULE_question = 6
	FlowParserRULE_as_id    = 7
	FlowParserRULE_block    = 8
	FlowParserRULE_as       = 9
	FlowParserRULE_with     = 10
	FlowParserRULE_caption  = 11
	FlowParserRULE_op       = 12
)

// IFlowsContext is an interface to support dynamic dispatch.
type IFlowsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFlowsContext differentiates from other interfaces.
	IsFlowsContext()
}

type FlowsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFlowsContext() *FlowsContext {
	var p = new(FlowsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlowParserRULE_flows
	return p
}

func (*FlowsContext) IsFlowsContext() {}

func NewFlowsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FlowsContext {
	var p = new(FlowsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowParserRULE_flows

	return p
}

func (s *FlowsContext) GetParser() antlr.Parser { return s.parser }

func (s *FlowsContext) EOF() antlr.TerminalNode {
	return s.GetToken(FlowParserEOF, 0)
}

func (s *FlowsContext) AllFlow() []IFlowContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFlowContext)(nil)).Elem())
	var tst = make([]IFlowContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFlowContext)
		}
	}

	return tst
}

func (s *FlowsContext) Flow(i int) IFlowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFlowContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFlowContext)
}

func (s *FlowsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FlowsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FlowsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.EnterFlows(s)
	}
}

func (s *FlowsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.ExitFlows(s)
	}
}

func (p *FlowParser) Flows() (localctx IFlowsContext) {
	localctx = NewFlowsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FlowParserRULE_flows)
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
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FlowParserT__0)|(1<<FlowParserT__2)|(1<<FlowParserTEXT)|(1<<FlowParserLINK_PREV)|(1<<FlowParserLINK_NEXT)|(1<<FlowParserID))) != 0) {
		{
			p.SetState(26)
			p.Flow()
		}

		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(31)
		p.Match(FlowParserEOF)
	}

	return localctx
}

// IFlowContext is an interface to support dynamic dispatch.
type IFlowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFlowContext differentiates from other interfaces.
	IsFlowContext()
}

type FlowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFlowContext() *FlowContext {
	var p = new(FlowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlowParserRULE_flow
	return p
}

func (*FlowContext) IsFlowContext() {}

func NewFlowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FlowContext {
	var p = new(FlowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowParserRULE_flow

	return p
}

func (s *FlowContext) GetParser() antlr.Parser { return s.parser }

func (s *FlowContext) END() antlr.TerminalNode {
	return s.GetToken(FlowParserEND, 0)
}

func (s *FlowContext) AllStmt() []IStmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStmtContext)(nil)).Elem())
	var tst = make([]IStmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStmtContext)
		}
	}

	return tst
}

func (s *FlowContext) Stmt(i int) IStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *FlowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FlowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FlowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.EnterFlow(s)
	}
}

func (s *FlowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.ExitFlow(s)
	}
}

func (p *FlowParser) Flow() (localctx IFlowContext) {
	localctx = NewFlowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FlowParserRULE_flow)
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
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FlowParserT__0)|(1<<FlowParserT__2)|(1<<FlowParserTEXT)|(1<<FlowParserLINK_PREV)|(1<<FlowParserLINK_NEXT)|(1<<FlowParserID))) != 0) {
		{
			p.SetState(33)
			p.Stmt()
		}

		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(38)
		p.Match(FlowParserEND)
	}

	return localctx
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	p.RuleIndex = FlowParserRULE_stmt
	return p
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Node() INodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INodeContext)
}

func (s *StmtContext) Edge() IEdgeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEdgeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEdgeContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (p *FlowParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FlowParserRULE_stmt)

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

	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(40)
			p.Node()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(41)
			p.Edge()
		}

	}

	return localctx
}

// IEdgeContext is an interface to support dynamic dispatch.
type IEdgeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEdgeContext differentiates from other interfaces.
	IsEdgeContext()
}

type EdgeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEdgeContext() *EdgeContext {
	var p = new(EdgeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlowParserRULE_edge
	return p
}

func (*EdgeContext) IsEdgeContext() {}

func NewEdgeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EdgeContext {
	var p = new(EdgeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowParserRULE_edge

	return p
}

func (s *EdgeContext) GetParser() antlr.Parser { return s.parser }

func (s *EdgeContext) Node() INodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INodeContext)
}

func (s *EdgeContext) Edge_rhs() IEdge_rhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEdge_rhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEdge_rhsContext)
}

func (s *EdgeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EdgeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EdgeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.EnterEdge(s)
	}
}

func (s *EdgeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.ExitEdge(s)
	}
}

func (p *FlowParser) Edge() (localctx IEdgeContext) {
	localctx = NewEdgeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FlowParserRULE_edge)

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
		p.SetState(44)
		p.Node()
	}
	{
		p.SetState(45)
		p.Edge_rhs()
	}

	return localctx
}

// IEdge_rhsContext is an interface to support dynamic dispatch.
type IEdge_rhsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEdge_rhsContext differentiates from other interfaces.
	IsEdge_rhsContext()
}

type Edge_rhsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEdge_rhsContext() *Edge_rhsContext {
	var p = new(Edge_rhsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlowParserRULE_edge_rhs
	return p
}

func (*Edge_rhsContext) IsEdge_rhsContext() {}

func NewEdge_rhsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Edge_rhsContext {
	var p = new(Edge_rhsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowParserRULE_edge_rhs

	return p
}

func (s *Edge_rhsContext) GetParser() antlr.Parser { return s.parser }

func (s *Edge_rhsContext) AllOp() []IOpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOpContext)(nil)).Elem())
	var tst = make([]IOpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOpContext)
		}
	}

	return tst
}

func (s *Edge_rhsContext) Op(i int) IOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOpContext)
}

func (s *Edge_rhsContext) AllNode() []INodeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INodeContext)(nil)).Elem())
	var tst = make([]INodeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INodeContext)
		}
	}

	return tst
}

func (s *Edge_rhsContext) Node(i int) INodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INodeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INodeContext)
}

func (s *Edge_rhsContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(FlowParserAND)
}

func (s *Edge_rhsContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(FlowParserAND, i)
}

func (s *Edge_rhsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Edge_rhsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Edge_rhsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.EnterEdge_rhs(s)
	}
}

func (s *Edge_rhsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.ExitEdge_rhs(s)
	}
}

func (p *FlowParser) Edge_rhs() (localctx IEdge_rhsContext) {
	localctx = NewEdge_rhsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FlowParserRULE_edge_rhs)
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

	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == FlowParserT__7 || _la == FlowParserT__8 {
			{
				p.SetState(47)
				p.Op()
			}
			{
				p.SetState(48)
				p.Node()
			}

			p.SetState(52)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == FlowParserT__7 || _la == FlowParserT__8 {
			{
				p.SetState(54)
				p.Op()
			}
			{
				p.SetState(55)
				p.Node()
			}
			{
				p.SetState(56)
				p.Match(FlowParserAND)
			}
			{
				p.SetState(57)
				p.Node()
			}

			p.SetState(61)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// INodeContext is an interface to support dynamic dispatch.
type INodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNodeContext differentiates from other interfaces.
	IsNodeContext()
}

type NodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNodeContext() *NodeContext {
	var p = new(NodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlowParserRULE_node
	return p
}

func (*NodeContext) IsNodeContext() {}

func NewNodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NodeContext {
	var p = new(NodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowParserRULE_node

	return p
}

func (s *NodeContext) GetParser() antlr.Parser { return s.parser }

func (s *NodeContext) LINK_PREV() antlr.TerminalNode {
	return s.GetToken(FlowParserLINK_PREV, 0)
}

func (s *NodeContext) LINK_NEXT() antlr.TerminalNode {
	return s.GetToken(FlowParserLINK_NEXT, 0)
}

func (s *NodeContext) ID() antlr.TerminalNode {
	return s.GetToken(FlowParserID, 0)
}

func (s *NodeContext) Question() IQuestionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuestionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuestionContext)
}

func (s *NodeContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *NodeContext) TEXT() antlr.TerminalNode {
	return s.GetToken(FlowParserTEXT, 0)
}

func (s *NodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.EnterNode(s)
	}
}

func (s *NodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.ExitNode(s)
	}
}

func (p *FlowParser) Node() (localctx INodeContext) {
	localctx = NewNodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FlowParserRULE_node)

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

	p.SetState(71)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FlowParserLINK_PREV:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(65)
			p.Match(FlowParserLINK_PREV)
		}

	case FlowParserLINK_NEXT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)
			p.Match(FlowParserLINK_NEXT)
		}

	case FlowParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(67)
			p.Match(FlowParserID)
		}

	case FlowParserT__0:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(68)
			p.Question()
		}

	case FlowParserT__2:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(69)
			p.Block()
		}

	case FlowParserTEXT:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(70)
			p.Match(FlowParserTEXT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IQuestionContext is an interface to support dynamic dispatch.
type IQuestionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuestionContext differentiates from other interfaces.
	IsQuestionContext()
}

type QuestionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuestionContext() *QuestionContext {
	var p = new(QuestionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlowParserRULE_question
	return p
}

func (*QuestionContext) IsQuestionContext() {}

func NewQuestionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuestionContext {
	var p = new(QuestionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowParserRULE_question

	return p
}

func (s *QuestionContext) GetParser() antlr.Parser { return s.parser }

func (s *QuestionContext) STRING() antlr.TerminalNode {
	return s.GetToken(FlowParserSTRING, 0)
}

func (s *QuestionContext) As_id() IAs_idContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAs_idContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAs_idContext)
}

func (s *QuestionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuestionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuestionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.EnterQuestion(s)
	}
}

func (s *QuestionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.ExitQuestion(s)
	}
}

func (p *FlowParser) Question() (localctx IQuestionContext) {
	localctx = NewQuestionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FlowParserRULE_question)
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
		p.SetState(73)
		p.Match(FlowParserT__0)
	}
	{
		p.SetState(74)
		p.Match(FlowParserSTRING)
	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FlowParserT__4 || _la == FlowParserT__5 {
		{
			p.SetState(75)
			p.As_id()
		}

	}
	{
		p.SetState(78)
		p.Match(FlowParserT__1)
	}

	return localctx
}

// IAs_idContext is an interface to support dynamic dispatch.
type IAs_idContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAs_idContext differentiates from other interfaces.
	IsAs_idContext()
}

type As_idContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAs_idContext() *As_idContext {
	var p = new(As_idContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlowParserRULE_as_id
	return p
}

func (*As_idContext) IsAs_idContext() {}

func NewAs_idContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *As_idContext {
	var p = new(As_idContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowParserRULE_as_id

	return p
}

func (s *As_idContext) GetParser() antlr.Parser { return s.parser }

func (s *As_idContext) As() IAsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAsContext)
}

func (s *As_idContext) ID() antlr.TerminalNode {
	return s.GetToken(FlowParserID, 0)
}

func (s *As_idContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *As_idContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *As_idContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.EnterAs_id(s)
	}
}

func (s *As_idContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.ExitAs_id(s)
	}
}

func (p *FlowParser) As_id() (localctx IAs_idContext) {
	localctx = NewAs_idContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FlowParserRULE_as_id)

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
		p.SetState(80)
		p.As()
	}
	{
		p.SetState(81)
		p.Match(FlowParserID)
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlowParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) STRING() antlr.TerminalNode {
	return s.GetToken(FlowParserSTRING, 0)
}

func (s *BlockContext) As_id() IAs_idContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAs_idContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAs_idContext)
}

func (s *BlockContext) Caption() ICaptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICaptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICaptionContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *FlowParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FlowParserRULE_block)
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
		p.SetState(83)
		p.Match(FlowParserT__2)
	}
	{
		p.SetState(84)
		p.Match(FlowParserSTRING)
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FlowParserT__4 || _la == FlowParserT__5 {
		{
			p.SetState(85)
			p.As_id()
		}

	}
	{
		p.SetState(88)
		p.Match(FlowParserT__3)
	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FlowParserT__6 {
		{
			p.SetState(89)
			p.Caption()
		}

	}

	return localctx
}

// IAsContext is an interface to support dynamic dispatch.
type IAsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAsContext differentiates from other interfaces.
	IsAsContext()
}

type AsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsContext() *AsContext {
	var p = new(AsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlowParserRULE_as
	return p
}

func (*AsContext) IsAsContext() {}

func NewAsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsContext {
	var p = new(AsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowParserRULE_as

	return p
}

func (s *AsContext) GetParser() antlr.Parser { return s.parser }
func (s *AsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.EnterAs(s)
	}
}

func (s *AsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.ExitAs(s)
	}
}

func (p *FlowParser) As() (localctx IAsContext) {
	localctx = NewAsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FlowParserRULE_as)
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
		p.SetState(92)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FlowParserT__4 || _la == FlowParserT__5) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IWithContext is an interface to support dynamic dispatch.
type IWithContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWithContext differentiates from other interfaces.
	IsWithContext()
}

type WithContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWithContext() *WithContext {
	var p = new(WithContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlowParserRULE_with
	return p
}

func (*WithContext) IsWithContext() {}

func NewWithContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WithContext {
	var p = new(WithContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowParserRULE_with

	return p
}

func (s *WithContext) GetParser() antlr.Parser { return s.parser }
func (s *WithContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WithContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WithContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.EnterWith(s)
	}
}

func (s *WithContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.ExitWith(s)
	}
}

func (p *FlowParser) With() (localctx IWithContext) {
	localctx = NewWithContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, FlowParserRULE_with)

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
		p.SetState(94)
		p.Match(FlowParserT__6)
	}

	return localctx
}

// ICaptionContext is an interface to support dynamic dispatch.
type ICaptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCaptionContext differentiates from other interfaces.
	IsCaptionContext()
}

type CaptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaptionContext() *CaptionContext {
	var p = new(CaptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlowParserRULE_caption
	return p
}

func (*CaptionContext) IsCaptionContext() {}

func NewCaptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaptionContext {
	var p = new(CaptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowParserRULE_caption

	return p
}

func (s *CaptionContext) GetParser() antlr.Parser { return s.parser }

func (s *CaptionContext) With() IWithContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWithContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWithContext)
}

func (s *CaptionContext) TEXT() antlr.TerminalNode {
	return s.GetToken(FlowParserTEXT, 0)
}

func (s *CaptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.EnterCaption(s)
	}
}

func (s *CaptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.ExitCaption(s)
	}
}

func (p *FlowParser) Caption() (localctx ICaptionContext) {
	localctx = NewCaptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, FlowParserRULE_caption)

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
		p.SetState(96)
		p.With()
	}
	{
		p.SetState(97)
		p.Match(FlowParserTEXT)
	}

	return localctx
}

// IOpContext is an interface to support dynamic dispatch.
type IOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpContext differentiates from other interfaces.
	IsOpContext()
}

type OpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpContext() *OpContext {
	var p = new(OpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlowParserRULE_op
	return p
}

func (*OpContext) IsOpContext() {}

func NewOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpContext {
	var p = new(OpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowParserRULE_op

	return p
}

func (s *OpContext) GetParser() antlr.Parser { return s.parser }
func (s *OpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.EnterOp(s)
	}
}

func (s *OpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.ExitOp(s)
	}
}

func (p *FlowParser) Op() (localctx IOpContext) {
	localctx = NewOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, FlowParserRULE_op)
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
		p.SetState(99)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FlowParserT__7 || _la == FlowParserT__8) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
