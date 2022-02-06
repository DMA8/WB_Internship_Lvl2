package pkg

import (
	"fmt"
	"strings"
	"io/ioutil"
	"log"
)

//Cuter stores CmdLine and inp line to cut
type Cuter struct{
	CmdLine		*CmdLine
	InpLines	[]string
	Output		[]string //для тестов
}

//NewCuter creates Cuter
func NewCuter(cmd *CmdLine) *Cuter{
	return &Cuter{
		CmdLine: cmd,
	}
}

//CutFiles cuts given files
func (c *Cuter)CutFiles(){
	for _, v := range c.CmdLine.Files {
		c.InpLines = readFile(v, false)
		c.cut()
	}
}

func(c *Cuter)cut(){
	for _, v := range c.InpLines{
		if c.CmdLine.OnlyStringsWithSep && !strings.Contains(v, c.CmdLine.Delimiter) {
			continue
		}
		splitted := strings.Split(v, c.CmdLine.Delimiter)
		if len(splitted) < c.CmdLine.Field {
			fmt.Println()
			//c.Output = append(c.Output, "\n") // для тестов
		} else {
			fmt.Println(splitted[c.CmdLine.Field - 1])
			c.Output = append(c.Output, splitted[c.CmdLine.Field - 1]) // для тестов
		}
	}
}

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
