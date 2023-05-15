// Code generated from parser/Bistat.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Bistat

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BistatListener is a complete listener for a parse tree produced by BistatParser.
type BistatListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterVarDeclaration is called when entering the varDeclaration production.
	EnterVarDeclaration(c *VarDeclarationContext)

	// EnterVar_type is called when entering the var_type production.
	EnterVar_type(c *Var_typeContext)

	// EnterFuncDef is called when entering the funcDef production.
	EnterFuncDef(c *FuncDefContext)

	// EnterFuncBlockStart is called when entering the funcBlockStart production.
	EnterFuncBlockStart(c *FuncBlockStartContext)

	// EnterFuncBlockEnd is called when entering the funcBlockEnd production.
	EnterFuncBlockEnd(c *FuncBlockEndContext)

	// EnterParamDeclaration is called when entering the paramDeclaration production.
	EnterParamDeclaration(c *ParamDeclarationContext)

	// EnterMain is called when entering the main production.
	EnterMain(c *MainContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterMatrixAssignment is called when entering the matrixAssignment production.
	EnterMatrixAssignment(c *MatrixAssignmentContext)

	// EnterListAssignment is called when entering the listAssignment production.
	EnterListAssignment(c *ListAssignmentContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterForLoop is called when entering the forLoop production.
	EnterForLoop(c *ForLoopContext)

	// EnterForHeader is called when entering the forHeader production.
	EnterForHeader(c *ForHeaderContext)

	// EnterForExprEnd is called when entering the forExprEnd production.
	EnterForExprEnd(c *ForExprEndContext)

	// EnterWhileLoop is called when entering the whileLoop production.
	EnterWhileLoop(c *WhileLoopContext)

	// EnterWhileExprEnd is called when entering the whileExprEnd production.
	EnterWhileExprEnd(c *WhileExprEndContext)

	// EnterConditional is called when entering the conditional production.
	EnterConditional(c *ConditionalContext)

	// EnterIfStmt is called when entering the ifStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterElseIfStmt is called when entering the elseIfStmt production.
	EnterElseIfStmt(c *ElseIfStmtContext)

	// EnterCondExprEnd is called when entering the condExprEnd production.
	EnterCondExprEnd(c *CondExprEndContext)

	// EnterElseStmt is called when entering the elseStmt production.
	EnterElseStmt(c *ElseStmtContext)

	// EnterReturnStmt is called when entering the returnStmt production.
	EnterReturnStmt(c *ReturnStmtContext)

	// EnterSpecialFunction is called when entering the specialFunction production.
	EnterSpecialFunction(c *SpecialFunctionContext)

	// EnterInputRead is called when entering the inputRead production.
	EnterInputRead(c *InputReadContext)

	// EnterPrint is called when entering the print production.
	EnterPrint(c *PrintContext)

	// EnterListAdd is called when entering the listAdd production.
	EnterListAdd(c *ListAddContext)

	// EnterListPop is called when entering the listPop production.
	EnterListPop(c *ListPopContext)

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

	// EnterSMode is called when entering the sMode production.
	EnterSMode(c *SModeContext)

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

	// EnterUnaryMinus is called when entering the unaryMinus production.
	EnterUnaryMinus(c *UnaryMinusContext)

	// EnterNestedExpression is called when entering the nestedExpression production.
	EnterNestedExpression(c *NestedExpressionContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterIndexing is called when entering the indexing production.
	EnterIndexing(c *IndexingContext)

	// EnterVarCons is called when entering the varCons production.
	EnterVarCons(c *VarConsContext)

	// EnterOpSec is called when entering the opSec production.
	EnterOpSec(c *OpSecContext)

	// EnterOpFirst is called when entering the opFirst production.
	EnterOpFirst(c *OpFirstContext)

	// EnterLogicOperator is called when entering the logicOperator production.
	EnterLogicOperator(c *LogicOperatorContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitVarDeclaration is called when exiting the varDeclaration production.
	ExitVarDeclaration(c *VarDeclarationContext)

	// ExitVar_type is called when exiting the var_type production.
	ExitVar_type(c *Var_typeContext)

	// ExitFuncDef is called when exiting the funcDef production.
	ExitFuncDef(c *FuncDefContext)

	// ExitFuncBlockStart is called when exiting the funcBlockStart production.
	ExitFuncBlockStart(c *FuncBlockStartContext)

	// ExitFuncBlockEnd is called when exiting the funcBlockEnd production.
	ExitFuncBlockEnd(c *FuncBlockEndContext)

	// ExitParamDeclaration is called when exiting the paramDeclaration production.
	ExitParamDeclaration(c *ParamDeclarationContext)

	// ExitMain is called when exiting the main production.
	ExitMain(c *MainContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitMatrixAssignment is called when exiting the matrixAssignment production.
	ExitMatrixAssignment(c *MatrixAssignmentContext)

	// ExitListAssignment is called when exiting the listAssignment production.
	ExitListAssignment(c *ListAssignmentContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitForLoop is called when exiting the forLoop production.
	ExitForLoop(c *ForLoopContext)

	// ExitForHeader is called when exiting the forHeader production.
	ExitForHeader(c *ForHeaderContext)

	// ExitForExprEnd is called when exiting the forExprEnd production.
	ExitForExprEnd(c *ForExprEndContext)

	// ExitWhileLoop is called when exiting the whileLoop production.
	ExitWhileLoop(c *WhileLoopContext)

	// ExitWhileExprEnd is called when exiting the whileExprEnd production.
	ExitWhileExprEnd(c *WhileExprEndContext)

	// ExitConditional is called when exiting the conditional production.
	ExitConditional(c *ConditionalContext)

	// ExitIfStmt is called when exiting the ifStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitElseIfStmt is called when exiting the elseIfStmt production.
	ExitElseIfStmt(c *ElseIfStmtContext)

	// ExitCondExprEnd is called when exiting the condExprEnd production.
	ExitCondExprEnd(c *CondExprEndContext)

	// ExitElseStmt is called when exiting the elseStmt production.
	ExitElseStmt(c *ElseStmtContext)

	// ExitReturnStmt is called when exiting the returnStmt production.
	ExitReturnStmt(c *ReturnStmtContext)

	// ExitSpecialFunction is called when exiting the specialFunction production.
	ExitSpecialFunction(c *SpecialFunctionContext)

	// ExitInputRead is called when exiting the inputRead production.
	ExitInputRead(c *InputReadContext)

	// ExitPrint is called when exiting the print production.
	ExitPrint(c *PrintContext)

	// ExitListAdd is called when exiting the listAdd production.
	ExitListAdd(c *ListAddContext)

	// ExitListPop is called when exiting the listPop production.
	ExitListPop(c *ListPopContext)

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

	// ExitSMode is called when exiting the sMode production.
	ExitSMode(c *SModeContext)

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

	// ExitUnaryMinus is called when exiting the unaryMinus production.
	ExitUnaryMinus(c *UnaryMinusContext)

	// ExitNestedExpression is called when exiting the nestedExpression production.
	ExitNestedExpression(c *NestedExpressionContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitIndexing is called when exiting the indexing production.
	ExitIndexing(c *IndexingContext)

	// ExitVarCons is called when exiting the varCons production.
	ExitVarCons(c *VarConsContext)

	// ExitOpSec is called when exiting the opSec production.
	ExitOpSec(c *OpSecContext)

	// ExitOpFirst is called when exiting the opFirst production.
	ExitOpFirst(c *OpFirstContext)

	// ExitLogicOperator is called when exiting the logicOperator production.
	ExitLogicOperator(c *LogicOperatorContext)
}
