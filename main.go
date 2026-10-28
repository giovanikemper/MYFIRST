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

	"github.com/jchv/go-webview2"
)

// A tela do jogo vai dentro do executável, então o .exe funciona sozinho.
//
//go:embed index.html
var paginaInicial []byte

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
	// Porta 0: o sistema escolhe uma livre, então dá para abrir o jogo várias vezes.
	ouvinte, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		fmt.Println("Não foi possível iniciar o jogo:", err)
		return
	}
	url := "http://" + ouvinte.Addr().String()

	http.HandleFunc("/", home)
	http.HandleFunc("/novo-jogo", novoJogo)
	go http.Serve(ouvinte, nil)

	janela := webview2.NewWithOptions(webview2.WebViewOptions{
		AutoFocus: true,
		WindowOptions: webview2.WindowOptions{
			Title:  "Jogo da Memória",
			Width:  600,
			Height: 780,
			Center: true,
		},
	})
	if janela == nil {
		// Sem o WebView2 instalado: usa o navegador como plano B.
		abrirNavegador(url)
		select {}
	}
	defer janela.Destroy()

	// Fechar a janela encerra o programa.
	janela.Navigate(url)
	janela.Run()
}
