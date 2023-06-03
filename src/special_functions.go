package src

import (
	parser "bistat/parser"
	"strconv"
)

func (l *bistatListener) EnterSpecialFunction(ctx *parser.SpecialFunctionContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}
	l.pCtx.POperPush(int(Other))
}

func (l *bistatListener) ExitSpecialFunction(ctx *parser.SpecialFunctionContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}
	l.pCtx.POperPop()

	if ctx.Plot() != nil {
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
	iAddrMgr := l.pCtx.vm.tempIntAddressMgr

	if isLocal {
		addrMgr = l.pCtx.vm.localRefAddressMgr
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
		idx := l.pCtx.pO[len(l.pCtx.pO)-expectedSize-1]
		iAddr, _ := iAddrMgr.GetNext()
		l.pCtx.vm.PushQuad(NewQuad(Verify, idx.Address, rType.Address, rType.FirstDim))
		l.pCtx.vm.PushQuad(NewQuad(Multiplication, l.pCtx.ConstIntUpsert(rType.FirstDim), idx.Address, iAddr))
		secondAddr, _ := iAddrMgr.GetNext()
		l.pCtx.vm.PushQuad(NewQuad(Sum, l.pCtx.ConstIntUpsert(rType.Address), iAddr, secondAddr))

		startAddr := secondAddr
		for i > 0 {
			elem := l.pCtx.POTop()
			l.pCtx.POPop()
			if rType.PType != elem.PType {
				// todo: csecondAddrst where appropriate
				l.pCtx.SemanticError("Cannot assign to " + varName + " because of type mismatch: " + PTypeToString(elem.PType) + " != " + PTypeToString(rType.PType))
				return
			}
			refAddr, _ := addrMgr.GetNext()
			l.pCtx.vm.PushQuad(NewQuad(Sum, startAddr, l.pCtx.ConstIntUpsert(i-1), refAddr))
			l.pCtx.vm.PushQuad(NewQuad(Assign, refAddr, -1, elem.Address))
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
			if rType.PType != val.PType {
				// todo: csecondAddrst where appropriate
				l.pCtx.SemanticError("Cannot assign to " + varName + " because of type mismatch: " + PTypeToString(val.PType) + " != " + PTypeToString(rType.PType))
				return
			}

			if val.FirstDim > 0 {
				if val.FirstDim != rType.SecondDim {
					l.pCtx.SemanticError("Dimension mismatch when assigning to " + varName)
				}
				i := val.FirstDim
				startAddr := val.Address
				l.pCtx.vm.PushQuad(NewQuad(Verify, idx.Address, rType.Address, rType.FirstDim))
				lStartAddr, _ := iAddrMgr.GetNext()
				l.pCtx.vm.PushQuad(NewQuad(RefMul, l.pCtx.ConstIntUpsert(rType.FirstDim), idx.Address, lStartAddr))
				secLStartAddr, _ := addrMgr.GetNext()
				l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.ConstIntUpsert(rType.Address), lStartAddr, secLStartAddr))
				rAddrMgr := l.pCtx.vm.globalRefAddressMgr

				for i > 0 {
					lAddr, _ := addrMgr.GetNext()
					rAddr, _ := rAddrMgr.GetNext()
					l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.IgnoreIfRef(startAddr), l.pCtx.ConstIntUpsert(i-1), rAddr))
					l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.IgnoreIfRef(secLStartAddr), l.pCtx.ConstIntUpsert(i-1), lAddr))
					l.pCtx.vm.PushQuad(NewQuad(Assign, lAddr, -1, rAddr))
					i = i - 1
				}
			} else {
				l.pCtx.vm.PushQuad(NewQuad(Verify, idx.Address, rType.Address, rType.FirstDim))
				refAddr, _ := addrMgr.GetNext()
				l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.IgnoreIfRef(rType.Address), idx.Address, refAddr))
				l.pCtx.vm.PushQuad(NewQuad(Assign, refAddr, -1, val.Address))
			}
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
			l.pCtx.vm.PushQuad(NewQuad(RefMul, l.pCtx.ConstIntUpsert(rType.FirstDim), firstIdx.Address, addr))
			secondAddr, _ := iAddrMgr.GetNext()
			l.pCtx.vm.PushQuad(NewQuad(RefSum, addr, secondIdx.Address, secondAddr))
			refAddr, _ := addrMgr.GetNext()
			l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.IgnoreIfRef(rType.Address), secondAddr, refAddr))
			l.pCtx.vm.PushQuad(NewQuad(Assign, refAddr, -1, val.Address))
		}
	}
}

