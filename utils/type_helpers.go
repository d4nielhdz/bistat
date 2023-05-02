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
	if len(semanticCube) >= int(op1) {
		return Undefined
	}
	if len(semanticCube[op1]) >= int(op2) {
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
