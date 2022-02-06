package pkg

import (
	"fmt"
	"math/rand"
	"time"
)

type ChatBot struct {
	Next	ChainMember
}

func (chatbot *ChatBot)Execute(client *Client) {
	rand.Seed(time.Now().UnixMicro())
	if a := rand.Intn(100); a > 75 {
		client.ProblemSolved = true
		fmt.Println("Problem solved with bot's help")
	} else {
		chatbot.Next.Execute(client)
		fmt.Println("Bot has not helped to the client. Problem is passed to the operator")
	}
}
