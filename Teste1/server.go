package main

import (
	"bufio"
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

type Serveroptions struct {
	ip   string
	port string
}

func initArgs(dados *Serveroptions) {
	flag.StringVar(&dados.ip, "ip", "", "help")
	flag.StringVar(&dados.port, "port", "", "help")
	flag.Parse()
}

func main() {

	var serverInstania Serveroptions
	initArgs(&serverInstania)

	fmt.Println("Launching server...")

	ln, _ := net.Listen("tcp", serverInstania.ip+":"+serverInstania.port)

	conn, _ := ln.Accept()
	//Diffie-Hellman

	X := converte(Diffie_part1(conn))
	Y := converte(Diffie_part1(conn))

	Bob := rand.Intn(1000-10) + 10
	//fmt.Println(X, Y, Bob)

	Bob_Ra := (Y ^ Bob) % X
	fmt.Println("Bob_ra:", Bob_Ra)

	Alice_Ra, _ := bufio.NewReader(conn).ReadString('\n')
	Alice_Ra = strings.TrimRight(Alice_Ra, "\n")
	fmt.Println("Alice_Ra: ", Alice_Ra)
	time.Sleep(100 * time.Millisecond)

	fmt.Fprintln(conn, Bob_Ra)
	//Diffie-Hellman

	k := (Y ^ (converte(Alice_Ra) * Bob_Ra)) % X
	//k := (converte(Alice_Ra) ^ Bob) % X
	fmt.Println(k)

	for {
		defer fmt.Println("We have problems")
		//Nome do cliente
		NomeCliente, _ := bufio.NewReader(conn).ReadString('\n')
		NomeCliente = strings.TrimRight(NomeCliente, "\n")
		//fmt.Print(string(message))
		//Mensagem do cliente
		MenssagemCliente, _ := bufio.NewReader(conn).ReadString('\n')
		//MenssagemCliente = strings.TrimRight(MenssagemCliente, "\n")

		Hash, _ := bufio.NewReader(conn).ReadString('\n')
		Hash = strings.TrimRight(Hash, "\n")
		fmt.Println(Hash)
		fmt.Println(len(Hash))
		if Hash == Md5EmptyHash(NomeCliente+MenssagemCliente) {

			fmt.Println("[" + string(NomeCliente) + "] " + string(MenssagemCliente))

			conn.Write([]byte("OK\n"))
		} else {
			conn.Write([]byte("We have problems\n"))
		}
	}
}

func Md5EmptyHash(message string) string {
	h := md5.New()
	io.WriteString(h, message)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func converte(valor string) int {
	x, _ := strconv.Atoi(valor)
	return x
}

func Diffie_part1(conexao net.Conn) string {
	//Função usada para receber o "q" depois o "p" do Diffie
	aux, _ := bufio.NewReader(conexao).ReadString('\n')
	return strings.TrimRight(aux, "\n")
}
