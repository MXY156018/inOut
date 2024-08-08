package utils

import "reflect"

// 查找目标在数组的下标
//
// 泛型有可能效率相对较低
func ArrayIndexOf(dest interface{}, target interface{}) int {
	if dest == nil {
		return -1
	}
	arr := reflect.ValueOf(dest)
	for i := 0; i < arr.Len(); i++ {
		v := arr.Index(i)
		if v.Interface() == target {
			return i
		}
	}
	return -1
}

func ArrayIndexOfUint64(dest []uint64, target uint64) int {
	if dest == nil {
		return -1
	}

	for i := 0; i < len(dest); i++ {
		if dest[i] == target {
			return i
		}
	}
	return -1
}

func ArrayIndexOfInt64(dest []int64, target int64) int {
	if dest == nil {
		return -1
	}

	for i := 0; i < len(dest); i++ {
		if dest[i] == target {
			return i
		}
	}
	return -1
}

func ArrayIndexOfInt(dest []int, target int) int {
	if dest == nil {
		return -1
	}

	for i := 0; i < len(dest); i++ {
		if dest[i] == target {
			return i
		}
	}
	return -1
}

func ArrayIndexOfUint(dest []uint, target uint) int {
	if dest == nil {
		return -1
	}

	for i := 0; i < len(dest); i++ {
		if dest[i] == target {
			return i
		}
	}
	return -1
}

func ArrayIndexOfUint32(dest []uint32, target uint32) int {
	if dest == nil {
		return -1
	}

	for i := 0; i < len(dest); i++ {
		if dest[i] == target {
			return i
		}
	}
	return -1
}

func ArrayIndexOfInt32(dest []int32, target int32) int {
	if dest == nil {
		return -1
	}

	for i := 0; i < len(dest); i++ {
		if dest[i] == target {
			return i
		}
	}
	return -1
}

func ArrayIndexOfUint16(dest []uint16, target uint16) int {
	if dest == nil {
		return -1
	}

	for i := 0; i < len(dest); i++ {
		if dest[i] == target {
			return i
		}
	}
	return -1
}

func ArrayIndexOfInt16(dest []int16, target int16) int {
	if dest == nil {
		return -1
	}

	for i := 0; i < len(dest); i++ {
		if dest[i] == target {
			return i
		}
	}
	return -1
}

func ArrayIndexOfUint8(dest []uint8, target uint8) int {
	if dest == nil {
		return -1
	}

	for i := 0; i < len(dest); i++ {
		if dest[i] == target {
			return i
		}
	}
	return -1
}

func ArrayIndexOfInt8(dest []int8, target int8) int {
	if dest == nil {
		return -1
	}

	for i := 0; i < len(dest); i++ {
		if dest[i] == target {
			return i
		}
	}
	return -1
}

func ArrayIndexOfString(dest []string, target string) int {
	if dest == nil {
		return -1
	}

	for i := 0; i < len(dest); i++ {
		if dest[i] == target {
			return i
		}
	}
	return -1
}
