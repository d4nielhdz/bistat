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

func (l *bistatListener) EnterLogicOperator(ctx *parser.LogicOperatorContext) {
	fmt.Println("entered logic operator")
	fmt.Println(ctx.GetText())
	l.pCtx.POperPush(int(OpFromString(ctx.GetText())))
}

func (l *bistatListener) ExitExp(ctx *parser.ExpContext) {
	fmt.Println("exited exp")
	fmt.Println(ctx.GetText())
	if !l.pCtx.POperIsEmpty() {
		oper := Op(l.pCtx.POperTop())
		if oper == And || oper == Or || oper == Gt || oper == Lt || oper == Ge || oper == Le || oper == Eq || oper == Ne {
			l.pCtx.HandleGenerateQuadForExpression()
		}
	}
}

func (l *bistatListener) EnterOpSec(ctx *parser.OpSecContext) {
	fmt.Println("entered opSec")
	fmt.Println(ctx.GetText())
	l.pCtx.POperPush(int(OpFromString(ctx.GetText())))
}

func (l *bistatListener) ExitTerm(ctx *parser.TermContext) {
	fmt.Println("exited term")
	fmt.Println(ctx.GetText())
	if !l.pCtx.POperIsEmpty() {
		oper := Op(l.pCtx.POperTop())
		if oper == Sum || oper == Subtraction || oper == UnaryMinus {
			l.pCtx.HandleGenerateQuadForExpression()
		}
	}
}

func (l *bistatListener) EnterOpFirst(ctx *parser.OpFirstContext) {
	fmt.Println("entered op first")
	fmt.Println(ctx.GetText())
	l.pCtx.POperPush(int(OpFromString(ctx.GetText())))
}

func (l *bistatListener) ExitFactor(ctx *parser.FactorContext) {
	fmt.Println("entered factor")
	fmt.Println(ctx.GetText())
	if !l.pCtx.POperIsEmpty() {
		oper := Op(l.pCtx.POperTop())
		if oper == Multiplication || oper == Division || oper == Modulus || oper == UnaryMinus {
			l.pCtx.HandleGenerateQuadForExpression()
		}
	}
}

func (l *bistatListener) EnterUnaryMinus(ctx *parser.UnaryMinusContext) {
	fmt.Println("entered unary minus")
	fmt.Println(ctx.GetText())
	l.pCtx.POperPush(int(UnaryMinus))
}

func (l *bistatListener) EnterNestedExpression(ctx *parser.NestedExpressionContext) {
	fmt.Println("entered nested expression")
	fmt.Println(ctx.GetText())
	l.pCtx.POperPush(int(UndefinedOp))
}

func (l *bistatListener) ExitNestedExpression(ctx *parser.NestedExpressionContext) {
	fmt.Println("exited nested expression")
	fmt.Println(ctx.GetText())
	l.pCtx.POperPop()
}

func (l *bistatListener) EnterSpecialFunction(ctx *parser.SpecialFunctionContext) {
	fmt.Println("entered special function")
	fmt.Println(ctx.GetText())
}

func (l *bistatListener) EnterVarCons(ctx *parser.VarConsContext) {
	fmt.Println("entered varCons")
	rType := NewRType(Bool)
	if ctx.ID() != nil {
		entry, ok := l.pCtx.GetRTypeFromVarName(ctx.ID().GetText())
		if ok {
			rType = entry
		} else {
			l.pCtx.SemanticError("Variable " + ctx.ID().GetText() + " not found in scope")
		}
	} else {
		rType = l.pCtx.GetRTypeFromVarConsContext(ctx)
	}

	fmt.Println(ctx.ID())
	l.pCtx.POPush(rType)
}

func (l *bistatListener) ExitAssignment(ctx *parser.AssignmentContext) {
	fmt.Println("exited assignment")
	fmt.Println(ctx.GetText())
}

// func (l *bistatListener) EnterSpecialFunction(ctx *parser.SpecialFunctionContext) {
// 	fmt.Println("Entered special function")
// }

// func (l *bistatListener) EnterAssignment(ctx *parser.AssignmentContext) {
// 	fmt.Println("Entered assignment")
// 	fmt.Println(ctx.VAR_CONS())
// }

func (l *bistatListener) ExitProgram(ctx *parser.ProgramContext) {
	fmt.Println("Exited program")
	l.pCtx.PrintAddrTable()
	l.pCtx.PrintQuads()
}
