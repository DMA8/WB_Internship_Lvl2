package pkg

import (
	"fmt"
	"math/rand"
	"time"
)

type Operator struct {
	Next	ChainMember
}

func (o *Operator)Execute(client *Client){
	rand.Seed(time.Now().UnixMicro())
	if rand.Intn(100) > 50 {
		fmt.Println("Problem is solved by the operator")
		client.ProblemSolved = true
	} else {
		o.Next.Execute(client)
		fmt.Println("Problem is passed by operator to specialist")
	}
}