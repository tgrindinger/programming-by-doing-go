package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func PrintSquareRoot(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprintln(writer, "SQUARE ROOT!")
	num := 0
	for ok := true; ok; ok = num < 0 {
		if num < 0 {
			fmt.Fprintln(writer, "You can't take the square root of a negative number, silly.")
			fmt.Fprint(writer, "Try again: ")
		} else {
			fmt.Fprint(writer, "Enter a number: ")
		}
		scanner.Scan()
		num, _ = strconv.Atoi(scanner.Text())
	}
	fmt.Fprintf(writer, "The square root of %d is %f.\n", num, math.Sqrt(float64(num)))
}

func main() {
	PrintSquareRoot(os.Stdin, os.Stdout)
}
