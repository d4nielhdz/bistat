package main

import (
	"strconv"
)

// Handles array op codes
func (eCtx *ECtx) HandleRefSum() {
	quad := eCtx.GetCurrentQuad()
	destination := quad.Destination
	addr1 := eCtx.GetDerefed(quad.Op1)
	addr2 := eCtx.GetDerefed(quad.Op2)

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

func (eCtx *ECtx) HandleRefMul() {
	quad := eCtx.GetCurrentQuad()
	destination := quad.Destination
	addr1 := eCtx.GetDerefed(quad.Op1)
	addr2 := eCtx.GetDerefed(quad.Op2)

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

func (eCtx *ECtx) HandleVerify() {
	quad := eCtx.GetCurrentQuad()
	addr1 := eCtx.GetDerefed(quad.Op1)
	val := eCtx.GetIntFromAddress(addr1)
	// fmt.Println("ver", quad.Op1, val, quad.Destination)
	if int(val) >= quad.Destination {
		eCtx.NewError("Access " + strconv.Itoa(int(val)) + " out of bounds for element with dimension " + strconv.Itoa(quad.Destination))
	}
}
