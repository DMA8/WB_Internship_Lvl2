package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
//	"../pkg"
	"os/exec"
	"strings"
	"io"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		wd, err := os.Getwd() 
		fmt.Print("\033[32m", wd, "$ \033[37m")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		err = RunCommand(cmdString)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func RunCommand(commandStr string) io.Reader {
	buf := &bytes.Buffer{}
	commandStr = strings.TrimSuffix(commandStr, "\n")
	arrCommandStr := strings.Fields(commandStr)
	if len(arrCommandStr) < 1 {
		return nil
	}
	switch arrCommandStr[0] {
	case "exit":
		os.Exit(0)
	case "cd":
		os.Chdir(arrCommandStr[1])
		return nil
	default:
	}
	cmd := exec.Command(arrCommandStr[0], arrCommandStr[1:]...) // pwd echo kill
	cmd.Stderr = os.Stderr
	cmd.Stdout = buf
	cmd.Run()
	return buf
}

