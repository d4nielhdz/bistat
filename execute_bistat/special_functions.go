package main

import (
	"bistat/src"
	"fmt"
	"math"
	"sort"

	"github.com/guptarohit/asciigraph"
)

// Handles all built in functions

func (eCtx *ECtx) HandlePrint() {
	quad := eCtx.GetCurrentQuad()
	addr := eCtx.GetDerefed(quad.Op1)
	// fmt.Println(eCtx.IP, "#")
	// fmt.Println("t", quad.Op1)
	// fmt.Println("addr", addr)
	pType := GetPTypeFromAddress(addr)

	// fmt.Println(src.PTypeToString(pType))
	if pType == src.Int {
		// fmt.Println("print")
		val := eCtx.GetIntFromAddress(addr)
		fmt.Print(val)
	} else if pType == src.Float {
		val := eCtx.GetFloatFromAddress(addr)
		fmt.Print(val)
	} else if pType == src.Bool {
		val := eCtx.GetBoolFromAddress(addr)
		fmt.Print(val)
	} else if pType == src.String {
		val := eCtx.GetStringFromAddress(addr)
		fmt.Print(val)
	}
}

func (eCtx *ECtx) HandlePrintN() {
	fmt.Print("\n")
}

func (eCtx *ECtx) HandleCos() {
	quad := eCtx.GetCurrentQuad()
	addr1 := eCtx.GetDerefed(quad.Op1)
	addr2 := eCtx.GetDerefed(quad.Destination)

	pType := GetPTypeFromAddress(addr1)
	var cos float64

	if pType == src.Int {
		val := eCtx.GetIntFromAddress(addr1)
		cos = math.Cos(float64(val))
	} else if pType == src.Float {
		val := eCtx.GetFloatFromAddress(addr1)
		cos = math.Cos(val)
	}

	eCtx.StoreFloatAtAddress(cos, addr2)
}

func (eCtx *ECtx) HandleSin() {
	quad := eCtx.GetCurrentQuad()
	addr1 := eCtx.GetDerefed(quad.Op1)
	addr2 := eCtx.GetDerefed(quad.Destination)

	pType := GetPTypeFromAddress(addr1)
	var sin float64

	if pType == src.Int {
		val := eCtx.GetIntFromAddress(addr1)
		sin = math.Sin(float64(val))
	} else if pType == src.Float {
		val := eCtx.GetFloatFromAddress(addr1)
		sin = math.Sin(val)
	}

	eCtx.StoreFloatAtAddress(sin, addr2)
}

func (eCtx *ECtx) HandleTan() {
	quad := eCtx.GetCurrentQuad()
	addr1 := eCtx.GetDerefed(quad.Op1)
	addr2 := eCtx.GetDerefed(quad.Destination)

	pType := GetPTypeFromAddress(addr1)
	var tan float64

	if pType == src.Int {
		val := eCtx.GetIntFromAddress(addr1)
		tan = math.Tan(float64(val))
	} else if pType == src.Float {
		val := eCtx.GetFloatFromAddress(addr1)
		tan = math.Tan(val)
	}

	eCtx.StoreFloatAtAddress(tan, addr2)
}

func (eCtx *ECtx) HandleSqrt() {
	quad := eCtx.GetCurrentQuad()
	addr1 := eCtx.GetDerefed(quad.Op1)
	addr2 := eCtx.GetDerefed(quad.Destination)

	pType := GetPTypeFromAddress(addr1)

	if pType == src.Int {
		val := eCtx.GetIntFromAddress(addr1)
		sqrted := math.Sqrt(float64(val))
		eCtx.StoreIntAtAddress(int64(sqrted), addr2)
	} else if pType == src.Float {
		val := eCtx.GetFloatFromAddress(addr1)
		sqrted := math.Sqrt(float64(val))
		eCtx.StoreFloatAtAddress(sqrted, addr2)
	}
}

func (eCtx *ECtx) HandleFloor() {
	quad := eCtx.GetCurrentQuad()
	addr1 := eCtx.GetDerefed(quad.Op1)
	addr2 := eCtx.GetDerefed(quad.Destination)

	var floored float64

	val := eCtx.GetFloatFromAddress(addr1)
	floored = math.Floor(val)

	eCtx.StoreIntAtAddress(int64(floored), addr2)
}

func (eCtx *ECtx) HandleCeil() {
	quad := eCtx.GetCurrentQuad()
	addr1 := eCtx.GetDerefed(quad.Op1)
	addr2 := eCtx.GetDerefed(quad.Destination)

	var ceiled float64

	val := eCtx.GetFloatFromAddress(addr1)
	ceiled = math.Ceil(val)

	eCtx.StoreIntAtAddress(int64(ceiled), addr2)
}

