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
	"math"
	"math/cmplx"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

type ArgumentosCliente struct {
	nomeCliente    string
	ip             string
	port           string
	repeticoes     int
	tamanhoPalavra int
}
type ComplexDiffie struct {
	ParteReal       float64
	ParteImaginaria float64
}
type EstruturaDiffie struct {
	choice ComplexDiffie
	raiz   ComplexDiffie
	modulo ComplexDiffie
	A      ComplexDiffie
	B      ComplexDiffie
	key    ComplexDiffie
}

func InicializaArgs(dados *ArgumentosCliente) {
	flag.StringVar(&dados.nomeCliente, "nome", "", "Aqui voce digita -nome <nomeDoCliente>. Ex Alice")
	flag.StringVar(&dados.ip, "ip", "", "Aqui voce digita -ip <ip>. Ex: -ip 127.0.0.1")
	flag.StringVar(&dados.port, "port", "", "Aqui voce digita -port <porta>. Ex -port 8033")
	flag.IntVar(&dados.repeticoes, "repeticoes", 10, "Aqui voce digita -repeticoes <numero inteiro>. Ex -repeticoes 100")
	flag.IntVar(&dados.tamanhoPalavra, "tamanhoPalavra", 10, "Aqui voce digita -tamanhoPalavra <numero inteiro>. Ex -tamanhoPalavra 10")
	flag.Parse()
}
func InicializaEstruturaDiffie(auxiliar *EstruturaDiffie) {
	rand.Seed(time.Now().UnixNano() * 1000)
	//Etapa1
	auxiliar.modulo.ParteReal = rand.Float64()
	auxiliar.modulo.ParteImaginaria = rand.Float64()
	auxiliar.raiz.ParteReal = rand.Float64()
	auxiliar.raiz.ParteImaginaria = rand.Float64()
	//Etapa2
	auxiliar.choice.ParteReal = rand.Float64()
	auxiliar.choice.ParteImaginaria = rand.Float64()
	//Etapa3
	auxiliar.A.ParteReal, auxiliar.A.ParteImaginaria = calculo(auxiliar.raiz, auxiliar.choice, auxiliar.modulo)

}

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func main() {
	var contador = 0
	var ClienteInstancia ArgumentosCliente
	InicializaArgs(&ClienteInstancia)
	var InstanciaDiffie EstruturaDiffie
	InicializaEstruturaDiffie(&InstanciaDiffie)

	conn, _ := net.Dial("tcp", ClienteInstancia.ip+":"+ClienteInstancia.port)
	//Etapa1
	fmt.Fprintln(conn, InstanciaDiffie.choice.ParteReal)
	time.Sleep(100 * time.Millisecond)
	fmt.Fprintln(conn, InstanciaDiffie.choice.ParteImaginaria)
	time.Sleep(100 * time.Millisecond)
	//Etapa2
	fmt.Fprintln(conn, InstanciaDiffie.raiz.ParteReal)
	time.Sleep(100 * time.Millisecond)
	fmt.Fprintln(conn, InstanciaDiffie.raiz.ParteImaginaria)
	time.Sleep(100 * time.Millisecond)
	//Etapa4
	fmt.Fprintln(conn, InstanciaDiffie.A.ParteReal)
	time.Sleep(100 * time.Millisecond)
	fmt.Fprintln(conn, InstanciaDiffie.A.ParteImaginaria)
	//Etapa5

	InstanciaDiffie.B.ParteReal, _ = strconv.ParseFloat(breakline(recebefloat(conn)), 64)
	InstanciaDiffie.B.ParteImaginaria, _ = strconv.ParseFloat(breakline(recebefloat(conn)), 64)

	//Etapa6
	InstanciaDiffie.key.ParteReal, InstanciaDiffie.key.ParteImaginaria = calculo2(InstanciaDiffie.raiz, InstanciaDiffie.B, InstanciaDiffie.A, InstanciaDiffie.modulo)
	fmt.Println("key: ", InstanciaDiffie.key)

	for i := 0; i < ClienteInstancia.repeticoes; i++ {
		fmt.Printf("Digite aqui: ")
		//reader := bufio.NewReader(os.Stdin)

		sendMessage(StringWithCharset(ClienteInstancia.tamanhoPalavra, charset)+"\n", conn, ClienteInstancia.nomeCliente, contador, float64ToByte(InstanciaDiffie.key.ParteImaginaria))
		contador++
		MenssagemDeConfirmacao, _ := bufio.NewReader(conn).ReadString('\n')
		if breakline(MenssagemDeConfirmacao) != "Ok...Adeus" {
			fmt.Println(breakline(MenssagemDeConfirmacao))
		} else {
			conn.Close()
			break
		}
	}
	sendMessage("quit\n", conn, ClienteInstancia.nomeCliente, contador, float64ToByte(InstanciaDiffie.key.ParteImaginaria))
}

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func sendMessage(mensagem string, conn net.Conn, nome_cliente string, contador int, secret []byte) {
	//Funcao para enviar para servidor {Nome, Mensagem, HMAC(MD5(Nome+Mensagem+contador),key)}
	fmt.Fprintln(conn, nome_cliente)
	time.Sleep(100 * time.Millisecond)
	//text, _ := r.ReadString('\n')
	text := mensagem
	fmt.Fprintf(conn, text)
	time.Sleep(100 * time.Millisecond)
	fmt.Fprintln(conn, HMAC(Md5EmptyHash(nome_cliente+text+strconv.Itoa(contador)), secret))
	//fmt.Println(HMAC(Md5EmptyHash(nome_cliente+text+strconv.Itoa(contador)), secret))

}
func calculo(raiz, choice, modulo ComplexDiffie) (float64, float64) {
	//Faz terceira etapa do diffie
	aux := cmplx.Pow(complex(raiz.ParteReal, raiz.ParteImaginaria), complex(choice.ParteReal, choice.ParteImaginaria))
	aux = (aux - (complex(choice.ParteReal, choice.ParteImaginaria) * (aux / complex(choice.ParteReal, choice.ParteImaginaria))))
	return real(aux), imag(aux)
}
func calculo2(raiz, choice, choice2, modulo ComplexDiffie) (float64, float64) {
	//Faz sexta etapa do Diffie
	aux := cmplx.Pow(complex(raiz.ParteReal, raiz.ParteImaginaria), (complex(choice.ParteReal, choice.ParteImaginaria) * complex(choice2.ParteReal, choice2.ParteImaginaria)))
	aux = (aux - (complex(choice.ParteReal, choice.ParteImaginaria) * (aux / complex(choice.ParteReal, choice.ParteImaginaria))))
	return real(aux), imag(aux)
}
func breakline(entrada string) string {
	//Remove quebra linha da mensagem pois eh um caracter invisivel
	return strings.Replace(entrada, "\n", "", -1)
}
func recebefloat(conn net.Conn) string {
	//Recebe parte real ou imaginaria do numero
	X, _ := bufio.NewReader(conn).ReadString('\n')
	return X
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
