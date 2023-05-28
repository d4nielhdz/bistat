// Generated from /Users/daniel/compiladores/bistat/parser/Bistat.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class BistatParser extends Parser {
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
		T__38=39, T__39=40, T__40=41, T__41=42, T__42=43, T__43=44, T__44=45, 
		T__45=46, T__46=47, T__47=48, T__48=49, T__49=50, T__50=51, T__51=52, 
		T__52=53, T__53=54, T__54=55, T__55=56, T__56=57, T__57=58, WS=59, CARDINALITY=60, 
		TYPE_PRIMITIVE=61, BOOL_CONS=62, ID=63, INT_CONS=64, NUMBER=65, STRING_CONS=66, 
		FLOAT_CONS=67;
	public static final int
		RULE_program = 0, RULE_varDeclaration = 1, RULE_var_type = 2, RULE_funcDef = 3, 
		RULE_funcBlockStart = 4, RULE_funcBlockEnd = 5, RULE_paramDeclaration = 6, 
		RULE_main = 7, RULE_stmt = 8, RULE_assignment = 9, RULE_matrixAssignment = 10, 
		RULE_listAssignment = 11, RULE_comment = 12, RULE_forLoop = 13, RULE_forHeader = 14, 
		RULE_forExprEnd = 15, RULE_whileLoop = 16, RULE_whileExprEnd = 17, RULE_conditional = 18, 
		RULE_ifStmt = 19, RULE_elseIfStmt = 20, RULE_condExprEnd = 21, RULE_elseStmt = 22, 
		RULE_returnStmt = 23, RULE_specialFunction = 24, RULE_inputRead = 25, 
		RULE_print = 26, RULE_listAdd = 27, RULE_listPop = 28, RULE_length = 29, 
		RULE_range = 30, RULE_plot = 31, RULE_sum = 32, RULE_min = 33, RULE_prod = 34, 
		RULE_avg = 35, RULE_sMode = 36, RULE_median = 37, RULE_sin = 38, RULE_tan = 39, 
		RULE_cos = 40, RULE_sort = 41, RULE_sqrt = 42, RULE_floor = 43, RULE_ceil = 44, 
		RULE_abs = 45, RULE_not = 46, RULE_listAccess = 47, RULE_listAssign = 48, 
		RULE_expression = 49, RULE_exp = 50, RULE_term = 51, RULE_factor = 52, 
		RULE_unaryMinus = 53, RULE_nestedExpression = 54, RULE_functionCall = 55, 
		RULE_varCons = 56, RULE_opSec = 57, RULE_opFirst = 58, RULE_logicOperator = 59;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "varDeclaration", "var_type", "funcDef", "funcBlockStart", 
			"funcBlockEnd", "paramDeclaration", "main", "stmt", "assignment", "matrixAssignment", 
			"listAssignment", "comment", "forLoop", "forHeader", "forExprEnd", "whileLoop", 
			"whileExprEnd", "conditional", "ifStmt", "elseIfStmt", "condExprEnd", 
			"elseStmt", "returnStmt", "specialFunction", "inputRead", "print", "listAdd", 
			"listPop", "length", "range", "plot", "sum", "min", "prod", "avg", "sMode", 
			"median", "sin", "tan", "cos", "sort", "sqrt", "floor", "ceil", "abs", 
			"not", "listAccess", "listAssign", "expression", "exp", "term", "factor", 
			"unaryMinus", "nestedExpression", "functionCall", "varCons", "opSec", 
			"opFirst", "logicOperator"
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
			"'sqrt'", "'floor'", "'ceil'", "'abs'", "'not'", "'listAccess'", "'listAssign'", 
			"'-'", "'+'", "'/'", "'*'", "'%'", "'<'", "'>'", "'<='", "'>='", "'=='", 
			"'!='", "'&&'", "'||'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, "WS", 
			"CARDINALITY", "TYPE_PRIMITIVE", "BOOL_CONS", "ID", "INT_CONS", "NUMBER", 
			"STRING_CONS", "FLOAT_CONS"
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

	@Override
	public String getGrammarFileName() { return "Bistat.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public BistatParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class ProgramContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(BistatParser.ID, 0); }
		public MainContext main() {
			return getRuleContext(MainContext.class,0);
		}
		public TerminalNode EOF() { return getToken(BistatParser.EOF, 0); }
		public List<VarDeclarationContext> varDeclaration() {
			return getRuleContexts(VarDeclarationContext.class);
		}
		public VarDeclarationContext varDeclaration(int i) {
			return getRuleContext(VarDeclarationContext.class,i);
		}
		public List<FuncDefContext> funcDef() {
			return getRuleContexts(FuncDefContext.class);
		}
		public FuncDefContext funcDef(int i) {
			return getRuleContext(FuncDefContext.class,i);
		}
		public ProgramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program; }
	}

	public final ProgramContext program() throws RecognitionException {
		ProgramContext _localctx = new ProgramContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_program);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(120);
			match(T__0);
			setState(121);
			match(ID);
			setState(122);
			match(T__1);
			setState(126);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__2) {
				{
				{
				setState(123);
				varDeclaration();
				}
				}
				setState(128);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(132);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__3) {
				{
				{
				setState(129);
				funcDef();
				}
				}
				setState(134);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(135);
			main();
			setState(136);
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

	public static class VarDeclarationContext extends ParserRuleContext {
		public Var_typeContext var_type() {
			return getRuleContext(Var_typeContext.class,0);
		}
		public TerminalNode ID() { return getToken(BistatParser.ID, 0); }
		public VarDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varDeclaration; }
	}

	public final VarDeclarationContext varDeclaration() throws RecognitionException {
		VarDeclarationContext _localctx = new VarDeclarationContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_varDeclaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(138);
			match(T__2);
			setState(139);
			var_type();
			setState(140);
			match(ID);
			setState(141);
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

	public static class Var_typeContext extends ParserRuleContext {
		public TerminalNode TYPE_PRIMITIVE() { return getToken(BistatParser.TYPE_PRIMITIVE, 0); }
		public TerminalNode CARDINALITY() { return getToken(BistatParser.CARDINALITY, 0); }
		public Var_typeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_var_type; }
	}

	public final Var_typeContext var_type() throws RecognitionException {
		Var_typeContext _localctx = new Var_typeContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_var_type);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(143);
			match(TYPE_PRIMITIVE);
			{
			setState(145);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==CARDINALITY) {
				{
				setState(144);
				match(CARDINALITY);
				}
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

	public static class FuncDefContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(BistatParser.ID, 0); }
		public TerminalNode TYPE_PRIMITIVE() { return getToken(BistatParser.TYPE_PRIMITIVE, 0); }
		public FuncBlockStartContext funcBlockStart() {
			return getRuleContext(FuncBlockStartContext.class,0);
		}
		public FuncBlockEndContext funcBlockEnd() {
			return getRuleContext(FuncBlockEndContext.class,0);
		}
		public List<ParamDeclarationContext> paramDeclaration() {
			return getRuleContexts(ParamDeclarationContext.class);
		}
		public ParamDeclarationContext paramDeclaration(int i) {
			return getRuleContext(ParamDeclarationContext.class,i);
		}
		public List<VarDeclarationContext> varDeclaration() {
			return getRuleContexts(VarDeclarationContext.class);
		}
		public VarDeclarationContext varDeclaration(int i) {
			return getRuleContext(VarDeclarationContext.class,i);
		}
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public FuncDefContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcDef; }
	}

	public final FuncDefContext funcDef() throws RecognitionException {
		FuncDefContext _localctx = new FuncDefContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_funcDef);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(147);
			match(T__3);
			setState(148);
			match(ID);
			setState(149);
			match(T__4);
			setState(153);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__2) {
				{
				{
				setState(150);
				paramDeclaration();
				}
				}
				setState(155);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(156);
			match(T__5);
			setState(157);
			match(T__6);
			setState(158);
			match(TYPE_PRIMITIVE);
			setState(162);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__2) {
				{
				{
				setState(159);
				varDeclaration();
				}
				}
				setState(164);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(165);
			funcBlockStart();
			setState(167); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(166);
				stmt();
				}
				}
				setState(169); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__14) | (1L << T__15) | (1L << T__17) | (1L << T__18) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << T__42) | (1L << T__43) | (1L << T__44) | (1L << ID))) != 0) );
			setState(171);
			funcBlockEnd();
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

	public static class FuncBlockStartContext extends ParserRuleContext {
		public FuncBlockStartContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcBlockStart; }
	}

	public final FuncBlockStartContext funcBlockStart() throws RecognitionException {
		FuncBlockStartContext _localctx = new FuncBlockStartContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_funcBlockStart);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(173);
			match(T__7);
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

	public static class FuncBlockEndContext extends ParserRuleContext {
		public FuncBlockEndContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcBlockEnd; }
	}

	public final FuncBlockEndContext funcBlockEnd() throws RecognitionException {
		FuncBlockEndContext _localctx = new FuncBlockEndContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_funcBlockEnd);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(175);
			match(T__8);
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

	public static class ParamDeclarationContext extends ParserRuleContext {
		public TerminalNode TYPE_PRIMITIVE() { return getToken(BistatParser.TYPE_PRIMITIVE, 0); }
		public TerminalNode ID() { return getToken(BistatParser.ID, 0); }
		public ParamDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_paramDeclaration; }
	}

	public final ParamDeclarationContext paramDeclaration() throws RecognitionException {
		ParamDeclarationContext _localctx = new ParamDeclarationContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_paramDeclaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(177);
			match(T__2);
			setState(178);
			match(TYPE_PRIMITIVE);
			setState(179);
			match(ID);
			setState(180);
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

	public static class MainContext extends ParserRuleContext {
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public MainContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_main; }
	}

	public final MainContext main() throws RecognitionException {
		MainContext _localctx = new MainContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_main);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(182);
			match(T__9);
			setState(183);
			match(T__4);
			setState(184);
			match(T__5);
			setState(185);
			match(T__7);
			setState(187); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(186);
				stmt();
				}
				}
				setState(189); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__14) | (1L << T__15) | (1L << T__17) | (1L << T__18) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << T__42) | (1L << T__43) | (1L << T__44) | (1L << ID))) != 0) );
			setState(191);
			match(T__8);
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
		public AssignmentContext assignment() {
			return getRuleContext(AssignmentContext.class,0);
		}
		public SpecialFunctionContext specialFunction() {
			return getRuleContext(SpecialFunctionContext.class,0);
		}
		public FunctionCallContext functionCall() {
			return getRuleContext(FunctionCallContext.class,0);
		}
		public ReturnStmtContext returnStmt() {
			return getRuleContext(ReturnStmtContext.class,0);
		}
		public ConditionalContext conditional() {
			return getRuleContext(ConditionalContext.class,0);
		}
		public WhileLoopContext whileLoop() {
			return getRuleContext(WhileLoopContext.class,0);
		}
		public ForLoopContext forLoop() {
			return getRuleContext(ForLoopContext.class,0);
		}
		public CommentContext comment() {
			return getRuleContext(CommentContext.class,0);
		}
		public StmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stmt; }
	}

	public final StmtContext stmt() throws RecognitionException {
		StmtContext _localctx = new StmtContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_stmt);
		try {
			setState(205);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__20:
			case T__21:
			case T__22:
			case T__23:
			case T__24:
			case T__25:
			case T__26:
			case T__27:
			case T__28:
			case T__29:
			case T__30:
			case T__31:
			case T__32:
			case T__33:
			case T__34:
			case T__35:
			case T__36:
			case T__37:
			case T__38:
			case T__39:
			case T__40:
			case T__41:
			case T__42:
			case T__43:
			case T__44:
			case ID:
				enterOuterAlt(_localctx, 1);
				{
				{
				setState(197);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
				case 1:
					{
					setState(193);
					assignment();
					}
					break;
				case 2:
					{
					setState(194);
					specialFunction();
					}
					break;
				case 3:
					{
					setState(195);
					functionCall();
					}
					break;
				case 4:
					{
					setState(196);
					returnStmt();
					}
					break;
				}
				setState(199);
				match(T__1);
				}
				}
				break;
			case T__18:
				enterOuterAlt(_localctx, 2);
				{
				setState(201);
				conditional();
				}
				break;
			case T__17:
				enterOuterAlt(_localctx, 3);
				{
				setState(202);
				whileLoop();
				}
				break;
			case T__15:
				enterOuterAlt(_localctx, 4);
				{
				setState(203);
				forLoop();
				}
				break;
			case T__14:
				enterOuterAlt(_localctx, 5);
				{
				setState(204);
				comment();
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

	public static class AssignmentContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(BistatParser.ID, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ListAssignmentContext listAssignment() {
			return getRuleContext(ListAssignmentContext.class,0);
		}
		public MatrixAssignmentContext matrixAssignment() {
			return getRuleContext(MatrixAssignmentContext.class,0);
		}
		public AssignmentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignment; }
	}

	public final AssignmentContext assignment() throws RecognitionException {
		AssignmentContext _localctx = new AssignmentContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_assignment);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(207);
			match(ID);
			setState(208);
			match(T__10);
			setState(212);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				{
				setState(209);
				expression();
				}
				break;
			case 2:
				{
				setState(210);
				listAssignment();
				}
				break;
			case 3:
				{
				setState(211);
				matrixAssignment();
				}
				break;
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

	public static class MatrixAssignmentContext extends ParserRuleContext {
		public List<ListAssignmentContext> listAssignment() {
			return getRuleContexts(ListAssignmentContext.class);
		}
		public ListAssignmentContext listAssignment(int i) {
			return getRuleContext(ListAssignmentContext.class,i);
		}
		public MatrixAssignmentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_matrixAssignment; }
	}

	public final MatrixAssignmentContext matrixAssignment() throws RecognitionException {
		MatrixAssignmentContext _localctx = new MatrixAssignmentContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_matrixAssignment);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(214);
			match(T__11);
			setState(215);
			listAssignment();
			setState(220);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__12) {
				{
				{
				setState(216);
				match(T__12);
				setState(217);
				listAssignment();
				}
				}
				setState(222);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(223);
			match(T__13);
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

	public static class ListAssignmentContext extends ParserRuleContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public ListAssignmentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listAssignment; }
	}

	public final ListAssignmentContext listAssignment() throws RecognitionException {
		ListAssignmentContext _localctx = new ListAssignmentContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_listAssignment);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(225);
			match(T__11);
			{
			setState(226);
			expression();
			}
			setState(231);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__12) {
				{
				{
				setState(227);
				match(T__12);
				{
				setState(228);
				expression();
				}
				}
				}
				setState(233);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(234);
			match(T__13);
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

	public static class CommentContext extends ParserRuleContext {
		public CommentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_comment; }
	}

	public final CommentContext comment() throws RecognitionException {
		CommentContext _localctx = new CommentContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_comment);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(236);
			match(T__14);
			setState(238); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(237);
				_la = _input.LA(1);
				if ( _la <= 0 || (_la==T__14) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				}
				setState(240); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__1) | (1L << T__2) | (1L << T__3) | (1L << T__4) | (1L << T__5) | (1L << T__6) | (1L << T__7) | (1L << T__8) | (1L << T__9) | (1L << T__10) | (1L << T__11) | (1L << T__12) | (1L << T__13) | (1L << T__15) | (1L << T__16) | (1L << T__17) | (1L << T__18) | (1L << T__19) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << T__42) | (1L << T__43) | (1L << T__44) | (1L << T__45) | (1L << T__46) | (1L << T__47) | (1L << T__48) | (1L << T__49) | (1L << T__50) | (1L << T__51) | (1L << T__52) | (1L << T__53) | (1L << T__54) | (1L << T__55) | (1L << T__56) | (1L << T__57) | (1L << WS) | (1L << CARDINALITY) | (1L << TYPE_PRIMITIVE) | (1L << BOOL_CONS) | (1L << ID))) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & ((1L << (INT_CONS - 64)) | (1L << (NUMBER - 64)) | (1L << (STRING_CONS - 64)) | (1L << (FLOAT_CONS - 64)))) != 0) );
			setState(242);
			match(T__14);
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

	public static class ForLoopContext extends ParserRuleContext {
		public ForHeaderContext forHeader() {
			return getRuleContext(ForHeaderContext.class,0);
		}
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public ForLoopContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forLoop; }
	}

	public final ForLoopContext forLoop() throws RecognitionException {
		ForLoopContext _localctx = new ForLoopContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_forLoop);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(244);
			forHeader();
			setState(245);
			match(T__7);
			setState(247); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(246);
				stmt();
				}
				}
				setState(249); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__14) | (1L << T__15) | (1L << T__17) | (1L << T__18) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << T__42) | (1L << T__43) | (1L << T__44) | (1L << ID))) != 0) );
			setState(251);
			match(T__8);
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

	public static class ForHeaderContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(BistatParser.ID, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ForHeaderContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forHeader; }
	}

	public final ForHeaderContext forHeader() throws RecognitionException {
		ForHeaderContext _localctx = new ForHeaderContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_forHeader);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(253);
			match(T__15);
			setState(254);
			match(T__4);
			setState(255);
			match(ID);
			setState(256);
			match(T__16);
			setState(257);
			expression();
			setState(258);
			match(T__5);
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

	public static class ForExprEndContext extends ParserRuleContext {
		public ForExprEndContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forExprEnd; }
	}

	public final ForExprEndContext forExprEnd() throws RecognitionException {
		ForExprEndContext _localctx = new ForExprEndContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_forExprEnd);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(260);
			match(T__5);
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

	public static class WhileLoopContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public WhileExprEndContext whileExprEnd() {
			return getRuleContext(WhileExprEndContext.class,0);
		}
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public WhileLoopContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_whileLoop; }
	}

	public final WhileLoopContext whileLoop() throws RecognitionException {
		WhileLoopContext _localctx = new WhileLoopContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_whileLoop);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(262);
			match(T__17);
			setState(263);
			match(T__4);
			setState(264);
			expression();
			setState(265);
			whileExprEnd();
			setState(266);
			match(T__7);
			setState(268); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(267);
				stmt();
				}
				}
				setState(270); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__14) | (1L << T__15) | (1L << T__17) | (1L << T__18) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << T__42) | (1L << T__43) | (1L << T__44) | (1L << ID))) != 0) );
			setState(272);
			match(T__8);
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

	public static class WhileExprEndContext extends ParserRuleContext {
		public WhileExprEndContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_whileExprEnd; }
	}

	public final WhileExprEndContext whileExprEnd() throws RecognitionException {
		WhileExprEndContext _localctx = new WhileExprEndContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_whileExprEnd);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(274);
			match(T__5);
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

	public static class ConditionalContext extends ParserRuleContext {
		public IfStmtContext ifStmt() {
			return getRuleContext(IfStmtContext.class,0);
		}
		public List<ElseIfStmtContext> elseIfStmt() {
			return getRuleContexts(ElseIfStmtContext.class);
		}
		public ElseIfStmtContext elseIfStmt(int i) {
			return getRuleContext(ElseIfStmtContext.class,i);
		}
		public ElseStmtContext elseStmt() {
			return getRuleContext(ElseStmtContext.class,0);
		}
		public ConditionalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_conditional; }
	}

	public final ConditionalContext conditional() throws RecognitionException {
		ConditionalContext _localctx = new ConditionalContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_conditional);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(276);
			ifStmt();
			setState(280);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(277);
					elseIfStmt();
					}
					} 
				}
				setState(282);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
			}
			setState(284);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__19) {
				{
				setState(283);
				elseStmt();
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

	public static class IfStmtContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public CondExprEndContext condExprEnd() {
			return getRuleContext(CondExprEndContext.class,0);
		}
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public IfStmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifStmt; }
	}

	public final IfStmtContext ifStmt() throws RecognitionException {
		IfStmtContext _localctx = new IfStmtContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_ifStmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(286);
			match(T__18);
			setState(287);
			match(T__4);
			setState(288);
			expression();
			setState(289);
			condExprEnd();
			setState(290);
			match(T__7);
			setState(292); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(291);
				stmt();
				}
				}
				setState(294); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__14) | (1L << T__15) | (1L << T__17) | (1L << T__18) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << T__42) | (1L << T__43) | (1L << T__44) | (1L << ID))) != 0) );
			setState(296);
			match(T__8);
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

	public static class ElseIfStmtContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public CondExprEndContext condExprEnd() {
			return getRuleContext(CondExprEndContext.class,0);
		}
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public ElseIfStmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_elseIfStmt; }
	}

	public final ElseIfStmtContext elseIfStmt() throws RecognitionException {
		ElseIfStmtContext _localctx = new ElseIfStmtContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_elseIfStmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(298);
			match(T__19);
			setState(299);
			match(T__18);
			setState(300);
			match(T__4);
			setState(301);
			expression();
			setState(302);
			condExprEnd();
			setState(303);
			match(T__7);
			setState(305); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(304);
				stmt();
				}
				}
				setState(307); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__14) | (1L << T__15) | (1L << T__17) | (1L << T__18) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << T__42) | (1L << T__43) | (1L << T__44) | (1L << ID))) != 0) );
			setState(309);
			match(T__8);
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

	public static class CondExprEndContext extends ParserRuleContext {
		public CondExprEndContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_condExprEnd; }
	}

	public final CondExprEndContext condExprEnd() throws RecognitionException {
		CondExprEndContext _localctx = new CondExprEndContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_condExprEnd);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(311);
			match(T__5);
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

	public static class ElseStmtContext extends ParserRuleContext {
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public ElseStmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_elseStmt; }
	}

	public final ElseStmtContext elseStmt() throws RecognitionException {
		ElseStmtContext _localctx = new ElseStmtContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_elseStmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(313);
			match(T__19);
			setState(314);
			match(T__7);
			setState(316); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(315);
				stmt();
				}
				}
				setState(318); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__14) | (1L << T__15) | (1L << T__17) | (1L << T__18) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << T__42) | (1L << T__43) | (1L << T__44) | (1L << ID))) != 0) );
			setState(320);
			match(T__8);
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

	public static class ReturnStmtContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ReturnStmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_returnStmt; }
	}

	public final ReturnStmtContext returnStmt() throws RecognitionException {
		ReturnStmtContext _localctx = new ReturnStmtContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_returnStmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(322);
			match(T__20);
			setState(323);
			expression();
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

	public static class SpecialFunctionContext extends ParserRuleContext {
		public InputReadContext inputRead() {
			return getRuleContext(InputReadContext.class,0);
		}
		public PrintContext print() {
			return getRuleContext(PrintContext.class,0);
		}
		public ListAddContext listAdd() {
			return getRuleContext(ListAddContext.class,0);
		}
		public ListAccessContext listAccess() {
			return getRuleContext(ListAccessContext.class,0);
		}
		public ListAssignContext listAssign() {
			return getRuleContext(ListAssignContext.class,0);
		}
		public ListPopContext listPop() {
			return getRuleContext(ListPopContext.class,0);
		}
		public LengthContext length() {
			return getRuleContext(LengthContext.class,0);
		}
		public RangeContext range() {
			return getRuleContext(RangeContext.class,0);
		}
		public PlotContext plot() {
			return getRuleContext(PlotContext.class,0);
		}
		public SumContext sum() {
			return getRuleContext(SumContext.class,0);
		}
		public MinContext min() {
			return getRuleContext(MinContext.class,0);
		}
		public ProdContext prod() {
			return getRuleContext(ProdContext.class,0);
		}
		public AvgContext avg() {
			return getRuleContext(AvgContext.class,0);
		}
		public SModeContext sMode() {
			return getRuleContext(SModeContext.class,0);
		}
		public MedianContext median() {
			return getRuleContext(MedianContext.class,0);
		}
		public SinContext sin() {
			return getRuleContext(SinContext.class,0);
		}
		public CosContext cos() {
			return getRuleContext(CosContext.class,0);
		}
		public TanContext tan() {
			return getRuleContext(TanContext.class,0);
		}
		public SortContext sort() {
			return getRuleContext(SortContext.class,0);
		}
		public SqrtContext sqrt() {
			return getRuleContext(SqrtContext.class,0);
		}
		public FloorContext floor() {
			return getRuleContext(FloorContext.class,0);
		}
		public CeilContext ceil() {
			return getRuleContext(CeilContext.class,0);
		}
		public AbsContext abs() {
			return getRuleContext(AbsContext.class,0);
		}
		public NotContext not() {
			return getRuleContext(NotContext.class,0);
		}
		public SpecialFunctionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_specialFunction; }
	}

	public final SpecialFunctionContext specialFunction() throws RecognitionException {
		SpecialFunctionContext _localctx = new SpecialFunctionContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_specialFunction);
		try {
			setState(349);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__21:
				enterOuterAlt(_localctx, 1);
				{
				setState(325);
				inputRead();
				}
				break;
			case T__22:
				enterOuterAlt(_localctx, 2);
				{
				setState(326);
				print();
				}
				break;
			case T__23:
				enterOuterAlt(_localctx, 3);
				{
				setState(327);
				listAdd();
				}
				break;
			case T__43:
				enterOuterAlt(_localctx, 4);
				{
				setState(328);
				listAccess();
				}
				break;
			case T__44:
				enterOuterAlt(_localctx, 5);
				{
				setState(329);
				listAssign();
				}
				break;
			case T__24:
				enterOuterAlt(_localctx, 6);
				{
				setState(330);
				listPop();
				}
				break;
			case T__25:
				enterOuterAlt(_localctx, 7);
				{
				setState(331);
				length();
				}
				break;
			case T__26:
				enterOuterAlt(_localctx, 8);
				{
				setState(332);
				range();
				}
				break;
			case T__27:
				enterOuterAlt(_localctx, 9);
				{
				setState(333);
				plot();
				}
				break;
			case T__28:
				enterOuterAlt(_localctx, 10);
				{
				setState(334);
				sum();
				}
				break;
			case T__29:
				enterOuterAlt(_localctx, 11);
				{
				setState(335);
				min();
				}
				break;
			case T__30:
				enterOuterAlt(_localctx, 12);
				{
				setState(336);
				prod();
				}
				break;
			case T__31:
				enterOuterAlt(_localctx, 13);
				{
				setState(337);
				avg();
				}
				break;
			case T__32:
				enterOuterAlt(_localctx, 14);
				{
				setState(338);
				sMode();
				}
				break;
			case T__33:
				enterOuterAlt(_localctx, 15);
				{
				setState(339);
				median();
				}
				break;
			case T__34:
				enterOuterAlt(_localctx, 16);
				{
				setState(340);
				sin();
				}
				break;
			case T__36:
				enterOuterAlt(_localctx, 17);
				{
				setState(341);
				cos();
				}
				break;
			case T__35:
				enterOuterAlt(_localctx, 18);
				{
				setState(342);
				tan();
				}
				break;
			case T__37:
				enterOuterAlt(_localctx, 19);
				{
				setState(343);
				sort();
				}
				break;
			case T__38:
				enterOuterAlt(_localctx, 20);
				{
				setState(344);
				sqrt();
				}
				break;
			case T__39:
				enterOuterAlt(_localctx, 21);
				{
				setState(345);
				floor();
				}
				break;
			case T__40:
				enterOuterAlt(_localctx, 22);
				{
				setState(346);
				ceil();
				}
				break;
			case T__41:
				enterOuterAlt(_localctx, 23);
				{
				setState(347);
				abs();
				}
				break;
			case T__42:
				enterOuterAlt(_localctx, 24);
				{
				setState(348);
				not();
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

	public static class InputReadContext extends ParserRuleContext {
		public List<TerminalNode> ID() { return getTokens(BistatParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(BistatParser.ID, i);
		}
		public InputReadContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_inputRead; }
	}

	public final InputReadContext inputRead() throws RecognitionException {
		InputReadContext _localctx = new InputReadContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_inputRead);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(351);
			match(T__21);
			setState(352);
			match(T__4);
			setState(353);
			match(ID);
			setState(358);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__12) {
				{
				{
				setState(354);
				match(T__12);
				setState(355);
				match(ID);
				}
				}
				setState(360);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(361);
			match(T__5);
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

	public static class PrintContext extends ParserRuleContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public PrintContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_print; }
	}

	public final PrintContext print() throws RecognitionException {
		PrintContext _localctx = new PrintContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_print);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(363);
			match(T__22);
			setState(364);
			match(T__4);
			setState(365);
			expression();
			setState(370);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__12) {
				{
				{
				setState(366);
				match(T__12);
				setState(367);
				expression();
				}
				}
				setState(372);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(373);
			match(T__5);
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

	public static class ListAddContext extends ParserRuleContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public ListAddContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listAdd; }
	}

	public final ListAddContext listAdd() throws RecognitionException {
		ListAddContext _localctx = new ListAddContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_listAdd);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(375);
			match(T__23);
			setState(376);
			match(T__4);
			setState(377);
			expression();
			setState(378);
			match(T__12);
			setState(379);
			expression();
			setState(380);
			match(T__5);
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

	public static class ListPopContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ListPopContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listPop; }
	}

	public final ListPopContext listPop() throws RecognitionException {
		ListPopContext _localctx = new ListPopContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_listPop);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(382);
			match(T__24);
			setState(383);
			match(T__4);
			setState(384);
			expression();
			setState(385);
			match(T__5);
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

	public static class LengthContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public LengthContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_length; }
	}

	public final LengthContext length() throws RecognitionException {
		LengthContext _localctx = new LengthContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_length);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(387);
			match(T__25);
			setState(388);
			match(T__4);
			setState(389);
			expression();
			setState(390);
			match(T__5);
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

	public static class RangeContext extends ParserRuleContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public RangeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_range; }
	}

	public final RangeContext range() throws RecognitionException {
		RangeContext _localctx = new RangeContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_range);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(392);
			match(T__26);
			setState(393);
			match(T__4);
			setState(394);
			expression();
			setState(397);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__12) {
				{
				setState(395);
				match(T__12);
				setState(396);
				expression();
				}
			}

			setState(399);
			match(T__5);
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

	public static class PlotContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public PlotContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_plot; }
	}

	public final PlotContext plot() throws RecognitionException {
		PlotContext _localctx = new PlotContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_plot);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(401);
			match(T__27);
			setState(402);
			match(T__4);
			setState(403);
			expression();
			setState(404);
			match(T__5);
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

	public static class SumContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public SumContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sum; }
	}

	public final SumContext sum() throws RecognitionException {
		SumContext _localctx = new SumContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_sum);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(406);
			match(T__28);
			setState(407);
			match(T__4);
			setState(408);
			expression();
			setState(409);
			match(T__5);
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

	public static class MinContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public MinContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_min; }
	}

	public final MinContext min() throws RecognitionException {
		MinContext _localctx = new MinContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_min);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(411);
			match(T__29);
			setState(412);
			match(T__4);
			setState(413);
			expression();
			setState(414);
			match(T__5);
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

	public static class ProdContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ProdContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_prod; }
	}

	public final ProdContext prod() throws RecognitionException {
		ProdContext _localctx = new ProdContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_prod);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(416);
			match(T__30);
			setState(417);
			match(T__4);
			setState(418);
			expression();
			setState(419);
			match(T__5);
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

	public static class AvgContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public AvgContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_avg; }
	}

	public final AvgContext avg() throws RecognitionException {
		AvgContext _localctx = new AvgContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_avg);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(421);
			match(T__31);
			setState(422);
			match(T__4);
			setState(423);
			expression();
			setState(424);
			match(T__5);
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

	public static class SModeContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public SModeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sMode; }
	}

	public final SModeContext sMode() throws RecognitionException {
		SModeContext _localctx = new SModeContext(_ctx, getState());
		enterRule(_localctx, 72, RULE_sMode);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(426);
			match(T__32);
			setState(427);
			match(T__4);
			setState(428);
			expression();
			setState(429);
			match(T__5);
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

	public static class MedianContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public MedianContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_median; }
	}

	public final MedianContext median() throws RecognitionException {
		MedianContext _localctx = new MedianContext(_ctx, getState());
		enterRule(_localctx, 74, RULE_median);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(431);
			match(T__33);
			setState(432);
			match(T__4);
			setState(433);
			expression();
			setState(434);
			match(T__5);
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

	public static class SinContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public SinContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sin; }
	}

	public final SinContext sin() throws RecognitionException {
		SinContext _localctx = new SinContext(_ctx, getState());
		enterRule(_localctx, 76, RULE_sin);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(436);
			match(T__34);
			setState(437);
			match(T__4);
			setState(438);
			expression();
			setState(439);
			match(T__5);
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

	public static class TanContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TanContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tan; }
	}

	public final TanContext tan() throws RecognitionException {
		TanContext _localctx = new TanContext(_ctx, getState());
		enterRule(_localctx, 78, RULE_tan);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(441);
			match(T__35);
			setState(442);
			match(T__4);
			setState(443);
			expression();
			setState(444);
			match(T__5);
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

	public static class CosContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public CosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cos; }
	}

	public final CosContext cos() throws RecognitionException {
		CosContext _localctx = new CosContext(_ctx, getState());
		enterRule(_localctx, 80, RULE_cos);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(446);
			match(T__36);
			setState(447);
			match(T__4);
			setState(448);
			expression();
			setState(449);
			match(T__5);
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

	public static class SortContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public SortContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sort; }
	}

	public final SortContext sort() throws RecognitionException {
		SortContext _localctx = new SortContext(_ctx, getState());
		enterRule(_localctx, 82, RULE_sort);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(451);
			match(T__37);
			setState(452);
			match(T__4);
			setState(453);
			expression();
			setState(454);
			match(T__5);
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

	public static class SqrtContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public SqrtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sqrt; }
	}

	public final SqrtContext sqrt() throws RecognitionException {
		SqrtContext _localctx = new SqrtContext(_ctx, getState());
		enterRule(_localctx, 84, RULE_sqrt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(456);
			match(T__38);
			setState(457);
			match(T__4);
			setState(458);
			expression();
			setState(459);
			match(T__5);
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

	public static class FloorContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public FloorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_floor; }
	}

	public final FloorContext floor() throws RecognitionException {
		FloorContext _localctx = new FloorContext(_ctx, getState());
		enterRule(_localctx, 86, RULE_floor);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(461);
			match(T__39);
			setState(462);
			match(T__4);
			setState(463);
			expression();
			setState(464);
			match(T__5);
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

	public static class CeilContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public CeilContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ceil; }
	}

	public final CeilContext ceil() throws RecognitionException {
		CeilContext _localctx = new CeilContext(_ctx, getState());
		enterRule(_localctx, 88, RULE_ceil);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(466);
			match(T__40);
			setState(467);
			match(T__4);
			setState(468);
			expression();
			setState(469);
			match(T__5);
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

	public static class AbsContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public AbsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_abs; }
	}

	public final AbsContext abs() throws RecognitionException {
		AbsContext _localctx = new AbsContext(_ctx, getState());
		enterRule(_localctx, 90, RULE_abs);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(471);
			match(T__41);
			setState(472);
			match(T__4);
			setState(473);
			expression();
			setState(474);
			match(T__5);
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

	public static class NotContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public NotContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_not; }
	}

	public final NotContext not() throws RecognitionException {
		NotContext _localctx = new NotContext(_ctx, getState());
		enterRule(_localctx, 92, RULE_not);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(476);
			match(T__42);
			setState(477);
			match(T__4);
			setState(478);
			expression();
			setState(479);
			match(T__5);
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

	public static class ListAccessContext extends ParserRuleContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public ListAccessContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listAccess; }
	}

	public final ListAccessContext listAccess() throws RecognitionException {
		ListAccessContext _localctx = new ListAccessContext(_ctx, getState());
		enterRule(_localctx, 94, RULE_listAccess);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(481);
			match(T__43);
			setState(482);
			match(T__4);
			setState(483);
			expression();
			setState(484);
			match(T__12);
			setState(485);
			expression();
			setState(488);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__12) {
				{
				setState(486);
				match(T__12);
				setState(487);
				expression();
				}
			}

			setState(490);
			match(T__5);
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

	public static class ListAssignContext extends ParserRuleContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public ListAssignContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listAssign; }
	}

	public final ListAssignContext listAssign() throws RecognitionException {
		ListAssignContext _localctx = new ListAssignContext(_ctx, getState());
		enterRule(_localctx, 96, RULE_listAssign);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(492);
			match(T__44);
			setState(493);
			match(T__4);
			setState(494);
			expression();
			setState(495);
			match(T__12);
			setState(496);
			expression();
			setState(497);
			match(T__12);
			setState(498);
			expression();
			setState(501);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__12) {
				{
				setState(499);
				match(T__12);
				setState(500);
				expression();
				}
			}

			setState(503);
			match(T__5);
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

	public static class ExpressionContext extends ParserRuleContext {
		public List<ExpContext> exp() {
			return getRuleContexts(ExpContext.class);
		}
		public ExpContext exp(int i) {
			return getRuleContext(ExpContext.class,i);
		}
		public LogicOperatorContext logicOperator() {
			return getRuleContext(LogicOperatorContext.class,0);
		}
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		ExpressionContext _localctx = new ExpressionContext(_ctx, getState());
		enterRule(_localctx, 98, RULE_expression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(505);
			exp();
			setState(509);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__50) | (1L << T__51) | (1L << T__52) | (1L << T__53) | (1L << T__54) | (1L << T__55) | (1L << T__56) | (1L << T__57))) != 0)) {
				{
				setState(506);
				logicOperator();
				setState(507);
				exp();
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

	public static class ExpContext extends ParserRuleContext {
		public List<TermContext> term() {
			return getRuleContexts(TermContext.class);
		}
		public TermContext term(int i) {
			return getRuleContext(TermContext.class,i);
		}
		public List<OpSecContext> opSec() {
			return getRuleContexts(OpSecContext.class);
		}
		public OpSecContext opSec(int i) {
			return getRuleContext(OpSecContext.class,i);
		}
		public ExpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_exp; }
	}

	public final ExpContext exp() throws RecognitionException {
		ExpContext _localctx = new ExpContext(_ctx, getState());
		enterRule(_localctx, 100, RULE_exp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(511);
			term();
			setState(517);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__45 || _la==T__46) {
				{
				{
				setState(512);
				opSec();
				setState(513);
				term();
				}
				}
				setState(519);
				_errHandler.sync(this);
				_la = _input.LA(1);
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

	public static class TermContext extends ParserRuleContext {
		public List<FactorContext> factor() {
			return getRuleContexts(FactorContext.class);
		}
		public FactorContext factor(int i) {
			return getRuleContext(FactorContext.class,i);
		}
		public List<OpFirstContext> opFirst() {
			return getRuleContexts(OpFirstContext.class);
		}
		public OpFirstContext opFirst(int i) {
			return getRuleContext(OpFirstContext.class,i);
		}
		public TermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_term; }
	}

	public final TermContext term() throws RecognitionException {
		TermContext _localctx = new TermContext(_ctx, getState());
		enterRule(_localctx, 102, RULE_term);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(520);
			factor();
			setState(526);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__47) | (1L << T__48) | (1L << T__49))) != 0)) {
				{
				{
				setState(521);
				opFirst();
				setState(522);
				factor();
				}
				}
				setState(528);
				_errHandler.sync(this);
				_la = _input.LA(1);
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

	public static class FactorContext extends ParserRuleContext {
		public NestedExpressionContext nestedExpression() {
			return getRuleContext(NestedExpressionContext.class,0);
		}
		public SpecialFunctionContext specialFunction() {
			return getRuleContext(SpecialFunctionContext.class,0);
		}
		public FunctionCallContext functionCall() {
			return getRuleContext(FunctionCallContext.class,0);
		}
		public VarConsContext varCons() {
			return getRuleContext(VarConsContext.class,0);
		}
		public UnaryMinusContext unaryMinus() {
			return getRuleContext(UnaryMinusContext.class,0);
		}
		public FactorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_factor; }
	}

	public final FactorContext factor() throws RecognitionException {
		FactorContext _localctx = new FactorContext(_ctx, getState());
		enterRule(_localctx, 104, RULE_factor);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(530);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__45) {
				{
				setState(529);
				unaryMinus();
				}
			}

			setState(536);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,30,_ctx) ) {
			case 1:
				{
				setState(532);
				nestedExpression();
				}
				break;
			case 2:
				{
				setState(533);
				specialFunction();
				}
				break;
			case 3:
				{
				setState(534);
				functionCall();
				}
				break;
			case 4:
				{
				setState(535);
				varCons();
				}
				break;
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

	public static class UnaryMinusContext extends ParserRuleContext {
		public UnaryMinusContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_unaryMinus; }
	}

	public final UnaryMinusContext unaryMinus() throws RecognitionException {
		UnaryMinusContext _localctx = new UnaryMinusContext(_ctx, getState());
		enterRule(_localctx, 106, RULE_unaryMinus);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(538);
			match(T__45);
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

	public static class NestedExpressionContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public NestedExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_nestedExpression; }
	}

	public final NestedExpressionContext nestedExpression() throws RecognitionException {
		NestedExpressionContext _localctx = new NestedExpressionContext(_ctx, getState());
		enterRule(_localctx, 108, RULE_nestedExpression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(540);
			match(T__4);
			setState(541);
			expression();
			setState(542);
			match(T__5);
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

	public static class FunctionCallContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(BistatParser.ID, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public FunctionCallContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionCall; }
	}

	public final FunctionCallContext functionCall() throws RecognitionException {
		FunctionCallContext _localctx = new FunctionCallContext(_ctx, getState());
		enterRule(_localctx, 110, RULE_functionCall);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(544);
			match(ID);
			setState(545);
			match(T__4);
			setState(554);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (((((_la - 5)) & ~0x3f) == 0 && ((1L << (_la - 5)) & ((1L << (T__4 - 5)) | (1L << (T__21 - 5)) | (1L << (T__22 - 5)) | (1L << (T__23 - 5)) | (1L << (T__24 - 5)) | (1L << (T__25 - 5)) | (1L << (T__26 - 5)) | (1L << (T__27 - 5)) | (1L << (T__28 - 5)) | (1L << (T__29 - 5)) | (1L << (T__30 - 5)) | (1L << (T__31 - 5)) | (1L << (T__32 - 5)) | (1L << (T__33 - 5)) | (1L << (T__34 - 5)) | (1L << (T__35 - 5)) | (1L << (T__36 - 5)) | (1L << (T__37 - 5)) | (1L << (T__38 - 5)) | (1L << (T__39 - 5)) | (1L << (T__40 - 5)) | (1L << (T__41 - 5)) | (1L << (T__42 - 5)) | (1L << (T__43 - 5)) | (1L << (T__44 - 5)) | (1L << (T__45 - 5)) | (1L << (BOOL_CONS - 5)) | (1L << (ID - 5)) | (1L << (INT_CONS - 5)) | (1L << (STRING_CONS - 5)) | (1L << (FLOAT_CONS - 5)))) != 0)) {
				{
				setState(546);
				expression();
				setState(551);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__12) {
					{
					{
					setState(547);
					match(T__12);
					setState(548);
					expression();
					}
					}
					setState(553);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(556);
			match(T__5);
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

	public static class VarConsContext extends ParserRuleContext {
		public TerminalNode STRING_CONS() { return getToken(BistatParser.STRING_CONS, 0); }
		public TerminalNode FLOAT_CONS() { return getToken(BistatParser.FLOAT_CONS, 0); }
		public TerminalNode INT_CONS() { return getToken(BistatParser.INT_CONS, 0); }
		public TerminalNode BOOL_CONS() { return getToken(BistatParser.BOOL_CONS, 0); }
		public TerminalNode ID() { return getToken(BistatParser.ID, 0); }
		public VarConsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varCons; }
	}

	public final VarConsContext varCons() throws RecognitionException {
		VarConsContext _localctx = new VarConsContext(_ctx, getState());
		enterRule(_localctx, 112, RULE_varCons);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(558);
			_la = _input.LA(1);
			if ( !(((((_la - 62)) & ~0x3f) == 0 && ((1L << (_la - 62)) & ((1L << (BOOL_CONS - 62)) | (1L << (ID - 62)) | (1L << (INT_CONS - 62)) | (1L << (STRING_CONS - 62)) | (1L << (FLOAT_CONS - 62)))) != 0)) ) {
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

	public static class OpSecContext extends ParserRuleContext {
		public OpSecContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_opSec; }
	}

	public final OpSecContext opSec() throws RecognitionException {
		OpSecContext _localctx = new OpSecContext(_ctx, getState());
		enterRule(_localctx, 114, RULE_opSec);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(560);
			_la = _input.LA(1);
			if ( !(_la==T__45 || _la==T__46) ) {
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

	public static class OpFirstContext extends ParserRuleContext {
		public OpFirstContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_opFirst; }
	}

	public final OpFirstContext opFirst() throws RecognitionException {
		OpFirstContext _localctx = new OpFirstContext(_ctx, getState());
		enterRule(_localctx, 116, RULE_opFirst);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(562);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__47) | (1L << T__48) | (1L << T__49))) != 0)) ) {
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

	public static class LogicOperatorContext extends ParserRuleContext {
		public LogicOperatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_logicOperator; }
	}

	public final LogicOperatorContext logicOperator() throws RecognitionException {
		LogicOperatorContext _localctx = new LogicOperatorContext(_ctx, getState());
		enterRule(_localctx, 118, RULE_logicOperator);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(564);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__50) | (1L << T__51) | (1L << T__52) | (1L << T__53) | (1L << T__54) | (1L << T__55) | (1L << T__56) | (1L << T__57))) != 0)) ) {
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3E\u0239\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\3\2\3\2\3\2\3\2\7\2\177\n\2\f\2\16\2\u0082\13\2\3\2\7\2\u0085\n\2\f\2"+
		"\16\2\u0088\13\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\4\3\4\5\4\u0094\n\4"+
		"\3\5\3\5\3\5\3\5\7\5\u009a\n\5\f\5\16\5\u009d\13\5\3\5\3\5\3\5\3\5\7\5"+
		"\u00a3\n\5\f\5\16\5\u00a6\13\5\3\5\3\5\6\5\u00aa\n\5\r\5\16\5\u00ab\3"+
		"\5\3\5\3\6\3\6\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\6\t\u00be"+
		"\n\t\r\t\16\t\u00bf\3\t\3\t\3\n\3\n\3\n\3\n\5\n\u00c8\n\n\3\n\3\n\3\n"+
		"\3\n\3\n\3\n\5\n\u00d0\n\n\3\13\3\13\3\13\3\13\3\13\5\13\u00d7\n\13\3"+
		"\f\3\f\3\f\3\f\7\f\u00dd\n\f\f\f\16\f\u00e0\13\f\3\f\3\f\3\r\3\r\3\r\3"+
		"\r\7\r\u00e8\n\r\f\r\16\r\u00eb\13\r\3\r\3\r\3\16\3\16\6\16\u00f1\n\16"+
		"\r\16\16\16\u00f2\3\16\3\16\3\17\3\17\3\17\6\17\u00fa\n\17\r\17\16\17"+
		"\u00fb\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\21\3\21\3\22\3\22"+
		"\3\22\3\22\3\22\3\22\6\22\u010f\n\22\r\22\16\22\u0110\3\22\3\22\3\23\3"+
		"\23\3\24\3\24\7\24\u0119\n\24\f\24\16\24\u011c\13\24\3\24\5\24\u011f\n"+
		"\24\3\25\3\25\3\25\3\25\3\25\3\25\6\25\u0127\n\25\r\25\16\25\u0128\3\25"+
		"\3\25\3\26\3\26\3\26\3\26\3\26\3\26\3\26\6\26\u0134\n\26\r\26\16\26\u0135"+
		"\3\26\3\26\3\27\3\27\3\30\3\30\3\30\6\30\u013f\n\30\r\30\16\30\u0140\3"+
		"\30\3\30\3\31\3\31\3\31\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3"+
		"\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3"+
		"\32\5\32\u0160\n\32\3\33\3\33\3\33\3\33\3\33\7\33\u0167\n\33\f\33\16\33"+
		"\u016a\13\33\3\33\3\33\3\34\3\34\3\34\3\34\3\34\7\34\u0173\n\34\f\34\16"+
		"\34\u0176\13\34\3\34\3\34\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\36\3\36"+
		"\3\36\3\36\3\36\3\37\3\37\3\37\3\37\3\37\3 \3 \3 \3 \3 \5 \u0190\n \3"+
		" \3 \3!\3!\3!\3!\3!\3\"\3\"\3\"\3\"\3\"\3#\3#\3#\3#\3#\3$\3$\3$\3$\3$"+
		"\3%\3%\3%\3%\3%\3&\3&\3&\3&\3&\3\'\3\'\3\'\3\'\3\'\3(\3(\3(\3(\3(\3)\3"+
		")\3)\3)\3)\3*\3*\3*\3*\3*\3+\3+\3+\3+\3+\3,\3,\3,\3,\3,\3-\3-\3-\3-\3"+
		"-\3.\3.\3.\3.\3.\3/\3/\3/\3/\3/\3\60\3\60\3\60\3\60\3\60\3\61\3\61\3\61"+
		"\3\61\3\61\3\61\3\61\5\61\u01eb\n\61\3\61\3\61\3\62\3\62\3\62\3\62\3\62"+
		"\3\62\3\62\3\62\3\62\5\62\u01f8\n\62\3\62\3\62\3\63\3\63\3\63\3\63\5\63"+
		"\u0200\n\63\3\64\3\64\3\64\3\64\7\64\u0206\n\64\f\64\16\64\u0209\13\64"+
		"\3\65\3\65\3\65\3\65\7\65\u020f\n\65\f\65\16\65\u0212\13\65\3\66\5\66"+
		"\u0215\n\66\3\66\3\66\3\66\3\66\5\66\u021b\n\66\3\67\3\67\38\38\38\38"+
		"\39\39\39\39\39\79\u0228\n9\f9\169\u022b\139\59\u022d\n9\39\39\3:\3:\3"+
		";\3;\3<\3<\3=\3=\3=\2\2>\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&("+
		"*,.\60\62\64\668:<>@BDFHJLNPRTVXZ\\^`bdfhjlnprtvx\2\7\3\2\21\21\4\2@B"+
		"DE\3\2\60\61\3\2\62\64\3\2\65<\2\u023b\2z\3\2\2\2\4\u008c\3\2\2\2\6\u0091"+
		"\3\2\2\2\b\u0095\3\2\2\2\n\u00af\3\2\2\2\f\u00b1\3\2\2\2\16\u00b3\3\2"+
		"\2\2\20\u00b8\3\2\2\2\22\u00cf\3\2\2\2\24\u00d1\3\2\2\2\26\u00d8\3\2\2"+
		"\2\30\u00e3\3\2\2\2\32\u00ee\3\2\2\2\34\u00f6\3\2\2\2\36\u00ff\3\2\2\2"+
		" \u0106\3\2\2\2\"\u0108\3\2\2\2$\u0114\3\2\2\2&\u0116\3\2\2\2(\u0120\3"+
		"\2\2\2*\u012c\3\2\2\2,\u0139\3\2\2\2.\u013b\3\2\2\2\60\u0144\3\2\2\2\62"+
		"\u015f\3\2\2\2\64\u0161\3\2\2\2\66\u016d\3\2\2\28\u0179\3\2\2\2:\u0180"+
		"\3\2\2\2<\u0185\3\2\2\2>\u018a\3\2\2\2@\u0193\3\2\2\2B\u0198\3\2\2\2D"+
		"\u019d\3\2\2\2F\u01a2\3\2\2\2H\u01a7\3\2\2\2J\u01ac\3\2\2\2L\u01b1\3\2"+
		"\2\2N\u01b6\3\2\2\2P\u01bb\3\2\2\2R\u01c0\3\2\2\2T\u01c5\3\2\2\2V\u01ca"+
		"\3\2\2\2X\u01cf\3\2\2\2Z\u01d4\3\2\2\2\\\u01d9\3\2\2\2^\u01de\3\2\2\2"+
		"`\u01e3\3\2\2\2b\u01ee\3\2\2\2d\u01fb\3\2\2\2f\u0201\3\2\2\2h\u020a\3"+
		"\2\2\2j\u0214\3\2\2\2l\u021c\3\2\2\2n\u021e\3\2\2\2p\u0222\3\2\2\2r\u0230"+
		"\3\2\2\2t\u0232\3\2\2\2v\u0234\3\2\2\2x\u0236\3\2\2\2z{\7\3\2\2{|\7A\2"+
		"\2|\u0080\7\4\2\2}\177\5\4\3\2~}\3\2\2\2\177\u0082\3\2\2\2\u0080~\3\2"+
		"\2\2\u0080\u0081\3\2\2\2\u0081\u0086\3\2\2\2\u0082\u0080\3\2\2\2\u0083"+
		"\u0085\5\b\5\2\u0084\u0083\3\2\2\2\u0085\u0088\3\2\2\2\u0086\u0084\3\2"+
		"\2\2\u0086\u0087\3\2\2\2\u0087\u0089\3\2\2\2\u0088\u0086\3\2\2\2\u0089"+
		"\u008a\5\20\t\2\u008a\u008b\7\2\2\3\u008b\3\3\2\2\2\u008c\u008d\7\5\2"+
		"\2\u008d\u008e\5\6\4\2\u008e\u008f\7A\2\2\u008f\u0090\7\4\2\2\u0090\5"+
		"\3\2\2\2\u0091\u0093\7?\2\2\u0092\u0094\7>\2\2\u0093\u0092\3\2\2\2\u0093"+
		"\u0094\3\2\2\2\u0094\7\3\2\2\2\u0095\u0096\7\6\2\2\u0096\u0097\7A\2\2"+
		"\u0097\u009b\7\7\2\2\u0098\u009a\5\16\b\2\u0099\u0098\3\2\2\2\u009a\u009d"+
		"\3\2\2\2\u009b\u0099\3\2\2\2\u009b\u009c\3\2\2\2\u009c\u009e\3\2\2\2\u009d"+
		"\u009b\3\2\2\2\u009e\u009f\7\b\2\2\u009f\u00a0\7\t\2\2\u00a0\u00a4\7?"+
		"\2\2\u00a1\u00a3\5\4\3\2\u00a2\u00a1\3\2\2\2\u00a3\u00a6\3\2\2\2\u00a4"+
		"\u00a2\3\2\2\2\u00a4\u00a5\3\2\2\2\u00a5\u00a7\3\2\2\2\u00a6\u00a4\3\2"+
		"\2\2\u00a7\u00a9\5\n\6\2\u00a8\u00aa\5\22\n\2\u00a9\u00a8\3\2\2\2\u00aa"+
		"\u00ab\3\2\2\2\u00ab\u00a9\3\2\2\2\u00ab\u00ac\3\2\2\2\u00ac\u00ad\3\2"+
		"\2\2\u00ad\u00ae\5\f\7\2\u00ae\t\3\2\2\2\u00af\u00b0\7\n\2\2\u00b0\13"+
		"\3\2\2\2\u00b1\u00b2\7\13\2\2\u00b2\r\3\2\2\2\u00b3\u00b4\7\5\2\2\u00b4"+
		"\u00b5\7?\2\2\u00b5\u00b6\7A\2\2\u00b6\u00b7\7\4\2\2\u00b7\17\3\2\2\2"+
		"\u00b8\u00b9\7\f\2\2\u00b9\u00ba\7\7\2\2\u00ba\u00bb\7\b\2\2\u00bb\u00bd"+
		"\7\n\2\2\u00bc\u00be\5\22\n\2\u00bd\u00bc\3\2\2\2\u00be\u00bf\3\2\2\2"+
		"\u00bf\u00bd\3\2\2\2\u00bf\u00c0\3\2\2\2\u00c0\u00c1\3\2\2\2\u00c1\u00c2"+
		"\7\13\2\2\u00c2\21\3\2\2\2\u00c3\u00c8\5\24\13\2\u00c4\u00c8\5\62\32\2"+
		"\u00c5\u00c8\5p9\2\u00c6\u00c8\5\60\31\2\u00c7\u00c3\3\2\2\2\u00c7\u00c4"+
		"\3\2\2\2\u00c7\u00c5\3\2\2\2\u00c7\u00c6\3\2\2\2\u00c8\u00c9\3\2\2\2\u00c9"+
		"\u00ca\7\4\2\2\u00ca\u00d0\3\2\2\2\u00cb\u00d0\5&\24\2\u00cc\u00d0\5\""+
		"\22\2\u00cd\u00d0\5\34\17\2\u00ce\u00d0\5\32\16\2\u00cf\u00c7\3\2\2\2"+
		"\u00cf\u00cb\3\2\2\2\u00cf\u00cc\3\2\2\2\u00cf\u00cd\3\2\2\2\u00cf\u00ce"+
		"\3\2\2\2\u00d0\23\3\2\2\2\u00d1\u00d2\7A\2\2\u00d2\u00d6\7\r\2\2\u00d3"+
		"\u00d7\5d\63\2\u00d4\u00d7\5\30\r\2\u00d5\u00d7\5\26\f\2\u00d6\u00d3\3"+
		"\2\2\2\u00d6\u00d4\3\2\2\2\u00d6\u00d5\3\2\2\2\u00d7\25\3\2\2\2\u00d8"+
		"\u00d9\7\16\2\2\u00d9\u00de\5\30\r\2\u00da\u00db\7\17\2\2\u00db\u00dd"+
		"\5\30\r\2\u00dc\u00da\3\2\2\2\u00dd\u00e0\3\2\2\2\u00de\u00dc\3\2\2\2"+
		"\u00de\u00df\3\2\2\2\u00df\u00e1\3\2\2\2\u00e0\u00de\3\2\2\2\u00e1\u00e2"+
		"\7\20\2\2\u00e2\27\3\2\2\2\u00e3\u00e4\7\16\2\2\u00e4\u00e9\5d\63\2\u00e5"+
		"\u00e6\7\17\2\2\u00e6\u00e8\5d\63\2\u00e7\u00e5\3\2\2\2\u00e8\u00eb\3"+
		"\2\2\2\u00e9\u00e7\3\2\2\2\u00e9\u00ea\3\2\2\2\u00ea\u00ec\3\2\2\2\u00eb"+
		"\u00e9\3\2\2\2\u00ec\u00ed\7\20\2\2\u00ed\31\3\2\2\2\u00ee\u00f0\7\21"+
		"\2\2\u00ef\u00f1\n\2\2\2\u00f0\u00ef\3\2\2\2\u00f1\u00f2\3\2\2\2\u00f2"+
		"\u00f0\3\2\2\2\u00f2\u00f3\3\2\2\2\u00f3\u00f4\3\2\2\2\u00f4\u00f5\7\21"+
		"\2\2\u00f5\33\3\2\2\2\u00f6\u00f7\5\36\20\2\u00f7\u00f9\7\n\2\2\u00f8"+
		"\u00fa\5\22\n\2\u00f9\u00f8\3\2\2\2\u00fa\u00fb\3\2\2\2\u00fb\u00f9\3"+
		"\2\2\2\u00fb\u00fc\3\2\2\2\u00fc\u00fd\3\2\2\2\u00fd\u00fe\7\13\2\2\u00fe"+
		"\35\3\2\2\2\u00ff\u0100\7\22\2\2\u0100\u0101\7\7\2\2\u0101\u0102\7A\2"+
		"\2\u0102\u0103\7\23\2\2\u0103\u0104\5d\63\2\u0104\u0105\7\b\2\2\u0105"+
		"\37\3\2\2\2\u0106\u0107\7\b\2\2\u0107!\3\2\2\2\u0108\u0109\7\24\2\2\u0109"+
		"\u010a\7\7\2\2\u010a\u010b\5d\63\2\u010b\u010c\5$\23\2\u010c\u010e\7\n"+
		"\2\2\u010d\u010f\5\22\n\2\u010e\u010d\3\2\2\2\u010f\u0110\3\2\2\2\u0110"+
		"\u010e\3\2\2\2\u0110\u0111\3\2\2\2\u0111\u0112\3\2\2\2\u0112\u0113\7\13"+
		"\2\2\u0113#\3\2\2\2\u0114\u0115\7\b\2\2\u0115%\3\2\2\2\u0116\u011a\5("+
		"\25\2\u0117\u0119\5*\26\2\u0118\u0117\3\2\2\2\u0119\u011c\3\2\2\2\u011a"+
		"\u0118\3\2\2\2\u011a\u011b\3\2\2\2\u011b\u011e\3\2\2\2\u011c\u011a\3\2"+
		"\2\2\u011d\u011f\5.\30\2\u011e\u011d\3\2\2\2\u011e\u011f\3\2\2\2\u011f"+
		"\'\3\2\2\2\u0120\u0121\7\25\2\2\u0121\u0122\7\7\2\2\u0122\u0123\5d\63"+
		"\2\u0123\u0124\5,\27\2\u0124\u0126\7\n\2\2\u0125\u0127\5\22\n\2\u0126"+
		"\u0125\3\2\2\2\u0127\u0128\3\2\2\2\u0128\u0126\3\2\2\2\u0128\u0129\3\2"+
		"\2\2\u0129\u012a\3\2\2\2\u012a\u012b\7\13\2\2\u012b)\3\2\2\2\u012c\u012d"+
		"\7\26\2\2\u012d\u012e\7\25\2\2\u012e\u012f\7\7\2\2\u012f\u0130\5d\63\2"+
		"\u0130\u0131\5,\27\2\u0131\u0133\7\n\2\2\u0132\u0134\5\22\n\2\u0133\u0132"+
		"\3\2\2\2\u0134\u0135\3\2\2\2\u0135\u0133\3\2\2\2\u0135\u0136\3\2\2\2\u0136"+
		"\u0137\3\2\2\2\u0137\u0138\7\13\2\2\u0138+\3\2\2\2\u0139\u013a\7\b\2\2"+
		"\u013a-\3\2\2\2\u013b\u013c\7\26\2\2\u013c\u013e\7\n\2\2\u013d\u013f\5"+
		"\22\n\2\u013e\u013d\3\2\2\2\u013f\u0140\3\2\2\2\u0140\u013e\3\2\2\2\u0140"+
		"\u0141\3\2\2\2\u0141\u0142\3\2\2\2\u0142\u0143\7\13\2\2\u0143/\3\2\2\2"+
		"\u0144\u0145\7\27\2\2\u0145\u0146\5d\63\2\u0146\61\3\2\2\2\u0147\u0160"+
		"\5\64\33\2\u0148\u0160\5\66\34\2\u0149\u0160\58\35\2\u014a\u0160\5`\61"+
		"\2\u014b\u0160\5b\62\2\u014c\u0160\5:\36\2\u014d\u0160\5<\37\2\u014e\u0160"+
		"\5> \2\u014f\u0160\5@!\2\u0150\u0160\5B\"\2\u0151\u0160\5D#\2\u0152\u0160"+
		"\5F$\2\u0153\u0160\5H%\2\u0154\u0160\5J&\2\u0155\u0160\5L\'\2\u0156\u0160"+
		"\5N(\2\u0157\u0160\5R*\2\u0158\u0160\5P)\2\u0159\u0160\5T+\2\u015a\u0160"+
		"\5V,\2\u015b\u0160\5X-\2\u015c\u0160\5Z.\2\u015d\u0160\5\\/\2\u015e\u0160"+
		"\5^\60\2\u015f\u0147\3\2\2\2\u015f\u0148\3\2\2\2\u015f\u0149\3\2\2\2\u015f"+
		"\u014a\3\2\2\2\u015f\u014b\3\2\2\2\u015f\u014c\3\2\2\2\u015f\u014d\3\2"+
		"\2\2\u015f\u014e\3\2\2\2\u015f\u014f\3\2\2\2\u015f\u0150\3\2\2\2\u015f"+
		"\u0151\3\2\2\2\u015f\u0152\3\2\2\2\u015f\u0153\3\2\2\2\u015f\u0154\3\2"+
		"\2\2\u015f\u0155\3\2\2\2\u015f\u0156\3\2\2\2\u015f\u0157\3\2\2\2\u015f"+
		"\u0158\3\2\2\2\u015f\u0159\3\2\2\2\u015f\u015a\3\2\2\2\u015f\u015b\3\2"+
		"\2\2\u015f\u015c\3\2\2\2\u015f\u015d\3\2\2\2\u015f\u015e\3\2\2\2\u0160"+
		"\63\3\2\2\2\u0161\u0162\7\30\2\2\u0162\u0163\7\7\2\2\u0163\u0168\7A\2"+
		"\2\u0164\u0165\7\17\2\2\u0165\u0167\7A\2\2\u0166\u0164\3\2\2\2\u0167\u016a"+
		"\3\2\2\2\u0168\u0166\3\2\2\2\u0168\u0169\3\2\2\2\u0169\u016b\3\2\2\2\u016a"+
		"\u0168\3\2\2\2\u016b\u016c\7\b\2\2\u016c\65\3\2\2\2\u016d\u016e\7\31\2"+
		"\2\u016e\u016f\7\7\2\2\u016f\u0174\5d\63\2\u0170\u0171\7\17\2\2\u0171"+
		"\u0173\5d\63\2\u0172\u0170\3\2\2\2\u0173\u0176\3\2\2\2\u0174\u0172\3\2"+
		"\2\2\u0174\u0175\3\2\2\2\u0175\u0177\3\2\2\2\u0176\u0174\3\2\2\2\u0177"+
		"\u0178\7\b\2\2\u0178\67\3\2\2\2\u0179\u017a\7\32\2\2\u017a\u017b\7\7\2"+
		"\2\u017b\u017c\5d\63\2\u017c\u017d\7\17\2\2\u017d\u017e\5d\63\2\u017e"+
		"\u017f\7\b\2\2\u017f9\3\2\2\2\u0180\u0181\7\33\2\2\u0181\u0182\7\7\2\2"+
		"\u0182\u0183\5d\63\2\u0183\u0184\7\b\2\2\u0184;\3\2\2\2\u0185\u0186\7"+
		"\34\2\2\u0186\u0187\7\7\2\2\u0187\u0188\5d\63\2\u0188\u0189\7\b\2\2\u0189"+
		"=\3\2\2\2\u018a\u018b\7\35\2\2\u018b\u018c\7\7\2\2\u018c\u018f\5d\63\2"+
		"\u018d\u018e\7\17\2\2\u018e\u0190\5d\63\2\u018f\u018d\3\2\2\2\u018f\u0190"+
		"\3\2\2\2\u0190\u0191\3\2\2\2\u0191\u0192\7\b\2\2\u0192?\3\2\2\2\u0193"+
		"\u0194\7\36\2\2\u0194\u0195\7\7\2\2\u0195\u0196\5d\63\2\u0196\u0197\7"+
		"\b\2\2\u0197A\3\2\2\2\u0198\u0199\7\37\2\2\u0199\u019a\7\7\2\2\u019a\u019b"+
		"\5d\63\2\u019b\u019c\7\b\2\2\u019cC\3\2\2\2\u019d\u019e\7 \2\2\u019e\u019f"+
		"\7\7\2\2\u019f\u01a0\5d\63\2\u01a0\u01a1\7\b\2\2\u01a1E\3\2\2\2\u01a2"+
		"\u01a3\7!\2\2\u01a3\u01a4\7\7\2\2\u01a4\u01a5\5d\63\2\u01a5\u01a6\7\b"+
		"\2\2\u01a6G\3\2\2\2\u01a7\u01a8\7\"\2\2\u01a8\u01a9\7\7\2\2\u01a9\u01aa"+
		"\5d\63\2\u01aa\u01ab\7\b\2\2\u01abI\3\2\2\2\u01ac\u01ad\7#\2\2\u01ad\u01ae"+
		"\7\7\2\2\u01ae\u01af\5d\63\2\u01af\u01b0\7\b\2\2\u01b0K\3\2\2\2\u01b1"+
		"\u01b2\7$\2\2\u01b2\u01b3\7\7\2\2\u01b3\u01b4\5d\63\2\u01b4\u01b5\7\b"+
		"\2\2\u01b5M\3\2\2\2\u01b6\u01b7\7%\2\2\u01b7\u01b8\7\7\2\2\u01b8\u01b9"+
		"\5d\63\2\u01b9\u01ba\7\b\2\2\u01baO\3\2\2\2\u01bb\u01bc\7&\2\2\u01bc\u01bd"+
		"\7\7\2\2\u01bd\u01be\5d\63\2\u01be\u01bf\7\b\2\2\u01bfQ\3\2\2\2\u01c0"+
		"\u01c1\7\'\2\2\u01c1\u01c2\7\7\2\2\u01c2\u01c3\5d\63\2\u01c3\u01c4\7\b"+
		"\2\2\u01c4S\3\2\2\2\u01c5\u01c6\7(\2\2\u01c6\u01c7\7\7\2\2\u01c7\u01c8"+
		"\5d\63\2\u01c8\u01c9\7\b\2\2\u01c9U\3\2\2\2\u01ca\u01cb\7)\2\2\u01cb\u01cc"+
		"\7\7\2\2\u01cc\u01cd\5d\63\2\u01cd\u01ce\7\b\2\2\u01ceW\3\2\2\2\u01cf"+
		"\u01d0\7*\2\2\u01d0\u01d1\7\7\2\2\u01d1\u01d2\5d\63\2\u01d2\u01d3\7\b"+
		"\2\2\u01d3Y\3\2\2\2\u01d4\u01d5\7+\2\2\u01d5\u01d6\7\7\2\2\u01d6\u01d7"+
		"\5d\63\2\u01d7\u01d8\7\b\2\2\u01d8[\3\2\2\2\u01d9\u01da\7,\2\2\u01da\u01db"+
		"\7\7\2\2\u01db\u01dc\5d\63\2\u01dc\u01dd\7\b\2\2\u01dd]\3\2\2\2\u01de"+
		"\u01df\7-\2\2\u01df\u01e0\7\7\2\2\u01e0\u01e1\5d\63\2\u01e1\u01e2\7\b"+
		"\2\2\u01e2_\3\2\2\2\u01e3\u01e4\7.\2\2\u01e4\u01e5\7\7\2\2\u01e5\u01e6"+
		"\5d\63\2\u01e6\u01e7\7\17\2\2\u01e7\u01ea\5d\63\2\u01e8\u01e9\7\17\2\2"+
		"\u01e9\u01eb\5d\63\2\u01ea\u01e8\3\2\2\2\u01ea\u01eb\3\2\2\2\u01eb\u01ec"+
		"\3\2\2\2\u01ec\u01ed\7\b\2\2\u01eda\3\2\2\2\u01ee\u01ef\7/\2\2\u01ef\u01f0"+
		"\7\7\2\2\u01f0\u01f1\5d\63\2\u01f1\u01f2\7\17\2\2\u01f2\u01f3\5d\63\2"+
		"\u01f3\u01f4\7\17\2\2\u01f4\u01f7\5d\63\2\u01f5\u01f6\7\17\2\2\u01f6\u01f8"+
		"\5d\63\2\u01f7\u01f5\3\2\2\2\u01f7\u01f8\3\2\2\2\u01f8\u01f9\3\2\2\2\u01f9"+
		"\u01fa\7\b\2\2\u01fac\3\2\2\2\u01fb\u01ff\5f\64\2\u01fc\u01fd\5x=\2\u01fd"+
		"\u01fe\5f\64\2\u01fe\u0200\3\2\2\2\u01ff\u01fc\3\2\2\2\u01ff\u0200\3\2"+
		"\2\2\u0200e\3\2\2\2\u0201\u0207\5h\65\2\u0202\u0203\5t;\2\u0203\u0204"+
		"\5h\65\2\u0204\u0206\3\2\2\2\u0205\u0202\3\2\2\2\u0206\u0209\3\2\2\2\u0207"+
		"\u0205\3\2\2\2\u0207\u0208\3\2\2\2\u0208g\3\2\2\2\u0209\u0207\3\2\2\2"+
		"\u020a\u0210\5j\66\2\u020b\u020c\5v<\2\u020c\u020d\5j\66\2\u020d\u020f"+
		"\3\2\2\2\u020e\u020b\3\2\2\2\u020f\u0212\3\2\2\2\u0210\u020e\3\2\2\2\u0210"+
		"\u0211\3\2\2\2\u0211i\3\2\2\2\u0212\u0210\3\2\2\2\u0213\u0215\5l\67\2"+
		"\u0214\u0213\3\2\2\2\u0214\u0215\3\2\2\2\u0215\u021a\3\2\2\2\u0216\u021b"+
		"\5n8\2\u0217\u021b\5\62\32\2\u0218\u021b\5p9\2\u0219\u021b\5r:\2\u021a"+
		"\u0216\3\2\2\2\u021a\u0217\3\2\2\2\u021a\u0218\3\2\2\2\u021a\u0219\3\2"+
		"\2\2\u021bk\3\2\2\2\u021c\u021d\7\60\2\2\u021dm\3\2\2\2\u021e\u021f\7"+
		"\7\2\2\u021f\u0220\5d\63\2\u0220\u0221\7\b\2\2\u0221o\3\2\2\2\u0222\u0223"+
		"\7A\2\2\u0223\u022c\7\7\2\2\u0224\u0229\5d\63\2\u0225\u0226\7\17\2\2\u0226"+
		"\u0228\5d\63\2\u0227\u0225\3\2\2\2\u0228\u022b\3\2\2\2\u0229\u0227\3\2"+
		"\2\2\u0229\u022a\3\2\2\2\u022a\u022d\3\2\2\2\u022b\u0229\3\2\2\2\u022c"+
		"\u0224\3\2\2\2\u022c\u022d\3\2\2\2\u022d\u022e\3\2\2\2\u022e\u022f\7\b"+
		"\2\2\u022fq\3\2\2\2\u0230\u0231\t\3\2\2\u0231s\3\2\2\2\u0232\u0233\t\4"+
		"\2\2\u0233u\3\2\2\2\u0234\u0235\t\5\2\2\u0235w\3\2\2\2\u0236\u0237\t\6"+
		"\2\2\u0237y\3\2\2\2#\u0080\u0086\u0093\u009b\u00a4\u00ab\u00bf\u00c7\u00cf"+
		"\u00d6\u00de\u00e9\u00f2\u00fb\u0110\u011a\u011e\u0128\u0135\u0140\u015f"+
		"\u0168\u0174\u018f\u01ea\u01f7\u01ff\u0207\u0210\u0214\u021a\u0229\u022c";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}