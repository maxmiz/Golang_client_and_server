package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

type options struct {
	nomeCliente string
}

func initArgs(dados *options) {
	flag.StringVar(&dados.nomeCliente, "nome", "", "help")
	flag.Parse()
	//fmt.Println(dados.nomeCliente)
}

func main() {
	var ross options
	initArgs(&ross)
	fmt.Println(ross.nomeCliente)

	conn, _ := net.Dial("tcp", "127.0.0.1:8082")
	for {

		reader := bufio.NewReader(os.Stdin)

		//Cliente
		sendClient(ross.nomeCliente, conn, "Cliente:")

		//Mensagem
		sendMessage(reader, conn, "Mensagem:")

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}

}

func sendMessage(r *bufio.Reader, conexao net.Conn, mensagemAux string) {
	fmt.Printf(mensagemAux)
	text, _ := r.ReadString('\n')
	fmt.Fprintf(conexao, text+"\n")
}

func sendClient(text string, conexao net.Conn, mensagemAux string) {
	//fmt.Printf(mensagemAux)
	fmt.Fprintf(conexao, text+"\n")
}
