package pkg

import "fmt"

type Kettle struct {
	Volume	float64
	IsHot	bool
}

func (k *Kettle)Boil() {
	k.IsHot = true
}

type Film struct {
	Name 	string
	Ref		string
}

func (f *Film) FindRefForFreeFilm() {
	fmt.Printf("Looking for free film %s\n", f.Name)
	f.Ref = fmt.Sprintf("https:/google.com/%s", f.Name)
}

type CellPhone struct {
	SilentMode	bool
}

func (c *CellPhone)SwitchSilent() {
	c.SilentMode = true
}