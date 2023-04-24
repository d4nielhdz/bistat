package main

import (
	parser "bistat/parser"
	utils "bistat/utils"
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type bistatListener struct {
	*parser.BaseBistatListener
}

type PCtx struct {
	scopes   []string
	funcDir  map[string]utils.RType
	varTable map[string]map[string]utils.RType
}

func NewPCtx() PCtx {
	return PCtx{
		scopes:   make([]string, 0),
		funcDir:  make(map[string]utils.RType),
		varTable: make(map[string]map[string]utils.RType),
	}
}

func (pCtx PCtx) GetCurrentScope() string {
	return pCtx.scopes[len(pCtx.scopes)-1]
}

var pCtx = NewPCtx()

func NewBistatListener() *bistatListener {
	return new(bistatListener)
}

func (l *bistatListener) EnterProgram(ctx *parser.ProgramContext) {
	fmt.Println("Entered program")
	fmt.Println(ctx.ID())
	scopes := append(pCtx.scopes, ctx.ID().GetText())
	pCtx.scopes = scopes
	pCtx.funcDir[ctx.ID().GetText()] = utils.NewRType(utils.Void)
}

func (l *bistatListener) EnterVarDeclaration(ctx *parser.VarDeclarationContext) {
	fmt.Println("Entered variable declaration")
	currScope := pCtx.GetCurrentScope()
	fmt.Println(ctx.ID())
	var vt = ctx.Var_type()
	var resolved = utils.ResolveType(vt)
	_, exists := pCtx.varTable[currScope][ctx.ID().GetText()]
	if exists {
		fmt.Println("variable already exists")
	} else {
		pCtx.varTable[currScope][ctx.ID().GetText()] = resolved
	}
	fmt.Println(vt.TYPE_PRIMITIVE())
	fmt.Println(vt.CARDINALITY())
	fmt.Println(resolved)

}

func (l *bistatListener) EnterParamDeclaration(ctx *parser.ParamDeclarationContext) {
	fmt.Println("Entered param declaration")
	fmt.Println(ctx.ID())
	var vt = ctx.Var_type()
	var resolved = utils.ResolveType(vt)
	fmt.Println(vt.TYPE_PRIMITIVE())
	fmt.Println(vt.CARDINALITY())
	fmt.Println(resolved)

}

// func (this *bistatListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
// 	fmt.Println("every rule")
// 	fmt.Println(ctx.G)
// 	fmt.Println("end")
// }

// func (l *bistatListener) EnterStmt(ctx *parser.StmtContext) {
// 	fmt.Println("Entered statement")
// }

func (l *bistatListener) EnterFuncDef(ctx *parser.FuncDefContext) {
	fmt.Println("Entered function definition")
	fmt.Println(ctx.ID())
	var vt = ctx.Var_type()
	var resolved = utils.ResolveType(vt)
	fmt.Println(vt.TYPE_PRIMITIVE())
	fmt.Println(vt.CARDINALITY())
	fmt.Println(resolved)

}

// func (l *bistatListener) EnterSpecialFunction(ctx *parser.SpecialFunctionContext) {
// 	fmt.Println("Entered special function")
// }

// func (l *bistatListener) EnterAssignment(ctx *parser.AssignmentContext) {
// 	fmt.Println("Entered assignment")
// 	fmt.Println(ctx.VAR_CONS())
// }

// func (l *bistatListener) EnterReturnStmt(ctx *parser.ReturnStmtContext) {
// 	fmt.Println("Entered return stmt")
// }

// func (l *bistatListener) EnterFunctionCall(ctx *parser.FunctionCallContext) {
// 	fmt.Println("Entered function call")
// }

// func (l *bistatListener) EnterPrint(ctx *parser.PrintContext) {
// 	fmt.Println("Entered print")
// }

// func (l *bistatListener) ExitProgram(ctx *parser.ProgramContext) {
// 	fmt.Println("Exited program")
// }

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewBistatLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewBistatParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	antlr.ParseTreeWalkerDefault.Walk(NewBistatListener(), p.Program())
}
