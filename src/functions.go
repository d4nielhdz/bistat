package src

import (
	parser "bistat/parser"
	"strconv"
)

func (l *bistatListener) EnterFuncDef(ctx *parser.FuncDefContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	funcName := ctx.ID().GetText()
	pType := PTypeFromString(ctx.TYPE_PRIMITIVE().GetText())
	addrMgr := l.pCtx.GetCorrespondingAddressManager(pType)
	resolved := NewRType(pType)
	addr, ok := addrMgr.GetNext()
	if !ok {
		l.pCtx.SemanticError("Out of memory")
		return
	}
	resolved.Address = addr
	l.pCtx.AddVarToScope("main", funcName, resolved)
	l.pCtx.AddToAddrTable(resolved.Address, ctx.ID().GetText())
	l.pCtx.AddScope(funcName)
	l.pCtx.AddFunction(funcName, RTypeToFuncData(resolved))
	params := make([]RType, 0)
	localIntVars := 0
	localFloatVars := 0
	localBoolVars := 0
	localStringVars := 0

	for _, p := range ctx.AllParamDeclaration() {
		pType := PTypeFromString(p.TYPE_PRIMITIVE().GetText())
		if pType == Void {
			l.pCtx.SemanticError("Can't have void parameters (" + p.ID().GetText() + ")")
			return
		}
		addrMgr := l.pCtx.GetCorrespondingAddressManager(pType)
		resolved := NewRType(pType)
		if resolved.FirstDim > 0 {
			l.pCtx.SemanticError("Can't use arrays as parameters in function definition")
			return
		}
		addr, ok := addrMgr.GetNext()
		if !ok {
			l.pCtx.SemanticError("Out of memory")
			return
		}
		resolved.Address = addr
		l.pCtx.AddVarToScope(l.pCtx.GetCurrentScope(), p.ID().GetText(), resolved)
		l.pCtx.AddToAddrTable(resolved.Address, p.ID().GetText())

		params = append(params, resolved)
		switch pType {
		case Int:
			localIntVars += 1
			// break
		case Float:
			localFloatVars += 1
			// break
		case String:
			localStringVars += 1
			// break
		default:
			localBoolVars += 1
		}
	}

	l.pCtx.AddParams(funcName, params)
	data := l.pCtx.funcDir[funcName]
	data.Params = len(params)
	data.LocalStringVars = localStringVars
	data.LocalBoolVars = localBoolVars
	data.LocalFloatVars = localFloatVars
	data.LocalIntVars = localIntVars
	data.Idx = len(l.pCtx.functions)
	l.pCtx.functions = append(l.pCtx.functions, funcName)
	l.pCtx.funcDir[funcName] = data
}

func (l *bistatListener) EnterFuncBlockStart(ctx *parser.FuncBlockStartContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}
	funcName := l.pCtx.GetCurrentScope()
	data := l.pCtx.funcDir[funcName]
	data.LocalIntVars = data.LocalIntVars + l.pCtx.vm.localIntAddressMgr.GetSize()
	data.LocalFloatVars = data.LocalFloatVars + l.pCtx.vm.localFloatAddressMgr.GetSize()
	data.LocalBoolVars = data.LocalBoolVars + l.pCtx.vm.localBoolAddressMgr.GetSize()
	data.LocalStringVars = data.LocalStringVars + l.pCtx.vm.localStringAddressMgr.GetSize()

	data.FuncStart = len(l.pCtx.vm.Quads())
	l.pCtx.funcDir[funcName] = data
}

