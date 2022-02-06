package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// ДОБАВИТЬ ТАЙМАУТ
func main() {
	allArgs := os.Args[1:]
	timeOut := strings.Split(allArgs[0], "=")[1]
	timeOut = strings.TrimSuffix(timeOut, "s")
	host := allArgs[1]
	port := allArgs[2]
	nSec, _ := strconv.Atoi(timeOut)
	TCPClient(host, port, int32(nSec))
}

//TCPClient connects via tcp to given socket. 
func TCPClient(host, port string, timoutSec int32){
	var wg sync.WaitGroup
	conn, err := net.DialTimeout("tcp", host + ":" + port, time.Second * time.Duration(timoutSec))
	if err != nil {
		log.Fatal(err)
	}
	wg.Add(1)
	go msgSender(&wg, conn)
	go msgGetter(&wg, conn)
	wg.Wait()
}

func msgSender(wg *sync.WaitGroup, connect net.Conn){
	stdInReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Text to send")
		text, err := stdInReader.ReadString('\n')
		if err != nil {
			if err == io.EOF{
				connect.Close()
				fmt.Println("Connect is closed. Exit...")
				wg.Done()
				return
			}
			log.Fatal(err)
		}
		fmt.Fprint(connect, text)
	}
}

func msgGetter(wg *sync.WaitGroup, connect net.Conn){
	connReader := bufio.NewReader(connect)
	for {
		message, err := connReader.ReadString('\n')
		if err != nil {
			if err == io.EOF{
				connect.Close()
				fmt.Println("Server is down ")
				wg.Done()
				return
			} 
			log.Fatal(err)
		}
		message = strings.Trim(message, "\n")
		fmt.Println(connect.RemoteAddr().String(), message)
	}
}
