package main

import (
	"fmt"
	"io"
	"os"
)

func PrintElements(writer io.Writer) {
	arr := make([]int, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = -113
	}
	for i := 0; i < len(arr); i++ {
		fmt.Fprintf(writer, "Slot %d contains a %d\n", i, arr[i])
	}
}

func main() {
	PrintElements(os.Stdout)
}
