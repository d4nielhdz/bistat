// Code generated from Bistat.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Bistat

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BistatListener is a complete listener for a parse tree produced by BistatParser.
type BistatListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)
}
