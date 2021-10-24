package main

import (
	"fmt"
	"io"
)

type TicTacToe struct {
	board [][]int
	turns int
}

func NewTicTacToe() *TicTacToe {
	board := make([][]int, 3)
	for i := 0; i < len(board); i++ {
		board[i] = make([]int, 3)
		for j := 0; j < len(board[i]); j++ {
			board[i][j] = ' '
		}
	}
	return &TicTacToe{board, 0}
}

func (t *TicTacToe) isWinner(p int) bool {
	winner := int(' ')
	if (t.board[0][0] == t.board[0][1] && t.board[0][0] == t.board[0][2]) ||
			(t.board[0][0] == t.board[1][0] && t.board[0][0] == t.board[2][0]) {
		winner = t.board[0][0]
	}
	if (t.board[2][2] == t.board[2][1] && t.board[2][2] == t.board[2][0]) ||
			(t.board[2][2] == t.board[1][2] && t.board[2][2] == t.board[0][2]) {
		winner = t.board[2][2]
	}
	if (t.board[1][1] == t.board[2][1] && t.board[1][1] == t.board[0][1]) ||
			(t.board[1][1] == t.board[1][0] && t.board[1][1] == t.board[1][2]) ||
			(t.board[1][1] == t.board[0][0] && t.board[1][1] == t.board[2][2]) ||
			(t.board[1][1] == t.board[0][2] && t.board[1][1] == t.board[2][0]) {
		winner = t.board[1][1]
	}
	return winner == p
}

func (t *TicTacToe) isFull() bool {
	return t.numTurns() == 9
}

func (t *TicTacToe) isCat() bool {
	return t.isFull() && t.isWinner(' ')
}

func (t *TicTacToe) isValid(r, c int) bool {
	return r >= 0 && r < len(t.board) &&
		c >= 0 && c < len(t.board[r])
}

func (t *TicTacToe) numTurns() int {
	return t.turns
}

func (t *TicTacToe) playerAt(r, c int) int {
	return t.board[r][c]
}

func (t *TicTacToe) displayBoard(writer io.Writer) {
	fmt.Fprintln(writer)
	for r := 0; r < len(t.board); r++ {
		fmt.Fprintf(writer, "\t%d ", r)
		for c := 0; c < len(t.board[r]); c++ {
			fmt.Fprintf(writer, "%c ", t.board[r][c])
		}
		fmt.Fprintln(writer)
	}
	fmt.Fprintln(writer, "\t  0 1 2 ")
	fmt.Fprintln(writer)
}

func (t *TicTacToe) playMove(p, r, c int) {
	t.board[r][c] = p
	t.turns++
}
