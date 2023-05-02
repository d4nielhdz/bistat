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
		T__38=39, T__39=40, T__40=41, T__41=42, T__42=43, T__43=44, WS=45, CARDINALITY=46, 
		TYPE_PRIMITIVE=47, ID=48, INT_CONS=49, NUMBER=50, BOOL_CONS=51, STRING_CONS=52, 
		FLOAT_CONS=53, OP_SEC=54, OP_FIRST=55, LOGIC_OPERATOR=56;
	public static final int
		RULE_program = 0, RULE_varDeclaration = 1, RULE_var_type = 2, RULE_funcDef = 3, 
		RULE_funcEnd = 4, RULE_paramDeclaration = 5, RULE_main = 6, RULE_stmt = 7, 
		RULE_assignment = 8, RULE_nested_stmt = 9, RULE_matrixAssignment = 10, 
		RULE_listAssignment = 11, RULE_comment = 12, RULE_forLoop = 13, RULE_whileLoop = 14, 
		RULE_conditional = 15, RULE_returnStmt = 16, RULE_specialFunction = 17, 
		RULE_inputRead = 18, RULE_print = 19, RULE_listAdd = 20, RULE_listPop = 21, 
		RULE_length = 22, RULE_range = 23, RULE_plot = 24, RULE_sum = 25, RULE_min = 26, 
		RULE_prod = 27, RULE_avg = 28, RULE_sMode = 29, RULE_median = 30, RULE_sin = 31, 
		RULE_tan = 32, RULE_cos = 33, RULE_sort = 34, RULE_sqrt = 35, RULE_floor = 36, 
		RULE_ceil = 37, RULE_abs = 38, RULE_not = 39, RULE_expression = 40, RULE_exp = 41, 
		RULE_term = 42, RULE_factor = 43, RULE_functionCall = 44, RULE_indexing = 45, 
		RULE_var_cons = 46;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "varDeclaration", "var_type", "funcDef", "funcEnd", "paramDeclaration", 
			"main", "stmt", "assignment", "nested_stmt", "matrixAssignment", "listAssignment", 
			"comment", "forLoop", "whileLoop", "conditional", "returnStmt", "specialFunction", 
			"inputRead", "print", "listAdd", "listPop", "length", "range", "plot", 
			"sum", "min", "prod", "avg", "sMode", "median", "sin", "tan", "cos", 
			"sort", "sqrt", "floor", "ceil", "abs", "not", "expression", "exp", "term", 
			"factor", "functionCall", "indexing", "var_cons"
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
			"TYPE_PRIMITIVE", "ID", "INT_CONS", "NUMBER", "BOOL_CONS", "STRING_CONS", 
			"FLOAT_CONS", "OP_SEC", "OP_FIRST", "LOGIC_OPERATOR"
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
			setState(94);
			match(T__0);
			setState(95);
			match(ID);
			setState(96);
			match(T__1);
			setState(100);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__2) {
				{
				{
				setState(97);
				varDeclaration();
				}
				}
				setState(102);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(106);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__3) {
				{
				{
				setState(103);
				funcDef();
				}
				}
				setState(108);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(109);
			main();
			setState(110);
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
			setState(112);
			match(T__2);
			setState(113);
			var_type();
			setState(114);
			match(ID);
			setState(115);
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
			setState(117);
			match(TYPE_PRIMITIVE);
			{
			setState(119);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==CARDINALITY) {
				{
				setState(118);
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
		public FuncEndContext funcEnd() {
			return getRuleContext(FuncEndContext.class,0);
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
			setState(121);
			match(T__3);
			setState(122);
			match(ID);
			setState(123);
			match(T__4);
			setState(127);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__2) {
				{
				{
				setState(124);
				paramDeclaration();
				}
				}
				setState(129);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(130);
			match(T__5);
			setState(131);
			match(T__6);
			setState(132);
			var_type();
			setState(136);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__2) {
				{
				{
				setState(133);
				varDeclaration();
				}
				}
				setState(138);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(139);
			match(T__7);
			setState(141); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(140);
				stmt();
				}
				}
				setState(143); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__14) | (1L << T__15) | (1L << T__17) | (1L << T__18) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << T__42) | (1L << ID))) != 0) );
			setState(145);
			funcEnd();
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

	public static class FuncEndContext extends ParserRuleContext {
		public FuncEndContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcEnd; }
	}

	public final FuncEndContext funcEnd() throws RecognitionException {
		FuncEndContext _localctx = new FuncEndContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_funcEnd);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(147);
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
		enterRule(_localctx, 10, RULE_paramDeclaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(149);
			match(T__2);
			setState(150);
			var_type();
			setState(151);
			match(ID);
			setState(152);
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
		enterRule(_localctx, 12, RULE_main);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(154);
			match(T__9);
			setState(155);
			match(T__4);
			setState(156);
			match(T__5);
			setState(157);
			match(T__7);
			setState(159); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(158);
				stmt();
				}
				}
				setState(161); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__14) | (1L << T__15) | (1L << T__17) | (1L << T__18) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << T__42) | (1L << ID))) != 0) );
			setState(163);
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
		enterRule(_localctx, 14, RULE_stmt);
		try {
			setState(177);
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
				setState(169);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
				case 1:
					{
					setState(165);
					assignment();
					}
					break;
				case 2:
					{
					setState(166);
					specialFunction();
					}
					break;
				case 3:
					{
					setState(167);
					functionCall();
					}
					break;
				case 4:
					{
					setState(168);
					returnStmt();
					}
					break;
				}
				setState(171);
				match(T__1);
				}
				}
				break;
			case T__18:
				enterOuterAlt(_localctx, 2);
				{
				setState(173);
				conditional();
				}
				break;
			case T__17:
				enterOuterAlt(_localctx, 3);
				{
				setState(174);
				whileLoop();
				}
				break;
			case T__15:
				enterOuterAlt(_localctx, 4);
				{
				setState(175);
				forLoop();
				}
				break;
			case T__14:
				enterOuterAlt(_localctx, 5);
				{
				setState(176);
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
		enterRule(_localctx, 16, RULE_assignment);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(179);
			match(ID);
			setState(180);
			match(T__10);
			setState(184);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				{
				setState(181);
				expression();
				}
				break;
			case 2:
				{
				setState(182);
				listAssignment();
				}
				break;
			case 3:
				{
				setState(183);
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

	public static class Nested_stmtContext extends ParserRuleContext {
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
		public CommentContext comment() {
			return getRuleContext(CommentContext.class,0);
		}
		public Nested_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_nested_stmt; }
	}

	public final Nested_stmtContext nested_stmt() throws RecognitionException {
		Nested_stmtContext _localctx = new Nested_stmtContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_nested_stmt);
		try {
			setState(197);
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
				setState(190);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
				case 1:
					{
					setState(186);
					assignment();
					}
					break;
				case 2:
					{
					setState(187);
					specialFunction();
					}
					break;
				case 3:
					{
					setState(188);
					functionCall();
					}
					break;
				case 4:
					{
					setState(189);
					returnStmt();
					}
					break;
				}
				setState(192);
				match(T__1);
				}
				}
				break;
			case T__18:
				enterOuterAlt(_localctx, 2);
				{
				setState(194);
				conditional();
				}
				break;
			case T__17:
				enterOuterAlt(_localctx, 3);
				{
				setState(195);
				whileLoop();
				}
				break;
			case T__14:
				enterOuterAlt(_localctx, 4);
				{
				setState(196);
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
			setState(199);
			match(T__11);
			setState(200);
			listAssignment();
			setState(205);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__12) {
				{
				{
				setState(201);
				match(T__12);
				setState(202);
				listAssignment();
				}
				}
				setState(207);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(208);
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
		public List<Var_consContext> var_cons() {
			return getRuleContexts(Var_consContext.class);
		}
		public Var_consContext var_cons(int i) {
			return getRuleContext(Var_consContext.class,i);
		}
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
			setState(210);
			match(T__11);
			setState(213);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				{
				setState(211);
				var_cons();
				}
				break;
			case 2:
				{
				setState(212);
				expression();
				}
				break;
			}
			setState(222);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__12) {
				{
				{
				setState(215);
				match(T__12);
				setState(218);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
				case 1:
					{
					setState(216);
					var_cons();
					}
					break;
				case 2:
					{
					setState(217);
					expression();
					}
					break;
				}
				}
				}
				setState(224);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(225);
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
			setState(227);
			match(T__14);
			setState(229); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(228);
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
				setState(231); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__1) | (1L << T__2) | (1L << T__3) | (1L << T__4) | (1L << T__5) | (1L << T__6) | (1L << T__7) | (1L << T__8) | (1L << T__9) | (1L << T__10) | (1L << T__11) | (1L << T__12) | (1L << T__13) | (1L << T__15) | (1L << T__16) | (1L << T__17) | (1L << T__18) | (1L << T__19) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << T__42) | (1L << T__43) | (1L << WS) | (1L << CARDINALITY) | (1L << TYPE_PRIMITIVE) | (1L << ID) | (1L << INT_CONS) | (1L << NUMBER) | (1L << BOOL_CONS) | (1L << STRING_CONS) | (1L << FLOAT_CONS) | (1L << OP_SEC) | (1L << OP_FIRST) | (1L << LOGIC_OPERATOR))) != 0) );
			setState(233);
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
		public TerminalNode ID() { return getToken(BistatParser.ID, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public List<Nested_stmtContext> nested_stmt() {
			return getRuleContexts(Nested_stmtContext.class);
		}
		public Nested_stmtContext nested_stmt(int i) {
			return getRuleContext(Nested_stmtContext.class,i);
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
			setState(235);
			match(T__15);
			setState(236);
			match(T__4);
			setState(237);
			match(ID);
			setState(238);
			match(T__16);
			setState(239);
			expression();
			setState(240);
			match(T__5);
			setState(241);
			match(T__7);
			setState(243); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(242);
				nested_stmt();
				}
				}
				setState(245); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__14) | (1L << T__17) | (1L << T__18) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << T__42) | (1L << ID))) != 0) );
			setState(247);
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

	public static class WhileLoopContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
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
		enterRule(_localctx, 28, RULE_whileLoop);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(249);
			match(T__17);
			setState(250);
			match(T__4);
			setState(251);
			expression();
			setState(252);
			match(T__5);
			setState(253);
			match(T__7);
			setState(255); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(254);
				stmt();
				}
				}
				setState(257); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__14) | (1L << T__15) | (1L << T__17) | (1L << T__18) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << T__42) | (1L << ID))) != 0) );
			setState(259);
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

	public static class ConditionalContext extends ParserRuleContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public ConditionalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_conditional; }
	}

	public final ConditionalContext conditional() throws RecognitionException {
		ConditionalContext _localctx = new ConditionalContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_conditional);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(261);
			match(T__18);
			setState(262);
			match(T__4);
			setState(263);
			expression();
			setState(264);
			match(T__5);
			setState(265);
			match(T__7);
			setState(267); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(266);
				stmt();
				}
				}
				setState(269); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__14) | (1L << T__15) | (1L << T__17) | (1L << T__18) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << T__42) | (1L << ID))) != 0) );
			setState(271);
			match(T__8);
			setState(287);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,21,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(272);
					match(T__19);
					setState(273);
					match(T__18);
					setState(274);
					match(T__4);
					setState(275);
					expression();
					setState(276);
					match(T__5);
					setState(277);
					match(T__7);
					setState(279); 
					_errHandler.sync(this);
					_la = _input.LA(1);
					do {
						{
						{
						setState(278);
						stmt();
						}
						}
						setState(281); 
						_errHandler.sync(this);
						_la = _input.LA(1);
					} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__14) | (1L << T__15) | (1L << T__17) | (1L << T__18) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << T__42) | (1L << ID))) != 0) );
					setState(283);
					match(T__8);
					}
					} 
				}
				setState(289);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,21,_ctx);
			}
			setState(299);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__19) {
				{
				setState(290);
				match(T__19);
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
		enterRule(_localctx, 32, RULE_returnStmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(301);
			match(T__20);
			setState(302);
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
		enterRule(_localctx, 34, RULE_specialFunction);
		try {
			setState(326);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__21:
				enterOuterAlt(_localctx, 1);
				{
				setState(304);
				inputRead();
				}
				break;
			case T__22:
				enterOuterAlt(_localctx, 2);
				{
				setState(305);
				print();
				}
				break;
			case T__23:
				enterOuterAlt(_localctx, 3);
				{
				setState(306);
				listAdd();
				}
				break;
			case T__24:
				enterOuterAlt(_localctx, 4);
				{
				setState(307);
				listPop();
				}
				break;
			case T__25:
				enterOuterAlt(_localctx, 5);
				{
				setState(308);
				length();
				}
				break;
			case T__26:
				enterOuterAlt(_localctx, 6);
				{
				setState(309);
				range();
				}
				break;
			case T__27:
				enterOuterAlt(_localctx, 7);
				{
				setState(310);
				plot();
				}
				break;
			case T__28:
				enterOuterAlt(_localctx, 8);
				{
				setState(311);
				sum();
				}
				break;
			case T__29:
				enterOuterAlt(_localctx, 9);
				{
				setState(312);
				min();
				}
				break;
			case T__30:
				enterOuterAlt(_localctx, 10);
				{
				setState(313);
				prod();
				}
				break;
			case T__31:
				enterOuterAlt(_localctx, 11);
				{
				setState(314);
				avg();
				}
				break;
			case T__32:
				enterOuterAlt(_localctx, 12);
				{
				setState(315);
				sMode();
				}
				break;
			case T__33:
				enterOuterAlt(_localctx, 13);
				{
				setState(316);
				median();
				}
				break;
			case T__34:
				enterOuterAlt(_localctx, 14);
				{
				setState(317);
				sin();
				}
				break;
			case T__36:
				enterOuterAlt(_localctx, 15);
				{
				setState(318);
				cos();
				}
				break;
			case T__35:
				enterOuterAlt(_localctx, 16);
				{
				setState(319);
				tan();
				}
				break;
			case T__37:
				enterOuterAlt(_localctx, 17);
				{
				setState(320);
				sort();
				}
				break;
			case T__38:
				enterOuterAlt(_localctx, 18);
				{
				setState(321);
				sqrt();
				}
				break;
			case T__39:
				enterOuterAlt(_localctx, 19);
				{
				setState(322);
				floor();
				}
				break;
			case T__40:
				enterOuterAlt(_localctx, 20);
				{
				setState(323);
				ceil();
				}
				break;
			case T__41:
				enterOuterAlt(_localctx, 21);
				{
				setState(324);
				abs();
				}
				break;
			case T__42:
				enterOuterAlt(_localctx, 22);
				{
				setState(325);
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
		enterRule(_localctx, 36, RULE_inputRead);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(328);
			match(T__21);
			setState(329);
			match(T__4);
			setState(330);
			match(ID);
			setState(335);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__12) {
				{
				{
				setState(331);
				match(T__12);
				setState(332);
				match(ID);
				}
				}
				setState(337);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(338);
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
		enterRule(_localctx, 38, RULE_print);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(340);
			match(T__22);
			setState(341);
			match(T__4);
			setState(342);
			expression();
			setState(347);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__12) {
				{
				{
				setState(343);
				match(T__12);
				setState(344);
				expression();
				}
				}
				setState(349);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(350);
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
		enterRule(_localctx, 40, RULE_listAdd);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(352);
			match(T__23);
			setState(353);
			match(T__4);
			setState(354);
			expression();
			setState(355);
			match(T__12);
			setState(356);
			expression();
			setState(357);
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
		enterRule(_localctx, 42, RULE_listPop);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(359);
			match(T__24);
			setState(360);
			match(T__4);
			setState(361);
			expression();
			setState(362);
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
		enterRule(_localctx, 44, RULE_length);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(364);
			match(T__25);
			setState(365);
			match(T__4);
			setState(366);
			expression();
			setState(367);
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
		enterRule(_localctx, 46, RULE_range);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(369);
			match(T__26);
			setState(370);
			match(T__4);
			setState(371);
			expression();
			setState(374);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__12) {
				{
				setState(372);
				match(T__12);
				setState(373);
				expression();
				}
			}

			setState(376);
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
		enterRule(_localctx, 48, RULE_plot);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(378);
			match(T__27);
			setState(379);
			match(T__4);
			setState(380);
			expression();
			setState(381);
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
		enterRule(_localctx, 50, RULE_sum);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(383);
			match(T__28);
			setState(384);
			match(T__4);
			setState(385);
			expression();
			setState(386);
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
		enterRule(_localctx, 52, RULE_min);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(388);
			match(T__29);
			setState(389);
			match(T__4);
			setState(390);
			expression();
			setState(391);
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
		enterRule(_localctx, 54, RULE_prod);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(393);
			match(T__30);
			setState(394);
			match(T__4);
			setState(395);
			expression();
			setState(396);
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
		enterRule(_localctx, 56, RULE_avg);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(398);
			match(T__31);
			setState(399);
			match(T__4);
			setState(400);
			expression();
			setState(401);
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
		enterRule(_localctx, 58, RULE_sMode);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(403);
			match(T__32);
			setState(404);
			match(T__4);
			setState(405);
			expression();
			setState(406);
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
		enterRule(_localctx, 60, RULE_median);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(408);
			match(T__33);
			setState(409);
			match(T__4);
			setState(410);
			expression();
			setState(411);
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
		enterRule(_localctx, 62, RULE_sin);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(413);
			match(T__34);
			setState(414);
			match(T__4);
			setState(415);
			expression();
			setState(416);
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
		enterRule(_localctx, 64, RULE_tan);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(418);
			match(T__35);
			setState(419);
			match(T__4);
			setState(420);
			expression();
			setState(421);
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
		enterRule(_localctx, 66, RULE_cos);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(423);
			match(T__36);
			setState(424);
			match(T__4);
			setState(425);
			expression();
			setState(426);
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
		enterRule(_localctx, 68, RULE_sort);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(428);
			match(T__37);
			setState(429);
			match(T__4);
			setState(430);
			expression();
			setState(431);
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
		enterRule(_localctx, 70, RULE_sqrt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(433);
			match(T__38);
			setState(434);
			match(T__4);
			setState(435);
			expression();
			setState(436);
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
		enterRule(_localctx, 72, RULE_floor);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(438);
			match(T__39);
			setState(439);
			match(T__4);
			setState(440);
			expression();
			setState(441);
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
		enterRule(_localctx, 74, RULE_ceil);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(443);
			match(T__40);
			setState(444);
			match(T__4);
			setState(445);
			expression();
			setState(446);
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
		enterRule(_localctx, 76, RULE_abs);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(448);
			match(T__41);
			setState(449);
			match(T__4);
			setState(450);
			expression();
			setState(451);
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
		enterRule(_localctx, 78, RULE_not);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(453);
			match(T__42);
			setState(454);
			match(T__4);
			setState(455);
			expression();
			setState(456);
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
		public TerminalNode LOGIC_OPERATOR() { return getToken(BistatParser.LOGIC_OPERATOR, 0); }
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		ExpressionContext _localctx = new ExpressionContext(_ctx, getState());
		enterRule(_localctx, 80, RULE_expression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(458);
			exp();
			setState(461);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LOGIC_OPERATOR) {
				{
				setState(459);
				match(LOGIC_OPERATOR);
				setState(460);
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
		public TermContext term() {
			return getRuleContext(TermContext.class,0);
		}
		public TerminalNode OP_SEC() { return getToken(BistatParser.OP_SEC, 0); }
		public ExpContext exp() {
			return getRuleContext(ExpContext.class,0);
		}
		public ExpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_exp; }
	}

	public final ExpContext exp() throws RecognitionException {
		ExpContext _localctx = new ExpContext(_ctx, getState());
		enterRule(_localctx, 82, RULE_exp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(463);
			term();
			setState(466);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==OP_SEC) {
				{
				setState(464);
				match(OP_SEC);
				setState(465);
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

	public static class TermContext extends ParserRuleContext {
		public FactorContext factor() {
			return getRuleContext(FactorContext.class,0);
		}
		public TerminalNode OP_FIRST() { return getToken(BistatParser.OP_FIRST, 0); }
		public TermContext term() {
			return getRuleContext(TermContext.class,0);
		}
		public TermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_term; }
	}

	public final TermContext term() throws RecognitionException {
		TermContext _localctx = new TermContext(_ctx, getState());
		enterRule(_localctx, 84, RULE_term);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(468);
			factor();
			setState(471);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==OP_FIRST) {
				{
				setState(469);
				match(OP_FIRST);
				setState(470);
				term();
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

	public static class FactorContext extends ParserRuleContext {
		public IndexingContext indexing() {
			return getRuleContext(IndexingContext.class,0);
		}
		public SpecialFunctionContext specialFunction() {
			return getRuleContext(SpecialFunctionContext.class,0);
		}
		public FunctionCallContext functionCall() {
			return getRuleContext(FunctionCallContext.class,0);
		}
		public Var_consContext var_cons() {
			return getRuleContext(Var_consContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public FactorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_factor; }
	}

	public final FactorContext factor() throws RecognitionException {
		FactorContext _localctx = new FactorContext(_ctx, getState());
		enterRule(_localctx, 86, RULE_factor);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(474);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__43) {
				{
				setState(473);
				match(T__43);
				}
			}

			setState(484);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,32,_ctx) ) {
			case 1:
				{
				{
				setState(476);
				match(T__4);
				setState(477);
				expression();
				setState(478);
				match(T__5);
				}
				}
				break;
			case 2:
				{
				setState(480);
				indexing();
				}
				break;
			case 3:
				{
				setState(481);
				specialFunction();
				}
				break;
			case 4:
				{
				setState(482);
				functionCall();
				}
				break;
			case 5:
				{
				setState(483);
				var_cons();
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
		enterRule(_localctx, 88, RULE_functionCall);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(486);
			match(ID);
			setState(487);
			match(T__4);
			setState(496);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__4) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << T__42) | (1L << T__43) | (1L << ID) | (1L << INT_CONS) | (1L << BOOL_CONS) | (1L << STRING_CONS) | (1L << FLOAT_CONS))) != 0)) {
				{
				setState(488);
				expression();
				setState(493);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__12) {
					{
					{
					setState(489);
					match(T__12);
					setState(490);
					expression();
					}
					}
					setState(495);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(498);
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
		enterRule(_localctx, 90, RULE_indexing);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(500);
			match(ID);
			setState(501);
			match(T__11);
			setState(502);
			expression();
			setState(503);
			match(T__13);
			setState(508);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__11) {
				{
				setState(504);
				match(T__11);
				setState(505);
				expression();
				setState(506);
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

	public static class Var_consContext extends ParserRuleContext {
		public TerminalNode STRING_CONS() { return getToken(BistatParser.STRING_CONS, 0); }
		public TerminalNode FLOAT_CONS() { return getToken(BistatParser.FLOAT_CONS, 0); }
		public TerminalNode INT_CONS() { return getToken(BistatParser.INT_CONS, 0); }
		public TerminalNode BOOL_CONS() { return getToken(BistatParser.BOOL_CONS, 0); }
		public TerminalNode ID() { return getToken(BistatParser.ID, 0); }
		public Var_consContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_var_cons; }
	}

	public final Var_consContext var_cons() throws RecognitionException {
		Var_consContext _localctx = new Var_consContext(_ctx, getState());
		enterRule(_localctx, 92, RULE_var_cons);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(510);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << ID) | (1L << INT_CONS) | (1L << BOOL_CONS) | (1L << STRING_CONS) | (1L << FLOAT_CONS))) != 0)) ) {
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3:\u0203\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\3\2\3\2\3\2\3\2\7\2e\n\2\f\2\16\2h\13"+
		"\2\3\2\7\2k\n\2\f\2\16\2n\13\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\4\3\4"+
		"\5\4z\n\4\3\5\3\5\3\5\3\5\7\5\u0080\n\5\f\5\16\5\u0083\13\5\3\5\3\5\3"+
		"\5\3\5\7\5\u0089\n\5\f\5\16\5\u008c\13\5\3\5\3\5\6\5\u0090\n\5\r\5\16"+
		"\5\u0091\3\5\3\5\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\6\b\u00a2"+
		"\n\b\r\b\16\b\u00a3\3\b\3\b\3\t\3\t\3\t\3\t\5\t\u00ac\n\t\3\t\3\t\3\t"+
		"\3\t\3\t\3\t\5\t\u00b4\n\t\3\n\3\n\3\n\3\n\3\n\5\n\u00bb\n\n\3\13\3\13"+
		"\3\13\3\13\5\13\u00c1\n\13\3\13\3\13\3\13\3\13\3\13\5\13\u00c8\n\13\3"+
		"\f\3\f\3\f\3\f\7\f\u00ce\n\f\f\f\16\f\u00d1\13\f\3\f\3\f\3\r\3\r\3\r\5"+
		"\r\u00d8\n\r\3\r\3\r\3\r\5\r\u00dd\n\r\7\r\u00df\n\r\f\r\16\r\u00e2\13"+
		"\r\3\r\3\r\3\16\3\16\6\16\u00e8\n\16\r\16\16\16\u00e9\3\16\3\16\3\17\3"+
		"\17\3\17\3\17\3\17\3\17\3\17\3\17\6\17\u00f6\n\17\r\17\16\17\u00f7\3\17"+
		"\3\17\3\20\3\20\3\20\3\20\3\20\3\20\6\20\u0102\n\20\r\20\16\20\u0103\3"+
		"\20\3\20\3\21\3\21\3\21\3\21\3\21\3\21\6\21\u010e\n\21\r\21\16\21\u010f"+
		"\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\6\21\u011a\n\21\r\21\16\21\u011b"+
		"\3\21\3\21\7\21\u0120\n\21\f\21\16\21\u0123\13\21\3\21\3\21\3\21\6\21"+
		"\u0128\n\21\r\21\16\21\u0129\3\21\3\21\5\21\u012e\n\21\3\22\3\22\3\22"+
		"\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23"+
		"\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\5\23\u0149\n\23\3\24\3\24\3\24"+
		"\3\24\3\24\7\24\u0150\n\24\f\24\16\24\u0153\13\24\3\24\3\24\3\25\3\25"+
		"\3\25\3\25\3\25\7\25\u015c\n\25\f\25\16\25\u015f\13\25\3\25\3\25\3\26"+
		"\3\26\3\26\3\26\3\26\3\26\3\26\3\27\3\27\3\27\3\27\3\27\3\30\3\30\3\30"+
		"\3\30\3\30\3\31\3\31\3\31\3\31\3\31\5\31\u0179\n\31\3\31\3\31\3\32\3\32"+
		"\3\32\3\32\3\32\3\33\3\33\3\33\3\33\3\33\3\34\3\34\3\34\3\34\3\34\3\35"+
		"\3\35\3\35\3\35\3\35\3\36\3\36\3\36\3\36\3\36\3\37\3\37\3\37\3\37\3\37"+
		"\3 \3 \3 \3 \3 \3!\3!\3!\3!\3!\3\"\3\"\3\"\3\"\3\"\3#\3#\3#\3#\3#\3$\3"+
		"$\3$\3$\3$\3%\3%\3%\3%\3%\3&\3&\3&\3&\3&\3\'\3\'\3\'\3\'\3\'\3(\3(\3("+
		"\3(\3(\3)\3)\3)\3)\3)\3*\3*\3*\5*\u01d0\n*\3+\3+\3+\5+\u01d5\n+\3,\3,"+
		"\3,\5,\u01da\n,\3-\5-\u01dd\n-\3-\3-\3-\3-\3-\3-\3-\3-\5-\u01e7\n-\3."+
		"\3.\3.\3.\3.\7.\u01ee\n.\f.\16.\u01f1\13.\5.\u01f3\n.\3.\3.\3/\3/\3/\3"+
		"/\3/\3/\3/\3/\5/\u01ff\n/\3\60\3\60\3\60\2\2\61\2\4\6\b\n\f\16\20\22\24"+
		"\26\30\32\34\36 \"$&(*,.\60\62\64\668:<>@BDFHJLNPRTVXZ\\^\2\4\3\2\21\21"+
		"\4\2\62\63\65\67\2\u0218\2`\3\2\2\2\4r\3\2\2\2\6w\3\2\2\2\b{\3\2\2\2\n"+
		"\u0095\3\2\2\2\f\u0097\3\2\2\2\16\u009c\3\2\2\2\20\u00b3\3\2\2\2\22\u00b5"+
		"\3\2\2\2\24\u00c7\3\2\2\2\26\u00c9\3\2\2\2\30\u00d4\3\2\2\2\32\u00e5\3"+
		"\2\2\2\34\u00ed\3\2\2\2\36\u00fb\3\2\2\2 \u0107\3\2\2\2\"\u012f\3\2\2"+
		"\2$\u0148\3\2\2\2&\u014a\3\2\2\2(\u0156\3\2\2\2*\u0162\3\2\2\2,\u0169"+
		"\3\2\2\2.\u016e\3\2\2\2\60\u0173\3\2\2\2\62\u017c\3\2\2\2\64\u0181\3\2"+
		"\2\2\66\u0186\3\2\2\28\u018b\3\2\2\2:\u0190\3\2\2\2<\u0195\3\2\2\2>\u019a"+
		"\3\2\2\2@\u019f\3\2\2\2B\u01a4\3\2\2\2D\u01a9\3\2\2\2F\u01ae\3\2\2\2H"+
		"\u01b3\3\2\2\2J\u01b8\3\2\2\2L\u01bd\3\2\2\2N\u01c2\3\2\2\2P\u01c7\3\2"+
		"\2\2R\u01cc\3\2\2\2T\u01d1\3\2\2\2V\u01d6\3\2\2\2X\u01dc\3\2\2\2Z\u01e8"+
		"\3\2\2\2\\\u01f6\3\2\2\2^\u0200\3\2\2\2`a\7\3\2\2ab\7\62\2\2bf\7\4\2\2"+
		"ce\5\4\3\2dc\3\2\2\2eh\3\2\2\2fd\3\2\2\2fg\3\2\2\2gl\3\2\2\2hf\3\2\2\2"+
		"ik\5\b\5\2ji\3\2\2\2kn\3\2\2\2lj\3\2\2\2lm\3\2\2\2mo\3\2\2\2nl\3\2\2\2"+
		"op\5\16\b\2pq\7\2\2\3q\3\3\2\2\2rs\7\5\2\2st\5\6\4\2tu\7\62\2\2uv\7\4"+
		"\2\2v\5\3\2\2\2wy\7\61\2\2xz\7\60\2\2yx\3\2\2\2yz\3\2\2\2z\7\3\2\2\2{"+
		"|\7\6\2\2|}\7\62\2\2}\u0081\7\7\2\2~\u0080\5\f\7\2\177~\3\2\2\2\u0080"+
		"\u0083\3\2\2\2\u0081\177\3\2\2\2\u0081\u0082\3\2\2\2\u0082\u0084\3\2\2"+
		"\2\u0083\u0081\3\2\2\2\u0084\u0085\7\b\2\2\u0085\u0086\7\t\2\2\u0086\u008a"+
		"\5\6\4\2\u0087\u0089\5\4\3\2\u0088\u0087\3\2\2\2\u0089\u008c\3\2\2\2\u008a"+
		"\u0088\3\2\2\2\u008a\u008b\3\2\2\2\u008b\u008d\3\2\2\2\u008c\u008a\3\2"+
		"\2\2\u008d\u008f\7\n\2\2\u008e\u0090\5\20\t\2\u008f\u008e\3\2\2\2\u0090"+
		"\u0091\3\2\2\2\u0091\u008f\3\2\2\2\u0091\u0092\3\2\2\2\u0092\u0093\3\2"+
		"\2\2\u0093\u0094\5\n\6\2\u0094\t\3\2\2\2\u0095\u0096\7\13\2\2\u0096\13"+
		"\3\2\2\2\u0097\u0098\7\5\2\2\u0098\u0099\5\6\4\2\u0099\u009a\7\62\2\2"+
		"\u009a\u009b\7\4\2\2\u009b\r\3\2\2\2\u009c\u009d\7\f\2\2\u009d\u009e\7"+
		"\7\2\2\u009e\u009f\7\b\2\2\u009f\u00a1\7\n\2\2\u00a0\u00a2\5\20\t\2\u00a1"+
		"\u00a0\3\2\2\2\u00a2\u00a3\3\2\2\2\u00a3\u00a1\3\2\2\2\u00a3\u00a4\3\2"+
		"\2\2\u00a4\u00a5\3\2\2\2\u00a5\u00a6\7\13\2\2\u00a6\17\3\2\2\2\u00a7\u00ac"+
		"\5\22\n\2\u00a8\u00ac\5$\23\2\u00a9\u00ac\5Z.\2\u00aa\u00ac\5\"\22\2\u00ab"+
		"\u00a7\3\2\2\2\u00ab\u00a8\3\2\2\2\u00ab\u00a9\3\2\2\2\u00ab\u00aa\3\2"+
		"\2\2\u00ac\u00ad\3\2\2\2\u00ad\u00ae\7\4\2\2\u00ae\u00b4\3\2\2\2\u00af"+
		"\u00b4\5 \21\2\u00b0\u00b4\5\36\20\2\u00b1\u00b4\5\34\17\2\u00b2\u00b4"+
		"\5\32\16\2\u00b3\u00ab\3\2\2\2\u00b3\u00af\3\2\2\2\u00b3\u00b0\3\2\2\2"+
		"\u00b3\u00b1\3\2\2\2\u00b3\u00b2\3\2\2\2\u00b4\21\3\2\2\2\u00b5\u00b6"+
		"\7\62\2\2\u00b6\u00ba\7\r\2\2\u00b7\u00bb\5R*\2\u00b8\u00bb\5\30\r\2\u00b9"+
		"\u00bb\5\26\f\2\u00ba\u00b7\3\2\2\2\u00ba\u00b8\3\2\2\2\u00ba\u00b9\3"+
		"\2\2\2\u00bb\23\3\2\2\2\u00bc\u00c1\5\22\n\2\u00bd\u00c1\5$\23\2\u00be"+
		"\u00c1\5Z.\2\u00bf\u00c1\5\"\22\2\u00c0\u00bc\3\2\2\2\u00c0\u00bd\3\2"+
		"\2\2\u00c0\u00be\3\2\2\2\u00c0\u00bf\3\2\2\2\u00c1\u00c2\3\2\2\2\u00c2"+
		"\u00c3\7\4\2\2\u00c3\u00c8\3\2\2\2\u00c4\u00c8\5 \21\2\u00c5\u00c8\5\36"+
		"\20\2\u00c6\u00c8\5\32\16\2\u00c7\u00c0\3\2\2\2\u00c7\u00c4\3\2\2\2\u00c7"+
		"\u00c5\3\2\2\2\u00c7\u00c6\3\2\2\2\u00c8\25\3\2\2\2\u00c9\u00ca\7\16\2"+
		"\2\u00ca\u00cf\5\30\r\2\u00cb\u00cc\7\17\2\2\u00cc\u00ce\5\30\r\2\u00cd"+
		"\u00cb\3\2\2\2\u00ce\u00d1\3\2\2\2\u00cf\u00cd\3\2\2\2\u00cf\u00d0\3\2"+
		"\2\2\u00d0\u00d2\3\2\2\2\u00d1\u00cf\3\2\2\2\u00d2\u00d3\7\20\2\2\u00d3"+
		"\27\3\2\2\2\u00d4\u00d7\7\16\2\2\u00d5\u00d8\5^\60\2\u00d6\u00d8\5R*\2"+
		"\u00d7\u00d5\3\2\2\2\u00d7\u00d6\3\2\2\2\u00d8\u00e0\3\2\2\2\u00d9\u00dc"+
		"\7\17\2\2\u00da\u00dd\5^\60\2\u00db\u00dd\5R*\2\u00dc\u00da\3\2\2\2\u00dc"+
		"\u00db\3\2\2\2\u00dd\u00df\3\2\2\2\u00de\u00d9\3\2\2\2\u00df\u00e2\3\2"+
		"\2\2\u00e0\u00de\3\2\2\2\u00e0\u00e1\3\2\2\2\u00e1\u00e3\3\2\2\2\u00e2"+
		"\u00e0\3\2\2\2\u00e3\u00e4\7\20\2\2\u00e4\31\3\2\2\2\u00e5\u00e7\7\21"+
		"\2\2\u00e6\u00e8\n\2\2\2\u00e7\u00e6\3\2\2\2\u00e8\u00e9\3\2\2\2\u00e9"+
		"\u00e7\3\2\2\2\u00e9\u00ea\3\2\2\2\u00ea\u00eb\3\2\2\2\u00eb\u00ec\7\21"+
		"\2\2\u00ec\33\3\2\2\2\u00ed\u00ee\7\22\2\2\u00ee\u00ef\7\7\2\2\u00ef\u00f0"+
		"\7\62\2\2\u00f0\u00f1\7\23\2\2\u00f1\u00f2\5R*\2\u00f2\u00f3\7\b\2\2\u00f3"+
		"\u00f5\7\n\2\2\u00f4\u00f6\5\24\13\2\u00f5\u00f4\3\2\2\2\u00f6\u00f7\3"+
		"\2\2\2\u00f7\u00f5\3\2\2\2\u00f7\u00f8\3\2\2\2\u00f8\u00f9\3\2\2\2\u00f9"+
		"\u00fa\7\13\2\2\u00fa\35\3\2\2\2\u00fb\u00fc\7\24\2\2\u00fc\u00fd\7\7"+
		"\2\2\u00fd\u00fe\5R*\2\u00fe\u00ff\7\b\2\2\u00ff\u0101\7\n\2\2\u0100\u0102"+
		"\5\20\t\2\u0101\u0100\3\2\2\2\u0102\u0103\3\2\2\2\u0103\u0101\3\2\2\2"+
		"\u0103\u0104\3\2\2\2\u0104\u0105\3\2\2\2\u0105\u0106\7\13\2\2\u0106\37"+
		"\3\2\2\2\u0107\u0108\7\25\2\2\u0108\u0109\7\7\2\2\u0109\u010a\5R*\2\u010a"+
		"\u010b\7\b\2\2\u010b\u010d\7\n\2\2\u010c\u010e\5\20\t\2\u010d\u010c\3"+
		"\2\2\2\u010e\u010f\3\2\2\2\u010f\u010d\3\2\2\2\u010f\u0110\3\2\2\2\u0110"+
		"\u0111\3\2\2\2\u0111\u0121\7\13\2\2\u0112\u0113\7\26\2\2\u0113\u0114\7"+
		"\25\2\2\u0114\u0115\7\7\2\2\u0115\u0116\5R*\2\u0116\u0117\7\b\2\2\u0117"+
		"\u0119\7\n\2\2\u0118\u011a\5\20\t\2\u0119\u0118\3\2\2\2\u011a\u011b\3"+
		"\2\2\2\u011b\u0119\3\2\2\2\u011b\u011c\3\2\2\2\u011c\u011d\3\2\2\2\u011d"+
		"\u011e\7\13\2\2\u011e\u0120\3\2\2\2\u011f\u0112\3\2\2\2\u0120\u0123\3"+
		"\2\2\2\u0121\u011f\3\2\2\2\u0121\u0122\3\2\2\2\u0122\u012d\3\2\2\2\u0123"+
		"\u0121\3\2\2\2\u0124\u0125\7\26\2\2\u0125\u0127\7\n\2\2\u0126\u0128\5"+
		"\20\t\2\u0127\u0126\3\2\2\2\u0128\u0129\3\2\2\2\u0129\u0127\3\2\2\2\u0129"+
		"\u012a\3\2\2\2\u012a\u012b\3\2\2\2\u012b\u012c\7\13\2\2\u012c\u012e\3"+
		"\2\2\2\u012d\u0124\3\2\2\2\u012d\u012e\3\2\2\2\u012e!\3\2\2\2\u012f\u0130"+
		"\7\27\2\2\u0130\u0131\5R*\2\u0131#\3\2\2\2\u0132\u0149\5&\24\2\u0133\u0149"+
		"\5(\25\2\u0134\u0149\5*\26\2\u0135\u0149\5,\27\2\u0136\u0149\5.\30\2\u0137"+
		"\u0149\5\60\31\2\u0138\u0149\5\62\32\2\u0139\u0149\5\64\33\2\u013a\u0149"+
		"\5\66\34\2\u013b\u0149\58\35\2\u013c\u0149\5:\36\2\u013d\u0149\5<\37\2"+
		"\u013e\u0149\5> \2\u013f\u0149\5@!\2\u0140\u0149\5D#\2\u0141\u0149\5B"+
		"\"\2\u0142\u0149\5F$\2\u0143\u0149\5H%\2\u0144\u0149\5J&\2\u0145\u0149"+
		"\5L\'\2\u0146\u0149\5N(\2\u0147\u0149\5P)\2\u0148\u0132\3\2\2\2\u0148"+
		"\u0133\3\2\2\2\u0148\u0134\3\2\2\2\u0148\u0135\3\2\2\2\u0148\u0136\3\2"+
		"\2\2\u0148\u0137\3\2\2\2\u0148\u0138\3\2\2\2\u0148\u0139\3\2\2\2\u0148"+
		"\u013a\3\2\2\2\u0148\u013b\3\2\2\2\u0148\u013c\3\2\2\2\u0148\u013d\3\2"+
		"\2\2\u0148\u013e\3\2\2\2\u0148\u013f\3\2\2\2\u0148\u0140\3\2\2\2\u0148"+
		"\u0141\3\2\2\2\u0148\u0142\3\2\2\2\u0148\u0143\3\2\2\2\u0148\u0144\3\2"+
		"\2\2\u0148\u0145\3\2\2\2\u0148\u0146\3\2\2\2\u0148\u0147\3\2\2\2\u0149"+
		"%\3\2\2\2\u014a\u014b\7\30\2\2\u014b\u014c\7\7\2\2\u014c\u0151\7\62\2"+
		"\2\u014d\u014e\7\17\2\2\u014e\u0150\7\62\2\2\u014f\u014d\3\2\2\2\u0150"+
		"\u0153\3\2\2\2\u0151\u014f\3\2\2\2\u0151\u0152\3\2\2\2\u0152\u0154\3\2"+
		"\2\2\u0153\u0151\3\2\2\2\u0154\u0155\7\b\2\2\u0155\'\3\2\2\2\u0156\u0157"+
		"\7\31\2\2\u0157\u0158\7\7\2\2\u0158\u015d\5R*\2\u0159\u015a\7\17\2\2\u015a"+
		"\u015c\5R*\2\u015b\u0159\3\2\2\2\u015c\u015f\3\2\2\2\u015d\u015b\3\2\2"+
		"\2\u015d\u015e\3\2\2\2\u015e\u0160\3\2\2\2\u015f\u015d\3\2\2\2\u0160\u0161"+
		"\7\b\2\2\u0161)\3\2\2\2\u0162\u0163\7\32\2\2\u0163\u0164\7\7\2\2\u0164"+
		"\u0165\5R*\2\u0165\u0166\7\17\2\2\u0166\u0167\5R*\2\u0167\u0168\7\b\2"+
		"\2\u0168+\3\2\2\2\u0169\u016a\7\33\2\2\u016a\u016b\7\7\2\2\u016b\u016c"+
		"\5R*\2\u016c\u016d\7\b\2\2\u016d-\3\2\2\2\u016e\u016f\7\34\2\2\u016f\u0170"+
		"\7\7\2\2\u0170\u0171\5R*\2\u0171\u0172\7\b\2\2\u0172/\3\2\2\2\u0173\u0174"+
		"\7\35\2\2\u0174\u0175\7\7\2\2\u0175\u0178\5R*\2\u0176\u0177\7\17\2\2\u0177"+
		"\u0179\5R*\2\u0178\u0176\3\2\2\2\u0178\u0179\3\2\2\2\u0179\u017a\3\2\2"+
		"\2\u017a\u017b\7\b\2\2\u017b\61\3\2\2\2\u017c\u017d\7\36\2\2\u017d\u017e"+
		"\7\7\2\2\u017e\u017f\5R*\2\u017f\u0180\7\b\2\2\u0180\63\3\2\2\2\u0181"+
		"\u0182\7\37\2\2\u0182\u0183\7\7\2\2\u0183\u0184\5R*\2\u0184\u0185\7\b"+
		"\2\2\u0185\65\3\2\2\2\u0186\u0187\7 \2\2\u0187\u0188\7\7\2\2\u0188\u0189"+
		"\5R*\2\u0189\u018a\7\b\2\2\u018a\67\3\2\2\2\u018b\u018c\7!\2\2\u018c\u018d"+
		"\7\7\2\2\u018d\u018e\5R*\2\u018e\u018f\7\b\2\2\u018f9\3\2\2\2\u0190\u0191"+
		"\7\"\2\2\u0191\u0192\7\7\2\2\u0192\u0193\5R*\2\u0193\u0194\7\b\2\2\u0194"+
		";\3\2\2\2\u0195\u0196\7#\2\2\u0196\u0197\7\7\2\2\u0197\u0198\5R*\2\u0198"+
		"\u0199\7\b\2\2\u0199=\3\2\2\2\u019a\u019b\7$\2\2\u019b\u019c\7\7\2\2\u019c"+
		"\u019d\5R*\2\u019d\u019e\7\b\2\2\u019e?\3\2\2\2\u019f\u01a0\7%\2\2\u01a0"+
		"\u01a1\7\7\2\2\u01a1\u01a2\5R*\2\u01a2\u01a3\7\b\2\2\u01a3A\3\2\2\2\u01a4"+
		"\u01a5\7&\2\2\u01a5\u01a6\7\7\2\2\u01a6\u01a7\5R*\2\u01a7\u01a8\7\b\2"+
		"\2\u01a8C\3\2\2\2\u01a9\u01aa\7\'\2\2\u01aa\u01ab\7\7\2\2\u01ab\u01ac"+
		"\5R*\2\u01ac\u01ad\7\b\2\2\u01adE\3\2\2\2\u01ae\u01af\7(\2\2\u01af\u01b0"+
		"\7\7\2\2\u01b0\u01b1\5R*\2\u01b1\u01b2\7\b\2\2\u01b2G\3\2\2\2\u01b3\u01b4"+
		"\7)\2\2\u01b4\u01b5\7\7\2\2\u01b5\u01b6\5R*\2\u01b6\u01b7\7\b\2\2\u01b7"+
		"I\3\2\2\2\u01b8\u01b9\7*\2\2\u01b9\u01ba\7\7\2\2\u01ba\u01bb\5R*\2\u01bb"+
		"\u01bc\7\b\2\2\u01bcK\3\2\2\2\u01bd\u01be\7+\2\2\u01be\u01bf\7\7\2\2\u01bf"+
		"\u01c0\5R*\2\u01c0\u01c1\7\b\2\2\u01c1M\3\2\2\2\u01c2\u01c3\7,\2\2\u01c3"+
		"\u01c4\7\7\2\2\u01c4\u01c5\5R*\2\u01c5\u01c6\7\b\2\2\u01c6O\3\2\2\2\u01c7"+
		"\u01c8\7-\2\2\u01c8\u01c9\7\7\2\2\u01c9\u01ca\5R*\2\u01ca\u01cb\7\b\2"+
		"\2\u01cbQ\3\2\2\2\u01cc\u01cf\5T+\2\u01cd\u01ce\7:\2\2\u01ce\u01d0\5T"+
		"+\2\u01cf\u01cd\3\2\2\2\u01cf\u01d0\3\2\2\2\u01d0S\3\2\2\2\u01d1\u01d4"+
		"\5V,\2\u01d2\u01d3\78\2\2\u01d3\u01d5\5T+\2\u01d4\u01d2\3\2\2\2\u01d4"+
		"\u01d5\3\2\2\2\u01d5U\3\2\2\2\u01d6\u01d9\5X-\2\u01d7\u01d8\79\2\2\u01d8"+
		"\u01da\5V,\2\u01d9\u01d7\3\2\2\2\u01d9\u01da\3\2\2\2\u01daW\3\2\2\2\u01db"+
		"\u01dd\7.\2\2\u01dc\u01db\3\2\2\2\u01dc\u01dd\3\2\2\2\u01dd\u01e6\3\2"+
		"\2\2\u01de\u01df\7\7\2\2\u01df\u01e0\5R*\2\u01e0\u01e1\7\b\2\2\u01e1\u01e7"+
		"\3\2\2\2\u01e2\u01e7\5\\/\2\u01e3\u01e7\5$\23\2\u01e4\u01e7\5Z.\2\u01e5"+
		"\u01e7\5^\60\2\u01e6\u01de\3\2\2\2\u01e6\u01e2\3\2\2\2\u01e6\u01e3\3\2"+
		"\2\2\u01e6\u01e4\3\2\2\2\u01e6\u01e5\3\2\2\2\u01e7Y\3\2\2\2\u01e8\u01e9"+
		"\7\62\2\2\u01e9\u01f2\7\7\2\2\u01ea\u01ef\5R*\2\u01eb\u01ec\7\17\2\2\u01ec"+
		"\u01ee\5R*\2\u01ed\u01eb\3\2\2\2\u01ee\u01f1\3\2\2\2\u01ef\u01ed\3\2\2"+
		"\2\u01ef\u01f0\3\2\2\2\u01f0\u01f3\3\2\2\2\u01f1\u01ef\3\2\2\2\u01f2\u01ea"+
		"\3\2\2\2\u01f2\u01f3\3\2\2\2\u01f3\u01f4\3\2\2\2\u01f4\u01f5\7\b\2\2\u01f5"+
		"[\3\2\2\2\u01f6\u01f7\7\62\2\2\u01f7\u01f8\7\16\2\2\u01f8\u01f9\5R*\2"+
		"\u01f9\u01fe\7\20\2\2\u01fa\u01fb\7\16\2\2\u01fb\u01fc\5R*\2\u01fc\u01fd"+
		"\7\20\2\2\u01fd\u01ff\3\2\2\2\u01fe\u01fa\3\2\2\2\u01fe\u01ff\3\2\2\2"+
		"\u01ff]\3\2\2\2\u0200\u0201\t\3\2\2\u0201_\3\2\2\2&fly\u0081\u008a\u0091"+
		"\u00a3\u00ab\u00b3\u00ba\u00c0\u00c7\u00cf\u00d7\u00dc\u00e0\u00e9\u00f7"+
		"\u0103\u010f\u011b\u0121\u0129\u012d\u0148\u0151\u015d\u0178\u01cf\u01d4"+
		"\u01d9\u01dc\u01e6\u01ef\u01f2\u01fe";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}