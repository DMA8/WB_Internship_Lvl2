package pkg

import "fmt"

type DeadWariorState struct{
	Warior	*Warior
}

func (d *DeadWariorState)Heal(incomeHealt int){
	fmt.Printf("Warior %s is dead. Health can't be restored\n", d.Warior.Name)
}

func (d *DeadWariorState)GetDamage(incomeDamage int, e *Warior){
	fmt.Printf("Warior %s is dead. You can't damage dead players\n", d.Warior.Name)
}

func (d *DeadWariorState)DoDamage(outcomeDamage int, e *Warior) {
	fmt.Printf("Warior %s is dead. It can not damage anyone\n", d.Warior.Name)
}