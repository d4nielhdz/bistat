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
		T__52=53, T__53=54, T__54=55, T__55=56, WS=57, CARDINALITY=58, TYPE_PRIMITIVE=59, 
		BOOL_CONS=60, ID=61, INT_CONS=62, NUMBER=63, STRING_CONS=64, FLOAT_CONS=65;
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
		RULE_abs = 45, RULE_not = 46, RULE_expression = 47, RULE_exp = 48, RULE_term = 49, 
		RULE_factor = 50, RULE_unaryMinus = 51, RULE_nestedExpression = 52, RULE_functionCall = 53, 
		RULE_indexing = 54, RULE_varCons = 55, RULE_opSec = 56, RULE_opFirst = 57, 
		RULE_logicOperator = 58;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "varDeclaration", "var_type", "funcDef", "funcBlockStart", 
			"funcBlockEnd", "paramDeclaration", "main", "stmt", "assignment", "matrixAssignment", 
			"listAssignment", "comment", "forLoop", "forHeader", "forExprEnd", "whileLoop", 
			"whileExprEnd", "conditional", "ifStmt", "elseIfStmt", "condExprEnd", 
			"elseStmt", "returnStmt", "specialFunction", "inputRead", "print", "listAdd", 
			"listPop", "length", "range", "plot", "sum", "min", "prod", "avg", "sMode", 
			"median", "sin", "tan", "cos", "sort", "sqrt", "floor", "ceil", "abs", 
			"not", "expression", "exp", "term", "factor", "unaryMinus", "nestedExpression", 
			"functionCall", "indexing", "varCons", "opSec", "opFirst", "logicOperator"
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
			"'sqrt'", "'floor'", "'ceil'", "'abs'", "'not'", "'-'", "'+'", "'/'", 
			"'*'", "'%'", "'<'", "'>'", "'<='", "'>='", "'=='", "'!='", "'&&'", "'||'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, "WS", "CARDINALITY", 
			"TYPE_PRIMITIVE", "BOOL_CONS", "ID", "INT_CONS", "NUMBER", "STRING_CONS", 
			"FLOAT_CONS"
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
			setState(118);
			match(T__0);
			setState(119);
			match(ID);
			setState(120);
			match(T__1);
			setState(124);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__2) {
				{
				{
				setState(121);
				varDeclaration();
				}
				}
				setState(126);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(130);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__3) {
				{
				{
				setState(127);
				funcDef();
				}
				}
				setState(132);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(133);
			main();
			setState(134);
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
			setState(136);
			match(T__2);
			setState(137);
			var_type();
			setState(138);
			match(ID);
			setState(139);
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
			setState(141);
			match(TYPE_PRIMITIVE);
			{
			setState(143);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==CARDINALITY) {
				{
				setState(142);
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
		public Var_typeContext var_type() {
			return getRuleContext(Var_typeContext.class,0);
		}
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
			setState(145);
			match(T__3);
			setState(146);
			match(ID);
			setState(147);
			match(T__4);
			setState(151);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__2) {
				{
				{
				setState(148);
				paramDeclaration();
				}
				}
				setState(153);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(154);
			match(T__5);
			setState(155);
			match(T__6);
			setState(156);
			var_type();
			setState(160);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__2) {
				{
				{
				setState(157);
				varDeclaration();
				}
				}
				setState(162);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(163);
			funcBlockStart();
			setState(165); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(164);
				stmt();
				}
				}
				setState(167); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__14) | (1L << T__15) | (1L << T__17) | (1L << T__18) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << T__42) | (1L << ID))) != 0) );
			setState(169);
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
			setState(171);
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
			setState(173);
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
		public Var_typeContext var_type() {
			return getRuleContext(Var_typeContext.class,0);
		}
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
			setState(175);
			match(T__2);
			setState(176);
			var_type();
			setState(177);
			match(ID);
			setState(178);
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
			setState(180);
			match(T__9);
			setState(181);
			match(T__4);
			setState(182);
			match(T__5);
			setState(183);
			match(T__7);
			setState(185); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(184);
				stmt();
				}
				}
				setState(187); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__14) | (1L << T__15) | (1L << T__17) | (1L << T__18) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << T__42) | (1L << ID))) != 0) );
			setState(189);
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
			setState(203);
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
			case ID:
				enterOuterAlt(_localctx, 1);
				{
				{
				setState(195);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
				case 1:
					{
					setState(191);
					assignment();
					}
					break;
				case 2:
					{
					setState(192);
					specialFunction();
					}
					break;
				case 3:
					{
					setState(193);
					functionCall();
					}
					break;
				case 4:
					{
					setState(194);
					returnStmt();
					}
					break;
				}
				setState(197);
				match(T__1);
				}
				}
				break;
			case T__18:
				enterOuterAlt(_localctx, 2);
				{
				setState(199);
				conditional();
				}
				break;
			case T__17:
				enterOuterAlt(_localctx, 3);
				{
				setState(200);
				whileLoop();
				}
				break;
			case T__15:
				enterOuterAlt(_localctx, 4);
				{
				setState(201);
				forLoop();
				}
				break;
			case T__14:
				enterOuterAlt(_localctx, 5);
				{
				setState(202);
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
		public IndexingContext indexing() {
			return getRuleContext(IndexingContext.class,0);
		}
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
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				{
				setState(205);
				indexing();
				}
				break;
			case 2:
				{
				setState(206);
				match(ID);
				}
				break;
			}
			setState(209);
			match(T__10);
			setState(213);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				{
				setState(210);
				expression();
				}
				break;
			case 2:
				{
				setState(211);
				listAssignment();
				}
				break;
			case 3:
				{
				setState(212);
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
			setState(215);
			match(T__11);
			setState(216);
			listAssignment();
			setState(221);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__12) {
				{
				{
				setState(217);
				match(T__12);
				setState(218);
				listAssignment();
				}
				}
				setState(223);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(224);
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
			setState(226);
			match(T__11);
			{
			setState(227);
			expression();
			}
			setState(232);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__12) {
				{
				{
				setState(228);
				match(T__12);
				{
				setState(229);
				expression();
				}
				}
				}
				setState(234);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(235);
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
			setState(237);
			match(T__14);
			setState(239); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(238);
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
				setState(241); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__1) | (1L << T__2) | (1L << T__3) | (1L << T__4) | (1L << T__5) | (1L << T__6) | (1L << T__7) | (1L << T__8) | (1L << T__9) | (1L << T__10) | (1L << T__11) | (1L << T__12) | (1L << T__13) | (1L << T__15) | (1L << T__16) | (1L << T__17) | (1L << T__18) | (1L << T__19) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << T__42) | (1L << T__43) | (1L << T__44) | (1L << T__45) | (1L << T__46) | (1L << T__47) | (1L << T__48) | (1L << T__49) | (1L << T__50) | (1L << T__51) | (1L << T__52) | (1L << T__53) | (1L << T__54) | (1L << T__55) | (1L << WS) | (1L << CARDINALITY) | (1L << TYPE_PRIMITIVE) | (1L << BOOL_CONS) | (1L << ID) | (1L << INT_CONS) | (1L << NUMBER))) != 0) || _la==STRING_CONS || _la==FLOAT_CONS );
			setState(243);
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
			setState(245);
			forHeader();
			setState(246);
			match(T__7);
			setState(248); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(247);
				stmt();
				}
				}
				setState(250); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__14) | (1L << T__15) | (1L << T__17) | (1L << T__18) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << T__42) | (1L << ID))) != 0) );
			setState(252);
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
			setState(254);
			match(T__15);
			setState(255);
			match(T__4);
			setState(256);
			match(ID);
			setState(257);
			match(T__16);
			setState(258);
			expression();
			setState(259);
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
			setState(261);
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
			setState(263);
			match(T__17);
			setState(264);
			match(T__4);
			setState(265);
			expression();
			setState(266);
			whileExprEnd();
			setState(267);
			match(T__7);
			setState(269); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(268);
				stmt();
				}
				}
				setState(271); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__14) | (1L << T__15) | (1L << T__17) | (1L << T__18) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << T__42) | (1L << ID))) != 0) );
			setState(273);
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
			setState(275);
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
			setState(277);
			ifStmt();
			setState(281);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,16,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(278);
					elseIfStmt();
					}
					} 
				}
				setState(283);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,16,_ctx);
			}
			setState(285);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__19) {
				{
				setState(284);
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
			setState(287);
			match(T__18);
			setState(288);
			match(T__4);
			setState(289);
			expression();
			setState(290);
			condExprEnd();
			setState(291);
			match(T__7);
			setState(293); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(292);
				stmt();
				}
				}
				setState(295); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__14) | (1L << T__15) | (1L << T__17) | (1L << T__18) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << T__42) | (1L << ID))) != 0) );
			setState(297);
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
			setState(299);
			match(T__19);
			setState(300);
			match(T__18);
			setState(301);
			match(T__4);
			setState(302);
			expression();
			setState(303);
			condExprEnd();
			setState(304);
			match(T__7);
			setState(306); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(305);
				stmt();
				}
				}
				setState(308); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__14) | (1L << T__15) | (1L << T__17) | (1L << T__18) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << T__42) | (1L << ID))) != 0) );
			setState(310);
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
			setState(312);
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
			setState(314);
			match(T__19);
			setState(315);
			match(T__7);
			setState(317); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(316);
				stmt();
				}
				}
				setState(319); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__14) | (1L << T__15) | (1L << T__17) | (1L << T__18) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << T__42) | (1L << ID))) != 0) );
			setState(321);
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
			setState(323);
			match(T__20);
			setState(324);
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
			setState(348);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__21:
				enterOuterAlt(_localctx, 1);
				{
				setState(326);
				inputRead();
				}
				break;
			case T__22:
				enterOuterAlt(_localctx, 2);
				{
				setState(327);
				print();
				}
				break;
			case T__23:
				enterOuterAlt(_localctx, 3);
				{
				setState(328);
				listAdd();
				}
				break;
			case T__24:
				enterOuterAlt(_localctx, 4);
				{
				setState(329);
				listPop();
				}
				break;
			case T__25:
				enterOuterAlt(_localctx, 5);
				{
				setState(330);
				length();
				}
				break;
			case T__26:
				enterOuterAlt(_localctx, 6);
				{
				setState(331);
				range();
				}
				break;
			case T__27:
				enterOuterAlt(_localctx, 7);
				{
				setState(332);
				plot();
				}
				break;
			case T__28:
				enterOuterAlt(_localctx, 8);
				{
				setState(333);
				sum();
				}
				break;
			case T__29:
				enterOuterAlt(_localctx, 9);
				{
				setState(334);
				min();
				}
				break;
			case T__30:
				enterOuterAlt(_localctx, 10);
				{
				setState(335);
				prod();
				}
				break;
			case T__31:
				enterOuterAlt(_localctx, 11);
				{
				setState(336);
				avg();
				}
				break;
			case T__32:
				enterOuterAlt(_localctx, 12);
				{
				setState(337);
				sMode();
				}
				break;
			case T__33:
				enterOuterAlt(_localctx, 13);
				{
				setState(338);
				median();
				}
				break;
			case T__34:
				enterOuterAlt(_localctx, 14);
				{
				setState(339);
				sin();
				}
				break;
			case T__36:
				enterOuterAlt(_localctx, 15);
				{
				setState(340);
				cos();
				}
				break;
			case T__35:
				enterOuterAlt(_localctx, 16);
				{
				setState(341);
				tan();
				}
				break;
			case T__37:
				enterOuterAlt(_localctx, 17);
				{
				setState(342);
				sort();
				}
				break;
			case T__38:
				enterOuterAlt(_localctx, 18);
				{
				setState(343);
				sqrt();
				}
				break;
			case T__39:
				enterOuterAlt(_localctx, 19);
				{
				setState(344);
				floor();
				}
				break;
			case T__40:
				enterOuterAlt(_localctx, 20);
				{
				setState(345);
				ceil();
				}
				break;
			case T__41:
				enterOuterAlt(_localctx, 21);
				{
				setState(346);
				abs();
				}
				break;
			case T__42:
				enterOuterAlt(_localctx, 22);
				{
				setState(347);
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
			setState(350);
			match(T__21);
			setState(351);
			match(T__4);
			setState(352);
			match(ID);
			setState(357);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__12) {
				{
				{
				setState(353);
				match(T__12);
				setState(354);
				match(ID);
				}
				}
				setState(359);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(360);
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
			setState(362);
			match(T__22);
			setState(363);
			match(T__4);
			setState(364);
			expression();
			setState(369);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__12) {
				{
				{
				setState(365);
				match(T__12);
				setState(366);
				expression();
				}
				}
				setState(371);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(372);
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
			setState(374);
			match(T__23);
			setState(375);
			match(T__4);
			setState(376);
			expression();
			setState(377);
			match(T__12);
			setState(378);
			expression();
			setState(379);
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
			setState(381);
			match(T__24);
			setState(382);
			match(T__4);
			setState(383);
			expression();
			setState(384);
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
			setState(386);
			match(T__25);
			setState(387);
			match(T__4);
			setState(388);
			expression();
			setState(389);
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
			setState(391);
			match(T__26);
			setState(392);
			match(T__4);
			setState(393);
			expression();
			setState(396);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__12) {
				{
				setState(394);
				match(T__12);
				setState(395);
				expression();
				}
			}

			setState(398);
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
			setState(400);
			match(T__27);
			setState(401);
			match(T__4);
			setState(402);
			expression();
			setState(403);
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
			setState(405);
			match(T__28);
			setState(406);
			match(T__4);
			setState(407);
			expression();
			setState(408);
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
			setState(410);
			match(T__29);
			setState(411);
			match(T__4);
			setState(412);
			expression();
			setState(413);
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
			setState(415);
			match(T__30);
			setState(416);
			match(T__4);
			setState(417);
			expression();
			setState(418);
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
			setState(420);
			match(T__31);
			setState(421);
			match(T__4);
			setState(422);
			expression();
			setState(423);
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
			setState(425);
			match(T__32);
			setState(426);
			match(T__4);
			setState(427);
			expression();
			setState(428);
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
			setState(430);
			match(T__33);
			setState(431);
			match(T__4);
			setState(432);
			expression();
			setState(433);
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
			setState(435);
			match(T__34);
			setState(436);
			match(T__4);
			setState(437);
			expression();
			setState(438);
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
			setState(440);
			match(T__35);
			setState(441);
			match(T__4);
			setState(442);
			expression();
			setState(443);
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
			setState(445);
			match(T__36);
			setState(446);
			match(T__4);
			setState(447);
			expression();
			setState(448);
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
			setState(450);
			match(T__37);
			setState(451);
			match(T__4);
			setState(452);
			expression();
			setState(453);
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
			setState(455);
			match(T__38);
			setState(456);
			match(T__4);
			setState(457);
			expression();
			setState(458);
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
			setState(460);
			match(T__39);
			setState(461);
			match(T__4);
			setState(462);
			expression();
			setState(463);
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
			setState(465);
			match(T__40);
			setState(466);
			match(T__4);
			setState(467);
			expression();
			setState(468);
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
			setState(470);
			match(T__41);
			setState(471);
			match(T__4);
			setState(472);
			expression();
			setState(473);
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
			setState(475);
			match(T__42);
			setState(476);
			match(T__4);
			setState(477);
			expression();
			setState(478);
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
		enterRule(_localctx, 94, RULE_expression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(480);
			exp();
			setState(484);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__48) | (1L << T__49) | (1L << T__50) | (1L << T__51) | (1L << T__52) | (1L << T__53) | (1L << T__54) | (1L << T__55))) != 0)) {
				{
				setState(481);
				logicOperator();
				setState(482);
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
		enterRule(_localctx, 96, RULE_exp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(486);
			term();
			setState(492);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__43 || _la==T__44) {
				{
				{
				setState(487);
				opSec();
				setState(488);
				term();
				}
				}
				setState(494);
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
		enterRule(_localctx, 98, RULE_term);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(495);
			factor();
			setState(501);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__45) | (1L << T__46) | (1L << T__47))) != 0)) {
				{
				{
				setState(496);
				opFirst();
				setState(497);
				factor();
				}
				}
				setState(503);
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
		public IndexingContext indexing() {
			return getRuleContext(IndexingContext.class,0);
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
		enterRule(_localctx, 100, RULE_factor);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(505);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__43) {
				{
				setState(504);
				unaryMinus();
				}
			}

			setState(512);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,29,_ctx) ) {
			case 1:
				{
				setState(507);
				nestedExpression();
				}
				break;
			case 2:
				{
				setState(508);
				indexing();
				}
				break;
			case 3:
				{
				setState(509);
				specialFunction();
				}
				break;
			case 4:
				{
				setState(510);
				functionCall();
				}
				break;
			case 5:
				{
				setState(511);
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
		enterRule(_localctx, 102, RULE_unaryMinus);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(514);
			match(T__43);
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
		enterRule(_localctx, 104, RULE_nestedExpression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(516);
			match(T__4);
			setState(517);
			expression();
			setState(518);
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
		enterRule(_localctx, 106, RULE_functionCall);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(520);
			match(ID);
			setState(521);
			match(T__4);
			setState(530);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (((((_la - 5)) & ~0x3f) == 0 && ((1L << (_la - 5)) & ((1L << (T__4 - 5)) | (1L << (T__21 - 5)) | (1L << (T__22 - 5)) | (1L << (T__23 - 5)) | (1L << (T__24 - 5)) | (1L << (T__25 - 5)) | (1L << (T__26 - 5)) | (1L << (T__27 - 5)) | (1L << (T__28 - 5)) | (1L << (T__29 - 5)) | (1L << (T__30 - 5)) | (1L << (T__31 - 5)) | (1L << (T__32 - 5)) | (1L << (T__33 - 5)) | (1L << (T__34 - 5)) | (1L << (T__35 - 5)) | (1L << (T__36 - 5)) | (1L << (T__37 - 5)) | (1L << (T__38 - 5)) | (1L << (T__39 - 5)) | (1L << (T__40 - 5)) | (1L << (T__41 - 5)) | (1L << (T__42 - 5)) | (1L << (T__43 - 5)) | (1L << (BOOL_CONS - 5)) | (1L << (ID - 5)) | (1L << (INT_CONS - 5)) | (1L << (STRING_CONS - 5)) | (1L << (FLOAT_CONS - 5)))) != 0)) {
				{
				setState(522);
				expression();
				setState(527);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__12) {
					{
					{
					setState(523);
					match(T__12);
					setState(524);
					expression();
					}
					}
					setState(529);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(532);
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

	public static class IndexingContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(BistatParser.ID, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public IndexingContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_indexing; }
	}

	public final IndexingContext indexing() throws RecognitionException {
		IndexingContext _localctx = new IndexingContext(_ctx, getState());
		enterRule(_localctx, 108, RULE_indexing);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(534);
			match(ID);
			setState(535);
			match(T__11);
			setState(536);
			expression();
			setState(537);
			match(T__13);
			setState(542);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__11) {
				{
				setState(538);
				match(T__11);
				setState(539);
				expression();
				setState(540);
				match(T__13);
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
		enterRule(_localctx, 110, RULE_varCons);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(544);
			_la = _input.LA(1);
			if ( !(((((_la - 60)) & ~0x3f) == 0 && ((1L << (_la - 60)) & ((1L << (BOOL_CONS - 60)) | (1L << (ID - 60)) | (1L << (INT_CONS - 60)) | (1L << (STRING_CONS - 60)) | (1L << (FLOAT_CONS - 60)))) != 0)) ) {
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
		enterRule(_localctx, 112, RULE_opSec);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(546);
			_la = _input.LA(1);
			if ( !(_la==T__43 || _la==T__44) ) {
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
		enterRule(_localctx, 114, RULE_opFirst);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(548);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__45) | (1L << T__46) | (1L << T__47))) != 0)) ) {
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
		enterRule(_localctx, 116, RULE_logicOperator);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(550);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__48) | (1L << T__49) | (1L << T__50) | (1L << T__51) | (1L << T__52) | (1L << T__53) | (1L << T__54) | (1L << T__55))) != 0)) ) {
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3C\u022b\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\3\2\3"+
		"\2\3\2\3\2\7\2}\n\2\f\2\16\2\u0080\13\2\3\2\7\2\u0083\n\2\f\2\16\2\u0086"+
		"\13\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\4\3\4\5\4\u0092\n\4\3\5\3\5\3"+
		"\5\3\5\7\5\u0098\n\5\f\5\16\5\u009b\13\5\3\5\3\5\3\5\3\5\7\5\u00a1\n\5"+
		"\f\5\16\5\u00a4\13\5\3\5\3\5\6\5\u00a8\n\5\r\5\16\5\u00a9\3\5\3\5\3\6"+
		"\3\6\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\6\t\u00bc\n\t\r\t"+
		"\16\t\u00bd\3\t\3\t\3\n\3\n\3\n\3\n\5\n\u00c6\n\n\3\n\3\n\3\n\3\n\3\n"+
		"\3\n\5\n\u00ce\n\n\3\13\3\13\5\13\u00d2\n\13\3\13\3\13\3\13\3\13\5\13"+
		"\u00d8\n\13\3\f\3\f\3\f\3\f\7\f\u00de\n\f\f\f\16\f\u00e1\13\f\3\f\3\f"+
		"\3\r\3\r\3\r\3\r\7\r\u00e9\n\r\f\r\16\r\u00ec\13\r\3\r\3\r\3\16\3\16\6"+
		"\16\u00f2\n\16\r\16\16\16\u00f3\3\16\3\16\3\17\3\17\3\17\6\17\u00fb\n"+
		"\17\r\17\16\17\u00fc\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\21"+
		"\3\21\3\22\3\22\3\22\3\22\3\22\3\22\6\22\u0110\n\22\r\22\16\22\u0111\3"+
		"\22\3\22\3\23\3\23\3\24\3\24\7\24\u011a\n\24\f\24\16\24\u011d\13\24\3"+
		"\24\5\24\u0120\n\24\3\25\3\25\3\25\3\25\3\25\3\25\6\25\u0128\n\25\r\25"+
		"\16\25\u0129\3\25\3\25\3\26\3\26\3\26\3\26\3\26\3\26\3\26\6\26\u0135\n"+
		"\26\r\26\16\26\u0136\3\26\3\26\3\27\3\27\3\30\3\30\3\30\6\30\u0140\n\30"+
		"\r\30\16\30\u0141\3\30\3\30\3\31\3\31\3\31\3\32\3\32\3\32\3\32\3\32\3"+
		"\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3"+
		"\32\3\32\3\32\5\32\u015f\n\32\3\33\3\33\3\33\3\33\3\33\7\33\u0166\n\33"+
		"\f\33\16\33\u0169\13\33\3\33\3\33\3\34\3\34\3\34\3\34\3\34\7\34\u0172"+
		"\n\34\f\34\16\34\u0175\13\34\3\34\3\34\3\35\3\35\3\35\3\35\3\35\3\35\3"+
		"\35\3\36\3\36\3\36\3\36\3\36\3\37\3\37\3\37\3\37\3\37\3 \3 \3 \3 \3 \5"+
		" \u018f\n \3 \3 \3!\3!\3!\3!\3!\3\"\3\"\3\"\3\"\3\"\3#\3#\3#\3#\3#\3$"+
		"\3$\3$\3$\3$\3%\3%\3%\3%\3%\3&\3&\3&\3&\3&\3\'\3\'\3\'\3\'\3\'\3(\3(\3"+
		"(\3(\3(\3)\3)\3)\3)\3)\3*\3*\3*\3*\3*\3+\3+\3+\3+\3+\3,\3,\3,\3,\3,\3"+
		"-\3-\3-\3-\3-\3.\3.\3.\3.\3.\3/\3/\3/\3/\3/\3\60\3\60\3\60\3\60\3\60\3"+
		"\61\3\61\3\61\3\61\5\61\u01e7\n\61\3\62\3\62\3\62\3\62\7\62\u01ed\n\62"+
		"\f\62\16\62\u01f0\13\62\3\63\3\63\3\63\3\63\7\63\u01f6\n\63\f\63\16\63"+
		"\u01f9\13\63\3\64\5\64\u01fc\n\64\3\64\3\64\3\64\3\64\3\64\5\64\u0203"+
		"\n\64\3\65\3\65\3\66\3\66\3\66\3\66\3\67\3\67\3\67\3\67\3\67\7\67\u0210"+
		"\n\67\f\67\16\67\u0213\13\67\5\67\u0215\n\67\3\67\3\67\38\38\38\38\38"+
		"\38\38\38\58\u0221\n8\39\39\3:\3:\3;\3;\3<\3<\3<\2\2=\2\4\6\b\n\f\16\20"+
		"\22\24\26\30\32\34\36 \"$&(*,.\60\62\64\668:<>@BDFHJLNPRTVXZ\\^`bdfhj"+
		"lnprtv\2\7\3\2\21\21\4\2>@BC\3\2./\3\2\60\62\3\2\63:\2\u022d\2x\3\2\2"+
		"\2\4\u008a\3\2\2\2\6\u008f\3\2\2\2\b\u0093\3\2\2\2\n\u00ad\3\2\2\2\f\u00af"+
		"\3\2\2\2\16\u00b1\3\2\2\2\20\u00b6\3\2\2\2\22\u00cd\3\2\2\2\24\u00d1\3"+
		"\2\2\2\26\u00d9\3\2\2\2\30\u00e4\3\2\2\2\32\u00ef\3\2\2\2\34\u00f7\3\2"+
		"\2\2\36\u0100\3\2\2\2 \u0107\3\2\2\2\"\u0109\3\2\2\2$\u0115\3\2\2\2&\u0117"+
		"\3\2\2\2(\u0121\3\2\2\2*\u012d\3\2\2\2,\u013a\3\2\2\2.\u013c\3\2\2\2\60"+
		"\u0145\3\2\2\2\62\u015e\3\2\2\2\64\u0160\3\2\2\2\66\u016c\3\2\2\28\u0178"+
		"\3\2\2\2:\u017f\3\2\2\2<\u0184\3\2\2\2>\u0189\3\2\2\2@\u0192\3\2\2\2B"+
		"\u0197\3\2\2\2D\u019c\3\2\2\2F\u01a1\3\2\2\2H\u01a6\3\2\2\2J\u01ab\3\2"+
		"\2\2L\u01b0\3\2\2\2N\u01b5\3\2\2\2P\u01ba\3\2\2\2R\u01bf\3\2\2\2T\u01c4"+
		"\3\2\2\2V\u01c9\3\2\2\2X\u01ce\3\2\2\2Z\u01d3\3\2\2\2\\\u01d8\3\2\2\2"+
		"^\u01dd\3\2\2\2`\u01e2\3\2\2\2b\u01e8\3\2\2\2d\u01f1\3\2\2\2f\u01fb\3"+
		"\2\2\2h\u0204\3\2\2\2j\u0206\3\2\2\2l\u020a\3\2\2\2n\u0218\3\2\2\2p\u0222"+
		"\3\2\2\2r\u0224\3\2\2\2t\u0226\3\2\2\2v\u0228\3\2\2\2xy\7\3\2\2yz\7?\2"+
		"\2z~\7\4\2\2{}\5\4\3\2|{\3\2\2\2}\u0080\3\2\2\2~|\3\2\2\2~\177\3\2\2\2"+
		"\177\u0084\3\2\2\2\u0080~\3\2\2\2\u0081\u0083\5\b\5\2\u0082\u0081\3\2"+
		"\2\2\u0083\u0086\3\2\2\2\u0084\u0082\3\2\2\2\u0084\u0085\3\2\2\2\u0085"+
		"\u0087\3\2\2\2\u0086\u0084\3\2\2\2\u0087\u0088\5\20\t\2\u0088\u0089\7"+
		"\2\2\3\u0089\3\3\2\2\2\u008a\u008b\7\5\2\2\u008b\u008c\5\6\4\2\u008c\u008d"+
		"\7?\2\2\u008d\u008e\7\4\2\2\u008e\5\3\2\2\2\u008f\u0091\7=\2\2\u0090\u0092"+
		"\7<\2\2\u0091\u0090\3\2\2\2\u0091\u0092\3\2\2\2\u0092\7\3\2\2\2\u0093"+
		"\u0094\7\6\2\2\u0094\u0095\7?\2\2\u0095\u0099\7\7\2\2\u0096\u0098\5\16"+
		"\b\2\u0097\u0096\3\2\2\2\u0098\u009b\3\2\2\2\u0099\u0097\3\2\2\2\u0099"+
		"\u009a\3\2\2\2\u009a\u009c\3\2\2\2\u009b\u0099\3\2\2\2\u009c\u009d\7\b"+
		"\2\2\u009d\u009e\7\t\2\2\u009e\u00a2\5\6\4\2\u009f\u00a1\5\4\3\2\u00a0"+
		"\u009f\3\2\2\2\u00a1\u00a4\3\2\2\2\u00a2\u00a0\3\2\2\2\u00a2\u00a3\3\2"+
		"\2\2\u00a3\u00a5\3\2\2\2\u00a4\u00a2\3\2\2\2\u00a5\u00a7\5\n\6\2\u00a6"+
		"\u00a8\5\22\n\2\u00a7\u00a6\3\2\2\2\u00a8\u00a9\3\2\2\2\u00a9\u00a7\3"+
		"\2\2\2\u00a9\u00aa\3\2\2\2\u00aa\u00ab\3\2\2\2\u00ab\u00ac\5\f\7\2\u00ac"+
		"\t\3\2\2\2\u00ad\u00ae\7\n\2\2\u00ae\13\3\2\2\2\u00af\u00b0\7\13\2\2\u00b0"+
		"\r\3\2\2\2\u00b1\u00b2\7\5\2\2\u00b2\u00b3\5\6\4\2\u00b3\u00b4\7?\2\2"+
		"\u00b4\u00b5\7\4\2\2\u00b5\17\3\2\2\2\u00b6\u00b7\7\f\2\2\u00b7\u00b8"+
		"\7\7\2\2\u00b8\u00b9\7\b\2\2\u00b9\u00bb\7\n\2\2\u00ba\u00bc\5\22\n\2"+
		"\u00bb\u00ba\3\2\2\2\u00bc\u00bd\3\2\2\2\u00bd\u00bb\3\2\2\2\u00bd\u00be"+
		"\3\2\2\2\u00be\u00bf\3\2\2\2\u00bf\u00c0\7\13\2\2\u00c0\21\3\2\2\2\u00c1"+
		"\u00c6\5\24\13\2\u00c2\u00c6\5\62\32\2\u00c3\u00c6\5l\67\2\u00c4\u00c6"+
		"\5\60\31\2\u00c5\u00c1\3\2\2\2\u00c5\u00c2\3\2\2\2\u00c5\u00c3\3\2\2\2"+
		"\u00c5\u00c4\3\2\2\2\u00c6\u00c7\3\2\2\2\u00c7\u00c8\7\4\2\2\u00c8\u00ce"+
		"\3\2\2\2\u00c9\u00ce\5&\24\2\u00ca\u00ce\5\"\22\2\u00cb\u00ce\5\34\17"+
		"\2\u00cc\u00ce\5\32\16\2\u00cd\u00c5\3\2\2\2\u00cd\u00c9\3\2\2\2\u00cd"+
		"\u00ca\3\2\2\2\u00cd\u00cb\3\2\2\2\u00cd\u00cc\3\2\2\2\u00ce\23\3\2\2"+
		"\2\u00cf\u00d2\5n8\2\u00d0\u00d2\7?\2\2\u00d1\u00cf\3\2\2\2\u00d1\u00d0"+
		"\3\2\2\2\u00d2\u00d3\3\2\2\2\u00d3\u00d7\7\r\2\2\u00d4\u00d8\5`\61\2\u00d5"+
		"\u00d8\5\30\r\2\u00d6\u00d8\5\26\f\2\u00d7\u00d4\3\2\2\2\u00d7\u00d5\3"+
		"\2\2\2\u00d7\u00d6\3\2\2\2\u00d8\25\3\2\2\2\u00d9\u00da\7\16\2\2\u00da"+
		"\u00df\5\30\r\2\u00db\u00dc\7\17\2\2\u00dc\u00de\5\30\r\2\u00dd\u00db"+
		"\3\2\2\2\u00de\u00e1\3\2\2\2\u00df\u00dd\3\2\2\2\u00df\u00e0\3\2\2\2\u00e0"+
		"\u00e2\3\2\2\2\u00e1\u00df\3\2\2\2\u00e2\u00e3\7\20\2\2\u00e3\27\3\2\2"+
		"\2\u00e4\u00e5\7\16\2\2\u00e5\u00ea\5`\61\2\u00e6\u00e7\7\17\2\2\u00e7"+
		"\u00e9\5`\61\2\u00e8\u00e6\3\2\2\2\u00e9\u00ec\3\2\2\2\u00ea\u00e8\3\2"+
		"\2\2\u00ea\u00eb\3\2\2\2\u00eb\u00ed\3\2\2\2\u00ec\u00ea\3\2\2\2\u00ed"+
		"\u00ee\7\20\2\2\u00ee\31\3\2\2\2\u00ef\u00f1\7\21\2\2\u00f0\u00f2\n\2"+
		"\2\2\u00f1\u00f0\3\2\2\2\u00f2\u00f3\3\2\2\2\u00f3\u00f1\3\2\2\2\u00f3"+
		"\u00f4\3\2\2\2\u00f4\u00f5\3\2\2\2\u00f5\u00f6\7\21\2\2\u00f6\33\3\2\2"+
		"\2\u00f7\u00f8\5\36\20\2\u00f8\u00fa\7\n\2\2\u00f9\u00fb\5\22\n\2\u00fa"+
		"\u00f9\3\2\2\2\u00fb\u00fc\3\2\2\2\u00fc\u00fa\3\2\2\2\u00fc\u00fd\3\2"+
		"\2\2\u00fd\u00fe\3\2\2\2\u00fe\u00ff\7\13\2\2\u00ff\35\3\2\2\2\u0100\u0101"+
		"\7\22\2\2\u0101\u0102\7\7\2\2\u0102\u0103\7?\2\2\u0103\u0104\7\23\2\2"+
		"\u0104\u0105\5`\61\2\u0105\u0106\7\b\2\2\u0106\37\3\2\2\2\u0107\u0108"+
		"\7\b\2\2\u0108!\3\2\2\2\u0109\u010a\7\24\2\2\u010a\u010b\7\7\2\2\u010b"+
		"\u010c\5`\61\2\u010c\u010d\5$\23\2\u010d\u010f\7\n\2\2\u010e\u0110\5\22"+
		"\n\2\u010f\u010e\3\2\2\2\u0110\u0111\3\2\2\2\u0111\u010f\3\2\2\2\u0111"+
		"\u0112\3\2\2\2\u0112\u0113\3\2\2\2\u0113\u0114\7\13\2\2\u0114#\3\2\2\2"+
		"\u0115\u0116\7\b\2\2\u0116%\3\2\2\2\u0117\u011b\5(\25\2\u0118\u011a\5"+
		"*\26\2\u0119\u0118\3\2\2\2\u011a\u011d\3\2\2\2\u011b\u0119\3\2\2\2\u011b"+
		"\u011c\3\2\2\2\u011c\u011f\3\2\2\2\u011d\u011b\3\2\2\2\u011e\u0120\5."+
		"\30\2\u011f\u011e\3\2\2\2\u011f\u0120\3\2\2\2\u0120\'\3\2\2\2\u0121\u0122"+
		"\7\25\2\2\u0122\u0123\7\7\2\2\u0123\u0124\5`\61\2\u0124\u0125\5,\27\2"+
		"\u0125\u0127\7\n\2\2\u0126\u0128\5\22\n\2\u0127\u0126\3\2\2\2\u0128\u0129"+
		"\3\2\2\2\u0129\u0127\3\2\2\2\u0129\u012a\3\2\2\2\u012a\u012b\3\2\2\2\u012b"+
		"\u012c\7\13\2\2\u012c)\3\2\2\2\u012d\u012e\7\26\2\2\u012e\u012f\7\25\2"+
		"\2\u012f\u0130\7\7\2\2\u0130\u0131\5`\61\2\u0131\u0132\5,\27\2\u0132\u0134"+
		"\7\n\2\2\u0133\u0135\5\22\n\2\u0134\u0133\3\2\2\2\u0135\u0136\3\2\2\2"+
		"\u0136\u0134\3\2\2\2\u0136\u0137\3\2\2\2\u0137\u0138\3\2\2\2\u0138\u0139"+
		"\7\13\2\2\u0139+\3\2\2\2\u013a\u013b\7\b\2\2\u013b-\3\2\2\2\u013c\u013d"+
		"\7\26\2\2\u013d\u013f\7\n\2\2\u013e\u0140\5\22\n\2\u013f\u013e\3\2\2\2"+
		"\u0140\u0141\3\2\2\2\u0141\u013f\3\2\2\2\u0141\u0142\3\2\2\2\u0142\u0143"+
		"\3\2\2\2\u0143\u0144\7\13\2\2\u0144/\3\2\2\2\u0145\u0146\7\27\2\2\u0146"+
		"\u0147\5`\61\2\u0147\61\3\2\2\2\u0148\u015f\5\64\33\2\u0149\u015f\5\66"+
		"\34\2\u014a\u015f\58\35\2\u014b\u015f\5:\36\2\u014c\u015f\5<\37\2\u014d"+
		"\u015f\5> \2\u014e\u015f\5@!\2\u014f\u015f\5B\"\2\u0150\u015f\5D#\2\u0151"+
		"\u015f\5F$\2\u0152\u015f\5H%\2\u0153\u015f\5J&\2\u0154\u015f\5L\'\2\u0155"+
		"\u015f\5N(\2\u0156\u015f\5R*\2\u0157\u015f\5P)\2\u0158\u015f\5T+\2\u0159"+
		"\u015f\5V,\2\u015a\u015f\5X-\2\u015b\u015f\5Z.\2\u015c\u015f\5\\/\2\u015d"+
		"\u015f\5^\60\2\u015e\u0148\3\2\2\2\u015e\u0149\3\2\2\2\u015e\u014a\3\2"+
		"\2\2\u015e\u014b\3\2\2\2\u015e\u014c\3\2\2\2\u015e\u014d\3\2\2\2\u015e"+
		"\u014e\3\2\2\2\u015e\u014f\3\2\2\2\u015e\u0150\3\2\2\2\u015e\u0151\3\2"+
		"\2\2\u015e\u0152\3\2\2\2\u015e\u0153\3\2\2\2\u015e\u0154\3\2\2\2\u015e"+
		"\u0155\3\2\2\2\u015e\u0156\3\2\2\2\u015e\u0157\3\2\2\2\u015e\u0158\3\2"+
		"\2\2\u015e\u0159\3\2\2\2\u015e\u015a\3\2\2\2\u015e\u015b\3\2\2\2\u015e"+
		"\u015c\3\2\2\2\u015e\u015d\3\2\2\2\u015f\63\3\2\2\2\u0160\u0161\7\30\2"+
		"\2\u0161\u0162\7\7\2\2\u0162\u0167\7?\2\2\u0163\u0164\7\17\2\2\u0164\u0166"+
		"\7?\2\2\u0165\u0163\3\2\2\2\u0166\u0169\3\2\2\2\u0167\u0165\3\2\2\2\u0167"+
		"\u0168\3\2\2\2\u0168\u016a\3\2\2\2\u0169\u0167\3\2\2\2\u016a\u016b\7\b"+
		"\2\2\u016b\65\3\2\2\2\u016c\u016d\7\31\2\2\u016d\u016e\7\7\2\2\u016e\u0173"+
		"\5`\61\2\u016f\u0170\7\17\2\2\u0170\u0172\5`\61\2\u0171\u016f\3\2\2\2"+
		"\u0172\u0175\3\2\2\2\u0173\u0171\3\2\2\2\u0173\u0174\3\2\2\2\u0174\u0176"+
		"\3\2\2\2\u0175\u0173\3\2\2\2\u0176\u0177\7\b\2\2\u0177\67\3\2\2\2\u0178"+
		"\u0179\7\32\2\2\u0179\u017a\7\7\2\2\u017a\u017b\5`\61\2\u017b\u017c\7"+
		"\17\2\2\u017c\u017d\5`\61\2\u017d\u017e\7\b\2\2\u017e9\3\2\2\2\u017f\u0180"+
		"\7\33\2\2\u0180\u0181\7\7\2\2\u0181\u0182\5`\61\2\u0182\u0183\7\b\2\2"+
		"\u0183;\3\2\2\2\u0184\u0185\7\34\2\2\u0185\u0186\7\7\2\2\u0186\u0187\5"+
		"`\61\2\u0187\u0188\7\b\2\2\u0188=\3\2\2\2\u0189\u018a\7\35\2\2\u018a\u018b"+
		"\7\7\2\2\u018b\u018e\5`\61\2\u018c\u018d\7\17\2\2\u018d\u018f\5`\61\2"+
		"\u018e\u018c\3\2\2\2\u018e\u018f\3\2\2\2\u018f\u0190\3\2\2\2\u0190\u0191"+
		"\7\b\2\2\u0191?\3\2\2\2\u0192\u0193\7\36\2\2\u0193\u0194\7\7\2\2\u0194"+
		"\u0195\5`\61\2\u0195\u0196\7\b\2\2\u0196A\3\2\2\2\u0197\u0198\7\37\2\2"+
		"\u0198\u0199\7\7\2\2\u0199\u019a\5`\61\2\u019a\u019b\7\b\2\2\u019bC\3"+
		"\2\2\2\u019c\u019d\7 \2\2\u019d\u019e\7\7\2\2\u019e\u019f\5`\61\2\u019f"+
		"\u01a0\7\b\2\2\u01a0E\3\2\2\2\u01a1\u01a2\7!\2\2\u01a2\u01a3\7\7\2\2\u01a3"+
		"\u01a4\5`\61\2\u01a4\u01a5\7\b\2\2\u01a5G\3\2\2\2\u01a6\u01a7\7\"\2\2"+
		"\u01a7\u01a8\7\7\2\2\u01a8\u01a9\5`\61\2\u01a9\u01aa\7\b\2\2\u01aaI\3"+
		"\2\2\2\u01ab\u01ac\7#\2\2\u01ac\u01ad\7\7\2\2\u01ad\u01ae\5`\61\2\u01ae"+
		"\u01af\7\b\2\2\u01afK\3\2\2\2\u01b0\u01b1\7$\2\2\u01b1\u01b2\7\7\2\2\u01b2"+
		"\u01b3\5`\61\2\u01b3\u01b4\7\b\2\2\u01b4M\3\2\2\2\u01b5\u01b6\7%\2\2\u01b6"+
		"\u01b7\7\7\2\2\u01b7\u01b8\5`\61\2\u01b8\u01b9\7\b\2\2\u01b9O\3\2\2\2"+
		"\u01ba\u01bb\7&\2\2\u01bb\u01bc\7\7\2\2\u01bc\u01bd\5`\61\2\u01bd\u01be"+
		"\7\b\2\2\u01beQ\3\2\2\2\u01bf\u01c0\7\'\2\2\u01c0\u01c1\7\7\2\2\u01c1"+
		"\u01c2\5`\61\2\u01c2\u01c3\7\b\2\2\u01c3S\3\2\2\2\u01c4\u01c5\7(\2\2\u01c5"+
		"\u01c6\7\7\2\2\u01c6\u01c7\5`\61\2\u01c7\u01c8\7\b\2\2\u01c8U\3\2\2\2"+
		"\u01c9\u01ca\7)\2\2\u01ca\u01cb\7\7\2\2\u01cb\u01cc\5`\61\2\u01cc\u01cd"+
		"\7\b\2\2\u01cdW\3\2\2\2\u01ce\u01cf\7*\2\2\u01cf\u01d0\7\7\2\2\u01d0\u01d1"+
		"\5`\61\2\u01d1\u01d2\7\b\2\2\u01d2Y\3\2\2\2\u01d3\u01d4\7+\2\2\u01d4\u01d5"+
		"\7\7\2\2\u01d5\u01d6\5`\61\2\u01d6\u01d7\7\b\2\2\u01d7[\3\2\2\2\u01d8"+
		"\u01d9\7,\2\2\u01d9\u01da\7\7\2\2\u01da\u01db\5`\61\2\u01db\u01dc\7\b"+
		"\2\2\u01dc]\3\2\2\2\u01dd\u01de\7-\2\2\u01de\u01df\7\7\2\2\u01df\u01e0"+
		"\5`\61\2\u01e0\u01e1\7\b\2\2\u01e1_\3\2\2\2\u01e2\u01e6\5b\62\2\u01e3"+
		"\u01e4\5v<\2\u01e4\u01e5\5b\62\2\u01e5\u01e7\3\2\2\2\u01e6\u01e3\3\2\2"+
		"\2\u01e6\u01e7\3\2\2\2\u01e7a\3\2\2\2\u01e8\u01ee\5d\63\2\u01e9\u01ea"+
		"\5r:\2\u01ea\u01eb\5d\63\2\u01eb\u01ed\3\2\2\2\u01ec\u01e9\3\2\2\2\u01ed"+
		"\u01f0\3\2\2\2\u01ee\u01ec\3\2\2\2\u01ee\u01ef\3\2\2\2\u01efc\3\2\2\2"+
		"\u01f0\u01ee\3\2\2\2\u01f1\u01f7\5f\64\2\u01f2\u01f3\5t;\2\u01f3\u01f4"+
		"\5f\64\2\u01f4\u01f6\3\2\2\2\u01f5\u01f2\3\2\2\2\u01f6\u01f9\3\2\2\2\u01f7"+
		"\u01f5\3\2\2\2\u01f7\u01f8\3\2\2\2\u01f8e\3\2\2\2\u01f9\u01f7\3\2\2\2"+
		"\u01fa\u01fc\5h\65\2\u01fb\u01fa\3\2\2\2\u01fb\u01fc\3\2\2\2\u01fc\u0202"+
		"\3\2\2\2\u01fd\u0203\5j\66\2\u01fe\u0203\5n8\2\u01ff\u0203\5\62\32\2\u0200"+
		"\u0203\5l\67\2\u0201\u0203\5p9\2\u0202\u01fd\3\2\2\2\u0202\u01fe\3\2\2"+
		"\2\u0202\u01ff\3\2\2\2\u0202\u0200\3\2\2\2\u0202\u0201\3\2\2\2\u0203g"+
		"\3\2\2\2\u0204\u0205\7.\2\2\u0205i\3\2\2\2\u0206\u0207\7\7\2\2\u0207\u0208"+
		"\5`\61\2\u0208\u0209\7\b\2\2\u0209k\3\2\2\2\u020a\u020b\7?\2\2\u020b\u0214"+
		"\7\7\2\2\u020c\u0211\5`\61\2\u020d\u020e\7\17\2\2\u020e\u0210\5`\61\2"+
		"\u020f\u020d\3\2\2\2\u0210\u0213\3\2\2\2\u0211\u020f\3\2\2\2\u0211\u0212"+
		"\3\2\2\2\u0212\u0215\3\2\2\2\u0213\u0211\3\2\2\2\u0214\u020c\3\2\2\2\u0214"+
		"\u0215\3\2\2\2\u0215\u0216\3\2\2\2\u0216\u0217\7\b\2\2\u0217m\3\2\2\2"+
		"\u0218\u0219\7?\2\2\u0219\u021a\7\16\2\2\u021a\u021b\5`\61\2\u021b\u0220"+
		"\7\20\2\2\u021c\u021d\7\16\2\2\u021d\u021e\5`\61\2\u021e\u021f\7\20\2"+
		"\2\u021f\u0221\3\2\2\2\u0220\u021c\3\2\2\2\u0220\u0221\3\2\2\2\u0221o"+
		"\3\2\2\2\u0222\u0223\t\3\2\2\u0223q\3\2\2\2\u0224\u0225\t\4\2\2\u0225"+
		"s\3\2\2\2\u0226\u0227\t\5\2\2\u0227u\3\2\2\2\u0228\u0229\t\6\2\2\u0229"+
		"w\3\2\2\2#~\u0084\u0091\u0099\u00a2\u00a9\u00bd\u00c5\u00cd\u00d1\u00d7"+
		"\u00df\u00ea\u00f3\u00fc\u0111\u011b\u011f\u0129\u0136\u0141\u015e\u0167"+
		"\u0173\u018e\u01e6\u01ee\u01f7\u01fb\u0202\u0211\u0214\u0220";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}