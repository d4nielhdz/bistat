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
	vt := ctx.Var_type()
	pType := PTypeFromString(vt.TYPE_PRIMITIVE().GetText())
	addrMgr := l.pCtx.GetCorrespondingAddressManager(pType)
	resolved := l.pCtx.ResolveType(vt, addrMgr)
	l.pCtx.AddVarToScope("main", funcName, resolved)
	l.pCtx.AddToAddrTable(resolved.address, ctx.ID().GetText())
	l.pCtx.AddScope(funcName)
	l.pCtx.AddFunction(funcName, RTypeToFuncData(resolved))
	params := make([]RType, 0)
	localIntVars := 0
	localFloatVars := 0
	localBoolVars := 0
	localStringVars := 0

	for _, p := range ctx.AllParamDeclaration() {
		pType := PTypeFromString(p.Var_type().TYPE_PRIMITIVE().GetText())
		addrMgr := l.pCtx.GetCorrespondingAddressManager(pType)
		resolved := l.pCtx.ResolveType(p.Var_type(), addrMgr)
		params = append(params, resolved)
		switch pType {
		case Int:
			localIntVars += 1
			break
		case Float:
			localFloatVars += 1
			break
		case String:
			localStringVars += 1
			break
		default:
			localBoolVars += 1
		}
	}

	l.pCtx.AddParams(funcName, params)
	data := l.pCtx.funcDir[funcName]
	data.params = len(params)
	data.localStringVars = localStringVars
	data.localBoolVars = localBoolVars
	data.localFloatVars = localFloatVars
	data.localIntVars = localIntVars
	data.idx = len(l.pCtx.functions)
	l.pCtx.functions = append(l.pCtx.functions, funcName)
	l.pCtx.funcDir[funcName] = data
}

func (l *bistatListener) EnterParamDeclaration(ctx *parser.ParamDeclarationContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	vt := ctx.Var_type()
	pType := PTypeFromString(vt.TYPE_PRIMITIVE().GetText())
	if pType == Void {
		l.pCtx.SemanticError("Can't have void parameters (" + ctx.ID().GetText() + ")")
		return
	}
	addrMgr := l.pCtx.GetCorrespondingAddressManager(pType)
	resolved := l.pCtx.ResolveType(vt, addrMgr)
	l.pCtx.AddVarToScope(l.pCtx.GetCurrentScope(), ctx.ID().GetText(), resolved)
	l.pCtx.AddToAddrTable(resolved.address, ctx.ID().GetText())
}

func (l *bistatListener) EnterFuncBlockStart(ctx *parser.FuncBlockStartContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}
	funcName := l.pCtx.GetCurrentScope()
	data := l.pCtx.funcDir[funcName]
	data.localIntVars = data.localIntVars + l.pCtx.vm.localIntAddressMgr.GetSize()
	data.localFloatVars = data.localFloatVars + l.pCtx.vm.localFloatAddressMgr.GetSize()
	data.localBoolVars = data.localBoolVars + l.pCtx.vm.localBoolAddressMgr.GetSize()
	data.localStringVars = data.localStringVars + l.pCtx.vm.localStringAddressMgr.GetSize()

	data.funcStart = len(l.pCtx.vm.Quads())
	l.pCtx.funcDir[funcName] = data
}

func (l *bistatListener) EnterFuncBlockEnd(ctx *parser.FuncBlockEndContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}
	funcName := l.pCtx.GetCurrentScope()
	data := l.pCtx.funcDir[funcName]
	data.tempIntVars = data.tempIntVars + l.pCtx.vm.tempIntAddressMgr.GetSize()
	data.tempFloatVars = data.tempFloatVars + l.pCtx.vm.tempFloatAddressMgr.GetSize()
	data.tempBoolVars = data.tempBoolVars + l.pCtx.vm.tempBoolAddressMgr.GetSize()
	data.tempStringVars = data.tempStringVars + l.pCtx.vm.tempStringAddressMgr.GetSize()
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
}

func (l *bistatListener) ExitFunctionCall(ctx *parser.FunctionCallContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	funcName := ctx.ID().GetText()
	data, found := l.pCtx.funcDir[funcName]
	if !found {
		l.pCtx.SemanticError("Function " + funcName + " was not defined")
		return
	}
	l.pCtx.vm.PushQuad(NewQuad(Era, data.idx, -1, -1))

	numArgs := len(ctx.AllExpression())
	if numArgs != data.params {
		l.pCtx.SemanticError("Number of arguments in function call to " + funcName + " doesn't match the number of parameters with which it was defined")
		return
	}
	for numArgs != 0 {
		rType := l.pCtx.POTop()
		paramtype := l.pCtx.paramTable[funcName][numArgs-1]
		// todo: check array coherence
		if paramtype.pType != rType.pType {
			l.pCtx.SemanticError("Type mismatch between argument and parameter in argument #" + strconv.Itoa(numArgs) + " in function call to " + funcName)
			return
		}
		l.pCtx.vm.PushQuad(NewQuad(Param, rType.address, numArgs-1, -1))
		numArgs--
		l.pCtx.POPop()
	}
	l.pCtx.vm.PushQuad(NewQuad(GoSub, data.idx, -1, -1))
}

// todo: handle return stmt
