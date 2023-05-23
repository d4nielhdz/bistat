package main

import src "bistat/src"

func (eCtx *ECtx) GetIntFromAddress(addr int) int64 {
	if addr >= src.CONST_INT_START && addr < src.CONST_FLOAT_START {
		return eCtx.ConstMemory.ints[addr-eCtx.ConstMemory.baseIntAddr]
	} else if addr >= src.GLOBAL_INT_START && addr < src.GLOBAL_FLOAT_START {
		return eCtx.GlobalMemory.ints[addr-eCtx.GlobalMemory.baseIntAddr]
	} else if addr >= src.TEMP_INT_START && addr < src.TEMP_FLOAT_START {
		if len(eCtx.StackSegment) > 0 {
			return eCtx.GetStackSegment().tempMemory.ints[addr-eCtx.GetStackSegment().tempMemory.baseIntAddr]
		}
		return eCtx.TempMemory.ints[addr-eCtx.TempMemory.baseIntAddr]
	}
	return eCtx.GetStackSegment().localMemory.ints[addr-eCtx.GetStackSegment().localMemory.baseIntAddr]
}

func (eCtx *ECtx) GetFloatFromAddress(addr int) float64 {
	if addr >= src.CONST_FLOAT_START && addr < src.CONST_STRING_START {
		return eCtx.ConstMemory.floats[addr-eCtx.ConstMemory.baseFloatAddr]
	} else if addr >= src.GLOBAL_FLOAT_START && addr < src.GLOBAL_STRING_START {
		return eCtx.GlobalMemory.floats[addr-eCtx.GlobalMemory.baseFloatAddr]
	} else if addr >= src.TEMP_FLOAT_START && addr < src.TEMP_STRING_START {
		if len(eCtx.StackSegment) > 0 {
			return eCtx.GetStackSegment().tempMemory.floats[addr-eCtx.GetStackSegment().tempMemory.baseFloatAddr]
		}
		return eCtx.TempMemory.floats[addr-eCtx.TempMemory.baseFloatAddr]
	}
	return eCtx.GetStackSegment().localMemory.floats[addr-eCtx.GetStackSegment().localMemory.baseFloatAddr]
}
func (eCtx *ECtx) GetStringFromAddress(addr int) string {
	if addr >= src.CONST_STRING_START && addr < src.CONST_BOOL_START {
		return eCtx.ConstMemory.strings[addr-eCtx.ConstMemory.baseStringAddr]
	} else if addr >= src.GLOBAL_STRING_START && addr < src.GLOBAL_BOOL_START {
		return eCtx.GlobalMemory.strings[addr-eCtx.GlobalMemory.baseStringAddr]
	} else if addr >= src.TEMP_STRING_START && addr < src.TEMP_BOOL_START {
		if len(eCtx.StackSegment) > 0 {
			return eCtx.GetStackSegment().tempMemory.strings[addr-eCtx.GetStackSegment().tempMemory.baseStringAddr]
		}
		return eCtx.TempMemory.strings[addr-eCtx.TempMemory.baseStringAddr]
	}
	return eCtx.GetStackSegment().localMemory.strings[addr-eCtx.GetStackSegment().localMemory.baseStringAddr]
}
func (eCtx *ECtx) GetBoolFromAddress(addr int) bool {
	if addr >= src.CONST_BOOL_START && addr < src.CONST_BOOL_END {
		return eCtx.ConstMemory.bools[addr-eCtx.ConstMemory.baseBoolAddr]
	} else if addr >= src.GLOBAL_BOOL_START && addr < src.GLOBAL_BOOL_END {
		return eCtx.GlobalMemory.bools[addr-eCtx.GlobalMemory.baseBoolAddr]
	} else if addr >= src.TEMP_BOOL_START && addr < src.CONST_INT_START {
		if len(eCtx.StackSegment) > 0 {
			return eCtx.GetStackSegment().tempMemory.bools[addr-eCtx.GetStackSegment().tempMemory.baseBoolAddr]
		}
		return eCtx.TempMemory.bools[addr-eCtx.TempMemory.baseBoolAddr]
	}
	return eCtx.GetStackSegment().localMemory.bools[addr-eCtx.GetStackSegment().localMemory.baseBoolAddr]
}

