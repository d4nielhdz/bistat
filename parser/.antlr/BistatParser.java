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
		T__38=39, T__39=40, T__40=41, T__41=42, WS=43, NON_VOID_TYPE=44, RETURN_TYPE=45, 
		PARAM_TYPE=46, ID=47, PARAM_CARDINALITY=48, CARDINALITY=49, VAR_CONS=50, 
		BOOL_CONS=51, STRING_CONS=52, FLOAT_CONS=53, INT_CONS=54, OP_SEC=55, OP_FIRST=56, 
		LOGIC_OPERATOR=57, NUMBER=58;
	public static final int
		RULE_program = 0, RULE_varDeclaration = 1, RULE_funcDef = 2, RULE_main = 3, 
		RULE_stmt = 4, RULE_paramDeclaration = 5, RULE_assignment = 6, RULE_nested_stmt = 7, 
		RULE_matrixAssignment = 8, RULE_listAssignment = 9, RULE_comment = 10, 
		RULE_forLoop = 11, RULE_whileLoop = 12, RULE_conditional = 13, RULE_specialFunction = 14, 
		RULE_inputRead = 15, RULE_print = 16, RULE_listAdd = 17, RULE_listPop = 18, 
		RULE_length = 19, RULE_range = 20, RULE_plot = 21, RULE_sum = 22, RULE_min = 23, 
		RULE_prod = 24, RULE_avg = 25, RULE_sMode = 26, RULE_median = 27, RULE_sin = 28, 
		RULE_tan = 29, RULE_cos = 30, RULE_sort = 31, RULE_sqrt = 32, RULE_floor = 33, 
		RULE_ceil = 34, RULE_abs = 35, RULE_not = 36, RULE_expression = 37, RULE_exp = 38, 
		RULE_term = 39, RULE_factor = 40, RULE_functionCall = 41, RULE_indexing = 42;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "varDeclaration", "funcDef", "main", "stmt", "paramDeclaration", 
			"assignment", "nested_stmt", "matrixAssignment", "listAssignment", "comment", 
			"forLoop", "whileLoop", "conditional", "specialFunction", "inputRead", 
			"print", "listAdd", "listPop", "length", "range", "plot", "sum", "min", 
			"prod", "avg", "sMode", "median", "sin", "tan", "cos", "sort", "sqrt", 
			"floor", "ceil", "abs", "not", "expression", "exp", "term", "factor", 
			"functionCall", "indexing"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'Program'", "';'", "'var'", "'func'", "'('", "')'", "':'", "'{'", 
			"'}'", "'main'", "'='", "'['", "','", "']'", "'#'", "'for'", "'in'", 
			"'while'", "'if'", "'else'", "'read'", "'print'", "'listAdd'", "'listPop'", 
			"'length'", "'range'", "'plot'", "'sum'", "'min'", "'prod'", "'avg'", 
			"'sMode'", "'median'", "'sin'", "'tan'", "'cos'", "'sort'", "'sqrt'", 
			"'floor'", "'ceil'", "'abs'", "'not'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, "WS", "NON_VOID_TYPE", "RETURN_TYPE", 
			"PARAM_TYPE", "ID", "PARAM_CARDINALITY", "CARDINALITY", "VAR_CONS", "BOOL_CONS", 
			"STRING_CONS", "FLOAT_CONS", "INT_CONS", "OP_SEC", "OP_FIRST", "LOGIC_OPERATOR", 
			"NUMBER"
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
			setState(86);
			match(T__0);
			setState(87);
			match(ID);
			setState(88);
			match(T__1);
			setState(92);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__2) {
				{
				{
				setState(89);
				varDeclaration();
				}
				}
				setState(94);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(98);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__3) {
				{
				{
				setState(95);
				funcDef();
				}
				}
				setState(100);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(101);
			main();
			setState(102);
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
		public TerminalNode NON_VOID_TYPE() { return getToken(BistatParser.NON_VOID_TYPE, 0); }
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
			setState(104);
			match(T__2);
			setState(105);
			match(NON_VOID_TYPE);
			setState(106);
			match(ID);
			setState(107);
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

	public static class FuncDefContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(BistatParser.ID, 0); }
		public TerminalNode RETURN_TYPE() { return getToken(BistatParser.RETURN_TYPE, 0); }
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
		enterRule(_localctx, 4, RULE_funcDef);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(109);
			match(T__3);
			setState(110);
			match(ID);
			setState(111);
			match(T__4);
			setState(115);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__2) {
				{
				{
				setState(112);
				paramDeclaration();
				}
				}
				setState(117);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(118);
			match(T__5);
			setState(119);
			match(T__6);
			setState(120);
			match(RETURN_TYPE);
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
			setState(127);
			match(T__7);
			setState(129); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(128);
				stmt();
				}
				}
				setState(131); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__14) | (1L << T__15) | (1L << T__17) | (1L << T__18) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << ID))) != 0) );
			setState(133);
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
		enterRule(_localctx, 6, RULE_main);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(135);
			match(T__9);
			setState(136);
			match(T__4);
			setState(137);
			match(T__5);
			setState(138);
			match(T__7);
			setState(140); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(139);
				stmt();
				}
				}
				setState(142); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__14) | (1L << T__15) | (1L << T__17) | (1L << T__18) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << ID))) != 0) );
			setState(144);
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
		enterRule(_localctx, 8, RULE_stmt);
		try {
			setState(157);
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
			case ID:
				enterOuterAlt(_localctx, 1);
				{
				{
				setState(149);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
				case 1:
					{
					setState(146);
					assignment();
					}
					break;
				case 2:
					{
					setState(147);
					specialFunction();
					}
					break;
				case 3:
					{
					setState(148);
					functionCall();
					}
					break;
				}
				setState(151);
				match(T__1);
				}
				}
				break;
			case T__18:
				enterOuterAlt(_localctx, 2);
				{
				setState(153);
				conditional();
				}
				break;
			case T__17:
				enterOuterAlt(_localctx, 3);
				{
				setState(154);
				whileLoop();
				}
				break;
			case T__15:
				enterOuterAlt(_localctx, 4);
				{
				setState(155);
				forLoop();
				}
				break;
			case T__14:
				enterOuterAlt(_localctx, 5);
				{
				setState(156);
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

	public static class ParamDeclarationContext extends ParserRuleContext {
		public TerminalNode PARAM_TYPE() { return getToken(BistatParser.PARAM_TYPE, 0); }
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
			setState(159);
			match(T__2);
			setState(160);
			match(PARAM_TYPE);
			setState(161);
			match(ID);
			setState(162);
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

	public static class AssignmentContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(BistatParser.ID, 0); }
		public TerminalNode VAR_CONS() { return getToken(BistatParser.VAR_CONS, 0); }
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
		enterRule(_localctx, 12, RULE_assignment);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(164);
			match(ID);
			setState(165);
			match(T__10);
			setState(169);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				{
				setState(166);
				match(VAR_CONS);
				}
				break;
			case 2:
				{
				setState(167);
				listAssignment();
				}
				break;
			case 3:
				{
				setState(168);
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
		public PrintContext print() {
			return getRuleContext(PrintContext.class,0);
		}
		public SpecialFunctionContext specialFunction() {
			return getRuleContext(SpecialFunctionContext.class,0);
		}
		public FunctionCallContext functionCall() {
			return getRuleContext(FunctionCallContext.class,0);
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
		enterRule(_localctx, 14, RULE_nested_stmt);
		try {
			setState(182);
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
			case ID:
				enterOuterAlt(_localctx, 1);
				{
				{
				setState(175);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
				case 1:
					{
					setState(171);
					assignment();
					}
					break;
				case 2:
					{
					setState(172);
					print();
					}
					break;
				case 3:
					{
					setState(173);
					specialFunction();
					}
					break;
				case 4:
					{
					setState(174);
					functionCall();
					}
					break;
				}
				setState(177);
				match(T__1);
				}
				}
				break;
			case T__18:
				enterOuterAlt(_localctx, 2);
				{
				setState(179);
				conditional();
				}
				break;
			case T__17:
				enterOuterAlt(_localctx, 3);
				{
				setState(180);
				whileLoop();
				}
				break;
			case T__14:
				enterOuterAlt(_localctx, 4);
				{
				setState(181);
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
		enterRule(_localctx, 16, RULE_matrixAssignment);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(184);
			match(T__11);
			setState(185);
			listAssignment();
			setState(190);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__12) {
				{
				{
				setState(186);
				match(T__12);
				setState(187);
				listAssignment();
				}
				}
				setState(192);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(193);
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
		public List<TerminalNode> VAR_CONS() { return getTokens(BistatParser.VAR_CONS); }
		public TerminalNode VAR_CONS(int i) {
			return getToken(BistatParser.VAR_CONS, i);
		}
		public ListAssignmentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listAssignment; }
	}

	public final ListAssignmentContext listAssignment() throws RecognitionException {
		ListAssignmentContext _localctx = new ListAssignmentContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_listAssignment);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(195);
			match(T__11);
			setState(196);
			match(VAR_CONS);
			setState(201);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__12) {
				{
				{
				setState(197);
				match(T__12);
				setState(198);
				match(VAR_CONS);
				}
				}
				setState(203);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(204);
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
		enterRule(_localctx, 20, RULE_comment);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(206);
			match(T__14);
			setState(208); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(207);
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
				setState(210); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__1) | (1L << T__2) | (1L << T__3) | (1L << T__4) | (1L << T__5) | (1L << T__6) | (1L << T__7) | (1L << T__8) | (1L << T__9) | (1L << T__10) | (1L << T__11) | (1L << T__12) | (1L << T__13) | (1L << T__15) | (1L << T__16) | (1L << T__17) | (1L << T__18) | (1L << T__19) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << WS) | (1L << NON_VOID_TYPE) | (1L << RETURN_TYPE) | (1L << PARAM_TYPE) | (1L << ID) | (1L << PARAM_CARDINALITY) | (1L << CARDINALITY) | (1L << VAR_CONS) | (1L << BOOL_CONS) | (1L << STRING_CONS) | (1L << FLOAT_CONS) | (1L << INT_CONS) | (1L << OP_SEC) | (1L << OP_FIRST) | (1L << LOGIC_OPERATOR) | (1L << NUMBER))) != 0) );
			setState(212);
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
		enterRule(_localctx, 22, RULE_forLoop);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(214);
			match(T__15);
			setState(215);
			match(T__4);
			setState(216);
			match(ID);
			setState(217);
			match(T__16);
			setState(218);
			expression();
			setState(219);
			match(T__5);
			setState(220);
			match(T__7);
			setState(222); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(221);
				nested_stmt();
				}
				}
				setState(224); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__14) | (1L << T__17) | (1L << T__18) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << ID))) != 0) );
			setState(226);
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
		enterRule(_localctx, 24, RULE_whileLoop);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(228);
			match(T__17);
			setState(229);
			match(T__4);
			setState(230);
			expression();
			setState(231);
			match(T__5);
			setState(232);
			match(T__7);
			setState(234); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(233);
				stmt();
				}
				}
				setState(236); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__14) | (1L << T__15) | (1L << T__17) | (1L << T__18) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << ID))) != 0) );
			setState(238);
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
		enterRule(_localctx, 26, RULE_conditional);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(240);
			match(T__18);
			setState(241);
			match(T__4);
			setState(242);
			expression();
			setState(243);
			match(T__5);
			setState(244);
			match(T__7);
			setState(246); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(245);
				stmt();
				}
				}
				setState(248); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__14) | (1L << T__15) | (1L << T__17) | (1L << T__18) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << ID))) != 0) );
			setState(250);
			match(T__8);
			setState(266);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,18,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(251);
					match(T__19);
					setState(252);
					match(T__18);
					setState(253);
					match(T__4);
					setState(254);
					expression();
					setState(255);
					match(T__5);
					setState(256);
					match(T__7);
					setState(258); 
					_errHandler.sync(this);
					_la = _input.LA(1);
					do {
						{
						{
						setState(257);
						stmt();
						}
						}
						setState(260); 
						_errHandler.sync(this);
						_la = _input.LA(1);
					} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__14) | (1L << T__15) | (1L << T__17) | (1L << T__18) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << ID))) != 0) );
					setState(262);
					match(T__8);
					}
					} 
				}
				setState(268);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,18,_ctx);
			}
			setState(278);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__19) {
				{
				setState(269);
				match(T__19);
				setState(270);
				match(T__7);
				setState(272); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(271);
					stmt();
					}
					}
					setState(274); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__14) | (1L << T__15) | (1L << T__17) | (1L << T__18) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << ID))) != 0) );
				setState(276);
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
		enterRule(_localctx, 28, RULE_specialFunction);
		try {
			setState(302);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__20:
				enterOuterAlt(_localctx, 1);
				{
				setState(280);
				inputRead();
				}
				break;
			case T__21:
				enterOuterAlt(_localctx, 2);
				{
				setState(281);
				print();
				}
				break;
			case T__22:
				enterOuterAlt(_localctx, 3);
				{
				setState(282);
				listAdd();
				}
				break;
			case T__23:
				enterOuterAlt(_localctx, 4);
				{
				setState(283);
				listPop();
				}
				break;
			case T__24:
				enterOuterAlt(_localctx, 5);
				{
				setState(284);
				length();
				}
				break;
			case T__25:
				enterOuterAlt(_localctx, 6);
				{
				setState(285);
				range();
				}
				break;
			case T__26:
				enterOuterAlt(_localctx, 7);
				{
				setState(286);
				plot();
				}
				break;
			case T__27:
				enterOuterAlt(_localctx, 8);
				{
				setState(287);
				sum();
				}
				break;
			case T__28:
				enterOuterAlt(_localctx, 9);
				{
				setState(288);
				min();
				}
				break;
			case T__29:
				enterOuterAlt(_localctx, 10);
				{
				setState(289);
				prod();
				}
				break;
			case T__30:
				enterOuterAlt(_localctx, 11);
				{
				setState(290);
				avg();
				}
				break;
			case T__31:
				enterOuterAlt(_localctx, 12);
				{
				setState(291);
				sMode();
				}
				break;
			case T__32:
				enterOuterAlt(_localctx, 13);
				{
				setState(292);
				median();
				}
				break;
			case T__33:
				enterOuterAlt(_localctx, 14);
				{
				setState(293);
				sin();
				}
				break;
			case T__35:
				enterOuterAlt(_localctx, 15);
				{
				setState(294);
				cos();
				}
				break;
			case T__34:
				enterOuterAlt(_localctx, 16);
				{
				setState(295);
				tan();
				}
				break;
			case T__36:
				enterOuterAlt(_localctx, 17);
				{
				setState(296);
				sort();
				}
				break;
			case T__37:
				enterOuterAlt(_localctx, 18);
				{
				setState(297);
				sqrt();
				}
				break;
			case T__38:
				enterOuterAlt(_localctx, 19);
				{
				setState(298);
				floor();
				}
				break;
			case T__39:
				enterOuterAlt(_localctx, 20);
				{
				setState(299);
				ceil();
				}
				break;
			case T__40:
				enterOuterAlt(_localctx, 21);
				{
				setState(300);
				abs();
				}
				break;
			case T__41:
				enterOuterAlt(_localctx, 22);
				{
				setState(301);
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
		enterRule(_localctx, 30, RULE_inputRead);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(304);
			match(T__20);
			setState(305);
			match(T__4);
			setState(306);
			match(ID);
			setState(311);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__12) {
				{
				{
				setState(307);
				match(T__12);
				setState(308);
				match(ID);
				}
				}
				setState(313);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(314);
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
		enterRule(_localctx, 32, RULE_print);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(316);
			match(T__21);
			setState(317);
			match(T__4);
			setState(318);
			expression();
			setState(323);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__12) {
				{
				{
				setState(319);
				match(T__12);
				setState(320);
				expression();
				}
				}
				setState(325);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(326);
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
		enterRule(_localctx, 34, RULE_listAdd);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(328);
			match(T__22);
			setState(329);
			match(T__4);
			setState(330);
			expression();
			setState(331);
			match(T__12);
			setState(332);
			expression();
			setState(333);
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
		enterRule(_localctx, 36, RULE_listPop);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(335);
			match(T__23);
			setState(336);
			match(T__4);
			setState(337);
			expression();
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
		enterRule(_localctx, 38, RULE_length);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(340);
			match(T__24);
			setState(341);
			match(T__4);
			setState(342);
			expression();
			setState(343);
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
		enterRule(_localctx, 40, RULE_range);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(345);
			match(T__25);
			setState(346);
			match(T__4);
			setState(347);
			expression();
			setState(350);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__12) {
				{
				setState(348);
				match(T__12);
				setState(349);
				expression();
				}
			}

			setState(352);
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
		enterRule(_localctx, 42, RULE_plot);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(354);
			match(T__26);
			setState(355);
			match(T__4);
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
		enterRule(_localctx, 44, RULE_sum);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(359);
			match(T__27);
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
		enterRule(_localctx, 46, RULE_min);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(364);
			match(T__28);
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
		enterRule(_localctx, 48, RULE_prod);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(369);
			match(T__29);
			setState(370);
			match(T__4);
			setState(371);
			expression();
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
		enterRule(_localctx, 50, RULE_avg);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(374);
			match(T__30);
			setState(375);
			match(T__4);
			setState(376);
			expression();
			setState(377);
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
		enterRule(_localctx, 52, RULE_sMode);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(379);
			match(T__31);
			setState(380);
			match(T__4);
			setState(381);
			expression();
			setState(382);
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
		enterRule(_localctx, 54, RULE_median);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(384);
			match(T__32);
			setState(385);
			match(T__4);
			setState(386);
			expression();
			setState(387);
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
		enterRule(_localctx, 56, RULE_sin);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(389);
			match(T__33);
			setState(390);
			match(T__4);
			setState(391);
			expression();
			setState(392);
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
		enterRule(_localctx, 58, RULE_tan);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(394);
			match(T__34);
			setState(395);
			match(T__4);
			setState(396);
			expression();
			setState(397);
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
		enterRule(_localctx, 60, RULE_cos);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(399);
			match(T__35);
			setState(400);
			match(T__4);
			setState(401);
			expression();
			setState(402);
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
		enterRule(_localctx, 62, RULE_sort);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(404);
			match(T__36);
			setState(405);
			match(T__4);
			setState(406);
			expression();
			setState(407);
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
		enterRule(_localctx, 64, RULE_sqrt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(409);
			match(T__37);
			setState(410);
			match(T__4);
			setState(411);
			expression();
			setState(412);
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
		enterRule(_localctx, 66, RULE_floor);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(414);
			match(T__38);
			setState(415);
			match(T__4);
			setState(416);
			expression();
			setState(417);
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
		enterRule(_localctx, 68, RULE_ceil);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(419);
			match(T__39);
			setState(420);
			match(T__4);
			setState(421);
			expression();
			setState(422);
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
		enterRule(_localctx, 70, RULE_abs);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(424);
			match(T__40);
			setState(425);
			match(T__4);
			setState(426);
			expression();
			setState(427);
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
		enterRule(_localctx, 72, RULE_not);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(429);
			match(T__41);
			setState(430);
			match(T__4);
			setState(431);
			expression();
			setState(432);
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
		enterRule(_localctx, 74, RULE_expression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(434);
			exp();
			setState(437);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LOGIC_OPERATOR) {
				{
				setState(435);
				match(LOGIC_OPERATOR);
				setState(436);
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
		enterRule(_localctx, 76, RULE_exp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(439);
			term();
			setState(442);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==OP_SEC) {
				{
				setState(440);
				match(OP_SEC);
				setState(441);
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
		enterRule(_localctx, 78, RULE_term);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(444);
			factor();
			setState(447);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==OP_FIRST) {
				{
				setState(445);
				match(OP_FIRST);
				setState(446);
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
		public TerminalNode VAR_CONS() { return getToken(BistatParser.VAR_CONS, 0); }
		public TerminalNode OP_SEC() { return getToken(BistatParser.OP_SEC, 0); }
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
		enterRule(_localctx, 80, RULE_factor);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(450);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==OP_SEC) {
				{
				setState(449);
				match(OP_SEC);
				}
			}

			setState(460);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,29,_ctx) ) {
			case 1:
				{
				{
				setState(452);
				match(T__4);
				setState(453);
				expression();
				setState(454);
				match(T__5);
				}
				}
				break;
			case 2:
				{
				setState(456);
				indexing();
				}
				break;
			case 3:
				{
				setState(457);
				specialFunction();
				}
				break;
			case 4:
				{
				setState(458);
				functionCall();
				}
				break;
			case 5:
				{
				setState(459);
				match(VAR_CONS);
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
		enterRule(_localctx, 82, RULE_functionCall);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(462);
			match(ID);
			setState(463);
			match(T__4);
			setState(472);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__4) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << ID) | (1L << VAR_CONS) | (1L << OP_SEC))) != 0)) {
				{
				setState(464);
				expression();
				setState(469);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__12) {
					{
					{
					setState(465);
					match(T__12);
					setState(466);
					expression();
					}
					}
					setState(471);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

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
		enterRule(_localctx, 84, RULE_indexing);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(476);
			match(ID);
			setState(477);
			match(T__11);
			setState(478);
			expression();
			setState(479);
			match(T__13);
			setState(484);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__11) {
				{
				setState(480);
				match(T__11);
				setState(481);
				expression();
				setState(482);
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

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3<\u01e9\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\3\2\3\2\3\2\3\2\7\2]\n\2\f\2\16\2`\13\2\3\2\7\2c\n\2\f\2\16\2f\13"+
		"\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\7\4t\n\4\f\4\16\4w"+
		"\13\4\3\4\3\4\3\4\3\4\7\4}\n\4\f\4\16\4\u0080\13\4\3\4\3\4\6\4\u0084\n"+
		"\4\r\4\16\4\u0085\3\4\3\4\3\5\3\5\3\5\3\5\3\5\6\5\u008f\n\5\r\5\16\5\u0090"+
		"\3\5\3\5\3\6\3\6\3\6\5\6\u0098\n\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6\u00a0\n"+
		"\6\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\5\b\u00ac\n\b\3\t\3\t\3\t\3"+
		"\t\5\t\u00b2\n\t\3\t\3\t\3\t\3\t\3\t\5\t\u00b9\n\t\3\n\3\n\3\n\3\n\7\n"+
		"\u00bf\n\n\f\n\16\n\u00c2\13\n\3\n\3\n\3\13\3\13\3\13\3\13\7\13\u00ca"+
		"\n\13\f\13\16\13\u00cd\13\13\3\13\3\13\3\f\3\f\6\f\u00d3\n\f\r\f\16\f"+
		"\u00d4\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\6\r\u00e1\n\r\r\r\16\r"+
		"\u00e2\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16\6\16\u00ed\n\16\r\16\16\16"+
		"\u00ee\3\16\3\16\3\17\3\17\3\17\3\17\3\17\3\17\6\17\u00f9\n\17\r\17\16"+
		"\17\u00fa\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\6\17\u0105\n\17\r\17"+
		"\16\17\u0106\3\17\3\17\7\17\u010b\n\17\f\17\16\17\u010e\13\17\3\17\3\17"+
		"\3\17\6\17\u0113\n\17\r\17\16\17\u0114\3\17\3\17\5\17\u0119\n\17\3\20"+
		"\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20"+
		"\3\20\3\20\3\20\3\20\3\20\3\20\3\20\5\20\u0131\n\20\3\21\3\21\3\21\3\21"+
		"\3\21\7\21\u0138\n\21\f\21\16\21\u013b\13\21\3\21\3\21\3\22\3\22\3\22"+
		"\3\22\3\22\7\22\u0144\n\22\f\22\16\22\u0147\13\22\3\22\3\22\3\23\3\23"+
		"\3\23\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25"+
		"\3\25\3\26\3\26\3\26\3\26\3\26\5\26\u0161\n\26\3\26\3\26\3\27\3\27\3\27"+
		"\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\32\3\32"+
		"\3\32\3\32\3\32\3\33\3\33\3\33\3\33\3\33\3\34\3\34\3\34\3\34\3\34\3\35"+
		"\3\35\3\35\3\35\3\35\3\36\3\36\3\36\3\36\3\36\3\37\3\37\3\37\3\37\3\37"+
		"\3 \3 \3 \3 \3 \3!\3!\3!\3!\3!\3\"\3\"\3\"\3\"\3\"\3#\3#\3#\3#\3#\3$\3"+
		"$\3$\3$\3$\3%\3%\3%\3%\3%\3&\3&\3&\3&\3&\3\'\3\'\3\'\5\'\u01b8\n\'\3("+
		"\3(\3(\5(\u01bd\n(\3)\3)\3)\5)\u01c2\n)\3*\5*\u01c5\n*\3*\3*\3*\3*\3*"+
		"\3*\3*\3*\5*\u01cf\n*\3+\3+\3+\3+\3+\7+\u01d6\n+\f+\16+\u01d9\13+\5+\u01db"+
		"\n+\3+\3+\3,\3,\3,\3,\3,\3,\3,\3,\5,\u01e7\n,\3,\2\2-\2\4\6\b\n\f\16\20"+
		"\22\24\26\30\32\34\36 \"$&(*,.\60\62\64\668:<>@BDFHJLNPRTV\2\3\3\2\21"+
		"\21\2\u01fe\2X\3\2\2\2\4j\3\2\2\2\6o\3\2\2\2\b\u0089\3\2\2\2\n\u009f\3"+
		"\2\2\2\f\u00a1\3\2\2\2\16\u00a6\3\2\2\2\20\u00b8\3\2\2\2\22\u00ba\3\2"+
		"\2\2\24\u00c5\3\2\2\2\26\u00d0\3\2\2\2\30\u00d8\3\2\2\2\32\u00e6\3\2\2"+
		"\2\34\u00f2\3\2\2\2\36\u0130\3\2\2\2 \u0132\3\2\2\2\"\u013e\3\2\2\2$\u014a"+
		"\3\2\2\2&\u0151\3\2\2\2(\u0156\3\2\2\2*\u015b\3\2\2\2,\u0164\3\2\2\2."+
		"\u0169\3\2\2\2\60\u016e\3\2\2\2\62\u0173\3\2\2\2\64\u0178\3\2\2\2\66\u017d"+
		"\3\2\2\28\u0182\3\2\2\2:\u0187\3\2\2\2<\u018c\3\2\2\2>\u0191\3\2\2\2@"+
		"\u0196\3\2\2\2B\u019b\3\2\2\2D\u01a0\3\2\2\2F\u01a5\3\2\2\2H\u01aa\3\2"+
		"\2\2J\u01af\3\2\2\2L\u01b4\3\2\2\2N\u01b9\3\2\2\2P\u01be\3\2\2\2R\u01c4"+
		"\3\2\2\2T\u01d0\3\2\2\2V\u01de\3\2\2\2XY\7\3\2\2YZ\7\61\2\2Z^\7\4\2\2"+
		"[]\5\4\3\2\\[\3\2\2\2]`\3\2\2\2^\\\3\2\2\2^_\3\2\2\2_d\3\2\2\2`^\3\2\2"+
		"\2ac\5\6\4\2ba\3\2\2\2cf\3\2\2\2db\3\2\2\2de\3\2\2\2eg\3\2\2\2fd\3\2\2"+
		"\2gh\5\b\5\2hi\7\2\2\3i\3\3\2\2\2jk\7\5\2\2kl\7.\2\2lm\7\61\2\2mn\7\4"+
		"\2\2n\5\3\2\2\2op\7\6\2\2pq\7\61\2\2qu\7\7\2\2rt\5\f\7\2sr\3\2\2\2tw\3"+
		"\2\2\2us\3\2\2\2uv\3\2\2\2vx\3\2\2\2wu\3\2\2\2xy\7\b\2\2yz\7\t\2\2z~\7"+
		"/\2\2{}\5\4\3\2|{\3\2\2\2}\u0080\3\2\2\2~|\3\2\2\2~\177\3\2\2\2\177\u0081"+
		"\3\2\2\2\u0080~\3\2\2\2\u0081\u0083\7\n\2\2\u0082\u0084\5\n\6\2\u0083"+
		"\u0082\3\2\2\2\u0084\u0085\3\2\2\2\u0085\u0083\3\2\2\2\u0085\u0086\3\2"+
		"\2\2\u0086\u0087\3\2\2\2\u0087\u0088\7\13\2\2\u0088\7\3\2\2\2\u0089\u008a"+
		"\7\f\2\2\u008a\u008b\7\7\2\2\u008b\u008c\7\b\2\2\u008c\u008e\7\n\2\2\u008d"+
		"\u008f\5\n\6\2\u008e\u008d\3\2\2\2\u008f\u0090\3\2\2\2\u0090\u008e\3\2"+
		"\2\2\u0090\u0091\3\2\2\2\u0091\u0092\3\2\2\2\u0092\u0093\7\13\2\2\u0093"+
		"\t\3\2\2\2\u0094\u0098\5\16\b\2\u0095\u0098\5\36\20\2\u0096\u0098\5T+"+
		"\2\u0097\u0094\3\2\2\2\u0097\u0095\3\2\2\2\u0097\u0096\3\2\2\2\u0098\u0099"+
		"\3\2\2\2\u0099\u009a\7\4\2\2\u009a\u00a0\3\2\2\2\u009b\u00a0\5\34\17\2"+
		"\u009c\u00a0\5\32\16\2\u009d\u00a0\5\30\r\2\u009e\u00a0\5\26\f\2\u009f"+
		"\u0097\3\2\2\2\u009f\u009b\3\2\2\2\u009f\u009c\3\2\2\2\u009f\u009d\3\2"+
		"\2\2\u009f\u009e\3\2\2\2\u00a0\13\3\2\2\2\u00a1\u00a2\7\5\2\2\u00a2\u00a3"+
		"\7\60\2\2\u00a3\u00a4\7\61\2\2\u00a4\u00a5\7\4\2\2\u00a5\r\3\2\2\2\u00a6"+
		"\u00a7\7\61\2\2\u00a7\u00ab\7\r\2\2\u00a8\u00ac\7\64\2\2\u00a9\u00ac\5"+
		"\24\13\2\u00aa\u00ac\5\22\n\2\u00ab\u00a8\3\2\2\2\u00ab\u00a9\3\2\2\2"+
		"\u00ab\u00aa\3\2\2\2\u00ac\17\3\2\2\2\u00ad\u00b2\5\16\b\2\u00ae\u00b2"+
		"\5\"\22\2\u00af\u00b2\5\36\20\2\u00b0\u00b2\5T+\2\u00b1\u00ad\3\2\2\2"+
		"\u00b1\u00ae\3\2\2\2\u00b1\u00af\3\2\2\2\u00b1\u00b0\3\2\2\2\u00b2\u00b3"+
		"\3\2\2\2\u00b3\u00b4\7\4\2\2\u00b4\u00b9\3\2\2\2\u00b5\u00b9\5\34\17\2"+
		"\u00b6\u00b9\5\32\16\2\u00b7\u00b9\5\26\f\2\u00b8\u00b1\3\2\2\2\u00b8"+
		"\u00b5\3\2\2\2\u00b8\u00b6\3\2\2\2\u00b8\u00b7\3\2\2\2\u00b9\21\3\2\2"+
		"\2\u00ba\u00bb\7\16\2\2\u00bb\u00c0\5\24\13\2\u00bc\u00bd\7\17\2\2\u00bd"+
		"\u00bf\5\24\13\2\u00be\u00bc\3\2\2\2\u00bf\u00c2\3\2\2\2\u00c0\u00be\3"+
		"\2\2\2\u00c0\u00c1\3\2\2\2\u00c1\u00c3\3\2\2\2\u00c2\u00c0\3\2\2\2\u00c3"+
		"\u00c4\7\20\2\2\u00c4\23\3\2\2\2\u00c5\u00c6\7\16\2\2\u00c6\u00cb\7\64"+
		"\2\2\u00c7\u00c8\7\17\2\2\u00c8\u00ca\7\64\2\2\u00c9\u00c7\3\2\2\2\u00ca"+
		"\u00cd\3\2\2\2\u00cb\u00c9\3\2\2\2\u00cb\u00cc\3\2\2\2\u00cc\u00ce\3\2"+
		"\2\2\u00cd\u00cb\3\2\2\2\u00ce\u00cf\7\20\2\2\u00cf\25\3\2\2\2\u00d0\u00d2"+
		"\7\21\2\2\u00d1\u00d3\n\2\2\2\u00d2\u00d1\3\2\2\2\u00d3\u00d4\3\2\2\2"+
		"\u00d4\u00d2\3\2\2\2\u00d4\u00d5\3\2\2\2\u00d5\u00d6\3\2\2\2\u00d6\u00d7"+
		"\7\21\2\2\u00d7\27\3\2\2\2\u00d8\u00d9\7\22\2\2\u00d9\u00da\7\7\2\2\u00da"+
		"\u00db\7\61\2\2\u00db\u00dc\7\23\2\2\u00dc\u00dd\5L\'\2\u00dd\u00de\7"+
		"\b\2\2\u00de\u00e0\7\n\2\2\u00df\u00e1\5\20\t\2\u00e0\u00df\3\2\2\2\u00e1"+
		"\u00e2\3\2\2\2\u00e2\u00e0\3\2\2\2\u00e2\u00e3\3\2\2\2\u00e3\u00e4\3\2"+
		"\2\2\u00e4\u00e5\7\13\2\2\u00e5\31\3\2\2\2\u00e6\u00e7\7\24\2\2\u00e7"+
		"\u00e8\7\7\2\2\u00e8\u00e9\5L\'\2\u00e9\u00ea\7\b\2\2\u00ea\u00ec\7\n"+
		"\2\2\u00eb\u00ed\5\n\6\2\u00ec\u00eb\3\2\2\2\u00ed\u00ee\3\2\2\2\u00ee"+
		"\u00ec\3\2\2\2\u00ee\u00ef\3\2\2\2\u00ef\u00f0\3\2\2\2\u00f0\u00f1\7\13"+
		"\2\2\u00f1\33\3\2\2\2\u00f2\u00f3\7\25\2\2\u00f3\u00f4\7\7\2\2\u00f4\u00f5"+
		"\5L\'\2\u00f5\u00f6\7\b\2\2\u00f6\u00f8\7\n\2\2\u00f7\u00f9\5\n\6\2\u00f8"+
		"\u00f7\3\2\2\2\u00f9\u00fa\3\2\2\2\u00fa\u00f8\3\2\2\2\u00fa\u00fb\3\2"+
		"\2\2\u00fb\u00fc\3\2\2\2\u00fc\u010c\7\13\2\2\u00fd\u00fe\7\26\2\2\u00fe"+
		"\u00ff\7\25\2\2\u00ff\u0100\7\7\2\2\u0100\u0101\5L\'\2\u0101\u0102\7\b"+
		"\2\2\u0102\u0104\7\n\2\2\u0103\u0105\5\n\6\2\u0104\u0103\3\2\2\2\u0105"+
		"\u0106\3\2\2\2\u0106\u0104\3\2\2\2\u0106\u0107\3\2\2\2\u0107\u0108\3\2"+
		"\2\2\u0108\u0109\7\13\2\2\u0109\u010b\3\2\2\2\u010a\u00fd\3\2\2\2\u010b"+
		"\u010e\3\2\2\2\u010c\u010a\3\2\2\2\u010c\u010d\3\2\2\2\u010d\u0118\3\2"+
		"\2\2\u010e\u010c\3\2\2\2\u010f\u0110\7\26\2\2\u0110\u0112\7\n\2\2\u0111"+
		"\u0113\5\n\6\2\u0112\u0111\3\2\2\2\u0113\u0114\3\2\2\2\u0114\u0112\3\2"+
		"\2\2\u0114\u0115\3\2\2\2\u0115\u0116\3\2\2\2\u0116\u0117\7\13\2\2\u0117"+
		"\u0119\3\2\2\2\u0118\u010f\3\2\2\2\u0118\u0119\3\2\2\2\u0119\35\3\2\2"+
		"\2\u011a\u0131\5 \21\2\u011b\u0131\5\"\22\2\u011c\u0131\5$\23\2\u011d"+
		"\u0131\5&\24\2\u011e\u0131\5(\25\2\u011f\u0131\5*\26\2\u0120\u0131\5,"+
		"\27\2\u0121\u0131\5.\30\2\u0122\u0131\5\60\31\2\u0123\u0131\5\62\32\2"+
		"\u0124\u0131\5\64\33\2\u0125\u0131\5\66\34\2\u0126\u0131\58\35\2\u0127"+
		"\u0131\5:\36\2\u0128\u0131\5> \2\u0129\u0131\5<\37\2\u012a\u0131\5@!\2"+
		"\u012b\u0131\5B\"\2\u012c\u0131\5D#\2\u012d\u0131\5F$\2\u012e\u0131\5"+
		"H%\2\u012f\u0131\5J&\2\u0130\u011a\3\2\2\2\u0130\u011b\3\2\2\2\u0130\u011c"+
		"\3\2\2\2\u0130\u011d\3\2\2\2\u0130\u011e\3\2\2\2\u0130\u011f\3\2\2\2\u0130"+
		"\u0120\3\2\2\2\u0130\u0121\3\2\2\2\u0130\u0122\3\2\2\2\u0130\u0123\3\2"+
		"\2\2\u0130\u0124\3\2\2\2\u0130\u0125\3\2\2\2\u0130\u0126\3\2\2\2\u0130"+
		"\u0127\3\2\2\2\u0130\u0128\3\2\2\2\u0130\u0129\3\2\2\2\u0130\u012a\3\2"+
		"\2\2\u0130\u012b\3\2\2\2\u0130\u012c\3\2\2\2\u0130\u012d\3\2\2\2\u0130"+
		"\u012e\3\2\2\2\u0130\u012f\3\2\2\2\u0131\37\3\2\2\2\u0132\u0133\7\27\2"+
		"\2\u0133\u0134\7\7\2\2\u0134\u0139\7\61\2\2\u0135\u0136\7\17\2\2\u0136"+
		"\u0138\7\61\2\2\u0137\u0135\3\2\2\2\u0138\u013b\3\2\2\2\u0139\u0137\3"+
		"\2\2\2\u0139\u013a\3\2\2\2\u013a\u013c\3\2\2\2\u013b\u0139\3\2\2\2\u013c"+
		"\u013d\7\b\2\2\u013d!\3\2\2\2\u013e\u013f\7\30\2\2\u013f\u0140\7\7\2\2"+
		"\u0140\u0145\5L\'\2\u0141\u0142\7\17\2\2\u0142\u0144\5L\'\2\u0143\u0141"+
		"\3\2\2\2\u0144\u0147\3\2\2\2\u0145\u0143\3\2\2\2\u0145\u0146\3\2\2\2\u0146"+
		"\u0148\3\2\2\2\u0147\u0145\3\2\2\2\u0148\u0149\7\b\2\2\u0149#\3\2\2\2"+
		"\u014a\u014b\7\31\2\2\u014b\u014c\7\7\2\2\u014c\u014d\5L\'\2\u014d\u014e"+
		"\7\17\2\2\u014e\u014f\5L\'\2\u014f\u0150\7\b\2\2\u0150%\3\2\2\2\u0151"+
		"\u0152\7\32\2\2\u0152\u0153\7\7\2\2\u0153\u0154\5L\'\2\u0154\u0155\7\b"+
		"\2\2\u0155\'\3\2\2\2\u0156\u0157\7\33\2\2\u0157\u0158\7\7\2\2\u0158\u0159"+
		"\5L\'\2\u0159\u015a\7\b\2\2\u015a)\3\2\2\2\u015b\u015c\7\34\2\2\u015c"+
		"\u015d\7\7\2\2\u015d\u0160\5L\'\2\u015e\u015f\7\17\2\2\u015f\u0161\5L"+
		"\'\2\u0160\u015e\3\2\2\2\u0160\u0161\3\2\2\2\u0161\u0162\3\2\2\2\u0162"+
		"\u0163\7\b\2\2\u0163+\3\2\2\2\u0164\u0165\7\35\2\2\u0165\u0166\7\7\2\2"+
		"\u0166\u0167\5L\'\2\u0167\u0168\7\b\2\2\u0168-\3\2\2\2\u0169\u016a\7\36"+
		"\2\2\u016a\u016b\7\7\2\2\u016b\u016c\5L\'\2\u016c\u016d\7\b\2\2\u016d"+
		"/\3\2\2\2\u016e\u016f\7\37\2\2\u016f\u0170\7\7\2\2\u0170\u0171\5L\'\2"+
		"\u0171\u0172\7\b\2\2\u0172\61\3\2\2\2\u0173\u0174\7 \2\2\u0174\u0175\7"+
		"\7\2\2\u0175\u0176\5L\'\2\u0176\u0177\7\b\2\2\u0177\63\3\2\2\2\u0178\u0179"+
		"\7!\2\2\u0179\u017a\7\7\2\2\u017a\u017b\5L\'\2\u017b\u017c\7\b\2\2\u017c"+
		"\65\3\2\2\2\u017d\u017e\7\"\2\2\u017e\u017f\7\7\2\2\u017f\u0180\5L\'\2"+
		"\u0180\u0181\7\b\2\2\u0181\67\3\2\2\2\u0182\u0183\7#\2\2\u0183\u0184\7"+
		"\7\2\2\u0184\u0185\5L\'\2\u0185\u0186\7\b\2\2\u01869\3\2\2\2\u0187\u0188"+
		"\7$\2\2\u0188\u0189\7\7\2\2\u0189\u018a\5L\'\2\u018a\u018b\7\b\2\2\u018b"+
		";\3\2\2\2\u018c\u018d\7%\2\2\u018d\u018e\7\7\2\2\u018e\u018f\5L\'\2\u018f"+
		"\u0190\7\b\2\2\u0190=\3\2\2\2\u0191\u0192\7&\2\2\u0192\u0193\7\7\2\2\u0193"+
		"\u0194\5L\'\2\u0194\u0195\7\b\2\2\u0195?\3\2\2\2\u0196\u0197\7\'\2\2\u0197"+
		"\u0198\7\7\2\2\u0198\u0199\5L\'\2\u0199\u019a\7\b\2\2\u019aA\3\2\2\2\u019b"+
		"\u019c\7(\2\2\u019c\u019d\7\7\2\2\u019d\u019e\5L\'\2\u019e\u019f\7\b\2"+
		"\2\u019fC\3\2\2\2\u01a0\u01a1\7)\2\2\u01a1\u01a2\7\7\2\2\u01a2\u01a3\5"+
		"L\'\2\u01a3\u01a4\7\b\2\2\u01a4E\3\2\2\2\u01a5\u01a6\7*\2\2\u01a6\u01a7"+
		"\7\7\2\2\u01a7\u01a8\5L\'\2\u01a8\u01a9\7\b\2\2\u01a9G\3\2\2\2\u01aa\u01ab"+
		"\7+\2\2\u01ab\u01ac\7\7\2\2\u01ac\u01ad\5L\'\2\u01ad\u01ae\7\b\2\2\u01ae"+
		"I\3\2\2\2\u01af\u01b0\7,\2\2\u01b0\u01b1\7\7\2\2\u01b1\u01b2\5L\'\2\u01b2"+
		"\u01b3\7\b\2\2\u01b3K\3\2\2\2\u01b4\u01b7\5N(\2\u01b5\u01b6\7;\2\2\u01b6"+
		"\u01b8\5N(\2\u01b7\u01b5\3\2\2\2\u01b7\u01b8\3\2\2\2\u01b8M\3\2\2\2\u01b9"+
		"\u01bc\5P)\2\u01ba\u01bb\79\2\2\u01bb\u01bd\5N(\2\u01bc\u01ba\3\2\2\2"+
		"\u01bc\u01bd\3\2\2\2\u01bdO\3\2\2\2\u01be\u01c1\5R*\2\u01bf\u01c0\7:\2"+
		"\2\u01c0\u01c2\5P)\2\u01c1\u01bf\3\2\2\2\u01c1\u01c2\3\2\2\2\u01c2Q\3"+
		"\2\2\2\u01c3\u01c5\79\2\2\u01c4\u01c3\3\2\2\2\u01c4\u01c5\3\2\2\2\u01c5"+
		"\u01ce\3\2\2\2\u01c6\u01c7\7\7\2\2\u01c7\u01c8\5L\'\2\u01c8\u01c9\7\b"+
		"\2\2\u01c9\u01cf\3\2\2\2\u01ca\u01cf\5V,\2\u01cb\u01cf\5\36\20\2\u01cc"+
		"\u01cf\5T+\2\u01cd\u01cf\7\64\2\2\u01ce\u01c6\3\2\2\2\u01ce\u01ca\3\2"+
		"\2\2\u01ce\u01cb\3\2\2\2\u01ce\u01cc\3\2\2\2\u01ce\u01cd\3\2\2\2\u01cf"+
		"S\3\2\2\2\u01d0\u01d1\7\61\2\2\u01d1\u01da\7\7\2\2\u01d2\u01d7\5L\'\2"+
		"\u01d3\u01d4\7\17\2\2\u01d4\u01d6\5L\'\2\u01d5\u01d3\3\2\2\2\u01d6\u01d9"+
		"\3\2\2\2\u01d7\u01d5\3\2\2\2\u01d7\u01d8\3\2\2\2\u01d8\u01db\3\2\2\2\u01d9"+
		"\u01d7\3\2\2\2\u01da\u01d2\3\2\2\2\u01da\u01db\3\2\2\2\u01db\u01dc\3\2"+
		"\2\2\u01dc\u01dd\7\b\2\2\u01ddU\3\2\2\2\u01de\u01df\7\61\2\2\u01df\u01e0"+
		"\7\16\2\2\u01e0\u01e1\5L\'\2\u01e1\u01e6\7\20\2\2\u01e2\u01e3\7\16\2\2"+
		"\u01e3\u01e4\5L\'\2\u01e4\u01e5\7\20\2\2\u01e5\u01e7\3\2\2\2\u01e6\u01e2"+
		"\3\2\2\2\u01e6\u01e7\3\2\2\2\u01e7W\3\2\2\2#^du~\u0085\u0090\u0097\u009f"+
		"\u00ab\u00b1\u00b8\u00c0\u00cb\u00d4\u00e2\u00ee\u00fa\u0106\u010c\u0114"+
		"\u0118\u0130\u0139\u0145\u0160\u01b7\u01bc\u01c1\u01c4\u01ce\u01d7\u01da"+
		"\u01e6";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}