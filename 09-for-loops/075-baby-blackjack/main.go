package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func PlayBlackjack(random *rand.Rand, writer io.Writer) {
	fmt.Fprintln(writer, "Baby Blackjack!")
	fmt.Fprintln(writer)
	player1 := 1 + random.Intn(10)
	player2 := 1 + random.Intn(10)
	playerSum := player1 + player2
	dealer1 := 1 + random.Intn(10)
	dealer2 := 1 + random.Intn(10)
	dealerSum := dealer1 + dealer2
	fmt.Fprintf(writer, "You drew %d and %d.\n", player1, player2)
	fmt.Fprintf(writer, "Your total is %d.\n", playerSum)
	fmt.Fprintln(writer)
	fmt.Fprintf(writer, "The dealer has %d and %d.\n", dealer1, dealer2)
	fmt.Fprintf(writer, "Dealer's total is %d.\n", dealerSum)
	fmt.Fprintln(writer)
	if playerSum > dealerSum {
		fmt.Fprintln(writer, "YOU WIN!")
	} else {
		fmt.Fprintln(writer, "YOU LOSE!")
	}
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	PlayBlackjack(random, os.Stdout)
}
