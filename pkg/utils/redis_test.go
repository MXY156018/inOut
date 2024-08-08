package utils_test

import (
	"testing"

	"mall-pkg/utils"
)

func Test_RedisFormatKey(t *testing.T) {
	key := utils.RedisFormatKey("asset", 10000)
	if key != "asset:10000" {
		t.Fatal()
	}
	key = utils.RedisFormatKey("10", 10000)
	if key != "10:10000" {
		t.Fatal()
	}
}
