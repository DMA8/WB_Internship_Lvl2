package main

import (
	"fmt"

	"../pkg"
)

/*
Фабричный метод — это порождающий паттерн проектирования,
который определяет общий интерфейс для создания объектов в суперклассе,
позволяя подклассам изменять тип создаваемых объектов.
*/

func main(){
	myFactory := pkg.VehicleFactory{}
	myCar := myFactory.NewVehicle("car")
	fmt.Println(myCar)
	myTruck := myFactory.NewVehicle("truck")
	fmt.Println(myTruck)
}