package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func PrintCollatz(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprint(writer, "Starting number: ")
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	steps := 0
	largest := num
	fmt.Fprintf(writer, "%d\t", num)
	for num > 1 {
		if num % 2 == 0 {
			num /= 2
		} else {
			num = 3 * num + 1
		}
		fmt.Fprintf(writer, "%d\t", num)
		steps++
		if num > largest {
			largest = num
		}
	}
	fmt.Fprintln(writer)
	fmt.Fprintln(writer)
	fmt.Fprintf(writer, "Terminated after %d steps. The largest value was %d.\n", steps, largest)
}

func main() {
	PrintCollatz(os.Stdin, os.Stdout)
}
