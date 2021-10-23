package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func PrintSum(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprint(writer, "Open which file: ")
	scanner.Scan()
	filename := scanner.Text()
	fmt.Fprintf(writer, "Reading numbers from file %q\n", filename)
	fileReader, _ := os.Open(filename)
	fileScanner := bufio.NewScanner(fileReader)
	for fileScanner.Scan() {
		fmt.Fprintln(writer, fileScanner.Text())
	}
}

func main() {
	PrintSum(os.Stdin, os.Stdout)
}
