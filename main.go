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

	// // Start the parsing process
	// program := p.Program()

	// // Get the parse tree
	// parseTree := program.GetRuleContext()

	// // Print the parse tree
	// fmt.Println(parseTree.ToStringTree([]string{}, p))
	antlr.ParseTreeWalkerDefault.Walk(src.NewBistatListener(), p.Program())
}
