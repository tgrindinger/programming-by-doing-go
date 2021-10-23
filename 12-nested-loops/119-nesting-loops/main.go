package main

import (
	"fmt"
	"io"
	"os"
)

func PrintStuff(writer io.Writer) {
	for c := 'A'; c <= 'E'; c++ {
		for n := 1; n <= 3; n++ {
			fmt.Fprintf(writer, "%c %d\n", c, n)
		}
	}
	fmt.Fprintln(writer)
	fmt.Fprintln(writer)
	for a := 1; a <= 3; a++ {
		for b := 1; b <= 3; b++ {
			fmt.Fprintf(writer, "%d-%d ", a, b)
		}
	}
	fmt.Fprintln(writer)
}

func main() {
	PrintStuff(os.Stdout)
}
