package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func PlayNim(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprint(writer, "Player 1, enter your name: ")
	scanner.Scan()
	player1 := scanner.Text()
	fmt.Fprint(writer, "Player 2, enter your name: ")
	scanner.Scan()
	player2 := scanner.Text()
	player := player1
	a := 3
	b := 4
	c := 5
	fmt.Fprintf(writer, "A: %d    B: %d    C: %d\n", a, b, c)
	fmt.Fprintln(writer)
	for ok := true; ok; ok = a + b + c > 1 {
		fmt.Fprintf(writer, "%s, choose a pile: ", player)
		pile := ""
		for pile == "" {
			scanner.Scan()
			pile = scanner.Text()
			if (pile == "A" && a == 0) || 
					(pile == "B" && b == 0) ||
					(pile == "C" && c == 0) {
				fmt.Fprintln(writer)
				fmt.Fprintf(writer, "Nice try, %s. That pile is empty. Choose again: ", player)
				pile = ""
			}
		}
		fmt.Fprintf(writer, "How many to remove from pile %s: ", pile)
		num := 0
		for num == 0 {
			scanner.Scan()
			num, _ = strconv.Atoi(scanner.Text())
			if (pile == "A" && a < num) ||
					(pile == "B" && b < num) ||
					(pile == "C" && c < num) {
				fmt.Fprintln(writer)
				fmt.Fprintf(writer, "Pile %s doesn't have that many. Try again: ", pile)
				num = 0
			}
		}
		if pile == "A" {
			a -= num
		} else if pile == "B" {
			b -= num
		} else if pile == "C" {
			c -= num
		}
		fmt.Fprintln(writer)
		fmt.Fprintf(writer, "A: %d    B: %d    C: %d\n", a, b, c)
		fmt.Fprintln(writer)
		if player == player1 {
			player = player2
		} else {
			player = player1
		}
	}
	if a + b + c == 1 {
		var otherPlayer string
		if player == player1 {
			otherPlayer = player2
		} else {
			otherPlayer = player1
		}
		fmt.Fprintf(writer, "%s, you must take the last remaining counter, so you lose. %s wins!\n", player, otherPlayer)
	} else {
		fmt.Fprintf(writer, "%s, there are no counters left, so you WIN!\n", player)
	}
}

func main() {
	PlayNim(os.Stdin, os.Stdout)
}
