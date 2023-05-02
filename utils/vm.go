package utils

const GLOBAL_INT_START = 0
const GLOBAL_FLOAT_START = 10000
const GLOBAL_STRING_START = 20000
const GLOBAL_BOOL_START = 30000
const GLOBAL_BOOL_END = 39999
const LOCAL_INT_START = 40000
const LOCAL_FLOAT_START = 50000
const LOCAL_STRING_START = 60000
const LOCAL_BOOL_START = 70000
const TEMP_INT_START = 80000
const TEMP_FLOAT_START = 90000
const TEMP_STRING_START = 100000
const TEMP_BOOL_START = 110000
const CONST_INT_START = 120000
const CONST_FLOAT_START = 130000
const CONST_STRING_START = 140000
const CONST_BOOL_START = 150000
const CONST_BOOL_END = 159999

type VM struct {
	quads                                                                                                                                                                                                                                                                                                                                                  []Quad
	globalIntAddressMgr, localIntAddressMgr, tempIntAddressMgr, constIntAddressMgr, globalFloatAddressMgr, localFloatAddressMgr, tempFloatAddressMgr, constFloatAddressMgr, globalStringAddressMgr, localStringAddressMgr, tempStringAddressMgr, constStringAddressMgr, globalBoolAddressMgr, localBoolAddressMgr, tempBoolAddressMgr, constBoolAddressMgr *AddressManager
}

func (vm *VM) Quads() []Quad {
	return vm.quads
}

func (vm *VM) GlobalIntAddressMgr() *AddressManager {
	return vm.globalIntAddressMgr
}

func (vm *VM) LocalIntAddressMgr() *AddressManager {
	return vm.localIntAddressMgr
}

func (vm *VM) TempIntAddressMgr() *AddressManager {
	return vm.tempIntAddressMgr
}

func (vm *VM) ConstIntAddressMgr() *AddressManager {
	return vm.constIntAddressMgr
}

func (vm *VM) GlobalFloatAddressMgr() *AddressManager {
	return vm.globalFloatAddressMgr
}

func (vm *VM) LocalFloatAddressMgr() *AddressManager {
	return vm.localFloatAddressMgr
}

func (vm *VM) TempFloatAddressMgr() *AddressManager {
	return vm.tempFloatAddressMgr
}

func (vm *VM) ConstFloatAddressMgr() *AddressManager {
	return vm.constFloatAddressMgr
}

func (vm *VM) GlobalStringAddressMgr() *AddressManager {
	return vm.globalStringAddressMgr
}

func (vm *VM) LocalStringAddressMgr() *AddressManager {
	return vm.localStringAddressMgr
}

func (vm *VM) TempStringAddressMgr() *AddressManager {
	return vm.tempStringAddressMgr
}

func (vm *VM) ConstStringAddressMgr() *AddressManager {
	return vm.constStringAddressMgr
}

func (vm *VM) GlobalBoolAddressMgr() *AddressManager {
	return vm.globalBoolAddressMgr
}

func (vm *VM) LocalBoolAddressMgr() *AddressManager {
	return vm.localBoolAddressMgr
}

func (vm *VM) TempBoolAddressMgr() *AddressManager {
	return vm.tempBoolAddressMgr
}

func (vm *VM) ConstBoolAddressMgr() *AddressManager {
	return vm.constBoolAddressMgr
}

func NewVM() VM {
	return VM{
		quads:                 make([]Quad, 0),
		globalIntAddressMgr:   NewAddressManager(GLOBAL_INT_START, GLOBAL_FLOAT_START-1),
		globalFloatAddressMgr: NewAddressManager(GLOBAL_FLOAT_START, GLOBAL_BOOL_START-1), globalStringAddressMgr: NewAddressManager(GLOBAL_STRING_START, GLOBAL_BOOL_START-1), globalBoolAddressMgr: NewAddressManager(GLOBAL_BOOL_START, LOCAL_INT_START-1),
		localIntAddressMgr:    NewAddressManager(LOCAL_INT_START, LOCAL_FLOAT_START-1),
		localFloatAddressMgr:  NewAddressManager(LOCAL_FLOAT_START, LOCAL_STRING_START-1),
		localStringAddressMgr: NewAddressManager(LOCAL_STRING_START, LOCAL_BOOL_START-1),
		localBoolAddressMgr:   NewAddressManager(LOCAL_BOOL_START, TEMP_INT_START-1),
		tempIntAddressMgr:     NewAddressManager(TEMP_INT_START, TEMP_FLOAT_START-1),
		tempFloatAddressMgr:   NewAddressManager(TEMP_FLOAT_START, TEMP_STRING_START-1),
		tempStringAddressMgr:  NewAddressManager(TEMP_STRING_START, TEMP_BOOL_START-1),
		tempBoolAddressMgr:    NewAddressManager(TEMP_BOOL_START, CONST_INT_START-1),
		constIntAddressMgr:    NewAddressManager(CONST_INT_START, CONST_FLOAT_START-1),
		constFloatAddressMgr:  NewAddressManager(CONST_FLOAT_START, CONST_STRING_START-1),
		constStringAddressMgr: NewAddressManager(CONST_STRING_START, CONST_BOOL_START-1),
		constBoolAddressMgr:   NewAddressManager(CONST_BOOL_START, CONST_BOOL_END),
	}
}

type Quad struct {
	op                    Op
	op1, op2, destination int
}

func NewQuad(op Op, op1 int, op2 int, destination int) Quad {
	return Quad{op: op, op1: op1, op2: op2, destination: destination}
}

func (vm *VM) PushQuad(quad Quad) {
	vm.quads = append(vm.quads, quad)
}

type AddressManager struct {
	first, curr, last int
}

func NewAddressManager(first int, last int) *AddressManager {
	return &AddressManager{first: first, last: last, curr: first}
}

func (am *AddressManager) GetNext() (int, bool) {
	curr := am.curr
	if curr > am.last {
		return -1, false
	}
	am.curr = curr + 1
	return curr, true
}

func (am *AddressManager) Reset() {
	am.curr = am.first
}
