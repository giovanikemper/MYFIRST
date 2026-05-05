package main

import (
    "encoding/json"
    "net/http"
    "math/rand"
    "time"
    "os"
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

    texto := "Parece confiável 😎"
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

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    http.ListenAndServe(":"+port, nil)
}
