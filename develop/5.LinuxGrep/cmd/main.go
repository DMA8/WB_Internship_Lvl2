package main

import (
	"../pkg"
	"os"
	"strings"
)

func main(){
	inputString := strings.Join(os.Args[1:], " ")
	cmd := pkg.NewCmdLine(inputString)
	greper := pkg.NewGreper(cmd)
	greper.Main()
}
