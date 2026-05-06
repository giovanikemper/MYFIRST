package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

type Jogada struct {
	Escolha string `json:"escolha"`
}

type Resultado struct {
	Jogador    string `json:"jogador"`
	Computador string `json:"computador"`
	Resultado  string `json:"resultado"`
}

var opcoes = []string{"pedra", "papel", "tesoura"}

var emojis = map[string]string{
	"pedra":   "🪨",
	"papel":   "📄",
	"tesoura": "✂️",
}

func jogar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var jogada Jogada
	if err := json.NewDecoder(r.Body).Decode(&jogada); err != nil {
		http.Error(w, "jogada inválida", http.StatusBadRequest)
		return
	}

	computador := opcoes[rand.Intn(3)]
	jogador := jogada.Escolha

	resultado := "Empate! 🤝"
	if (jogador == "pedra" && computador == "tesoura") ||
		(jogador == "papel" && computador == "pedra") ||
		(jogador == "tesoura" && computador == "papel") {
		resultado = "Você ganhou! 🎉"
	} else if jogador != computador {
		resultado = "Você perdeu! 😢"
	}

	json.NewEncoder(w).Encode(Resultado{
		Jogador:    emojis[jogador] + " " + jogador,
		Computador: emojis[computador] + " " + computador,
		Resultado:  resultado,
	})
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/jogar", jogar)
	http.ListenAndServe(":3000", nil)
}
