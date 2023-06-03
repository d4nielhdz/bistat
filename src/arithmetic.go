package src

import (
	parser "bistat/parser"
)

func (l *bistatListener) EnterVarCons(ctx *parser.VarConsContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

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

	l.pCtx.POPush(rType)
}

func (l *bistatListener) EnterLogicOperator(ctx *parser.LogicOperatorContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	l.pCtx.POperPush(int(OpFromString(ctx.GetText())))
}

func (l *bistatListener) ExitExp(ctx *parser.ExpContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

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

	l.pCtx.POperPush(int(OpFromString(ctx.GetText())))
}

func (l *bistatListener) ExitTerm(ctx *parser.TermContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

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

	l.pCtx.POperPush(int(OpFromString(ctx.GetText())))
}

func (l *bistatListener) ExitFactor(ctx *parser.FactorContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

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

	l.pCtx.POperPush(int(UnaryMinus))
}

func (l *bistatListener) EnterNestedExpression(ctx *parser.NestedExpressionContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	l.pCtx.POperPush(int(Other))
}

func (l *bistatListener) ExitNestedExpression(ctx *parser.NestedExpressionContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	l.pCtx.POperPop()
}
