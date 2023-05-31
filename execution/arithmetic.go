package main

import (
	src "bistat/src"
)

func (eCtx *ECtx) HandleSum() {
	quad := eCtx.GetCurrentQuad()
	destination := quad.Destination
	addr1 := eCtx.GetDerefed(quad.Op1)
	addr2 := eCtx.GetDerefed(quad.Op2)

	pType1 := GetPTypeFromAddress(addr1)
	pType2 := GetPTypeFromAddress(addr2)

	if pType1 == src.Float || pType2 == src.Float {
		if pType1 == src.Float && pType2 == src.Float {
			val1 := eCtx.GetFloatFromAddress(addr1)
			val2 := eCtx.GetFloatFromAddress(addr2)
			result := val1 + val2
			// fmt.Println(val1, val2, result, destination)
			eCtx.StoreFloatAtAddress(result, destination)
		} else if pType1 == src.Float {
			val1 := eCtx.GetFloatFromAddress(addr1)
			var val2 int64 = 0
			if eCtx.AddrIsRef(quad.Op2) {
				// val2 = int64(addr2)
				val2 = eCtx.GetIntFromAddress(addr2)
			} else {
				val2 = eCtx.GetIntFromAddress(addr2)
			}
			result := val1 + float64(val2)
			eCtx.StoreFloatAtAddress(result, destination)
		} else {
			var val1 int64 = 0
			if eCtx.AddrIsRef(quad.Op1) {
				// val1 = int64(addr1)
				val1 = eCtx.GetIntFromAddress(addr1)
			} else {
				val1 = eCtx.GetIntFromAddress(addr1)
			}
			val2 := eCtx.GetFloatFromAddress(addr2)
			result := float64(val1) + val2
			eCtx.StoreFloatAtAddress(result, destination)
		}
	} else {
		// fmt.Println(quad.Op1, quad.Op2, addr1, addr2)
		var val1 int64
		if eCtx.AddrIsRef(quad.Op1) {
			// val1 = int64(addr1)
			val1 = eCtx.GetIntFromAddress(addr1)
		} else {
			val1 = eCtx.GetIntFromAddress(addr1)
		}
		var val2 int64
		if eCtx.AddrIsRef(quad.Op2) {
			// val2 = int64(addr2)
			val2 = eCtx.GetIntFromAddress(addr2)
		} else {
			val2 = eCtx.GetIntFromAddress(addr2)
		}
		result := val1 + val2
		// fmt.Println(eCtx.IP, "- sum", addr1, addr2, val1, val2, result, destination)
		if eCtx.AddrIsRef(destination) {
			eCtx.StoreRefAtAddress(int(result), destination)
		} else {
			eCtx.StoreIntAtAddress(result, destination)
		}
	}
}

func (eCtx *ECtx) HandleSubtraction() {
	quad := eCtx.GetCurrentQuad()
	destination := eCtx.GetDerefed(quad.Destination)
	addr1 := eCtx.GetDerefed(quad.Op1)
	addr2 := eCtx.GetDerefed(quad.Op2)

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
			var val2 int64 = 0
			if eCtx.AddrIsRef(quad.Op2) {
				// val2 = int64(addr2)
				val2 = eCtx.GetIntFromAddress(addr2)
			} else {
				val2 = eCtx.GetIntFromAddress(addr2)
			}
			result := val1 - float64(val2)
			eCtx.StoreFloatAtAddress(result, destination)
		} else {
			var val1 int64 = 0
			if eCtx.AddrIsRef(quad.Op1) {
				// val1 = int64(addr1)
				val1 = eCtx.GetIntFromAddress(addr1)
			} else {
				val1 = eCtx.GetIntFromAddress(addr1)
			}
			val2 := eCtx.GetFloatFromAddress(addr2)
			result := float64(val1) - val2
			eCtx.StoreFloatAtAddress(result, destination)
		}
	} else {
		var val1 int64
		if eCtx.AddrIsRef(quad.Op1) {
			// val1 = int64(addr1)
			val1 = eCtx.GetIntFromAddress(addr1)
		} else {
			val1 = eCtx.GetIntFromAddress(addr1)
		}
		var val2 int64
		if eCtx.AddrIsRef(quad.Op2) {
			// val2 = int64(addr2)
			val2 = eCtx.GetIntFromAddress(addr2)
		} else {
			val2 = eCtx.GetIntFromAddress(addr2)
		}
		result := val1 - val2
		eCtx.StoreIntAtAddress(result, destination)
	}
}

