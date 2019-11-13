package main

import (
	"bufio"
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

type options struct {
	nomeCliente string
	ip          string
	port        string
}

func initArgs(dados *options) {
	flag.StringVar(&dados.nomeCliente, "nome", "", "help")
	flag.StringVar(&dados.ip, "ip", "", "help")
	flag.StringVar(&dados.port, "port", "", "help")
	flag.Parse()
	//fmt.Println(dados.nomeCliente)
}

func main() {
	var ross options
	initArgs(&ross)
	fmt.Println(ross.nomeCliente)

	conn, _ := net.Dial("tcp", ross.ip+":"+ross.port)
	for {

		reader := bufio.NewReader(os.Stdin)

		//Cliente
		sendClient(ross.nomeCliente, conn)

		//Mensagem
		sendMessage(reader, conn, "Mensagem:", ross.nomeCliente)

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
		fmt.Println()
	}

}

func sendMessage(r *bufio.Reader, conexao net.Conn, mensagemAux string, nomeCliente string) {
	fmt.Printf(mensagemAux)
	text, _ := r.ReadString('\n')

	fmt.Fprintf(conexao, text+"\n")
	time.Sleep(1000 * time.Millisecond)
	fmt.Fprintf(conexao, Md5EmptyHash(nomeCliente+text)+"\n")

}

func sendClient(text string, conexao net.Conn) {
	fmt.Fprintf(conexao, text+"\n")
}

func Md5EmptyHash(message string) string {
	h := md5.New()
	io.WriteString(h, message)
	return fmt.Sprintf("%x", h.Sum(nil))
}
