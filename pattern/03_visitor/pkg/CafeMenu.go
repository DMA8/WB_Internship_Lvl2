package pkg

type Menu struct {
	MenuItems	[]MenuItem
}

type MenuItem struct {
	Name		string
	Cost		float64
	Popularity	int
	Ingridiends	[]Ingridient
}

type Ingridient struct {
	Name		string
	Calories	float64
	Carbs		float64
	Protein		float64
	Fat			float64
}