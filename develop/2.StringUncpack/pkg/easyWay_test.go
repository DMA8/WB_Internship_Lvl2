package pkg

import (
	"testing"
	"fmt"
)
func TestEasyUnpack(t *testing.T) {
	test1 := "a4bc2d5e"
	if val, _ := EasyUnpack(test1); val != "aaaabccddddde" {
		fmt.Println(val)
		t.Error("TestEasyUnpack 1 is not passed")
	}
	test2 := "abcd"
	if val, _ := EasyUnpack(test2); val != "abcd" {
		t.Error("TestEasyUnpack 2 is not passed")
	}
	test3 := "45"
	if val, err := EasyUnpack(test3); val != "" || err == nil {
		t.Error("TestEasyUnpack 3 is not passed")
	}
	test4 := ""
	if val, _ := EasyUnpack(test4); val != "" {
		t.Error("TestEasyUnpack 4 is not passed")
	}
	test5 := `qwe\4\5`
	if val, err := EasyUnpack(test5); val != "qwe45" && err != nil {
		t.Error("TestEasyUnpack 5 is not passed")
	}
	test6 := `qwe\45`
	if val, err := EasyUnpack(test6); val != "qwe44444" && err != nil {
		t.Error("TestEasyUnpack 6 is not passed")
	}
	test7 := `qwe\\5`
	if val, err := EasyUnpack(test7); val != `qwe\\\\\` && err != nil {
		t.Error("TestEasyUnpack 7 is not passed")
	}
}