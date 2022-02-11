package pkg

import (
	"fmt"
)

type GameActor struct {
	Name	string
}

func (g GameActor)Jump(){
	fmt.Printf("I am GameActor with name %s. And I just have jumped!\n", g.Name)
}
