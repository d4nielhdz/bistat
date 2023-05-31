package main

import (
	src "bistat/src"
)

func (eCtx *ECtx) HandleAssign() {
	quad := eCtx.GetCurrentQuad()
	rAddr := eCtx.GetDerefed(quad.Destination)
	lAddr := eCtx.GetDerefed(quad.Op1)

	pType := GetPTypeFromAddress(lAddr)
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
