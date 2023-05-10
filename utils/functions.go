package utils

import (
	parser "bistat/parser"
	"fmt"
)

func (l *bistatListener) EnterParamDeclaration(ctx *parser.ParamDeclarationContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	// todo: catch type errors
	fmt.Println("Entered param declaration")
	fmt.Println(ctx.ID())
	vt := ctx.Var_type()
	pType := PTypeFromString(vt.TYPE_PRIMITIVE().GetText())
	addrMgr := l.pCtx.GetCorrespondingAddressManager(pType)
	resolved := l.pCtx.ResolveType(vt, addrMgr)
	l.pCtx.AddVarToScope(l.pCtx.GetCurrentScope(), ctx.ID().GetText(), resolved)
	l.pCtx.AddToAddrTable(resolved.address, ctx.ID().GetText())
	fmt.Println(vt.TYPE_PRIMITIVE())
	fmt.Println(vt.CARDINALITY())
	fmt.Println(resolved)
}

func (l *bistatListener) EnterFuncDef(ctx *parser.FuncDefContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	fmt.Println("Entered function definition")
	fmt.Println(ctx.ID())
	funcName := ctx.ID().GetText()
	vt := ctx.Var_type()
	pType := PTypeFromString(vt.TYPE_PRIMITIVE().GetText())
	addrMgr := l.pCtx.GetCorrespondingAddressManager(pType)
	resolved := l.pCtx.ResolveType(vt, addrMgr)
	l.pCtx.AddToAddrTable(resolved.address, ctx.ID().GetText())
	l.pCtx.AddScope(funcName)
	l.pCtx.AddFunction(funcName, resolved)
	fmt.Println(vt.TYPE_PRIMITIVE())
	fmt.Println(vt.CARDINALITY())
	fmt.Println(resolved)

}

func (l *bistatListener) EnterFuncEnd(ctx *parser.FuncEndContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	fmt.Println("entered func end")
	l.pCtx.PopCurrentScope()
	// l.pCtx.RemoveFunction()
}
