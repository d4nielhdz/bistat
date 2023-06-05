package main

// Handles jumps
func (eCtx *ECtx) HandleGotoF() {
	quad := eCtx.GetCurrentQuad()
	val := eCtx.GetBoolFromAddress(eCtx.GetDerefed(quad.Op1))
	if !val {
		eCtx.IP = quad.Destination
	} else {
		eCtx.IP++
	}
}

func (eCtx *ECtx) HandleGotoT() {
	quad := eCtx.GetCurrentQuad()
	val := eCtx.GetBoolFromAddress(eCtx.GetDerefed(quad.Op1))
	if val {
		eCtx.IP = quad.Destination
	} else {
		eCtx.IP++
	}
}
