package main

import (
	"fmt"
	"io"
	"os"
)

func AskQuestions(reader io.Reader, writer io.Writer) {
	var dump string
	fmt.Fprintln(writer, "Give me a word!")
	fmt.Fscanln(reader, &dump)
	fmt.Fprintln(writer, "Give me a second word!")
	fmt.Fscanln(reader, &dump)
	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "Great, now your favorite number?")
	fmt.Fscanln(reader, &dump)
	fmt.Fprintln(writer, "And your second-favorite number...")
	fmt.Fscanln(reader, &dump)
	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "Whew!  Wasn't that fun?")
}

func main() {
	AskQuestions(os.Stdin, os.Stdout)
}
