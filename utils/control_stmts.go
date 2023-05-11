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
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	end := l.pCtx.JumpsTop()
	l.pCtx.JumpsPop()
	ret := l.pCtx.JumpsTop()
	l.pCtx.JumpsPop()

	l.pCtx.vm.PushQuad(NewQuad(Goto, -1, -1, ret))
	l.pCtx.vm.quads[end].destination = len(l.pCtx.vm.quads)

}

func (l *bistatListener) EnterForHeader(ctx *parser.ForHeaderContext) {
	idxAddr, _ := l.pCtx.vm.tempIntAddressMgr.GetNext()
	rType := NewRType(Int)
	rType.address = idxAddr
	entry, found := l.pCtx.consTable["0"]
	if !found {
		next, okAddr := l.pCtx.vm.constIntAddressMgr.GetNext()
		if !okAddr {
			l.pCtx.SemanticError("Out of memory")
		}
		zero := NewRType(Int)
		zero.firstDim = 0
		zero.secondDim = 0
		zero.address = next
		l.pCtx.consTable["0"] = zero
		l.pCtx.addrTable[zero.address] = "0"
		entry = zero
	}
	l.pCtx.vm.PushQuad(NewQuad(Assign, entry.address, -1, idxAddr))
	l.pCtx.POPush(rType)
}

func (l *bistatListener) ExitForHeader(ctx *parser.ForHeaderContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	varName := ctx.ID().GetText()
	_, found := l.pCtx.GetRTypeFromVarName(varName)
	if !found {
		l.pCtx.SemanticError("Variable " + varName + " not found")
		return
	}
	o := l.pCtx.POTop()
	l.pCtx.POPop()
	vc := l.pCtx.POTop()
	if o.pType != vc.pType {
		// todo: validate array type
		l.pCtx.SemanticError("Control variable must be the same type as expression variable")
		return
	}
	lenAddr, _ := l.pCtx.vm.tempIntAddressMgr.GetNext()
	boolAddr, _ := l.pCtx.vm.tempBoolAddressMgr.GetNext()
	l.pCtx.JumpsPush(len(l.pCtx.vm.Quads()))
	l.pCtx.vm.PushQuad(NewQuad(Length, o.address, -1, lenAddr))
	l.pCtx.vm.PushQuad(NewQuad(Lt, vc.address, lenAddr, boolAddr))
	currQuad := len(l.pCtx.vm.quads)
	l.pCtx.JumpsPush(currQuad)
	l.pCtx.vm.PushQuad(NewQuad(GotoF, boolAddr, -1, -1))

}

func (l *bistatListener) ExitForLoop(ctx *parser.ForLoopContext) {
	vc := l.pCtx.POTop()
	entry, found := l.pCtx.consTable["1"]
	if !found {
		next, okAddr := l.pCtx.vm.constIntAddressMgr.GetNext()
		if !okAddr {
			l.pCtx.SemanticError("Out of memory")
		}
		one := NewRType(Int)
		one.firstDim = 0
		one.secondDim = 0
		one.address = next
		l.pCtx.consTable["1"] = one
		l.pCtx.addrTable[one.address] = "1"
		entry = one
	}
	l.pCtx.vm.PushQuad(NewQuad(Sum, vc.address, entry.address, vc.address))
	l.pCtx.POPop()
	end := l.pCtx.JumpsTop()
	l.pCtx.JumpsPop()
	ret := l.pCtx.JumpsTop()
	l.pCtx.JumpsPop()
	l.pCtx.vm.PushQuad(NewQuad(Goto, -1, -1, ret))
	l.pCtx.vm.quads[end].destination = len(l.pCtx.vm.Quads())
}
