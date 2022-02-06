// parses inputCmd line and sets flags
// it works in context of grep
package pkg

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type CmdLine struct {
	Files		[]string
	RawPattern	string
	Pattern		string
	NoFlags		bool
	Flags
}

type Flags struct {
	After			int
	Before			int
	Context			int
	Count			bool
	IgnoreCase		bool
	Invert			bool
	Fixed			bool
	LineNum			bool
}

func NewCmdLine(inpStr string) *CmdLine {
	cmd := CmdLine{}
	cmd.initCmdLine(inpStr)
	blankFlag := Flags{}
	if cmd.Flags == blankFlag{
			cmd.NoFlags = true
		}
	return &cmd
}

func (c *CmdLine)initCmdLine(inpStr string) {
	operands := strings.Split(inpStr, " ")
	for _, v := range operands {
		if len(v) > 0 {
		if v[0] == '-' {
			c.SetFlags(inpStr)
		} else {
			if len(c.Pattern) == 0{
				c.Pattern = v
				out := ""
				for _, c := range v {
					if c == '\\' || c == '.' || c == '[' || c == ']' || c == '^' {
						out += "\\"
					} 
					out += string(c)
				}
				c.RawPattern = out
			} else {
				c.Files = append(c.Files, v)
			}
		}
	}
	}
}

func (c *CmdLine)SetFlags(inpStr string) {
	args := strings.Split(inpStr, " ")
	if len(args[0]) > 1 && args[0][0] != '-' {
		return
	}
	for _, arg := range args{
		if arg[0] != '-' {
			continue
		}
	for i := 0; i < len(arg); i++ {
		switch arg[i]{
		case 'A':
			SetNumeredFlag(arg, &c.Flags.After, &i)
		case 'B':
			SetNumeredFlag(arg, &c.Flags.Before, &i)
		case 'C':
			SetNumeredFlag(arg, &c.Flags.Context, &i)
		case 'c':
			c.Flags.Count = true
		case 'i':
			c.Flags.IgnoreCase = true
		case 'v':
			c.Flags.Invert = true
		case 'F':
			c.Flags.Fixed = true
		case 'n':
			c.Flags.LineNum = true
		}
	}
	}
}

func SetNumeredFlag(rawCmdString string, structField, outerInd *int){
	if !(rawCmdString[*outerInd + 1] >= '0' && rawCmdString[*outerInd + 1] <= '9') {
		log.Fatal("set column number for k flag")
	} else {
		*outerInd++
		startInd := *outerInd
		for *outerInd < len(rawCmdString) &&
		rawCmdString[*outerInd] >= '0' &&
		rawCmdString[*outerInd] <= '9' {
			*outerInd++
		}
		nCol, _ := strconv.Atoi(rawCmdString[startInd: *outerInd])
		*structField = nCol
		*outerInd--
	}
}

func GetNextLine(input *bufio.Reader) string{
	a, err := input.ReadString('\n') // Readstring изначально выделяет больший буффер, чем ReadLine
	if err != nil {
		log.Fatal(err)
	}
	return a
}

//bufio.Reader потому что Scanner не работает на больших строках
func CreateReader(fileName string) *bufio.Reader{
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	input := bufio.NewReader(file)
	return input
}