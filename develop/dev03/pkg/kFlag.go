package pkg

import (
	"fmt"
	"strconv"
	"strings"
)

// extracts cirtain column from multiple strings
// if there's no column - null value
func ExtractColumnMap(allStrings []string, nColumn int) map[string][]string{
	outMap := make(map[string][]string)
	for _, v := range allStrings {
		tmp := strings.Split(v, " ")
		if len(tmp) < nColumn {
			outMap[""] = append(outMap[""], v)
		} else {
			outMap[tmp[nColumn - 1]] = append(outMap[tmp[nColumn - 1]], v)
		}
	}
	return outMap
}

// extracts keys from map and sort it
func GetKeySortedCollection(M map[string][]string) []string{
	out := make([]string,0)
	for key, _ := range M{
		out = append(out, key)
	}
	QsortStrings(out, 0, len(out) - 1)
	return out
}

func IntGetKeySortedCollection(M map[string][]string) (map[int][]string, []int){ // отсортировать слайс внутри значений 
	out := make(map[int][]string)
	keys:= make([]int, 0)
	for key, value := range M{
		i := 0
		for i < len(key) && key[i] >= '0' && key[i] <= '9' {
			i++
		}
		num, _ := strconv.Atoi(key[ : i])
		out[num] = append(out[num], value...)
	}
	for key, _ := range out {
		keys = append(keys, key)
	}
	QsortInt(keys, 0, len(out) - 1)
	fmt.Println(keys)
	return out, keys
}
