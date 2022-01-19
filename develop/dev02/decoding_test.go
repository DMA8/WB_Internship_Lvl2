package main

import (
	"fmt"
	"testing"
)

func TestGetCloseLength(t *testing.T) {
	test1 := "a4bc2d5e"
	if GetCloseLength([]rune(test1)) != 13 {
		t.Error("test 1 is not passed")
	}
	test2 := "a"
	if GetCloseLength([]rune(test2)) != 1 {
		t.Error("test 2 is not passed")
	}
	test3 := `qwe\4\5`
	if GetCloseLength([]rune(test3)) != 5 {
		t.Error("test 3 is not passed")
	}
	test4 := `qwe\45`
	if GetCloseLength([]rune(test4)) != 8 {
		t.Error("test 4 is not passed")
	}
}

func TestUnpack(t *testing.T) {
	// test1 := "a4bc2d5e"
	// if val, _ := Unpack(test1); val != "aaaabccddddde" {
	// 	fmt.Println(val)
	// 	t.Error("TestUnpack 1 is not passed")
	// }
	test2 := "abcd"
	if val, _ := Unpack(test2); val != "abcd" {
		t.Error("TestUnpack 2 is not passed")
	}
	test3 := "45"
	if val, err := Unpack(test3); val != "" || err == nil {
		t.Error("TestUnpack 3 is not passed")
	}
	test4 := ""
	if val, _ := Unpack(test4); val != "" {
		t.Error("TestUnpack 4 is not passed")
	}
	test5 := `qwe\4\5`
	if val, err := Unpack(test5); val != "qwe45" && err != nil {
		t.Error("TestUnpack 5 is not passed")
	}
	test6 := `qwe\45`
	if val, err := Unpack(test6); val != "qwe44444" && err != nil {
		t.Error("TestUnpack 6 is not passed")
	}
	test7 := `qwe\\5`
	if val, err := Unpack(test7); val != `qwe\\\\\` && err != nil {
		t.Error("TestUnpack 7 is not passed")
	}
}
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

func TestEscapeAtoi(t *testing.T) {
	ind := 0
	test1 := `\45`
	if EscapeAtoi([]rune(test1), &ind) != 5 {
		t.Error("TestEscapeAtoi 1 is not passed")
	}
	test2 := `\4`
	if EscapeAtoi([]rune(test2), &ind) != 1 {
		t.Error("TestEscapeAtoi 2 is not passed")
	}
}

func TestAtoiStable(t *testing.T) {
	ind := 0
	test1 := `45`
	if val := AtoiStable([]rune(test1), &ind); val != 45 {
		t.Error("TestAtoiStable 1 is not passed")
	}
	test2 := `4`
	if val := AtoiStable([]rune(test2), &ind); val != 4 {
		t.Error("TestAtoiStable 2 is not passed")
	}
	test3 := `24asd`
	if val := AtoiStable([]rune(test3), &ind); val != 24 {
		t.Error("TestAtoiStable 3 is not passed")
	}
}

// func TestMultiplyRune(t *testing.T) {
// 	outInd := 0
// 	test1 := "a2"
// 	outRunes := make([]rune, 100)
// 	if res := MultiplyRune([]rune(test1), outRunes, &)
// }
