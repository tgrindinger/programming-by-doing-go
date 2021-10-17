package main

import (
	"fmt"
	"io"
	"os"
)

func AskForNumbers(reader io.Reader, writer io.Writer) {
	var firstString, secondString, thirdString string
	var first, second, third float64
	fmt.Fprintf(writer, "What is your first number? ")
	fmt.Fscanln(reader, &firstString)
	fmt.Sscanf(firstString, "%f", &first)
	fmt.Fprintf(writer, "What is your second number? ")
	fmt.Fscanln(reader, &secondString)
	fmt.Sscanf(secondString, "%f", &second)
	fmt.Fprintf(writer, "What is your third number? ")
	fmt.Fscanln(reader, &thirdString)
	fmt.Sscanf(thirdString, "%f", &third)
	fmt.Fprintln(writer)
	fmt.Fprintf(writer, "( %.1f + %.1f + %.1f ) / 2 is... %.1f\n", first, second, third, (first + second + third) / 2)
}

func main() {
	AskForNumbers(os.Stdin, os.Stdout)
}
