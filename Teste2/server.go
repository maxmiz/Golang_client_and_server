package main

import (
	"bufio"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"math/cmplx"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	. "github.com/logrusorgru/aurora"
)

type ArgumentosServer struct {
	ip   string
	port string
}
type ComplexDiffie struct {
	ParteReal       float64
	ParteImaginaria float64
}
type EstruturaDiffie struct {
	choice ComplexDiffie
	raiz   ComplexDiffie
	modulo ComplexDiffie
	B      ComplexDiffie
	A      ComplexDiffie
	key    ComplexDiffie
}

func InicializaArgumentos(dados *ArgumentosServer) {
	flag.StringVar(&dados.ip, "ip", "", "help")
	flag.StringVar(&dados.port, "port", "", "help")
	flag.Parse()
}

func main() {
	var NomeCliente string
	var contador = 0
	var serverInstania ArgumentosServer
	InicializaArgumentos(&serverInstania)
	var InstanciaDiffie EstruturaDiffie
	InicializaEstruturaDiffie(&InstanciaDiffie)
	fmt.Println("Launching server...")

	ln, _ := net.Listen("tcp", serverInstania.ip+":"+serverInstania.port)

	conn, _ := ln.Accept()
	//Etapa1
	InstanciaDiffie.modulo.ParteReal, _ = strconv.ParseFloat(breakline(recebefloat(conn)), 64)
	InstanciaDiffie.modulo.ParteImaginaria, _ = strconv.ParseFloat(breakline(recebefloat(conn)), 64)

	InstanciaDiffie.raiz.ParteReal, _ = strconv.ParseFloat(breakline(recebefloat(conn)), 64)
	InstanciaDiffie.raiz.ParteImaginaria, _ = strconv.ParseFloat(breakline(recebefloat(conn)), 64)
	//Etapa2
	rand.Seed(time.Now().UnixNano())
	fmt.Println(time.Now().UnixNano())
	//generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	InstanciaDiffie.choice.ParteReal = rand.Float64()
	InstanciaDiffie.choice.ParteImaginaria = rand.Float64()
	//Etapa3
	InstanciaDiffie.B.ParteReal, InstanciaDiffie.B.ParteImaginaria = calculo(InstanciaDiffie.raiz, InstanciaDiffie.choice, InstanciaDiffie.modulo)
	//Etapa4
	InstanciaDiffie.A.ParteReal, _ = strconv.ParseFloat(breakline(recebefloat(conn)), 64)
	InstanciaDiffie.A.ParteImaginaria, _ = strconv.ParseFloat(breakline(recebefloat(conn)), 64)
	//Etapa5
	fmt.Fprintln(conn, InstanciaDiffie.B.ParteReal)
	time.Sleep(100 * time.Millisecond)
	fmt.Fprintln(conn, InstanciaDiffie.B.ParteImaginaria)

	fmt.Println("Alice_from_server: ", InstanciaDiffie.A)

	fmt.Println("Bob_from_server: ", InstanciaDiffie.B)
	//Etapa6
	InstanciaDiffie.key.ParteReal, InstanciaDiffie.key.ParteImaginaria = calculo2(InstanciaDiffie.raiz, InstanciaDiffie.B, InstanciaDiffie.A, InstanciaDiffie.modulo)
	fmt.Println("key: ", InstanciaDiffie.key)

	for {
		NomeCliente, _ = bufio.NewReader(conn).ReadString('\n')
		mensagem, _ := bufio.NewReader(conn).ReadString('\n')
		HMACFromClient, _ := bufio.NewReader(conn).ReadString('\n')
		MD5 := Md5EmptyHash(breakline(NomeCliente) + mensagem)

		if breakline(mensagem) != "quit" {

			if breakline(HMACFromClient) == HMAC(Md5EmptyHash(breakline(NomeCliente)+mensagem+strconv.Itoa(contador)), float64ToByte(InstanciaDiffie.key.ParteImaginaria)) && VereficaArquivo(breakline(NomeCliente)+".dat", MD5) == false {
				//fmt.Println("["+breakline(NomeCliente)+"]: ", breakline(mensagem))
				CriaArquivo(breakline(NomeCliente)+".dat", MD5)
				imprimir(NomeCliente, mensagem)
				fmt.Fprintln(conn, "OK")
			} else {
				fmt.Fprintln(conn, "Errroouuuu !!!!")
			}
			contador++

		} else {
			fmt.Fprintln(conn, "Ok...Adeus")
			os.Remove(breakline(NomeCliente) + ".dat")
			conn.Close()
			break
		}
	}
}
func imprimir(nome string, mensagem string) {
	fmt.Println(Bold(Green("[")), Gray(1-1, breakline(nome)).BgGray(24-1), Bold(Green("]:")), " ", Cyan(breakline(mensagem)))
}
func breakline(entrada string) string {
	return strings.Replace(entrada, "\n", "", -1)
}
func recebefloat(conn net.Conn) string {
	X, _ := bufio.NewReader(conn).ReadString('\n')
	return X
}
func InicializaEstruturaDiffie(InstanciaDiffieiliar *EstruturaDiffie) {
	rand.Seed(time.Now().UnixNano())
	InstanciaDiffieiliar.modulo.ParteReal = rand.Float64()
	InstanciaDiffieiliar.modulo.ParteImaginaria = rand.Float64()
}
func calculo(raiz, choice, modulo ComplexDiffie) (float64, float64) {
	//Faz terceira etapa do diffie
	InstanciaDiffie := cmplx.Pow(complex(raiz.ParteReal, raiz.ParteImaginaria), complex(choice.ParteReal, choice.ParteImaginaria))
	InstanciaDiffie = (InstanciaDiffie - (complex(choice.ParteReal, choice.ParteImaginaria) * (InstanciaDiffie / complex(choice.ParteReal, choice.ParteImaginaria))))
	return real(InstanciaDiffie), imag(InstanciaDiffie)
}
func calculo2(raiz, choice, choice2, modulo ComplexDiffie) (float64, float64) {
	//Faz terceira etapa do diffie
	InstanciaDiffie := cmplx.Pow(complex(raiz.ParteReal, raiz.ParteImaginaria), (complex(choice.ParteReal, choice.ParteImaginaria) * complex(choice2.ParteReal, choice2.ParteImaginaria)))
	InstanciaDiffie = (InstanciaDiffie - (complex(choice.ParteReal, choice.ParteImaginaria) * (InstanciaDiffie / complex(choice.ParteReal, choice.ParteImaginaria))))
	return real(InstanciaDiffie), imag(InstanciaDiffie)
}
func Md5EmptyHash(message string) string {
	h := md5.New()
	io.WriteString(h, message)
	return fmt.Sprintf("%x", h.Sum(nil))
}
func HMAC(MD5 string, secret []byte) string {
	h := hmac.New(sha256.New, []byte(secret))

	// Write Data to it
	h.Write([]byte(MD5))
	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}
func float64ToByte(f float64) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], math.Float64bits(f))
	return buf[:]
}
func CriaArquivo(arquivo string, MD5 string) {
	f, err := os.OpenFile(arquivo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(MD5 + "\n"); err != nil {
		log.Println(err)
	}
}
func VereficaArquivo(arquivo string, MD5 string) bool {

	file, _ := os.Open(arquivo)

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if scanner.Text() == MD5 {
			fmt.Println("arquivo: ", scanner.Text())
			fmt.Println(len(scanner.Text()))
			fmt.Println("servidor: ", MD5)
			fmt.Println(len(MD5))
			return true
		}
	}
	return false

}
