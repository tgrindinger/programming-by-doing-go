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
	for _, card := range match.gameTrick() {
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

func readChoice(scanner *bufio.Scanner, writer io.Writer, match IMatch) int {
	var choice int
	for ok := true; ok; ok = !match.validPlay(choice - 1) {
		fmt.Fprint(writer, "Discard which card: ")
		scanner.Scan()
		choice, _ = strconv.Atoi(scanner.Text())
		if !match.validPlay(choice - 1) {
			fmt.Fprintln(writer, "You are unable to play that card.")
			fmt.Fprintln(writer, "You may only play cards of the same suit that has already been played.")
			fmt.Fprintln(writer, "If you have no matching cards, you may play any other card.")
			fmt.Fprintln(writer, "You may not break hearts as the first play unless you only have hearts.")
			fmt.Fprintln(writer)
		}
	}
	return choice - 1
}

func playTrick(scanner *bufio.Scanner, writer io.Writer, match IMatch) {
	for !match.matchOver() {
		printTrick(writer, match)
		fmt.Fprintf(writer, "Player %d, your move. Current score: %d\n", match.currentPlayer().index, match.currentPlayer().currentScore())
		fmt.Fprintln(writer, match.currentPlayer().printHand())
		choice := readChoice(scanner, writer, match)
		match.playCard(choice)
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
