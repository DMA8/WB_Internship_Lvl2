package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"./pkg"
	//"os"
	"strings"
)

// уникальные строки - сортируем, печатаем строку, хэш строки записываем в словарь
func main() {
	//inputString1 := "-k1n -r -u texts/1e.txt texts/an.txt texts/ex.txt"
	inputString2 := "-k1u texts/noNums.txt texts/1e.txt texts/ex.txt"
	cmdFlags2 := pkg.NewCmdLine(inputString2)
	allStrings2 := readFile(cmdFlags2.Files)
//	SortEngine(*cmdFlags, allStrings)
	SortEngine(*cmdFlags2, allStrings2)

}

func SortEngine(cmds pkg.CmdLine, allStrings []string) {
	var (
		KeysSortedInt []int
		KeysSortedStr []string
		Mapa map[string][]string
		MapSorted map[int][]string
	)
	if cmds.NoFlags{
		pkg.QsortStrings(allStrings, 0, len(allStrings) - 1)
		for _, v := range allStrings {
			fmt.Println(v)
		}
	} else if cmds.Flags.NumberSort {
		if cmds.Flags.PivotColumn == 0 {
			cmds.Flags.PivotColumn = 1
		}
		mapa := pkg.ExtractColumnMap(allStrings, cmds.Flags.PivotColumn)
		MapSorted, KeysSortedInt = pkg.IntGetKeySortedCollection(mapa)
		for _, key := range KeysSortedInt {
			if len(MapSorted[key]) > 1{
				pkg.QsortStrings(MapSorted[key], 0, len(MapSorted[key]) - 1)
			}
		}
	} else{
		Mapa = pkg.ExtractColumnMap(allStrings, cmds.PivotColumn)
		KeysSortedStr = pkg.GetKeySortedCollection(Mapa)
		for _, key := range KeysSortedStr {
			if key == "" {
				pkg.QsortStrings(Mapa[key], 0, len(Mapa[key]) - 1)
			}
		}
	}
	if cmds.Flags.UniqueLines && !cmds.Flags.ReversedOrder { //-u
		printedLines := make(map[string]int)
		if cmds.Flags.NumberSort{
			for _, key := range KeysSortedInt{
				if len(MapSorted[key]) > 1{
					for _, str := range MapSorted[key] {
						if _, ok := printedLines[str]; !ok {
							printedLines[str] = 1
							fmt.Println(str)
						}
					}
				} else {
					if _, ok := printedLines[MapSorted[key][0]]; !ok {
						printedLines[MapSorted[key][0]] = 1
						fmt.Println(MapSorted[key][0])
					}
				}
			}
		} else {
			for _, key := range KeysSortedStr{
				if len(Mapa[key]) > 1{
					for _, str := range Mapa[key] {
						if _, ok := printedLines[str]; !ok {
							printedLines[str] = 1
							fmt.Println(str)
						}
					}
				} else {
					if _, ok := printedLines[Mapa[key][0]]; !ok {
						printedLines[Mapa[key][0]] = 1
						fmt.Println(Mapa[key][0])
					}
				}
			}
		}
	} else if cmds.Flags.UniqueLines && cmds.Flags.ReversedOrder { //-ur
		printedLines := make(map[string]int)
		if cmds.Flags.NumberSort{
			for i := len(KeysSortedInt) - 1; i > 0; i--{
				if len(MapSorted[KeysSortedInt[i]]) > 1{
					for j := len(MapSorted[KeysSortedInt[i]]) - 1; j > 0; j-- {
						if _, ok := printedLines[MapSorted[KeysSortedInt[i]][j]]; !ok {
							printedLines[MapSorted[KeysSortedInt[i]][j]] = 1
							fmt.Println(MapSorted[KeysSortedInt[i]][j])
						}
					}
				} else {
					if _, ok := printedLines[MapSorted[KeysSortedInt[i]][0]]; !ok {
						printedLines[MapSorted[KeysSortedInt[i]][0]] = 1
						fmt.Println(MapSorted[KeysSortedInt[i]][0])
					}
				}
			}
		} else {
			for i := len(KeysSortedStr) - 1; i > 0; i--{
				if len(Mapa[KeysSortedStr[i]]) > 1{
					for j := len(Mapa[KeysSortedStr[i]]) - 1; j > 0 ; j-- {
						if _, ok := printedLines[Mapa[KeysSortedStr[i]][j]]; !ok {
							printedLines[Mapa[KeysSortedStr[i]][j]] = 1
							fmt.Println(Mapa[KeysSortedStr[i]][j])
						}
					}
				} else {
					if _, ok := printedLines[Mapa[KeysSortedStr[i]][0]]; !ok {
						printedLines[Mapa[KeysSortedStr[i]][0]] = 1
						fmt.Println(Mapa[KeysSortedStr[i]][0])
					}
				}
			}
		} 
	} else if !cmds.Flags.UniqueLines && cmds.Flags.ReversedOrder { //-r
		if cmds.Flags.NumberSort{
			for i := len(KeysSortedInt) - 1; i > 0; i--{
				for j := len(MapSorted[KeysSortedInt[i]]) - 1; j > 0; j-- {
						fmt.Println(MapSorted[KeysSortedInt[i]][j])
				}
			}
		} else {
			for i := len(KeysSortedStr) - 1; i > 0; i--{
				for j := len(Mapa[KeysSortedStr[i]]) - 1; j > 0 ; j-- {
						fmt.Println(Mapa[KeysSortedStr[i]][j])
				}
			}
		} 
	} else if !cmds.Flags.UniqueLines && !cmds.Flags.ReversedOrder {
		if cmds.Flags.NumberSort{
			for i := 0; i < len(KeysSortedInt); i++{
				for j := 0; j < len(MapSorted[KeysSortedInt[i]]); j++ {
						fmt.Println(MapSorted[KeysSortedInt[i]][j])
				}
			}
		} else {
			for i := 0; i < len(KeysSortedStr); i++{
				for j := 0; j < len(Mapa[KeysSortedStr[i]]); j++ {
						fmt.Println(Mapa[KeysSortedStr[i]][j])
				}
			}
		} 
	}
}

func readFile(names []string) []string{
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



