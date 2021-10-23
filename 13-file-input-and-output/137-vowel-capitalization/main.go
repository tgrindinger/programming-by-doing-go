package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CapitalizeVowels(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprint(writer, "Open which file: ")
	scanner.Scan()
	filename := scanner.Text()
	fmt.Fprintln(writer)
	file, _ := os.Open(filename)
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	success := fileScanner.Scan()
	for success {
		for _, c := range fileScanner.Text() {
			if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
				c -= 'a'
				c += 'A'
			}
			fmt.Fprintf(writer, "%c", c)
		}
		fmt.Fprintln(writer)
		success = fileScanner.Scan()
	}
	fmt.Fprintln(writer)
}

func main() {
	CapitalizeVowels(os.Stdin, os.Stdout)
}
