package src

import (
	parser "bistat/parser"
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
}

// func (l *bistatListener) ExitListAssign(ctx *parser.ListAssignContext) {
// 	if len(l.pCtx.semanticErrors) > 0 {
// 		return
// 	}
// 	varName := ctx.ID().GetText()
// 	rType, found := l.pCtx.GetRTypeFromVarName(varName)
// 	if !found {
// 		l.pCtx.SemanticError("Variable " + varName + " not defined")
// 		return
// 	}
// 	if rType.FirstDim <= 0 {
// 		l.pCtx.SemanticError("Variable " + varName + " is not an array")
// 		return
// 	}

// 	addrMgr := l.pCtx.GetCorrespondingRefAddressManager(rType.Address)
// 	iAddrMgr := l.pCtx.vm.tempIntAddressMgr

// 	if ctx.ListAssignment() != nil {
// 		expectedSize := 0
// 		if rType.SecondDim <= 0 {
// 			l.pCtx.SemanticError("Variable " + varName + " is not a matrix")
// 			return
// 		}
// 		if rType.SecondDim != len(ctx.ListAssignment().AllExpression()) {
// 			l.pCtx.SemanticError("Incorrect number of elements in assignment to " + varName + " expected " + strconv.Itoa(rType.SecondDim) + ", got " + strconv.Itoa(len(ctx.ListAssignment().AllExpression())))
// 			return
// 		}
// 		expectedSize = rType.SecondDim
// 		i := expectedSize
// 		idx := l.pCtx.pO[len(l.pCtx.pO)-expectedSize-1]
// 		iAddr, _ := addrMgr.GetNext()
// 		l.pCtx.vm.PushQuad(NewQuad(Verify, idx.Address, rType.Address, rType.SecondDim))
// 		l.pCtx.vm.PushQuad(NewQuad(RefMul, l.pCtx.ConstIntUpsert(rType.SecondDim), idx.Address, iAddr))
// 		secondAddr, _ := addrMgr.GetNext()
// 		l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.IgnoreIfRef(rType.Address), l.pCtx.IgnoreIfRef(iAddr), secondAddr))

// 		startAddr := secondAddr
// 		for i > 0 {
// 			elem := l.pCtx.POTop()
// 			l.pCtx.POPop()
// 			if rType.PType != elem.PType {
// 				// todo: cast where appropriate
// 				l.pCtx.SemanticError("Cannot assign to " + varName + " because of type mismatch: " + PTypeToString(elem.PType) + " != " + PTypeToString(rType.PType))
// 				return
// 			}
// 			refAddr, _ := addrMgr.GetNext()
// 			l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.IgnoreIfRef(startAddr), l.pCtx.ConstIntUpsert(i-1), refAddr))
// 			l.pCtx.vm.PushQuad(NewQuad(Assign, refAddr, -1, elem.Address))
// 			i = i - 1
// 		}
// 		l.pCtx.POPop()
// 	} else {
// 		levels := len(ctx.AllExpression())

// 		if levels == 2 {
// 			val := l.pCtx.POTop()
// 			l.pCtx.POPop()
// 			idx := l.pCtx.POTop()
// 			l.pCtx.POPop()
// 			if rType.PType != val.PType {
// 				// todo: cast where appropriate
// 				l.pCtx.SemanticError("Cannot assign to " + varName + " because of type mismatch: " + PTypeToString(val.PType) + " != " + PTypeToString(rType.PType))
// 				return
// 			}

// 			if val.FirstDim > 0 {
// 				if val.FirstDim != rType.SecondDim {
// 					l.pCtx.SemanticError("Dimension mismatch when assigning to " + varName)
// 				}
// 				i := val.FirstDim
// 				startAddr := val.Address
// 				l.pCtx.vm.PushQuad(NewQuad(Verify, idx.Address, rType.Address, rType.FirstDim))
// 				lStartAddr, _ := iAddrMgr.GetNext()
// 				l.pCtx.vm.PushQuad(NewQuad(RefMul, l.pCtx.ConstIntUpsert(rType.SecondDim), idx.Address, lStartAddr))
// 				secLStartAddr, _ := addrMgr.GetNext()
// 				l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.IgnoreIfRef(rType.Address), lStartAddr, secLStartAddr))
// 				rAddrMgr := l.pCtx.GetCorrespondingRefAddressManager(val.Address)

