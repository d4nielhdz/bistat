package src

import (
	parser "bistat/parser"
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
	fData.idx = 0
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
		l.pCtx.AddToAddrTable(resolved.address, ctx.ID().GetText())
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
	if lRType.pType != rRType.pType {
		// todo: cast where appropriate
		l.pCtx.SemanticError("Cannot assign to " + ctx.ID().GetText() + " because of type mismatch: " + PTypeToString(rRType.pType) + " != " + PTypeToString(lRType.pType))
		return
	}
	quad := NewQuad(Assign, rRType.address, -1, lRType.address)
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
	if arr.firstDim == 0 {
		l.pCtx.SemanticError("Variable " + varName + " isn't a matrix or array")
		return
	}
	levels := len(ctx.AllExpression())
	if levels == 1 {
		idx := l.pCtx.POTop()
		l.pCtx.POPop()
		l.pCtx.vm.PushQuad(NewQuad(Verify, idx.address, arr.address, arr.firstDim))
		// addr is a ref
		addr, _ := l.pCtx.vm.tempIntAddressMgr.GetNext()
		indexed := NewRType(arr.pType)
		if arr.secondDim != 0 {
			l.pCtx.vm.PushQuad(NewQuad(Multiplication, arr.address, idx.address, addr))
			indexed.firstDim = arr.secondDim
		} else {
			l.pCtx.vm.PushQuad(NewQuad(Sum, arr.address, idx.address, addr))
		}
		indexed.address = addr
		indexed.isRef = true

		// todo: figure out how/when to calculate end address
		l.pCtx.POPush(indexed)
	} else {
		if arr.secondDim == 0 {
			l.pCtx.SemanticError("Variable " + varName + " isn't a matrix")
			return
		}
		secondIdx := l.pCtx.POTop()
		l.pCtx.POPop()
		firstIdx := l.pCtx.POTop()
		l.pCtx.POPop()
		l.pCtx.vm.PushQuad(NewQuad(Verify, firstIdx.address, arr.address, arr.firstDim))
		l.pCtx.vm.PushQuad(NewQuad(Verify, secondIdx.address, arr.address, arr.secondDim))

		// addr is a ref
		addr, _ := l.pCtx.vm.tempIntAddressMgr.GetNext()
		l.pCtx.vm.PushQuad(NewQuad(Multiplication, arr.address, firstIdx.address, addr))
		secondAddr, _ := l.pCtx.vm.tempIntAddressMgr.GetNext()
		l.pCtx.vm.PushQuad(NewQuad(Sum, addr, secondIdx.address, secondAddr))

		indexed := NewRType(arr.pType)
		indexed.address = addr
		indexed.isRef = true
		// todo: figure out how/when to calculate end address
		l.pCtx.POPush(indexed)
	}
}

func (l *bistatListener) EnterMain(ctx *parser.MainContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}
	fData := l.pCtx.funcDir["main"]
	fData.funcStart = len(l.pCtx.vm.Quads())
	l.pCtx.funcDir["main"] = fData
	l.pCtx.vm.quads[0].op1 = len(l.pCtx.vm.Quads())
}

func (l *bistatListener) ExitProgram(ctx *parser.ProgramContext) {
	l.pCtx.vm.PushQuad(NewQuad(End, -1, -1, -1))
	l.pCtx.PrintAddrTable()
	l.pCtx.PrintQuads()
	l.pCtx.PrintErrors()
}
