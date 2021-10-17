package main

import (
	"fmt"
	"io"
	"os"
)

func AskQuestions(reader io.Reader, writer io.Writer) {
	fmt.Fprintln(writer, "Hello.  What is your name?")
	var name, ageString, salaryString string
	fmt.Fscanln(reader, &name)
	fmt.Fprintln(writer)
	fmt.Fprintf(writer, "Hi, %s!  How old are you?\n", name)
	var age int
	fmt.Fscanln(reader, &ageString)
	fmt.Sscanf(ageString, "%d", &age)
	fmt.Fprintln(writer)
	fmt.Fprintf(writer, "So you're %d, eh?  That's not old at all!\n", age)
	fmt.Fprintf(writer, "How much do you make, %s?\n", name)
	var salary float64
	fmt.Fscanln(reader, &salaryString)
	fmt.Sscanf(salaryString, "%f", &salary)
	fmt.Fprintln(writer)
	fmt.Fprintf(writer, "%.1f!  I hope that's per hour and not per year! LOL!\n", salary)
}

func main() {
	AskQuestions(os.Stdin, os.Stdout)
}
