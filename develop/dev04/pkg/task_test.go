package pkg

import "testing"

func TestFindAnagram(t *testing.T) {
	inpData1 := 
}


func TestCheckAnagram(t *testing.T){
	w1 := "листок"
	w1Anagrams := []string{"столик", "слиток"}
	for _, v := range w1Anagrams{
		if !checkAnagram([]rune(w1), []rune(v)){
			t.Error("Test1 Failed")
		}
	}
	w2 := "пятка"
	w2Anagrams := []string{"тяпка", "пятка"}
	for _, v := range w2Anagrams{
		if !checkAnagram([]rune(w2), []rune(v)){
			t.Error("Test2 Failed")
		}
	}
	w3 := "арбуз"
	w3Compare := "дыняя"
	if checkAnagram([]rune(w3), []rune(w3Compare)){
		t.Error("Test3 Failed")
	}
}

