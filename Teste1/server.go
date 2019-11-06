package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	fmt.Println("Launching server...")

	ln, _ := net.Listen("tcp", ":8082")

	// accept connection on port
	conn, _ := ln.Accept()

	// run loop forever (or until ctrl-c)
	for {

		//Nome do cliente
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Client:", string(message))
		//Mensagem do cliente
		message, _ = bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received:", string(message))

		conn.Write([]byte("OK\n"))
	}
}
