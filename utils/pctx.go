package utils

type PCtx struct {
	scopes    []string
	funcDir   map[string]RType
	varTable  map[string]map[string]RType
	consTable map[string]RType
}

func NewPCtx() PCtx {
	return PCtx{
		scopes:    make([]string, 0),
		funcDir:   make(map[string]RType),
		varTable:  make(map[string]map[string]RType),
		consTable: make(map[string]RType),
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

func (pCtx *PCtx) AddScope(scope string) {
	scopes := append(pCtx.scopes, scope)
	pCtx.scopes = scopes
}

func (pCtx *PCtx) SetScope(scopes []string) {
	pCtx.scopes = scopes
}

func (pCtx *PCtx) AddFunction(funcName string, rType RType) {
	pCtx.funcDir[funcName] = rType
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
