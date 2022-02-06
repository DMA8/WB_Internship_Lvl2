package pkg

import "fmt"

type Bike struct {
	Vehicle
	Name	string
	Firm	string
}

func (b Bike)Move(){
	fmt.Printf("%s is moving\n", b.Name)
}

func (b Bike)Flip(){
	fmt.Printf("bike.Name = %s. And It has flipped right now!", b.Name)
}

func NewBike()*Bike{
	return &Bike{
		Vehicle{"Light"},
		"Gazulya",
		"Zaporozhets",
	}
}