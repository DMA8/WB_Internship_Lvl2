package main

import (
	"../pkg"
)

func main() {
	inpStr := "-f3 -s -d' ' ../texts/t"
	cmdLine := pkg.NewCmdLine(inpStr)
	myCut := pkg.NewCuter(cmdLine)
	myCut.CutFiles()
}
