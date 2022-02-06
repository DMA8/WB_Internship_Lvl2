package pkg

type ChainMember interface {
	Execute(*Client)
}