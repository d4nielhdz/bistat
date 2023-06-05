package src

import (
	parser "bistat/parser"
	"encoding/gob"
	"fmt"
	"os"
	"strconv"
)

type bistatListener struct {
	*parser.BaseBistatListener
	pCtx PCtx
}

func NewBistatListener(filename string, printQuads bool) *bistatListener {
	l := new(bistatListener)
	l.pCtx = NewPCtx(printQuads, filename)
	return l
}

func (l *bistatListener) EnterProgram(ctx *parser.ProgramContext) {
	l.pCtx.AddScope("main")
	l.pCtx.AddFunction("main", NewFuncData(Void))
	fData := l.pCtx.funcDir["main"]
	fData.Idx = 0
	l.pCtx.funcDir["main"] = fData
	l.pCtx.vm.PushQuad(NewQuad(Goto, -1, -1, 0))
}

func (l *bistatListener) EnterVarDeclaration(ctx *parser.VarDeclarationContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}

	currScope := l.pCtx.GetCurrentScope()
	vt := ctx.Var_type()
	pType := PTypeFromString(vt.TYPE_PRIMITIVE().GetText())
	if pType == Void {
		l.pCtx.SemanticError("Can't declare void variables")
		return
	}
	addrMgr := l.pCtx.GetCorrespondingAddressManager(pType)
	resolved := l.pCtx.ResolveType(vt, addrMgr)
	varName := ctx.ID().GetText()
	_, exists := l.pCtx.GetVarInScope(currScope, varName)
	if exists {
		l.pCtx.SemanticError("Variable " + varName + " already declared")
		return
	} else {
		l.pCtx.AddVarToScope(currScope, varName, resolved)
		l.pCtx.AddToAddrTable(resolved.Address, ctx.ID().GetText())
	}
}
func (l *bistatListener) EnterAssignment(ctx *parser.AssignmentContext) {
	// fmt.Println("Enter assignment")

}
func (l *bistatListener) ExitAssignment(ctx *parser.AssignmentContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}
	varName := ""
	lRType := NewRType(Bool)
	if ctx.ID() != nil {
		varName = ctx.ID().GetText()
		lType, found := l.pCtx.GetRTypeFromVarName(varName)
		if !found {
			l.pCtx.SemanticError("Variable " + varName + " not defined")
			return
		}
		lRType = lType
	} else if ctx.Indexing() != nil {
		varName = ctx.Indexing().ID().GetText()
		rightSize := 1
		if ctx.ListAssignment() != nil {
			rightSize = len(ctx.ListAssignment().AllExpression())
		} else if ctx.MatrixAssignment() != nil {
			l.pCtx.SemanticError("Cannot assign a matrix with an indexing expression on the left side")
			return
		}
		lRType = l.pCtx.pO[len(l.pCtx.pO)-rightSize-1]
	}

	if ctx.ListAssignment() != nil || ctx.MatrixAssignment() != nil {
		expectedSize := 0

		if ctx.MatrixAssignment() != nil {
			if lRType.SecondDim <= 0 {
				l.pCtx.SemanticError("Variable " + varName + " is not a matrix")
				return
			}
			expectedSize = lRType.FirstDim * lRType.SecondDim
			rows := len(ctx.MatrixAssignment().AllListAssignment())
			if lRType.FirstDim != rows {
				l.pCtx.SemanticError("Incorrect number of rows in assignment to " + varName + " expected " + strconv.Itoa(lRType.FirstDim) + ", got " + strconv.Itoa(rows))
				return
			}
			for i, row := range ctx.MatrixAssignment().AllListAssignment() {
				cols := len(row.AllExpression())
				if cols != lRType.SecondDim {
					l.pCtx.SemanticError("Incorrect number of columns in assignment to " + varName + " in row " + strconv.Itoa(i) + "; expected " + strconv.Itoa(lRType.SecondDim) + ", got " + strconv.Itoa(cols))
					return
				}
			}
		} else {
			if lRType.FirstDim <= 0 {
				l.pCtx.SemanticError("Variable " + varName + " is not an array")
				return
			}
			if lRType.FirstDim != len(ctx.ListAssignment().AllExpression()) {
				l.pCtx.SemanticError("Incorrect number of elements in assignment to " + varName + " expected " + strconv.Itoa(lRType.FirstDim) + ", got " + strconv.Itoa(len(ctx.ListAssignment().AllExpression())))
				return
			}
			expectedSize = lRType.FirstDim
		}
		startAddr := lRType.Address
		i := expectedSize
		addrMgr := l.pCtx.GetCorrespondingRefAddressManager(lRType.Address)
		for i > 0 {
			elem := l.pCtx.POTop()
			l.pCtx.POPop()
			refAddr, _ := addrMgr.GetNext()
			l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.IgnoreIfRef(startAddr), l.pCtx.ConstIntUpsert(i-1), refAddr))
			l.pCtx.vm.PushQuad(NewQuad(Assign, refAddr, -1, elem.Address))
			i = i - 1
		}
		// if ctx.Indexing() != nil {
		// 	l.pCtx.POPop()
		// }
	} else {
		rRType := l.pCtx.POTop()
		l.pCtx.POPop()
		if lRType.PType != rRType.PType {
			// todo: cast where appropriate
			l.pCtx.SemanticError("Cannot assign to " + varName + " because of type mismatch: " + PTypeToString(rRType.PType) + " != " + PTypeToString(lRType.PType))
			return
		}
		if rRType.FirstDim > 0 {
			expectedSize := lRType.FirstDim
			if lRType.FirstDim != rRType.FirstDim {
				l.pCtx.SemanticError("Dimension mismatch when assigning to " + varName + ": expected " + strconv.Itoa(lRType.FirstDim) + " got " + strconv.Itoa(rRType.FirstDim))
				return
			}
			if lRType.SecondDim > 0 {
				expectedSize *= rRType.SecondDim
			}
			if lRType.SecondDim != rRType.SecondDim {
				l.pCtx.SemanticError("Dimension mismatch when assigning to " + varName + ": expected " + strconv.Itoa(lRType.SecondDim) + " got " + strconv.Itoa(rRType.SecondDim))
				return
			}

			addrMgr := l.pCtx.GetCorrespondingRefAddressManager(lRType.Address)

			i := 0
			startAddr := rRType.Address
			lStartAddr := lRType.Address
			rAddrMgr := l.pCtx.GetCorrespondingRefAddressManager(rRType.Address)

			for i != expectedSize {
				lAddr, _ := addrMgr.GetNext()
				rAddr, _ := rAddrMgr.GetNext()
				l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.IgnoreIfRef(startAddr), l.pCtx.ConstIntUpsert(i), rAddr))
				l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.IgnoreIfRef(lStartAddr), l.pCtx.ConstIntUpsert(i), lAddr))
				l.pCtx.vm.PushQuad(NewQuad(Assign, lAddr, -1, rAddr))
				i++
			}
		} else {
			quad := NewQuad(Assign, lRType.Address, -1, rRType.Address)
			l.pCtx.vm.PushQuad(quad)
		}
	}
	if ctx.Indexing() != nil {
		l.pCtx.POPop()
	}
}