func (eCtx *ECtx) HandleAbs() {
	quad := eCtx.GetCurrentQuad()
	addr1 := eCtx.GetDerefed(quad.Op1)
	addr2 := eCtx.GetDerefed(quad.Destination)

	pType := GetPTypeFromAddress(addr1)

	if pType == src.Int {
		val := eCtx.GetIntFromAddress(addr1)
		absed := math.Abs(float64(val))
		eCtx.StoreIntAtAddress(int64(absed), addr2)
	} else if pType == src.Float {
		val := eCtx.GetFloatFromAddress(addr1)
		absed := math.Abs(val)
		eCtx.StoreFloatAtAddress(absed, addr2)
	}
}

func (eCtx *ECtx) HandleNot() {
	quad := eCtx.GetCurrentQuad()
	addr1 := eCtx.GetDerefed(quad.Op1)
	addr2 := eCtx.GetDerefed(quad.Destination)

	val := eCtx.GetBoolFromAddress(addr1)
	eCtx.StoreBoolAtAddress(!val, addr2)
}

func (eCtx *ECtx) HandleListSum() {
	quad := eCtx.GetCurrentQuad()
	addr1 := eCtx.GetDerefed(quad.Op1)
	destination := eCtx.GetDerefed(quad.Destination)
	firstDim := quad.Op2
	secondDim := quad.Aux
	pType := GetPTypeFromAddress(addr1)
	size := firstDim
	resultSize := 1
	len := firstDim

	if secondDim > 0 {
		size *= secondDim
		resultSize = firstDim
		len = secondDim
	}

	if pType == src.Int {
		original := eCtx.GetIntListFromAddress(addr1, size)
		result := make([]int64, resultSize)
		i := 0
		var acum int64 = 0
		for i != size {
			acum += original[i]
			if (i+1)%len == 0 {
				idx := ((i + 1) / len) - 1
				result[idx] = acum
				acum = 0
			}
			i++
		}
		eCtx.StoreIntListAtAddress(destination, result)
	} else {
		original := eCtx.GetFloatListFromAddress(addr1, size)
		result := make([]float64, resultSize)
		i := 0
		var acum float64 = 0
		for i != size {
			acum += original[i]
			if (i+1)%len == 0 {
				idx := ((i + 1) / len) - 1
				result[idx] = acum
				acum = 0
			}
			i++
		}
		eCtx.StoreFloatListAtAddress(destination, result)
	}
}

func (eCtx *ECtx) HandleProd() {
	quad := eCtx.GetCurrentQuad()
	addr1 := eCtx.GetDerefed(quad.Op1)
	destination := eCtx.GetDerefed(quad.Destination)
	firstDim := quad.Op2
	secondDim := quad.Aux
	pType := GetPTypeFromAddress(addr1)
	size := firstDim
	resultSize := 1
	len := firstDim

	if secondDim > 0 {
		size *= secondDim
		resultSize = firstDim
		len = secondDim
	}

	if pType == src.Int {
		original := eCtx.GetIntListFromAddress(addr1, size)
		result := make([]int64, resultSize)
		i := 0
		var acum int64 = 1
		for i != size {
			acum *= original[i]
			if (i+1)%len == 0 {
				idx := ((i + 1) / len) - 1
				result[idx] = acum
				acum = 1
			}
			i++
		}
		eCtx.StoreIntListAtAddress(destination, result)
	} else {
		original := eCtx.GetFloatListFromAddress(addr1, size)
		result := make([]float64, resultSize)
		i := 0
		var acum float64 = 1
		for i != size {
			acum *= original[i]
			if (i+1)%len == 0 {
				idx := ((i + 1) / len) - 1
				result[idx] = acum
				acum = 1
			}
			i++
		}
		eCtx.StoreFloatListAtAddress(destination, result)
	}
}

