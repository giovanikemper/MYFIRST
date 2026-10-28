package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"os/exec"
	"runtime"
)

// A tela do jogo vai dentro do executável, então o .exe funciona sozinho.
//
//go:embed index.html
var paginaInicial []byte

const endereco = "localhost:3000"

type Tabuleiro struct {
	Cartas []string `json:"cartas"`
}

var emojis = []string{"🐶", "🐱", "🐸", "🦊", "🐼", "🐨", "🦁", "🐯"}

func novoJogo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pares := append(emojis, emojis...)
	rand.Shuffle(len(pares), func(i, j int) {
		pares[i], pares[j] = pares[j], pares[i]
	})

	json.NewEncoder(w).Encode(Tabuleiro{Cartas: pares})
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(paginaInicial)
}

func abrirNavegador(url string) {
	switch runtime.GOOS {
	case "windows":
		exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		exec.Command("open", url).Start()
	default:
		exec.Command("xdg-open", url).Start()
	}
}

func main() {
	url := "http://" + endereco

	ouvinte, err := net.Listen("tcp", endereco)
	if err != nil {
		// O jogo já está aberto em outra janela: só abre o navegador nele.
		abrirNavegador(url)
		return
	}

	http.HandleFunc("/", home)
	http.HandleFunc("/novo-jogo", novoJogo)

	fmt.Println("Jogo da Memória rodando em", url)
	fmt.Println("Feche esta janela para encerrar o jogo.")
	abrirNavegador(url)
	http.Serve(ouvinte, nil)
}
