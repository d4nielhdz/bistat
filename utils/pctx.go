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
	pO             []int
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
		pO:             make([]int, 0),
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

func (pCtx *PCtx) PO() []int {
	return pCtx.pO
}

func (pCtx *PCtx) POper() []int {
	return pCtx.pOper
}

func (pCtx *PCtx) POTop() int {
	return pCtx.pO[len(pCtx.pO)]
}

func (pCtx *PCtx) POperTop() int {
	return pCtx.pOper[len(pCtx.pOper)]
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

func (pCtx *PCtx) POPush(op int) {
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
