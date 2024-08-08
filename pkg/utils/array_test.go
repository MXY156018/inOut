package utils_test

import (
	"mall-pkg/utils"
	"testing"
)

func Test_ArrayIndexOf1(t *testing.T) {
	var target int = 10
	var dest []int = []int{10, 20, 30}

	idx := utils.ArrayIndexOf(dest, target)
	if idx != 0 {
		t.Fatal()
	}

	idx = utils.ArrayIndexOf(nil, target)
	if idx != -1 {
		t.Fatal()
	}

	idx = utils.ArrayIndexOf(dest, 1)
	if idx != -1 {
		t.Fatal()
	}
}

func Test_ArrayIndexOfUint64(t *testing.T) {
	var target uint64 = 10
	var dest []uint64 = []uint64{10, 20, 30}

	idx := utils.ArrayIndexOfUint64(dest, target)
	if idx != 0 {
		t.Fatal()
	}

	idx = utils.ArrayIndexOfUint64(nil, target)
	if idx != -1 {
		t.Fatal()
	}

	idx = utils.ArrayIndexOfUint64(dest, 1)
	if idx != -1 {
		t.Fatal()
	}
}

func Test_ArrayIndexOfInt64(t *testing.T) {
	var target int64 = 10
	var dest []int64 = []int64{10, 20, 30}

	idx := utils.ArrayIndexOfInt64(dest, target)
	if idx != 0 {
		t.Fatal()
	}

	idx = utils.ArrayIndexOfInt64(nil, target)
	if idx != -1 {
		t.Fatal()
	}

	idx = utils.ArrayIndexOfInt64(dest, 1)
	if idx != -1 {
		t.Fatal()
	}
}

func Test_ArrayIndexOfInt(t *testing.T) {
	var target int = 10
	var dest []int = []int{10, 20, 30}

	idx := utils.ArrayIndexOfInt(dest, target)
	if idx != 0 {
		t.Fatal()
	}

	idx = utils.ArrayIndexOfInt(nil, target)
	if idx != -1 {
		t.Fatal()
	}

	idx = utils.ArrayIndexOfInt(dest, 1)
	if idx != -1 {
		t.Fatal()
	}
}

func Test_ArrayIndexOfUint(t *testing.T) {
	var target uint = 10
	var dest []uint = []uint{10, 20, 30}

	idx := utils.ArrayIndexOfUint(dest, target)
	if idx != 0 {
		t.Fatal()
	}

	idx = utils.ArrayIndexOfUint(nil, target)
	if idx != -1 {
		t.Fatal()
	}

	idx = utils.ArrayIndexOfUint(dest, 1)
	if idx != -1 {
		t.Fatal()
	}
}

func Test_ArrayIndexOfUint32(t *testing.T) {
	var target uint32 = 10
	var dest []uint32 = []uint32{10, 20, 30}

	idx := utils.ArrayIndexOfUint32(dest, target)
	if idx != 0 {
		t.Fatal()
	}

	idx = utils.ArrayIndexOfUint32(nil, target)
	if idx != -1 {
		t.Fatal()
	}

	idx = utils.ArrayIndexOfUint32(dest, 1)
	if idx != -1 {
		t.Fatal()
	}
}

func Test_ArrayIndexOfInt32(t *testing.T) {
	var target int32 = 10
	var dest []int32 = []int32{10, 20, 30}

	idx := utils.ArrayIndexOfInt32(dest, target)
	if idx != 0 {
		t.Fatal()
	}

	idx = utils.ArrayIndexOfInt32(nil, target)
	if idx != -1 {
		t.Fatal()
	}

	idx = utils.ArrayIndexOfInt32(dest, 1)
	if idx != -1 {
		t.Fatal()
	}
}

func Test_ArrayIndexOfInt16(t *testing.T) {
	var target int16 = 10
	var dest []int16 = []int16{10, 20, 30}

	idx := utils.ArrayIndexOfInt16(dest, target)
	if idx != 0 {
		t.Fatal()
	}

	idx = utils.ArrayIndexOfInt16(nil, target)
	if idx != -1 {
		t.Fatal()
	}

	idx = utils.ArrayIndexOfInt16(dest, 1)
	if idx != -1 {
		t.Fatal()
	}
}

func Test_ArrayIndexOfUint16(t *testing.T) {
	var target uint16 = 10
	var dest []uint16 = []uint16{10, 20, 30}

	idx := utils.ArrayIndexOfUint16(dest, target)
	if idx != 0 {
		t.Fatal()
	}

	idx = utils.ArrayIndexOfUint16(nil, target)
	if idx != -1 {
		t.Fatal()
	}

	idx = utils.ArrayIndexOfUint16(dest, 1)
	if idx != -1 {
		t.Fatal()
	}
}

func Test_ArrayIndexOfUint8(t *testing.T) {
	var target uint8 = 10
	var dest []uint8 = []uint8{10, 20, 30}

	idx := utils.ArrayIndexOfUint8(dest, target)
	if idx != 0 {
		t.Fatal()
	}

	idx = utils.ArrayIndexOfUint8(nil, target)
	if idx != -1 {
		t.Fatal()
	}

	idx = utils.ArrayIndexOfUint8(dest, 1)
	if idx != -1 {
		t.Fatal()
	}
}

func Test_ArrayIndexOfInt8(t *testing.T) {
	var target int8 = 10
	var dest []int8 = []int8{10, 20, 30}

	idx := utils.ArrayIndexOfInt8(dest, target)
	if idx != 0 {
		t.Fatal()
	}

	idx = utils.ArrayIndexOfInt8(nil, target)
	if idx != -1 {
		t.Fatal()
	}

	idx = utils.ArrayIndexOfInt8(dest, 1)
	if idx != -1 {
		t.Fatal()
	}
}

func Test_ArrayIndexOfString(t *testing.T) {
	var target string = "10"
	var dest []string = []string{"10", "20", "30"}

	idx := utils.ArrayIndexOfString(dest, target)
	if idx != 0 {
		t.Fatal()
	}

	idx = utils.ArrayIndexOfString(nil, target)
	if idx != -1 {
		t.Fatal()
	}

	idx = utils.ArrayIndexOfString(dest, "1")
	if idx != -1 {
		t.Fatal()
	}
}
