# Trababalho de Segurança de Sistemas: Golang_client_and_server
Membros: Rafael Beltran e Daniel Temp

    Intruções de execução:
    Opção 1:
        Executar manualmente servidor e cliente:
            1º Para executar o servidor
            go run ~/Documents/Golang_client_and_server/server/server.go -ip 127.0.0.1 -port 8082

            2º
            obs: Para enviar mensagens manualmente
            go run ~/Documents/Golang_client_and_server/client/client.go -nome Alice -ip 127.0.0.1 -port 8082

            ou

            2º
            obs: para enviar mensagens aleatorias automaticamente
            go run ~/Documents/Golang_client_and_server/clientPseudoRandom/client.go -nome Alice -ip 127.0.0.1 -port 8082 -repeticoes 20 -tamanhoPalavra 10
    Opção 2:
        Executar automaticamente através dos arquivos shell
            1º Para executar o servidor
            ./server_run.sh

            2º Para executar o cliente manualmente. No diretorio client.
            ./client_run.sh

            ou

            2º Para executar o cliente automaticamente. No diretorio clientPseudoRandom
            ./client_run.sh

    Opção 3:
        Executar mais automaticamente ainda.
            1º Executar ./use.sh ira iniciar o servidor automaticamente depois o cliente.

            ou

            2º Executar ./useRandomStrings.sh ira iniciar o servidor automaticamente depois o cliente. O cliente iniciado enviar mensagens automaticamente para o servidor.

    Requisitos do SO:
    xterm

    Requisitos externos do código:
    "github.com/logrusorgru/aurora"
    "math/cmplx"
    "crypto/md5"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
