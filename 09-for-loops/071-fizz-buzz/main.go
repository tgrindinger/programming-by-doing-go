package main

import (
	"fmt"
	"io"
	"os"
)

func PrintFizzBuzz(writer io.Writer) {
	for i := 1; i <= 100; i++ {
		if i % 15 == 0 {
			fmt.Fprintln(writer, "FizzBuzz")
		} else if i % 3 == 0 {
			fmt.Fprintln(writer, "Fizz")
		} else if i % 5 == 0 {
			fmt.Fprintln(writer, "Buzz")
		} else {
			fmt.Fprintf(writer, "%d\n", i)
		}
	}
}

func main() {
	PrintFizzBuzz(os.Stdout)
}
