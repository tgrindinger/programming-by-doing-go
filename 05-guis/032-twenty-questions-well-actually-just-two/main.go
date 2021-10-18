package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func AskQuestions(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprintln(writer, "TWO QUESTIONS!")
	fmt.Fprintln(writer, "Think of an object, and I'll try to guess it.")
	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "Question 1) Is it animal, vegetable, or mineral?")
	fmt.Fprint(writer, "> ")
	scanner.Scan()
	answer1 := scanner.Text()
	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "Question 2) Is it bigger than a breadbox?")
	fmt.Fprint(writer, "> ")
	scanner.Scan()
	answer2 := scanner.Text()
	fmt.Fprintln(writer)
	var guess string
	if answer1 == "animal" {
		if answer2 == "yes" {
			guess = "moose"
		} else {
			guess = "squirrel"
		}
	} else if answer1 == "vegetable" {
		if answer2 == "yes" {
			guess = "watermelon"
		} else {
			guess = "carrot"
		}
	} else {
		if answer2 == "yes" {
			guess = "Camaro"
		} else {
			guess = "paper clip"
		}
	}
	fmt.Fprintf(writer, "My guess is that you are thinking of a %s.\n", guess)
	fmt.Fprintln(writer, "I would ask you if I'm right, but I don't actually care.")
}

func main() {
	AskQuestions(os.Stdin, os.Stdout)
}
