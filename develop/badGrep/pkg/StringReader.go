package pkg
// package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

type ReaderImproved struct {
	Reader	bufio.Reader // Можно было и наследоваться, но композиция
}

func (r *ReaderImproved)ReadNString(n int) []string{
	outStrings := make([]string, 0, n)
	for i := 0; i < n; i++ {
		str, err := r.Reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		outStrings = append(outStrings, str)
	}
	return outStrings
}


// func main() {
// 	file := "../t"
// 	opened := openFile(file)
// 	stringRead := StringReader(opened)
// 	my1string, _ := stringRead.ReadString('\n')
// 	time.Sleep(100)
// 	my2string, _ := stringRead.ReadString('\n')
// 	fmt.Print(my1string, my2string)
// 	fewStrings := GetNStrings(stringRead, 5)
// 	fmt.Println(fewStrings)
// }

func StringReader(file *os.File) *bufio.Reader{
	return bufio.NewReader(file)
}

func openFile(fname string) *os.File {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