// 				for i > 0 {
// 					lAddr, _ := addrMgr.GetNext()
// 					rAddr, _ := rAddrMgr.GetNext()
// 					l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.IgnoreIfRef(startAddr), l.pCtx.ConstIntUpsert(i-1), rAddr))
// 					l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.IgnoreIfRef(secLStartAddr), l.pCtx.ConstIntUpsert(i-1), lAddr))
// 					l.pCtx.vm.PushQuad(NewQuad(Assign, lAddr, -1, rAddr))
// 					i = i - 1
// 				}
// 			} else {
// 				l.pCtx.vm.PushQuad(NewQuad(Verify, idx.Address, rType.Address, rType.FirstDim))
// 				refAddr, _ := addrMgr.GetNext()
// 				l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.IgnoreIfRef(rType.Address), idx.Address, refAddr))
// 				l.pCtx.vm.PushQuad(NewQuad(Assign, refAddr, -1, val.Address))
// 			}
// 		} else if levels == 3 {
// 			if rType.SecondDim == 0 {
// 				l.pCtx.SemanticError("Variable " + varName + " isn't a matrix")
// 				return
// 			}
// 			val := l.pCtx.POTop()
// 			l.pCtx.POPop()

// 			secondIdx := l.pCtx.POTop()
// 			l.pCtx.POPop()
// 			firstIdx := l.pCtx.POTop()
// 			l.pCtx.POPop()
// 			l.pCtx.vm.PushQuad(NewQuad(Verify, firstIdx.Address, rType.Address, rType.FirstDim))
// 			l.pCtx.vm.PushQuad(NewQuad(Verify, secondIdx.Address, rType.Address, rType.SecondDim))

// 			addr, _ := iAddrMgr.GetNext()
// 			l.pCtx.vm.PushQuad(NewQuad(RefMul, l.pCtx.ConstIntUpsert(rType.SecondDim), firstIdx.Address, addr))
// 			secondAddr, _ := iAddrMgr.GetNext()
// 			l.pCtx.vm.PushQuad(NewQuad(RefSum, addr, secondIdx.Address, secondAddr))
// 			refAddr, _ := addrMgr.GetNext()
// 			l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.IgnoreIfRef(rType.Address), secondAddr, refAddr))
// 			l.pCtx.vm.PushQuad(NewQuad(Assign, refAddr, -1, val.Address))
// 		}
// 	}
// }

func (l *bistatListener) ExitPrint(ctx *parser.PrintContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}
	if !l.pCtx.POperIsEmpty() && l.pCtx.POperTop() != int(Other) {
		l.pCtx.SemanticError("Cannot use 'print' inside an expression")
		return
	}
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
			addrMgr := l.pCtx.GetCorrespondingRefAddressManager(elem.Address)
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

// func (l *bistatListener) EnterListAccess(ctx *parser.ListAccessContext) {
// 	if len(l.pCtx.semanticErrors) > 0 {
// 		return
// 	}
// 	l.pCtx.POperPush(int(Other))
// }

// func (l *bistatListener) ExitListAccess(ctx *parser.ListAccessContext) {
// 	if len(l.pCtx.semanticErrors) > 0 {
// 		return
// 	}
// 	l.pCtx.POperPop()

// 	varName := ctx.ID().GetText()
// 	arr, found := l.pCtx.GetRTypeFromVarName(varName)
// 	if !found {
// 		l.pCtx.SemanticError("Variable " + varName + " was not declared")
// 		return
// 	}
// 	if arr.FirstDim == 0 {
// 		l.pCtx.SemanticError("Variable " + varName + " isn't a matrix or array")
// 		return
// 	}
// 	levels := len(ctx.AllExpression())

// 	if levels == 1 {
// 		idx := l.pCtx.POTop()
// 		l.pCtx.POPop()
// 		l.pCtx.vm.PushQuad(NewQuad(Verify, idx.Address, arr.Address, arr.FirstDim))
// 		addrMgr := l.pCtx.GetCorrespondingRefAddressManager(arr.Address)
// 		iAddrMgr := l.pCtx.vm.tempIntAddressMgr

// 		refAddr, _ := addrMgr.GetNext()
// 		indexed := NewRType(arr.PType)
// 		if arr.SecondDim != 0 {
// 			// ref is array
// 			iAddr, _ := iAddrMgr.GetNext()
// 			l.pCtx.vm.PushQuad(NewQuad(RefMul, l.pCtx.ConstIntUpsert(arr.SecondDim), idx.Address, iAddr))
// 			l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.IgnoreIfRef(arr.Address), iAddr, refAddr))

// 			indexed.FirstDim = arr.SecondDim
// 		} else {
// 			// ref is scalar
// 			l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.ConstIntUpsert(arr.Address), idx.Address, refAddr))
// 		}
// 		indexed.Address = refAddr

// 		l.pCtx.POPush(indexed)
// 	} else {
// 		if arr.SecondDim <= 0 {
// 			l.pCtx.SemanticError("Variable " + varName + " isn't a matrix")
// 			return
// 		}
// 		secondIdx := l.pCtx.POTop()
// 		l.pCtx.POPop()
// 		firstIdx := l.pCtx.POTop()
// 		l.pCtx.POPop()
// 		l.pCtx.vm.PushQuad(NewQuad(Verify, firstIdx.Address, arr.Address, arr.FirstDim))
// 		l.pCtx.vm.PushQuad(NewQuad(Verify, secondIdx.Address, arr.Address, arr.SecondDim))

