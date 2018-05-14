// Code generated from ./parser/Flow.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Flow

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FlowListener is a complete listener for a parse tree produced by FlowParser.
type FlowListener interface {
	antlr.ParseTreeListener

	// EnterFlows is called when entering the flows production.
	EnterFlows(c *FlowsContext)

	// EnterFlow is called when entering the flow production.
	EnterFlow(c *FlowContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterEdge is called when entering the edge production.
	EnterEdge(c *EdgeContext)

	// EnterEdge_rhs is called when entering the edge_rhs production.
	EnterEdge_rhs(c *Edge_rhsContext)

	// EnterNode is called when entering the node production.
	EnterNode(c *NodeContext)

	// EnterQuestion is called when entering the question production.
	EnterQuestion(c *QuestionContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterOp is called when entering the op production.
	EnterOp(c *OpContext)

	// ExitFlows is called when exiting the flows production.
	ExitFlows(c *FlowsContext)

	// ExitFlow is called when exiting the flow production.
	ExitFlow(c *FlowContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitEdge is called when exiting the edge production.
	ExitEdge(c *EdgeContext)

	// ExitEdge_rhs is called when exiting the edge_rhs production.
	ExitEdge_rhs(c *Edge_rhsContext)

	// ExitNode is called when exiting the node production.
	ExitNode(c *NodeContext)

	// ExitQuestion is called when exiting the question production.
	ExitQuestion(c *QuestionContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitOp is called when exiting the op production.
	ExitOp(c *OpContext)
}
