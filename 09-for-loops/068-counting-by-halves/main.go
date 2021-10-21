package main

import (
	"fmt"
	"io"
	"os"
)

func PrintCount(writer io.Writer) {
	fmt.Fprintln(writer, "x")
	fmt.Fprintln(writer, "------")
	for i := -10.0; i <= 10.0; i += 0.5 {
		fmt.Fprintf(writer, "%.1f\n", i)
	}
}

func main() {
	PrintCount(os.Stdout)
}
