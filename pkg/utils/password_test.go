package utils_test

import (
	"mall-pkg/utils"
	"testing"
)

func Test_Password(t *testing.T) {
	pwd := utils.BcryptHash("123456")
	t.Log(pwd)
	isOk := utils.BcryptCheck("123456", pwd)
	if !isOk {
		t.Fail()
	}
}