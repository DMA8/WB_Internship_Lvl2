package pkg

import (
	"testing"
)
func TestNewCmdLine(t *testing.T){
	test1 := "-k10Murh text1 text2"
	correct1 := CmdLine{
		Files: make([]string, 0, 2),
		NoFlags: false,
		Flags: Flags{
			KeyColumnSort: 	true,
			PivotColumn: 	10,
			NumberSort:		false,
			ReversedOrder:	true,
			UniqueLines:	true,
			MonthSort:		true,
			SuffixSort:		true,
			CheckSort:		false,
		},
	}
	correct1.Files = append(correct1.Files, "text1", "text2")
	res1 := NewCmdLine(test1)
	if res1.Flags != correct1.Flags{
		t.Error("Failed1")
	}
	if !compareSliceStrings(res1.Files, correct1.Files){
		t.Error("Failed1 files problem!")
	}
}