func (l *bistatListener) EnterIndexing(ctx *parser.IndexingContext) {
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}
	l.pCtx.POperPush(int(Other))
}

func (l *bistatListener) ExitIndexing(ctx *parser.IndexingContext) {
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
		if idx.PType != Int {
			l.pCtx.SemanticError("Expression in indexing must be of type int")
			return
		}
		l.pCtx.POPop()
		l.pCtx.vm.PushQuad(NewQuad(Verify, idx.Address, arr.Address, arr.FirstDim))
		addrMgr := l.pCtx.GetCorrespondingRefAddressManager(arr.Address)

		indexed := NewRType(arr.PType)
		refAddr, _ := addrMgr.GetNext()
		if arr.SecondDim != 0 {
			l.pCtx.vm.PushQuad(NewQuad(RefMul, l.pCtx.ConstIntUpsert(arr.SecondDim), idx.Address, refAddr))
			endAddr, _ := addrMgr.GetNext()
			l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.IgnoreIfRef(arr.Address), refAddr, endAddr))

			indexed.FirstDim = arr.SecondDim
			indexed.Address = endAddr
		} else {
			l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.IgnoreIfRef(arr.Address), idx.Address, refAddr))
			indexed.Address = refAddr
		}

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

		if firstIdx.PType != Int || secondIdx.PType != Int {
			l.pCtx.SemanticError("Expressions in indexing must be of type int")
			return
		}

		l.pCtx.vm.PushQuad(NewQuad(Verify, firstIdx.Address, arr.Address, arr.FirstDim))
		l.pCtx.vm.PushQuad(NewQuad(Verify, secondIdx.Address, arr.Address, arr.SecondDim))

		addrMgr := l.pCtx.GetCorrespondingRefAddressManager(arr.Address)
		addr, _ := addrMgr.GetNext()
		l.pCtx.vm.PushQuad(NewQuad(RefMul, l.pCtx.ConstIntUpsert(arr.SecondDim), firstIdx.Address, addr))
		secondAddr, _ := addrMgr.GetNext()
		l.pCtx.vm.PushQuad(NewQuad(RefSum, addr, secondIdx.Address, secondAddr))
		refAddr, _ := addrMgr.GetNext()
		l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.IgnoreIfRef(arr.Address), secondAddr, refAddr))

		indexed := NewRType(arr.PType)
		indexed.Address = refAddr
		l.pCtx.POPush(indexed)
	}
}

