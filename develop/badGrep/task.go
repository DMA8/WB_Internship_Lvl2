package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"./pkg"
)
//Strings CONTAINS когда ищем не по регулярке
// Считать число рун, без явного преобразования - utf8.RuneCountInString(str)


func main() {
	cmd := pkg.NewCmdLine(`-A3B2 -B3F . file3 file4`)
	pattern := pkg.InitRegexCompile(*cmd)
	file, _ := os.Open("t")
	q := pkg.QueueLines{"", make([]string, 0, cmd.Before), make([]string, 0, cmd.After),  cmd.Before, cmd.After}
	defer file.Close()
	//a := bufio.NewReader(file)
	a := pkg.ReaderImproved{*bufio.NewReader(file)}
	t := a.ReadNString(2)
	fmt.Println(t)
	for line, err2 := a.Reader.ReadString('\n'); err2 != io.EOF; {
		if err2 != nil && err2 != io.EOF{
			log.Fatal(err2)
		}
		res := pattern.Match([]byte(line))
		if res {
			for _, v := range q.BeforeStrings{
				fmt.Print(v)
			}
			fmt.Print(line)
			for _, v := range q.AfterStrings{
				fmt.Print(v)
			}
		}
		if cmd.Before > 0 {
			q.AddStringBefore(line)
		}
		if cmd.After > 0 {
			q.AddStringAfter(line)
		}
		if err2 == io.EOF {
			break
		}
		line, err2 = a.Reader.ReadString('\n')
	}
}

