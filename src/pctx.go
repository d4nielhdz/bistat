package src

import (
	parser "bistat/parser"
	"fmt"
	"regexp"
	"strconv"
)

/*
This struct contains methods used by many rules such as searching for variables, functions, scope, pO, pOper, jumps and condJumps manipulation
*/
type PCtx struct {
	scopes         []string
	functions      []string
	funcDir        map[string]FuncData
	paramTable     map[string][]RType
	addrTable      map[int]string
	varTable       map[string]map[string]RType
	consTable      map[string]RType
	semanticErrors []string
	pO             []RType
	pOper          []int
	vm             VM
	jumps          []int
	condJumps      [][]int
	printQuads     bool
	Filename       string
}

func NewPCtx(printQuads bool, filename string) PCtx {
	return PCtx{
		scopes:         make([]string, 0),
		functions:      []string{"main"},
		funcDir:        make(map[string]FuncData),
		paramTable:     make(map[string][]RType),
		varTable:       make(map[string]map[string]RType),
		consTable:      make(map[string]RType),
		semanticErrors: make([]string, 0),
		vm:             NewVM(),
		addrTable:      make(map[int]string),
		pO:             make([]RType, 0),
		pOper:          make([]int, 0),
		jumps:          make([]int, 0),
		printQuads:     printQuads,
		Filename:       filename,
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
	fmt.Println("printing addr table")
	for addr, varName := range pCtx.addrTable {
		fmt.Println(varName)
		fmt.Println(addr)
		for k := range pCtx.varTable {
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
		rType.FirstDim = 0
		rType.SecondDim = 0
		rType.Address = next
		pCtx.consTable[ctx.INT_CONS().GetText()] = rType
		pCtx.addrTable[rType.Address] = ctx.INT_CONS().GetText()
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
		rType.FirstDim = 0
		rType.SecondDim = 0
		rType.Address = next
		pCtx.consTable[ctx.FLOAT_CONS().GetText()] = rType
		pCtx.addrTable[rType.Address] = ctx.FLOAT_CONS().GetText()
		return rType
	} else if ctx.STRING_CONS() != nil {
		pattern := "^\"(.*)\"$"
		regex := regexp.MustCompile(pattern)
		matches := regex.FindStringSubmatch(ctx.STRING_CONS().GetText())
		result := matches[1]
		entry, found := pCtx.consTable[result]
		if found {
			return entry
		}
		next, okAddr := pCtx.vm.constStringAddressMgr.GetNext()
		if !okAddr {
			pCtx.SemanticError("Out of memory")
		}
		rType := NewRType(String)
		rType.FirstDim = 0
		rType.SecondDim = 0
		rType.Address = next
		pCtx.consTable[result] = rType
		pCtx.addrTable[rType.Address] = result
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
	rType.FirstDim = 0
	rType.SecondDim = 0
	rType.Address = next
	pCtx.consTable[ctx.BOOL_CONS().GetText()] = rType
	pCtx.addrTable[rType.Address] = ctx.BOOL_CONS().GetText()

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

// This avoids creating a reference to a reference
func (pCtx *PCtx) IgnoreIfRef(val int) int {
	if val >= GLOBAL_REF_START && val < LOCAL_REF_END {
		return val
	}
	return pCtx.ConstIntUpsert(val)
}

// Used for adding an int value to constant memory if not present
func (pCtx *PCtx) ConstIntUpsert(val int) int {
	entry, found := pCtx.consTable[strconv.Itoa(val)]
	if !found {
		next, okAddr := pCtx.vm.constIntAddressMgr.GetNext()
		if !okAddr {
			pCtx.SemanticError("Out of memory")
		}
		rType := NewRType(Int)
		rType.FirstDim = 0
		rType.SecondDim = 0
		rType.Address = next
		pCtx.consTable[strconv.Itoa(val)] = rType
		pCtx.addrTable[rType.Address] = strconv.Itoa(val)
		entry = rType
	}
	return entry.Address
}

// Used for adding a string value to constant memory if not present
func (pCtx *PCtx) ConstStringUpsert(val string) int {
	entry, found := pCtx.consTable[val]
	if !found {
		next, okAddr := pCtx.vm.constStringAddressMgr.GetNext()
		if !okAddr {
			pCtx.SemanticError("Out of memory")
		}
		rType := NewRType(String)
		rType.FirstDim = 0
		rType.SecondDim = 0
		rType.Address = next
		pCtx.consTable[val] = rType
		pCtx.addrTable[rType.Address] = val
		entry = rType
	}
	return entry.Address
}

func (pCtx *PCtx) IsLocalVar(varName string) bool {
	currScope := pCtx.GetCurrentScope()
	_, ok := pCtx.varTable[currScope][varName]
	return ok && len(pCtx.scopes) > 1
}

func (pCtx *PCtx) IsLocalAddr(addr int) bool {
	return (addr >= LOCAL_INT_START && addr < LOCAL_FLOAT_START) || (addr >= LOCAL_FLOAT_START && addr < LOCAL_STRING_START) || (addr >= LOCAL_STRING_START && addr < LOCAL_BOOL_START) || (addr >= LOCAL_BOOL_START && addr < TEMP_INT_START)
}

func (pCtx *PCtx) IsLocalRef(addr int) bool {
	return (addr >= LOCAL_REF_START && addr < LOCAL_REF_END)
}

func (pCtx *PCtx) IsRef(addr int) bool {
	return (addr >= GLOBAL_REF_START && addr < LOCAL_REF_START) || (addr >= LOCAL_REF_START && addr < LOCAL_INT_START)
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

func (pCtx *PCtx) JumpsTop() int {
	return pCtx.jumps[len(pCtx.jumps)-1]
}

func (pCtx *PCtx) JumpsPop() {
	jumps := pCtx.jumps
	popped := jumps[:len(jumps)-1]
	pCtx.jumps = popped
}

func (pCtx *PCtx) JumpsPush(op int) {
	pCtx.jumps = append(pCtx.jumps, op)
}

func (pCtx *PCtx) JumpsIsEmpty() bool {
	return len(pCtx.jumps) == 0
}

func (pCtx *PCtx) JumpsSize() int {
	return len(pCtx.jumps)
}

func (pCtx *PCtx) CondJumpsTop() int {
	last := len(pCtx.condJumps[len(pCtx.condJumps)-1]) - 1
	return pCtx.condJumps[len(pCtx.condJumps)-1][last]
}

func (pCtx *PCtx) CondJumpsPop() {
	condJumps := pCtx.condJumps[len(pCtx.condJumps)-1]
	popped := condJumps[:len(condJumps)-1]
	pCtx.condJumps[len(pCtx.condJumps)-1] = popped
}

func (pCtx *PCtx) CondJumpsPush(op int) {
	curr := append(pCtx.condJumps[len(pCtx.condJumps)-1], op)
	pCtx.condJumps[len(pCtx.condJumps)-1] = curr
}

func (pCtx *PCtx) PopCondJumps() {
	condJumps := pCtx.condJumps
	popped := condJumps[:len(condJumps)-1]
	pCtx.condJumps = popped
}

func (pCtx *PCtx) PushCondJumps() {
	curr := append(pCtx.condJumps, make([]int, 0))
	pCtx.condJumps = curr
}

func (pCtx *PCtx) CondJumpsIsEmpty() bool {
	if len(pCtx.condJumps) == 0 {
		return false
	}
	return len(pCtx.condJumps[len(pCtx.condJumps)-1]) == 0
}

func (pCtx *PCtx) CondJumpsSize() int {
	if len(pCtx.condJumps) == 0 {
		return 0
	}
	return len(pCtx.condJumps[len(pCtx.condJumps)-1])
}

func (pCtx *PCtx) PrintQuads() {
	for i, quad := range pCtx.vm.Quads() {
		Op1 := pCtx.GetVarnameAtAddress(quad.Op1)
		Op2 := pCtx.GetVarnameAtAddress(quad.Op2)
		des := pCtx.GetVarnameAtAddress(quad.Destination)
		fmt.Println(i, OpToString(quad.Op), Op1, "#"+strconv.Itoa(quad.Op1), Op2, "#"+strconv.Itoa(quad.Op2), des, "#"+strconv.Itoa(quad.Destination))
	}
}

// This method adds an address and dimensions to a variable
func (pCtx *PCtx) ResolveType(vt parser.IVar_typeContext, addrMgr *AddressManager) RType {
	pType := PTypeFromString(vt.TYPE_PRIMITIVE().GetText())
	rType := RType{PType: pType}
	card := ""
	if vt.CARDINALITY() != nil {
		card = vt.CARDINALITY().GetText()
		re := regexp.MustCompile(`(\[)(\d*)\](\[)?(\d*)\]?`)
		match := re.FindStringSubmatch(card)
		if match[1] != "" {
			if match[2] != "" {
				intV, _ := strconv.Atoi(match[2])
				rType.FirstDim = intV
			} else {
				rType.FirstDim = -1
			}
		}
		if match[3] != "" {
			if match[4] != "" {
				intV, _ := strconv.Atoi(match[4])
				rType.SecondDim = intV
			} else {
				rType.SecondDim = -1
			}
		}
	}
	addr, ok := addrMgr.GetNext()
	if rType.FirstDim > 0 {
		size := rType.FirstDim
		i := 0
		if rType.SecondDim > 0 {
			size *= rType.SecondDim
		}
		for i < size-1 {
			_, ok := addrMgr.GetNext()
			if !ok {
				pCtx.SemanticError("Out of memory")
			}
			i += 1
		}
	}
	if !ok {
		pCtx.SemanticError("Out of memory")
	}
	rType.Address = addr
	// rType.print()
	return rType
}

func (pCtx *PCtx) AddToAddrTable(addr int, varName string) {
	pCtx.addrTable[addr] = varName
}

// Gets adequate address manager based on type and scope
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

func (pCtx *PCtx) GetCorrespondingRefAddressManager(addr int) *AddressManager {
	if pCtx.IsLocalAddr(addr) || pCtx.IsLocalRef(addr) || len(pCtx.scopes) > 1 {
		return pCtx.vm.localRefAddressMgr
	}
	return pCtx.vm.globalRefAddressMgr
}

func (pCtx *PCtx) PrintPo() {
	for _, po := range pCtx.pO {
		po.print()
	}
}

// Method used by many rules involving arithmetic and logic operations when quads have to be generated
func (pCtx *PCtx) HandleGenerateQuadForExpression() {
	oper := Op(pCtx.POperTop())
	// pCtx.PrintPo()
	pCtx.POperPop()
	o1 := pCtx.POTop()
	pCtx.POPop()
	if oper == UnaryMinus {
		if o1.PType != Int && o1.PType != Float {
			pCtx.SemanticError("Can't use minus on type " + PTypeToString(o1.PType))
			return
		}
		des, ok := pCtx.GetCorrespondingTempAddressManager(o1.PType).GetNext()
		if !ok {
			pCtx.SemanticError("Out of memory")
			return
		}

		quad := NewQuad(oper, o1.Address, -1, des)
		pCtx.vm.PushQuad(quad)
		rType := NewRType(o1.PType)
		rType.Address = des
		pCtx.POPush(rType)
		return
	}
	o2 := pCtx.POTop()
	pCtx.POPop()
	resultingType := SemanticCubeLookup(o1.PType, o2.PType, oper)
	if resultingType == Undefined {
		pCtx.SemanticError("Cannot perform " + OpToString(oper) + " on " + PTypeToString(o1.PType) + " and " + PTypeToString(o2.PType))
		return
	}
	des, ok := pCtx.GetCorrespondingTempAddressManager(resultingType).GetNext()
	if !ok {
		pCtx.SemanticError("Out of memory")
		return
	}
	quad := NewQuad(oper, o2.Address, o1.Address, des)
	pCtx.vm.PushQuad(quad)
	rType := NewRType(resultingType)
	rType.Address = des
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

func (pCtx *PCtx) AddFunction(funcName string, funcData FuncData) {
	pCtx.funcDir[funcName] = funcData
	pCtx.varTable[funcName] = make(map[string]RType)
}

func (pCtx *PCtx) AddParams(funcName string, params []RType) {
	pCtx.paramTable[funcName] = params
}

func (pCtx *PCtx) GetVarInScope(scope string, varName string) (RType, bool) {
	rType, exists := pCtx.varTable[scope][varName]
	return rType, exists
}

func (pCtx *PCtx) AddVarToScope(scope string, varName string, rType RType) {
	pCtx.varTable[scope][varName] = rType
}

func (pCtx *PCtx) RemoveFunction(funcName string) {
	// delete(pCtx.funcDir, funcName)
	delete(pCtx.varTable, funcName)
}

func (pCtx *PCtx) GenerateQuadsForNonAggregateFunction(op Op, resultType PType) {
	if len(pCtx.semanticErrors) > 0 {
		return
	}
	o := pCtx.POTop()
	pCtx.POPop()
	if op == Not && o.PType != Bool {
		pCtx.SemanticError(OpToString(op) + " can only be called with boolean expressions")
		return
	}
	if (op == Ceil || op == Floor) && o.PType != Float {
		pCtx.SemanticError(OpToString(op) + " can only be called with expressions of type float")
		return
	}
	if (op == Sin || op == Tan || op == Cos || op == Abs || op == Sqrt) && o.PType != Float && o.PType != Int {
		pCtx.SemanticError(OpToString(op) + " can only be called with expressions of type int or float")
		return
	}
	result := NewRType(resultType)
	if o.FirstDim > 0 {
		size := o.FirstDim
		if o.SecondDim > 0 {
			size *= o.SecondDim
		}
		addrMgr := pCtx.GetCorrespondingRefAddressManager(o.Address)
		tempAddMgr := pCtx.GetCorrespondingTempAddressManager(resultType)
		result.FirstDim = o.FirstDim
		result.SecondDim = o.SecondDim
		i := 0
		for i < size {
			refAddr, _ := addrMgr.GetNext()
			tempAddr, _ := tempAddMgr.GetNext()
			if i == 0 {
				result.Address = tempAddr
			}
			pCtx.vm.PushQuad(NewQuad(RefSum, pCtx.IgnoreIfRef(o.Address), pCtx.ConstIntUpsert(i), refAddr))
			pCtx.vm.PushQuad(NewQuad(op, refAddr, -1, tempAddr))
			i++
		}
	} else {
		addr, _ := pCtx.GetCorrespondingTempAddressManager(resultType).GetNext()
		result.Address = addr
		pCtx.vm.PushQuad(NewQuad(op, o.Address, -1, addr))
	}
	pCtx.POPush(result)
}

// Used by all aggregate function rules
func (pCtx *PCtx) GenerateQuadsForAggregateFunction(op Op, resultType PType) {
	if len(pCtx.semanticErrors) > 0 {
		return
	}
	o := pCtx.POTop()
	pCtx.POPop()
	if o.PType != Float && o.PType != Int {
		pCtx.SemanticError(OpToString(op) + " can only be called with lists of type int or float")
		return
	}
	if o.FirstDim <= 0 {
		pCtx.SemanticError(OpToString(op) + " can only be called on arrays or matrices")
		return
	}
	quad := NewQuad(op, o.Address, o.FirstDim, 0)
	result := NewRType(resultType)

	if o.SecondDim > 0 {
		quad.Aux = o.SecondDim
		result.FirstDim = o.FirstDim
		i := 0
		addrMgr := pCtx.GetCorrespondingTempAddressManager(resultType)
		for i != o.FirstDim {
			addr, _ := addrMgr.GetNext()
			if i == 0 {
				quad.Destination = addr
				result.Address = addr
			}
			i++
		}
	} else {
		addrMgr := pCtx.GetCorrespondingTempAddressManager(resultType)
		addr, _ := addrMgr.GetNext()
		quad.Destination = addr
		result.Address = addr
	}
	pCtx.vm.PushQuad(quad)
	pCtx.POPush(result)
}
