package pkg

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

//Greper is a struct for seaching keywords in files
type Greper struct{
	Flags			*CmdLine // Список файлов и флаги
	InpLines		[]string
	FoundLines		[]int
	RegEngie		*regexp.Regexp
	FileToReadInd	int
	AllFilesDone	bool
	Output			[]string //сохраняем найденные строки (для тестов)
}

//NewGreper is constructor of Greper
func NewGreper(cmd *CmdLine)*Greper{
	g := &Greper{
		Flags: cmd,
		InpLines: nil,
		FoundLines: nil,
		FileToReadInd: 0,
		AllFilesDone: false,
		Output: make([]string, 0),
	}
	g.initRegexCompile()
	return g
}

//Main executes search with given flags in given files
func (g *Greper)Main(){
	for i:= 0; i < len(g.Flags.Files); i++{
		g.readFile()
		g.findIndexMatchedStrings()
		g.shwoStrings()
		g.FoundLines = []int{}
	}
}

func (g *Greper)shwoStrings(){
	if g.Flags.Count{
		fmt.Println(len(g.FoundLines))
		return
	}
	alreadyPrinted := make(map[int]bool)
	for _, v := range g.FoundLines{
		for i := v - g.Flags.Before; i <= v; i++ {
			if i < 0 {
				i = 0
			}
			alreadyPrinted[i] = true
		}
		for i := v + g.Flags.After; i > v; i-- {
			if i > len(g.InpLines) - 1{
				i = len(g.InpLines) - 1
			}
			alreadyPrinted[i] = true
		}
	}
	sortedKeys := make([]int, 0, len(alreadyPrinted))
	for k := range alreadyPrinted{
		sortedKeys = append(sortedKeys, k)
	}
	sort.Ints(sortedKeys)
	if !g.Flags.Invert{
		for _, v := range sortedKeys{
			g.printLine(v, g.InpLines[v])
		}
	}else{
		for i, v := range g.InpLines{
			if ok := alreadyPrinted[i]; ok{
				continue
			} else {
				g.printLine(i, v)
			}
		}
	}

}

func (g *Greper)printLine(ind int, str string){
	outLine := ""
	if len(g.Flags.Files) > 1 {
		outLine += g.Flags.Files[g.FileToReadInd - 1]
	}
	if g.Flags.LineNum{
		outLine += ":" + strconv.Itoa(ind + 1) + ": "
	}
	fmt.Printf("%s%s\n", outLine, str)
	g.Output = append(g.Output, outLine + str) // for tests
}

func (g *Greper)findIndexMatchedStrings(){
	if g.Flags.Fixed{
		for i, v := range g.InpLines{
			if strings.Contains(v, g.Flags.Pattern){
				g.FoundLines = append(g.FoundLines, i)
			}
		}
	} else {
		for i, v := range g.InpLines{
			match := g.RegEngie.Match([]byte(v))
			if match{
				g.FoundLines = append(g.FoundLines, i)
			}
		}
	}
}

func (g *Greper)initRegexCompile(){
	var regEx string
	if g.Flags.Fixed {
		regEx = g.Flags.RawPattern
	} else {
		regEx = g.Flags.Pattern
	}
	pattern, err := regexp.Compile(regEx)
	if err != nil {
		log.Fatal(err)
	}
	g.RegEngie = pattern
}

func (g *Greper)readFile()bool{
	if g.FileToReadInd >= len(g.Flags.Files){
		g.AllFilesDone = true
		return false
	}
	g.InpLines = readFile(g.Flags.Files[g.FileToReadInd], g.Flags.IgnoreCase)
	g.FileToReadInd++
	return true
}
