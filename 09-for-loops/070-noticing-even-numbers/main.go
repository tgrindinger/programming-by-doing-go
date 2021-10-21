package main

import (
	"fmt"
	"io"
	"os"
)

func PrintEvenNumbers(writer io.Writer) {
	for i := 1; i <= 20; i++ {
		if i % 2 == 0 {
			fmt.Fprintf(writer, "%d <\n", i)
		} else {
			fmt.Fprintf(writer, "%d\n", i)
		}
	}
}

func main() {
	PrintEvenNumbers(os.Stdout)
}
