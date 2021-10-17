package main

import (
	"fmt"
	"io"
	"os"
)

func AskQuestions(reader io.Reader, writer io.Writer) {
	fmt.Fprintf(writer, "How old are you? ")
	var ageString, height, weightString string
	var age int
	var weight float64
	fmt.Fscanln(reader, &ageString)
	fmt.Sscanf(ageString, "%d", &age)
	fmt.Fprintf(writer, "How tall are you? ")
	fmt.Fscanln(reader, &height)
	fmt.Fprintf(writer, "How much do you weigh? ")
	fmt.Fscanln(reader, &weightString)
	fmt.Sscanf(weightString, "%f", &weight)
	fmt.Fprintf(writer, "So you're %d old, %s tall and %.0f heavy.", age, height, weight)
}

func main() {
	AskQuestions(os.Stdin, os.Stdout)
}