package utils

import (
	parser "bistat/parser"
	"fmt"
	"regexp"
	"strconv"
)

type PCtx struct {
	scopes         []string
	funcDir        map[string]RType
	addrTable      map[int]string
	varTable       map[string]map[string]RType
	consTable      map[string]RType
	semanticErrors []string
	pO             []RType
	pOper          []int
	vm             VM
}

func NewPCtx() PCtx {
	return PCtx{
		scopes:         make([]string, 0),
		funcDir:        make(map[string]RType),
		varTable:       make(map[string]map[string]RType),
		consTable:      make(map[string]RType),
		semanticErrors: make([]string, 0),
		vm:             NewVM(),
		addrTable:      make(map[int]string),
		pO:             make([]RType, 0),
		pOper:          make([]int, 0),
	}
}

func (pCtx PCtx) GetCurrentScope() string {
	return pCtx.scopes[len(pCtx.scopes)-1]
}

func (pCtx *PCtx) PopCurrentScope() {
	scopes := pCtx.scopes
	popped := scopes[:len(scopes)-1]
	pCtx.scopes = popped
}

func (pCtx *PCtx) PrintAddrTable() {
	fmt.Println("-------")
	fmt.Println("printing adrr table")
	for addr, varName := range pCtx.addrTable {
		fmt.Println(varName)
		fmt.Println(addr)
		for k, _ := range pCtx.varTable {
			if rType, ok := pCtx.varTable[k][varName]; ok {
				rType.print()
				fmt.Println("_____")
			}
		}
	}
}

func (pCtx *PCtx) GetVarnameAtAddress(addr int) string {
	vn, ok := pCtx.addrTable[addr]
	if !ok {
		return strconv.Itoa(addr)
	}
	return vn
}

func (pCtx *PCtx) PrintErrors() {
	for _, err := range pCtx.semanticErrors {
		fmt.Println(err)
	}
}

func (pCtx *PCtx) GetRTypeFromVarConsContext(ctx *parser.VarConsContext) RType {
	if ctx.INT_CONS() != nil {
		entry, found := pCtx.consTable[ctx.INT_CONS().GetText()]
		if found {
			return entry
		}
		next, okAddr := pCtx.vm.constIntAddressMgr.GetNext()
		if !okAddr {
			pCtx.SemanticError("Out of memory")
		}
		rType := NewRType(Int)
		rType.firstDim = 0
		rType.secondDim = 0
		rType.address = next
		pCtx.consTable[ctx.INT_CONS().GetText()] = rType
		pCtx.addrTable[rType.address] = ctx.INT_CONS().GetText()
		return rType
	} else if ctx.FLOAT_CONS() != nil {
		entry, found := pCtx.consTable[ctx.FLOAT_CONS().GetText()]
		if found {
			return entry
		}
		next, okAddr := pCtx.vm.constFloatAddressMgr.GetNext()
		if !okAddr {
			pCtx.SemanticError("Out of memory")
		}
		rType := NewRType(Float)
		rType.firstDim = 0
		rType.secondDim = 0
		rType.address = next
		pCtx.consTable[ctx.FLOAT_CONS().GetText()] = rType
		pCtx.addrTable[rType.address] = ctx.FLOAT_CONS().GetText()
		return rType
	} else if ctx.STRING_CONS() != nil {
		entry, found := pCtx.consTable[ctx.STRING_CONS().GetText()]
		if found {
			return entry
		}
		next, okAddr := pCtx.vm.constStringAddressMgr.GetNext()
		if !okAddr {
			pCtx.SemanticError("Out of memory")
		}
		rType := NewRType(String)
		rType.firstDim = 0
		rType.secondDim = 0
		rType.address = next
		pCtx.consTable[ctx.STRING_CONS().GetText()] = rType
		pCtx.addrTable[rType.address] = ctx.STRING_CONS().GetText()
		return rType
	}
	entry, found := pCtx.consTable[ctx.BOOL_CONS().GetText()]
	if found {
		return entry
	}
	next, okAddr := pCtx.vm.constBoolAddressMgr.GetNext()
	if !okAddr {
		pCtx.SemanticError("Out of memory")
	}
	rType := NewRType(Bool)
	rType.firstDim = 0
	rType.secondDim = 0
	rType.address = next
	pCtx.consTable[ctx.BOOL_CONS().GetText()] = rType
	pCtx.addrTable[rType.address] = ctx.BOOL_CONS().GetText()

	return rType
}

func (pCtx *PCtx) GetRTypeFromVarName(varName string) (RType, bool) {
	currScope := pCtx.GetCurrentScope()
	entry, ok := pCtx.varTable[currScope][varName]
	if ok {
		return entry, true
	} else {
		if len(pCtx.scopes) > 1 {
			mainScope := pCtx.scopes[0]
			entry, ok := pCtx.varTable[mainScope][varName]
			if ok {
				return entry, true
			}
		}
		return NewRType(Int), false
	}
}

func (pCtx *PCtx) PO() []RType {
	return pCtx.pO
}

func (pCtx *PCtx) POper() []int {
	return pCtx.pOper
}

func (pCtx *PCtx) POTop() RType {
	return pCtx.pO[len(pCtx.pO)-1]
}

func (pCtx *PCtx) POperTop() int {
	return pCtx.pOper[len(pCtx.pOper)-1]
}

func (pCtx *PCtx) POPop() {
	pO := pCtx.pO
	popped := pO[:len(pO)-1]
	pCtx.pO = popped
}

func (pCtx *PCtx) POperPop() {
	pOper := pCtx.pOper
	popped := pOper[:len(pOper)-1]
	pCtx.pOper = popped
}

