package main

import (
	"bistat/src"
	"fmt"
	"math"
)

func (eCtx *ECtx) HandlePrint() {
	quad := eCtx.GetCurrentQuad()
	addr := eCtx.GetDerefed(quad.Op1)
	// fmt.Println(eCtx.IP, "#")
	// fmt.Println("t", quad.Op1)
	// fmt.Println("addr", addr)
	pType := GetPTypeFromAddress(addr)

	// fmt.Println(src.PTypeToString(pType))
	if pType == src.Int {
		// fmt.Println("print")
		val := eCtx.GetIntFromAddress(addr)
		fmt.Print(val)
	} else if pType == src.Float {
		val := eCtx.GetFloatFromAddress(addr)
		fmt.Print(val)
	} else if pType == src.Bool {
		val := eCtx.GetBoolFromAddress(addr)
		fmt.Print(val)
	} else if pType == src.String {
		val := eCtx.GetStringFromAddress(addr)
		fmt.Print(val)
	}
}

func (eCtx *ECtx) HandlePrintN() {
	fmt.Print("\n")
}

func (eCtx *ECtx) HandleCos() {
	quad := eCtx.GetCurrentQuad()
	addr1 := eCtx.GetDerefed(quad.Op1)
	addr2 := eCtx.GetDerefed(quad.Destination)

	pType := GetPTypeFromAddress(addr1)
	var cos float64

	if pType == src.Int {
		val := eCtx.GetIntFromAddress(addr1)
		cos = math.Cos(float64(val))
	} else if pType == src.Float {
		val := eCtx.GetFloatFromAddress(addr1)
		cos = math.Cos(val)
	}

	eCtx.StoreFloatAtAddress(cos, addr2)
}

func (eCtx *ECtx) HandleSin() {
	quad := eCtx.GetCurrentQuad()
	addr1 := eCtx.GetDerefed(quad.Op1)
	addr2 := eCtx.GetDerefed(quad.Destination)

	pType := GetPTypeFromAddress(addr1)
	var sin float64

	if pType == src.Int {
		val := eCtx.GetIntFromAddress(addr1)
		sin = math.Sin(float64(val))
	} else if pType == src.Float {
		val := eCtx.GetFloatFromAddress(addr1)
		sin = math.Sin(val)
	}

	eCtx.StoreFloatAtAddress(sin, addr2)
}

func (eCtx *ECtx) HandleTan() {
	quad := eCtx.GetCurrentQuad()
	addr1 := eCtx.GetDerefed(quad.Op1)
	addr2 := eCtx.GetDerefed(quad.Destination)

	pType := GetPTypeFromAddress(addr1)
	var tan float64

	if pType == src.Int {
		val := eCtx.GetIntFromAddress(addr1)
		tan = math.Tan(float64(val))
	} else if pType == src.Float {
		val := eCtx.GetFloatFromAddress(addr1)
		tan = math.Tan(val)
	}

	eCtx.StoreFloatAtAddress(tan, addr2)
}

func (eCtx *ECtx) HandleSqrt() {
	quad := eCtx.GetCurrentQuad()
	addr1 := eCtx.GetDerefed(quad.Op1)
	addr2 := eCtx.GetDerefed(quad.Destination)

	pType := GetPTypeFromAddress(addr1)

	if pType == src.Int {
		val := eCtx.GetIntFromAddress(addr1)
		sqrted := math.Sqrt(float64(val))
		eCtx.StoreIntAtAddress(int64(sqrted), addr2)
	} else if pType == src.Float {
		val := eCtx.GetFloatFromAddress(addr1)
		sqrted := math.Sqrt(float64(val))
		eCtx.StoreFloatAtAddress(sqrted, addr2)
	}
}

func (eCtx *ECtx) HandleFloor() {
	quad := eCtx.GetCurrentQuad()
	addr1 := eCtx.GetDerefed(quad.Op1)
	addr2 := eCtx.GetDerefed(quad.Destination)

	var floored float64

	val := eCtx.GetFloatFromAddress(addr1)
	floored = math.Floor(val)

	eCtx.StoreFloatAtAddress(floored, addr2)
}

func (eCtx *ECtx) HandleCeil() {
	quad := eCtx.GetCurrentQuad()
	addr1 := eCtx.GetDerefed(quad.Op1)
	addr2 := eCtx.GetDerefed(quad.Destination)

	var ceiled float64

	val := eCtx.GetFloatFromAddress(addr1)
	ceiled = math.Ceil(val)

	eCtx.StoreFloatAtAddress(ceiled, addr2)
}

func (eCtx *ECtx) HandleAbs() {
	quad := eCtx.GetCurrentQuad()
	addr1 := eCtx.GetDerefed(quad.Op1)
	addr2 := eCtx.GetDerefed(quad.Destination)

	pType := GetPTypeFromAddress(addr1)

	if pType == src.Int {
		val := eCtx.GetIntFromAddress(addr1)
		absed := math.Abs(float64(val))
		eCtx.StoreIntAtAddress(int64(absed), addr2)
	} else if pType == src.Float {
		val := eCtx.GetFloatFromAddress(addr1)
		absed := math.Abs(val)
		eCtx.StoreFloatAtAddress(absed, addr2)
	}
}

func (eCtx *ECtx) HandleNot() {
	quad := eCtx.GetCurrentQuad()
	addr1 := eCtx.GetDerefed(quad.Op1)
	addr2 := eCtx.GetDerefed(quad.Destination)

	val := eCtx.GetBoolFromAddress(addr1)
	eCtx.StoreBoolAtAddress(!val, addr2)
}
