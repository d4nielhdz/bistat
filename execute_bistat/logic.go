package main

import (
	src "bistat/src"
)

// Handles logic op codes

func (eCtx *ECtx) HandleAnd() {
	quad := eCtx.GetCurrentQuad()
	addr1 := eCtx.GetDerefed(quad.Op1)
	addr2 := eCtx.GetDerefed(quad.Op2)

	val1 := eCtx.GetBoolFromAddress(addr1)
	val2 := eCtx.GetBoolFromAddress(addr2)
	result := val1 && val2
	eCtx.StoreBoolAtAddress(result, eCtx.GetDerefed(quad.Destination))
}

func (eCtx *ECtx) HandleOr() {
	quad := eCtx.GetCurrentQuad()
	addr1 := eCtx.GetDerefed(quad.Op1)
	addr2 := eCtx.GetDerefed(quad.Op2)

	val1 := eCtx.GetBoolFromAddress(addr1)
	val2 := eCtx.GetBoolFromAddress(addr2)
	result := val1 || val2
	eCtx.StoreBoolAtAddress(result, quad.Destination)
}

func (eCtx *ECtx) HandleGt() {
	quad := eCtx.GetCurrentQuad()
	addr1 := eCtx.GetDerefed(quad.Op1)
	addr2 := eCtx.GetDerefed(quad.Op2)
	pType1 := GetPTypeFromAddress(addr1)
	pType2 := GetPTypeFromAddress(addr2)

	if pType1 == src.Float || pType2 == src.Float {
		if pType1 == src.Float && pType2 == src.Float {
			val1 := eCtx.GetFloatFromAddress(addr1)
			val2 := eCtx.GetFloatFromAddress(addr2)
			result := val1 > val2
			eCtx.StoreBoolAtAddress(result, quad.Destination)
		} else if pType1 == src.Float {
			val1 := eCtx.GetFloatFromAddress(addr1)
			val2 := eCtx.GetIntFromAddress(addr2)
			result := val1 > float64(val2)
			eCtx.StoreBoolAtAddress(result, quad.Destination)
		} else {
			val1 := eCtx.GetIntFromAddress(addr1)
			val2 := eCtx.GetFloatFromAddress(addr2)
			result := float64(val1) > val2
			eCtx.StoreBoolAtAddress(result, quad.Destination)
		}
	} else {
		val1 := eCtx.GetIntFromAddress(addr1)
		val2 := eCtx.GetIntFromAddress(addr2)
		result := val1 > val2
		eCtx.StoreBoolAtAddress(result, quad.Destination)
	}
}

func (eCtx *ECtx) HandleLt() {
	quad := eCtx.GetCurrentQuad()
	addr1 := eCtx.GetDerefed(quad.Op1)
	addr2 := eCtx.GetDerefed(quad.Op2)
	pType1 := GetPTypeFromAddress(addr1)
	pType2 := GetPTypeFromAddress(addr2)

	if pType1 == src.Float || pType2 == src.Float {
		if pType1 == src.Float && pType2 == src.Float {
			val1 := eCtx.GetFloatFromAddress(addr1)
			val2 := eCtx.GetFloatFromAddress(addr2)
			result := val1 < val2
			eCtx.StoreBoolAtAddress(result, quad.Destination)
		} else if pType1 == src.Float {
			val1 := eCtx.GetFloatFromAddress(addr1)
			val2 := eCtx.GetIntFromAddress(addr2)
			result := val1 < float64(val2)
			eCtx.StoreBoolAtAddress(result, quad.Destination)
		} else {
			val1 := eCtx.GetIntFromAddress(addr1)
			val2 := eCtx.GetFloatFromAddress(addr2)
			result := float64(val1) < val2
			eCtx.StoreBoolAtAddress(result, quad.Destination)
		}
	} else {
		val1 := eCtx.GetIntFromAddress(addr1)
		val2 := eCtx.GetIntFromAddress(addr2)
		result := val1 < val2
		eCtx.StoreBoolAtAddress(result, quad.Destination)
	}
}

func (eCtx *ECtx) HandleGe() {
	quad := eCtx.GetCurrentQuad()
	addr1 := eCtx.GetDerefed(quad.Op1)
	addr2 := eCtx.GetDerefed(quad.Op2)
	pType1 := GetPTypeFromAddress(addr1)
	pType2 := GetPTypeFromAddress(addr2)

	if pType1 == src.Float || pType2 == src.Float {
		if pType1 == src.Float && pType2 == src.Float {
			val1 := eCtx.GetFloatFromAddress(addr1)
			val2 := eCtx.GetFloatFromAddress(addr2)
			result := val1 >= val2
			eCtx.StoreBoolAtAddress(result, quad.Destination)
		} else if pType1 == src.Float {
			val1 := eCtx.GetFloatFromAddress(addr1)
			val2 := eCtx.GetIntFromAddress(addr2)
			result := val1 >= float64(val2)
			eCtx.StoreBoolAtAddress(result, quad.Destination)
		} else {
			val1 := eCtx.GetIntFromAddress(addr1)
			val2 := eCtx.GetFloatFromAddress(addr2)
			result := float64(val1) >= val2
			eCtx.StoreBoolAtAddress(result, quad.Destination)
		}
	} else {
		val1 := eCtx.GetIntFromAddress(addr1)
		val2 := eCtx.GetIntFromAddress(addr2)
		result := val1 >= val2
		eCtx.StoreBoolAtAddress(result, quad.Destination)
	}
}

