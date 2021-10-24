package main

import (
	"fmt"
	"io"
	"os"
)

func initBoard(board [][]int) {
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			board[r][c] = ' '
		}
	}
}

func displayBoard(writer io.Writer, board [][]int) {
	for r := 0; r < 3; r++ {
		fmt.Fprintf(writer, "\t%d ", r)
		for c := 0; c < 3; c++ {
			fmt.Fprintf(writer, "%c ", board[r][c])
		}
		fmt.Fprintln(writer)
	}
	fmt.Fprintln(writer, "\t  0 1 2 ")
}

func winner(board [][]int) int {
	winner := int(' ')
	if (board[0][0] == board[0][1] && board[0][0] == board[0][2]) ||
			(board[0][0] == board[1][0] && board[0][0] == board[2][0]) {
		winner = board[0][0]
	}
	if (board[2][2] == board[2][1] && board[2][2] == board[2][0]) ||
			(board[2][2] == board[1][2] && board[2][2] == board[0][2]) {
		winner = board[2][2]
	}
	if (board[1][1] == board[2][1] && board[1][1] == board[0][1]) ||
			(board[1][1] == board[1][0] && board[1][1] == board[1][2]) ||
			(board[1][1] == board[0][0] && board[1][1] == board[2][2]) ||
			(board[1][1] == board[0][2] && board[1][1] == board[2][0]) {
		winner = board[1][1]
	}
	return winner
}

func gameOver(board [][]int) bool {
	filled := true
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if board[r][c] == ' ' {
				filled = false
			}
		}
	}
	return filled || winner(board) != ' '
}

func PlayGame(reader io.Reader, writer io.Writer) {
	board := make([][]int, 3)
	for i := 0; i < 3; i++ {
		board[i] = make([]int, 3)
	}
	initBoard(board)
	displayBoard(writer, board)
	var r, c int
	player := 'X'
	for ok := true; ok; ok = !gameOver(board) {
		fmt.Fprintf(writer, "'%c', choose your location (row, column): ", player)
		fmt.Fscanln(reader, &r, &c)
		board[r][c] = int(player)
		fmt.Fprintln(writer)
		displayBoard(writer, board)
		fmt.Fprintln(writer)
		if player == 'X' {
			player = 'O'
		} else {
			player = 'X'
		}
	}
	if winner(board) == ' ' {
		fmt.Fprintln(writer, "The game is a tie.")
	} else {
		fmt.Fprintf(writer, "The winner is %c!\n", winner(board))
	}
}

func main() {
	PlayGame(os.Stdin, os.Stdout)
}
