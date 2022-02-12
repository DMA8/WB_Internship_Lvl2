package main

import (
	"fmt"

	"../pkg"
)

func main() {
	words := []string{
		"осел",
		"пятак",
		"тяпка",
		"листоК",
		"пятка",
		"тяпка",
		"пятка",
		"баклажан",
		"дом",
		"СЕЛО",
		"СтоЛик",
		"Слиток",
		"Слиток",
	}
	a := pkg.FindAnagram(words)
	fmt.Println(a)
}
