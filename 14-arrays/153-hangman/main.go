package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func ShowWord(writer io.Writer, word string, revealed []bool) {
	fmt.Fprint(writer, "Word:   ")
	for i := 0; i < len(word); i++ {
		if revealed[i] {
			fmt.Fprintf(writer, "%c ", word[i])
		} else {
			fmt.Fprint(writer, "_ ")
		}
	}
	fmt.Fprintln(writer)
}

func ShowMisses(writer io.Writer, misses []byte) {
	fmt.Fprint(writer, "Misses: ")
	for i := 0; i < len(misses); i++ {
		fmt.Fprintf(writer, "%c", misses[i])
	}
	fmt.Fprintln(writer)
}

func containsGuess(word string, guess byte) bool {
	for i := 0; i < len(word); i++ {
		if word[i] == guess {
			return true
		}
	}
	return false
}

func revealWord(word string, revealed []bool, guess byte) {
	for i := 0; i < len(word); i++ {
		if word[i] == guess {
			revealed[i] = true
		}
	}
}

func gameOver(revealed []bool, misses []byte, missCount int) bool {
	if missCount == len(misses) {
		return true
	}
	for i := 0; i < len(revealed); i++ {
		if !revealed[i] {
			return false
		}
	}
	return true
}

func PlayGame(random *rand.Rand, scanner *bufio.Scanner, writer io.Writer) {
	words := []string { "leviathan", "porcupine", "alabama", "xylophone", "jaguar" }
	word := words[random.Intn(len(words))]
	revealed := make([]bool, len(word))
	misses := make([]byte, len(word))
	for i := 0; i < len(misses); i++ {
		misses[i] = ' '
	}
	missCount := 0

	for ok := true; ok; ok = !gameOver(revealed, misses, missCount) {
		fmt.Fprintln(writer, "-=-=-=-=-=-=-=-=-=-=-=-=-=-")
		fmt.Fprintln(writer)
		ShowWord(writer, word, revealed)
		fmt.Fprintln(writer)
		ShowMisses(writer, misses)
		fmt.Fprintln(writer)
		fmt.Fprint(writer, "Guess:  ")
		scanner.Scan()
		fmt.Fprintln(writer)
		guess := scanner.Text()[0]
		if containsGuess(word, guess) {
			revealWord(word, revealed, guess)
		} else {
			misses[missCount] = guess
			missCount++
		}
	}
	fmt.Fprintln(writer, "-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	fmt.Fprintln(writer)
	ShowWord(writer, word, revealed)
	fmt.Fprintln(writer)
	ShowMisses(writer, misses)
	fmt.Fprintln(writer)
	if missCount == len(misses) {
		fmt.Fprintln(writer, "Oops, you lost!")
	} else {
		fmt.Fprintln(writer, "YOU GOT IT!")
	}
	fmt.Fprintln(writer)
}

func PlayGames(random *rand.Rand, reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	var choice string
	for ok := true; ok; ok = choice == "again" {
		PlayGame(random, scanner, writer)
		fmt.Fprint(writer, "Play \"again\" or \"quit\"? ")
		scanner.Scan()
		choice = scanner.Text()
	}
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	PlayGames(random, os.Stdin, os.Stdout)
}
