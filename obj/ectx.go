package main

import (
	src "bistat/src"
	"fmt"
)

type MemorySegment struct {
	floats  []float64
	ints    []int64
	strings []string
	bools   []bool
}

func NewMemorySegment(fSize int, iSize int, sSize int, bSize int) *MemorySegment {
	return &MemorySegment{
		floats:  make([]float64, fSize),
		ints:    make([]int64, iSize),
		strings: make([]string, sSize),
		bools:   make([]bool, bSize),
	}
}

type ECtx struct {
	IP           int
	Quads        []src.Quad
	FuncDir      map[string]src.FuncData
	ConsTable    map[string]src.RType
	ConstMemory  *MemorySegment
	GlobalMemory *MemorySegment
	StackSegment []MemorySegment
	Errors       []string
}

func NewEctx(objCode src.ObjCode) ECtx {
	return ECtx{
		IP:           0,
		Quads:        objCode.Quads,
		FuncDir:      objCode.FuncDir,
		ConsTable:    objCode.ConsTable,
		ConstMemory:  NewMemorySegment(objCode.ConstFloatSize, objCode.ConstIntSize, objCode.ConstStringSize, objCode.ConstBoolSize),
		GlobalMemory: NewMemorySegment(objCode.FloatSize, objCode.IntSize, objCode.StringSize, objCode.BoolSize),
		StackSegment: make([]MemorySegment, 0),
		Errors:       make([]string, 0),
	}
}

func (eCtx *ECtx) NewError(err string) {
	eCtx.Errors = append(eCtx.Errors, err)
}

func (eCtx *ECtx) PrintErrors() {
	for _, e := range eCtx.Errors {
		fmt.Println(e)
	}
}

func (eCtx *ECtx) Run() {

}
