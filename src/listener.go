package src

import (
	parser "bistat/parser"
	"encoding/gob"
	"fmt"
	"os"
)

type bistatListener struct {
	*parser.BaseBistatListener
	pCtx PCtx
}

func NewBistatListener() *bistatListener {
	l := new(bistatListener)
	l.pCtx = NewPCtx()
	return l
}

func (l *bistatListener) EnterProgram(ctx *parser.ProgramContext) {
	l.pCtx.AddScope("main")
	l.pCtx.AddFunction("main", NewFuncData(Void))
	fData := l.pCtx.funcDir["main"]
	fData.Idx = 0
	l.pCtx.funcDir["main"] = fData
	l.pCtx.vm.PushQuad(NewQuad(Goto, 0, -1, -1))
}

func (l *bistatListener) EnterVarDeclaration(ctx *parser.VarDeclarationContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	// todo: catch type errors
	currScope := l.pCtx.GetCurrentScope()
	vt := ctx.Var_type()
	pType := PTypeFromString(vt.TYPE_PRIMITIVE().GetText())
	if pType == Void {
		l.pCtx.SemanticError("Can't declare void variables")
		return
	}
	addrMgr := l.pCtx.GetCorrespondingAddressManager(pType)
	resolved := l.pCtx.ResolveType(vt, addrMgr)
	varName := ctx.ID().GetText()
	_, exists := l.pCtx.GetVarInScope(currScope, varName)
	if exists {
		l.pCtx.SemanticError("Variable " + varName + " already declared")
		return
	} else {
		l.pCtx.AddVarToScope(currScope, varName, resolved)
		l.pCtx.AddToAddrTable(resolved.Address, ctx.ID().GetText())
	}
}

func (l *bistatListener) ExitAssignment(ctx *parser.AssignmentContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	rRType := l.pCtx.POTop()
	l.pCtx.POPop()
	lRType, found := l.pCtx.GetRTypeFromVarName(ctx.ID().GetText())
	if !found {
		l.pCtx.SemanticError("Variable " + ctx.ID().GetText() + " not defined")
		return
	}
	if lRType.PType != rRType.PType {
		// todo: cast where appropriate
		l.pCtx.SemanticError("Cannot assign to " + ctx.ID().GetText() + " because of type mismatch: " + PTypeToString(rRType.PType) + " != " + PTypeToString(lRType.PType))
		return
	}
	quad := NewQuad(Assign, rRType.Address, -1, lRType.Address)
	l.pCtx.vm.PushQuad(quad)
}

func (l *bistatListener) ExitIndexing(ctx *parser.IndexingContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	varName := ctx.ID().GetText()
	arr, found := l.pCtx.GetRTypeFromVarName(varName)
	if !found {
		l.pCtx.SemanticError("Variable " + varName + " was not declared")
		return
	}
	if arr.FirstDim == 0 {
		l.pCtx.SemanticError("Variable " + varName + " isn't a matrix or array")
		return
	}
	levels := len(ctx.AllExpression())
	if levels == 1 {
		idx := l.pCtx.POTop()
		l.pCtx.POPop()
		l.pCtx.vm.PushQuad(NewQuad(Verify, idx.Address, arr.Address, arr.FirstDim))
		// addr is a ref
		addr, _ := l.pCtx.GetCorrespondingTempAddressManager(arr.PType).GetNext()
		indexed := NewRType(arr.PType)
		if arr.SecondDim != 0 {
			l.pCtx.vm.PushQuad(NewQuad(Multiplication, arr.Address, idx.Address, addr))
			indexed.FirstDim = arr.SecondDim
		} else {
			l.pCtx.vm.PushQuad(NewQuad(Sum, arr.Address, idx.Address, addr))
		}
		indexed.Address = addr
		indexed.IsRef = true

		// todo: figure out how/when to calculate end address
		l.pCtx.POPush(indexed)
	} else {
		if arr.SecondDim == 0 {
			l.pCtx.SemanticError("Variable " + varName + " isn't a matrix")
			return
		}
		secondIdx := l.pCtx.POTop()
		l.pCtx.POPop()
		firstIdx := l.pCtx.POTop()
		l.pCtx.POPop()
		l.pCtx.vm.PushQuad(NewQuad(Verify, firstIdx.Address, arr.Address, arr.FirstDim))
		l.pCtx.vm.PushQuad(NewQuad(Verify, secondIdx.Address, arr.Address, arr.SecondDim))

		// addr is a ref
		addr, _ := l.pCtx.GetCorrespondingTempAddressManager(arr.PType).GetNext()
		l.pCtx.vm.PushQuad(NewQuad(Multiplication, arr.Address, firstIdx.Address, addr))
		secondAddr, _ := l.pCtx.vm.tempIntAddressMgr.GetNext()
		l.pCtx.vm.PushQuad(NewQuad(Sum, addr, secondIdx.Address, secondAddr))

		indexed := NewRType(arr.PType)
		indexed.Address = addr
		indexed.IsRef = true
		// todo: figure out how/when to calculate end address
		l.pCtx.POPush(indexed)
	}
}

func (l *bistatListener) EnterMain(ctx *parser.MainContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}
	fData := l.pCtx.funcDir["main"]
	fData.FuncStart = len(l.pCtx.vm.Quads())
	l.pCtx.funcDir["main"] = fData
	l.pCtx.vm.quads[0].Op1 = len(l.pCtx.vm.Quads())
}

func (l *bistatListener) ExitProgram(ctx *parser.ProgramContext) {
	l.pCtx.vm.PushQuad(NewQuad(End, -1, -1, -1))
	l.pCtx.PrintAddrTable()
	l.pCtx.PrintQuads()
	l.pCtx.PrintErrors()
	if len(l.pCtx.semanticErrors) == 0 {
		RegisterTypes()
		file, err := os.Create("./obj.gob")
		objCode := NewObjCode(l.pCtx.funcDir, l.pCtx.consTable, l.pCtx.vm.quads)
		objCode.BoolSize = l.pCtx.vm.globalBoolAddressMgr.GetSize()
		objCode.StringSize = l.pCtx.vm.globalStringAddressMgr.GetSize()
		objCode.FloatSize = l.pCtx.vm.globalFloatAddressMgr.GetSize()
		objCode.IntSize = l.pCtx.vm.globalIntAddressMgr.GetSize()
		objCode.ConstBoolSize = l.pCtx.vm.constBoolAddressMgr.GetSize()
		objCode.ConstStringSize = l.pCtx.vm.constStringAddressMgr.GetSize()
		objCode.ConstFloatSize = l.pCtx.vm.constFloatAddressMgr.GetSize()
		objCode.ConstIntSize = l.pCtx.vm.constIntAddressMgr.GetSize()
		objCode.TempBoolSize = l.pCtx.vm.tempBoolAddressMgr.GetSize()
		objCode.TempStringSize = l.pCtx.vm.tempStringAddressMgr.GetSize()
		objCode.TempFloatSize = l.pCtx.vm.tempFloatAddressMgr.GetSize()
		objCode.TempIntSize = l.pCtx.vm.tempIntAddressMgr.GetSize()
		if err == nil {
			encoder := gob.NewEncoder(file)
			encoder.Encode(objCode)
		} else {
			fmt.Println(err)
		}
		file.Close()
	}
}

func RegisterTypes() {
	gob.Register(ObjCode{})
	gob.Register(FuncData{})
	gob.Register(RType{})
	gob.Register(Quad{})
}
