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
		RULE_with = 10, RULE_caption = 11, RULE_op = 12;
	public static final String[] ruleNames = {
		"flows", "flow", "stmt", "edge", "edge_rhs", "node", "question", "as_id", 
		"block", "as", "with", "caption", "op"
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
			setState(27); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(26);
				flow();
				}
				}
				setState(29); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__2) | (1L << TEXT) | (1L << LINK_PREV) | (1L << LINK_NEXT) | (1L << ID))) != 0) );
			setState(31);
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
			setState(34); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(33);
				stmt();
				}
				}
				setState(36); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__2) | (1L << TEXT) | (1L << LINK_PREV) | (1L << LINK_NEXT) | (1L << ID))) != 0) );
			setState(38);
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
			setState(42);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(40);
				node();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(41);
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
			setState(44);
			node();
			setState(45);
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
			setState(63);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(50); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(47);
					op();
					setState(48);
					node();
					}
					}
					setState(52); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==T__7 || _la==T__8 );
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(59); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(54);
					op();
					setState(55);
					node();
					setState(56);
					match(AND);
					setState(57);
					node();
					}
					}
					setState(61); 
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
			setState(71);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LINK_PREV:
				enterOuterAlt(_localctx, 1);
				{
				setState(65);
				match(LINK_PREV);
				}
				break;
			case LINK_NEXT:
				enterOuterAlt(_localctx, 2);
				{
				setState(66);
				match(LINK_NEXT);
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 3);
				{
				setState(67);
				match(ID);
				}
				break;
			case T__0:
				enterOuterAlt(_localctx, 4);
				{
				setState(68);
				question();
				}
				break;
			case T__2:
				enterOuterAlt(_localctx, 5);
				{
				setState(69);
				block();
				}
				break;
			case TEXT:
				enterOuterAlt(_localctx, 6);
				{
				setState(70);
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
			setState(73);
			match(T__0);
			setState(74);
			match(STRING);
			setState(76);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__4 || _la==T__5) {
				{
				setState(75);
				as_id();
				}
			}

			setState(78);
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
			setState(80);
			as();
			setState(81);
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
			setState(83);
			match(T__2);
			setState(84);
			match(STRING);
			setState(86);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__4 || _la==T__5) {
				{
				setState(85);
				as_id();
				}
			}

			setState(88);
			match(T__3);
			setState(90);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__6) {
				{
				setState(89);
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
			setState(92);
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

	public static class WithContext extends ParserRuleContext {
		public WithContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_with; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof FlowListener ) ((FlowListener)listener).enterWith(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof FlowListener ) ((FlowListener)listener).exitWith(this);
		}
	}

	public final WithContext with() throws RecognitionException {
		WithContext _localctx = new WithContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_with);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(94);
			match(T__6);
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
		public WithContext with() {
			return getRuleContext(WithContext.class,0);
		}
		public TerminalNode TEXT() { return getToken(FlowParser.TEXT, 0); }
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
		enterRule(_localctx, 22, RULE_caption);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(96);
			with();
			setState(97);
			match(TEXT);
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
		enterRule(_localctx, 24, RULE_op);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(99);
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\24h\4\2\t\2\4\3\t"+
		"\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t\13\4"+
		"\f\t\f\4\r\t\r\4\16\t\16\3\2\6\2\36\n\2\r\2\16\2\37\3\2\3\2\3\3\6\3%\n"+
		"\3\r\3\16\3&\3\3\3\3\3\4\3\4\5\4-\n\4\3\5\3\5\3\5\3\6\3\6\3\6\6\6\65\n"+
		"\6\r\6\16\6\66\3\6\3\6\3\6\3\6\3\6\6\6>\n\6\r\6\16\6?\5\6B\n\6\3\7\3\7"+
		"\3\7\3\7\3\7\3\7\5\7J\n\7\3\b\3\b\3\b\5\bO\n\b\3\b\3\b\3\t\3\t\3\t\3\n"+
		"\3\n\3\n\5\nY\n\n\3\n\3\n\5\n]\n\n\3\13\3\13\3\f\3\f\3\r\3\r\3\r\3\16"+
		"\3\16\3\16\2\2\17\2\4\6\b\n\f\16\20\22\24\26\30\32\2\4\3\2\7\b\3\2\n\13"+
		"\2h\2\35\3\2\2\2\4$\3\2\2\2\6,\3\2\2\2\b.\3\2\2\2\nA\3\2\2\2\fI\3\2\2"+
		"\2\16K\3\2\2\2\20R\3\2\2\2\22U\3\2\2\2\24^\3\2\2\2\26`\3\2\2\2\30b\3\2"+
		"\2\2\32e\3\2\2\2\34\36\5\4\3\2\35\34\3\2\2\2\36\37\3\2\2\2\37\35\3\2\2"+
		"\2\37 \3\2\2\2 !\3\2\2\2!\"\7\2\2\3\"\3\3\2\2\2#%\5\6\4\2$#\3\2\2\2%&"+
		"\3\2\2\2&$\3\2\2\2&\'\3\2\2\2\'(\3\2\2\2()\7\21\2\2)\5\3\2\2\2*-\5\f\7"+
		"\2+-\5\b\5\2,*\3\2\2\2,+\3\2\2\2-\7\3\2\2\2./\5\f\7\2/\60\5\n\6\2\60\t"+
		"\3\2\2\2\61\62\5\32\16\2\62\63\5\f\7\2\63\65\3\2\2\2\64\61\3\2\2\2\65"+
		"\66\3\2\2\2\66\64\3\2\2\2\66\67\3\2\2\2\67B\3\2\2\289\5\32\16\29:\5\f"+
		"\7\2:;\7\r\2\2;<\5\f\7\2<>\3\2\2\2=8\3\2\2\2>?\3\2\2\2?=\3\2\2\2?@\3\2"+
		"\2\2@B\3\2\2\2A\64\3\2\2\2A=\3\2\2\2B\13\3\2\2\2CJ\7\17\2\2DJ\7\20\2\2"+
		"EJ\7\22\2\2FJ\5\16\b\2GJ\5\22\n\2HJ\7\f\2\2IC\3\2\2\2ID\3\2\2\2IE\3\2"+
		"\2\2IF\3\2\2\2IG\3\2\2\2IH\3\2\2\2J\r\3\2\2\2KL\7\3\2\2LN\7\16\2\2MO\5"+
		"\20\t\2NM\3\2\2\2NO\3\2\2\2OP\3\2\2\2PQ\7\4\2\2Q\17\3\2\2\2RS\5\24\13"+
		"\2ST\7\22\2\2T\21\3\2\2\2UV\7\5\2\2VX\7\16\2\2WY\5\20\t\2XW\3\2\2\2XY"+
		"\3\2\2\2YZ\3\2\2\2Z\\\7\6\2\2[]\5\30\r\2\\[\3\2\2\2\\]\3\2\2\2]\23\3\2"+
		"\2\2^_\t\2\2\2_\25\3\2\2\2`a\7\t\2\2a\27\3\2\2\2bc\5\26\f\2cd\7\f\2\2"+
		"d\31\3\2\2\2ef\t\3\2\2f\33\3\2\2\2\f\37&,\66?AINX\\";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}