package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func PlayNim(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	a := 3
	b := 3
	c := 3
	fmt.Fprintf(writer, "A: %d    B: %d    C: %d\n", a, b, c)
	fmt.Fprintln(writer)
	for ok := true; ok; ok = a > 0 || b > 0 || c > 0 {
		fmt.Fprint(writer, "Choose a pile: ")
		scanner.Scan()
		pile := scanner.Text()
		fmt.Fprintf(writer, "How many to remove from pile %s: ", pile)
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		if pile == "A" {
			a -= num
		} else if pile == "B" {
			b -= num
		} else if pile == "C" {
			c -= num
		}
		fmt.Fprintln(writer)
		fmt.Fprintf(writer, "A: %d    B: %d    C: %d\n", a, b, c)
		fmt.Fprintln(writer)
	}
	fmt.Fprintln(writer, "All piles are empty. Good job!")
}

func main() {
	PlayNim(os.Stdin, os.Stdout)
}
