package pkg

import (
	"fmt"
	"time"
)
type SleepingWariorState struct{
	Warior	*Warior
}

func (s *SleepingWariorState)Heal(incomeHealt int){
	fmt.Printf("Player %s is healing while sleeping...", s.Warior.Name)
	for i := 0; i < int(incomeHealt); i++ {
		time.Sleep(time.Millisecond * 10)
		s.Warior.Health += 5
		s.Warior.Energy += 10
	}
	fmt.Printf("Health of player %s is increased on %d\n", s.Warior.Name, incomeHealt)
}

func (f *SleepingWariorState)GetDamage(incomeDamage int, Enemy *Warior){
	f.Warior.CurrentState = f.Warior.DeadWariorState
	fmt.Printf("Warior %s has been killed while sleeping by %s\n", Enemy.Name)
}

func (s *SleepingWariorState)DoDamage(outcomeDamage int, e *Warior) {
	fmt.Printf("Warior %s is sleeping. It can not damage anyone\n", s.Warior.Name)
}