package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Futuro struct {
	Mensagem string `json:"mensagem"`
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func futuro(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())

	mensagens := []string{
		"Em breve você vai tomar uma decisão que muda tudo.",
		"Alguém inesperado vai cruzar seu caminho.",
		"Uma oportunidade grande está chegando.",
		"Você vai ganhar dinheiro de um jeito inesperado.",
	}

	msg := mensagens[rand.Intn(len(mensagens))]

	json.NewEncoder(w).Encode(Futuro{
		Mensagem: msg,
	})
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/futuro", futuro)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(":"+port, nil)
}
