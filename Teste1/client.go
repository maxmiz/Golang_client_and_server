package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func send_message(r *bufio.Reader, conexao net.Conn, mensagem_aux string) {
	fmt.Printf(mensagem_aux)
	text, _ := r.ReadString('\n')
	fmt.Fprintf(conexao, mensagem_aux+" "+text+"\n")
}

func main() {

	conn, _ := net.Dial("tcp", "127.0.0.1:8082")
	for {

		reader := bufio.NewReader(os.Stdin)
		//fmt.Print("Text to send: ")

		//Cliente
		send_message(reader, conn, "Cliente:")

		//Mensagem
		send_message(reader, conn, "Mensagem:")

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}

}
