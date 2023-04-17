// Code generated from parser/Bistat.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Bistat

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BistatListener is a complete listener for a parse tree produced by BistatParser.
type BistatListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterVar_declaration is called when entering the var_declaration production.
	EnterVar_declaration(c *Var_declarationContext)

	// EnterFunc_def is called when entering the func_def production.
	EnterFunc_def(c *Func_defContext)

	// EnterMain is called when entering the main production.
	EnterMain(c *MainContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterParam_declaration is called when entering the param_declaration production.
	EnterParam_declaration(c *Param_declarationContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterNested_stmt is called when entering the nested_stmt production.
	EnterNested_stmt(c *Nested_stmtContext)

	// EnterMatrix_assignment is called when entering the matrix_assignment production.
	EnterMatrix_assignment(c *Matrix_assignmentContext)

	// EnterList_assignment is called when entering the list_assignment production.
	EnterList_assignment(c *List_assignmentContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterFor_loop is called when entering the for_loop production.
	EnterFor_loop(c *For_loopContext)

	// EnterWhile_loop is called when entering the while_loop production.
	EnterWhile_loop(c *While_loopContext)

	// EnterConditional is called when entering the conditional production.
	EnterConditional(c *ConditionalContext)

	// EnterSpecial_function is called when entering the special_function production.
	EnterSpecial_function(c *Special_functionContext)

	// EnterInput_read is called when entering the input_read production.
	EnterInput_read(c *Input_readContext)

	// EnterPrint is called when entering the print production.
	EnterPrint(c *PrintContext)

	// EnterList_add is called when entering the list_add production.
	EnterList_add(c *List_addContext)

	// EnterList_pop is called when entering the list_pop production.
	EnterList_pop(c *List_popContext)

	// EnterLength is called when entering the length production.
	EnterLength(c *LengthContext)

	// EnterRange is called when entering the range production.
	EnterRange(c *RangeContext)

	// EnterPlot is called when entering the plot production.
	EnterPlot(c *PlotContext)

	// EnterSum is called when entering the sum production.
	EnterSum(c *SumContext)

	// EnterMin is called when entering the min production.
	EnterMin(c *MinContext)

	// EnterProd is called when entering the prod production.
	EnterProd(c *ProdContext)

	// EnterAvg is called when entering the avg production.
	EnterAvg(c *AvgContext)

	// EnterS_mode is called when entering the s_mode production.
	EnterS_mode(c *S_modeContext)

	// EnterMedian is called when entering the median production.
	EnterMedian(c *MedianContext)

	// EnterSin is called when entering the sin production.
	EnterSin(c *SinContext)

	// EnterTan is called when entering the tan production.
	EnterTan(c *TanContext)

	// EnterCos is called when entering the cos production.
	EnterCos(c *CosContext)

	// EnterSort is called when entering the sort production.
	EnterSort(c *SortContext)

	// EnterSqrt is called when entering the sqrt production.
	EnterSqrt(c *SqrtContext)

	// EnterFloor is called when entering the floor production.
	EnterFloor(c *FloorContext)

	// EnterCeil is called when entering the ceil production.
	EnterCeil(c *CeilContext)

	// EnterAbs is called when entering the abs production.
	EnterAbs(c *AbsContext)

	// EnterNot is called when entering the not production.
	EnterNot(c *NotContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterExp is called when entering the exp production.
	EnterExp(c *ExpContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// EnterFunction_call is called when entering the function_call production.
	EnterFunction_call(c *Function_callContext)

	// EnterIndexing is called when entering the indexing production.
	EnterIndexing(c *IndexingContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitVar_declaration is called when exiting the var_declaration production.
	ExitVar_declaration(c *Var_declarationContext)

	// ExitFunc_def is called when exiting the func_def production.
	ExitFunc_def(c *Func_defContext)

	// ExitMain is called when exiting the main production.
	ExitMain(c *MainContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitParam_declaration is called when exiting the param_declaration production.
	ExitParam_declaration(c *Param_declarationContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitNested_stmt is called when exiting the nested_stmt production.
	ExitNested_stmt(c *Nested_stmtContext)

	// ExitMatrix_assignment is called when exiting the matrix_assignment production.
	ExitMatrix_assignment(c *Matrix_assignmentContext)

	// ExitList_assignment is called when exiting the list_assignment production.
	ExitList_assignment(c *List_assignmentContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitFor_loop is called when exiting the for_loop production.
	ExitFor_loop(c *For_loopContext)

	// ExitWhile_loop is called when exiting the while_loop production.
	ExitWhile_loop(c *While_loopContext)

	// ExitConditional is called when exiting the conditional production.
	ExitConditional(c *ConditionalContext)

	// ExitSpecial_function is called when exiting the special_function production.
	ExitSpecial_function(c *Special_functionContext)

	// ExitInput_read is called when exiting the input_read production.
	ExitInput_read(c *Input_readContext)

	// ExitPrint is called when exiting the print production.
	ExitPrint(c *PrintContext)

	// ExitList_add is called when exiting the list_add production.
	ExitList_add(c *List_addContext)

	// ExitList_pop is called when exiting the list_pop production.
	ExitList_pop(c *List_popContext)

	// ExitLength is called when exiting the length production.
	ExitLength(c *LengthContext)

	// ExitRange is called when exiting the range production.
	ExitRange(c *RangeContext)

	// ExitPlot is called when exiting the plot production.
	ExitPlot(c *PlotContext)

	// ExitSum is called when exiting the sum production.
	ExitSum(c *SumContext)

	// ExitMin is called when exiting the min production.
	ExitMin(c *MinContext)

	// ExitProd is called when exiting the prod production.
	ExitProd(c *ProdContext)

	// ExitAvg is called when exiting the avg production.
	ExitAvg(c *AvgContext)

	// ExitS_mode is called when exiting the s_mode production.
	ExitS_mode(c *S_modeContext)

	// ExitMedian is called when exiting the median production.
	ExitMedian(c *MedianContext)

	// ExitSin is called when exiting the sin production.
	ExitSin(c *SinContext)

	// ExitTan is called when exiting the tan production.
	ExitTan(c *TanContext)

	// ExitCos is called when exiting the cos production.
	ExitCos(c *CosContext)

	// ExitSort is called when exiting the sort production.
	ExitSort(c *SortContext)

	// ExitSqrt is called when exiting the sqrt production.
	ExitSqrt(c *SqrtContext)

	// ExitFloor is called when exiting the floor production.
	ExitFloor(c *FloorContext)

	// ExitCeil is called when exiting the ceil production.
	ExitCeil(c *CeilContext)

	// ExitAbs is called when exiting the abs production.
	ExitAbs(c *AbsContext)

	// ExitNot is called when exiting the not production.
	ExitNot(c *NotContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitExp is called when exiting the exp production.
	ExitExp(c *ExpContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)

	// ExitFunction_call is called when exiting the function_call production.
	ExitFunction_call(c *Function_callContext)

	// ExitIndexing is called when exiting the indexing production.
	ExitIndexing(c *IndexingContext)
}
