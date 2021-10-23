package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func PrintMultiples(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprint(writer, "Choose a number: ")
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	fmt.Fprintln(writer)
	for i := 1; i <= 12; i++ {
		fmt.Fprintf(writer, "%dx%d = %d\n", num, i, num * i)
	}
}

func main() {
	PrintMultiples(os.Stdin, os.Stdout)
}