func (l *bistatListener) EnterMain(ctx *parser.MainContext) {
	// fmt.Println("Enter main")
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}
	fData := l.pCtx.funcDir["main"]
	fData.FuncStart = len(l.pCtx.vm.Quads())
	l.pCtx.funcDir["main"] = fData
	l.pCtx.vm.quads[0].Destination = len(l.pCtx.vm.Quads())
}

/*
Loasing all data needed for execution
*/
func (l *bistatListener) ExitProgram(ctx *parser.ProgramContext) {
	l.pCtx.vm.PushQuad(NewQuad(End, -1, -1, -1))
	// l.pCtx.PrintAddrTable()
	if l.pCtx.printQuads {
		l.pCtx.PrintQuads()
	}
	l.pCtx.PrintErrors()
	if len(l.pCtx.semanticErrors) == 0 {
		RegisterTypes()
		file, err := os.Create("./" + l.pCtx.Filename + ".gob")
		objCode := NewObjCode(l.pCtx.funcDir, l.pCtx.consTable, l.pCtx.vm.quads, l.pCtx.functions)
		objCode.BoolSize = l.pCtx.vm.globalBoolAddressMgr.GetSize()
		objCode.StringSize = l.pCtx.vm.globalStringAddressMgr.GetSize()
		objCode.FloatSize = l.pCtx.vm.globalFloatAddressMgr.GetSize()
		objCode.IntSize = l.pCtx.vm.globalIntAddressMgr.GetSize()
		objCode.ConstBoolSize = l.pCtx.vm.constBoolAddressMgr.GetSize()
		objCode.ConstStringSize = l.pCtx.vm.constStringAddressMgr.GetSize()
		objCode.ConstFloatSize = l.pCtx.vm.constFloatAddressMgr.GetSize()
		objCode.ConstIntSize = l.pCtx.vm.constIntAddressMgr.GetSize()
		objCode.TempBoolSize = l.pCtx.vm.tempBoolAddressMgr.GetSize()
		objCode.TempStringSize = l.pCtx.vm.tempStringAddressMgr.GetSize()
		objCode.TempFloatSize = l.pCtx.vm.tempFloatAddressMgr.GetSize()
		objCode.TempIntSize = l.pCtx.vm.tempIntAddressMgr.GetSize()
		objCode.GlobalRefSize = l.pCtx.vm.globalRefAddressMgr.GetSize()

		if err == nil {
			encoder := gob.NewEncoder(file)
			encoder.Encode(objCode)
		} else {
			fmt.Println(err)
		}
		file.Close()
	}
}

// Required for writing .gob file
func RegisterTypes() {
	gob.Register(ObjCode{})
	gob.Register(FuncData{})
	gob.Register(RType{})
	gob.Register(Quad{})
}
