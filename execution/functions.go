package main

import "bistat/src"

func (eCtx *ECtx) HandleEra() {
	quad := eCtx.GetCurrentQuad()
	funcData := eCtx.FuncDir[eCtx.Functions[quad.Op1]]
	stackSegment := NewStackSegment(funcData.LocalFloatVars, funcData.LocalIntVars, funcData.LocalStringVars, funcData.LocalBoolVars, funcData.LocalRefs, src.LOCAL_FLOAT_START, src.LOCAL_INT_START, src.LOCAL_STRING_START, src.LOCAL_BOOL_START, src.LOCAL_REF_START, funcData.TempFloatVars, funcData.TempIntVars, funcData.TempStringVars, funcData.TempBoolVars, src.TEMP_FLOAT_START, src.TEMP_INT_START, src.TEMP_STRING_START, src.TEMP_BOOL_START)
	eCtx.PushStackSegment(stackSegment)
}

func (eCtx *ECtx) HandleReturn() {
	eCtx.HandleAssign()
	eCtx.IP = eCtx.Jumps[len(eCtx.Jumps)-1]
	popped := eCtx.Jumps[:len(eCtx.Jumps)-1]
	eCtx.Jumps = popped
	newStack := eCtx.StackSegment[:len(eCtx.StackSegment)-1]
	eCtx.StackSegment = newStack
}

func (eCtx *ECtx) HandleEndFunc() {
	eCtx.IP = eCtx.Jumps[len(eCtx.Jumps)-1]
	popped := eCtx.Jumps[:len(eCtx.Jumps)-1]
	eCtx.Jumps = popped
	newStack := eCtx.StackSegment[:len(eCtx.StackSegment)-1]
	eCtx.StackSegment = newStack
}

func (eCtx *ECtx) HandleGoSub() {
	quad := eCtx.GetCurrentQuad()
	funcIdx := quad.Op1
	funcData := eCtx.FuncDir[eCtx.Functions[funcIdx]]
	eCtx.Jumps = append(eCtx.Jumps, eCtx.IP+1)
	eCtx.IP = funcData.FuncStart
}

func (eCtx *ECtx) HandleParam() {
	quad := eCtx.GetCurrentQuad()
	rAddr := eCtx.GetDerefed(quad.Destination)
	lAddr := eCtx.GetDerefed(quad.Op1)

	pType := GetPTypeFromAddress(lAddr)
	// fmt.Println("#", lAddr, src.PTypeToString(pType))
	if pType == src.Int {
		val := eCtx.GetIntFromAddressForParam(rAddr)
		// fmt.Println(quad.Op1, lAddr, rAddr, val)
		eCtx.StoreIntAtAddress(val, lAddr)
	} else if pType == src.Float {
		val := eCtx.GetFloatFromAddressForParam(rAddr)
		// fmt.Println(quad.Op1, lAddr, rAddr, val)
		eCtx.StoreFloatAtAddress(val, lAddr)
	} else if pType == src.String {
		val := eCtx.GetStringFromAddressForParam(rAddr)
		eCtx.StoreStringAtAddress(val, lAddr)
	} else {
		val := eCtx.GetBoolFromAddressForParam(rAddr)
		eCtx.StoreBoolAtAddress(val, lAddr)
	}
}
