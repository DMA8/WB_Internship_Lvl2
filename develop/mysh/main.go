package main

import (
	"fmt"
	"io"
	"io/fs"
	"bufio"
	"log"
	"os"
	"strings"
	"syscall"
	"errors"
	"os/exec"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	for {
		wd, _ := os.Getwd() 
		fmt.Print("\033[32m", wd, ">$ \033[37m")
		line, _ := reader.ReadString('\n')
		if line == "exit\n" {
			break
		}
		if line == "\n" {
			continue
		}
		HandleForks(line[:len(line)-1])
		fmt.Println("------------------------------------------")

	}
}

func HandleForks(inpLine string){
	forks := strings.Split(inpLine, "&")
	for i := 0; i < len(forks) - 1; i++{
		pid, _, err := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
		if err != 0 {
			log.Fatal(err)
		}
		if pid == 0 {
			fmt.Println(os.Getegid())
			HandlePipes(forks[i])
			fmt.Println("Done")
			os.Exit(0)
		}
	}
	HandlePipes(forks[len(forks) - 1])
}

func HandlePipes(inpCommands string) {
	var resultOutFromPipe io.Reader

	commands := strings.Split(inpCommands, "|")
	waits := []func() error{}
	returnedFromPrevCommand, wait := HandleCommand(commands[0], os.Stdin)
	waits = append(waits, wait)

	for i := 1; i < len(commands) - 1; i++ {
		returnedFromCommand, wait := HandleCommand(commands[i], returnedFromPrevCommand)
		waits = append(waits, wait)
		returnedFromPrevCommand = returnedFromCommand
	}
	if len(commands) > 1 {
		resultOutFromPipe, wait = HandleCommand(commands[len(commands) - 1], returnedFromPrevCommand)
		waits = append(waits, wait)
	} else {
		resultOutFromPipe, wait = returnedFromPrevCommand, nil
	}
	if resultOutFromPipe != nil {
		_, err := io.Copy(os.Stdout, resultOutFromPipe)
		if err != nil && !errors.Is(err, fs.ErrClosed){
			log.Fatal(err)
		}
	}
	for _, wait := range waits{
		if wait != nil {
			wait()
		}
	}
}

func HandleCommand(commandStr string, in io.Reader, args ...string) (io.Reader, func()error) {
	commands := strings.Fields(commandStr)
	if len(commands) == 0 {
		return nil, nil
	}
	switch commands[0]{
	case "cd":
		err := os.Chdir(commands[1])
		if err != nil {
			log.Fatal(err)
		}
		return nil, nil
	default:
		cmd := exec.Command(commands[0], commands[1:]...)
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
}