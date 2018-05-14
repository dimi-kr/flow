// Generated from ./parser/Flow.g4 by ANTLR 4.7.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class FlowLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.7.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, TEXT=7, AND=8, STRING=9, 
		LINK_PREV=10, LINK_NEXT=11, END=12, ID=13, COMMENT=14, WS=15;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	public static final String[] ruleNames = {
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "TEXT", "AND", "STRING", 
		"LINK_PREV", "LINK_NEXT", "END", "ID", "COMMENT", "WS"
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


	public FlowLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Flow.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\21i\b\1\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\3\2\3\2\3\3\3\3\3\4"+
		"\3\4\3\5\3\5\3\6\3\6\3\6\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\t\3\t\3\n\3\n\3"+
		"\n\3\n\7\n:\n\n\f\n\16\n=\13\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3"+
		"\f\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\16\3\16\7\16Q\n\16\f\16\16\16T\13\16"+
		"\3\17\3\17\3\17\3\17\6\17Z\n\17\r\17\16\17[\3\17\5\17_\n\17\3\17\3\17"+
		"\3\20\6\20d\n\20\r\20\16\20e\3\20\3\20\4;[\2\21\3\3\5\4\7\5\t\6\13\7\r"+
		"\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21\3\2\6\4\2C\\c|\5"+
		"\2\62;C\\c|\3\3\f\f\5\2\13\f\17\17\"\"\2m\2\3\3\2\2\2\2\5\3\2\2\2\2\7"+
		"\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2"+
		"\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2"+
		"\35\3\2\2\2\2\37\3\2\2\2\3!\3\2\2\2\5#\3\2\2\2\7%\3\2\2\2\t\'\3\2\2\2"+
		"\13)\3\2\2\2\r,\3\2\2\2\17/\3\2\2\2\21\63\3\2\2\2\23\65\3\2\2\2\25@\3"+
		"\2\2\2\27F\3\2\2\2\31L\3\2\2\2\33N\3\2\2\2\35U\3\2\2\2\37c\3\2\2\2!\""+
		"\7>\2\2\"\4\3\2\2\2#$\7@\2\2$\6\3\2\2\2%&\7]\2\2&\b\3\2\2\2\'(\7_\2\2"+
		"(\n\3\2\2\2)*\7/\2\2*+\7@\2\2+\f\3\2\2\2,-\7/\2\2-.\7/\2\2.\16\3\2\2\2"+
		"/\60\7$\2\2\60\61\5\23\n\2\61\62\7$\2\2\62\20\3\2\2\2\63\64\7.\2\2\64"+
		"\22\3\2\2\2\65;\7$\2\2\66\67\7^\2\2\67:\7$\2\28:\13\2\2\29\66\3\2\2\2"+
		"98\3\2\2\2:=\3\2\2\2;<\3\2\2\2;9\3\2\2\2<>\3\2\2\2=;\3\2\2\2>?\7$\2\2"+
		"?\24\3\2\2\2@A\7&\2\2AB\7r\2\2BC\7t\2\2CD\7g\2\2DE\7x\2\2E\26\3\2\2\2"+
		"FG\7&\2\2GH\7p\2\2HI\7g\2\2IJ\7z\2\2JK\7v\2\2K\30\3\2\2\2LM\7=\2\2M\32"+
		"\3\2\2\2NR\t\2\2\2OQ\t\3\2\2PO\3\2\2\2QT\3\2\2\2RP\3\2\2\2RS\3\2\2\2S"+
		"\34\3\2\2\2TR\3\2\2\2UV\7\61\2\2VW\7\61\2\2WY\3\2\2\2XZ\13\2\2\2YX\3\2"+
		"\2\2Z[\3\2\2\2[\\\3\2\2\2[Y\3\2\2\2\\^\3\2\2\2]_\t\4\2\2^]\3\2\2\2_`\3"+
		"\2\2\2`a\b\17\2\2a\36\3\2\2\2bd\t\5\2\2cb\3\2\2\2de\3\2\2\2ec\3\2\2\2"+
		"ef\3\2\2\2fg\3\2\2\2gh\b\20\2\2h \3\2\2\2\t\29;R[^e\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}