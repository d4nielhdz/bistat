// Code generated from parser/Bistat.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Bistat

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseBistatListener is a complete listener for a parse tree produced by BistatParser.
type BaseBistatListener struct{}

var _ BistatListener = &BaseBistatListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBistatListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBistatListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBistatListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBistatListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseBistatListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseBistatListener) ExitProgram(ctx *ProgramContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseBistatListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseBistatListener) ExitStmt(ctx *StmtContext) {}

// EnterNested_stmt is called when production nested_stmt is entered.
func (s *BaseBistatListener) EnterNested_stmt(ctx *Nested_stmtContext) {}

// ExitNested_stmt is called when production nested_stmt is exited.
func (s *BaseBistatListener) ExitNested_stmt(ctx *Nested_stmtContext) {}

// EnterPrint is called when production print is entered.
func (s *BaseBistatListener) EnterPrint(ctx *PrintContext) {}

// ExitPrint is called when production print is exited.
func (s *BaseBistatListener) ExitPrint(ctx *PrintContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBistatListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBistatListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMain is called when production main is entered.
func (s *BaseBistatListener) EnterMain(ctx *MainContext) {}

// ExitMain is called when production main is exited.
func (s *BaseBistatListener) ExitMain(ctx *MainContext) {}

// EnterFunc_def is called when production func_def is entered.
func (s *BaseBistatListener) EnterFunc_def(ctx *Func_defContext) {}

// ExitFunc_def is called when production func_def is exited.
func (s *BaseBistatListener) ExitFunc_def(ctx *Func_defContext) {}

// EnterVar_declaration is called when production var_declaration is entered.
func (s *BaseBistatListener) EnterVar_declaration(ctx *Var_declarationContext) {}

// ExitVar_declaration is called when production var_declaration is exited.
func (s *BaseBistatListener) ExitVar_declaration(ctx *Var_declarationContext) {}

// EnterParam_declaration is called when production param_declaration is entered.
func (s *BaseBistatListener) EnterParam_declaration(ctx *Param_declarationContext) {}

// ExitParam_declaration is called when production param_declaration is exited.
func (s *BaseBistatListener) ExitParam_declaration(ctx *Param_declarationContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseBistatListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseBistatListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterMatrix_assignment is called when production matrix_assignment is entered.
func (s *BaseBistatListener) EnterMatrix_assignment(ctx *Matrix_assignmentContext) {}

// ExitMatrix_assignment is called when production matrix_assignment is exited.
func (s *BaseBistatListener) ExitMatrix_assignment(ctx *Matrix_assignmentContext) {}

// EnterList_assignment is called when production list_assignment is entered.
func (s *BaseBistatListener) EnterList_assignment(ctx *List_assignmentContext) {}

// ExitList_assignment is called when production list_assignment is exited.
func (s *BaseBistatListener) ExitList_assignment(ctx *List_assignmentContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseBistatListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseBistatListener) ExitComment(ctx *CommentContext) {}

// EnterFor is called when production for is entered.
func (s *BaseBistatListener) EnterFor(ctx *ForContext) {}

// ExitFor is called when production for is exited.
func (s *BaseBistatListener) ExitFor(ctx *ForContext) {}

// EnterWhile is called when production while is entered.
func (s *BaseBistatListener) EnterWhile(ctx *WhileContext) {}

// ExitWhile is called when production while is exited.
func (s *BaseBistatListener) ExitWhile(ctx *WhileContext) {}

// EnterConditional is called when production conditional is entered.
func (s *BaseBistatListener) EnterConditional(ctx *ConditionalContext) {}

// ExitConditional is called when production conditional is exited.
func (s *BaseBistatListener) ExitConditional(ctx *ConditionalContext) {}

// EnterInput_read is called when production input_read is entered.
func (s *BaseBistatListener) EnterInput_read(ctx *Input_readContext) {}

// ExitInput_read is called when production input_read is exited.
func (s *BaseBistatListener) ExitInput_read(ctx *Input_readContext) {}

// EnterList_add is called when production list_add is entered.
func (s *BaseBistatListener) EnterList_add(ctx *List_addContext) {}

// ExitList_add is called when production list_add is exited.
func (s *BaseBistatListener) ExitList_add(ctx *List_addContext) {}

// EnterList_pop is called when production list_pop is entered.
func (s *BaseBistatListener) EnterList_pop(ctx *List_popContext) {}

// ExitList_pop is called when production list_pop is exited.
func (s *BaseBistatListener) ExitList_pop(ctx *List_popContext) {}

// EnterLength is called when production length is entered.
func (s *BaseBistatListener) EnterLength(ctx *LengthContext) {}

// ExitLength is called when production length is exited.
func (s *BaseBistatListener) ExitLength(ctx *LengthContext) {}

// EnterRange is called when production range is entered.
func (s *BaseBistatListener) EnterRange(ctx *RangeContext) {}

// ExitRange is called when production range is exited.
func (s *BaseBistatListener) ExitRange(ctx *RangeContext) {}

// EnterPlot is called when production plot is entered.
func (s *BaseBistatListener) EnterPlot(ctx *PlotContext) {}

// ExitPlot is called when production plot is exited.
func (s *BaseBistatListener) ExitPlot(ctx *PlotContext) {}

// EnterSum is called when production sum is entered.
func (s *BaseBistatListener) EnterSum(ctx *SumContext) {}

// ExitSum is called when production sum is exited.
func (s *BaseBistatListener) ExitSum(ctx *SumContext) {}

// EnterMin is called when production min is entered.
func (s *BaseBistatListener) EnterMin(ctx *MinContext) {}

// ExitMin is called when production min is exited.
func (s *BaseBistatListener) ExitMin(ctx *MinContext) {}

// EnterProd is called when production prod is entered.
func (s *BaseBistatListener) EnterProd(ctx *ProdContext) {}

// ExitProd is called when production prod is exited.
func (s *BaseBistatListener) ExitProd(ctx *ProdContext) {}

// EnterAvg is called when production avg is entered.
func (s *BaseBistatListener) EnterAvg(ctx *AvgContext) {}

// ExitAvg is called when production avg is exited.
func (s *BaseBistatListener) ExitAvg(ctx *AvgContext) {}

// EnterS_mode is called when production s_mode is entered.
func (s *BaseBistatListener) EnterS_mode(ctx *S_modeContext) {}

// ExitS_mode is called when production s_mode is exited.
func (s *BaseBistatListener) ExitS_mode(ctx *S_modeContext) {}

// EnterMedian is called when production median is entered.
func (s *BaseBistatListener) EnterMedian(ctx *MedianContext) {}

// ExitMedian is called when production median is exited.
func (s *BaseBistatListener) ExitMedian(ctx *MedianContext) {}

// EnterSin is called when production sin is entered.
func (s *BaseBistatListener) EnterSin(ctx *SinContext) {}

// ExitSin is called when production sin is exited.
func (s *BaseBistatListener) ExitSin(ctx *SinContext) {}

// EnterTan is called when production tan is entered.
func (s *BaseBistatListener) EnterTan(ctx *TanContext) {}

// ExitTan is called when production tan is exited.
func (s *BaseBistatListener) ExitTan(ctx *TanContext) {}

// EnterCos is called when production cos is entered.
func (s *BaseBistatListener) EnterCos(ctx *CosContext) {}

// ExitCos is called when production cos is exited.
func (s *BaseBistatListener) ExitCos(ctx *CosContext) {}

// EnterSort is called when production sort is entered.
func (s *BaseBistatListener) EnterSort(ctx *SortContext) {}

// ExitSort is called when production sort is exited.
func (s *BaseBistatListener) ExitSort(ctx *SortContext) {}

// EnterSqrt is called when production sqrt is entered.
func (s *BaseBistatListener) EnterSqrt(ctx *SqrtContext) {}

// ExitSqrt is called when production sqrt is exited.
func (s *BaseBistatListener) ExitSqrt(ctx *SqrtContext) {}

// EnterFloor is called when production floor is entered.
func (s *BaseBistatListener) EnterFloor(ctx *FloorContext) {}

// ExitFloor is called when production floor is exited.
func (s *BaseBistatListener) ExitFloor(ctx *FloorContext) {}

// EnterCeil is called when production ceil is entered.
func (s *BaseBistatListener) EnterCeil(ctx *CeilContext) {}

// ExitCeil is called when production ceil is exited.
func (s *BaseBistatListener) ExitCeil(ctx *CeilContext) {}

// EnterAbs is called when production abs is entered.
func (s *BaseBistatListener) EnterAbs(ctx *AbsContext) {}

// ExitAbs is called when production abs is exited.
func (s *BaseBistatListener) ExitAbs(ctx *AbsContext) {}

// EnterNot is called when production not is entered.
func (s *BaseBistatListener) EnterNot(ctx *NotContext) {}

// ExitNot is called when production not is exited.
func (s *BaseBistatListener) ExitNot(ctx *NotContext) {}

// EnterSpecial_function is called when production special_function is entered.
func (s *BaseBistatListener) EnterSpecial_function(ctx *Special_functionContext) {}

// ExitSpecial_function is called when production special_function is exited.
func (s *BaseBistatListener) ExitSpecial_function(ctx *Special_functionContext) {}

// EnterExp is called when production exp is entered.
func (s *BaseBistatListener) EnterExp(ctx *ExpContext) {}

// ExitExp is called when production exp is exited.
func (s *BaseBistatListener) ExitExp(ctx *ExpContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseBistatListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseBistatListener) ExitTerm(ctx *TermContext) {}

// EnterFactor is called when production factor is entered.
func (s *BaseBistatListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BaseBistatListener) ExitFactor(ctx *FactorContext) {}

// EnterFunction_call is called when production function_call is entered.
func (s *BaseBistatListener) EnterFunction_call(ctx *Function_callContext) {}

// ExitFunction_call is called when production function_call is exited.
func (s *BaseBistatListener) ExitFunction_call(ctx *Function_callContext) {}

// EnterIndexing is called when production indexing is entered.
func (s *BaseBistatListener) EnterIndexing(ctx *IndexingContext) {}

// ExitIndexing is called when production indexing is exited.
func (s *BaseBistatListener) ExitIndexing(ctx *IndexingContext) {}
