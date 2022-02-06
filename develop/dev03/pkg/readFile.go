package pkg

import (
	"io/ioutil"
	"log"
	"strings"
)

//ReadFile reads given files and returns lines
func ReadFile(names []string) []string{
	allLines := make([]string, 0)
	for _, v := range names {
		i, err := ioutil.ReadFile(v)
		if err != nil {
			log.Fatal(err)
		}
		allLines = append(allLines, strings.Split(string(i), "\n")...)
	}
	return allLines
}