func (l *bistatListener) ExitPrint(ctx *parser.PrintContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	if !l.pCtx.POperIsEmpty() && l.pCtx.POperTop() != int(Other) {
		l.pCtx.SemanticError("Cannot use 'print' inside an expression")
		return
	}
	ctx.AllExpression()
	toPrint := make([]RType, 0)
	for range ctx.AllExpression() {
		o := l.pCtx.POTop()
		l.pCtx.POPop()
		toPrint = append([]RType{o}, toPrint...)
	}
	ws := l.pCtx.ConstStringUpsert(" ")

	for idx, elem := range toPrint {
		if idx != 0 {
			l.pCtx.vm.PushQuad(NewQuad(Print, ws, -1, -1))
		}
		if elem.FirstDim > 0 {
			lBrack := l.pCtx.ConstStringUpsert("[")
			rBrack := l.pCtx.ConstStringUpsert("]")
			size := elem.FirstDim
			i := 0
			addrMgr := l.pCtx.vm.globalRefAddressMgr
			if elem.SecondDim > 0 {
				size *= elem.SecondDim
				l.pCtx.vm.PushQuad(NewQuad(Print, lBrack, -1, -1))
			}
			for i != size {
				if i != 0 {
					l.pCtx.vm.PushQuad(NewQuad(Print, ws, -1, -1))
				}

				if (i == 0) || (elem.SecondDim > 0 && i%elem.SecondDim == 0) {
					l.pCtx.vm.PushQuad(NewQuad(Print, lBrack, -1, -1))
				}
				addr, _ := addrMgr.GetNext()
				l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.IgnoreIfRef(elem.Address), l.pCtx.ConstIntUpsert(i), addr))
				l.pCtx.vm.PushQuad(NewQuad(Print, addr, -1, -1))

				if i != 0 && elem.SecondDim > 0 && (i+1)%elem.SecondDim == 0 {
					l.pCtx.vm.PushQuad(NewQuad(Print, rBrack, -1, -1))
				}
				i = i + 1
			}
			l.pCtx.vm.PushQuad(NewQuad(Print, rBrack, -1, -1))
		} else {
			quad := NewQuad(Print, elem.Address, -1, -1)
			l.pCtx.vm.PushQuad(quad)
		}
	}
	l.pCtx.vm.PushQuad(NewQuad(PrintN, -1, -1, -1))
}

func (l *bistatListener) EnterListAccess(ctx *parser.ListAccessContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}
	l.pCtx.POperPush(int(Other))
}

func (l *bistatListener) ExitListAccess(ctx *parser.ListAccessContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}
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
		iAddrMgr := l.pCtx.vm.tempIntAddressMgr
		if isLocal {
			addrMgr = l.pCtx.vm.localRefAddressMgr
		}
		refAddr, _ := addrMgr.GetNext()
		indexed := NewRType(arr.PType)
		if arr.SecondDim != 0 {
			// ref is array
			iAddr, _ := iAddrMgr.GetNext()
			l.pCtx.vm.PushQuad(NewQuad(RefMul, l.pCtx.ConstIntUpsert(arr.FirstDim), idx.Address, iAddr))
			l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.ConstIntUpsert(arr.Address), iAddr, refAddr))

			indexed.FirstDim = arr.SecondDim
		} else {
			// ref is scalar
			l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.ConstIntUpsert(arr.Address), idx.Address, refAddr))
		}
		indexed.Address = refAddr

		// todo: figure out how/when to calculate end address
		l.pCtx.POPush(indexed)
	} else {
		if arr.SecondDim <= 0 {
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
		iAddrMgr := l.pCtx.vm.tempIntAddressMgr

		if isLocal {
			addrMgr = l.pCtx.vm.localRefAddressMgr
		}
		addr, _ := iAddrMgr.GetNext()
		l.pCtx.vm.PushQuad(NewQuad(RefMul, l.pCtx.ConstIntUpsert(arr.FirstDim), firstIdx.Address, addr))
		secondAddr, _ := iAddrMgr.GetNext()
		l.pCtx.vm.PushQuad(NewQuad(RefSum, addr, secondIdx.Address, secondAddr))
		refAddr, _ := addrMgr.GetNext()
		l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.ConstIntUpsert(arr.Address), secondAddr, refAddr))

		indexed := NewRType(arr.PType)
		indexed.Address = refAddr
		// todo: figure out how/when to calculate end address
		l.pCtx.POPush(indexed)
	}
}

