package main

import (
	"os"
	"strings"

	"../pkg"
)

func main() {
	inputString := strings.Join(os.Args[1:], "`")
	cmdLine := pkg.NewCmdLine(inputString)
	myCut := pkg.NewCuter(cmdLine)
	myCut.CutFiles()
}
