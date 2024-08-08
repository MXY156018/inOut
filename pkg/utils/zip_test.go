package utils_test

import (
	"bytes"
	"encoding/base64"
	"mall-pkg/utils"
	"testing"
)

func Test_Zip_1(t *testing.T) {
	var data = "111245344444444444"
	for i := 0; i < 100; i++ {
		data += "dmqmdqklwiri2322354254543543534"
	}
	b := utils.Zip([]byte(data))
	t.Logf("after zip, length = %d\n", b.Len())
	dec, err := utils.Unzip(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(dec) != data {
		t.Fatalf("need %s, get %s", data, string(dec))
	}
}

func Test_Zip_2(t *testing.T) {
	var data = "111245344444444444dmqmdqklwiri2322354254543543534n.,l"
	t.Logf("%s", base64.StdEncoding.EncodeToString([]byte(data)))
	b := utils.Zip([]byte(data))
	t.Logf("before zip length = %d, after zip, length = %d\n", len(data), b.Len())
	b64 := base64.StdEncoding.EncodeToString(b.Bytes())
	bdec, _ := base64.StdEncoding.DecodeString(b64)
	b2 := bytes.NewBuffer(bdec)
	dec, err := utils.Unzip(b2)
	if err != nil {
		t.Fatal(err)
	}
	if string(dec) != data {
		t.Fatalf("need %s, get %s", data, string(dec))
	}
}
