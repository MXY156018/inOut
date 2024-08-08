package utils_test

import (
	"mall-pkg/utils"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_PhoneEncode(t *testing.T) {
	asrt := assert.New(t)
	phone := utils.PhoneEncode("")
	asrt.Equal("", phone)

	phone = utils.PhoneEncode("1")
	asrt.Equal("1", phone)

	phone = utils.PhoneEncode("13")
	asrt.Equal("13", phone)

	phone = utils.PhoneEncode("138")
	asrt.Equal("138", phone)

	phone = utils.PhoneEncode("1382")
	asrt.Equal("1382", phone)

	phone = utils.PhoneEncode("13823")
	asrt.Equal("13823", phone)

	phone = utils.PhoneEncode("138237")
	asrt.Equal("138237", phone)

	phone = utils.PhoneEncode("1382371")
	asrt.Equal("1382371", phone)

	phone = utils.PhoneEncode("13823710")
	asrt.Equal("138*3710", phone)

	phone = utils.PhoneEncode("138237104")
	asrt.Equal("138**7104", phone)

	phone = utils.PhoneEncode("1382371041")
	asrt.Equal("138***1041", phone)

	phone = utils.PhoneEncode("13823710414")
	asrt.Equal("138****0414", phone)

	phone = utils.PhoneEncode("138237104147")
	asrt.Equal("138****4147", phone)
}