func (eCtx *ECtx) HandleMax() {
	quad := eCtx.GetCurrentQuad()
	addr1 := eCtx.GetDerefed(quad.Op1)
	destination := eCtx.GetDerefed(quad.Destination)
	firstDim := quad.Op2
	secondDim := quad.Aux
	pType := GetPTypeFromAddress(addr1)
	size := firstDim
	resultSize := 1
	len := firstDim

	if secondDim > 0 {
		size *= secondDim
		resultSize = firstDim
		len = secondDim
	}

	if pType == src.Int {
		original := eCtx.GetIntListFromAddress(addr1, size)
		result := make([]int64, resultSize)
		i := 0
		var currMax int64 = math.MinInt64
		for i != size {
			currMax = int64(math.Max(float64(original[i]), float64(currMax)))
			if (i+1)%len == 0 {
				idx := ((i + 1) / len) - 1
				result[idx] = currMax
				currMax = math.MinInt64
			}
			i++
		}
		eCtx.StoreIntListAtAddress(destination, result)
	} else {
		original := eCtx.GetFloatListFromAddress(addr1, size)
		result := make([]float64, resultSize)
		i := 0
		var currMax float64 = -math.MaxFloat64
		for i != size {
			currMax = math.Max(original[i], currMax)
			if (i+1)%len == 0 {
				idx := ((i + 1) / len) - 1
				result[idx] = currMax
				currMax = float64(math.MinInt64)
			}
			i++
		}
		eCtx.StoreFloatListAtAddress(destination, result)
	}
}

func (eCtx *ECtx) HandleMin() {
	quad := eCtx.GetCurrentQuad()
	addr1 := eCtx.GetDerefed(quad.Op1)
	destination := eCtx.GetDerefed(quad.Destination)
	firstDim := quad.Op2
	secondDim := quad.Aux
	pType := GetPTypeFromAddress(addr1)
	size := firstDim
	resultSize := 1
	len := firstDim

	if secondDim > 0 {
		size *= secondDim
		resultSize = firstDim
		len = secondDim
	}

	if pType == src.Int {
		original := eCtx.GetIntListFromAddress(addr1, size)
		result := make([]int64, resultSize)
		i := 0
		var currMax int64 = math.MaxInt64
		for i != size {
			currMax = int64(math.Min(float64(original[i]), float64(currMax)))
			if (i+1)%len == 0 {
				idx := ((i + 1) / len) - 1
				result[idx] = currMax
				currMax = math.MaxInt64
			}
			i++
		}
		eCtx.StoreIntListAtAddress(destination, result)
	} else {
		original := eCtx.GetFloatListFromAddress(addr1, size)
		result := make([]float64, resultSize)
		i := 0
		var currMax float64 = math.MaxFloat64
		for i != size {
			currMax = math.Min(original[i], currMax)
			if (i+1)%len == 0 {
				idx := ((i + 1) / len) - 1
				result[idx] = currMax
				currMax = math.MaxInt64
			}
			i++
		}
		eCtx.StoreFloatListAtAddress(destination, result)
	}
}

func (eCtx *ECtx) HandleAvg() {
	quad := eCtx.GetCurrentQuad()
	addr1 := eCtx.GetDerefed(quad.Op1)
	destination := eCtx.GetDerefed(quad.Destination)
	firstDim := quad.Op2
	secondDim := quad.Aux
	pType := GetPTypeFromAddress(addr1)
	size := firstDim
	resultSize := 1
	n := firstDim

	if secondDim > 0 {
		size *= secondDim
		resultSize = firstDim
		n = secondDim
	}

	if pType == src.Int {
		original := eCtx.GetIntListFromAddress(addr1, size)
		result := make([]float64, resultSize)
		i := 0
		var acum int64 = 0
		for i != size {
			acum += original[i]
			if (i+1)%n == 0 {
				idx := ((i + 1) / n) - 1
				result[idx] = float64(acum) / float64(n)
				acum = 0
			}
			i++
		}
		eCtx.StoreFloatListAtAddress(destination, result)
	} else {
		// fmt.Println("--", quad.Op1, addr1, quad.Op2, quad.Aux)
		original := eCtx.GetFloatListFromAddress(addr1, size)
		result := make([]float64, resultSize)
		i := 0
		var acum float64 = 0
		for i != size {
			acum += original[i]
			if (i+1)%n == 0 {
				idx := ((i + 1) / n) - 1
				result[idx] = acum / float64(n)
				acum = 0
			}
			i++
		}
		eCtx.StoreFloatListAtAddress(destination, result)
	}
}

