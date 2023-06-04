package src

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

	if po.PType != Bool {
		l.pCtx.SemanticError("Expression type must be a boolean")
	}

	currQuad := len(l.pCtx.vm.Quads())

	l.pCtx.vm.PushQuad(NewQuad(GotoF, po.Address, 0, 0))
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
	l.pCtx.vm.Quads()[jump].Destination = currQuad

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
	l.pCtx.vm.Quads()[jump].Destination = currQuad

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
		l.pCtx.vm.Quads()[jump].Destination = currQuad
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

	if o.PType != Bool {
		l.pCtx.SemanticError("Expression type must be boolean")
		return
	}

	currQuad := len(l.pCtx.vm.Quads())
	l.pCtx.vm.PushQuad(NewQuad(GotoF, o.Address, -1, -1))
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
	l.pCtx.vm.quads[end].Destination = len(l.pCtx.vm.quads)

}

func (l *bistatListener) EnterForHeader(ctx *parser.ForHeaderContext) {
	idxAddr, _ := l.pCtx.vm.tempIntAddressMgr.GetNext()
	rType := NewRType(Int)
	rType.Address = idxAddr
	entry, found := l.pCtx.consTable["0"]
	if !found {
		next, okAddr := l.pCtx.vm.constIntAddressMgr.GetNext()
		if !okAddr {
			l.pCtx.SemanticError("Out of memory")
		}
		zero := NewRType(Int)
		zero.FirstDim = 0
		zero.SecondDim = 0
		zero.Address = next
		l.pCtx.consTable["0"] = zero
		l.pCtx.addrTable[zero.Address] = "0"
		entry = zero
	}
	l.pCtx.vm.PushQuad(NewQuad(Assign, idxAddr, -1, entry.Address))
	l.pCtx.POPush(rType)
}

func (l *bistatListener) ExitForHeader(ctx *parser.ForHeaderContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	varName := ctx.ID().GetText()
	elem, found := l.pCtx.GetRTypeFromVarName(varName)
	if !found {
		l.pCtx.SemanticError("Variable " + varName + " not found")
		return
	}
	o := l.pCtx.POTop()
	l.pCtx.POPop()
	vc := l.pCtx.POTop()
	if o.FirstDim <= 0 || o.SecondDim > 0 {
		l.pCtx.SemanticError("For loops can only be used in arrays")
		return
	}
	if o.PType != vc.PType {
		l.pCtx.SemanticError("Control variable must be the same type as expression variable in for loop")
		return
	}
	lenAddr := l.pCtx.ConstIntUpsert(o.FirstDim)
	boolAddr, _ := l.pCtx.vm.tempBoolAddressMgr.GetNext()
	l.pCtx.JumpsPush(len(l.pCtx.vm.Quads()))
	l.pCtx.vm.PushQuad(NewQuad(Lt, vc.Address, lenAddr, boolAddr))
	currQuad := len(l.pCtx.vm.quads)
	l.pCtx.JumpsPush(currQuad)
	l.pCtx.vm.PushQuad(NewQuad(GotoF, boolAddr, -1, -1))
	addrMgr := l.pCtx.GetCorrespondingRefAddressManager(o.Address)
	refAddr, _ := addrMgr.GetNext()
	l.pCtx.vm.PushQuad(NewQuad(RefSum, vc.Address, l.pCtx.IgnoreIfRef(o.Address), refAddr))
	l.pCtx.vm.PushQuad(NewQuad(Assign, elem.Address, -1, refAddr))
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
		one.FirstDim = 0
		one.SecondDim = 0
		one.Address = next
		l.pCtx.consTable["1"] = one
		l.pCtx.addrTable[one.Address] = "1"
		entry = one
	}
	l.pCtx.vm.PushQuad(NewQuad(Sum, vc.Address, entry.Address, vc.Address))
	l.pCtx.POPop()
	end := l.pCtx.JumpsTop()
	l.pCtx.JumpsPop()
	ret := l.pCtx.JumpsTop()
	l.pCtx.JumpsPop()
	l.pCtx.vm.PushQuad(NewQuad(Goto, -1, -1, ret))
	l.pCtx.vm.quads[end].Destination = len(l.pCtx.vm.Quads())
}
