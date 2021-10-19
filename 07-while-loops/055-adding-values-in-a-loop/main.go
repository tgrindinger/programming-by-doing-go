package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func SumNumbers(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	total := 0
	fmt.Fprintln(writer, "I will add up the numbers you give me.")
	fmt.Fprint(writer, "Number: ")
	scanner.Scan()
	number, _ := strconv.Atoi(scanner.Text())
	for number != 0 {
		total += number
		fmt.Fprintf(writer, "The total so far is %d\n", total)
		fmt.Fprint(writer, "Number: ")
		scanner.Scan()
		number, _ = strconv.Atoi(scanner.Text())
	}
	fmt.Fprintln(writer)
	fmt.Fprintf(writer, "The total is %d.\n", total)
}

func main() {
	SumNumbers(os.Stdin, os.Stdout)
}
