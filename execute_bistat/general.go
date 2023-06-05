package main

import (
	src "bistat/src"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Handles assignment and read
func (eCtx *ECtx) HandleAssign() {
	quad := eCtx.GetCurrentQuad()
	rAddr := eCtx.GetDerefed(quad.Destination)
	lAddr := eCtx.GetDerefed(quad.Op1)

	pType := GetPTypeFromAddress(lAddr)
	// fmt.Println("assign", lAddr, rAddr)
	// fmt.Println("#", lAddr, src.PTypeToString(pType))
	if pType == src.Int {
		val := eCtx.GetIntFromAddress(rAddr)
		// fmt.Println(quad.Op1, lAddr, rAddr, val)
		eCtx.StoreIntAtAddress(val, lAddr)
	} else if pType == src.Float {
		val := eCtx.GetFloatFromAddress(rAddr)
		// fmt.Println(quad.Op1, lAddr, rAddr, val)
		eCtx.StoreFloatAtAddress(val, lAddr)
	} else if pType == src.String {
		val := eCtx.GetStringFromAddress(rAddr)
		eCtx.StoreStringAtAddress(val, lAddr)
	} else {
		val := eCtx.GetBoolFromAddress(rAddr)
		eCtx.StoreBoolAtAddress(val, lAddr)
	}
}

func (eCtx *ECtx) HandleInputRead() {
	quad := eCtx.GetCurrentQuad()
	addr := eCtx.GetDerefed(quad.Op1)
	pType := GetPTypeFromAddress(addr)
	reader := bufio.NewReader(os.Stdin)

	if pType == src.Int {
		var val int64
		fmt.Scan(&val)
		// fmt.Println(quad.Op1, lAddr, rAddr, val)
		eCtx.StoreIntAtAddress(val, addr)
	} else if pType == src.Float {
		var val float64
		fmt.Scan(&val)
		// fmt.Println(quad.Op1, lAddr, rAddr, val)
		eCtx.StoreFloatAtAddress(val, addr)
	} else if pType == src.String {
		val, _ := reader.ReadString('\n')
		eCtx.StoreStringAtAddress(strings.TrimSuffix(val, "\n"), addr)
	} else {
		var val bool
		fmt.Scan(&val)
		eCtx.StoreBoolAtAddress(val, addr)
	}
}

func (eCtx *ECtx) HandleEnd() {
}
