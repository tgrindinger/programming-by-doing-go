package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func PrintMessage(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprintln(writer, "Type in a message, and I'll display it five times.")
	fmt.Fprint(writer, "Message: ")
	scanner.Scan()
	msg := scanner.Text()
	for i := 2; i <= 20; i += 2 {
		fmt.Fprintf(writer, "%d. %s\n", i, msg)
	}
}

func main() {
	PrintMessage(os.Stdin, os.Stdout)
}
