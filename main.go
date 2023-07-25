package main

import (
	"fmt"
	"os"

	"training.go/hangman/dictionary"
	"training.go/hangman/hangman"
)

func main() {
	err := dictionary.Load("words.txt")
	if err != nil {
		fmt.Printf("Could not load dictionnary: %v\n", err)
		os.Exit(1)
	}

	hangman.YourWelcome()

	g, err := hangman.New(8, dictionary.PickWord())

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	guess := ""
	for {
		hangman.Draw(g, guess)

		switch g.State {
		case "won", "lost":
			os.Exit(0)
		}

		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Printf("Could not read from input: %v", err)
			os.Exit(1)
		}
		guess = l
		g.MakeAGuess(guess)
	}
}