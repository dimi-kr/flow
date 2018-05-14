// Generated from ./parser/Flow.g4 by ANTLR 4.7.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link FlowParser}.
 */
public interface FlowListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link FlowParser#flows}.
	 * @param ctx the parse tree
	 */
	void enterFlows(FlowParser.FlowsContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlowParser#flows}.
	 * @param ctx the parse tree
	 */
	void exitFlows(FlowParser.FlowsContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlowParser#flow}.
	 * @param ctx the parse tree
	 */
	void enterFlow(FlowParser.FlowContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlowParser#flow}.
	 * @param ctx the parse tree
	 */
	void exitFlow(FlowParser.FlowContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlowParser#stmt}.
	 * @param ctx the parse tree
	 */
	void enterStmt(FlowParser.StmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlowParser#stmt}.
	 * @param ctx the parse tree
	 */
	void exitStmt(FlowParser.StmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlowParser#edge}.
	 * @param ctx the parse tree
	 */
	void enterEdge(FlowParser.EdgeContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlowParser#edge}.
	 * @param ctx the parse tree
	 */
	void exitEdge(FlowParser.EdgeContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlowParser#edge_rhs}.
	 * @param ctx the parse tree
	 */
	void enterEdge_rhs(FlowParser.Edge_rhsContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlowParser#edge_rhs}.
	 * @param ctx the parse tree
	 */
	void exitEdge_rhs(FlowParser.Edge_rhsContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlowParser#node}.
	 * @param ctx the parse tree
	 */
	void enterNode(FlowParser.NodeContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlowParser#node}.
	 * @param ctx the parse tree
	 */
	void exitNode(FlowParser.NodeContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlowParser#question}.
	 * @param ctx the parse tree
	 */
	void enterQuestion(FlowParser.QuestionContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlowParser#question}.
	 * @param ctx the parse tree
	 */
	void exitQuestion(FlowParser.QuestionContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlowParser#block}.
	 * @param ctx the parse tree
	 */
	void enterBlock(FlowParser.BlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlowParser#block}.
	 * @param ctx the parse tree
	 */
	void exitBlock(FlowParser.BlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link FlowParser#op}.
	 * @param ctx the parse tree
	 */
	void enterOp(FlowParser.OpContext ctx);
	/**
	 * Exit a parse tree produced by {@link FlowParser#op}.
	 * @param ctx the parse tree
	 */
	void exitOp(FlowParser.OpContext ctx);
}