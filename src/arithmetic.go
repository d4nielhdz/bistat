package src

import (
	parser "bistat/parser"
	"fmt"
)

func (l *bistatListener) EnterLogicOperator(ctx *parser.LogicOperatorContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	fmt.Println("entered logic operator")
	fmt.Println(ctx.GetText())
	l.pCtx.POperPush(int(OpFromString(ctx.GetText())))
}

func (l *bistatListener) ExitExp(ctx *parser.ExpContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

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
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	fmt.Println("entered opSec")
	fmt.Println(ctx.GetText())
	l.pCtx.POperPush(int(OpFromString(ctx.GetText())))
}

func (l *bistatListener) ExitTerm(ctx *parser.TermContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

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
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	fmt.Println("entered op first")
	fmt.Println(ctx.GetText())
	l.pCtx.POperPush(int(OpFromString(ctx.GetText())))
}

func (l *bistatListener) ExitFactor(ctx *parser.FactorContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

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
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	fmt.Println("entered unary minus")
	fmt.Println(ctx.GetText())
	l.pCtx.POperPush(int(UnaryMinus))
}

func (l *bistatListener) EnterNestedExpression(ctx *parser.NestedExpressionContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	fmt.Println("entered nested expression")
	fmt.Println(ctx.GetText())
	l.pCtx.POperPush(int(UndefinedOp))
}

func (l *bistatListener) ExitNestedExpression(ctx *parser.NestedExpressionContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	fmt.Println("exited nested expression")
	fmt.Println(ctx.GetText())
	l.pCtx.POperPop()
}
