package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func PrintName(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprint(writer, "What is your name: ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Fprintln(writer)
	count := 10
	if name == "Mitchell" {
		count = 5
	}
	for i := 0; i < count; i++ {
		fmt.Fprintln(writer, name)
	}
}

func main() {
	PrintName(os.Stdin, os.Stdout)
}
