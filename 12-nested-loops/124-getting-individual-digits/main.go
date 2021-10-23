package main

import (
	"fmt"
	"io"
	"os"
)

func PrintPuzzleNumbers(writer io.Writer) {
	for i := 1; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			fmt.Fprintf(writer, "%d%d, %d+%d = %d\n", i, j, i, j, i + j)
		}
	}
}

func main() {
	PrintPuzzleNumbers(os.Stdout)
}
