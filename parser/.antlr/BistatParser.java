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
			setState(158);
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
				setState(150);
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
					print();
					}
					break;
				case 3:
					{
					setState(148);
					specialFunction();
					}
					break;
				case 4:
					{
					setState(149);
					functionCall();
					}
					break;
				}
				setState(152);
				match(T__1);
				}
				}
				break;
			case T__18:
				enterOuterAlt(_localctx, 2);
				{
				setState(154);
				conditional();
				}
				break;
			case T__17:
				enterOuterAlt(_localctx, 3);
				{
				setState(155);
				whileLoop();
				}
				break;
			case T__15:
				enterOuterAlt(_localctx, 4);
				{
				setState(156);
				forLoop();
				}
				break;
			case T__14:
				enterOuterAlt(_localctx, 5);
				{
				setState(157);
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
			setState(160);
			match(T__2);
			setState(161);
			match(PARAM_TYPE);
			setState(162);
			match(ID);
			setState(163);
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
			setState(165);
			match(ID);
			setState(166);
			match(T__10);
			setState(170);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				{
				setState(167);
				match(VAR_CONS);
				}
				break;
			case 2:
				{
				setState(168);
				listAssignment();
				}
				break;
			case 3:
				{
				setState(169);
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
			setState(183);
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
				setState(176);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
				case 1:
					{
					setState(172);
					assignment();
					}
					break;
				case 2:
					{
					setState(173);
					print();
					}
					break;
				case 3:
					{
					setState(174);
					specialFunction();
					}
					break;
				case 4:
					{
					setState(175);
					functionCall();
					}
					break;
				}
				setState(178);
				match(T__1);
				}
				}
				break;
			case T__18:
				enterOuterAlt(_localctx, 2);
				{
				setState(180);
				conditional();
				}
				break;
			case T__17:
				enterOuterAlt(_localctx, 3);
				{
				setState(181);
				whileLoop();
				}
				break;
			case T__14:
				enterOuterAlt(_localctx, 4);
				{
				setState(182);
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
			setState(185);
			match(T__11);
			setState(186);
			listAssignment();
			setState(191);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__12) {
				{
				{
				setState(187);
				match(T__12);
				setState(188);
				listAssignment();
				}
				}
				setState(193);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(194);
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
			setState(196);
			match(T__11);
			setState(197);
			match(VAR_CONS);
			setState(202);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__12) {
				{
				{
				setState(198);
				match(T__12);
				setState(199);
				match(VAR_CONS);
				}
				}
				setState(204);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(205);
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
			setState(207);
			match(T__14);
			setState(209); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(208);
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
				setState(211); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__1) | (1L << T__2) | (1L << T__3) | (1L << T__4) | (1L << T__5) | (1L << T__6) | (1L << T__7) | (1L << T__8) | (1L << T__9) | (1L << T__10) | (1L << T__11) | (1L << T__12) | (1L << T__13) | (1L << T__15) | (1L << T__16) | (1L << T__17) | (1L << T__18) | (1L << T__19) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << WS) | (1L << NON_VOID_TYPE) | (1L << RETURN_TYPE) | (1L << PARAM_TYPE) | (1L << ID) | (1L << PARAM_CARDINALITY) | (1L << CARDINALITY) | (1L << VAR_CONS) | (1L << BOOL_CONS) | (1L << STRING_CONS) | (1L << FLOAT_CONS) | (1L << INT_CONS) | (1L << OP_SEC) | (1L << OP_FIRST) | (1L << LOGIC_OPERATOR) | (1L << NUMBER))) != 0) );
			setState(213);
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
			setState(215);
			match(T__15);
			setState(216);
			match(T__4);
			setState(217);
			match(ID);
			setState(218);
			match(T__16);
			setState(219);
			expression();
			setState(220);
			match(T__5);
			setState(221);
			match(T__7);
			setState(223); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(222);
				nested_stmt();
				}
				}
				setState(225); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__14) | (1L << T__17) | (1L << T__18) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << ID))) != 0) );
			setState(227);
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
			setState(229);
			match(T__17);
			setState(230);
			match(T__4);
			setState(231);
			expression();
			setState(232);
			match(T__5);
			setState(233);
			match(T__7);
			setState(235); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(234);
				stmt();
				}
				}
				setState(237); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__14) | (1L << T__15) | (1L << T__17) | (1L << T__18) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << ID))) != 0) );
			setState(239);
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
			setState(241);
			match(T__18);
			setState(242);
			match(T__4);
			setState(243);
			expression();
			setState(244);
			match(T__5);
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
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__14) | (1L << T__15) | (1L << T__17) | (1L << T__18) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << ID))) != 0) );
			setState(251);
			match(T__8);
			setState(267);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,18,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(252);
					match(T__19);
					setState(253);
					match(T__18);
					setState(254);
					match(T__4);
					setState(255);
					expression();
					setState(256);
					match(T__5);
					setState(257);
					match(T__7);
					setState(259); 
					_errHandler.sync(this);
					_la = _input.LA(1);
					do {
						{
						{
						setState(258);
						stmt();
						}
						}
						setState(261); 
						_errHandler.sync(this);
						_la = _input.LA(1);
					} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__14) | (1L << T__15) | (1L << T__17) | (1L << T__18) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << ID))) != 0) );
					setState(263);
					match(T__8);
					}
					} 
				}
				setState(269);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,18,_ctx);
			}
			setState(279);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__19) {
				{
				setState(270);
				match(T__19);
				setState(271);
				match(T__7);
				setState(273); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(272);
					stmt();
					}
					}
					setState(275); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__14) | (1L << T__15) | (1L << T__17) | (1L << T__18) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << ID))) != 0) );
				setState(277);
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
			setState(303);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__20:
				enterOuterAlt(_localctx, 1);
				{
				setState(281);
				inputRead();
				}
				break;
			case T__21:
				enterOuterAlt(_localctx, 2);
				{
				setState(282);
				print();
				}
				break;
			case T__22:
				enterOuterAlt(_localctx, 3);
				{
				setState(283);
				listAdd();
				}
				break;
			case T__23:
				enterOuterAlt(_localctx, 4);
				{
				setState(284);
				listPop();
				}
				break;
			case T__24:
				enterOuterAlt(_localctx, 5);
				{
				setState(285);
				length();
				}
				break;
			case T__25:
				enterOuterAlt(_localctx, 6);
				{
				setState(286);
				range();
				}
				break;
			case T__26:
				enterOuterAlt(_localctx, 7);
				{
				setState(287);
				plot();
				}
				break;
			case T__27:
				enterOuterAlt(_localctx, 8);
				{
				setState(288);
				sum();
				}
				break;
			case T__28:
				enterOuterAlt(_localctx, 9);
				{
				setState(289);
				min();
				}
				break;
			case T__29:
				enterOuterAlt(_localctx, 10);
				{
				setState(290);
				prod();
				}
				break;
			case T__30:
				enterOuterAlt(_localctx, 11);
				{
				setState(291);
				avg();
				}
				break;
			case T__31:
				enterOuterAlt(_localctx, 12);
				{
				setState(292);
				sMode();
				}
				break;
			case T__32:
				enterOuterAlt(_localctx, 13);
				{
				setState(293);
				median();
				}
				break;
			case T__33:
				enterOuterAlt(_localctx, 14);
				{
				setState(294);
				sin();
				}
				break;
			case T__35:
				enterOuterAlt(_localctx, 15);
				{
				setState(295);
				cos();
				}
				break;
			case T__34:
				enterOuterAlt(_localctx, 16);
				{
				setState(296);
				tan();
				}
				break;
			case T__36:
				enterOuterAlt(_localctx, 17);
				{
				setState(297);
				sort();
				}
				break;
			case T__37:
				enterOuterAlt(_localctx, 18);
				{
				setState(298);
				sqrt();
				}
				break;
			case T__38:
				enterOuterAlt(_localctx, 19);
				{
				setState(299);
				floor();
				}
				break;
			case T__39:
				enterOuterAlt(_localctx, 20);
				{
				setState(300);
				ceil();
				}
				break;
			case T__40:
				enterOuterAlt(_localctx, 21);
				{
				setState(301);
				abs();
				}
				break;
			case T__41:
				enterOuterAlt(_localctx, 22);
				{
				setState(302);
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
			setState(305);
			match(T__20);
			setState(306);
			match(T__4);
			setState(307);
			match(ID);
			setState(312);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__12) {
				{
				{
				setState(308);
				match(T__12);
				setState(309);
				match(ID);
				}
				}
				setState(314);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(315);
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
			setState(317);
			match(T__21);
			setState(318);
			match(T__4);
			setState(319);
			expression();
			setState(324);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__12) {
				{
				{
				setState(320);
				match(T__12);
				setState(321);
				expression();
				}
				}
				setState(326);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(327);
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
			setState(329);
			match(T__22);
			setState(330);
			match(T__4);
			setState(331);
			expression();
			setState(332);
			match(T__12);
			setState(333);
			expression();
			setState(334);
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
			setState(336);
			match(T__23);
			setState(337);
			match(T__4);
			setState(338);
			expression();
			setState(339);
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
			setState(341);
			match(T__24);
			setState(342);
			match(T__4);
			setState(343);
			expression();
			setState(344);
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
			setState(346);
			match(T__25);
			setState(347);
			match(T__4);
			setState(348);
			expression();
			setState(351);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__12) {
				{
				setState(349);
				match(T__12);
				setState(350);
				expression();
				}
			}

			setState(353);
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
			setState(355);
			match(T__26);
			setState(356);
			match(T__4);
			setState(357);
			expression();
			setState(358);
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
			setState(360);
			match(T__27);
			setState(361);
			match(T__4);
			setState(362);
			expression();
			setState(363);
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
			setState(365);
			match(T__28);
			setState(366);
			match(T__4);
			setState(367);
			expression();
			setState(368);
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
			setState(370);
			match(T__29);
			setState(371);
			match(T__4);
			setState(372);
			expression();
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
			setState(375);
			match(T__30);
			setState(376);
			match(T__4);
			setState(377);
			expression();
			setState(378);
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
			setState(380);
			match(T__31);
			setState(381);
			match(T__4);
			setState(382);
			expression();
			setState(383);
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
			setState(385);
			match(T__32);
			setState(386);
			match(T__4);
			setState(387);
			expression();
			setState(388);
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
			setState(390);
			match(T__33);
			setState(391);
			match(T__4);
			setState(392);
			expression();
			setState(393);
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
			setState(395);
			match(T__34);
			setState(396);
			match(T__4);
			setState(397);
			expression();
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
			setState(400);
			match(T__35);
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
			setState(405);
			match(T__36);
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
			setState(410);
			match(T__37);
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
			setState(415);
			match(T__38);
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
			setState(420);
			match(T__39);
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
			setState(425);
			match(T__40);
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
			setState(430);
			match(T__41);
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
			setState(435);
			exp();
			setState(438);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LOGIC_OPERATOR) {
				{
				setState(436);
				match(LOGIC_OPERATOR);
				setState(437);
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
			setState(440);
			term();
			setState(443);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==OP_SEC) {
				{
				setState(441);
				match(OP_SEC);
				setState(442);
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
			setState(445);
			factor();
			setState(448);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==OP_FIRST) {
				{
				setState(446);
				match(OP_FIRST);
				setState(447);
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
			setState(451);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==OP_SEC) {
				{
				setState(450);
				match(OP_SEC);
				}
			}

			setState(461);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,29,_ctx) ) {
			case 1:
				{
				{
				setState(453);
				match(T__4);
				setState(454);
				expression();
				setState(455);
				match(T__5);
				}
				}
				break;
			case 2:
				{
				setState(457);
				indexing();
				}
				break;
			case 3:
				{
				setState(458);
				specialFunction();
				}
				break;
			case 4:
				{
				setState(459);
				functionCall();
				}
				break;
			case 5:
				{
				setState(460);
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
			setState(463);
			match(ID);
			setState(464);
			match(T__4);
			setState(473);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__4) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__23) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << T__34) | (1L << T__35) | (1L << T__36) | (1L << T__37) | (1L << T__38) | (1L << T__39) | (1L << T__40) | (1L << T__41) | (1L << ID) | (1L << VAR_CONS) | (1L << OP_SEC))) != 0)) {
				{
				setState(465);
				expression();
				setState(470);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__12) {
					{
					{
					setState(466);
					match(T__12);
					setState(467);
					expression();
					}
					}
					setState(472);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(475);
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
			setState(477);
			match(ID);
			setState(478);
			match(T__11);
			setState(479);
			expression();
			setState(480);
			match(T__13);
			setState(485);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__11) {
				{
				setState(481);
				match(T__11);
				setState(482);
				expression();
				setState(483);
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3<\u01ea\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\3\2\3\2\3\2\3\2\7\2]\n\2\f\2\16\2`\13\2\3\2\7\2c\n\2\f\2\16\2f\13"+
		"\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\7\4t\n\4\f\4\16\4w"+
		"\13\4\3\4\3\4\3\4\3\4\7\4}\n\4\f\4\16\4\u0080\13\4\3\4\3\4\6\4\u0084\n"+
		"\4\r\4\16\4\u0085\3\4\3\4\3\5\3\5\3\5\3\5\3\5\6\5\u008f\n\5\r\5\16\5\u0090"+
		"\3\5\3\5\3\6\3\6\3\6\3\6\5\6\u0099\n\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6\u00a1"+
		"\n\6\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\5\b\u00ad\n\b\3\t\3\t\3\t"+
		"\3\t\5\t\u00b3\n\t\3\t\3\t\3\t\3\t\3\t\5\t\u00ba\n\t\3\n\3\n\3\n\3\n\7"+
		"\n\u00c0\n\n\f\n\16\n\u00c3\13\n\3\n\3\n\3\13\3\13\3\13\3\13\7\13\u00cb"+
		"\n\13\f\13\16\13\u00ce\13\13\3\13\3\13\3\f\3\f\6\f\u00d4\n\f\r\f\16\f"+
		"\u00d5\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\6\r\u00e2\n\r\r\r\16\r"+
		"\u00e3\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16\6\16\u00ee\n\16\r\16\16\16"+
		"\u00ef\3\16\3\16\3\17\3\17\3\17\3\17\3\17\3\17\6\17\u00fa\n\17\r\17\16"+
		"\17\u00fb\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\6\17\u0106\n\17\r\17"+
		"\16\17\u0107\3\17\3\17\7\17\u010c\n\17\f\17\16\17\u010f\13\17\3\17\3\17"+
		"\3\17\6\17\u0114\n\17\r\17\16\17\u0115\3\17\3\17\5\17\u011a\n\17\3\20"+
		"\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20"+
		"\3\20\3\20\3\20\3\20\3\20\3\20\3\20\5\20\u0132\n\20\3\21\3\21\3\21\3\21"+
		"\3\21\7\21\u0139\n\21\f\21\16\21\u013c\13\21\3\21\3\21\3\22\3\22\3\22"+
		"\3\22\3\22\7\22\u0145\n\22\f\22\16\22\u0148\13\22\3\22\3\22\3\23\3\23"+
		"\3\23\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25"+
		"\3\25\3\26\3\26\3\26\3\26\3\26\5\26\u0162\n\26\3\26\3\26\3\27\3\27\3\27"+
		"\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\32\3\32"+
		"\3\32\3\32\3\32\3\33\3\33\3\33\3\33\3\33\3\34\3\34\3\34\3\34\3\34\3\35"+
		"\3\35\3\35\3\35\3\35\3\36\3\36\3\36\3\36\3\36\3\37\3\37\3\37\3\37\3\37"+
		"\3 \3 \3 \3 \3 \3!\3!\3!\3!\3!\3\"\3\"\3\"\3\"\3\"\3#\3#\3#\3#\3#\3$\3"+
		"$\3$\3$\3$\3%\3%\3%\3%\3%\3&\3&\3&\3&\3&\3\'\3\'\3\'\5\'\u01b9\n\'\3("+
		"\3(\3(\5(\u01be\n(\3)\3)\3)\5)\u01c3\n)\3*\5*\u01c6\n*\3*\3*\3*\3*\3*"+
		"\3*\3*\3*\5*\u01d0\n*\3+\3+\3+\3+\3+\7+\u01d7\n+\f+\16+\u01da\13+\5+\u01dc"+
		"\n+\3+\3+\3,\3,\3,\3,\3,\3,\3,\3,\5,\u01e8\n,\3,\2\2-\2\4\6\b\n\f\16\20"+
		"\22\24\26\30\32\34\36 \"$&(*,.\60\62\64\668:<>@BDFHJLNPRTV\2\3\3\2\21"+
		"\21\2\u0200\2X\3\2\2\2\4j\3\2\2\2\6o\3\2\2\2\b\u0089\3\2\2\2\n\u00a0\3"+
		"\2\2\2\f\u00a2\3\2\2\2\16\u00a7\3\2\2\2\20\u00b9\3\2\2\2\22\u00bb\3\2"+
		"\2\2\24\u00c6\3\2\2\2\26\u00d1\3\2\2\2\30\u00d9\3\2\2\2\32\u00e7\3\2\2"+
		"\2\34\u00f3\3\2\2\2\36\u0131\3\2\2\2 \u0133\3\2\2\2\"\u013f\3\2\2\2$\u014b"+
		"\3\2\2\2&\u0152\3\2\2\2(\u0157\3\2\2\2*\u015c\3\2\2\2,\u0165\3\2\2\2."+
		"\u016a\3\2\2\2\60\u016f\3\2\2\2\62\u0174\3\2\2\2\64\u0179\3\2\2\2\66\u017e"+
		"\3\2\2\28\u0183\3\2\2\2:\u0188\3\2\2\2<\u018d\3\2\2\2>\u0192\3\2\2\2@"+
		"\u0197\3\2\2\2B\u019c\3\2\2\2D\u01a1\3\2\2\2F\u01a6\3\2\2\2H\u01ab\3\2"+
		"\2\2J\u01b0\3\2\2\2L\u01b5\3\2\2\2N\u01ba\3\2\2\2P\u01bf\3\2\2\2R\u01c5"+
		"\3\2\2\2T\u01d1\3\2\2\2V\u01df\3\2\2\2XY\7\3\2\2YZ\7\61\2\2Z^\7\4\2\2"+
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
		"\t\3\2\2\2\u0094\u0099\5\16\b\2\u0095\u0099\5\"\22\2\u0096\u0099\5\36"+
		"\20\2\u0097\u0099\5T+\2\u0098\u0094\3\2\2\2\u0098\u0095\3\2\2\2\u0098"+
		"\u0096\3\2\2\2\u0098\u0097\3\2\2\2\u0099\u009a\3\2\2\2\u009a\u009b\7\4"+
		"\2\2\u009b\u00a1\3\2\2\2\u009c\u00a1\5\34\17\2\u009d\u00a1\5\32\16\2\u009e"+
		"\u00a1\5\30\r\2\u009f\u00a1\5\26\f\2\u00a0\u0098\3\2\2\2\u00a0\u009c\3"+
		"\2\2\2\u00a0\u009d\3\2\2\2\u00a0\u009e\3\2\2\2\u00a0\u009f\3\2\2\2\u00a1"+
		"\13\3\2\2\2\u00a2\u00a3\7\5\2\2\u00a3\u00a4\7\60\2\2\u00a4\u00a5\7\61"+
		"\2\2\u00a5\u00a6\7\4\2\2\u00a6\r\3\2\2\2\u00a7\u00a8\7\61\2\2\u00a8\u00ac"+
		"\7\r\2\2\u00a9\u00ad\7\64\2\2\u00aa\u00ad\5\24\13\2\u00ab\u00ad\5\22\n"+
		"\2\u00ac\u00a9\3\2\2\2\u00ac\u00aa\3\2\2\2\u00ac\u00ab\3\2\2\2\u00ad\17"+
		"\3\2\2\2\u00ae\u00b3\5\16\b\2\u00af\u00b3\5\"\22\2\u00b0\u00b3\5\36\20"+
		"\2\u00b1\u00b3\5T+\2\u00b2\u00ae\3\2\2\2\u00b2\u00af\3\2\2\2\u00b2\u00b0"+
		"\3\2\2\2\u00b2\u00b1\3\2\2\2\u00b3\u00b4\3\2\2\2\u00b4\u00b5\7\4\2\2\u00b5"+
		"\u00ba\3\2\2\2\u00b6\u00ba\5\34\17\2\u00b7\u00ba\5\32\16\2\u00b8\u00ba"+
		"\5\26\f\2\u00b9\u00b2\3\2\2\2\u00b9\u00b6\3\2\2\2\u00b9\u00b7\3\2\2\2"+
		"\u00b9\u00b8\3\2\2\2\u00ba\21\3\2\2\2\u00bb\u00bc\7\16\2\2\u00bc\u00c1"+
		"\5\24\13\2\u00bd\u00be\7\17\2\2\u00be\u00c0\5\24\13\2\u00bf\u00bd\3\2"+
		"\2\2\u00c0\u00c3\3\2\2\2\u00c1\u00bf\3\2\2\2\u00c1\u00c2\3\2\2\2\u00c2"+
		"\u00c4\3\2\2\2\u00c3\u00c1\3\2\2\2\u00c4\u00c5\7\20\2\2\u00c5\23\3\2\2"+
		"\2\u00c6\u00c7\7\16\2\2\u00c7\u00cc\7\64\2\2\u00c8\u00c9\7\17\2\2\u00c9"+
		"\u00cb\7\64\2\2\u00ca\u00c8\3\2\2\2\u00cb\u00ce\3\2\2\2\u00cc\u00ca\3"+
		"\2\2\2\u00cc\u00cd\3\2\2\2\u00cd\u00cf\3\2\2\2\u00ce\u00cc\3\2\2\2\u00cf"+
		"\u00d0\7\20\2\2\u00d0\25\3\2\2\2\u00d1\u00d3\7\21\2\2\u00d2\u00d4\n\2"+
		"\2\2\u00d3\u00d2\3\2\2\2\u00d4\u00d5\3\2\2\2\u00d5\u00d3\3\2\2\2\u00d5"+
		"\u00d6\3\2\2\2\u00d6\u00d7\3\2\2\2\u00d7\u00d8\7\21\2\2\u00d8\27\3\2\2"+
		"\2\u00d9\u00da\7\22\2\2\u00da\u00db\7\7\2\2\u00db\u00dc\7\61\2\2\u00dc"+
		"\u00dd\7\23\2\2\u00dd\u00de\5L\'\2\u00de\u00df\7\b\2\2\u00df\u00e1\7\n"+
		"\2\2\u00e0\u00e2\5\20\t\2\u00e1\u00e0\3\2\2\2\u00e2\u00e3\3\2\2\2\u00e3"+
		"\u00e1\3\2\2\2\u00e3\u00e4\3\2\2\2\u00e4\u00e5\3\2\2\2\u00e5\u00e6\7\13"+
		"\2\2\u00e6\31\3\2\2\2\u00e7\u00e8\7\24\2\2\u00e8\u00e9\7\7\2\2\u00e9\u00ea"+
		"\5L\'\2\u00ea\u00eb\7\b\2\2\u00eb\u00ed\7\n\2\2\u00ec\u00ee\5\n\6\2\u00ed"+
		"\u00ec\3\2\2\2\u00ee\u00ef\3\2\2\2\u00ef\u00ed\3\2\2\2\u00ef\u00f0\3\2"+
		"\2\2\u00f0\u00f1\3\2\2\2\u00f1\u00f2\7\13\2\2\u00f2\33\3\2\2\2\u00f3\u00f4"+
		"\7\25\2\2\u00f4\u00f5\7\7\2\2\u00f5\u00f6\5L\'\2\u00f6\u00f7\7\b\2\2\u00f7"+
		"\u00f9\7\n\2\2\u00f8\u00fa\5\n\6\2\u00f9\u00f8\3\2\2\2\u00fa\u00fb\3\2"+
		"\2\2\u00fb\u00f9\3\2\2\2\u00fb\u00fc\3\2\2\2\u00fc\u00fd\3\2\2\2\u00fd"+
		"\u010d\7\13\2\2\u00fe\u00ff\7\26\2\2\u00ff\u0100\7\25\2\2\u0100\u0101"+
		"\7\7\2\2\u0101\u0102\5L\'\2\u0102\u0103\7\b\2\2\u0103\u0105\7\n\2\2\u0104"+
		"\u0106\5\n\6\2\u0105\u0104\3\2\2\2\u0106\u0107\3\2\2\2\u0107\u0105\3\2"+
		"\2\2\u0107\u0108\3\2\2\2\u0108\u0109\3\2\2\2\u0109\u010a\7\13\2\2\u010a"+
		"\u010c\3\2\2\2\u010b\u00fe\3\2\2\2\u010c\u010f\3\2\2\2\u010d\u010b\3\2"+
		"\2\2\u010d\u010e\3\2\2\2\u010e\u0119\3\2\2\2\u010f\u010d\3\2\2\2\u0110"+
		"\u0111\7\26\2\2\u0111\u0113\7\n\2\2\u0112\u0114\5\n\6\2\u0113\u0112\3"+
		"\2\2\2\u0114\u0115\3\2\2\2\u0115\u0113\3\2\2\2\u0115\u0116\3\2\2\2\u0116"+
		"\u0117\3\2\2\2\u0117\u0118\7\13\2\2\u0118\u011a\3\2\2\2\u0119\u0110\3"+
		"\2\2\2\u0119\u011a\3\2\2\2\u011a\35\3\2\2\2\u011b\u0132\5 \21\2\u011c"+
		"\u0132\5\"\22\2\u011d\u0132\5$\23\2\u011e\u0132\5&\24\2\u011f\u0132\5"+
		"(\25\2\u0120\u0132\5*\26\2\u0121\u0132\5,\27\2\u0122\u0132\5.\30\2\u0123"+
		"\u0132\5\60\31\2\u0124\u0132\5\62\32\2\u0125\u0132\5\64\33\2\u0126\u0132"+
		"\5\66\34\2\u0127\u0132\58\35\2\u0128\u0132\5:\36\2\u0129\u0132\5> \2\u012a"+
		"\u0132\5<\37\2\u012b\u0132\5@!\2\u012c\u0132\5B\"\2\u012d\u0132\5D#\2"+
		"\u012e\u0132\5F$\2\u012f\u0132\5H%\2\u0130\u0132\5J&\2\u0131\u011b\3\2"+
		"\2\2\u0131\u011c\3\2\2\2\u0131\u011d\3\2\2\2\u0131\u011e\3\2\2\2\u0131"+
		"\u011f\3\2\2\2\u0131\u0120\3\2\2\2\u0131\u0121\3\2\2\2\u0131\u0122\3\2"+
		"\2\2\u0131\u0123\3\2\2\2\u0131\u0124\3\2\2\2\u0131\u0125\3\2\2\2\u0131"+
		"\u0126\3\2\2\2\u0131\u0127\3\2\2\2\u0131\u0128\3\2\2\2\u0131\u0129\3\2"+
		"\2\2\u0131\u012a\3\2\2\2\u0131\u012b\3\2\2\2\u0131\u012c\3\2\2\2\u0131"+
		"\u012d\3\2\2\2\u0131\u012e\3\2\2\2\u0131\u012f\3\2\2\2\u0131\u0130\3\2"+
		"\2\2\u0132\37\3\2\2\2\u0133\u0134\7\27\2\2\u0134\u0135\7\7\2\2\u0135\u013a"+
		"\7\61\2\2\u0136\u0137\7\17\2\2\u0137\u0139\7\61\2\2\u0138\u0136\3\2\2"+
		"\2\u0139\u013c\3\2\2\2\u013a\u0138\3\2\2\2\u013a\u013b\3\2\2\2\u013b\u013d"+
		"\3\2\2\2\u013c\u013a\3\2\2\2\u013d\u013e\7\b\2\2\u013e!\3\2\2\2\u013f"+
		"\u0140\7\30\2\2\u0140\u0141\7\7\2\2\u0141\u0146\5L\'\2\u0142\u0143\7\17"+
		"\2\2\u0143\u0145\5L\'\2\u0144\u0142\3\2\2\2\u0145\u0148\3\2\2\2\u0146"+
		"\u0144\3\2\2\2\u0146\u0147\3\2\2\2\u0147\u0149\3\2\2\2\u0148\u0146\3\2"+
		"\2\2\u0149\u014a\7\b\2\2\u014a#\3\2\2\2\u014b\u014c\7\31\2\2\u014c\u014d"+
		"\7\7\2\2\u014d\u014e\5L\'\2\u014e\u014f\7\17\2\2\u014f\u0150\5L\'\2\u0150"+
		"\u0151\7\b\2\2\u0151%\3\2\2\2\u0152\u0153\7\32\2\2\u0153\u0154\7\7\2\2"+
		"\u0154\u0155\5L\'\2\u0155\u0156\7\b\2\2\u0156\'\3\2\2\2\u0157\u0158\7"+
		"\33\2\2\u0158\u0159\7\7\2\2\u0159\u015a\5L\'\2\u015a\u015b\7\b\2\2\u015b"+
		")\3\2\2\2\u015c\u015d\7\34\2\2\u015d\u015e\7\7\2\2\u015e\u0161\5L\'\2"+
		"\u015f\u0160\7\17\2\2\u0160\u0162\5L\'\2\u0161\u015f\3\2\2\2\u0161\u0162"+
		"\3\2\2\2\u0162\u0163\3\2\2\2\u0163\u0164\7\b\2\2\u0164+\3\2\2\2\u0165"+
		"\u0166\7\35\2\2\u0166\u0167\7\7\2\2\u0167\u0168\5L\'\2\u0168\u0169\7\b"+
		"\2\2\u0169-\3\2\2\2\u016a\u016b\7\36\2\2\u016b\u016c\7\7\2\2\u016c\u016d"+
		"\5L\'\2\u016d\u016e\7\b\2\2\u016e/\3\2\2\2\u016f\u0170\7\37\2\2\u0170"+
		"\u0171\7\7\2\2\u0171\u0172\5L\'\2\u0172\u0173\7\b\2\2\u0173\61\3\2\2\2"+
		"\u0174\u0175\7 \2\2\u0175\u0176\7\7\2\2\u0176\u0177\5L\'\2\u0177\u0178"+
		"\7\b\2\2\u0178\63\3\2\2\2\u0179\u017a\7!\2\2\u017a\u017b\7\7\2\2\u017b"+
		"\u017c\5L\'\2\u017c\u017d\7\b\2\2\u017d\65\3\2\2\2\u017e\u017f\7\"\2\2"+
		"\u017f\u0180\7\7\2\2\u0180\u0181\5L\'\2\u0181\u0182\7\b\2\2\u0182\67\3"+
		"\2\2\2\u0183\u0184\7#\2\2\u0184\u0185\7\7\2\2\u0185\u0186\5L\'\2\u0186"+
		"\u0187\7\b\2\2\u01879\3\2\2\2\u0188\u0189\7$\2\2\u0189\u018a\7\7\2\2\u018a"+
		"\u018b\5L\'\2\u018b\u018c\7\b\2\2\u018c;\3\2\2\2\u018d\u018e\7%\2\2\u018e"+
		"\u018f\7\7\2\2\u018f\u0190\5L\'\2\u0190\u0191\7\b\2\2\u0191=\3\2\2\2\u0192"+
		"\u0193\7&\2\2\u0193\u0194\7\7\2\2\u0194\u0195\5L\'\2\u0195\u0196\7\b\2"+
		"\2\u0196?\3\2\2\2\u0197\u0198\7\'\2\2\u0198\u0199\7\7\2\2\u0199\u019a"+
		"\5L\'\2\u019a\u019b\7\b\2\2\u019bA\3\2\2\2\u019c\u019d\7(\2\2\u019d\u019e"+
		"\7\7\2\2\u019e\u019f\5L\'\2\u019f\u01a0\7\b\2\2\u01a0C\3\2\2\2\u01a1\u01a2"+
		"\7)\2\2\u01a2\u01a3\7\7\2\2\u01a3\u01a4\5L\'\2\u01a4\u01a5\7\b\2\2\u01a5"+
		"E\3\2\2\2\u01a6\u01a7\7*\2\2\u01a7\u01a8\7\7\2\2\u01a8\u01a9\5L\'\2\u01a9"+
		"\u01aa\7\b\2\2\u01aaG\3\2\2\2\u01ab\u01ac\7+\2\2\u01ac\u01ad\7\7\2\2\u01ad"+
		"\u01ae\5L\'\2\u01ae\u01af\7\b\2\2\u01afI\3\2\2\2\u01b0\u01b1\7,\2\2\u01b1"+
		"\u01b2\7\7\2\2\u01b2\u01b3\5L\'\2\u01b3\u01b4\7\b\2\2\u01b4K\3\2\2\2\u01b5"+
		"\u01b8\5N(\2\u01b6\u01b7\7;\2\2\u01b7\u01b9\5N(\2\u01b8\u01b6\3\2\2\2"+
		"\u01b8\u01b9\3\2\2\2\u01b9M\3\2\2\2\u01ba\u01bd\5P)\2\u01bb\u01bc\79\2"+
		"\2\u01bc\u01be\5N(\2\u01bd\u01bb\3\2\2\2\u01bd\u01be\3\2\2\2\u01beO\3"+
		"\2\2\2\u01bf\u01c2\5R*\2\u01c0\u01c1\7:\2\2\u01c1\u01c3\5P)\2\u01c2\u01c0"+
		"\3\2\2\2\u01c2\u01c3\3\2\2\2\u01c3Q\3\2\2\2\u01c4\u01c6\79\2\2\u01c5\u01c4"+
		"\3\2\2\2\u01c5\u01c6\3\2\2\2\u01c6\u01cf\3\2\2\2\u01c7\u01c8\7\7\2\2\u01c8"+
		"\u01c9\5L\'\2\u01c9\u01ca\7\b\2\2\u01ca\u01d0\3\2\2\2\u01cb\u01d0\5V,"+
		"\2\u01cc\u01d0\5\36\20\2\u01cd\u01d0\5T+\2\u01ce\u01d0\7\64\2\2\u01cf"+
		"\u01c7\3\2\2\2\u01cf\u01cb\3\2\2\2\u01cf\u01cc\3\2\2\2\u01cf\u01cd\3\2"+
		"\2\2\u01cf\u01ce\3\2\2\2\u01d0S\3\2\2\2\u01d1\u01d2\7\61\2\2\u01d2\u01db"+
		"\7\7\2\2\u01d3\u01d8\5L\'\2\u01d4\u01d5\7\17\2\2\u01d5\u01d7\5L\'\2\u01d6"+
		"\u01d4\3\2\2\2\u01d7\u01da\3\2\2\2\u01d8\u01d6\3\2\2\2\u01d8\u01d9\3\2"+
		"\2\2\u01d9\u01dc\3\2\2\2\u01da\u01d8\3\2\2\2\u01db\u01d3\3\2\2\2\u01db"+
		"\u01dc\3\2\2\2\u01dc\u01dd\3\2\2\2\u01dd\u01de\7\b\2\2\u01deU\3\2\2\2"+
		"\u01df\u01e0\7\61\2\2\u01e0\u01e1\7\16\2\2\u01e1\u01e2\5L\'\2\u01e2\u01e7"+
		"\7\20\2\2\u01e3\u01e4\7\16\2\2\u01e4\u01e5\5L\'\2\u01e5\u01e6\7\20\2\2"+
		"\u01e6\u01e8\3\2\2\2\u01e7\u01e3\3\2\2\2\u01e7\u01e8\3\2\2\2\u01e8W\3"+
		"\2\2\2#^du~\u0085\u0090\u0098\u00a0\u00ac\u00b2\u00b9\u00c1\u00cc\u00d5"+
		"\u00e3\u00ef\u00fb\u0107\u010d\u0115\u0119\u0131\u013a\u0146\u0161\u01b8"+
		"\u01bd\u01c2\u01c5\u01cf\u01d8\u01db\u01e7";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}