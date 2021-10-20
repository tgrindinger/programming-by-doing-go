package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func CheckRightTriangle(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprintln(writer, "Enter three integers:")
	fmt.Fprint(writer, "Side 1: ")
	scanner.Scan()
	side1, _ := strconv.Atoi(scanner.Text())
	var side2, side3 int
	for ok := true; ok; ok = side2 < side1 {
		fmt.Fprint(writer, "Side 2: ")
		scanner.Scan()
		side2, _ = strconv.Atoi(scanner.Text())
		if side2 < side1 {
			fmt.Fprintf(writer, "%d is smaller than %d.  Try again.\n", side2, side1)
		}
	}
	for ok := true; ok; ok = side3 < side2 {
		fmt.Fprint(writer, "Side 3: ")
		scanner.Scan()
		side3, _ = strconv.Atoi(scanner.Text())
		if side3 < side2 {
			fmt.Fprintf(writer, "%d is smaller than %d.  Try again.\n", side3, side2)
		}
	}
	fmt.Fprint(writer)
	fmt.Fprintf(writer, "Your three sides are %d %d %d\n", side1, side2, side3)
	if side1 * side1 + side2 * side2 == side3 * side3 {
		fmt.Fprintln(writer, "These sides *do* make a right triangle.  Yippy-skippy!")
	} else {
		fmt.Fprintln(writer, "NO!  These sides do not make a right triangle!")
	}
}

func main() {
	CheckRightTriangle(os.Stdin, os.Stdout)
}