func (eCtx *ECtx) HandleLe() {
	quad := eCtx.GetCurrentQuad()
	addr1 := eCtx.GetDerefed(quad.Op1)
	addr2 := eCtx.GetDerefed(quad.Op2)
	pType1 := GetPTypeFromAddress(addr1)
	pType2 := GetPTypeFromAddress(addr2)

	if pType1 == src.Float || pType2 == src.Float {
		if pType1 == src.Float && pType2 == src.Float {
			val1 := eCtx.GetFloatFromAddress(addr1)
			val2 := eCtx.GetFloatFromAddress(addr2)
			result := val1 <= val2
			eCtx.StoreBoolAtAddress(result, quad.Destination)
		} else if pType1 == src.Float {
			val1 := eCtx.GetFloatFromAddress(addr1)
			val2 := eCtx.GetIntFromAddress(addr2)
			result := val1 <= float64(val2)
			eCtx.StoreBoolAtAddress(result, quad.Destination)
		} else {
			val1 := eCtx.GetIntFromAddress(addr1)
			val2 := eCtx.GetFloatFromAddress(addr2)
			result := float64(val1) <= val2
			eCtx.StoreBoolAtAddress(result, quad.Destination)
		}
	} else {
		val1 := eCtx.GetIntFromAddress(addr1)
		val2 := eCtx.GetIntFromAddress(addr2)
		result := val1 <= val2
		eCtx.StoreBoolAtAddress(result, quad.Destination)
	}
}

func (eCtx *ECtx) HandleEq() {
	quad := eCtx.GetCurrentQuad()
	addr1 := eCtx.GetDerefed(quad.Op1)
	addr2 := eCtx.GetDerefed(quad.Op2)
	pType1 := GetPTypeFromAddress(addr1)
	pType2 := GetPTypeFromAddress(addr2)

	if pType1 == src.Float || pType2 == src.Float {
		if pType1 == src.Float && pType2 == src.Float {
			val1 := eCtx.GetFloatFromAddress(addr1)
			val2 := eCtx.GetFloatFromAddress(addr2)
			result := val1 == val2
			eCtx.StoreBoolAtAddress(result, quad.Destination)
		} else if pType1 == src.Float {
			val1 := eCtx.GetFloatFromAddress(addr1)
			val2 := eCtx.GetIntFromAddress(addr2)
			result := val1 == float64(val2)
			eCtx.StoreBoolAtAddress(result, quad.Destination)
		} else {
			val1 := eCtx.GetIntFromAddress(addr1)
			val2 := eCtx.GetFloatFromAddress(addr2)
			result := float64(val1) == val2
			eCtx.StoreBoolAtAddress(result, quad.Destination)
		}
	} else if pType1 == src.Int && pType2 == src.Int {
		val1 := eCtx.GetIntFromAddress(addr1)
		val2 := eCtx.GetIntFromAddress(addr2)
		result := val1 == val2
		eCtx.StoreBoolAtAddress(result, quad.Destination)
	} else {
		val1 := eCtx.GetBoolFromAddress(addr1)
		val2 := eCtx.GetBoolFromAddress(addr2)
		result := val1 == val2
		eCtx.StoreBoolAtAddress(result, quad.Destination)
	}
}

func (eCtx *ECtx) HandleNe() {
	quad := eCtx.GetCurrentQuad()
	addr1 := eCtx.GetDerefed(quad.Op1)
	addr2 := eCtx.GetDerefed(quad.Op2)
	pType1 := GetPTypeFromAddress(addr1)
	pType2 := GetPTypeFromAddress(addr2)

	if pType1 == src.Float || pType2 == src.Float {
		if pType1 == src.Float && pType2 == src.Float {
			val1 := eCtx.GetFloatFromAddress(addr1)
			val2 := eCtx.GetFloatFromAddress(addr2)
			result := val1 != val2
			eCtx.StoreBoolAtAddress(result, quad.Destination)
		} else if pType1 == src.Float {
			val1 := eCtx.GetFloatFromAddress(addr1)
			val2 := eCtx.GetIntFromAddress(addr2)
			result := val1 != float64(val2)
			eCtx.StoreBoolAtAddress(result, quad.Destination)
		} else {
			val1 := eCtx.GetIntFromAddress(addr1)
			val2 := eCtx.GetFloatFromAddress(addr2)
			result := float64(val1) != val2
			eCtx.StoreBoolAtAddress(result, quad.Destination)
		}
	} else if pType1 == src.Int && pType2 == src.Int {
		val1 := eCtx.GetIntFromAddress(addr1)
		val2 := eCtx.GetIntFromAddress(addr2)
		result := val1 != val2
		eCtx.StoreBoolAtAddress(result, quad.Destination)
	} else {
		val1 := eCtx.GetBoolFromAddress(addr1)
		val2 := eCtx.GetBoolFromAddress(addr2)
		result := val1 != val2
		eCtx.StoreBoolAtAddress(result, quad.Destination)
	}
}
