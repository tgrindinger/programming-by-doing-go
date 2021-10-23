package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func PrintSum(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprint(writer, "Which file would you like to read numbers from: ")
	scanner.Scan()
	filename := scanner.Text()
	fmt.Fprintf(writer, "Reading numbers from file %q", filename)
	fileReader, _ := os.Open(filename)
	fileScanner := bufio.NewScanner(fileReader)
	fileScanner.Scan()
	num1, _ := strconv.Atoi(fileScanner.Text())
	fileScanner.Scan()
	num2, _ := strconv.Atoi(fileScanner.Text())
	fileScanner.Scan()
	num3, _ := strconv.Atoi(fileScanner.Text())
	fmt.Fprintln(writer)
	fmt.Fprintf(writer, "%d + %d + %d = %d\n", num1, num2, num3, num1 + num2 + num3)
}

func main() {
	PrintSum(os.Stdin, os.Stdout)
}
