package main

import (
	"fmt"
	"io"
	"os"
)

func Sort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
}

func SortArray(writer io.Writer) {
	arr := []int{ 45, 87, 39, 32, 93, 86, 12, 44, 75, 50 }
	fmt.Fprint(writer, "before: ")
	for i := 0; i < len(arr); i++ {
		fmt.Fprintf(writer, "%d ", arr[i])
	}
	fmt.Fprintln(writer)
	Sort(arr)
	fmt.Fprint(writer, "after:  ")
	for i := 0; i < len(arr); i++ {
		fmt.Fprintf(writer, "%d ", arr[i])
	}
	fmt.Fprintln(writer)
}

func main() {
	SortArray(os.Stdout)
}
