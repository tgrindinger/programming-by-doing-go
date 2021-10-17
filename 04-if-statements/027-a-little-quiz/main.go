package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func GiveQuiz(reader io.Reader, writer io.Writer) {
	score := 0
	scanner := bufio.NewScanner(reader)
	fmt.Fprint(writer, "Are you ready for a quiz?  ")
	scanner.Scan()
	fmt.Fprintln(writer, "Okay, here it comes!")
	fmt.Fprintln(writer)

	fmt.Fprintln(writer, "Q1) What is the capital of Alaska?")
	fmt.Fprintln(writer, "        1) Melbourne")
	fmt.Fprintln(writer, "        2) Anchorage")
	fmt.Fprintln(writer, "        3) Juneau")
	fmt.Fprintln(writer)
	fmt.Fprint(writer, "> ")
	scanner.Scan()
	response, _ := strconv.Atoi(scanner.Text())
	fmt.Fprintln(writer)
	if response == 3 {
		fmt.Fprintln(writer, "That's right!")
		score++
	} else {
		fmt.Fprintln(writer, "Sorry, the answer is Juneau.")
	}
	fmt.Fprintln(writer)

	fmt.Fprintln(writer, "Q2) Can you store the value \"cat\" in a variable of type int?")
	fmt.Fprintln(writer, "        1) yes")
	fmt.Fprintln(writer, "        2) no")
	fmt.Fprintln(writer)
	fmt.Fprint(writer, "> ")
	scanner.Scan()
	response, _ = strconv.Atoi(scanner.Text())
	fmt.Fprintln(writer)
	if response == 2 {
		fmt.Fprintln(writer, "That's right!")
		score++
	} else {
		fmt.Fprintln(writer, "Sorry, \"cat\" is a string. ints can only store numbers.")
	}
	fmt.Fprintln(writer)

	fmt.Fprintln(writer, "Q3) What is the result of 9+6/3?")
	fmt.Fprintln(writer, "        1) 5")
	fmt.Fprintln(writer, "        2) 11")
	fmt.Fprintln(writer, "        3) 15/3")
	fmt.Fprintln(writer)
	fmt.Fprint(writer, "> ")
	scanner.Scan()
	response, _ = strconv.Atoi(scanner.Text())
	fmt.Fprintln(writer)
	if response == 2 {
		fmt.Fprintln(writer, "That's right!")
		score++
	} else {
		fmt.Fprintln(writer, "Sorry, the answer is 11.")
	}
	fmt.Fprintln(writer)
	fmt.Fprintln(writer)
	fmt.Fprintf(writer, "Overall, you got %d out of 3 correct.\n", score)
	fmt.Fprintln(writer, "Thanks for playing!")
}

func main() {
	GiveQuiz(os.Stdin, os.Stdout)
}