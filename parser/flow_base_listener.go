// Code generated from ./parser/Flow.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Flow

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFlowListener is a complete listener for a parse tree produced by FlowParser.
type BaseFlowListener struct{}

var _ FlowListener = &BaseFlowListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFlowListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFlowListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFlowListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFlowListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFlows is called when production flows is entered.
func (s *BaseFlowListener) EnterFlows(ctx *FlowsContext) {}

// ExitFlows is called when production flows is exited.
func (s *BaseFlowListener) ExitFlows(ctx *FlowsContext) {}

// EnterFlow is called when production flow is entered.
func (s *BaseFlowListener) EnterFlow(ctx *FlowContext) {}

// ExitFlow is called when production flow is exited.
func (s *BaseFlowListener) ExitFlow(ctx *FlowContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseFlowListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseFlowListener) ExitStmt(ctx *StmtContext) {}

// EnterEdge is called when production edge is entered.
func (s *BaseFlowListener) EnterEdge(ctx *EdgeContext) {}

// ExitEdge is called when production edge is exited.
func (s *BaseFlowListener) ExitEdge(ctx *EdgeContext) {}

// EnterEdge_rhs is called when production edge_rhs is entered.
func (s *BaseFlowListener) EnterEdge_rhs(ctx *Edge_rhsContext) {}

// ExitEdge_rhs is called when production edge_rhs is exited.
func (s *BaseFlowListener) ExitEdge_rhs(ctx *Edge_rhsContext) {}

// EnterNode is called when production node is entered.
func (s *BaseFlowListener) EnterNode(ctx *NodeContext) {}

// ExitNode is called when production node is exited.
func (s *BaseFlowListener) ExitNode(ctx *NodeContext) {}

// EnterQuestion is called when production question is entered.
func (s *BaseFlowListener) EnterQuestion(ctx *QuestionContext) {}

// ExitQuestion is called when production question is exited.
func (s *BaseFlowListener) ExitQuestion(ctx *QuestionContext) {}

// EnterAs_id is called when production as_id is entered.
func (s *BaseFlowListener) EnterAs_id(ctx *As_idContext) {}

// ExitAs_id is called when production as_id is exited.
func (s *BaseFlowListener) ExitAs_id(ctx *As_idContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseFlowListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseFlowListener) ExitBlock(ctx *BlockContext) {}

// EnterAs is called when production as is entered.
func (s *BaseFlowListener) EnterAs(ctx *AsContext) {}

// ExitAs is called when production as is exited.
func (s *BaseFlowListener) ExitAs(ctx *AsContext) {}

// EnterCaption is called when production caption is entered.
func (s *BaseFlowListener) EnterCaption(ctx *CaptionContext) {}

// ExitCaption is called when production caption is exited.
func (s *BaseFlowListener) ExitCaption(ctx *CaptionContext) {}

// EnterOp is called when production op is entered.
func (s *BaseFlowListener) EnterOp(ctx *OpContext) {}

// ExitOp is called when production op is exited.
func (s *BaseFlowListener) ExitOp(ctx *OpContext) {}
