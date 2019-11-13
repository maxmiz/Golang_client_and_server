package main

import (
	"bufio"
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	"golang.org/x/crypto/bcrypt"
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

	// accept connection on port
	conn, _ := ln.Accept()

	// run loop forever (or until ctrl-c)
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
		if Hash == Md5EmptyHash(NomeCliente+MenssagemCliente) {

			fmt.Println("[" + string(NomeCliente) + "] " + string(MenssagemCliente))
			//fmt.Println(Hash)

			conn.Write([]byte("OK\n"))
		} else {
			conn.Write([]byte("We have problems\n"))
		}
	}
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func Md5EmptyHash(message string) string {
	h := md5.New()
	io.WriteString(h, message)
	return fmt.Sprintf("%x", h.Sum(nil))
}
