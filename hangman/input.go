package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadGuess() (guess string, err error) {
	valid := false
	for !valid {
		fmt.Print("What is you letter ? ")
		i, err := reader.ReadString('\n')
		guess = i
		if err != nil {
			return i, err
		}
		guess = strings.TrimSpace(guess)
		if len(guess) != 1 {
			fmt.Printf("Invalid input letter. Letter = %v, len = %v\n", guess, len(guess))
			continue
		}
		valid = true
	}
	return guess, nil
}