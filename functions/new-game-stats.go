package functions

import game_function "hangman-classic/g-func"

type Infos struct {
	Word            string
	WordRune        []rune
	WordToPrint     string
	Propositon      string
	LetterSuggested []string
	Lives           int
	Hangman         string
	Name            string
	Difficulty      string
	Points          int
	Win             bool
	WinLose         string
}

func NewGamePrep(Difficulty string) (StartData Infos) {
	Word := ""
	var WordRune []rune
	switch Difficulty {
	case "easy":
		Word, WordRune = game_function.NewGamePrep([]string{"words.txt"})
	case "medium":
		Word, WordRune = game_function.NewGamePrep([]string{"words2.txt"})
	case "hard":
		Word, WordRune = game_function.NewGamePrep([]string{"words3.txt"})
	}

	StartData = Infos{
		Word:            Word,
		WordRune:        WordRune,
		WordToPrint:     WordToPrint(WordRune),
		Propositon:      "",
		LetterSuggested: []string{},
		Lives:           10,
		Hangman:         PrintHangMan(10),
		Name:            "",
		Difficulty:      Difficulty,
		Points:          0,
		Win:             false,
		WinLose:         "",
	}
	return StartData
}
