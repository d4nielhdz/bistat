package src

// This is what is stored in the binary .gob file after compilation
type ObjCode struct {
	FuncDir         map[string]FuncData
	ConsTable       map[string]RType
	Functions       []string
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
	functions []string,
) ObjCode {
	return ObjCode{Quads: quads, FuncDir: funcDir, ConsTable: consTable, Functions: functions}
}
