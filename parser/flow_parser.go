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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 17, 67, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 6, 2, 22, 10, 2, 13, 2, 14, 2,
	23, 3, 2, 3, 2, 3, 3, 6, 3, 29, 10, 3, 13, 3, 14, 3, 30, 3, 3, 3, 3, 3,
	4, 3, 4, 5, 4, 37, 10, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 6, 6, 45,
	10, 6, 13, 6, 14, 6, 46, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 55,
	10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3,
	10, 2, 2, 11, 2, 4, 6, 8, 10, 12, 14, 16, 18, 2, 3, 3, 2, 7, 8, 2, 66,
	2, 21, 3, 2, 2, 2, 4, 28, 3, 2, 2, 2, 6, 36, 3, 2, 2, 2, 8, 38, 3, 2, 2,
	2, 10, 44, 3, 2, 2, 2, 12, 54, 3, 2, 2, 2, 14, 56, 3, 2, 2, 2, 16, 60,
	3, 2, 2, 2, 18, 64, 3, 2, 2, 2, 20, 22, 5, 4, 3, 2, 21, 20, 3, 2, 2, 2,
	22, 23, 3, 2, 2, 2, 23, 21, 3, 2, 2, 2, 23, 24, 3, 2, 2, 2, 24, 25, 3,
	2, 2, 2, 25, 26, 7, 2, 2, 3, 26, 3, 3, 2, 2, 2, 27, 29, 5, 6, 4, 2, 28,
	27, 3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30, 28, 3, 2, 2, 2, 30, 31, 3, 2, 2,
	2, 31, 32, 3, 2, 2, 2, 32, 33, 7, 14, 2, 2, 33, 5, 3, 2, 2, 2, 34, 37,
	5, 12, 7, 2, 35, 37, 5, 8, 5, 2, 36, 34, 3, 2, 2, 2, 36, 35, 3, 2, 2, 2,
	37, 7, 3, 2, 2, 2, 38, 39, 5, 12, 7, 2, 39, 40, 5, 10, 6, 2, 40, 9, 3,
	2, 2, 2, 41, 42, 5, 18, 10, 2, 42, 43, 5, 12, 7, 2, 43, 45, 3, 2, 2, 2,
	44, 41, 3, 2, 2, 2, 45, 46, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 46, 47, 3,
	2, 2, 2, 47, 11, 3, 2, 2, 2, 48, 55, 7, 12, 2, 2, 49, 55, 7, 13, 2, 2,
	50, 55, 7, 15, 2, 2, 51, 55, 5, 14, 8, 2, 52, 55, 5, 16, 9, 2, 53, 55,
	7, 9, 2, 2, 54, 48, 3, 2, 2, 2, 54, 49, 3, 2, 2, 2, 54, 50, 3, 2, 2, 2,
	54, 51, 3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 54, 53, 3, 2, 2, 2, 55, 13, 3,
	2, 2, 2, 56, 57, 7, 3, 2, 2, 57, 58, 7, 11, 2, 2, 58, 59, 7, 4, 2, 2, 59,
	15, 3, 2, 2, 2, 60, 61, 7, 5, 2, 2, 61, 62, 7, 11, 2, 2, 62, 63, 7, 6,
	2, 2, 63, 17, 3, 2, 2, 2, 64, 65, 9, 2, 2, 2, 65, 19, 3, 2, 2, 2, 7, 23,
	30, 36, 46, 54,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'<'", "'>'", "'['", "']'", "'->'", "'--'", "", "','", "", "'$prev'",
	"'$next'", "';'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "TEXT", "AND", "STRING", "LINK_PREV", "LINK_NEXT",
	"END", "ID", "COMMENT", "WS",
}

var ruleNames = []string{
	"flows", "flow", "stmt", "edge", "edge_rhs", "node", "question", "block",
	"op",
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
	FlowParserTEXT      = 7
	FlowParserAND       = 8
	FlowParserSTRING    = 9
	FlowParserLINK_PREV = 10
	FlowParserLINK_NEXT = 11
	FlowParserEND       = 12
	FlowParserID        = 13
	FlowParserCOMMENT   = 14
	FlowParserWS        = 15
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
	FlowParserRULE_block    = 7
	FlowParserRULE_op       = 8
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
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FlowParserT__0)|(1<<FlowParserT__2)|(1<<FlowParserTEXT)|(1<<FlowParserLINK_PREV)|(1<<FlowParserLINK_NEXT)|(1<<FlowParserID))) != 0) {
		{
			p.SetState(18)
			p.Flow()
		}

		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(23)
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
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FlowParserT__0)|(1<<FlowParserT__2)|(1<<FlowParserTEXT)|(1<<FlowParserLINK_PREV)|(1<<FlowParserLINK_NEXT)|(1<<FlowParserID))) != 0) {
		{
			p.SetState(25)
			p.Stmt()
		}

		p.SetState(28)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(30)
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

	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(32)
			p.Node()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(33)
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
		p.SetState(36)
		p.Node()
	}
	{
		p.SetState(37)
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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == FlowParserT__4 || _la == FlowParserT__5 {
		{
			p.SetState(39)
			p.Op()
		}
		{
			p.SetState(40)
			p.Node()
		}

		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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

	p.SetState(52)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FlowParserLINK_PREV:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(46)
			p.Match(FlowParserLINK_PREV)
		}

	case FlowParserLINK_NEXT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(47)
			p.Match(FlowParserLINK_NEXT)
		}

	case FlowParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(48)
			p.Match(FlowParserID)
		}

	case FlowParserT__0:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(49)
			p.Question()
		}

	case FlowParserT__2:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(50)
			p.Block()
		}

	case FlowParserTEXT:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(51)
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
		p.SetState(54)
		p.Match(FlowParserT__0)
	}
	{
		p.SetState(55)
		p.Match(FlowParserSTRING)
	}
	{
		p.SetState(56)
		p.Match(FlowParserT__1)
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
	p.EnterRule(localctx, 14, FlowParserRULE_block)

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
		p.SetState(58)
		p.Match(FlowParserT__2)
	}
	{
		p.SetState(59)
		p.Match(FlowParserSTRING)
	}
	{
		p.SetState(60)
		p.Match(FlowParserT__3)
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
	p.EnterRule(localctx, 16, FlowParserRULE_op)
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
		p.SetState(62)
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
