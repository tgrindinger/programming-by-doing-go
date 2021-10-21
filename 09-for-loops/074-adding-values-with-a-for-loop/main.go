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
	fmt.Fprint(writer, "Number: ")
	scanner.Scan()
	number, _ := strconv.Atoi(scanner.Text())
	fmt.Fprintln(writer)
	for i := 1; i <= number; i++ {
		fmt.Fprintf(writer, "%d ", i)
		total += i
	}
	fmt.Fprintln(writer)
	fmt.Fprintf(writer, "The sum is %d.\n", total)
}

func main() {
	SumNumbers(os.Stdin, os.Stdout)
}