func (eCtx *ECtx) HandleMultiplication() {
	quad := eCtx.GetCurrentQuad()
	destination := quad.Destination
	addr1 := eCtx.GetDerefed(quad.Op1)
	addr2 := eCtx.GetDerefed(quad.Op2)

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
			var val2 int64 = 0
			if eCtx.AddrIsRef(quad.Op2) {
				// val2 = int64(addr2)
				val2 = eCtx.GetIntFromAddress(addr2)
			} else {
				val2 = eCtx.GetIntFromAddress(addr2)
			}
			result := val1 * float64(val2)
			eCtx.StoreFloatAtAddress(result, destination)
		} else {
			var val1 int64 = 0
			if eCtx.AddrIsRef(quad.Op1) {
				// val1 = int64(addr1)
				val1 = eCtx.GetIntFromAddress(addr1)
			} else {
				val1 = eCtx.GetIntFromAddress(addr1)
			}
			val2 := eCtx.GetFloatFromAddress(addr2)
			result := float64(val1) * val2
			eCtx.StoreFloatAtAddress(result, destination)
		}
	} else {
		var val1 int64
		if eCtx.AddrIsRef(quad.Op1) {
			// val1 = int64(addr1)
			val1 = eCtx.GetIntFromAddress(addr1)

		} else {
			val1 = eCtx.GetIntFromAddress(addr1)
		}
		var val2 int64
		if eCtx.AddrIsRef(quad.Op2) {
			// val2 = int64(addr2)
			val2 = eCtx.GetIntFromAddress(addr2)
		} else {
			val2 = eCtx.GetIntFromAddress(addr2)
		}
		result := val1 * val2
		// fmt.Println(eCtx.IP, "- mul", addr1, addr2, val1, val2, result, destination)
		if eCtx.AddrIsRef(destination) {
			eCtx.StoreRefAtAddress(int(result), destination)
		} else {
			eCtx.StoreIntAtAddress(result, destination)
		}
	}
}

func (eCtx *ECtx) HandleDivision() {
	quad := eCtx.GetCurrentQuad()
	destination := eCtx.GetDerefed(quad.Destination)
	addr1 := eCtx.GetDerefed(quad.Op1)
	addr2 := eCtx.GetDerefed(quad.Op2)

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
			var val2 int64 = 0
			if eCtx.AddrIsRef(quad.Op2) {
				// val2 = int64(addr2)
				val2 = eCtx.GetIntFromAddress(addr2)
			} else {
				val2 = eCtx.GetIntFromAddress(addr2)
			}
			result := val1 / float64(val2)
			eCtx.StoreFloatAtAddress(result, destination)
		} else {
			var val1 int64 = 0
			if eCtx.AddrIsRef(quad.Op1) {
				// val1 = int64(addr1)
				val1 = eCtx.GetIntFromAddress(addr1)
			} else {
				val1 = eCtx.GetIntFromAddress(addr1)
			}
			val2 := eCtx.GetFloatFromAddress(addr2)
			result := float64(val1) / val2
			eCtx.StoreFloatAtAddress(result, destination)
		}
	} else {
		var val1 int64
		if eCtx.AddrIsRef(quad.Op1) {
			// val1 = int64(addr1)
			val1 = eCtx.GetIntFromAddress(addr1)
		} else {
			val1 = eCtx.GetIntFromAddress(addr1)
		}
		var val2 int64
		if eCtx.AddrIsRef(quad.Op2) {
			// val2 = int64(addr2)
			val2 = eCtx.GetIntFromAddress(addr2)
		} else {
			val2 = eCtx.GetIntFromAddress(addr2)
		}
		result := val1 / val2
		eCtx.StoreIntAtAddress(result, destination)
	}
}

func (eCtx *ECtx) HandleModulus() {
	quad := eCtx.GetCurrentQuad()
	destination := eCtx.GetDerefed(quad.Destination)
	addr1 := eCtx.GetDerefed(quad.Op1)
	addr2 := eCtx.GetDerefed(quad.Op2)

	val1 := eCtx.GetIntFromAddress(addr1)
	val2 := eCtx.GetIntFromAddress(addr2)
	result := val1 % val2
	eCtx.StoreIntAtAddress(result, destination)
}

