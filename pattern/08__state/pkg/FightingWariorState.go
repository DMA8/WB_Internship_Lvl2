package pkg

import (
	"fmt"
	"time"
)

type FightingWariorState struct{
	Warior	*Warior
}

func (f *FightingWariorState)Heal(incomeHealt int){
	fmt.Printf("Player %s is healing while fighting...", f.Warior.Name)
	for i := 0; i < int(incomeHealt); i++ {
		time.Sleep(time.Millisecond * 5)
		f.Warior.Health += 1
		f.Warior.Energy += 1
	}
	fmt.Printf("Health of player %s is increased on %d\n", f.Warior.Name, incomeHealt)
}

func (f *FightingWariorState)GetDamage(incomeDamage int, Enemy *Warior){
	if f.Warior.Health - incomeDamage <= 0 {
		fmt.Printf("Player %s IS KILLED by enemy with name %s!!!\n", f.Warior.Name, Enemy.Name)
		f.Warior.CurrentState = f.Warior.DeadWariorState
		f.Warior.Health = 0
	} else {
		f.Warior.Health -= incomeDamage
		f.Warior.Energy -= 50
		if f.Warior.Energy < 0 {
			f.Warior.CurrentState = f.Warior.SleepingWariorState
		}
		fmt.Printf("Player %s has got %d damaged by enemy %s\n", f.Warior.Name, incomeDamage, Enemy.Name)
	}
}

func (f *FightingWariorState)DoDamage(outcomeDamage int, Enemy *Warior){
	fmt.Printf("Warior %s is damaging %s\n", f.Warior.Name, Enemy.Name)
	Enemy.GetDamage((outcomeDamage *f.Warior.Energy) / 100, f.Warior) // мощность удара зависит от энергии атакующего
	f.Warior.Energy -= 10
	if f.Warior.Energy < 0 {
		f.Warior.CurrentState = f.Warior.SleepingWariorState
	}
}

