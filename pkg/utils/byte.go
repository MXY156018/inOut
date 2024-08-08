package utils

import (
	"strconv"
	"strings"
)

// 解析字节
//
//str 字符串 如 1k 1m 1g 0x1000
//
//使用 utils.ParseByte("1k")
func ParseByte(str string) (int64, error) {
	str = strings.Trim(str, " ")
	str = strings.ToLower(str)
	var scale int64 = 1
	var base int = 10
	var units = map[string]int64{
		"k": 1024,
		"m": 1024 * 1024,
		"g": 1024 * 1024 * 1024,
		"t": 1024 * 1024 * 1024 * 1024,
	}
	if strings.Index(str, "0x") == 0 {
		base = 16
		str = strings.TrimLeft(str, "0x")
	}
	for k, v := range units {
		if strings.HasSuffix(str, k) {
			scale = v
			str = strings.TrimRight(str, k)
			break
		}
	}
	str = strings.Trim(str, " ")
	value, err := strconv.ParseInt(str, base, 64)
	if err != nil {
		return 0, err
	}
	return value * scale, nil
}
