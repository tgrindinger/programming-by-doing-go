package main

import (
	"fmt"
	"io"
	"os"
)

func AskQuestions(reader io.Reader, writer io.Writer) {
	var name, ageString string
	var age int
	fmt.Fprint(writer, "Hello.  What is your name? ")
	fmt.Fscanln(reader, &name)
	fmt.Fprintln(writer)
	fmt.Fprintf(writer, "Hi, %s!  How old are you? ", name)
	fmt.Fscanln(reader, &ageString)
	fmt.Sscanf(ageString, "%d", &age)
	fmt.Fprintln(writer)
	fmt.Fprintf(writer, "Did you know that in five years you will be %d years old?\n", age + 5)
	fmt.Fprintf(writer, "And five years ago you were %d! Imagine that!\n", age - 5)
}

func main() {
	AskQuestions(os.Stdin, os.Stdout)
}
