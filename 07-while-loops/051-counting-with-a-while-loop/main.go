package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Count(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprintln(writer, "Type in a message, and I'll display it several times.")
	fmt.Fprint(writer, "Message: ")
	scanner.Scan()
	msg := scanner.Text()
	fmt.Fprint(writer, "How many times? ")
	scanner.Scan()
	times, _ := strconv.Atoi(scanner.Text())
	for i := 10; i <= times * 10; i += 10 {
		fmt.Fprintf(writer, "%d. %s\n", i, msg)
	}
}

func main() {
	Count(os.Stdin, os.Stdout)
}
