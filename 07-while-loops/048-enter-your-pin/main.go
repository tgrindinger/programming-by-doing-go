package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func GetPin(reader io.Reader, writer io.Writer) {
	pin := 12345
	scanner := bufio.NewScanner(reader)
	fmt.Fprintln(writer, "WELCOME TO THE BANK OF MITCHELL.")
	fmt.Fprint(writer, "ENTER YOUR PIN: ")
	scanner.Scan()
	entry, _ := strconv.Atoi(scanner.Text())
	for entry != pin {
		fmt.Fprintln(writer)
		fmt.Fprintln(writer, "INCORRECT PIN. TRY AGAIN.")
		fmt.Fprint(writer, "ENTER YOUR PIN: ")
		scanner.Scan()
		entry, _ = strconv.Atoi(scanner.Text())
	}
	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "PIN ACCEPTED. YOU NOW HAVE ACCESS TO YOUR ACCOUNT.")
}

func main() {
	GetPin(os.Stdin, os.Stdout)
}