// 		iAddrMgr := l.pCtx.vm.tempIntAddressMgr

// 		addrMgr := l.pCtx.GetCorrespondingRefAddressManager(arr.Address)
// 		addr, _ := iAddrMgr.GetNext()
// 		l.pCtx.vm.PushQuad(NewQuad(RefMul, l.pCtx.ConstIntUpsert(arr.SecondDim), firstIdx.Address, addr))
// 		secondAddr, _ := iAddrMgr.GetNext()
// 		l.pCtx.vm.PushQuad(NewQuad(RefSum, addr, secondIdx.Address, secondAddr))
// 		refAddr, _ := addrMgr.GetNext()
// 		l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.IgnoreIfRef(arr.Address), secondAddr, refAddr))

// 		indexed := NewRType(arr.PType)
// 		indexed.Address = refAddr
// 		l.pCtx.POPush(indexed)
// 	}
// }

func (l *bistatListener) ExitInputRead(ctx *parser.InputReadContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	if !l.pCtx.POperIsEmpty() && l.pCtx.POperTop() != int(Other) {
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
			addrMgr := l.pCtx.GetCorrespondingRefAddressManager(o.Address)
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
	l.pCtx.GenerateQuadsForNonAggregateFunction(Cos, Float)
}

func (l *bistatListener) ExitSin(ctx *parser.SinContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}
	l.pCtx.GenerateQuadsForNonAggregateFunction(Sin, Float)
}

func (l *bistatListener) ExitTan(ctx *parser.TanContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}
	l.pCtx.GenerateQuadsForNonAggregateFunction(Tan, Float)
}

func (l *bistatListener) ExitSqrt(ctx *parser.SqrtContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}
	l.pCtx.GenerateQuadsForNonAggregateFunction(Sqrt, l.pCtx.POTop().PType)
}

func (l *bistatListener) ExitFloor(ctx *parser.FloorContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}
	l.pCtx.GenerateQuadsForNonAggregateFunction(Floor, Int)
}

func (l *bistatListener) ExitCeil(ctx *parser.CeilContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}
	l.pCtx.GenerateQuadsForNonAggregateFunction(Ceil, Int)
}

func (l *bistatListener) ExitAbs(ctx *parser.AbsContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}
	l.pCtx.GenerateQuadsForNonAggregateFunction(Abs, l.pCtx.POTop().PType)
}

func (l *bistatListener) ExitNot(ctx *parser.NotContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}
	l.pCtx.GenerateQuadsForNonAggregateFunction(Not, Bool)
}

func (l *bistatListener) ExitSum(ctx *parser.SumContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	l.pCtx.GenerateQuadsForAggregateFunction(ListSum, l.pCtx.POTop().PType)
}

func (l *bistatListener) ExitMin(ctx *parser.MinContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	l.pCtx.GenerateQuadsForAggregateFunction(Min, l.pCtx.POTop().PType)
}

func (l *bistatListener) ExitMax(ctx *parser.MaxContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	l.pCtx.GenerateQuadsForAggregateFunction(Max, l.pCtx.POTop().PType)
}

func (l *bistatListener) ExitProd(ctx *parser.ProdContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	l.pCtx.GenerateQuadsForAggregateFunction(Prod, l.pCtx.POTop().PType)
}

func (l *bistatListener) ExitAvg(ctx *parser.AvgContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	l.pCtx.GenerateQuadsForAggregateFunction(Avg, Float)
}

func (l *bistatListener) ExitSMode(ctx *parser.SModeContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	l.pCtx.GenerateQuadsForAggregateFunction(SMode, l.pCtx.POTop().PType)
}

func (l *bistatListener) ExitMedian(ctx *parser.MedianContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	l.pCtx.GenerateQuadsForAggregateFunction(Median, Float)
}

func (l *bistatListener) ExitPlot(ctx *parser.PlotContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	if len(l.pCtx.semanticErrors) > 0 {
		return
	}
	o := l.pCtx.POTop()
	l.pCtx.POPop()
	if o.PType != Float && o.PType != Int {
		l.pCtx.SemanticError(OpToString(Plot) + " can only be called with lists of type int or float")
		return
	}
	if o.FirstDim <= 0 {
		l.pCtx.SemanticError(OpToString(Plot) + " can only be called on arrays or matrices")
		return
	}
	quad := NewQuad(Plot, o.Address, o.FirstDim, o.SecondDim)
	l.pCtx.vm.PushQuad(quad)
}
