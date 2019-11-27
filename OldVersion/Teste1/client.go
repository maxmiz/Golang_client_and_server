package main

import (
	"bufio"
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"math/cmplx"
	"math/rand"
	"net"
	"strconv"
	"time"
)

type options struct {
	nomeCliente string
	ip          string
	port        string
}
type numero struct {
	primo  complex128
	raiz   complex128
	modulo complex128
}

func initArgs(dados *options) {
	flag.StringVar(&dados.nomeCliente, "nome", "", "help")
	flag.StringVar(&dados.ip, "ip", "", "help")
	flag.StringVar(&dados.port, "port", "", "help")
	flag.Parse()
}
func Gera(auxiliar *numero) {
	rand.Seed(time.Now().UnixNano())
	auxiliar.primo = complex(rand.Float64(), rand.Float64())
	auxiliar.raiz = complex(rand.Float64(), rand.Float64())
	auxiliar.modulo = complex(rand.Float64(), rand.Float64())
}

func main() {
	var ross options
	initArgs(&ross)
	fmt.Println(ross.nomeCliente)
	var vars_diffie numero
	Gera(&vars_diffie)
	fmt.Println(vars_diffie.primo)
	fmt.Println(vars_diffie.raiz)
	fmt.Println(vars_diffie.modulo)
	fmt.Println(calculo(vars_diffie.primo, vars_diffie.raiz, vars_diffie.modulo))

	conn, _ := net.Dial("tcp", ross.ip+":"+ross.port)
	//Diffie-Hellman
	//rand.Seed(time.Now().UnixNano())
	//A := rand.Intn(1000-10) + 10
	//B := rand.Intn(1000-10) + 10
	//fmt.Fprintln(conn, A)
	time.Sleep(100 * time.Millisecond)
	//fmt.Fprintln(conn, B)

	Envia(conn, vars_diffie.primo)
	time.Sleep(100 * time.Millisecond)
	Envia(conn, vars_diffie.raiz)
	/*
		Alice := rand.Intn(1000-10) + 10

		Alice_Ra := (B ^ Alice) % A

		fmt.Println("Alice_Ra: ", Alice_Ra)
		time.Sleep(100 * time.Millisecond)
		fmt.Fprintln(conn, Alice_Ra)
		time.Sleep(100 * time.Millisecond)
		Bob_Ra, _ := bufio.NewReader(conn).ReadString('\n')
		Bob_Ra = strings.TrimRight(Bob_Ra, "\n")
		fmt.Println("Bob_ra:", Bob_Ra)
		//k := (converte(Bob_Ra) ^ Alice_Ra) % A
		k := (B ^ (converte(Bob_Ra) * Alice_Ra)) % A
		fmt.Println(k)
		//Diffie-Hellman
		for {

			reader := bufio.NewReader(os.Stdin)

			//Cliente
			sendClient(ross.nomeCliente, conn)

			//Mensagem
			sendMessage(reader, conn, "Mensagem:", ross.nomeCliente)

			message, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Print("Message from server: " + message)
			fmt.Println()
		}*/

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

func converte(valor string) int {
	x, _ := strconv.Atoi(valor)
	return x
}

func Envia(conexao net.Conn, mensagem complex128) {
	//Envia mensagem para a conexao aberta
	fmt.Fprintln(conexao, mensagem)
	//err := binary.Write(conexao)
}

func calculo(primo, raiz, modulo complex128) complex128 {
	aux := cmplx.Pow(raiz, primo)
	return aux - (primo * (aux / primo))
}
