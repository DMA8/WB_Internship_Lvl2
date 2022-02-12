package pkg

import (
	"testing"
)

func TestCheckFileNameValid(t *testing.T) {
	if checkFileNameValid("!aw") {
		t.Error("Test1 Failed")
	}
	if !checkFileNameValid("goodName") {
		t.Error("Test2 Failed")
	}
	if !checkFileNameValid("index.html") {
		t.Error("Test3 Failed")
	}
	if !checkFileNameValid(".trew") {
		t.Error("Test4 Failed")
	}
}

func TestGetFileName(t *testing.T) {
	prefix := "/mnt/c/Users/DB/OneDrive/html/"
	
	URL1 := "https://www.google.com/"
	f1 := "index.html"
	t1 := getFileName(URL1)
	if prefix+f1 != t1 {
		t.Error("Test1 failed")
	}

	URL2 := "https://www.google.com/about"
	f2 := "about.html"
	t2 := getFileName(URL2)
	if prefix+f2 != t2 {
		t.Error("Test2 failed")
	}

	URL3 := "https://www.google.com/about/g.png"
	f3 := "g.png"
	t3 := getFileName(URL3)
	if prefix + f3 != t3 {
		t.Error("Test3 failed")
	}
}

func TestCheckAvailability(t *testing.T){
	if checkAvailabilyty("21345"){
		t.Error("Test1 failed")
	}
	if checkAvailabilyty("google.com"){
		t.Error("Test2 failed")
	}
	if !checkAvailabilyty("https://www.google.com/"){
		t.Error("Test3 failed")
	}
}

func TestHasSrcInHTMLLine(t *testing.T){
	
	line1 := `<img src="https://wiki.org/pics/redhat.png/>"`
	urls1 := []string{"https://wiki.org/pics/redhat.png"}
	urlsFound := getAllSrcInHTMLLine(line1, "www.wiki.org")
	if !compareSliceStrings(urls1, urlsFound) {
		t.Error("Test1 failed")
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
