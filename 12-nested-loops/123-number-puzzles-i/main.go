package main

import (
	"fmt"
	"io"
	"os"
)

func PrintPuzzleNumbers(writer io.Writer) {
	for i := 10; i <= 60; i++ {
		for j := 10; j <= 60; j++ {
			if i + j == 60 && i - j == 14 {
				fmt.Fprintf(writer, "%d %d\n", i, j)
			}
		}
	}
}

func main() {
	PrintPuzzleNumbers(os.Stdout)
}
