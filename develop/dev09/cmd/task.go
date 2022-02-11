package main

import (
	"fmt"
	"os"

	"../pkg"
)

func main() {
	myUrl := os.Args[1]
	pkg.MainDownload(myUrl)
	fmt.Println("done")
}
