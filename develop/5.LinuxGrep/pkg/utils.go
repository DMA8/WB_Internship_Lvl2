package pkg

import (
	"io/ioutil"
	"strings"
	"log"
)

func readFile(name string, toLower bool) []string{
	allLines := make([]string, 0)
	i, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}
	allLines = append(allLines, strings.Split(string(i), "\n")...)
	if toLower{
		for i := range allLines{
			allLines[i] = strings.ToLower(allLines[i])
		}
	}
	return allLines
}
