package pkg

import "testing"

func TestCheckAnagram(t *testing.T){
	w1 := "листок"
	w1Anagrams := []string{"столик", "слиток"}
	for _, v := range w1Anagrams{
		if !CheckAnagram([]rune(w1), []rune(v)){
			t.Error("Test1 Failed")
		}
	}
	w2 := "пятка"
	w2Anagrams := []string{"тяпка", "пятка"}
	for _, v := range w2Anagrams{
		if !CheckAnagram([]rune(w2), []rune(v)){
			t.Error("Test2 Failed")
		}
	}
	w3 := "арбуз"
	w3Compare := "дыняя"
	if CheckAnagram([]rune(w3), []rune(w3Compare)){
		t.Error("Test3 Failed")
	}
}