package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	tcpAddr , err := net.ResolveTCPAddr("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(fmt.Errorf("error when ResolveTCPAddr: %v", err))
	}

	listener , err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(fmt.Errorf("error when ListenTCP: %v", err))
	}

	conn, err := listener.Accept()

	if err != nil {
		log.Fatal(fmt.Errorf("error when establish connection: %v", err))
	}
	
	go io.Copy(conn, os.Stdin)
	io.Copy(os.Stdout, conn)
}
