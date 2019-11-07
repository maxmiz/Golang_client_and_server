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
		NomeCliente, _ := bufio.NewReader(conn).ReadString('\n')
		//fmt.Print(string(message))
		//Mensagem do cliente
		MenssagemCliente, _ := bufio.NewReader(conn).ReadString('\n')
		//fmt.Print("Message Received:", string(message))
		fmt.Println("[" + string(NomeCliente) + string(MenssagemCliente))

		conn.Write([]byte("OK\n"))
	}
}
