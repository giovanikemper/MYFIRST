package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

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
	http.ServeFile(w, r, "index.html")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/novo-jogo", novoJogo)
	http.ListenAndServe(":3000", nil)
}
