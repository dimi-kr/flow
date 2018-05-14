// Generated from ./parser/Flow.g4 by ANTLR 4.7.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class FlowParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.7.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, TEXT=7, AND=8, STRING=9, 
		LINK_PREV=10, LINK_NEXT=11, END=12, ID=13, COMMENT=14, WS=15;
	public static final int
		RULE_flows = 0, RULE_flow = 1, RULE_stmt = 2, RULE_edge = 3, RULE_edge_rhs = 4, 
		RULE_node = 5, RULE_question = 6, RULE_block = 7, RULE_op = 8;
	public static final String[] ruleNames = {
		"flows", "flow", "stmt", "edge", "edge_rhs", "node", "question", "block", 
		"op"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "'<'", "'>'", "'['", "']'", "'->'", "'--'", null, "','", null, "'$prev'", 
		"'$next'", "';'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, null, null, null, null, null, null, "TEXT", "AND", "STRING", "LINK_PREV", 
		"LINK_NEXT", "END", "ID", "COMMENT", "WS"
	};
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Flow.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public FlowParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}
	public static class FlowsContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(FlowParser.EOF, 0); }
		public List<FlowContext> flow() {
			return getRuleContexts(FlowContext.class);
		}
		public FlowContext flow(int i) {
			return getRuleContext(FlowContext.class,i);
		}
		public FlowsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_flows; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof FlowListener ) ((FlowListener)listener).enterFlows(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof FlowListener ) ((FlowListener)listener).exitFlows(this);
		}
	}

	public final FlowsContext flows() throws RecognitionException {
		FlowsContext _localctx = new FlowsContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_flows);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(19); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(18);
				flow();
				}
				}
				setState(21); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__2) | (1L << TEXT) | (1L << LINK_PREV) | (1L << LINK_NEXT) | (1L << ID))) != 0) );
			setState(23);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FlowContext extends ParserRuleContext {
		public TerminalNode END() { return getToken(FlowParser.END, 0); }
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public FlowContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_flow; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof FlowListener ) ((FlowListener)listener).enterFlow(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof FlowListener ) ((FlowListener)listener).exitFlow(this);
		}
	}

	public final FlowContext flow() throws RecognitionException {
		FlowContext _localctx = new FlowContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_flow);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(26); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(25);
				stmt();
				}
				}
				setState(28); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__2) | (1L << TEXT) | (1L << LINK_PREV) | (1L << LINK_NEXT) | (1L << ID))) != 0) );
			setState(30);
			match(END);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StmtContext extends ParserRuleContext {
		public NodeContext node() {
			return getRuleContext(NodeContext.class,0);
		}
		public EdgeContext edge() {
			return getRuleContext(EdgeContext.class,0);
		}
		public StmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stmt; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof FlowListener ) ((FlowListener)listener).enterStmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof FlowListener ) ((FlowListener)listener).exitStmt(this);
		}
	}

	public final StmtContext stmt() throws RecognitionException {
		StmtContext _localctx = new StmtContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_stmt);
		try {
			setState(34);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(32);
				node();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(33);
				edge();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class EdgeContext extends ParserRuleContext {
		public NodeContext node() {
			return getRuleContext(NodeContext.class,0);
		}
		public Edge_rhsContext edge_rhs() {
			return getRuleContext(Edge_rhsContext.class,0);
		}
		public EdgeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_edge; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof FlowListener ) ((FlowListener)listener).enterEdge(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof FlowListener ) ((FlowListener)listener).exitEdge(this);
		}
	}

	public final EdgeContext edge() throws RecognitionException {
		EdgeContext _localctx = new EdgeContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_edge);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(36);
			node();
			setState(37);
			edge_rhs();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Edge_rhsContext extends ParserRuleContext {
		public List<OpContext> op() {
			return getRuleContexts(OpContext.class);
		}
		public OpContext op(int i) {
			return getRuleContext(OpContext.class,i);
		}
		public List<NodeContext> node() {
			return getRuleContexts(NodeContext.class);
		}
		public NodeContext node(int i) {
			return getRuleContext(NodeContext.class,i);
		}
		public Edge_rhsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_edge_rhs; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof FlowListener ) ((FlowListener)listener).enterEdge_rhs(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof FlowListener ) ((FlowListener)listener).exitEdge_rhs(this);
		}
	}

	public final Edge_rhsContext edge_rhs() throws RecognitionException {
		Edge_rhsContext _localctx = new Edge_rhsContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_edge_rhs);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(42); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(39);
				op();
				setState(40);
				node();
				}
				}
				setState(44); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==T__4 || _la==T__5 );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class NodeContext extends ParserRuleContext {
		public TerminalNode LINK_PREV() { return getToken(FlowParser.LINK_PREV, 0); }
		public TerminalNode LINK_NEXT() { return getToken(FlowParser.LINK_NEXT, 0); }
		public TerminalNode ID() { return getToken(FlowParser.ID, 0); }
		public QuestionContext question() {
			return getRuleContext(QuestionContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode TEXT() { return getToken(FlowParser.TEXT, 0); }
		public NodeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_node; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof FlowListener ) ((FlowListener)listener).enterNode(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof FlowListener ) ((FlowListener)listener).exitNode(this);
		}
	}

	public final NodeContext node() throws RecognitionException {
		NodeContext _localctx = new NodeContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_node);
		try {
			setState(52);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LINK_PREV:
				enterOuterAlt(_localctx, 1);
				{
				setState(46);
				match(LINK_PREV);
				}
				break;
			case LINK_NEXT:
				enterOuterAlt(_localctx, 2);
				{
				setState(47);
				match(LINK_NEXT);
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 3);
				{
				setState(48);
				match(ID);
				}
				break;
			case T__0:
				enterOuterAlt(_localctx, 4);
				{
				setState(49);
				question();
				}
				break;
			case T__2:
				enterOuterAlt(_localctx, 5);
				{
				setState(50);
				block();
				}
				break;
			case TEXT:
				enterOuterAlt(_localctx, 6);
				{
				setState(51);
				match(TEXT);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class QuestionContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(FlowParser.STRING, 0); }
		public QuestionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_question; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof FlowListener ) ((FlowListener)listener).enterQuestion(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof FlowListener ) ((FlowListener)listener).exitQuestion(this);
		}
	}

	public final QuestionContext question() throws RecognitionException {
		QuestionContext _localctx = new QuestionContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_question);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(54);
			match(T__0);
			setState(55);
			match(STRING);
			setState(56);
			match(T__1);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class BlockContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(FlowParser.STRING, 0); }
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof FlowListener ) ((FlowListener)listener).enterBlock(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof FlowListener ) ((FlowListener)listener).exitBlock(this);
		}
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_block);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(58);
			match(T__2);
			setState(59);
			match(STRING);
			setState(60);
			match(T__3);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class OpContext extends ParserRuleContext {
		public OpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_op; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof FlowListener ) ((FlowListener)listener).enterOp(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof FlowListener ) ((FlowListener)listener).exitOp(this);
		}
	}

	public final OpContext op() throws RecognitionException {
		OpContext _localctx = new OpContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_op);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(62);
			_la = _input.LA(1);
			if ( !(_la==T__4 || _la==T__5) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\21C\4\2\t\2\4\3\t"+
		"\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\3\2\6\2\26"+
		"\n\2\r\2\16\2\27\3\2\3\2\3\3\6\3\35\n\3\r\3\16\3\36\3\3\3\3\3\4\3\4\5"+
		"\4%\n\4\3\5\3\5\3\5\3\6\3\6\3\6\6\6-\n\6\r\6\16\6.\3\7\3\7\3\7\3\7\3\7"+
		"\3\7\5\7\67\n\7\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\n\3\n\3\n\2\2\13\2\4"+
		"\6\b\n\f\16\20\22\2\3\3\2\7\b\2B\2\25\3\2\2\2\4\34\3\2\2\2\6$\3\2\2\2"+
		"\b&\3\2\2\2\n,\3\2\2\2\f\66\3\2\2\2\168\3\2\2\2\20<\3\2\2\2\22@\3\2\2"+
		"\2\24\26\5\4\3\2\25\24\3\2\2\2\26\27\3\2\2\2\27\25\3\2\2\2\27\30\3\2\2"+
		"\2\30\31\3\2\2\2\31\32\7\2\2\3\32\3\3\2\2\2\33\35\5\6\4\2\34\33\3\2\2"+
		"\2\35\36\3\2\2\2\36\34\3\2\2\2\36\37\3\2\2\2\37 \3\2\2\2 !\7\16\2\2!\5"+
		"\3\2\2\2\"%\5\f\7\2#%\5\b\5\2$\"\3\2\2\2$#\3\2\2\2%\7\3\2\2\2&\'\5\f\7"+
		"\2\'(\5\n\6\2(\t\3\2\2\2)*\5\22\n\2*+\5\f\7\2+-\3\2\2\2,)\3\2\2\2-.\3"+
		"\2\2\2.,\3\2\2\2./\3\2\2\2/\13\3\2\2\2\60\67\7\f\2\2\61\67\7\r\2\2\62"+
		"\67\7\17\2\2\63\67\5\16\b\2\64\67\5\20\t\2\65\67\7\t\2\2\66\60\3\2\2\2"+
		"\66\61\3\2\2\2\66\62\3\2\2\2\66\63\3\2\2\2\66\64\3\2\2\2\66\65\3\2\2\2"+
		"\67\r\3\2\2\289\7\3\2\29:\7\13\2\2:;\7\4\2\2;\17\3\2\2\2<=\7\5\2\2=>\7"+
		"\13\2\2>?\7\6\2\2?\21\3\2\2\2@A\t\2\2\2A\23\3\2\2\2\7\27\36$.\66";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}