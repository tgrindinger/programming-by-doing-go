package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func areaCircle(radius int) float64 {
	return math.Pi * float64(radius * radius)
}

func areaRectangle(length, width int) int {
	return length * width
}

func areaSquare(side int) int {
	return side * side
}

func areaTriangle(base, height int) float64 {
	return 0.5 * float64(base * height)
}

func RunCalculator(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprintln(writer, "Shape Area Calculator version 0.1 (c) 2005 Mitchell Sample Output, Inc.")
	var choice int
	for ok := true; ok; ok = choice != 5 {
		fmt.Fprintln(writer)
		fmt.Fprintln(writer, "-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
		fmt.Fprintln(writer)
		fmt.Fprintln(writer, "1) Triangle")
		fmt.Fprintln(writer, "2) Rectangle")
		fmt.Fprintln(writer, "3) Square")
		fmt.Fprintln(writer, "4) Circle")
		fmt.Fprintln(writer, "5) Quit")
		fmt.Fprint(writer, "Which shape: ")
		scanner.Scan()
		choice, _ = strconv.Atoi(scanner.Text())
		fmt.Fprintln(writer)
		if choice == 1 {
			fmt.Fprint(writer, "Base: ")
			scanner.Scan()
			base, _ := strconv.Atoi(scanner.Text())
			fmt.Fprint(writer, "Height: ")
			scanner.Scan()
			height, _ := strconv.Atoi(scanner.Text())
			fmt.Fprintln(writer)
			fmt.Fprintf(writer, "The area is %f.\n", areaTriangle(base, height))
		} else if choice == 2 {
			fmt.Fprint(writer, "Length: ")
			scanner.Scan()
			length, _ := strconv.Atoi(scanner.Text())
			fmt.Fprint(writer, "Width: ")
			scanner.Scan()
			width, _ := strconv.Atoi(scanner.Text())
			fmt.Fprintln(writer)
			fmt.Fprintf(writer, "The area is %d.\n", areaRectangle(length, width))
		} else if choice == 3 {
			fmt.Fprint(writer, "Side length: ")
			scanner.Scan()
			length, _ := strconv.Atoi(scanner.Text())
			fmt.Fprintln(writer)
			fmt.Fprintf(writer, "The area is %d.\n", areaSquare(length))
		} else if choice == 4 {
			fmt.Fprint(writer, "Radius: ")
			scanner.Scan()
			radius, _ := strconv.Atoi(scanner.Text())
			fmt.Fprintln(writer)
			fmt.Fprintf(writer, "The area is %f.\n", areaCircle(radius))
		}
	}
	fmt.Fprintln(writer, "Goodbye.")
}

func main() {
	RunCalculator(os.Stdin, os.Stdout)
}
