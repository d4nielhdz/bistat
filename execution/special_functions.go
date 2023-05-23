package main

import (
	"bistat/src"
	"fmt"
)

func (eCtx *ECtx) HandlePrint() {
	quad := eCtx.GetCurrentQuad()
	prevQuad := eCtx.Quads[eCtx.IP-1]
	printSpace := prevQuad.Op == src.Print
	if printSpace {
		fmt.Print(" ")
	}
	pType := GetPTypeFromAddress(quad.Op1)
	if pType == src.Int {
		val := eCtx.GetIntFromAddress(quad.Op1)
		fmt.Print(val)
	} else if pType == src.Float {
		val := eCtx.GetFloatFromAddress(quad.Op1)
		fmt.Print(val)
	} else if pType == src.Bool {
		val := eCtx.GetBoolFromAddress(quad.Op1)
		fmt.Print(val)
	} else if pType == src.String {
		val := eCtx.GetStringFromAddress(quad.Op1)
		fmt.Print(val)
	}
}

func (eCtx *ECtx) HandlePrintN() {
	fmt.Print("\n")
}
