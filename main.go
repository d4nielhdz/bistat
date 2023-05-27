package main

import (
	parser "bistat/parser"
	src "bistat/src"
	"os"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewBistatLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewBistatParser(stream)
	antlr.ParseTreeWalkerDefault.Walk(src.NewBistatListener(), p.Program())
}
