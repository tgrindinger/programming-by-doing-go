package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func ProcessInput(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprint(writer, "Please enter your current earth weight: ")
	scanner.Scan()
	weight, _ := strconv.ParseFloat(scanner.Text(), 64)
	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "I have information for the following planets:")
	fmt.Fprintln(writer, "   1. Venus   2. Mars    3. Jupiter")
	fmt.Fprintln(writer, "   4. Saturn  5. Uranus  6. Neptune")
	fmt.Fprintln(writer)
	fmt.Fprint(writer, "Which planet are you visiting? ")
	scanner.Scan()
	planet, _ := strconv.Atoi(scanner.Text())
	fmt.Fprintln(writer)
	relativeWeight := weight
	if planet == 1 {
		relativeWeight = weight * 0.78
	} else if planet == 2 {
		relativeWeight = weight * 0.39
	} else if planet == 3 {
		relativeWeight = weight * 2.65
	} else if planet == 4 {
		relativeWeight = weight * 1.17
	} else if planet == 5 {
		relativeWeight = weight * 1.05
	} else if planet == 6 {
		relativeWeight = weight * 1.23
	}
	fmt.Fprintf(writer, "Your weight would be %.2f pounds on that planet.\n", relativeWeight)
}

func main() {
	ProcessInput(os.Stdin, os.Stdout)
}
