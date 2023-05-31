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
	Jumps        []int
	Functions    []string
	Errors       []string
}

func NewEctx(objCode src.ObjCode) ECtx {
	return ECtx{
		IP:           0,
		Quads:        objCode.Quads,
		FuncDir:      objCode.FuncDir,
		ConsTable:    objCode.ConsTable,
		ConstMemory:  NewMemorySegment(objCode.ConstFloatSize, objCode.ConstIntSize, objCode.ConstStringSize, objCode.ConstBoolSize, 0, src.CONST_FLOAT_START, src.CONST_INT_START, src.CONST_STRING_START, src.CONST_BOOL_START, 0),
		TempMemory:   NewMemorySegment(objCode.TempFloatSize, objCode.TempIntSize, objCode.TempStringSize, objCode.TempBoolSize, 0, src.TEMP_FLOAT_START, src.TEMP_INT_START, src.TEMP_STRING_START, src.TEMP_BOOL_START, 0),
		GlobalMemory: NewMemorySegment(objCode.FloatSize, objCode.IntSize, objCode.StringSize, objCode.BoolSize, objCode.GlobalRefSize, src.GLOBAL_FLOAT_START, src.GLOBAL_INT_START, src.GLOBAL_STRING_START, src.GLOBAL_BOOL_START, src.GLOBAL_REF_START),
		StackSegment: make([]StackSegment, 0),
		Errors:       make([]string, 0),
		Functions:    objCode.Functions,
	}
}

func (eCtx *ECtx) PushStackSegment(stackSegment *StackSegment) {
	eCtx.StackSegment = append(eCtx.StackSegment, *stackSegment)
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
	// i := 0
	// eCtx.ConstMemory.Print()
	for eCtx.IP < len(eCtx.Quads) {
		if len(eCtx.Errors) > 0 {
			return
		}
		// fmt.Println("Running quad", eCtx.IP)
		eCtx.HandleQuad()
		if len(eCtx.StackSegment) > 1000000 {
			eCtx.NewError("Stack overflow")
			return
		}
		// i++
	}
	// fmt.Println("Global")
	// eCtx.GlobalMemory.Print()
	// fmt.Println("Temp")

	// eCtx.TempMemory.Print()

}

func (eCtx *ECtx) GetCurrentQuad() src.Quad {
	return eCtx.Quads[eCtx.IP]
}

func (eCtx *ECtx) GetStackSegment() *StackSegment {
	return &eCtx.StackSegment[len(eCtx.StackSegment)-1]
}

func (eCtx *ECtx) GetPreviousStackSegment() *StackSegment {
	return &eCtx.StackSegment[len(eCtx.StackSegment)-2]
}

func (eCtx *ECtx) HandleQuad() {
	quad := eCtx.GetCurrentQuad()
	if quad.Op != src.Print && quad.Op != src.PrintN {
		// fmt.Println("#", eCtx.IP)
	}
	switch quad.Op {
	case src.Goto:
		eCtx.IP = quad.Destination
		return
	case src.Assign:
		eCtx.HandleAssign()
	case src.Param:
		eCtx.HandleParam()
	case src.Return:
		eCtx.HandleReturn()
		return
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
	case src.Print:
		eCtx.HandlePrint()
	case src.PrintN:
		eCtx.HandlePrintN()
	case src.GotoF:
		eCtx.HandleGotoF()
		return
	case src.GotoT:
		eCtx.HandleGotoT()
		return
	case src.Era:
		eCtx.HandleEra()
	case src.GoSub:
		eCtx.HandleGoSub()
		return
	default:
		fmt.Println("Unhandled ", src.OpToString(quad.Op))
	}
	eCtx.IP++
}
