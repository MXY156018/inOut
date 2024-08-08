package utils_test

import (
	"mall-pkg/utils"
	"testing"
)

func Test_ParseByte(t *testing.T) {
	var need int64
	//////
	need = 1024
	v, err := utils.ParseByte("1k")
	if err != nil {
		t.Fatal(err)
	}
	if v != need {
		t.Fatalf("need %d get %d", need, v)
	}
	v, err = utils.ParseByte("1 k")
	if err != nil {
		t.Fatal(err)
	}
	if v != need {
		t.Fatalf("need %d get %d", need, v)
	}

	v, err = utils.ParseByte(" 1 k  ")
	if err != nil {
		t.Fatal(err)
	}
	if v != need {
		t.Fatalf("need %d get %d", need, v)
	}

	v, err = utils.ParseByte(" 1024  ")
	if err != nil {
		t.Fatal(err)
	}
	if v != need {
		t.Fatalf("need %d get %d", need, v)
	}

	v, err = utils.ParseByte(" 0x400  ")
	if err != nil {
		t.Fatal(err)
	}
	if v != need {
		t.Fatalf("need %d get %d", need, v)
	}

	////
	need = 1024 * 1024
	v, err = utils.ParseByte("1m")
	if err != nil {
		t.Fatal(err)
	}
	if v != need {
		t.Fatalf("need %d get %d", need, v)
	}

	////
	need = 1024 * 1024 * 1024
	v, err = utils.ParseByte("1g")
	if err != nil {
		t.Fatal(err)
	}
	if v != need {
		t.Fatalf("need %d get %d", need, v)
	}

	////
	need = 1024 * 1024 * 1024 * 1024
	v, err = utils.ParseByte("1t")
	if err != nil {
		t.Fatal(err)
	}
	if v != need {
		t.Fatalf("need %d get %d", need, v)
	}
}
