package utils

import (
	parser "bistat/parser"
)

func (l *bistatListener) EnterConditional(ctx *parser.ConditionalContext) {
	l.pCtx.PushCondJumps()
}

func (l *bistatListener) ExitCondExprEnd(ctx *parser.CondExprEndContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	po := l.pCtx.POTop()
	l.pCtx.POPop()

	if po.pType != Bool {
		l.pCtx.SemanticError("Expression type must be a boolean")
	}

	currQuad := len(l.pCtx.vm.Quads())

	l.pCtx.vm.PushQuad(NewQuad(GotoF, po.address, 0, 0))
	l.pCtx.CondJumpsPush(currQuad)
}

func (l *bistatListener) EnterElseIfStmt(ctx *parser.ElseIfStmtContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	l.pCtx.vm.PushQuad(NewQuad(Goto, 0, 0, 0))

	jump := l.pCtx.CondJumpsTop()
	currQuad := len(l.pCtx.vm.Quads())
	l.pCtx.CondJumpsPop()
	l.pCtx.vm.Quads()[jump].destination = currQuad

	l.pCtx.CondJumpsPush(currQuad - 1)
}

func (l *bistatListener) EnterElseStmt(ctx *parser.ElseStmtContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	l.pCtx.vm.PushQuad(NewQuad(Goto, 0, 0, 0))

	jump := l.pCtx.CondJumpsTop()
	currQuad := len(l.pCtx.vm.Quads())
	l.pCtx.CondJumpsPop()
	l.pCtx.vm.Quads()[jump].destination = currQuad

	l.pCtx.CondJumpsPush(currQuad - 1)
}

func (l *bistatListener) ExitConditional(ctx *parser.ConditionalContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	currQuad := len(l.pCtx.vm.Quads())
	for !l.pCtx.CondJumpsIsEmpty() {
		jump := l.pCtx.CondJumpsTop()
		l.pCtx.CondJumpsPop()
		l.pCtx.vm.Quads()[jump].destination = currQuad
	}

	l.pCtx.PopCondJumps()
}

func (l *bistatListener) EnterWhileLoop(ctx *parser.WhileLoopContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	l.pCtx.JumpsPush(len(l.pCtx.vm.Quads()))
}

func (l *bistatListener) ExitWhileExprEnd(ctx *parser.WhileExprEndContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	o := l.pCtx.POTop()
	l.pCtx.POPop()

	if o.pType != Bool {
		l.pCtx.SemanticError("Expression type must be boolean")
		return
	}

	currQuad := len(l.pCtx.vm.Quads())
	l.pCtx.vm.PushQuad(NewQuad(GotoF, o.address, -1, -1))
	l.pCtx.JumpsPush(currQuad)
}

func (l *bistatListener) ExitWhileLoop(ctx *parser.WhileLoopContext) {
	end := l.pCtx.JumpsTop()
	l.pCtx.JumpsPop()
	ret := l.pCtx.JumpsTop()
	l.pCtx.JumpsPop()

	l.pCtx.vm.PushQuad(NewQuad(Goto, -1, -1, ret))
	l.pCtx.vm.quads[end].destination = len(l.pCtx.vm.quads)

}
