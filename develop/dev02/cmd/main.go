package main

import (
	"fmt"
	"log"

	"../pkg"
)

func main() {
	var inputString string 
	//= `a4bc2d5e`
	fmt.Scan(&inputString)
	result, err := pkg.Unpack(inputString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result is|%s|\n", result)
	result2, err := pkg.EasyUnpack(inputString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result is|%s|\n", result2)
}