func (l *bistatListener) ExitInputRead(ctx *parser.InputReadContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	if !l.pCtx.POperIsEmpty() {
		l.pCtx.SemanticError("Cannot use 'read' inside an expression")
		return
	}
	for range ctx.AllExpression() {
		o := l.pCtx.POTop()
		if o.FirstDim > 0 {
			expectedSize := o.FirstDim
			if o.SecondDim > 0 {
				expectedSize *= o.SecondDim
			}
			addrMgr := l.pCtx.vm.globalRefAddressMgr
			i := 0
			for i != expectedSize {
				refAddr, _ := addrMgr.GetNext()
				l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.IgnoreIfRef(o.Address), l.pCtx.ConstIntUpsert(i), refAddr))
				l.pCtx.vm.PushQuad(NewQuad(InputRead, refAddr, -1, -1))
				i++
			}
		} else {
			quad := NewQuad(InputRead, o.Address, -1, -1)
			l.pCtx.vm.PushQuad(quad)
		}
		l.pCtx.POPop()
	}
}

func (l *bistatListener) ExitCos(ctx *parser.CosContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	o := l.pCtx.POTop()
	l.pCtx.POPop()

	if o.PType != Int && o.PType != Float {
		l.pCtx.SemanticError("cos can only be called with a float or int expression")
		return
	}

	result := NewRType(Float)
	if o.FirstDim > 0 {
		size := o.FirstDim
		if o.SecondDim > 0 {
			size *= o.SecondDim
		}
		addrMgr := l.pCtx.vm.globalRefAddressMgr
		if l.pCtx.IsLocalAddr(o.Address) || l.pCtx.IsLocalRef(o.Address) {
			addrMgr = l.pCtx.vm.localRefAddressMgr
		}
		i := 0
		result.FirstDim = o.FirstDim
		result.SecondDim = o.SecondDim
		for i < size {
			refAddr, _ := addrMgr.GetNext()
			if i == 0 {
				result.Address = refAddr
			}
			l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.IgnoreIfRef(o.Address), l.pCtx.ConstIntUpsert(i), refAddr))
			l.pCtx.vm.PushQuad(NewQuad(Cos, refAddr, -1, refAddr))
			i++
		}
	} else {
		addr, _ := l.pCtx.vm.tempFloatAddressMgr.GetNext()
		result.Address = addr
		l.pCtx.vm.PushQuad(NewQuad(Cos, o.Address, -1, addr))
	}
	l.pCtx.POPush(result)
}

func (l *bistatListener) ExitSin(ctx *parser.SinContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	o := l.pCtx.POTop()
	l.pCtx.POPop()

	if o.PType != Int && o.PType != Float {
		l.pCtx.SemanticError("sin can only be called with a float or int expression")
		return
	}

	result := NewRType(Float)
	if o.FirstDim > 0 {
		size := o.FirstDim
		if o.SecondDim > 0 {
			size *= o.SecondDim
		}
		addrMgr := l.pCtx.vm.globalRefAddressMgr
		if l.pCtx.IsLocalAddr(o.Address) || l.pCtx.IsLocalRef(o.Address) {
			addrMgr = l.pCtx.vm.localRefAddressMgr
		}
		i := 0
		result.FirstDim = o.FirstDim
		result.SecondDim = o.SecondDim
		for i < size {
			refAddr, _ := addrMgr.GetNext()
			if i == 0 {
				result.Address = refAddr
			}
			l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.IgnoreIfRef(o.Address), l.pCtx.ConstIntUpsert(i), refAddr))
			l.pCtx.vm.PushQuad(NewQuad(Sin, refAddr, -1, refAddr))
			i++
		}
	} else {
		addr, _ := l.pCtx.vm.tempFloatAddressMgr.GetNext()
		result.Address = addr
		l.pCtx.vm.PushQuad(NewQuad(Sin, o.Address, -1, addr))
	}
	l.pCtx.POPush(result)
}

