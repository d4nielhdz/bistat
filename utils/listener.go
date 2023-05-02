package utils

import (
	parser "bistat/parser"
	"fmt"
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
	fmt.Println("Entered program")
	fmt.Println(ctx.ID())
	l.pCtx.AddScope(ctx.ID().GetText())
	l.pCtx.AddFunction(ctx.ID().GetText(), NewRType(Void))
}

func (l *bistatListener) EnterVarDeclaration(ctx *parser.VarDeclarationContext) {
	// todo: catch type errors
	fmt.Println("Entered variable declaration")
	currScope := l.pCtx.GetCurrentScope()
	fmt.Println(ctx.ID())
	vt := ctx.Var_type()
	pType := PTypeFromString(vt.TYPE_PRIMITIVE().GetText())
	addrMgr := l.pCtx.GetCorrespondingAddressManager(pType)
	resolved := l.pCtx.ResolveType(vt, addrMgr)
	varName := ctx.ID().GetText()
	_, exists := l.pCtx.GetVarInScope(currScope, varName)
	if exists {
		fmt.Println("variable already exists")
	} else {
		l.pCtx.AddVarToScope(currScope, varName, resolved)
		l.pCtx.AddToAddrTable(resolved.address, ctx.ID().GetText())
	}
}

func (l *bistatListener) EnterParamDeclaration(ctx *parser.ParamDeclarationContext) {
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

// func (this *bistatListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
// 	fmt.Println("every rule")
// 	fmt.Println(ctx.G)
// 	fmt.Println("end")
// }

// func (l *bistatListener) EnterStmt(ctx *parser.StmtContext) {
// 	fmt.Println("Entered statement")
// }

func (l *bistatListener) EnterFuncDef(ctx *parser.FuncDefContext) {
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
	fmt.Println("entered func end")
	l.pCtx.PopCurrentScope()
	// l.pCtx.RemoveFunction()
}

// func (l *bistatListener) EnterSpecialFunction(ctx *parser.SpecialFunctionContext) {
// 	fmt.Println("Entered special function")
// }

// func (l *bistatListener) EnterAssignment(ctx *parser.AssignmentContext) {
// 	fmt.Println("Entered assignment")
// 	fmt.Println(ctx.VAR_CONS())
// }

// func (l *bistatListener) EnterReturnStmt(ctx *parser.ReturnStmtContext) {
// 	fmt.Println("Entered return stmt")
// }

// func (l *bistatListener) EnterFunctionCall(ctx *parser.FunctionCallContext) {
// 	fmt.Println("Entered function call")
// }

// func (l *bistatListener) EnterPrint(ctx *parser.PrintContext) {
// 	fmt.Println("Entered print")
// }

func (l *bistatListener) ExitProgram(ctx *parser.ProgramContext) {
	fmt.Println("Exited program")
	l.pCtx.PrintAddrTable()
}
