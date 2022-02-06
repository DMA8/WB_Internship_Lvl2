package pkg

import (
	"fmt"
	"math/rand"
	"time"
)

type Specialist struct {
	Next	ChainMember
}

func (s *Specialist)Execute(c *Client) {
	rand.Seed(time.Now().UnixMicro())
	if a := rand.Intn(100); a > 30 {
		fmt.Println("The problem is finaly solved by the specialist!")
		c.ProblemSolved = true
	} else {
		fmt.Println("The client's problem has not been solved at all!")
	}
}