func (l *bistatListener) ExitTan(ctx *parser.TanContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	o := l.pCtx.POTop()
	l.pCtx.POPop()

	if o.PType != Int && o.PType != Float {
		l.pCtx.SemanticError("tan can only be called with a float or int expression")
		return
	}

	result := NewRType(Float)
	if o.FirstDim > 0 {
		size := o.FirstDim
		if o.SecondDim > 0 {
			size *= o.SecondDim
		}
		addrMgr := l.pCtx.vm.globalRefAddressMgr
		if l.pCtx.IsLocalAddr(o.Address) || l.pCtx.IsLocalRef(o.Address) {
			addrMgr = l.pCtx.vm.localRefAddressMgr
		}
		i := 0
		result.FirstDim = o.FirstDim
		result.SecondDim = o.SecondDim
		for i < size {
			refAddr, _ := addrMgr.GetNext()
			if i == 0 {
				result.Address = refAddr
			}
			l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.IgnoreIfRef(o.Address), l.pCtx.ConstIntUpsert(i), refAddr))
			l.pCtx.vm.PushQuad(NewQuad(Tan, refAddr, -1, refAddr))
			i++
		}
	} else {
		addr, _ := l.pCtx.vm.tempFloatAddressMgr.GetNext()
		result.Address = addr
		l.pCtx.vm.PushQuad(NewQuad(Tan, o.Address, -1, addr))
	}
	l.pCtx.POPush(result)
}

func (l *bistatListener) ExitSqrt(ctx *parser.SqrtContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	o := l.pCtx.POTop()
	l.pCtx.POPop()

	if o.PType != Int && o.PType != Float {
		l.pCtx.SemanticError("sqrt can only be called with a float or int expression")
		return
	}

	result := NewRType(o.PType)
	if o.FirstDim > 0 {
		size := o.FirstDim
		if o.SecondDim > 0 {
			size *= o.SecondDim
		}
		addrMgr := l.pCtx.vm.globalRefAddressMgr
		if l.pCtx.IsLocalAddr(o.Address) || l.pCtx.IsLocalRef(o.Address) {
			addrMgr = l.pCtx.vm.localRefAddressMgr
		}
		i := 0
		result.FirstDim = o.FirstDim
		result.SecondDim = o.SecondDim
		for i < size {
			refAddr, _ := addrMgr.GetNext()
			if i == 0 {
				result.Address = refAddr
			}
			l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.IgnoreIfRef(o.Address), l.pCtx.ConstIntUpsert(i), refAddr))
			l.pCtx.vm.PushQuad(NewQuad(Sqrt, refAddr, -1, refAddr))
			i++
		}
	} else {
		addr, _ := l.pCtx.GetCorrespondingTempAddressManager(o.PType).GetNext()
		result.Address = addr
		l.pCtx.vm.PushQuad(NewQuad(Sqrt, o.Address, -1, addr))
	}
	l.pCtx.POPush(result)
}

func (l *bistatListener) ExitFloor(ctx *parser.FloorContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	o := l.pCtx.POTop()
	l.pCtx.POPop()

	if o.PType != Float {
		l.pCtx.SemanticError("floor can only be called with a float expression")
		return
	}

	result := NewRType(o.PType)
	if o.FirstDim > 0 {
		size := o.FirstDim
		if o.SecondDim > 0 {
			size *= o.SecondDim
		}
		addrMgr := l.pCtx.vm.globalRefAddressMgr
		if l.pCtx.IsLocalAddr(o.Address) || l.pCtx.IsLocalRef(o.Address) {
			addrMgr = l.pCtx.vm.localRefAddressMgr
		}
		i := 0
		result.FirstDim = o.FirstDim
		result.SecondDim = o.SecondDim
		for i < size {
			refAddr, _ := addrMgr.GetNext()
			if i == 0 {
				result.Address = refAddr
			}
			l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.IgnoreIfRef(o.Address), l.pCtx.ConstIntUpsert(i), refAddr))
			l.pCtx.vm.PushQuad(NewQuad(Floor, refAddr, -1, refAddr))
			i++
		}
	} else {
		addr, _ := l.pCtx.GetCorrespondingTempAddressManager(o.PType).GetNext()
		result.Address = addr
		l.pCtx.vm.PushQuad(NewQuad(Floor, o.Address, -1, addr))
	}
	l.pCtx.POPush(result)
}

