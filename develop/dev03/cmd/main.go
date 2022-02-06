package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"../pkg"
)

func main(){
	var inputString string
	scn := bufio.NewReader(os.Stdin)
	inputString, err := scn.ReadString('\n')
	if err != nil{
		log.Fatal(err)
	}
	inputString = strings.TrimSuffix(inputString, "\n")
	cmdFlags2 := pkg.NewCmdLine(inputString)
	allStrings := pkg.ReadFile(cmdFlags2.Files)
	mySorter := pkg.QuickSort{}
	sortObj := pkg.NewSortStrings(cmdFlags2, mySorter, allStrings)
	sortObj.Sort()
	fmt.Println(sortObj.SortedStrings)
}