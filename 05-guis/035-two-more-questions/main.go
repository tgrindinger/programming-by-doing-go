package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func AskQuestions(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprintln(writer, "TWO MORE QUESTIONS, BABY!")
	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "Think of something and I'll try to guess it!")
	fmt.Fprintln(writer)
	fmt.Fprint(writer, "Question 1) Does it stay inside or outside or both? ")
	scanner.Scan()
	location := scanner.Text()
	fmt.Fprint(writer, "Question 2) Is it a living thing? ")
	scanner.Scan()
	living := scanner.Text()
	fmt.Fprintln(writer)
	var answer string
	if location == "inside" && living == "yes" {
		answer = "houseplant"
	}
	if location == "inside" && living == "no" {
		answer = "shower curtain"
	}
	if location == "outside" && living == "yes" {
		answer = "bison"
	}
	if location == "outside" && living == "no" {
		answer = "billboard"
	}
	if location == "both" && living == "yes" {
		answer = "dog"
	}
	if location == "both" && living == "no" {
		answer = "cell phone"
	}
	fmt.Fprintf(writer, "Then what else could you be thinking of besides a %s?!?\n", answer)
}

func main() {
	AskQuestions(os.Stdin, os.Stdout)
}
