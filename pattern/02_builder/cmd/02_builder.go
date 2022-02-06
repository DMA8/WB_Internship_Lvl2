package main

import (
	"fmt"

	"../pkg"
)

/*
Строитель — это порождающий паттерн проектирования,
который позволяет создавать сложные объекты пошагово.
Строитель даёт возможность использовать один и тот же
код строительства для получения разных представлений
объектов.

Паттерн Строитель предлагает вынести конструирование
объекта за пределы его собственного класса, поручив это
дело отдельным объектам, называемым строителями.
*/
func main() {
	var myBuilder CarBuilder
	var buildInstance BuilderOfCars
	myBuilder = &buildInstance
	myBuilder.CreateCar()
	myBuilder.SetCarColor("red")
	myBuilder.SetEngine("VAZ")
	fmt.Println(myBuilder.GetCar())
}

type CarBuilder interface {
	SetEngine(string)
	SetWheels(string)
	SetMultimedia(string)
	SetCarColor(string)
	SetCarFirm(string)
	GetCar()pkg.Car
	CreateCar()
}

type BuilderOfCars struct {
	NewCar	*pkg.Car
}

func (c *BuilderOfCars)SetEngine(engineName string) {
	switch engineName {
	case "VAZ":
		c.NewCar.CEngine.Volume = 1.6
		c.NewCar.CEngine.HP = 98.9
		c.NewCar.CEngine.NCylynders = 4
		c.NewCar.CEngine.Mileage = 0
	case "1JZ":
		c.NewCar.CEngine.Volume = 2.5
		c.NewCar.CEngine.HP = 200.3
		c.NewCar.CEngine.NCylynders = 6
		c.NewCar.CEngine.Mileage = 1000000
	}
}

func (c *BuilderOfCars)SetWheels(wheelCost string) {
	switch wheelCost{
	case "cheap":
		c.NewCar.CWheels.Size = 14
		c.NewCar.CWheels.TireFirm = "Zaporozhskaya Kama"
		c.NewCar.CWheels.Mileage = 10000
	case "default":
		c.NewCar.CWheels.Size = 16
		c.NewCar.CWheels.TireFirm = "GoodTire"
		c.NewCar.CWheels.Mileage = 0
	case "top":
		c.NewCar.CWheels.Size = 22
		c.NewCar.CWheels.TireFirm = "BestTire"
		c.NewCar.CWheels.Mileage = 0
	}
}

func (c *BuilderOfCars)SetMultimedia(multimediaConf string) {
	switch multimediaConf{
	case "cheap":
		c.NewCar.CMultimedia.MonitorSize = 3
		c.NewCar.CMultimedia.NSpeakers = 1
		c.NewCar.CMultimedia.FirmSpeakers = "Lada" 
	case "tachka-prokachka":
		c.NewCar.CMultimedia.MonitorSize = 1 << 63
		c.NewCar.CMultimedia.NSpeakers = 1 << 31
		c.NewCar.CMultimedia.FirmSpeakers = "GodLike" 
	}
}

func (c *BuilderOfCars)SetCarColor(color string) {
	c.NewCar.Color = color
}

func (c *BuilderOfCars)SetCarFirm(firm string) {
	c.NewCar.Firm = firm
}

func (c *BuilderOfCars)CreateCar(){
	c.NewCar = &pkg.Car{}
}

func (c *BuilderOfCars)GetCar() pkg.Car {
	return *c.NewCar
}