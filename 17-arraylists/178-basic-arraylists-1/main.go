package main

import (
	"fmt"
	"io"
	"os"
)

func PrintArrayList(writer io.Writer) {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, -113)
	}
	for i := 0; i < len(arr); i++ {
		fmt.Fprintf(writer, "Slot %d contains a %d\n", i, arr[i])
	}
}

func main() {
	PrintArrayList(os.Stdout)
}
