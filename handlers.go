package main

import (
	"fmt"
	"hangman-web/functions"
	"net/http"
	"strings"
	"text/template"
)

var StartData functions.Infos

func Home(w http.ResponseWriter, r *http.Request) {
	StartData.Scores = functions.ReadScoreBoard()
	renderTemplate(w, "home", StartData)
}

func RequestDifficulty(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		renderTemplate(w, "home", nil)
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		// fmt.Fprintf(w, "%v", ScoreRank)
		// Setting up new game values by difficulty
		StartData = functions.NewGamePrep(r.FormValue("difficulty"))
		StartData.Scores = functions.ReadScoreBoard()
		StartData.Name = r.FormValue("username")
		StartData.WordToPrint = functions.WordToPrint(StartData.WordRune)
		http.Redirect(w, r, "/hangman", http.StatusSeeOther)
	default:
		fmt.Fprintf(w, "Only GET and POST")
	}
}

func Play(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "hangman", StartData)
}

func Request(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api/hangman" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		renderTemplate(w, "hangman", StartData)
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		StartData.Propositon = strings.ToUpper(r.FormValue("letter"))
		StartData = functions.Game(StartData)
		if StartData.WordToPrint == StartData.Word {
			StartData.Win = true
			StartData.Points = functions.Points(StartData)
			functions.Save(StartData)
		}
		renderTemplate(w, "hangman", StartData)
	default:
		fmt.Fprintf(w, "Only GET and POST")
	}
}

func renderTemplate(w http.ResponseWriter, file string, dataFiles any) {
	t, err := template.ParseFiles("./templates/" + file + ".gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, dataFiles)
}
