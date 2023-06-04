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
	printQuads := false
	if len(os.Args) > 2 && os.Args[2] == "1" {
		printQuads = true
	}
	antlr.ParseTreeWalkerDefault.Walk(src.NewBistatListener(os.Args[1], printQuads), p.Program())
}
