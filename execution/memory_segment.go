package main

import (
	src "bistat/src"
	"fmt"
	"strconv"
)

type MemorySegment struct {
	floats         []float64
	ints           []int64
	strings        []string
	bools          []bool
	baseFloatAddr  int
	baseIntAddr    int
	baseStringAddr int
	baseBoolAddr   int
}

func NewMemorySegment(fSize int, iSize int, sSize int, bSize int, baseFloatAddr int, baseIntAddr int, baseStringAddr int, baseBoolAddr int) *MemorySegment {
	return &MemorySegment{
		floats:         make([]float64, fSize),
		ints:           make([]int64, iSize),
		strings:        make([]string, sSize),
		bools:          make([]bool, bSize),
		baseFloatAddr:  baseFloatAddr,
		baseIntAddr:    baseIntAddr,
		baseBoolAddr:   baseBoolAddr,
		baseStringAddr: baseStringAddr,
	}
}
func (ms *MemorySegment) StoreVal(val string, addr int) {
	pType := GetPTypeFromAddress(addr)
	if pType == src.Int {
		converted, _ := strconv.Atoi(val)
		ms.ints[addr-ms.baseIntAddr] = int64(converted)
	} else if pType == src.Float {
		converted, _ := strconv.ParseFloat(val, 64)
		ms.floats[addr-ms.baseFloatAddr] = converted
	} else if pType == src.String {
		ms.strings[addr-ms.baseStringAddr] = val
	} else {
		converted, _ := strconv.ParseBool(val)
		ms.bools[addr-ms.baseBoolAddr] = converted
	}
}

func (ms *MemorySegment) Print() {
	fmt.Println("--print ints start")
	for _, val := range ms.ints {
		fmt.Println(val)
	}
	fmt.Println("--print ints end")
	fmt.Println("--print floats start")
	for _, val := range ms.floats {
		fmt.Println(val)
	}
	fmt.Println("--print floats end")
	fmt.Println("--print strings start")
	for _, val := range ms.strings {
		fmt.Println(val)
	}
	fmt.Println("--print strings end")
	fmt.Println("--print bools start")
	for _, val := range ms.bools {
		fmt.Println(val)
	}
	fmt.Println("--print bools end")
}

// func (ms *MemorySegment) StoreInt(val int64, addr int) {
// 	ms.ints[addr-ms.baseIntAddr] = val
// }

// func (ms *MemorySegment) StoreFloat(val float64, addr int) {
// 	ms.floats[addr-ms.baseFloatAddr] = val
// }

// func (ms *MemorySegment) StoreString(val string, addr int) {
// 	ms.strings[addr-ms.baseStringAddr] = val
// }

// func (ms *MemorySegment) StoreBool(val bool, addr int) {
// 	ms.bools[addr-ms.baseBoolAddr] = val
// }

type StackSegment struct {
	localMemory *MemorySegment
	tempMemory  *MemorySegment
}

func NewStackSegment(fSize int, iSize int, sSize int, bSize int, baseFloatAddr int, baseIntAddr int, baseStringAddr int, baseBoolAddr int, tempFSize int, tempISize int, tempSSize int, tempBSize int, baseTempFloatAddr int, baseTempIntAddr int, baseTempStringAddr int, baseTempBoolAddr int) *StackSegment {
	return &StackSegment{
		localMemory: NewMemorySegment(fSize, iSize, sSize, bSize, baseFloatAddr, baseIntAddr, baseStringAddr, baseBoolAddr),
		tempMemory:  NewMemorySegment(tempFSize, tempISize, tempSSize, tempBSize, baseTempFloatAddr, baseTempIntAddr, baseTempStringAddr, baseTempBoolAddr),
	}
}
