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

func printTrick(writer io.Writer, game *Game) {
	fmt.Fprintln(writer, "Current trick:")
	for _, card := range game.trick {
		fmt.Fprintf(writer, "%s\n", card)
	}
	fmt.Fprintln(writer)
}

func printScores(writer io.Writer, game *Game) {
	for i, s := range game.scores {
		fmt.Fprintf(writer, "%d) %d\n", i, s)
	}
	fmt.Fprintln(writer)
}

func printWinner(writer io.Writer, game *Game) {
	lowest := 0
	for i, s := range game.scores {
		if s < game.scores[lowest] {
			lowest = i
		}
	}
	fmt.Fprintf(writer, "Game over. Player %d wins!\n", lowest)
}

func readChoice(scanner *bufio.Scanner, writer io.Writer, game *Game) int {
	var choice int
	for ok := true; ok; ok = !game.validPlay(choice - 1) {
		fmt.Fprint(writer, "Discard which card: ")
		scanner.Scan()
		choice, _ = strconv.Atoi(scanner.Text())
		if !game.validPlay(choice - 1) {
			fmt.Fprintln(writer, "You are unable to play that card.")
			fmt.Fprintln(writer, "You may only play cards of the same suit that has already been played.")
			fmt.Fprintln(writer, "If you have no matching cards, you may play any other card.")
			fmt.Fprintln(writer, "You may not break hearts as the first play unless you only have hearts.")
			fmt.Fprintln(writer)
		}
	}
	return choice - 1
}

func playTrick(scanner *bufio.Scanner, writer io.Writer, game *Game) {
	for !game.matchOver() {
		printTrick(writer, game)
		fmt.Fprintf(writer, "Player %d, your move. Current score: %d (%d)\n", game.current, game.players[game.current].score(), game.scores[game.current])
		fmt.Fprintln(writer, game.players[game.current].printHand())
		choice := readChoice(scanner, writer, game)
		game.playCard(choice)
		if game.trickOver() {
			fmt.Fprintf(writer, "Player %d wins the trick!\n", game.trickWinner())
			game.collectTrick()
		}
	}
}

func playMatch(scanner *bufio.Scanner, writer io.Writer, game *Game) {
	fmt.Fprintln(writer, "Dealing cards...")
	game.dealCards()
	fmt.Fprintln(writer)
	game.doFirstTurn()
	playTrick(scanner, writer, game)
	fmt.Fprintln(writer, "Match over. Scores:")
	game.updateScores()
	printScores(writer, game)
}

func playGame(random *rand.Rand, reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	game := NewGame(random)
	fmt.Fprintln(writer, "Let's play some Hearts!")
	fmt.Fprintln(writer)
	for !game.gameOver() {
		playMatch(scanner, writer, game)
	}
	printWinner(writer, game)
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	playGame(random, os.Stdin, os.Stdout)
}
