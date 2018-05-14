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
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		TEXT=10, AND=11, STRING=12, LINK_PREV=13, LINK_NEXT=14, END=15, ID=16, 
		COMMENT=17, WS=18;
	public static final int
		RULE_flows = 0, RULE_flow = 1, RULE_stmt = 2, RULE_edge = 3, RULE_edge_rhs = 4, 
		RULE_node = 5, RULE_question = 6, RULE_as_id = 7, RULE_block = 8, RULE_as = 9, 
		RULE_caption = 10, RULE_op = 11;
	public static final String[] ruleNames = {
		"flows", "flow", "stmt", "edge", "edge_rhs", "node", "question", "as_id", 
		"block", "as", "caption", "op"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "'<'", "'>'", "'['", "']'", "' as '", "':'", "' with '", "'->'", 
		"'--'", null, "','", null, "'$prev'", "'$next'", "';'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, null, null, null, null, null, null, null, null, null, "TEXT", "AND", 
		"STRING", "LINK_PREV", "LINK_NEXT", "END", "ID", "COMMENT", "WS"
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
			setState(25); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(24);
				flow();
				}
				}
				setState(27); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__2) | (1L << TEXT) | (1L << LINK_PREV) | (1L << LINK_NEXT) | (1L << ID))) != 0) );
			setState(29);
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
			setState(32); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(31);
				stmt();
				}
				}
				setState(34); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__2) | (1L << TEXT) | (1L << LINK_PREV) | (1L << LINK_NEXT) | (1L << ID))) != 0) );
			setState(36);
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
			setState(40);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(38);
				node();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(39);
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
			setState(42);
			node();
			setState(43);
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
		public List<TerminalNode> AND() { return getTokens(FlowParser.AND); }
		public TerminalNode AND(int i) {
			return getToken(FlowParser.AND, i);
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
			setState(61);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(48); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(45);
					op();
					setState(46);
					node();
					}
					}
					setState(50); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==T__7 || _la==T__8 );
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(57); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(52);
					op();
					setState(53);
					node();
					setState(54);
					match(AND);
					setState(55);
					node();
					}
					}
					setState(59); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==T__7 || _la==T__8 );
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
			setState(69);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LINK_PREV:
				enterOuterAlt(_localctx, 1);
				{
				setState(63);
				match(LINK_PREV);
				}
				break;
			case LINK_NEXT:
				enterOuterAlt(_localctx, 2);
				{
				setState(64);
				match(LINK_NEXT);
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 3);
				{
				setState(65);
				match(ID);
				}
				break;
			case T__0:
				enterOuterAlt(_localctx, 4);
				{
				setState(66);
				question();
				}
				break;
			case T__2:
				enterOuterAlt(_localctx, 5);
				{
				setState(67);
				block();
				}
				break;
			case TEXT:
				enterOuterAlt(_localctx, 6);
				{
				setState(68);
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
		public As_idContext as_id() {
			return getRuleContext(As_idContext.class,0);
		}
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
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(71);
			match(T__0);
			setState(72);
			match(STRING);
			setState(74);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__4 || _la==T__5) {
				{
				setState(73);
				as_id();
				}
			}

			setState(76);
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

	public static class As_idContext extends ParserRuleContext {
		public AsContext as() {
			return getRuleContext(AsContext.class,0);
		}
		public TerminalNode ID() { return getToken(FlowParser.ID, 0); }
		public As_idContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_as_id; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof FlowListener ) ((FlowListener)listener).enterAs_id(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof FlowListener ) ((FlowListener)listener).exitAs_id(this);
		}
	}

	public final As_idContext as_id() throws RecognitionException {
		As_idContext _localctx = new As_idContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_as_id);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(78);
			as();
			setState(79);
			match(ID);
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
		public As_idContext as_id() {
			return getRuleContext(As_idContext.class,0);
		}
		public CaptionContext caption() {
			return getRuleContext(CaptionContext.class,0);
		}
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
		enterRule(_localctx, 16, RULE_block);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(81);
			match(T__2);
			setState(82);
			match(STRING);
			setState(84);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__4 || _la==T__5) {
				{
				setState(83);
				as_id();
				}
			}

			setState(86);
			match(T__3);
			setState(88);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__6) {
				{
				setState(87);
				caption();
				}
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

	public static class AsContext extends ParserRuleContext {
		public AsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_as; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof FlowListener ) ((FlowListener)listener).enterAs(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof FlowListener ) ((FlowListener)listener).exitAs(this);
		}
	}

	public final AsContext as() throws RecognitionException {
		AsContext _localctx = new AsContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_as);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(90);
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

	public static class CaptionContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(FlowParser.STRING, 0); }
		public CaptionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_caption; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof FlowListener ) ((FlowListener)listener).enterCaption(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof FlowListener ) ((FlowListener)listener).exitCaption(this);
		}
	}

	public final CaptionContext caption() throws RecognitionException {
		CaptionContext _localctx = new CaptionContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_caption);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(92);
			match(T__6);
			setState(93);
			match(STRING);
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
		enterRule(_localctx, 22, RULE_op);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(95);
			_la = _input.LA(1);
			if ( !(_la==T__7 || _la==T__8) ) {
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\24d\4\2\t\2\4\3\t"+
		"\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t\13\4"+
		"\f\t\f\4\r\t\r\3\2\6\2\34\n\2\r\2\16\2\35\3\2\3\2\3\3\6\3#\n\3\r\3\16"+
		"\3$\3\3\3\3\3\4\3\4\5\4+\n\4\3\5\3\5\3\5\3\6\3\6\3\6\6\6\63\n\6\r\6\16"+
		"\6\64\3\6\3\6\3\6\3\6\3\6\6\6<\n\6\r\6\16\6=\5\6@\n\6\3\7\3\7\3\7\3\7"+
		"\3\7\3\7\5\7H\n\7\3\b\3\b\3\b\5\bM\n\b\3\b\3\b\3\t\3\t\3\t\3\n\3\n\3\n"+
		"\5\nW\n\n\3\n\3\n\5\n[\n\n\3\13\3\13\3\f\3\f\3\f\3\r\3\r\3\r\2\2\16\2"+
		"\4\6\b\n\f\16\20\22\24\26\30\2\4\3\2\7\b\3\2\n\13\2e\2\33\3\2\2\2\4\""+
		"\3\2\2\2\6*\3\2\2\2\b,\3\2\2\2\n?\3\2\2\2\fG\3\2\2\2\16I\3\2\2\2\20P\3"+
		"\2\2\2\22S\3\2\2\2\24\\\3\2\2\2\26^\3\2\2\2\30a\3\2\2\2\32\34\5\4\3\2"+
		"\33\32\3\2\2\2\34\35\3\2\2\2\35\33\3\2\2\2\35\36\3\2\2\2\36\37\3\2\2\2"+
		"\37 \7\2\2\3 \3\3\2\2\2!#\5\6\4\2\"!\3\2\2\2#$\3\2\2\2$\"\3\2\2\2$%\3"+
		"\2\2\2%&\3\2\2\2&\'\7\21\2\2\'\5\3\2\2\2(+\5\f\7\2)+\5\b\5\2*(\3\2\2\2"+
		"*)\3\2\2\2+\7\3\2\2\2,-\5\f\7\2-.\5\n\6\2.\t\3\2\2\2/\60\5\30\r\2\60\61"+
		"\5\f\7\2\61\63\3\2\2\2\62/\3\2\2\2\63\64\3\2\2\2\64\62\3\2\2\2\64\65\3"+
		"\2\2\2\65@\3\2\2\2\66\67\5\30\r\2\678\5\f\7\289\7\r\2\29:\5\f\7\2:<\3"+
		"\2\2\2;\66\3\2\2\2<=\3\2\2\2=;\3\2\2\2=>\3\2\2\2>@\3\2\2\2?\62\3\2\2\2"+
		"?;\3\2\2\2@\13\3\2\2\2AH\7\17\2\2BH\7\20\2\2CH\7\22\2\2DH\5\16\b\2EH\5"+
		"\22\n\2FH\7\f\2\2GA\3\2\2\2GB\3\2\2\2GC\3\2\2\2GD\3\2\2\2GE\3\2\2\2GF"+
		"\3\2\2\2H\r\3\2\2\2IJ\7\3\2\2JL\7\16\2\2KM\5\20\t\2LK\3\2\2\2LM\3\2\2"+
		"\2MN\3\2\2\2NO\7\4\2\2O\17\3\2\2\2PQ\5\24\13\2QR\7\22\2\2R\21\3\2\2\2"+
		"ST\7\5\2\2TV\7\16\2\2UW\5\20\t\2VU\3\2\2\2VW\3\2\2\2WX\3\2\2\2XZ\7\6\2"+
		"\2Y[\5\26\f\2ZY\3\2\2\2Z[\3\2\2\2[\23\3\2\2\2\\]\t\2\2\2]\25\3\2\2\2^"+
		"_\7\t\2\2_`\7\16\2\2`\27\3\2\2\2ab\t\3\2\2b\31\3\2\2\2\f\35$*\64=?GLV"+
		"Z";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}