package main

import (
	"fmt"
	"io"
	"os"
)

func PrintCount(writer io.Writer) {
	fmt.Fprintln(writer, " x\ty")
	fmt.Fprintln(writer, "----------------")
	for x := -10.0; x <= 10.0; x += 0.5 {
		y := x * x
		fmt.Fprintf(writer, "%.1f\t%.2f\n", x, y)
	}
	fmt.Fprintln(writer)
}

func main() {
	PrintCount(os.Stdout)
}
