package pkg

import "fmt"

type Truck struct{
	Vehicle
	TruckName		string
	MaxWeight		float64
	CurrenWeight	float64
}

func (t Truck)Move(){
	fmt.Printf("%s is moving\n", t.TruckName)
}

func (t *Truck)TruckAddWeight(w float64) {
	if t.CurrenWeight + w < t.MaxWeight {
		t.CurrenWeight += w
		fmt.Printf("Weight is added to Truck %s. Current weight is %f, maxweight is %f\n", t.TruckName, t.CurrenWeight, t.MaxWeight)
	} else {
		fmt.Printf("Too much weight! max possible weight to add is %f\n", t.MaxWeight - t.CurrenWeight)
	}
}

func NewTruck()*Truck{
	return &Truck{
		Vehicle{"Heavy"},
		"KAMAZ",
		1000,
		0,
	}
}
