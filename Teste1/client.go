package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	for {

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		//Cliente
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, "Client: \n")

		//Mensagem
		text, _ = reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, text+"\n")

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}

}

func send_message(r *reader) {
	text, _ := r.ReadString('\n')
	fmt.Fprintf()
}
