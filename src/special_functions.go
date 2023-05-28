package src

import (
	parser "bistat/parser"
)

func (l *bistatListener) ExitSpecialFunction(ctx *parser.SpecialFunctionContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	if ctx.InputRead() != nil {
		if !l.pCtx.POperIsEmpty() {
			l.pCtx.SemanticError("Cannot use 'read' inside an expression")
			return
		}
		for range ctx.InputRead().AllID() {
			o := l.pCtx.POTop()
			quad := NewQuad(InputRead, o.Address, -1, -1)
			l.pCtx.vm.PushQuad(quad)
			l.pCtx.POPop()
		}
	} else if ctx.Print_() != nil {
		if !l.pCtx.POperIsEmpty() {
			l.pCtx.SemanticError("Cannot use 'print' inside an expression")
			return
		}
		ctx.Print_().AllExpression()
		for range ctx.Print_().AllExpression() {
			o := l.pCtx.POTop()
			quad := NewQuad(Print, o.Address, -1, -1)
			l.pCtx.vm.PushQuad(quad)
			l.pCtx.POPop()
		}
		l.pCtx.vm.PushQuad(NewQuad(PrintN, -1, -1, -1))
	} else if ctx.ListAdd() != nil {
		if !l.pCtx.POperIsEmpty() {
			l.pCtx.SemanticError("Cannot use 'listAdd' inside an arithmetic expression")
			return
		}
		o1 := l.pCtx.POTop()
		l.pCtx.POPop()
		o2 := l.pCtx.POTop()
		l.pCtx.POPop()
		quad := NewQuad(ListAdd, o2.Address, o1.Address, o2.Address)
		l.pCtx.vm.PushQuad(quad)
	} else if ctx.ListPop() != nil {
		if !l.pCtx.POperIsEmpty() {
			l.pCtx.SemanticError("Cannot use 'listPop' inside an arithmetic expression")
			return
		}
		o1 := l.pCtx.POTop()
		l.pCtx.POPop()
		quad := NewQuad(ListPop, o1.Address, -1, o1.Address)
		l.pCtx.vm.PushQuad(quad)
	} else if ctx.Length() != nil {
		// todo: verify array type
		o1 := l.pCtx.POTop()
		l.pCtx.POPop()
		addr, _ := l.pCtx.vm.tempIntAddressMgr.GetNext()
		rType := NewRType(Int)
		rType.Address = addr
		quad := NewQuad(Length, o1.Address, -1, addr)
		l.pCtx.vm.PushQuad(quad)
		l.pCtx.POPush(rType)
	} else if ctx.Range_() != nil {
		// todo: change when arrays work
		o1 := l.pCtx.POTop()
		l.pCtx.POPop()
		o2 := l.pCtx.POTop()
		l.pCtx.POPop()
		addr, _ := l.pCtx.vm.tempIntAddressMgr.GetNext()
		rType := NewRType(Int)
		rType.Address = addr
		quad := NewQuad(Length, o2.Address, o1.Address, addr)
		l.pCtx.vm.PushQuad(quad)
		l.pCtx.POPush(rType)
	} else if ctx.Plot() != nil {
		if !l.pCtx.POperIsEmpty() {
			l.pCtx.SemanticError("Cannot use 'plot' inside an arithmetic expression")
			return
		}
		o1 := l.pCtx.POTop()
		l.pCtx.POPop()
		quad := NewQuad(Plot, o1.Address, -1, -1)
		l.pCtx.vm.PushQuad(quad)
	} else if ctx.Sum() != nil {
		// todo: verify array type
		o1 := l.pCtx.POTop()
		l.pCtx.POPop()
		if o1.PType != Int && o1.PType != Float {
			l.pCtx.SemanticError("Cannot perform sum on " + PTypeToString(o1.PType))
		}
		addr, _ := l.pCtx.GetCorrespondingTempAddressManager(o1.PType).GetNext()
		rType := NewRType(o1.PType)
		rType.Address = addr
		quad := NewQuad(Sum, o1.Address, -1, addr)
		l.pCtx.vm.PushQuad(quad)
		l.pCtx.POPush(rType)
	} else if ctx.Min() != nil {
		// todo: verify array type
		o1 := l.pCtx.POTop()
		l.pCtx.POPop()
		if o1.PType != Int && o1.PType != Float {
			l.pCtx.SemanticError("Cannot perform min on " + PTypeToString(o1.PType))
		}
		addr, _ := l.pCtx.GetCorrespondingTempAddressManager(o1.PType).GetNext()
		rType := NewRType(o1.PType)
		rType.Address = addr
		quad := NewQuad(Min, o1.Address, -1, addr)
		l.pCtx.vm.PushQuad(quad)
		l.pCtx.POPush(rType)
	} else if ctx.Avg() != nil {
		// todo: verify array type
		o1 := l.pCtx.POTop()
		l.pCtx.POPop()
		if o1.PType != Int && o1.PType != Float {
			l.pCtx.SemanticError("Cannot perform avg on " + PTypeToString(o1.PType))
		}
		addr, _ := l.pCtx.GetCorrespondingTempAddressManager(o1.PType).GetNext()
		rType := NewRType(o1.PType)
		rType.Address = addr
		quad := NewQuad(Avg, o1.Address, -1, addr)
		l.pCtx.vm.PushQuad(quad)
		l.pCtx.POPush(rType)
	} else if ctx.Prod() != nil {
		// todo: verify array type
		o1 := l.pCtx.POTop()
		l.pCtx.POPop()
		if o1.PType != Int && o1.PType != Float {
			l.pCtx.SemanticError("Cannot perform prod on " + PTypeToString(o1.PType))
		}
		addr, _ := l.pCtx.GetCorrespondingTempAddressManager(o1.PType).GetNext()
		rType := NewRType(o1.PType)
		rType.Address = addr
		quad := NewQuad(Prod, o1.Address, -1, addr)
		l.pCtx.vm.PushQuad(quad)
		l.pCtx.POPush(rType)
	}
	// todo: handle missing functions
}

func HandleListAssign() {

}

func HandleListAccess() {

}