func (pCtx *PCtx) POPush(op RType) {
	pCtx.pO = append(pCtx.pO, op)
}

func (pCtx *PCtx) POperPush(op int) {
	pCtx.pOper = append(pCtx.pOper, op)
}

func (pCtx *PCtx) POIsEmpty() bool {
	return len(pCtx.pO) == 0
}

func (pCtx *PCtx) POperIsEmpty() bool {
	return len(pCtx.pOper) == 0
}

func (pCtx *PCtx) PrintQuads() {
	for i, quad := range pCtx.vm.Quads() {
		op1 := pCtx.GetVarnameAtAddress(quad.op1)
		op2 := pCtx.GetVarnameAtAddress(quad.op2)
		des := pCtx.GetVarnameAtAddress(quad.destination)
		fmt.Println(i, ". ", OpToString(quad.op), op1, op2, des)
	}
}

func (pCtx *PCtx) ResolveType(vt parser.IVar_typeContext, addrMgr *AddressManager) RType {
	fmt.Println("resolving")
	pType := PTypeFromString(vt.TYPE_PRIMITIVE().GetText())
	rType := RType{pType: pType}
	card := ""
	if vt.CARDINALITY() != nil {
		card = vt.CARDINALITY().GetText()
		re := regexp.MustCompile(`(\[)(\d*)\](\[)?(\d*)\]?`)
		match := re.FindStringSubmatch(card)
		if match[1] != "" {
			if match[2] != "" {
				intV, _ := strconv.Atoi(match[2])
				rType.firstDim = intV
			} else {
				rType.firstDim = -1
			}
		}
		if match[3] != "" {
			if match[4] != "" {
				intV, _ := strconv.Atoi(match[4])
				rType.secondDim = intV
			} else {
				rType.secondDim = -1
			}
		}
	}
	// fmt.Println(card)
	addr, ok := addrMgr.GetNext()
	if !ok {
		pCtx.SemanticError("Out of memory")
	}
	rType.address = addr
	rType.print()
	return rType
}

func (pCtx *PCtx) AddToAddrTable(addr int, varName string) {
	pCtx.addrTable[addr] = varName
}

func (pCtx *PCtx) GetCorrespondingAddressManager(pType PType) *AddressManager {
	if len(pCtx.scopes) == 1 {
		switch pType {
		case Int:
			return pCtx.vm.globalIntAddressMgr
		case Float:
			return pCtx.vm.globalFloatAddressMgr
		case String:
			return pCtx.vm.globalStringAddressMgr
		default:
			return pCtx.vm.globalBoolAddressMgr
		}
	}
	switch pType {
	case Int:
		return pCtx.vm.localIntAddressMgr
	case Float:
		return pCtx.vm.localFloatAddressMgr
	case String:
		return pCtx.vm.localStringAddressMgr
	default:
		return pCtx.vm.localBoolAddressMgr
	}
}

func (pCtx *PCtx) GetCorrespondingTempAddressManager(pType PType) *AddressManager {
	switch pType {
	case Int:
		return pCtx.vm.tempIntAddressMgr
	case Float:
		return pCtx.vm.tempFloatAddressMgr
	case String:
		return pCtx.vm.tempStringAddressMgr
	default:
		return pCtx.vm.tempBoolAddressMgr
	}
}

func (pCtx *PCtx) PrintPo() {
	fmt.Println("print po start")
	for _, po := range pCtx.pO {
		po.print()
	}
	fmt.Println("print po end")
}

func (pCtx *PCtx) HandleGenerateQuadForExpression() {
	oper := Op(pCtx.POperTop())
	pCtx.PrintPo()
	pCtx.POperPop()
	fmt.Println("Generating quad")
	o1 := pCtx.POTop()
	pCtx.POPop()
	o2 := pCtx.POTop()
	pCtx.POPop()
	resultingType := SemanticCubeLookup(o1.pType, o2.pType, oper)
	if resultingType == Undefined {
		pCtx.SemanticError("Cannot perform " + OpToString(oper) + " on " + PTypeToString(o1.pType) + " and " + PTypeToString(o2.pType))
		return
	}
	des, ok := pCtx.GetCorrespondingTempAddressManager(resultingType).GetNext()
	if !ok {
		pCtx.SemanticError("Out of memory")
	}
	quad := NewQuad(oper, o2.address, o1.address, des)
	pCtx.vm.PushQuad(quad)
	rType := NewRType(resultingType)
	rType.address = des
	pCtx.POPush(rType)
}

func (pCtx *PCtx) SemanticError(err string) {
	errors := append(pCtx.semanticErrors, err)
	pCtx.semanticErrors = errors
}

func (pCtx *PCtx) AddScope(scope string) {
	scopes := append(pCtx.scopes, scope)
	pCtx.scopes = scopes
}

func (pCtx *PCtx) SetScope(scopes []string) {
	pCtx.scopes = scopes
}

func (pCtx *PCtx) AddFunction(funcName string, rType RType) {
	pCtx.funcDir[funcName] = rType
	pCtx.varTable[funcName] = make(map[string]RType)
}

func (pCtx *PCtx) GetVarInScope(scope string, varName string) (RType, bool) {
	rType, exists := pCtx.varTable[scope][varName]
	return rType, exists
}

func (pCtx *PCtx) AddVarToScope(scope string, varName string, rType RType) {
	pCtx.varTable[scope][varName] = rType
}

func (pCtx *PCtx) RemoveFunction(funcName string) {
	delete(pCtx.funcDir, funcName)
	delete(pCtx.varTable, funcName)
}
