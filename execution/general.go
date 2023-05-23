package main

import (
	src "bistat/src"
	"fmt"
)

func (eCtx *ECtx) HandleAssign() {
	quad := eCtx.GetCurrentQuad()
	lAddr := quad.Destination
	rAddr := quad.Op1
	pType := GetPTypeFromAddress(lAddr)
	fmt.Println(lAddr, rAddr, src.PTypeToString(pType))
	if pType == src.Int {
		val := eCtx.GetIntFromAddress(rAddr)
		eCtx.StoreIntAtAddress(val, lAddr)
	} else if pType == src.Float {
		val := eCtx.GetFloatFromAddress(rAddr)
		eCtx.StoreFloatAtAddress(val, lAddr)
	} else if pType == src.String {
		val := eCtx.GetStringFromAddress(rAddr)
		eCtx.StoreStringAtAddress(val, lAddr)
	} else {
		val := eCtx.GetBoolFromAddress(rAddr)
		eCtx.StoreBoolAtAddress(val, lAddr)
	}
}
