package pkg


type Transporter interface { // можно расширить, если нужны другие методы
	Move()
}

type VehicleFactory struct{

}

func(v VehicleFactory)NewVehicle(className string) Transporter{ // нельзя вернуть тип Vehicle, потому что нет наследования(
	switch className{
	case "bike":
		return NewBike()
	case "truck":
		return NewTruck()
	case "car":
		return NewCar()
	}
	return &Vehicle{"SuperJet"}
}