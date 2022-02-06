package pkg

import "fmt"

type Car struct {
	Vehicle
	CarName	string
	Price	float64
}

func (c Car)Move(){
	fmt.Printf("%s is moving\n", c.CarName)
}

func(c Car)Beep(){
	fmt.Printf("I am Car %s and I'm doing beeeeep", c.CarName)
}

func NewCar()*Car{
	return &Car{
		Vehicle{"Car"},
		"Mashina",
		890,
	}
}