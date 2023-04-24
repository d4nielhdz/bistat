package utils

import (
	parser "bistat/parser"
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
)

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

func ResolveType(vt parser.IVar_typeContext) int {
	fmt.Println((vt).TYPE_PRIMITIVE())
	fmt.Println((vt).CARDINALITY())
	fmt.Println(Undefined)
	fmt.Println(Void)
	return 1
}
