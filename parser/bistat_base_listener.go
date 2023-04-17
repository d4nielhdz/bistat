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

// EnterVarDeclaration is called when production varDeclaration is entered.
func (s *BaseBistatListener) EnterVarDeclaration(ctx *VarDeclarationContext) {}

// ExitVarDeclaration is called when production varDeclaration is exited.
func (s *BaseBistatListener) ExitVarDeclaration(ctx *VarDeclarationContext) {}

// EnterFuncDef is called when production funcDef is entered.
func (s *BaseBistatListener) EnterFuncDef(ctx *FuncDefContext) {}

// ExitFuncDef is called when production funcDef is exited.
func (s *BaseBistatListener) ExitFuncDef(ctx *FuncDefContext) {}

// EnterMain is called when production main is entered.
func (s *BaseBistatListener) EnterMain(ctx *MainContext) {}

// ExitMain is called when production main is exited.
func (s *BaseBistatListener) ExitMain(ctx *MainContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseBistatListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseBistatListener) ExitStmt(ctx *StmtContext) {}

// EnterParamDeclaration is called when production paramDeclaration is entered.
func (s *BaseBistatListener) EnterParamDeclaration(ctx *ParamDeclarationContext) {}

// ExitParamDeclaration is called when production paramDeclaration is exited.
func (s *BaseBistatListener) ExitParamDeclaration(ctx *ParamDeclarationContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseBistatListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseBistatListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterNested_stmt is called when production nested_stmt is entered.
func (s *BaseBistatListener) EnterNested_stmt(ctx *Nested_stmtContext) {}

// ExitNested_stmt is called when production nested_stmt is exited.
func (s *BaseBistatListener) ExitNested_stmt(ctx *Nested_stmtContext) {}

// EnterMatrixAssignment is called when production matrixAssignment is entered.
func (s *BaseBistatListener) EnterMatrixAssignment(ctx *MatrixAssignmentContext) {}

// ExitMatrixAssignment is called when production matrixAssignment is exited.
func (s *BaseBistatListener) ExitMatrixAssignment(ctx *MatrixAssignmentContext) {}

// EnterListAssignment is called when production listAssignment is entered.
func (s *BaseBistatListener) EnterListAssignment(ctx *ListAssignmentContext) {}

// ExitListAssignment is called when production listAssignment is exited.
func (s *BaseBistatListener) ExitListAssignment(ctx *ListAssignmentContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseBistatListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseBistatListener) ExitComment(ctx *CommentContext) {}

// EnterForLoop is called when production forLoop is entered.
func (s *BaseBistatListener) EnterForLoop(ctx *ForLoopContext) {}

// ExitForLoop is called when production forLoop is exited.
func (s *BaseBistatListener) ExitForLoop(ctx *ForLoopContext) {}

// EnterWhileLoop is called when production whileLoop is entered.
func (s *BaseBistatListener) EnterWhileLoop(ctx *WhileLoopContext) {}

// ExitWhileLoop is called when production whileLoop is exited.
func (s *BaseBistatListener) ExitWhileLoop(ctx *WhileLoopContext) {}

// EnterConditional is called when production conditional is entered.
func (s *BaseBistatListener) EnterConditional(ctx *ConditionalContext) {}

// ExitConditional is called when production conditional is exited.
func (s *BaseBistatListener) ExitConditional(ctx *ConditionalContext) {}

// EnterReturnStmt is called when production returnStmt is entered.
func (s *BaseBistatListener) EnterReturnStmt(ctx *ReturnStmtContext) {}

// ExitReturnStmt is called when production returnStmt is exited.
func (s *BaseBistatListener) ExitReturnStmt(ctx *ReturnStmtContext) {}

// EnterSpecialFunction is called when production specialFunction is entered.
func (s *BaseBistatListener) EnterSpecialFunction(ctx *SpecialFunctionContext) {}

// ExitSpecialFunction is called when production specialFunction is exited.
func (s *BaseBistatListener) ExitSpecialFunction(ctx *SpecialFunctionContext) {}

// EnterInputRead is called when production inputRead is entered.
func (s *BaseBistatListener) EnterInputRead(ctx *InputReadContext) {}

// ExitInputRead is called when production inputRead is exited.
func (s *BaseBistatListener) ExitInputRead(ctx *InputReadContext) {}

// EnterPrint is called when production print is entered.
func (s *BaseBistatListener) EnterPrint(ctx *PrintContext) {}

// ExitPrint is called when production print is exited.
func (s *BaseBistatListener) ExitPrint(ctx *PrintContext) {}

// EnterListAdd is called when production listAdd is entered.
func (s *BaseBistatListener) EnterListAdd(ctx *ListAddContext) {}

// ExitListAdd is called when production listAdd is exited.
func (s *BaseBistatListener) ExitListAdd(ctx *ListAddContext) {}

// EnterListPop is called when production listPop is entered.
func (s *BaseBistatListener) EnterListPop(ctx *ListPopContext) {}

// ExitListPop is called when production listPop is exited.
func (s *BaseBistatListener) ExitListPop(ctx *ListPopContext) {}

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

// EnterSMode is called when production sMode is entered.
func (s *BaseBistatListener) EnterSMode(ctx *SModeContext) {}

// ExitSMode is called when production sMode is exited.
func (s *BaseBistatListener) ExitSMode(ctx *SModeContext) {}

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

// EnterExpression is called when production expression is entered.
func (s *BaseBistatListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBistatListener) ExitExpression(ctx *ExpressionContext) {}

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

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseBistatListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseBistatListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterIndexing is called when production indexing is entered.
func (s *BaseBistatListener) EnterIndexing(ctx *IndexingContext) {}

// ExitIndexing is called when production indexing is exited.
func (s *BaseBistatListener) ExitIndexing(ctx *IndexingContext) {}
