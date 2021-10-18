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
	fmt.Fprint(writer, "What is your gender (M or F): ")
	scanner.Scan()
	gender := scanner.Text()
	fmt.Fprint(writer, "First name: ")
	scanner.Scan()
	firstname := scanner.Text()
	fmt.Fprint(writer, "Last name: ")
	scanner.Scan()
	lastname := scanner.Text()
	fmt.Fprint(writer, "Age: ")
	scanner.Scan()
	age, _ := strconv.Atoi(scanner.Text())
	fmt.Fprintln(writer)
	var prefix string
	if age < 20 {
		prefix = firstname
	} else if gender == "M" {
		prefix = "Mr."
	} else {
		fmt.Fprintf(writer, "Are you married, %s (y or n)? ", firstname)
		scanner.Scan()
		fmt.Fprintln(writer)
		married := scanner.Text()
		if married == "y" {
			prefix = "Mrs."
		} else {
			prefix = "Ms."
		}
	}
	fmt.Fprintf(writer, "Then I shall call you %s %s.\n", prefix, lastname)
}

func main() {
	AskQuestions(os.Stdin, os.Stdout)
}