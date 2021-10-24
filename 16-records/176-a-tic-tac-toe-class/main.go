package main

import (
	"fmt"
	"io"
	"os"
)

func PlayGame(reader io.Reader, writer io.Writer) {
	ttt := NewTicTacToe()
	var r, c int
	p := int('X')
	for !(ttt.isWinner('X') || ttt.isWinner('O') || ttt.isFull()) {
		ttt.displayBoard(writer)
		fmt.Fprintf(writer, "'%c', choose your location (row, column): ", p)
		fmt.Fscanln(reader, &r, &c)
		for !ttt.isValid(r, c) || ttt.playerAt(r, c) != ' ' {
			if !ttt.isValid(r, c) {
				fmt.Fprintln(writer, "That is not a valid location. Try again.")
			} else if ttt.playerAt(r, c) != ' ' {
				fmt.Fprintln(writer, "That location is already full. Try again.")
			}
			fmt.Fprint(writer, "Choose your location (row, column): ")
			fmt.Fscanln(reader, &r, &c)
		}
		ttt.playMove(p, r, c)
		if p == 'X' {
			p = 'O'
		} else {
			p = 'X'
		}
	}
	ttt.displayBoard(writer)
	if ttt.isWinner('X') {
		fmt.Fprintf(writer, "X is the winner!\n")
	}
	if ttt.isWinner('O') {
		fmt.Fprintf(writer, "O is the winner!\n")
	}
	if ttt.isCat() {
		fmt.Fprintln(writer, "The game is a tie.")
	}
}

func main() {
	PlayGame(os.Stdin, os.Stdout)
}
