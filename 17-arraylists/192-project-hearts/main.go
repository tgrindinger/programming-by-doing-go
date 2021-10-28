package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func printTrick(writer io.Writer, match IMatch) {
	fmt.Fprintln(writer, "Current trick:")
	for _, card := range match.matchTrick() {
		fmt.Fprintf(writer, "%s\n", card)
	}
	fmt.Fprintln(writer)
}

func printScores(writer io.Writer, game IGame) {
	for i, s := range game.gameScores() {
		fmt.Fprintf(writer, "%d) %d\n", i, s)
	}
	fmt.Fprintln(writer)
}

func playTurn(scanner *bufio.Scanner, writer io.Writer, match IMatch) {
	var choice int
	var err error
	for ok := true; ok; ok = err != nil {
		fmt.Fprint(writer, "Discard which card: ")
		scanner.Scan()
		choice, _ = strconv.Atoi(scanner.Text())
		err = match.playCard(choice - 1)
		switch err {
		case ErrFirstPlay:     fmt.Fprintln(writer, "You cannot break hearts on the first play of a trick if you have any other suits in your hand.")
		case ErrNextPlay:      fmt.Fprintln(writer, "You must play a card of the same suit as the first card of the trick.")
		case ErrInvalidChoice: fmt.Fprintln(writer, "That's not a choice!")
		}
		fmt.Fprintln(writer)
	}
}

func playTrick(scanner *bufio.Scanner, writer io.Writer, match IMatch) {
	for !match.matchOver() {
		printTrick(writer, match)
		fmt.Fprintf(writer, "Player %d, your move. Current score: %d\n", match.currentPlayer().index, match.currentPlayer().currentScore())
		fmt.Fprintln(writer, match.currentPlayer().printHand())
		playTurn(scanner, writer, match)
		if match.trickOver() {
			fmt.Fprintf(writer, "Player %d wins the trick!\n", match.trickWinner())
			match.collectTrick()
		}
	}
}

func playMatch(scanner *bufio.Scanner, writer io.Writer, match IMatch) {
	fmt.Fprintln(writer, "Dealing cards...")
	match.dealCards()
	fmt.Fprintln(writer)
	match.doFirstTurn()
	playTrick(scanner, writer, match)
}

func playGame(random *rand.Rand, reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	game := NewGame(random)
	fmt.Fprintln(writer, "Let's play some Hearts!")
	fmt.Fprintln(writer)
	for !game.gameOver() {
		playMatch(scanner, writer, game.newMatch())
		fmt.Fprintln(writer, "Match over. Scores:")
		game.updateScores()
		printScores(writer, game)
	}
	fmt.Fprintf(writer, "Game over. Player %d wins!\n", game.gameWinner())
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	playGame(random, os.Stdin, os.Stdout)
}
