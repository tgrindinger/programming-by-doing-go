package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func PrintMessage(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprint(writer, "What is your message? ")
	scanner.Scan()
	msg := scanner.Text()
	lastIndex := len(msg) - 1
	fmt.Fprintln(writer)
	fmt.Fprintf(writer, "Your message is %d characters long.\n", len(msg))
	fmt.Fprintf(writer, "The first character is at position %d and is '%c'.\n", 0, msg[0])
	fmt.Fprintf(writer, "The last character is at position %d and is '%c'.\n", lastIndex, msg[lastIndex])
	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "Here are all the characters, one at a time:")
	fmt.Fprintln(writer)
	aCount := 0
	for i := 0; i < len(msg); i++ {
		fmt.Fprintf(writer, "\t%d - '%c'\n", i, msg[i])
		if msg[i] == 'a' || msg[i] == 'A' {
			aCount++
		}
	}
	fmt.Fprintln(writer)
	fmt.Fprintf(writer, "Your message contains the letter 'a' %d times. Isn't that interesting?\n", aCount)
}

func main() {
	PrintMessage(os.Stdin, os.Stdout)
}
