package pkg

type Warior struct{
	FightingWariorState	WariorState
	DeadWariorState		WariorState
	SleepingWariorState	WariorState

	CurrentState		WariorState

	Name				string
	Health				int
	Energy				int
}


func NewWarior(Name string, Health, Energy int) *Warior{
	warior := &Warior{
		Name: Name,
		Health: Health,
		Energy: Energy,
	}
	deadWariorState := &DeadWariorState{
		Warior:	warior,
	}
	fightingWariorState := &FightingWariorState{
		Warior: warior,
	}
	sleepingWariorState := &SleepingWariorState{
		Warior: warior,
	}

	warior.CurrentState = fightingWariorState
	warior.FightingWariorState = fightingWariorState
	warior.DeadWariorState = deadWariorState
	warior.SleepingWariorState = sleepingWariorState
	
	return warior
}

func (f *Warior)Heal(incomeHealt int){
	f.CurrentState.Heal(incomeHealt)
}

func (f *Warior)GetDamage(incomeDamage int, Enemy *Warior){
	f.CurrentState.GetDamage(incomeDamage, Enemy)
}

func (f *Warior)DoDamage(outcomeDamage int, Enemy *Warior){
	f.CurrentState.DoDamage(outcomeDamage, Enemy)
}