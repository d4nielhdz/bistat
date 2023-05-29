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
	addr := eCtx.GetDerefed(quad.Op1)
	// fmt.Println("addr", addr)
	pType := GetPTypeFromAddress(addr)
	if pType == src.Int {
		val := eCtx.GetIntFromAddress(addr)
		fmt.Print(val)
	} else if pType == src.Float {
		val := eCtx.GetFloatFromAddress(addr)
		fmt.Print(val)
	} else if pType == src.Bool {
		val := eCtx.GetBoolFromAddress(addr)
		fmt.Print(val)
	} else if pType == src.String {
		val := eCtx.GetStringFromAddress(addr)
		fmt.Print(val)
	}
}

func (eCtx *ECtx) HandlePrintN() {
	fmt.Print("\n")
}