func (eCtx *ECtx) HandleSMode() {
	quad := eCtx.GetCurrentQuad()
	addr1 := eCtx.GetDerefed(quad.Op1)
	destination := eCtx.GetDerefed(quad.Destination)
	firstDim := quad.Op2
	secondDim := quad.Aux
	pType := GetPTypeFromAddress(addr1)
	size := firstDim
	resultSize := 1
	len := firstDim

	if secondDim > 0 {
		size *= secondDim
		resultSize = firstDim
		len = secondDim
	}

	if pType == src.Int {
		original := eCtx.GetIntListFromAddress(addr1, size)
		result := make([]int64, resultSize)
		freqs := make(map[int64]int)
		i := 0
		for i != size {
			freqs[original[i]] += 1

			if (i+1)%len == 0 {
				maxFreq := 0
				var modeVal int64 = 0
				for val, freq := range freqs {
					if freq > maxFreq {
						maxFreq = freq
						modeVal = val
					}
				}

				idx := ((i + 1) / len) - 1
				result[idx] = modeVal
				freqs = make(map[int64]int)
			}
			i++
		}
		eCtx.StoreIntListAtAddress(destination, result)
	} else {
		original := eCtx.GetFloatListFromAddress(addr1, size)
		result := make([]float64, resultSize)
		freqs := make(map[float64]int)
		i := 0
		for i != size {
			freqs[original[i]] += 1

			if (i+1)%len == 0 {
				maxFreq := 0
				var modeVal float64 = 0
				for val, freq := range freqs {
					if freq > maxFreq {
						maxFreq = freq
						modeVal = val
					}
				}

				idx := ((i + 1) / len) - 1
				result[idx] = modeVal
				freqs = make(map[float64]int)
			}
			i++
		}
		eCtx.StoreFloatListAtAddress(destination, result)
	}
}

func (eCtx *ECtx) HandleMedian() {
	quad := eCtx.GetCurrentQuad()
	addr1 := eCtx.GetDerefed(quad.Op1)
	destination := eCtx.GetDerefed(quad.Destination)
	firstDim := quad.Op2
	secondDim := quad.Aux
	pType := GetPTypeFromAddress(addr1)
	size := firstDim
	resultSize := 1
	len := firstDim

	if secondDim > 0 {
		size *= secondDim
		resultSize = firstDim
		len = secondDim
	}

	if pType == src.Int {
		original := eCtx.GetIntListFromAddress(addr1, size)
		result := make([]float64, resultSize)
		ls := make([]int64, len)
		i := 0
		j := 0
		for i != size {
			ls[j] = original[i]
			var median float64 = 0
			if (j + 1) == len {
				sort.Slice(ls, func(k, l int) bool {
					return ls[k] < ls[l]
				})
				if len%2 == 0 {
					median = float64((ls[len/2-1] + ls[len/2])) / 2.0
				} else {
					median = float64(ls[len/2])
				}
				idx := ((i + 1) / len) - 1
				result[idx] = median
				ls = make([]int64, len)
				j = 0
			} else {
				j++
			}
			i++
		}
		eCtx.StoreFloatListAtAddress(destination, result)
	} else {
		original := eCtx.GetFloatListFromAddress(addr1, size)
		result := make([]float64, resultSize)
		ls := make([]float64, len)
		i := 0
		j := 0
		for i != size {
			ls[j] = original[i]
			var median float64 = 0
			if (j + 1) == len {
				sort.Slice(ls, func(k, l int) bool {
					return ls[k] < ls[l]
				})
				if len%2 == 0 {
					median = (ls[len/2-1] + ls[len/2]) / 2
				} else {
					median = ls[len/2]
				}
				idx := ((i + 1) / len) - 1
				result[idx] = median
				ls = make([]float64, len)
				j = 0
			} else {
				j++
			}
			i++
		}
		eCtx.StoreFloatListAtAddress(destination, result)
	}
}

func (eCtx *ECtx) HandlePlot() {
	quad := eCtx.GetCurrentQuad()
	addr1 := eCtx.GetDerefed(quad.Op1)
	firstDim := quad.Op2
	secondDim := quad.Destination
	pType := GetPTypeFromAddress(addr1)
	size := firstDim
	len := firstDim
	m := 1

	if secondDim > 0 {
		size *= secondDim
		len = secondDim
		m = firstDim
	}

	if pType == src.Int {
		data := make([][]float64, m)
		k := 0
		for k != m {
			data[k] = make([]float64, len)
			k++
		}
		i := 0
		j := 0
		original := eCtx.GetIntListFromAddress(addr1, size)
		for _, val := range original {
			data[i][j] = float64(val)
			if (j + 1) == len {
				j = 0
				i++
			} else {
				j++
			}
		}
		graph := asciigraph.PlotMany(data)
		fmt.Println(graph)
	} else {
		data := make([][]float64, m)
		k := 0
		for k != m {
			data[k] = make([]float64, len)
			k++
		}
		i := 0
		j := 0
		original := eCtx.GetFloatListFromAddress(addr1, size)
		for _, val := range original {
			data[i][j] = val
			if (j + 1) == len {
				j = 0
				i++
			} else {
				j++
			}
		}
		graph := asciigraph.PlotMany(data)
		fmt.Println(graph)
	}
}
