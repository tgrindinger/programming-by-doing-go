package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func AskQuestion(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprint(writer, "What's your last name? ")
	scanner.Scan()
	name := scanner.Text()
	if strings.Compare(name, "Carswell") <= 0 {
		fmt.Fprintf(writer, "You don't have to wait long, \"%s\".\n", name)
	} else if strings.Compare(name, "Jones") <= 0 {
		fmt.Fprintf(writer, "That's not bad, \"%s\".\n", name)
	} else if strings.Compare(name, "Smith") <= 0 {
		fmt.Fprintf(writer, "Looks like a bit of a wait, \"%s\".\n", name)
	} else if strings.Compare(name, "Young") <= 0 {
		fmt.Fprintf(writer, "It's gonna be a while, \"%s\".\n", name)
	} else {
		fmt.Fprintf(writer, "Not going anywhere for a while, \"%s\"?\n", name)
	}
}

func main() {
	AskQuestion(os.Stdin, os.Stdout)
}