func (eCtx *ECtx) HandleRefSum() {
	quad := eCtx.GetCurrentQuad()
	destination := quad.Destination
	addr1 := eCtx.GetDerefed(quad.Op1)
	addr2 := eCtx.GetDerefed(quad.Op2)

	pType1 := GetPTypeFromAddress(addr1)
	pType2 := GetPTypeFromAddress(addr2)

	if pType1 == src.Float || pType2 == src.Float {
		if pType1 == src.Float && pType2 == src.Float {
			val1 := eCtx.GetFloatFromAddress(addr1)
			val2 := eCtx.GetFloatFromAddress(addr2)
			result := val1 + val2
			// fmt.Println(val1, val2, result, destination)
			eCtx.StoreFloatAtAddress(result, destination)
		} else if pType1 == src.Float {
			val1 := eCtx.GetFloatFromAddress(addr1)
			var val2 int64 = 0
			if eCtx.AddrIsRef(quad.Op2) {
				val2 = int64(addr2)
			} else {
				val2 = eCtx.GetIntFromAddress(addr2)
			}
			result := val1 + float64(val2)
			eCtx.StoreFloatAtAddress(result, destination)
		} else {
			var val1 int64 = 0
			if eCtx.AddrIsRef(quad.Op1) {
				val1 = int64(addr1)
			} else {
				val1 = eCtx.GetIntFromAddress(addr1)
			}
			val2 := eCtx.GetFloatFromAddress(addr2)
			result := float64(val1) + val2
			eCtx.StoreFloatAtAddress(result, destination)
		}
	} else {
		// fmt.Println(quad.Op1, quad.Op2, addr1, addr2)
		var val1 int64
		if eCtx.AddrIsRef(quad.Op1) {
			val1 = int64(addr1)
		} else {
			val1 = eCtx.GetIntFromAddress(addr1)
		}
		var val2 int64
		if eCtx.AddrIsRef(quad.Op2) {
			val2 = int64(addr2)
		} else {
			val2 = eCtx.GetIntFromAddress(addr2)
		}
		result := val1 + val2
		// fmt.Println(eCtx.IP, "- sum", addr1, addr2, val1, val2, result, destination)
		if eCtx.AddrIsRef(destination) {
			eCtx.StoreRefAtAddress(int(result), destination)
		} else {
			eCtx.StoreIntAtAddress(result, destination)
		}
	}
}

func (eCtx *ECtx) HandleRefMul() {
	quad := eCtx.GetCurrentQuad()
	destination := quad.Destination
	addr1 := eCtx.GetDerefed(quad.Op1)
	addr2 := eCtx.GetDerefed(quad.Op2)

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
			var val2 int64 = 0
			if eCtx.AddrIsRef(quad.Op2) {
				val2 = int64(addr2)
			} else {
				val2 = eCtx.GetIntFromAddress(addr2)
			}
			result := val1 * float64(val2)
			eCtx.StoreFloatAtAddress(result, destination)
		} else {
			var val1 int64 = 0
			if eCtx.AddrIsRef(quad.Op1) {
				val1 = int64(addr1)
			} else {
				val1 = eCtx.GetIntFromAddress(addr1)
			}
			val2 := eCtx.GetFloatFromAddress(addr2)
			result := float64(val1) * val2
			eCtx.StoreFloatAtAddress(result, destination)
		}
	} else {
		var val1 int64
		if eCtx.AddrIsRef(quad.Op1) {
			val1 = int64(addr1)
			// val1 = eCtx.GetIntFromAddress(addr1)
		} else {
			val1 = eCtx.GetIntFromAddress(addr1)
		}
		var val2 int64
		if eCtx.AddrIsRef(quad.Op2) {
			val2 = int64(addr2)
			// val2 = eCtx.GetIntFromAddress(addr2)

		} else {
			val2 = eCtx.GetIntFromAddress(addr2)
		}
		result := val1 * val2
		// fmt.Println(eCtx.IP, "- mul", addr1, addr2, val1, val2, result, destination)
		if eCtx.AddrIsRef(destination) {
			eCtx.StoreRefAtAddress(int(result), destination)
		} else {
			eCtx.StoreIntAtAddress(result, destination)
		}
	}
}
