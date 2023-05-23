package main

import (
	src "bistat/src"
	"fmt"
)

type ECtx struct {
	IP           int
	Quads        []src.Quad
	FuncDir      map[string]src.FuncData
	ConsTable    map[string]src.RType
	ConstMemory  *MemorySegment
	TempMemory   *MemorySegment
	GlobalMemory *MemorySegment
	StackSegment []StackSegment
	Errors       []string
}

func NewEctx(objCode src.ObjCode) ECtx {
	return ECtx{
		IP:           0,
		Quads:        objCode.Quads,
		FuncDir:      objCode.FuncDir,
		ConsTable:    objCode.ConsTable,
		ConstMemory:  NewMemorySegment(objCode.ConstFloatSize, objCode.ConstIntSize, objCode.ConstStringSize, objCode.ConstBoolSize, src.CONST_FLOAT_START, src.CONST_INT_START, src.CONST_STRING_START, src.CONST_BOOL_START),
		TempMemory:   NewMemorySegment(objCode.TempFloatSize, objCode.TempIntSize, objCode.TempStringSize, objCode.TempBoolSize, src.TEMP_FLOAT_START, src.TEMP_INT_START, src.TEMP_STRING_START, src.TEMP_BOOL_START),
		GlobalMemory: NewMemorySegment(objCode.FloatSize, objCode.IntSize, objCode.StringSize, objCode.BoolSize, src.GLOBAL_FLOAT_START, src.GLOBAL_INT_START, src.GLOBAL_STRING_START, src.GLOBAL_BOOL_START),
		StackSegment: make([]StackSegment, 0),
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
	eCtx.LoadConstMemory()
	eCtx.ConstMemory.Print()
	for eCtx.IP < len(eCtx.Quads) {
		if len(eCtx.Errors) > 0 {
			return
		}
		fmt.Println("Running quad", eCtx.IP)
		eCtx.HandleQuad()
	}
	fmt.Println("Global")
	eCtx.GlobalMemory.Print()
	fmt.Println("Temp")

	eCtx.TempMemory.Print()

}

func (eCtx *ECtx) GetCurrentQuad() src.Quad {
	return eCtx.Quads[eCtx.IP]
}

func (eCtx *ECtx) GetStackSegment() *StackSegment {
	return &eCtx.StackSegment[len(eCtx.StackSegment)-1]
}

func (eCtx *ECtx) HandleQuad() {
	quad := eCtx.GetCurrentQuad()
	switch quad.Op {
	case src.Goto:
		eCtx.IP = quad.Op1
		return
	case src.Assign:
		eCtx.HandleAssign()
	case src.Sum:
		eCtx.HandleSum()
	case src.Subtraction:
		eCtx.HandleSubtraction()
	case src.Multiplication:
		eCtx.HandleMultiplication()
	case src.Division:
		eCtx.HandleDivision()
	case src.Modulus:
		eCtx.HandleModulus()
	case src.Gt:
		eCtx.HandleGt()
	case src.Lt:
		eCtx.HandleLt()
	case src.Ge:
		eCtx.HandleGe()
	case src.Le:
		eCtx.HandleLe()
	case src.Eq:
		eCtx.HandleEq()
	case src.Ne:
		eCtx.HandleNe()
	default:
		fmt.Println("Unhandled ", src.OpToString(quad.Op))
	}
	eCtx.IP++
}
