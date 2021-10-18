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
	fmt.Fprint(writer, "Your name: ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Fprint(writer, "Your age: ")
	scanner.Scan()
	age, _ := strconv.Atoi(scanner.Text())
	fmt.Fprintln(writer)
	if age < 16 {
		fmt.Fprintf(writer, "You can't drive, %s.\n", name)
	}
	if age >= 16 && age <= 17 {
		fmt.Fprintf(writer, "You can drive but not vote, %s.\n", name)
	}
	if age >= 18 && age <= 24 {
		fmt.Fprintf(writer, "You can vote but you can't rent a car, %s.\n", name)
	}
	if age >= 25 {
		fmt.Fprintf(writer, "You can do pretty much anything, %s.\n", name)
	}
}

func main() {
	AskQuestions(os.Stdin, os.Stdout)
}
