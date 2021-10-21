package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func PlayBlackjack(random *rand.Rand, reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprintln(writer, "Welcome to Mitchell's blackjack program!")
	player1 := 2 + random.Intn(10)
	player2 := 2 + random.Intn(10)
	playerSum := player1 + player2
	dealer1 := 2 + random.Intn(10)
	dealer2 := 2 + random.Intn(10)
	dealerSum := dealer1 + dealer2
	fmt.Fprintf(writer, "You get a %d and a %d.\n", player1, player2)
	fmt.Fprintf(writer, "Your total is %d.\n", playerSum)
	fmt.Fprintln(writer)
	fmt.Fprintf(writer, "The dealer has a %d showing and a hidden card.\n", dealer1)
	fmt.Fprintln(writer, "His total is hidden, too.")
	fmt.Fprintln(writer)

	var action string
	for ok := true; ok; ok = playerSum <= 21 && action == "hit" {
		fmt.Fprint(writer, "Would you like to \"hit\" or \"stay\"? ")
		scanner.Scan()
		action = scanner.Text()
		if action == "hit" {
			card := 2 + random.Intn(10)
			playerSum += card
			fmt.Fprintf(writer, "You drew a %d.\n", card)
			fmt.Fprintf(writer, "Your total is %d.\n", playerSum)
			fmt.Fprintln(writer)
		}
	}
	if playerSum > 21 {
		fmt.Fprintln(writer, "YOU BUST!")
	}

	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "Okay, dealer's turn.")
	fmt.Fprintf(writer, "His hidden card was a %d.\n", dealer2)
	fmt.Fprintf(writer, "His total was %d.\n", dealerSum)
	fmt.Fprintln(writer)
	for playerSum <= 21 && dealerSum <= 21 && dealerSum < playerSum {
		fmt.Fprintln(writer, "Dealer chooses to hit.")
		card := 2 + random.Intn(10)
		dealerSum += card
		fmt.Fprintf(writer, "He draws a %d.\n", card)
		fmt.Fprintf(writer, "His total is %d.\n", dealerSum)
		fmt.Fprintln(writer)
	}
	if dealerSum > 21 {
		fmt.Fprintln(writer, "DEALER BUSTS!")
	} else {
		fmt.Fprintln(writer, "Dealer stays.")
	}

	fmt.Fprintln(writer)
	fmt.Fprintf(writer, "Dealer total is %d.\n", dealerSum)
	fmt.Fprintf(writer, "Your total is %d.\n", playerSum)
	fmt.Fprintln(writer)

	if playerSum <= 21 && (dealerSum > 21 || playerSum > dealerSum) {
		fmt.Fprintln(writer, "YOU WIN!")
	} else {
		fmt.Fprintln(writer, "YOU LOSE!")
	}
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	PlayBlackjack(random, os.Stdin, os.Stdout)
}
