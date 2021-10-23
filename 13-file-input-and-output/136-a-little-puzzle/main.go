package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func PrintWebLine(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprint(writer, "Open which file: ")
	scanner.Scan()
	filename := scanner.Text()
	fmt.Fprintln(writer)
	file, _ := os.Open(filename)
	defer file.Close()
	var c int
	numRead, _ := fmt.Fscanf(file, "%c", &c)
	i := 1
	for numRead == 1 {
		if i % 3 == 0 {
			fmt.Fprintf(writer, "%c", c)
		}
		i++
		numRead, _ = fmt.Fscanf(file, "%c", &c)
	}
	fmt.Fprintln(writer)
}

func main() {
	PrintWebLine(os.Stdin, os.Stdout)
}
