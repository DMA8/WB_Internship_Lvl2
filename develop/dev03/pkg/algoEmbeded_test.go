package pkg

import (
	"testing"
)

func TestEmbededSort(t *testing.T){
	sorter := QuickSort{}

	testCase1 := []string{"yoba", "biba", "zoba", "aboba"}
	testCase1Result := []string{"aboba", "biba", "yoba", "zoba"}

	testCase2 := []int{9, 2, 3, 6, 1, 5, 4, 8, 7}
	testCase2Result := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	testCase3 := []float64{9, 2, 3, 6, 1, 5, 4, 8, 7}
	testCase3Result := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	////////////////////////////////////////////////
	sorter.Sort(testCase1)
	sorter.Sort(testCase2)
	sorter.Sort(testCase3)
	if !compareSliceStrings(testCase1, testCase1Result) {
		t.Error("Fail 1")
	}
	if !compareSliceInts(testCase2, testCase2Result) {
		t.Error("Fail 2")
	}
	if !compareSliceFloats(testCase3, testCase3Result) {
		t.Error("Fail 3")
	}
}


func compareSliceStrings(sl1, sl2 []string) bool{
	if len(sl1) != len(sl2){
		return false
	}
	for i, v := range sl1{
		if v != sl2[i]{
			return false
		}
	}
	return true
}

func compareSliceInts(sl1, sl2 []int) bool{
	if len(sl1) != len(sl2){
		return false
	}
	for i, v := range sl1{
		if v != sl2[i]{
			return false
		}
	}
	return true
}

func compareSliceFloats(sl1, sl2 []float64) bool{
	if len(sl1) != len(sl2){
		return false
	}
	for i, v := range sl1{
		if v != sl2[i]{
			return false
		}
	}
	return true
}

