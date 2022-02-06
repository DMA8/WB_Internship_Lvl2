package main

import (
	"../pkg"
	"bufio"
	"os"
	"log"
	"strings"
)

func main(){
	//inputString := "-Fni context ../texts/t"
	var inputString string
	scn := bufio.NewReader(os.Stdin)
	inputString, err := scn.ReadString('\n')
	if err != nil{
		log.Fatal(err)
	}
	inputString = strings.TrimSuffix(inputString, "\n")
	cmd := pkg.NewCmdLine(inputString)
	greper := pkg.NewGreper(cmd)
	greper.Main()
}