func (l *bistatListener) ExitCeil(ctx *parser.CeilContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	o := l.pCtx.POTop()
	l.pCtx.POPop()

	if o.PType != Float {
		l.pCtx.SemanticError("ceil can only be called with a float expression")
		return
	}

	result := NewRType(o.PType)
	if o.FirstDim > 0 {
		size := o.FirstDim
		if o.SecondDim > 0 {
			size *= o.SecondDim
		}
		addrMgr := l.pCtx.vm.globalRefAddressMgr
		if l.pCtx.IsLocalAddr(o.Address) || l.pCtx.IsLocalRef(o.Address) {
			addrMgr = l.pCtx.vm.localRefAddressMgr
		}
		i := 0
		result.FirstDim = o.FirstDim
		result.SecondDim = o.SecondDim
		for i < size {
			refAddr, _ := addrMgr.GetNext()
			if i == 0 {
				result.Address = refAddr
			}
			l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.IgnoreIfRef(o.Address), l.pCtx.ConstIntUpsert(i), refAddr))
			l.pCtx.vm.PushQuad(NewQuad(Ceil, refAddr, -1, refAddr))
			i++
		}
	} else {
		addr, _ := l.pCtx.GetCorrespondingTempAddressManager(o.PType).GetNext()
		result.Address = addr
		l.pCtx.vm.PushQuad(NewQuad(Ceil, o.Address, -1, addr))
	}
	l.pCtx.POPush(result)
}

func (l *bistatListener) ExitAbs(ctx *parser.AbsContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	o := l.pCtx.POTop()
	l.pCtx.POPop()

	if o.PType != Float && o.PType != Int {
		l.pCtx.SemanticError("abs can only be called with a float or int expression")
		return
	}

	result := NewRType(o.PType)
	if o.FirstDim > 0 {
		size := o.FirstDim
		if o.SecondDim > 0 {
			size *= o.SecondDim
		}
		addrMgr := l.pCtx.vm.globalRefAddressMgr
		if l.pCtx.IsLocalAddr(o.Address) || l.pCtx.IsLocalRef(o.Address) {
			addrMgr = l.pCtx.vm.localRefAddressMgr
		}
		i := 0
		result.FirstDim = o.FirstDim
		result.SecondDim = o.SecondDim
		for i < size {
			refAddr, _ := addrMgr.GetNext()
			if i == 0 {
				result.Address = refAddr
			}
			l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.IgnoreIfRef(o.Address), l.pCtx.ConstIntUpsert(i), refAddr))
			l.pCtx.vm.PushQuad(NewQuad(Abs, refAddr, -1, refAddr))
			i++
		}
	} else {
		addr, _ := l.pCtx.GetCorrespondingTempAddressManager(o.PType).GetNext()
		result.Address = addr
		l.pCtx.vm.PushQuad(NewQuad(Abs, o.Address, -1, addr))
	}
	l.pCtx.POPush(result)
}

func (l *bistatListener) ExitNot(ctx *parser.NotContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	o := l.pCtx.POTop()
	l.pCtx.POPop()

	if o.PType != Bool {
		l.pCtx.SemanticError("not can only be called with a boolean expression")
		return
	}

	result := NewRType(o.PType)
	if o.FirstDim > 0 {
		size := o.FirstDim
		if o.SecondDim > 0 {
			size *= o.SecondDim
		}
		addrMgr := l.pCtx.vm.globalRefAddressMgr
		if l.pCtx.IsLocalAddr(o.Address) || l.pCtx.IsLocalRef(o.Address) {
			addrMgr = l.pCtx.vm.localRefAddressMgr
		}
		i := 0
		result.FirstDim = o.FirstDim
		result.SecondDim = o.SecondDim
		for i < size {
			refAddr, _ := addrMgr.GetNext()
			if i == 0 {
				result.Address = refAddr
			}
			l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.IgnoreIfRef(o.Address), l.pCtx.ConstIntUpsert(i), refAddr))
			l.pCtx.vm.PushQuad(NewQuad(Not, refAddr, -1, refAddr))
			i++
		}
	} else {
		addr, _ := l.pCtx.GetCorrespondingTempAddressManager(o.PType).GetNext()
		result.Address = addr
		l.pCtx.vm.PushQuad(NewQuad(Not, o.Address, -1, addr))
	}
	l.pCtx.POPush(result)
}
