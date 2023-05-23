package main

import (
	src "bistat/src"
	"fmt"
)

func (eCtx *ECtx) LoadConstMemory() {
	for val, rType := range eCtx.ConsTable {
		eCtx.ConstMemory.StoreVal(val, rType.Address)
	}
}

func GetPTypeFromAddress(address int) src.PType {
	if (address >= src.GLOBAL_INT_START && address < src.GLOBAL_FLOAT_START) || (address >= src.LOCAL_INT_START && address < src.LOCAL_FLOAT_START) || (address >= src.CONST_INT_START && address < src.CONST_FLOAT_START) || (address >= src.TEMP_INT_START && address < src.TEMP_FLOAT_START) {
		return src.Int
	} else if (address >= src.GLOBAL_FLOAT_START && address < src.GLOBAL_STRING_START) || (address >= src.LOCAL_FLOAT_START && address < src.LOCAL_STRING_START) || (address >= src.CONST_FLOAT_START && address < src.CONST_STRING_START) || (address >= src.TEMP_FLOAT_START && address < src.TEMP_STRING_START) {
		return src.Float
	} else if (address >= src.GLOBAL_STRING_START && address < src.GLOBAL_BOOL_START) || (address >= src.LOCAL_STRING_START && address < src.LOCAL_BOOL_START) || (address >= src.CONST_STRING_START && address < src.CONST_BOOL_START) || (address >= src.TEMP_STRING_START && address < src.TEMP_BOOL_START) {
		return src.String
	}
	return src.Bool
}

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
