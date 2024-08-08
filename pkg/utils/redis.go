package utils

import "fmt"

// 格式化 redis Key
//
// 以 : 对字段进行分隔
//
// 参数 基本类型，不支持结构体，如果传入 对象类型，结果可能并不是你想要的
//
// 示例
//
// RedisFormatKey("asset", 10000)
func RedisFormatKey(args ...interface{}) string {
	key := ""
	for i := 0; i < len(args); i++ {
		if i > 0 {
			key += ":"
		}
		key += fmt.Sprintf("%v", args[i])
	}
	return key
}
