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
	fmt.Fprint(writer, "Count to: ")
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	for i := 0; i <= num; i++ {
		fmt.Fprintf(writer, "%d ", i)
	}
	fmt.Fprintln(writer)
}

func main() {
	PrintCount(os.Stdin, os.Stdout)
}
