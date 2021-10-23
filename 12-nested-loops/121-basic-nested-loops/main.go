package main

import (
	"fmt"
	"io"
	"os"
)

func PrintCounter(writer io.Writer) {
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			fmt.Fprintf(writer, "(%d,%d) ", i, j)
		}
		fmt.Fprintln(writer)
	}
}

func main() {
	PrintCounter(os.Stdout)
}
