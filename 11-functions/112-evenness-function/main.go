package main

import (
	"fmt"
	"io"
	"os"
)

func isEven(n int) bool {
	return n % 2 == 0
}

func isDivisibleBy3(n int) bool {
	return n % 3 == 0
}

func PrintEvenness(writer io.Writer) {
	for i := 1; i <= 20; i++ {
		fmt.Fprintf(writer, "%d ", i)
		if isEven(i) {
			fmt.Fprint(writer, "<")
		}
		if isDivisibleBy3(i) {
			fmt.Fprint(writer, "=")
		}
		fmt.Fprintln(writer)
	}
}

func main() {
	PrintEvenness(os.Stdout)
}