func (eCtx *ECtx) StoreIntAtAddress(val int64, addr int) {
	if addr >= src.CONST_INT_START && addr < src.CONST_FLOAT_START {
		eCtx.ConstMemory.ints[addr-eCtx.ConstMemory.baseIntAddr] = val
	} else if addr >= src.GLOBAL_INT_START && addr < src.GLOBAL_FLOAT_START {
		eCtx.GlobalMemory.ints[addr-eCtx.GlobalMemory.baseIntAddr] = val
	} else if addr >= src.TEMP_INT_START && addr < src.TEMP_FLOAT_START {
		if len(eCtx.StackSegment) > 0 {
			eCtx.GetStackSegment().tempMemory.ints[addr-eCtx.GetStackSegment().tempMemory.baseIntAddr] = val
		}
		eCtx.TempMemory.ints[addr-eCtx.TempMemory.baseIntAddr] = val
	} else {
		eCtx.GetStackSegment().localMemory.ints[addr-eCtx.GetStackSegment().localMemory.baseIntAddr] = val
	}
}

func (eCtx *ECtx) StoreFloatAtAddress(val float64, addr int) {
	if addr >= src.CONST_FLOAT_START && addr < src.CONST_STRING_START {
		eCtx.ConstMemory.floats[addr-eCtx.ConstMemory.baseFloatAddr] = val
	} else if addr >= src.GLOBAL_FLOAT_START && addr < src.GLOBAL_STRING_START {
		eCtx.GlobalMemory.floats[addr-eCtx.GlobalMemory.baseFloatAddr] = val
	} else if addr >= src.TEMP_FLOAT_START && addr < src.TEMP_STRING_START {
		if len(eCtx.StackSegment) > 0 {
			eCtx.GetStackSegment().tempMemory.floats[addr-eCtx.GetStackSegment().tempMemory.baseFloatAddr] = val
		}
		eCtx.TempMemory.floats[addr-eCtx.TempMemory.baseFloatAddr] = val
	} else {
		eCtx.GetStackSegment().localMemory.floats[addr-eCtx.GetStackSegment().localMemory.baseFloatAddr] = val
	}
}

func (eCtx *ECtx) StoreStringAtAddress(val string, addr int) {
	if addr >= src.CONST_STRING_START && addr < src.CONST_BOOL_START {
		eCtx.ConstMemory.strings[addr-eCtx.ConstMemory.baseStringAddr] = val
	} else if addr >= src.GLOBAL_STRING_START && addr < src.GLOBAL_BOOL_START {
		eCtx.GlobalMemory.strings[addr-eCtx.GlobalMemory.baseStringAddr] = val
	} else if addr >= src.TEMP_STRING_START && addr < src.TEMP_BOOL_START {
		if len(eCtx.StackSegment) > 0 {
			eCtx.GetStackSegment().tempMemory.strings[addr-eCtx.GetStackSegment().tempMemory.baseStringAddr] = val
		}
		eCtx.TempMemory.strings[addr-eCtx.TempMemory.baseStringAddr] = val
	} else {
		eCtx.GetStackSegment().localMemory.strings[addr-eCtx.GetStackSegment().localMemory.baseStringAddr] = val
	}
}

func (eCtx *ECtx) StoreBoolAtAddress(val bool, addr int) {
	if addr >= src.CONST_BOOL_START && addr < src.CONST_BOOL_END {
		eCtx.ConstMemory.bools[addr-eCtx.ConstMemory.baseBoolAddr] = val
	} else if addr >= src.GLOBAL_BOOL_START && addr < src.GLOBAL_BOOL_END {
		eCtx.GlobalMemory.bools[addr-eCtx.GlobalMemory.baseBoolAddr] = val
	} else if addr >= src.TEMP_BOOL_START && addr < src.CONST_INT_START {
		if len(eCtx.StackSegment) > 0 {
			eCtx.GetStackSegment().tempMemory.bools[addr-eCtx.GetStackSegment().tempMemory.baseBoolAddr] = val
		}
		eCtx.TempMemory.bools[addr-eCtx.TempMemory.baseBoolAddr] = val
	} else {
		eCtx.GetStackSegment().localMemory.bools[addr-eCtx.GetStackSegment().localMemory.baseBoolAddr] = val
	}
}
