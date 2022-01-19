package pkg

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

type CmdLine struct {
	Files	[]string
	NoFlags	bool
	Flags
}

type Flags struct {
	KeyColumnSort	bool
	PivotColumn		int
	NumberSort		bool
	ReversedOrder	bool
	UniqueLines		bool
}

func NewCmdLine(inpStr string) *CmdLine {
	cmd := CmdLine{}
	cmd.initCmdLine(inpStr)
	if !(cmd.Flags.KeyColumnSort || cmd.Flags.NumberSort ||
		cmd.Flags.ReversedOrder || cmd.Flags.UniqueLines) {
			cmd.NoFlags = true
		}
	return &cmd
}

func PrintCMD(c CmdLine){
	fmt.Println("files are ", c.Files)
	fmt.Println("-k ", c.Flags.KeyColumnSort)
	fmt.Println("numCol", c.Flags.PivotColumn)
	fmt.Println("-n ", c.Flags.NumberSort)
	fmt.Println("-r ", c.Flags.ReversedOrder)
	fmt.Println("-u ", c.Flags.UniqueLines)
}

func (c *CmdLine)initCmdLine(inpStr string) {
	operands := strings.Split(inpStr, " ")
	for _, v := range operands {
		if len(v) > 0 {
		if v[0] == '-' {
			SetFlags(c, inpStr)
		} else {
			c.Files = append(c.Files, v)
		}
	}
	}
}

func SetFlags(c *CmdLine, inpStr string) {
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
		case 'k':
			if !unicode.IsDigit(rune(arg[i + 1])) {
				log.Fatal("set column number for k flag")
			} else {
				c.Flags.KeyColumnSort = true
				i++
				startInd := i
				for i < len(arg) && arg[i] >= '0' && arg[i] <= '9' {
					i++
				}
				nCol, _ := strconv.Atoi(arg[startInd: i])
				c.Flags.PivotColumn = nCol
				i--
			}
		case 'n':
			c.Flags.NumberSort = true
		case 'r':
			c.Flags.ReversedOrder = true
		case 'u':
			c.Flags.UniqueLines = true
		}
	}
	}
}

