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
	tries := 1
	scanner := bufio.NewScanner(reader)
	fmt.Fprintln(writer, "WELCOME TO THE BANK OF MITCHELL.")
	fmt.Fprint(writer, "ENTER YOUR PIN: ")
	scanner.Scan()
	entry, _ := strconv.Atoi(scanner.Text())
	for entry != pin && tries < 3 {
		fmt.Fprintln(writer)
		fmt.Fprintln(writer, "INCORRECT PIN. TRY AGAIN.")
		fmt.Fprint(writer, "ENTER YOUR PIN: ")
		scanner.Scan()
		entry, _ = strconv.Atoi(scanner.Text())
		tries++
	}
	fmt.Fprintln(writer)
	if entry == pin {
		fmt.Fprintln(writer, "PIN ACCEPTED. YOU NOW HAVE ACCESS TO YOUR ACCOUNT.")
	} else {
		fmt.Fprintln(writer, "YOU HAVE RUN OUT OF TRIES. ACCOUNT LOCKED.")
	}
}

func main() {
	GetPin(os.Stdin, os.Stdout)
}
