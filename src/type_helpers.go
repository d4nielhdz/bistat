package src

import (
	"fmt"
)

type PType int
type Op int

const (
	Int PType = iota
	Float
	Bool
	String
	Void
	Undefined
)

func PTypeFromString(str string) PType {
	switch str {
	case "int":
		return Int
	case "float":
		return Float
	case "bool":
		return Bool
	case "string":
		return String
	case "void":
		return Void
	default:
		return Undefined
	}
}

func PTypeToString(pType PType) string {
	switch pType {
	case Int:
		return "int"
	case Float:
		return "float"
	case Bool:
		return "bool"
	case String:
		return "string"
	case Void:
		return "void"
	default:
		return "undefined"
	}
}

const (
	Sum Op = iota //
	Subtraction
	Multiplication
	Division
	Modulus
	And
	Or
	Gt
	Lt
	Ge
	Le
	Eq
	Ne
	Assign
	InputRead
	Print
	PrintN
	ListSum
	Min
	Max
	Prod
	Avg
	SMode
	Median
	Plot
	Sin
	Cos
	Tan
	Sqrt
	Floor
	Ceil
	Abs
	Not
	UnaryMinus
	Goto
	GotoT
	GotoF
	GoSub
	Era
	Param
	EndFunc
	End
	Return
	RefSum
	RefMul
	UndefinedOp
	Verify
	Other
)

func OpFromString(str string) Op {
	switch str {
	case "+":
		return Sum
	case "-":
		return Subtraction
	case "*":
		return Multiplication
	case "/":
		return Division
	case "%":
		return Modulus
	case "&&":
		return And
	case "||":
		return Or
	case ">":
		return Gt
	case "<":
		return Lt
	case ">=":
		return Ge
	case "<=":
		return Le
	case "==":
		return Eq
	case "!=":
		return Ne
	default:
		return UndefinedOp
	}
}

func OpToString(op Op) string {
	switch op {
	case Sum:
		return "Sum"
	case Subtraction:
		return "Subtraction"
	case Multiplication:
		return "Multiplication"
	case Division:
		return "Division"
	case Modulus:
		return "Modulus"
	case And:
		return "And"
	case Or:
		return "Or"
	case Gt:
		return "Gt"
	case Lt:
		return "Lt"
	case Ge:
		return "Ge"
	case Le:
		return "Le"
	case Eq:
		return "Eq"
	case Ne:
		return "Ne"
	case RefMul:
		return "RefMul"
	case RefSum:
		return "RefSum"
	case Assign:
		return "Assign"
	case InputRead:
		return "InputRead"
	case Print:
		return "Print"
	case PrintN:
		return "PrintN"
	case Plot:
		return "Plot"
	case ListSum:
		return "ListSum"
	case Min:
		return "Min"
	case Prod:
		return "Prod"
	case Avg:
		return "Avg"
	case SMode:
		return "SMode"
	case Median:
		return "Median"
	case Sin:
		return "Sin"
	case Cos:
		return "Cos"
	case Tan:
		return "Tan"
	case Sqrt:
		return "Sqrt"
	case Floor:
		return "Floor"
	case Ceil:
		return "Ceil"
	case Abs:
		return "Abs"
	case Not:
		return "Not"
	case UnaryMinus:
		return "UnaryMinus"
	case UndefinedOp:
		return "UndefinedOp"
	case Goto:
		return "Goto"
	case GotoF:
		return "GotoF"
	case GotoT:
		return "GotoT"
	case Other:
		return "Other"
	case GoSub:
		return "GoSub"
	case Era:
		return "Era"
	case Param:
		return "Param"
	case EndFunc:
		return "EndFunc"
	case End:
		return "End"
	case Return:
		return "Return"
	case Verify:
		return "Verify"
	default:
		return ""
	}
}

var semanticCube = [3][3][13]PType{
	{
		// Int Op Int
		{Int, Int, Int, Int, Int, Undefined, Undefined, Bool, Bool, Bool, Bool, Bool, Bool},
		// Int Op Float
		{Float, Float, Float, Float, Undefined, Undefined, Undefined, Bool, Bool, Bool, Bool, Bool, Bool},
		// Int Op Bool
		{Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined},
	},
	{
		// Float Op Int
		{Float, Float, Float, Float, Undefined, Undefined, Undefined, Bool, Bool, Bool, Bool, Bool, Bool},
		// Float Op Float
		{Float, Float, Float, Float, Undefined, Undefined, Undefined, Bool, Bool, Bool, Bool, Bool, Bool},
		// Float Op Bool
		{Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined},
	},
	{
		// Bool Op Int
		{Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined},
		// Bool Op Float
		{Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined},
		// Bool Op Bool
		{Undefined, Undefined, Undefined, Undefined, Undefined, Bool, Bool, Undefined, Undefined, Undefined, Undefined, Bool, Bool},
	},
}

func SemanticCubeLookup(op1 PType, op2 PType, op Op) PType {
	// fmt.Println("Checking cube ", int(op1), int(op2))
	if int(op1) >= len(semanticCube) {
		return Undefined
	}
	if int(op2) >= len(semanticCube[op1]) {
		return Undefined
	}
	return semanticCube[op1][op2][op]
}

type RType struct {
	PType PType
	// indicates size of first dimension. -1 for params, 0 for non arrays
	FirstDim int
	// indicates size of second dimension. -1 for params, 0 for non arrays
	SecondDim int
	Address   int
}

type FuncData struct {
	PType PType
	// indicates size of first dimension. -1 for params, 0 for non arrays
	FirstDim int
	// indicates size of second dimension. -1 for params, 0 for non arrays
	SecondDim       int
	Address         int
	Params          int
	LocalRefs       int
	LocalIntVars    int
	TempIntVars     int
	LocalFloatVars  int
	TempFloatVars   int
	LocalBoolVars   int
	TempBoolVars    int
	LocalStringVars int
	TempStringVars  int
	FuncStart       int
	Idx             int
}

func NewFuncData(pType PType) FuncData {
	return FuncData{PType: pType}
}

func NewRType(pType PType) RType {
	return RType{PType: pType}
}

func RTypeToFuncData(rType RType) FuncData {
	return FuncData{PType: rType.PType, FirstDim: rType.FirstDim, SecondDim: rType.SecondDim, Address: rType.Address}
}

func (rType RType) print() {
	fmt.Println("RType:")
	fmt.Println(rType.PType)
	fmt.Println("first dim", rType.FirstDim)
	fmt.Println("second dim", rType.SecondDim)
	fmt.Println("address #", rType.Address)
}
