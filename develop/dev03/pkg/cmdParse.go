package pkg

import (
	"log"
	"strconv"
	"strings"
	"unicode"
)

//CmdLine represents input flags and given files
type CmdLine struct {
	Files	[]string
	NoFlags	bool
	Flags
}

// Flags represents particular flags
type Flags struct {
	KeyColumnSort	bool
	PivotColumn		int
	NumberSort		bool
	ReversedOrder	bool
	UniqueLines		bool
	MonthSort		bool
	SuffixSort		bool
	CheckSort		bool
}

//NewCmdLine is a constructor for CmdLine struct
func NewCmdLine(inpStr string) *CmdLine {
	cmd := CmdLine{}
	cmd.initCmdLine(inpStr)
	if !(cmd.Flags.KeyColumnSort || cmd.Flags.NumberSort ||
		cmd.Flags.ReversedOrder || cmd.Flags.UniqueLines) {
			cmd.NoFlags = true
	}
	if cmd.Flags.PivotColumn == 0 {
		cmd.Flags.PivotColumn = 1
	}
	return &cmd
}

//initCmdLine inits parsing cmdLine
func (c *CmdLine)initCmdLine(inpStr string) {
	operands := strings.Split(inpStr, " ")
	for _, v := range operands {
		if len(v) > 0 {
			if v[0] == '-' {
				setFlags(c, inpStr)
			} else {
				c.Files = append(c.Files, v)
			}
		}
	}
}

// setFlags sets flags to CmdLine struct
func setFlags(c *CmdLine, inpStr string) {
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
		case 'c':
			c.Flags.CheckSort = true
		case 'M':
			c.Flags.MonthSort = true
		case 'h':
			c.Flags.SuffixSort = true
		}
	}
	}
	if c.Flags.PivotColumn == 0{
		c.Flags.PivotColumn = 1
	}
}

