package main

import (
	"fmt"
	"hangman-web/functions"
	"net/http"
)

const port = ":8080"

func main() {
	print(functions.HangmanStepLink(10))
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/api/", RequestDifficulty)
	http.HandleFunc("/", Home)
	http.HandleFunc("/api/hangman", Request)
	http.HandleFunc("/hangman", Play)

	functions.Openbrowser("http://localhost:8080")
	fmt.Println("Serveur lancé sur le port", port)
	http.ListenAndServe(port, nil)
}

// idée borne d'arcade

//TODO: win/lose affichage
//TODO: add a button to restart the game
