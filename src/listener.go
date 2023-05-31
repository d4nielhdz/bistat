package src

// todo: add fake bottom for array access and function calls
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

func NewBistatListener() *bistatListener {
	l := new(bistatListener)
	l.pCtx = NewPCtx()
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

	// todo: catch type errors
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
	// fmt.Println("Exit assignment")
	if len(l.pCtx.semanticErrors) > 0 {
		return
	}
	varName := ctx.ID().GetText()
	lRType, found := l.pCtx.GetRTypeFromVarName(varName)
	if !found {
		l.pCtx.SemanticError("Variable " + varName + " not defined")
		return
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
				if cols != lRType.FirstDim {
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
		rType, _ := l.pCtx.GetRTypeFromVarName(varName)
		startAddr := rType.Address
		i := expectedSize
		isLocal := l.pCtx.IsLocalVar(varName)
		addrMgr := l.pCtx.vm.globalRefAddressMgr
		if isLocal {
			addrMgr = l.pCtx.vm.localRefAddressMgr
		}
		for i > 0 {
			fmt.Println(i)
			elem := l.pCtx.POTop()
			l.pCtx.POPop()
			refAddr, _ := addrMgr.GetNext()
			// todo: store i-1 as const
			l.pCtx.vm.PushQuad(NewQuad(RefSum, l.pCtx.ConstIntUpsert(startAddr), l.pCtx.ConstIntUpsert(i-1), refAddr))
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
		quad := NewQuad(Assign, lRType.Address, -1, rRType.Address)
		l.pCtx.vm.PushQuad(quad)
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

func (l *bistatListener) ExitProgram(ctx *parser.ProgramContext) {
	l.pCtx.vm.PushQuad(NewQuad(End, -1, -1, -1))
	// l.pCtx.PrintAddrTable()
	l.pCtx.PrintQuads()
	l.pCtx.PrintErrors()
	if len(l.pCtx.semanticErrors) == 0 {
		RegisterTypes()
		file, err := os.Create("./obj.gob")
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

func RegisterTypes() {
	gob.Register(ObjCode{})
	gob.Register(FuncData{})
	gob.Register(RType{})
	gob.Register(Quad{})
}
