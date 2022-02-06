package pkg

import (
	"bytes"
	"strings"
	"log"
	"os"
	"io"
	"os/exec"
)
type Shell struct {
	CurrentWD	string
}


func (s *Shell)GetWD(){
	var err error
	s.CurrentWD, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
}



func HandleCommand2(commandStr string, in io.Reader, args ...string) (io.Reader, func()error) {
	splittedCmd 
	
	cmd := exec.Command(commandStr, args...)
	cmd.Stdin = in
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal("1", err)
	}

	go func(){
		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}
	}()
	return stdout, cmd.Wait
}