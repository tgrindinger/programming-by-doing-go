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
	fmt.Fprint(writer, "Reading numbers from file \"3nums.txt\"... ")
	scanner.Scan()
	num1, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	num2, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	num3, _ := strconv.Atoi(scanner.Text())
	fmt.Fprintln(writer, "done.")
	fmt.Fprintln(writer)
	fmt.Fprintf(writer, "%d + %d + %d = %d\n", num1, num2, num3, num1 + num2 + num3)
}

func main() {
	reader, _ := os.Open("3nums.txt")
	PrintSum(reader, os.Stdout)
	reader.Close()
}
