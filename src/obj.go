package src

type ObjCode struct {
	FuncDir         map[string]FuncData
	ConsTable       map[string]RType
	Quads           []Quad
	ConstIntSize    int
	ConstFloatSize  int
	ConstBoolSize   int
	ConstStringSize int
	TempIntSize     int
	TempFloatSize   int
	TempBoolSize    int
	TempStringSize  int
	IntSize         int
	FloatSize       int
	BoolSize        int
	StringSize      int
	GlobalRefSize   int
}

func NewObjCode(funcDir map[string]FuncData,
	consTable map[string]RType,
	quads []Quad,
) ObjCode {
	return ObjCode{Quads: quads, FuncDir: funcDir, ConsTable: consTable}
}
