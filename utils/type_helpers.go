package utils

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
	Sum Op = iota
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
	ListAdd
	ListPop
	Length
	Range
	Plot
	ListSum
	Min
	Prod
	Avg
	SMode
	Median
	Sin
	Cos
	Tan
	Sort
	Sqrt
	Floor
	Ceil
	Abs
	Not
	UnaryMinus
	UndefinedOp
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
	case Assign:
		return "Assign"
	case InputRead:
		return "InputRead"
	case Print:
		return "Print"
	case ListAdd:
		return "ListAdd"
	case ListPop:
		return "ListPop"
	case Length:
		return "Length"
	case Range:
		return "Range"
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
	case Sort:
		return "Sort"
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
	case Other:
		return "Other"
	default:
		return ""
	}
}

var semanticCube = [3][3][13]PType{
	{
		// Int Op Int
		{Int, Int, Int, Int, Int, Bool, Bool, Bool, Bool, Bool, Bool, Bool, Bool},
		// Int Op Float
		{Float, Float, Float, Float, Undefined, Bool, Bool, Bool, Bool, Bool, Bool, Bool, Bool},
		// Int Op Bool
		{Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined},
	},
	{
		// Float Op Int
		{Float, Float, Float, Float, Undefined, Bool, Bool, Bool, Bool, Bool, Bool, Bool, Bool},
		// Float Op Float
		{Float, Float, Float, Float, Undefined, Bool, Bool, Bool, Bool, Bool, Bool, Bool, Bool},
		// Float Op Bool
		{Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined},
	},
	{
		// Bool Op Int
		{Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined},
		// Bool Op Float
		{Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined, Undefined},
		// Bool Op Bool
		{Undefined, Undefined, Undefined, Undefined, Undefined, Bool, Bool, Bool, Bool, Bool, Bool, Bool, Bool},
	},
}

func SemanticCubeLookup(op1 PType, op2 PType, op Op) PType {
	fmt.Println("Checking cube ", int(op1), int(op2))
	if int(op1) >= len(semanticCube) {
		return Undefined
	}
	if int(op2) >= len(semanticCube[op1]) {
		return Undefined
	}
	return semanticCube[op1][op2][op]
}

type RType struct {
	pType PType
	// indicates size of first dimension. -1 for params, 0 for non arrays
	firstDim int
	// indicates size of second dimension. -1 for params, 0 for non arrays
	secondDim int
	address   int
}

func NewRType(pType PType) RType {
	return RType{pType: pType}
}

func (rType RType) print() {
	fmt.Println("RType:")
	fmt.Println(rType.pType)
	fmt.Println("first dim", rType.firstDim)
	fmt.Println("second dim", rType.secondDim)
	fmt.Println("address #", rType.address)
}
