package main

import (
    "encoding/json"
    "math/rand"
    "net/http"
    "time"
)

type Resultado struct {
    Chance int    `json:"chance"`
    Texto  string `json:"texto"`
}

func home(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.html")
}

func analisar(w http.ResponseWriter, r *http.Request) {
    rand.Seed(time.Now().UnixNano())

    chance := rand.Intn(100)

    texto := "Hmm... parece confiável 😎"
    if chance > 70 {
        texto = "⚠️ Alta chance de mentira 😂"
    } else if chance > 40 {
        texto = "🤨 Meio suspeito..."
    }

    json.NewEncoder(w).Encode(Resultado{
        Chance: chance,
        Texto:  texto,
    })
}

func main() {
    http.HandleFunc("/", home)
    http.HandleFunc("/analisar", analisar)

    http.ListenAndServe(":3000", nil)
}