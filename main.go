package main

import (
	parser "bistat/parser"
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type bistatListener struct {
	*parser.BaseBistatListener
}

func NewBistatListener() *bistatListener {
	return new(bistatListener)
}

func (l *bistatListener) EnterProgram(ctx *parser.ProgramContext) {
	fmt.Println("Entered program")
	fmt.Println(ctx.ID())
}

func (l *bistatListener) EnterVar_declaration(ctx *parser.Var_declarationContext) {
	fmt.Println("Entered variable declaration")
	fmt.Println(ctx.ID())
	fmt.Println(ctx.NON_VOID_TYPE())
}

func (l *bistatListener) EnterStmt(ctx *parser.StmtContext) {
	fmt.Println("Entered statement")
	fmt.Println(ctx.GetBaseRuleContext())
}

func (l *bistatListener) ExitProgram(ctx *parser.ProgramContext) {
	fmt.Println("Exited program")
	fmt.Println(ctx.GetBaseRuleContext())
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewBistatLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewBistatParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	antlr.ParseTreeWalkerDefault.Walk(NewBistatListener(), p.Program())
}
