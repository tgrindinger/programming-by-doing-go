package main

import (
	"fmt"
	"io"
	"os"
)

func PrintCount(writer io.Writer) {
	for j := 15; j <= 30; j += 3 {
		fmt.Fprintf(writer, "%d\n", j)
	}
}

func main() {
	PrintCount(os.Stdout)
}
