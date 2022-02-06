package pkg

type WariorState interface{
	GetDamage(int, *Warior)
	Heal(int)
	DoDamage(int, *Warior)
}