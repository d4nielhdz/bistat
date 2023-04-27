package main

import (
	parser "bistat/parser"
	utils "bistat/utils"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewBistatLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewBistatParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	antlr.ParseTreeWalkerDefault.Walk(utils.NewBistatListener(), p.Program())
}
