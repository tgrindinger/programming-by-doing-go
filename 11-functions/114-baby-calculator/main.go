package main

import (
	"fmt"
	"io"
	"os"
)

func RunCalculator(reader io.Reader, writer io.Writer) {
	var a, b float64
	for ok := true; ok; ok = a != 0 {
		var op string
		fmt.Fprint(writer, "> ")
		fmt.Fscanln(reader, &a, &op, &b)
		if a != 0 {
			var c float64
			if op == "+" {
				c = a + b
			} else if op == "-" {
				c = a - b
			} else if op == "*" {
				c = a * b
			} else if op == "/" {
				c = a / b
			}
			fmt.Fprintf(writer, "%f\n", c)
		}
	}
	fmt.Fprintln(writer, "Bye, now.")
}

func main() {
	RunCalculator(os.Stdin, os.Stdout)
}
