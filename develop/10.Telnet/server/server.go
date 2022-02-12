package main

import (
	"sync"
	"bufio"
	"net"
	"os"
	"io"
	"fmt"
	"log"
	"strings"
)

func main(){
	TCPServer()
}

func TCPServer(){
	var wg sync.WaitGroup
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	connect, err := listen.Accept()
	if err != nil {
		log.Fatal(err)
	}
	wg.Add(1)
	go msgSender(&wg, connect)
	go msgGetter(connect)
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
		fmt.Fprintf(connect, text)
	}
}

func msgGetter(connect net.Conn){
	connReader := bufio.NewReader(connect)
	for {
		message, err := connReader.ReadString('\n')
		if err != nil {
			if err == io.EOF{
				continue
			} 
			log.Fatal(err)
		}
		message = strings.Trim(message, "\n")
		fmt.Println(connect.RemoteAddr().String(), message)
	}
}