func (l *bistatListener) EnterFuncBlockEnd(ctx *parser.FuncBlockEndContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}
	funcName := l.pCtx.GetCurrentScope()
	data := l.pCtx.funcDir[funcName]
	data.LocalRefs = l.pCtx.vm.localRefAddressMgr.GetSize()
	data.TempIntVars = data.TempIntVars + l.pCtx.vm.tempIntAddressMgr.GetSize()
	data.TempFloatVars = data.TempFloatVars + l.pCtx.vm.tempFloatAddressMgr.GetSize()
	data.TempBoolVars = data.TempBoolVars + l.pCtx.vm.tempBoolAddressMgr.GetSize()
	data.TempStringVars = data.TempStringVars + l.pCtx.vm.tempStringAddressMgr.GetSize()
	l.pCtx.funcDir[funcName] = data

	l.pCtx.RemoveFunction(funcName)
	l.pCtx.PopCurrentScope()
	l.pCtx.vm.PushQuad(NewQuad(EndFunc, -1, -1, -1))
	l.pCtx.vm.localBoolAddressMgr.Reset()
	l.pCtx.vm.localStringAddressMgr.Reset()
	l.pCtx.vm.localIntAddressMgr.Reset()
	l.pCtx.vm.localFloatAddressMgr.Reset()
	l.pCtx.vm.tempBoolAddressMgr.Reset()
	l.pCtx.vm.tempStringAddressMgr.Reset()
	l.pCtx.vm.tempIntAddressMgr.Reset()
	l.pCtx.vm.tempFloatAddressMgr.Reset()
	l.pCtx.vm.localRefAddressMgr.Reset()
}

func (l *bistatListener) ExitReturnStmt(ctx *parser.ReturnStmtContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}
	o := l.pCtx.POTop()
	l.pCtx.POPop()
	funcName := l.pCtx.GetCurrentScope()
	fVar, _ := l.pCtx.GetVarInScope("main", funcName)
	if o.FirstDim > 0 || o.SecondDim > 0 {
		l.pCtx.SemanticError("Can't return an array or matrix in a function")
	}
	if o.PType != fVar.PType {
		l.pCtx.SemanticError("Incorrect return type for function " + funcName + ", expected " + PTypeToString(o.PType) + ", found " + PTypeToString(fVar.PType))
		return
	}
	l.pCtx.vm.PushQuad(NewQuad(Return, fVar.Address, -1, o.Address))
}
func (l *bistatListener) EnterFunctionCall(ctx *parser.FunctionCallContext) {
	l.pCtx.POperPush(int(Other))
}

func (l *bistatListener) ExitFunctionCall(ctx *parser.FunctionCallContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}
	l.pCtx.POperPop()

	funcName := ctx.ID().GetText()
	data, found := l.pCtx.funcDir[funcName]
	if !found {
		l.pCtx.SemanticError("Function " + funcName + " was not defined")
		return
	}
	l.pCtx.vm.PushQuad(NewQuad(Era, data.Idx, -1, -1))

	numArgs := len(ctx.AllExpression())
	if numArgs != data.Params {
		l.pCtx.SemanticError("Number of arguments in function call to " + funcName + " doesn't match the number of parameters with which it was defined")
		return
	}
	for numArgs != 0 {
		rType := l.pCtx.POTop()
		param := l.pCtx.paramTable[funcName][numArgs-1]
		if param.PType != rType.PType {
			l.pCtx.SemanticError("Type mismatch between argument and parameter in argument #" + strconv.Itoa(numArgs) + " in function call to " + funcName)
			return
		}
		l.pCtx.vm.PushQuad(NewQuad(Param, param.Address, -1, rType.Address))
		numArgs--
		l.pCtx.POPop()
	}
	l.pCtx.vm.PushQuad(NewQuad(GoSub, data.Idx, -1, -1))
	if data.PType != Void {
		fVar, _ := l.pCtx.GetVarInScope("main", funcName)
		addrMgr := l.pCtx.GetCorrespondingTempAddressManager(fVar.PType)
		addr, _ := addrMgr.GetNext()
		tmp := NewRType(fVar.PType)
		tmp.FirstDim = fVar.FirstDim
		tmp.SecondDim = fVar.SecondDim
		tmp.Address = addr
		l.pCtx.vm.PushQuad(NewQuad(Assign, addr, -1, fVar.Address))
		l.pCtx.POPush(tmp)
	}
}
