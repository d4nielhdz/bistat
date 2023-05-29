package src

import (
	parser "bistat/parser"
	"fmt"
	"strconv"
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

func (l *bistatListener) ExitListAssign(ctx *parser.ListAssignContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}
	varName := ctx.ID().GetText()
	rType, found := l.pCtx.GetRTypeFromVarName(varName)
	if !found {
		l.pCtx.SemanticError("Variable " + varName + " not defined")
		return
	}
	if rType.FirstDim <= 0 {
		l.pCtx.SemanticError("Variable " + varName + " is not an array")
		return
	}

	isLocal := l.pCtx.IsLocalVar(varName)
	addrMgr := l.pCtx.vm.globalRefAddressMgr
	iAddrMgr := l.pCtx.vm.globalIntAddressMgr

	if isLocal {
		addrMgr = l.pCtx.vm.localRefAddressMgr
		iAddrMgr = l.pCtx.vm.localIntAddressMgr
	}

	if ctx.ListAssignment() != nil {
		expectedSize := 0
		if rType.SecondDim <= 0 {
			l.pCtx.SemanticError("Variable " + varName + " is not a matrix")
			return
		}
		if rType.SecondDim != len(ctx.ListAssignment().AllExpression()) {
			l.pCtx.SemanticError("Incorrect number of elements in assignment to " + varName + " expected " + strconv.Itoa(rType.SecondDim) + ", got " + strconv.Itoa(len(ctx.ListAssignment().AllExpression())))
			return
		}
		expectedSize = rType.SecondDim
		i := expectedSize
		fmt.Println(expectedSize)
		idx := l.pCtx.pO[len(l.pCtx.pO)-expectedSize-1]
		iAddr, _ := iAddrMgr.GetNext()
		l.pCtx.vm.PushQuad(NewQuad(Verify, rType.Address, idx.Address, rType.FirstDim))
		l.pCtx.vm.PushQuad(NewQuad(Multiplication, rType.Address, idx.Address, iAddr))
		startAddr := iAddr
		for i > 0 {
			elem := l.pCtx.POTop()
			l.pCtx.POPop()
			if rType.PType != elem.PType {
				// todo: cast where appropriate
				l.pCtx.SemanticError("Cannot assign to " + varName + " because of type mismatch: " + PTypeToString(elem.PType) + " != " + PTypeToString(rType.PType))
				return
			}
			addr, _ := addrMgr.GetNext()
			l.pCtx.vm.PushQuad(NewQuad(Sum, startAddr, i-1, addr))
			l.pCtx.vm.PushQuad(NewQuad(Assign, addr, -1, elem.Address))
			i = i - 1
		}
		l.pCtx.POPop()
	} else {
		levels := len(ctx.AllExpression())

		if levels == 2 {
			val := l.pCtx.POTop()
			l.pCtx.POPop()

			idx := l.pCtx.POTop()
			l.pCtx.POPop()
			l.pCtx.vm.PushQuad(NewQuad(Verify, idx.Address, rType.Address, rType.FirstDim))
			addr, _ := addrMgr.GetNext()
			l.pCtx.vm.PushQuad(NewQuad(Sum, rType.Address, idx.Address, addr))
			l.pCtx.vm.PushQuad(NewQuad(Assign, addr, -1, val.Address))
		} else if levels == 3 {
			if rType.SecondDim == 0 {
				l.pCtx.SemanticError("Variable " + varName + " isn't a matrix")
				return
			}
			val := l.pCtx.POTop()
			l.pCtx.POPop()

			secondIdx := l.pCtx.POTop()
			l.pCtx.POPop()
			firstIdx := l.pCtx.POTop()
			l.pCtx.POPop()
			l.pCtx.vm.PushQuad(NewQuad(Verify, firstIdx.Address, rType.Address, rType.FirstDim))
			l.pCtx.vm.PushQuad(NewQuad(Verify, secondIdx.Address, rType.Address, rType.SecondDim))

			addr, _ := iAddrMgr.GetNext()
			l.pCtx.vm.PushQuad(NewQuad(Multiplication, rType.Address, firstIdx.Address, addr))
			secondAddr, _ := addrMgr.GetNext()
			l.pCtx.vm.PushQuad(NewQuad(Sum, addr, secondIdx.Address, secondAddr))
			l.pCtx.vm.PushQuad(NewQuad(Assign, secondAddr, -1, val.Address))
		}
	}
}

func (l *bistatListener) EnterListAccess(ctx *parser.ListAccessContext) {
	l.pCtx.POperPush(int(Other))
}

func (l *bistatListener) ExitListAccess(ctx *parser.ListAccessContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}
	fmt.Println("Exit list access")
	l.pCtx.POperPop()

	varName := ctx.ID().GetText()
	arr, found := l.pCtx.GetRTypeFromVarName(varName)
	if !found {
		l.pCtx.SemanticError("Variable " + varName + " was not declared")
		return
	}
	if arr.FirstDim == 0 {
		l.pCtx.SemanticError("Variable " + varName + " isn't a matrix or array")
		return
	}
	levels := len(ctx.AllExpression())

	if levels == 1 {
		idx := l.pCtx.POTop()
		l.pCtx.POPop()
		l.pCtx.vm.PushQuad(NewQuad(Verify, idx.Address, arr.Address, arr.FirstDim))
		// addr is a ref
		isLocal := l.pCtx.IsLocalVar(varName)
		addrMgr := l.pCtx.vm.globalRefAddressMgr
		if isLocal {
			addrMgr = l.pCtx.vm.localRefAddressMgr
		}
		addr, _ := addrMgr.GetNext()
		indexed := NewRType(arr.PType)
		if arr.SecondDim != 0 {
			// ref is array
			l.pCtx.vm.PushQuad(NewQuad(Multiplication, arr.Address, idx.Address, addr))
			indexed.FirstDim = arr.SecondDim
		} else {
			// ref is scalar
			l.pCtx.vm.PushQuad(NewQuad(Sum, arr.Address, idx.Address, addr))
		}
		indexed.Address = addr

		// todo: figure out how/when to calculate end address
		l.pCtx.POPush(indexed)
		fmt.Println("pushed arr")
	} else {
		if arr.SecondDim == 0 {
			l.pCtx.SemanticError("Variable " + varName + " isn't a matrix")
			return
		}
		secondIdx := l.pCtx.POTop()
		l.pCtx.POPop()
		firstIdx := l.pCtx.POTop()
		l.pCtx.POPop()
		l.pCtx.vm.PushQuad(NewQuad(Verify, firstIdx.Address, arr.Address, arr.FirstDim))
		l.pCtx.vm.PushQuad(NewQuad(Verify, secondIdx.Address, arr.Address, arr.SecondDim))

		// addr is a ref
		isLocal := l.pCtx.IsLocalVar(varName)
		addrMgr := l.pCtx.vm.globalRefAddressMgr
		iAddrMgr := l.pCtx.vm.globalIntAddressMgr

		if isLocal {
			addrMgr = l.pCtx.vm.localRefAddressMgr
			iAddrMgr = l.pCtx.vm.localIntAddressMgr
		}
		addr, _ := iAddrMgr.GetNext()
		l.pCtx.vm.PushQuad(NewQuad(Multiplication, arr.Address, firstIdx.Address, addr))
		secondAddr, _ := addrMgr.GetNext()
		l.pCtx.vm.PushQuad(NewQuad(Sum, addr, secondIdx.Address, secondAddr))

		indexed := NewRType(arr.PType)
		indexed.Address = addr
		// todo: figure out how/when to calculate end address
		l.pCtx.POPush(indexed)
		fmt.Println("pushed matr")
	}
}
