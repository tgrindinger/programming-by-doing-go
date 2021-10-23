package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func findNumbersLessThan56WithSumsOver10(writer io.Writer) {
	for i := 1; i <= 5; i++ {
		for j := 0; j <= 9; j++ {
			num := i * 10 + j
			if num <= 56 && i + j > 10 {
				fmt.Fprintf(writer, "%d\n", num)
			}
		}
	}
}

func findNumberWhereSubtractedReverseEqualsDigitSum(writer io.Writer) {
	for i := 1; i <= 5; i++ {
		for j := 0; j <= 9; j++ {
			num := i * 10 + j
			rev := j * 10 + i
			sum := i + j
			if num - rev == sum {
				fmt.Fprintf(writer, "%d\n", num)
			}
		}
	}
}

func PrintPuzzleNumbers(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	var choice int
	for ok := true; ok; ok = choice != 3 {
		fmt.Fprintln(writer, "1) Find two digit numbers <= 56 with sums of digits > 10")
		fmt.Fprintln(writer, "2) Find two digit number minus number reversed which equals sum of digits")
		fmt.Fprintln(writer, "3) Quit")
		fmt.Fprintln(writer)
		fmt.Fprint(writer, ">")
		scanner.Scan()
		choice, _ = strconv.Atoi(scanner.Text())
		if choice == 1 {
			fmt.Fprintln(writer)
			findNumbersLessThan56WithSumsOver10(writer)
			fmt.Fprintln(writer)
		} else if choice == 2 {
			fmt.Fprintln(writer)
			findNumberWhereSubtractedReverseEqualsDigitSum(writer)
			fmt.Fprintln(writer)
		}
	}
}

func main() {
	PrintPuzzleNumbers(os.Stdin, os.Stdout)
}
