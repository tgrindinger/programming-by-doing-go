package main

import (
	"fmt"
	"io"
	"os"
)

func PrintMessage(writer io.Writer) {
	for i := 1; i <= 10; i++ {
		fmt.Fprintf(writer, "%d. Mr. Mitchell is cool.\n", i)
	}
}

func main() {
	PrintMessage(os.Stdout)
}
