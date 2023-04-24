// Generated from /Users/daniel/compiladores/bistat/parser/Bistat.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class BistatLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, T__22=23, T__23=24, 
		T__24=25, T__25=26, T__26=27, T__27=28, T__28=29, T__29=30, T__30=31, 
		T__31=32, T__32=33, T__33=34, T__34=35, T__35=36, T__36=37, T__37=38, 
		T__38=39, T__39=40, T__40=41, T__41=42, T__42=43, T__43=44, WS=45, CARDINALITY=46, 
		TYPE_PRIMITIVE=47, ID=48, VAR_CONS=49, INT_CONS=50, NUMBER=51, BOOL_CONS=52, 
		STRING_CONS=53, FLOAT_CONS=54, OP_SEC=55, OP_FIRST=56, LOGIC_OPERATOR=57;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
			"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16", 
			"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24", 
			"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32", 
			"T__33", "T__34", "T__35", "T__36", "T__37", "T__38", "T__39", "T__40", 
			"T__41", "T__42", "T__43", "WS", "CARDINALITY", "TYPE_PRIMITIVE", "ID", 
			"VAR_CONS", "INT_CONS", "NUMBER", "BOOL_CONS", "STRING_CONS", "FLOAT_CONS", 
			"OP_SEC", "OP_FIRST", "LOGIC_OPERATOR", "Alpha"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'Program'", "';'", "'var'", "'func'", "'('", "')'", "':'", "'{'", 
			"'}'", "'main'", "'='", "'['", "','", "']'", "'#'", "'for'", "'in'", 
			"'while'", "'if'", "'else'", "'return'", "'read'", "'print'", "'listAdd'", 
			"'listPop'", "'length'", "'range'", "'plot'", "'sum'", "'min'", "'prod'", 
			"'avg'", "'sMode'", "'median'", "'sin'", "'tan'", "'cos'", "'sort'", 
			"'sqrt'", "'floor'", "'ceil'", "'abs'", "'not'", "'-'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, "WS", "CARDINALITY", 
			"TYPE_PRIMITIVE", "ID", "VAR_CONS", "INT_CONS", "NUMBER", "BOOL_CONS", 
			"STRING_CONS", "FLOAT_CONS", "OP_SEC", "OP_FIRST", "LOGIC_OPERATOR"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
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


	public BistatLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Bistat.g4"; }

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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2;\u01b1\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\3\2\3\2\3\2"+
		"\3\2\3\2\3\2\3\2\3\2\3\3\3\3\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\6\3"+
		"\6\3\7\3\7\3\b\3\b\3\t\3\t\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3"+
		"\r\3\r\3\16\3\16\3\17\3\17\3\20\3\20\3\21\3\21\3\21\3\21\3\22\3\22\3\22"+
		"\3\23\3\23\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\25"+
		"\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\27\3\27\3\27\3\27\3\27\3\30\3\30"+
		"\3\30\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\32\3\32"+
		"\3\32\3\32\3\32\3\32\3\32\3\32\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\34"+
		"\3\34\3\34\3\34\3\34\3\34\3\35\3\35\3\35\3\35\3\35\3\36\3\36\3\36\3\36"+
		"\3\37\3\37\3\37\3\37\3 \3 \3 \3 \3 \3!\3!\3!\3!\3\"\3\"\3\"\3\"\3\"\3"+
		"\"\3#\3#\3#\3#\3#\3#\3#\3$\3$\3$\3$\3%\3%\3%\3%\3&\3&\3&\3&\3\'\3\'\3"+
		"\'\3\'\3\'\3(\3(\3(\3(\3(\3)\3)\3)\3)\3)\3)\3*\3*\3*\3*\3*\3+\3+\3+\3"+
		"+\3,\3,\3,\3,\3-\3-\3.\6.\u0137\n.\r.\16.\u0138\3.\3.\3/\3/\5/\u013f\n"+
		"/\3/\3/\3/\3/\5/\u0145\n/\3/\5/\u0148\n/\3\60\3\60\3\60\3\60\3\60\3\60"+
		"\3\60\3\60\3\60\3\60\3\60\3\60\3\60\3\60\3\60\3\60\3\60\3\60\3\60\3\60"+
		"\3\60\3\60\5\60\u0160\n\60\3\61\3\61\6\61\u0164\n\61\r\61\16\61\u0165"+
		"\3\61\3\61\3\61\7\61\u016b\n\61\f\61\16\61\u016e\13\61\3\62\3\62\3\62"+
		"\3\62\3\62\5\62\u0175\n\62\3\63\6\63\u0178\n\63\r\63\16\63\u0179\3\64"+
		"\3\64\3\65\3\65\3\65\3\65\3\65\3\65\3\65\3\65\3\65\5\65\u0187\n\65\3\66"+
		"\3\66\7\66\u018b\n\66\f\66\16\66\u018e\13\66\3\66\3\66\3\67\6\67\u0193"+
		"\n\67\r\67\16\67\u0194\3\67\3\67\6\67\u0199\n\67\r\67\16\67\u019a\38\3"+
		"8\39\39\3:\3:\3:\3:\3:\3:\3:\3:\3:\3:\3:\3:\3:\5:\u01ae\n:\3;\3;\2\2<"+
		"\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20"+
		"\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37"+
		"= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62c\63e\64g\65i\66k\67m8o"+
		"9q:s;u\2\3\2\b\5\2\13\f\17\17\"\"\3\2$$\4\2--//\5\2\'\',,\61\61\4\2>>"+
		"@@\4\2C\\c|\2\u01cb\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2"+
		"\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3"+
		"\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2"+
		"\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2"+
		"\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2"+
		"\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2"+
		"\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q"+
		"\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2"+
		"\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2i\3\2\2\2"+
		"\2k\3\2\2\2\2m\3\2\2\2\2o\3\2\2\2\2q\3\2\2\2\2s\3\2\2\2\3w\3\2\2\2\5\177"+
		"\3\2\2\2\7\u0081\3\2\2\2\t\u0085\3\2\2\2\13\u008a\3\2\2\2\r\u008c\3\2"+
		"\2\2\17\u008e\3\2\2\2\21\u0090\3\2\2\2\23\u0092\3\2\2\2\25\u0094\3\2\2"+
		"\2\27\u0099\3\2\2\2\31\u009b\3\2\2\2\33\u009d\3\2\2\2\35\u009f\3\2\2\2"+
		"\37\u00a1\3\2\2\2!\u00a3\3\2\2\2#\u00a7\3\2\2\2%\u00aa\3\2\2\2\'\u00b0"+
		"\3\2\2\2)\u00b3\3\2\2\2+\u00b8\3\2\2\2-\u00bf\3\2\2\2/\u00c4\3\2\2\2\61"+
		"\u00ca\3\2\2\2\63\u00d2\3\2\2\2\65\u00da\3\2\2\2\67\u00e1\3\2\2\29\u00e7"+
		"\3\2\2\2;\u00ec\3\2\2\2=\u00f0\3\2\2\2?\u00f4\3\2\2\2A\u00f9\3\2\2\2C"+
		"\u00fd\3\2\2\2E\u0103\3\2\2\2G\u010a\3\2\2\2I\u010e\3\2\2\2K\u0112\3\2"+
		"\2\2M\u0116\3\2\2\2O\u011b\3\2\2\2Q\u0120\3\2\2\2S\u0126\3\2\2\2U\u012b"+
		"\3\2\2\2W\u012f\3\2\2\2Y\u0133\3\2\2\2[\u0136\3\2\2\2]\u013c\3\2\2\2_"+
		"\u015f\3\2\2\2a\u0163\3\2\2\2c\u0174\3\2\2\2e\u0177\3\2\2\2g\u017b\3\2"+
		"\2\2i\u0186\3\2\2\2k\u0188\3\2\2\2m\u0192\3\2\2\2o\u019c\3\2\2\2q\u019e"+
		"\3\2\2\2s\u01ad\3\2\2\2u\u01af\3\2\2\2wx\7R\2\2xy\7t\2\2yz\7q\2\2z{\7"+
		"i\2\2{|\7t\2\2|}\7c\2\2}~\7o\2\2~\4\3\2\2\2\177\u0080\7=\2\2\u0080\6\3"+
		"\2\2\2\u0081\u0082\7x\2\2\u0082\u0083\7c\2\2\u0083\u0084\7t\2\2\u0084"+
		"\b\3\2\2\2\u0085\u0086\7h\2\2\u0086\u0087\7w\2\2\u0087\u0088\7p\2\2\u0088"+
		"\u0089\7e\2\2\u0089\n\3\2\2\2\u008a\u008b\7*\2\2\u008b\f\3\2\2\2\u008c"+
		"\u008d\7+\2\2\u008d\16\3\2\2\2\u008e\u008f\7<\2\2\u008f\20\3\2\2\2\u0090"+
		"\u0091\7}\2\2\u0091\22\3\2\2\2\u0092\u0093\7\177\2\2\u0093\24\3\2\2\2"+
		"\u0094\u0095\7o\2\2\u0095\u0096\7c\2\2\u0096\u0097\7k\2\2\u0097\u0098"+
		"\7p\2\2\u0098\26\3\2\2\2\u0099\u009a\7?\2\2\u009a\30\3\2\2\2\u009b\u009c"+
		"\7]\2\2\u009c\32\3\2\2\2\u009d\u009e\7.\2\2\u009e\34\3\2\2\2\u009f\u00a0"+
		"\7_\2\2\u00a0\36\3\2\2\2\u00a1\u00a2\7%\2\2\u00a2 \3\2\2\2\u00a3\u00a4"+
		"\7h\2\2\u00a4\u00a5\7q\2\2\u00a5\u00a6\7t\2\2\u00a6\"\3\2\2\2\u00a7\u00a8"+
		"\7k\2\2\u00a8\u00a9\7p\2\2\u00a9$\3\2\2\2\u00aa\u00ab\7y\2\2\u00ab\u00ac"+
		"\7j\2\2\u00ac\u00ad\7k\2\2\u00ad\u00ae\7n\2\2\u00ae\u00af\7g\2\2\u00af"+
		"&\3\2\2\2\u00b0\u00b1\7k\2\2\u00b1\u00b2\7h\2\2\u00b2(\3\2\2\2\u00b3\u00b4"+
		"\7g\2\2\u00b4\u00b5\7n\2\2\u00b5\u00b6\7u\2\2\u00b6\u00b7\7g\2\2\u00b7"+
		"*\3\2\2\2\u00b8\u00b9\7t\2\2\u00b9\u00ba\7g\2\2\u00ba\u00bb\7v\2\2\u00bb"+
		"\u00bc\7w\2\2\u00bc\u00bd\7t\2\2\u00bd\u00be\7p\2\2\u00be,\3\2\2\2\u00bf"+
		"\u00c0\7t\2\2\u00c0\u00c1\7g\2\2\u00c1\u00c2\7c\2\2\u00c2\u00c3\7f\2\2"+
		"\u00c3.\3\2\2\2\u00c4\u00c5\7r\2\2\u00c5\u00c6\7t\2\2\u00c6\u00c7\7k\2"+
		"\2\u00c7\u00c8\7p\2\2\u00c8\u00c9\7v\2\2\u00c9\60\3\2\2\2\u00ca\u00cb"+
		"\7n\2\2\u00cb\u00cc\7k\2\2\u00cc\u00cd\7u\2\2\u00cd\u00ce\7v\2\2\u00ce"+
		"\u00cf\7C\2\2\u00cf\u00d0\7f\2\2\u00d0\u00d1\7f\2\2\u00d1\62\3\2\2\2\u00d2"+
		"\u00d3\7n\2\2\u00d3\u00d4\7k\2\2\u00d4\u00d5\7u\2\2\u00d5\u00d6\7v\2\2"+
		"\u00d6\u00d7\7R\2\2\u00d7\u00d8\7q\2\2\u00d8\u00d9\7r\2\2\u00d9\64\3\2"+
		"\2\2\u00da\u00db\7n\2\2\u00db\u00dc\7g\2\2\u00dc\u00dd\7p\2\2\u00dd\u00de"+
		"\7i\2\2\u00de\u00df\7v\2\2\u00df\u00e0\7j\2\2\u00e0\66\3\2\2\2\u00e1\u00e2"+
		"\7t\2\2\u00e2\u00e3\7c\2\2\u00e3\u00e4\7p\2\2\u00e4\u00e5\7i\2\2\u00e5"+
		"\u00e6\7g\2\2\u00e68\3\2\2\2\u00e7\u00e8\7r\2\2\u00e8\u00e9\7n\2\2\u00e9"+
		"\u00ea\7q\2\2\u00ea\u00eb\7v\2\2\u00eb:\3\2\2\2\u00ec\u00ed\7u\2\2\u00ed"+
		"\u00ee\7w\2\2\u00ee\u00ef\7o\2\2\u00ef<\3\2\2\2\u00f0\u00f1\7o\2\2\u00f1"+
		"\u00f2\7k\2\2\u00f2\u00f3\7p\2\2\u00f3>\3\2\2\2\u00f4\u00f5\7r\2\2\u00f5"+
		"\u00f6\7t\2\2\u00f6\u00f7\7q\2\2\u00f7\u00f8\7f\2\2\u00f8@\3\2\2\2\u00f9"+
		"\u00fa\7c\2\2\u00fa\u00fb\7x\2\2\u00fb\u00fc\7i\2\2\u00fcB\3\2\2\2\u00fd"+
		"\u00fe\7u\2\2\u00fe\u00ff\7O\2\2\u00ff\u0100\7q\2\2\u0100\u0101\7f\2\2"+
		"\u0101\u0102\7g\2\2\u0102D\3\2\2\2\u0103\u0104\7o\2\2\u0104\u0105\7g\2"+
		"\2\u0105\u0106\7f\2\2\u0106\u0107\7k\2\2\u0107\u0108\7c\2\2\u0108\u0109"+
		"\7p\2\2\u0109F\3\2\2\2\u010a\u010b\7u\2\2\u010b\u010c\7k\2\2\u010c\u010d"+
		"\7p\2\2\u010dH\3\2\2\2\u010e\u010f\7v\2\2\u010f\u0110\7c\2\2\u0110\u0111"+
		"\7p\2\2\u0111J\3\2\2\2\u0112\u0113\7e\2\2\u0113\u0114\7q\2\2\u0114\u0115"+
		"\7u\2\2\u0115L\3\2\2\2\u0116\u0117\7u\2\2\u0117\u0118\7q\2\2\u0118\u0119"+
		"\7t\2\2\u0119\u011a\7v\2\2\u011aN\3\2\2\2\u011b\u011c\7u\2\2\u011c\u011d"+
		"\7s\2\2\u011d\u011e\7t\2\2\u011e\u011f\7v\2\2\u011fP\3\2\2\2\u0120\u0121"+
		"\7h\2\2\u0121\u0122\7n\2\2\u0122\u0123\7q\2\2\u0123\u0124\7q\2\2\u0124"+
		"\u0125\7t\2\2\u0125R\3\2\2\2\u0126\u0127\7e\2\2\u0127\u0128\7g\2\2\u0128"+
		"\u0129\7k\2\2\u0129\u012a\7n\2\2\u012aT\3\2\2\2\u012b\u012c\7c\2\2\u012c"+
		"\u012d\7d\2\2\u012d\u012e\7u\2\2\u012eV\3\2\2\2\u012f\u0130\7p\2\2\u0130"+
		"\u0131\7q\2\2\u0131\u0132\7v\2\2\u0132X\3\2\2\2\u0133\u0134\7/\2\2\u0134"+
		"Z\3\2\2\2\u0135\u0137\t\2\2\2\u0136\u0135\3\2\2\2\u0137\u0138\3\2\2\2"+
		"\u0138\u0136\3\2\2\2\u0138\u0139\3\2\2\2\u0139\u013a\3\2\2\2\u013a\u013b"+
		"\b.\2\2\u013b\\\3\2\2\2\u013c\u013e\7]\2\2\u013d\u013f\5e\63\2\u013e\u013d"+
		"\3\2\2\2\u013e\u013f\3\2\2\2\u013f\u0140\3\2\2\2\u0140\u0141\7_\2\2\u0141"+
		"\u0147\3\2\2\2\u0142\u0144\7]\2\2\u0143\u0145\5e\63\2\u0144\u0143\3\2"+
		"\2\2\u0144\u0145\3\2\2\2\u0145\u0146\3\2\2\2\u0146\u0148\7_\2\2\u0147"+
		"\u0142\3\2\2\2\u0147\u0148\3\2\2\2\u0148^\3\2\2\2\u0149\u014a\7k\2\2\u014a"+
		"\u014b\7p\2\2\u014b\u0160\7v\2\2\u014c\u014d\7h\2\2\u014d\u014e\7n\2\2"+
		"\u014e\u014f\7q\2\2\u014f\u0150\7c\2\2\u0150\u0160\7v\2\2\u0151\u0152"+
		"\7u\2\2\u0152\u0153\7v\2\2\u0153\u0154\7t\2\2\u0154\u0155\7k\2\2\u0155"+
		"\u0156\7p\2\2\u0156\u0160\7i\2\2\u0157\u0158\7d\2\2\u0158\u0159\7q\2\2"+
		"\u0159\u015a\7q\2\2\u015a\u0160\7n\2\2\u015b\u015c\7x\2\2\u015c\u015d"+
		"\7q\2\2\u015d\u015e\7k\2\2\u015e\u0160\7f\2\2\u015f\u0149\3\2\2\2\u015f"+
		"\u014c\3\2\2\2\u015f\u0151\3\2\2\2\u015f\u0157\3\2\2\2\u015f\u015b\3\2"+
		"\2\2\u0160`\3\2\2\2\u0161\u0164\7a\2\2\u0162\u0164\5u;\2\u0163\u0161\3"+
		"\2\2\2\u0163\u0162\3\2\2\2\u0164\u0165\3\2\2\2\u0165\u0163\3\2\2\2\u0165"+
		"\u0166\3\2\2\2\u0166\u016c\3\2\2\2\u0167\u016b\5u;\2\u0168\u016b\5g\64"+
		"\2\u0169\u016b\7a\2\2\u016a\u0167\3\2\2\2\u016a\u0168\3\2\2\2\u016a\u0169"+
		"\3\2\2\2\u016b\u016e\3\2\2\2\u016c\u016a\3\2\2\2\u016c\u016d\3\2\2\2\u016d"+
		"b\3\2\2\2\u016e\u016c\3\2\2\2\u016f\u0175\5k\66\2\u0170\u0175\5m\67\2"+
		"\u0171\u0175\5e\63\2\u0172\u0175\5i\65\2\u0173\u0175\5a\61\2\u0174\u016f"+
		"\3\2\2\2\u0174\u0170\3\2\2\2\u0174\u0171\3\2\2\2\u0174\u0172\3\2\2\2\u0174"+
		"\u0173\3\2\2\2\u0175d\3\2\2\2\u0176\u0178\5g\64\2\u0177\u0176\3\2\2\2"+
		"\u0178\u0179\3\2\2\2\u0179\u0177\3\2\2\2\u0179\u017a\3\2\2\2\u017af\3"+
		"\2\2\2\u017b\u017c\4\62;\2\u017ch\3\2\2\2\u017d\u017e\7v\2\2\u017e\u017f"+
		"\7t\2\2\u017f\u0180\7w\2\2\u0180\u0187\7g\2\2\u0181\u0182\7h\2\2\u0182"+
		"\u0183\7c\2\2\u0183\u0184\7n\2\2\u0184\u0185\7u\2\2\u0185\u0187\7g\2\2"+
		"\u0186\u017d\3\2\2\2\u0186\u0181\3\2\2\2\u0187j\3\2\2\2\u0188\u018c\7"+
		"$\2\2\u0189\u018b\n\3\2\2\u018a\u0189\3\2\2\2\u018b\u018e\3\2\2\2\u018c"+
		"\u018a\3\2\2\2\u018c\u018d\3\2\2\2\u018d\u018f\3\2\2\2\u018e\u018c\3\2"+
		"\2\2\u018f\u0190\7$\2\2\u0190l\3\2\2\2\u0191\u0193\5g\64\2\u0192\u0191"+
		"\3\2\2\2\u0193\u0194\3\2\2\2\u0194\u0192\3\2\2\2\u0194\u0195\3\2\2\2\u0195"+
		"\u0196\3\2\2\2\u0196\u0198\7\60\2\2\u0197\u0199\5g\64\2\u0198\u0197\3"+
		"\2\2\2\u0199\u019a\3\2\2\2\u019a\u0198\3\2\2\2\u019a\u019b\3\2\2\2\u019b"+
		"n\3\2\2\2\u019c\u019d\t\4\2\2\u019dp\3\2\2\2\u019e\u019f\t\5\2\2\u019f"+
		"r\3\2\2\2\u01a0\u01ae\t\6\2\2\u01a1\u01a2\7>\2\2\u01a2\u01ae\7?\2\2\u01a3"+
		"\u01a4\7@\2\2\u01a4\u01ae\7?\2\2\u01a5\u01a6\7?\2\2\u01a6\u01ae\7?\2\2"+
		"\u01a7\u01a8\7#\2\2\u01a8\u01ae\7?\2\2\u01a9\u01aa\7(\2\2\u01aa\u01ae"+
		"\7(\2\2\u01ab\u01ac\7~\2\2\u01ac\u01ae\7~\2\2\u01ad\u01a0\3\2\2\2\u01ad"+
		"\u01a1\3\2\2\2\u01ad\u01a3\3\2\2\2\u01ad\u01a5\3\2\2\2\u01ad\u01a7\3\2"+
		"\2\2\u01ad\u01a9\3\2\2\2\u01ad\u01ab\3\2\2\2\u01aet\3\2\2\2\u01af\u01b0"+
		"\t\7\2\2\u01b0v\3\2\2\2\23\2\u0138\u013e\u0144\u0147\u015f\u0163\u0165"+
		"\u016a\u016c\u0174\u0179\u0186\u018c\u0194\u019a\u01ad\3\2\3\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}