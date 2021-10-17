package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func AskQuestions(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprint(writer, "Hey, what's your name? ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Fscanln(reader, &name)
	fmt.Fprintln(writer)
	fmt.Fprintf(writer, "Ok, %s, how old are you? ", name)
	scanner.Scan()
	age, _ := strconv.Atoi(scanner.Text())
	fmt.Fprintln(writer)
	if age < 16 {
		fmt.Fprintf(writer, "You can't drive, %s.\n", name)
	}
	if age < 18 {
		fmt.Fprintf(writer, "You can't vote, %s.\n", name)
	}
	if age < 25 {
		fmt.Fprintf(writer, "You can't rent a car, %s.\n", name)
	}
	if age >= 25 {
		fmt.Fprintf(writer, "You can do anything that's legal, %s.\n", name)
	}
}

func main() {
	AskQuestions(os.Stdin, os.Stdout)
}