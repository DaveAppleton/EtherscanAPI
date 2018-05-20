package etherscanAPI

import (
	"fmt"
	"testing"
)

func TestGetSource(t *testing.T) {

	code, err := GetSourceCode("0xba2184520A1cC49a6159c57e61E1844E085615B6", "")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	fmt.Println(code[:6])
	if code[:6] != "pragma" {
		t.Error(err)
		fmt.Println("[" + code + "]")
		t.Fail()
	}
}

func TestABI(t *testing.T) {
	code, err := GetABI("0xba2184520A1cC49a6159c57e61E1844E085615B6", "")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	fmt.Println(code[:6])
	if code[:1] != "[" {
		t.Error(err)
		fmt.Println("[" + code + "]")
		t.Fail()
	}
}

func TestVersion(t *testing.T) {
	ver, err := GetVersion("0xba2184520A1cC49a6159c57e61E1844E085615B6", "")
	fmt.Println(ver)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	if len(ver) < 5 {
		t.Error("Wrong Result : " + ver)
		t.Fail()
	}
}

func TestOptimisation(t *testing.T) {
	opt, err := GetOptimisation("0xba2184520A1cC49a6159c57e61E1844E085615B6", "")
	fmt.Println(opt)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	if !opt {
		t.Error("Wrong Result : false")
		t.Fail()

	}
}

func TestSwarm(t *testing.T) {
	ver, err := GetSWARM("0xba2184520A1cC49a6159c57e61E1844E085615B6", "")
	fmt.Println(ver)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	if len(ver) < 5 {
		t.Error("Wrong Result : " + ver)
		t.Fail()
	}
}
