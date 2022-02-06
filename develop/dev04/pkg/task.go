package pkg

import (
	"strings"
)

//Finds anagram in given words, groups it in map[string][]string
func FindAnagram(words []string) map[string][]string{
	for i := 0; i < len(words); i++ {
		words[i] = strings.ToLower(words[i])
	}
	outMap := make(map[string][]string)
	alreadyAdded := make(map[string]int)
	for i:= 0; i < len(words); i++{
		if _, ok := alreadyAdded[words[i]]; !ok {
			for j := 0; j < len(words); j++ {
				if i == j {
					continue
				}
				if _, ok := alreadyAdded[words[j]]; !ok && words[i] != words[j] &&
				CheckAnagram([]rune(words[i]), []rune(words[j])) {
					outMap[words[i]] = append(outMap[words[i]], words[j])
					alreadyAdded[words[j]] = 1
				}
			}
			alreadyAdded[words[i]] = 1
		}
	}
	return outMap
}

//CheckAnagram checks wether given 2 words are anagrams
func CheckAnagram(word1, word2 []rune) bool {
	if len(word1) != len(word2){
		return false
	}
	word1Sorted, word2Sorted := make([]rune, len(word1)), make([]rune, len(word2))
	copy(word1Sorted, word1)
	copy(word2Sorted, word2)
	QsortRunes(word1Sorted, 0, len(word1Sorted) - 1)
	QsortRunes(word2Sorted, 0, len(word1Sorted) - 1)
	for i := 0; i < len(word1Sorted); i++{
		if word1Sorted[i] != word2Sorted[i] {
			return false
		}
	}
	return true
}

//QsortRunes sorts []rune
func QsortRunes(a []rune, left, rigth int) {
	if rigth - left > 0 {
		pivotInd := partitionRunes(a, left, rigth)
		QsortRunes(a, left, pivotInd - 1)
		QsortRunes(a, pivotInd + 1, rigth)
	}
}

func partitionRunes(a []rune, left, rigth int) int {
	var i, pivot, wall int
	pivot = rigth 
	wall = left 
	for i = left; i < rigth; i++ {
		if a[i] < a[pivot] {
			a[i], a[wall] = a[wall], a[i]
			wall++
		}
	}
	a[pivot], a[wall] = a[wall], a[pivot]
	return wall
}
