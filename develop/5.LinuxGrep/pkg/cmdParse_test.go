package pkg

import (
	"testing"
)

func TestNewCmdLine(t *testing.T){
	test1 := `-A5 -B10 -ivF wordToFind text1 text2`
	correct1 := CmdLine{
		Files: make([]string, 0, 2),
		RawPattern:	"wordtofind",
		Pattern: "wordtofind",
		NoFlags: false,
		Flags: Flags{
			After: 5,
			Before: 10,
			Context: 0,
			Count: false,
			IgnoreCase: true,
			Invert: true,
			Fixed: true,
			LineNum: false,
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
