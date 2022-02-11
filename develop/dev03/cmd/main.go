package main

import (
	"fmt"
	"strings"
	"os"

	"../pkg"
)

func main(){
	inputString := strings.Join(os.Args[1:], " ")
	inputString = strings.TrimSuffix(inputString, "\n")
	cmdFlags2 := pkg.NewCmdLine(inputString)
	allStrings := pkg.ReadFile(cmdFlags2.Files)
	mySorter := pkg.QuickSort{}
	sortObj := pkg.NewSortStrings(cmdFlags2, mySorter, allStrings)
	sortObj.Sort()
	for _, v := range sortObj.SortedStrings{
		fmt.Println(v)
	}
}