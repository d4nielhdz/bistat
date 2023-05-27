package main

func (eCtx *ECtx) HandleGotoF() {
	quad := eCtx.GetCurrentQuad()
	val := eCtx.GetBoolFromAddress(quad.Op1)
	if !val {
		eCtx.IP = quad.Destination
	} else {
		eCtx.IP++
	}
}

func (eCtx *ECtx) HandleGotoT() {
	quad := eCtx.GetCurrentQuad()
	val := eCtx.GetBoolFromAddress(quad.Op1)
	if val {
		eCtx.IP = quad.Destination
	} else {
		eCtx.IP++
	}
}
