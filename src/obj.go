package src

type ObjCode struct {
	FuncDir   map[string]FuncData
	ConsTable map[string]RType
	Quads     []Quad
}

func NewObjCode(funcDir map[string]FuncData,
	consTable map[string]RType,
	quads []Quad,
) ObjCode {
	return ObjCode{Quads: quads, FuncDir: funcDir, ConsTable: consTable}
}
