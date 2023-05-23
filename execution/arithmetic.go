package main

import (
	src "bistat/src"
)

func (eCtx *ECtx) HandleSum() {
	quad := eCtx.GetCurrentQuad()
	destination := quad.Destination
	addr1 := quad.Op1
	addr2 := quad.Op2

	pType1 := GetPTypeFromAddress(addr1)
	pType2 := GetPTypeFromAddress(addr2)

	if pType1 == src.Float || pType2 == src.Float {
		if pType1 == src.Float && pType2 == src.Float {
			val1 := eCtx.GetFloatFromAddress(addr1)
			val2 := eCtx.GetFloatFromAddress(addr2)
			result := val1 + val2
			eCtx.StoreFloatAtAddress(result, destination)
		} else if pType1 == src.Float {
			val1 := eCtx.GetFloatFromAddress(addr1)
			val2 := eCtx.GetIntFromAddress(addr2)
			result := val1 + float64(val2)
			eCtx.StoreFloatAtAddress(result, destination)
		} else {
			val1 := eCtx.GetIntFromAddress(addr1)
			val2 := eCtx.GetFloatFromAddress(addr2)
			result := float64(val1) + val2
			eCtx.StoreFloatAtAddress(result, destination)
		}
	} else {
		val1 := eCtx.GetIntFromAddress(addr1)
		val2 := eCtx.GetIntFromAddress(addr2)
		result := val1 + val2
		eCtx.StoreIntAtAddress(result, destination)
	}
}

func (eCtx *ECtx) HandleSubtraction() {
	quad := eCtx.GetCurrentQuad()
	destination := quad.Destination
	addr1 := quad.Op1
	addr2 := quad.Op2

	pType1 := GetPTypeFromAddress(addr1)
	pType2 := GetPTypeFromAddress(addr2)

	if pType1 == src.Float || pType2 == src.Float {
		if pType1 == src.Float && pType2 == src.Float {
			val1 := eCtx.GetFloatFromAddress(addr1)
			val2 := eCtx.GetFloatFromAddress(addr2)
			result := val1 - val2
			eCtx.StoreFloatAtAddress(result, destination)
		} else if pType1 == src.Float {
			val1 := eCtx.GetFloatFromAddress(addr1)
			val2 := eCtx.GetIntFromAddress(addr2)
			result := val1 - float64(val2)
			eCtx.StoreFloatAtAddress(result, destination)
		} else {
			val1 := eCtx.GetIntFromAddress(addr1)
			val2 := eCtx.GetFloatFromAddress(addr2)
			result := float64(val1) - val2
			eCtx.StoreFloatAtAddress(result, destination)
		}
	} else {
		val1 := eCtx.GetIntFromAddress(addr1)
		val2 := eCtx.GetIntFromAddress(addr2)
		result := val1 - val2
		eCtx.StoreIntAtAddress(result, destination)
	}
}

func (eCtx *ECtx) HandleMultiplication() {
	quad := eCtx.GetCurrentQuad()
	destination := quad.Destination
	addr1 := quad.Op1
	addr2 := quad.Op2

	pType1 := GetPTypeFromAddress(addr1)
	pType2 := GetPTypeFromAddress(addr2)

	if pType1 == src.Float || pType2 == src.Float {
		if pType1 == src.Float && pType2 == src.Float {
			val1 := eCtx.GetFloatFromAddress(addr1)
			val2 := eCtx.GetFloatFromAddress(addr2)
			result := val1 * val2
			eCtx.StoreFloatAtAddress(result, destination)
		} else if pType1 == src.Float {
			val1 := eCtx.GetFloatFromAddress(addr1)
			val2 := eCtx.GetIntFromAddress(addr2)
			result := val1 * float64(val2)
			eCtx.StoreFloatAtAddress(result, destination)
		} else {
			val1 := eCtx.GetIntFromAddress(addr1)
			val2 := eCtx.GetFloatFromAddress(addr2)
			result := float64(val1) * val2
			eCtx.StoreFloatAtAddress(result, destination)
		}
	} else {
		val1 := eCtx.GetIntFromAddress(addr1)
		val2 := eCtx.GetIntFromAddress(addr2)
		result := val1 * val2
		eCtx.StoreIntAtAddress(result, destination)
	}
}

func (eCtx *ECtx) HandleDivision() {
	quad := eCtx.GetCurrentQuad()
	destination := quad.Destination
	addr1 := quad.Op1
	addr2 := quad.Op2

	pType1 := GetPTypeFromAddress(addr1)
	pType2 := GetPTypeFromAddress(addr2)

	if pType1 == src.Float || pType2 == src.Float {
		if pType1 == src.Float && pType2 == src.Float {
			val1 := eCtx.GetFloatFromAddress(addr1)
			val2 := eCtx.GetFloatFromAddress(addr2)
			result := val1 / val2
			eCtx.StoreFloatAtAddress(result, destination)
		} else if pType1 == src.Float {
			val1 := eCtx.GetFloatFromAddress(addr1)
			val2 := eCtx.GetIntFromAddress(addr2)
			result := val1 / float64(val2)
			eCtx.StoreFloatAtAddress(result, destination)
		} else {
			val1 := eCtx.GetIntFromAddress(addr1)
			val2 := eCtx.GetFloatFromAddress(addr2)
			result := float64(val1) / val2
			eCtx.StoreFloatAtAddress(result, destination)
		}
	} else {
		val1 := eCtx.GetIntFromAddress(addr1)
		val2 := eCtx.GetIntFromAddress(addr2)
		result := val1 / val2
		eCtx.StoreIntAtAddress(result, destination)
	}
}

func (eCtx *ECtx) HandleModulus() {
	quad := eCtx.GetCurrentQuad()
	destination := quad.Destination
	addr1 := quad.Op1
	addr2 := quad.Op2

	val1 := eCtx.GetIntFromAddress(addr1)
	val2 := eCtx.GetIntFromAddress(addr2)
	result := val1 % val2
	eCtx.StoreIntAtAddress(result, destination)
}
