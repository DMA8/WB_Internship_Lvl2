package pkg

import (
	"log"
	"strconv"
	"strings"
)

//CmdLine stores flags and files to cut
type CmdLine struct {
	Files		[]string
	Flags
}

//Flags stores flags
type Flags struct {
	Field				int
	Delimiter			string
	OnlyStringsWithSep	bool
}

//NewCmdLine creates new CmdLine obj
func NewCmdLine(inpStr string) *CmdLine {
	cmd := CmdLine{}
	cmd.initCmdLine(inpStr)
	if cmd.Delimiter == ""{
		cmd.Delimiter = "	"
	}
	return &cmd
}

func (c *CmdLine)initCmdLine(inpStr string) {
	operands := strings.Split(inpStr, " ")
	for _, v := range operands {
		if v == ""{
			continue
		}
		if v[0] == '-' {
			c.setFlags(inpStr)
		} else if v[0] != '\t' && v[0] != ' ' && v[0] != '"' && v[0] != '\''{
			c.Files = append(c.Files, v)
		}
	}
}

func (c *CmdLine)setFlags(inpStr string) {
	args := strings.Split(inpStr, " ")
	if len(args[0]) > 1 && args[0][0] != '-' {
		return
	}
	for _, arg := range args{
		if arg == ""{
			continue
		}
		if arg[0] != '-' {
			continue
		}
		for i := 0; i < len(arg); i++ {
			switch arg[i]{
			case 'f':
				setNumeredFlag(arg, &c.Flags.Field, &i)
			case 'd':
				c.catchDelimiter(inpStr)
			case 's':
				c.Flags.OnlyStringsWithSep = true
			}
		}
	}
}

func setNumeredFlag(rawCmdString string, structField, outerInd *int){
	if !(rawCmdString[*outerInd + 1] >= '0' && rawCmdString[*outerInd + 1] <= '9') {
		log.Fatal("Error flag f")
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

func (c *CmdLine)catchDelimiter(inpStr string){ // strings.Fields
	delim := ""
	for i := 0; i < len(inpStr) - 2; i++ {
		if inpStr[i:i+2] == "-d"{
			quote := inpStr[i + 2]
			for j := i + 3; j < len(inpStr); j++ {
				if inpStr[j] == quote{
					c.Delimiter = delim
					return
				} 
				delim += string(inpStr[j])
			}
		}
	}
}