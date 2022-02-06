package pkg

import (
	"testing"
)

func TestSort(t *testing.T){
	inputString := "../texts/suff"
	cmdFlags := NewCmdLine(inputString)
	sortAlgo := QuickSort{}
	mainObj := NewSortStrings(cmdFlags, sortAlgo, nil)
	
	testCase1 := []string{
		"5 Пятая строка",
		"3 Третья строка",
		"1 Строка",
	}
	testCase1Ans := []string{
		"1 Строка",
		"3 Третья строка",
		"5 Пятая строка",
	}
	
	mainObj.InpLines = testCase1
	mainObj.Sort()
	if !compareSliceStrings(mainObj.SortedStrings, testCase1Ans){
		t.Error("Test 1 Failed")
	}
	inputString = "-r ../texts/suff"
	cmdFlags = NewCmdLine(inputString)
	mainObj.Flags = cmdFlags
	testCase2Ans := []string{
		"5 Пятая строка",
		"3 Третья строка",
		"1 Строка",
	}
	mainObj.SortedStrings = nil
	mainObj.Sort()
	if !compareSliceStrings(mainObj.SortedStrings, testCase2Ans){
		t.Error("Test 1 Failed")
	}
}

func TestSafeFloatConvert(t *testing.T){
	testCase1 := "532.1"
	testCase2 := "532.12.kgnkrng3223"
	testCase3 := "skjdfnejfkn"
	testCase4 := "54K"
	testCase5 := "12G"
	t1Ans := 532.1
	t2Ans := 532.12
	t3Ans := 0.0
	t4Ans := 54.0
	t5Ans := 12.0

	if a := safeFloatConvert(testCase1); a != t1Ans{
		t.Error("Fail1")
	}
	if a := safeFloatConvert(testCase2); a != t2Ans{
		t.Error("Fail2")
	}
	if a := safeFloatConvert(testCase3); a != t3Ans{
		t.Error("Fail3")
	}
	if a := safeFloatConvert(testCase4); a != t4Ans{
		t.Error("Fail4")
	}
	if a := safeFloatConvert(testCase5); a != t5Ans{
		t.Error("Fail3")
	}
}
