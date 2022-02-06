package pkg

type Car struct {
	CEngine		Engine
	CWheels		Wheels
	CMultimedia	Multimedia
	Color		string
	Firm		string
}

type Engine struct {
	HP			float64
	Volume		float64
	Mileage		float64
	NCylynders	int
}

type Wheels struct {
	Size		float64
	TireFirm	string
	Mileage		float64
}

type Multimedia struct {
	MonitorSize		float64
	NSpeakers		int
	FirmSpeakers	string
}
