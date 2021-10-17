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
	fmt.Fprint(writer, "Hey, what's your name? (Sorry, I keep forgetting.) ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Fprintf(writer, "Ok, %s, how old are you? ", name)
	scanner.Scan()
	age, _ := strconv.Atoi(scanner.Text())
	fmt.Fprintln(writer)
	if age < 16 {
		fmt.Fprintf(writer, "You can't drive, %s.\n", name)
	} else if age < 18 {
		fmt.Fprintf(writer, "You can drive but not vote, %s.\n", name)
	} else if age < 25 {
		fmt.Fprintf(writer, "You can vote but not rent a car, %s.\n", name)
	} else {
		fmt.Fprintf(writer, "You can do pretty much anything, %s.\n", name)
	}
}

func main() {
	AskQuestions(os.Stdin, os.Stdout)
}
