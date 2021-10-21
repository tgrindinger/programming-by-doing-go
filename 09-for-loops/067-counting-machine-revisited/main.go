package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func PrintCount(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprint(writer, "Count from: ")
	scanner.Scan()
	from, _ := strconv.Atoi(scanner.Text())
	fmt.Fprint(writer, "Count to: ")
	scanner.Scan()
	to, _ := strconv.Atoi(scanner.Text())
	fmt.Fprint(writer, "Count by: ")
	scanner.Scan()
	by, _ := strconv.Atoi(scanner.Text())
	fmt.Fprintln(writer)
	for i := from; i <= to; i += by {
		fmt.Fprintf(writer, "%d ", i)
	}
	fmt.Fprintln(writer)
}

func main() {
	PrintCount(os.Stdin, os.Stdout)
}
