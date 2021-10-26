package main

import (
	"fmt"
	"io"
	"os"
)

func PrintArrayList(writer io.Writer) {
	arr := []int{}
	arr = append(arr, -113)
	arr = append(arr, -113)
	arr = append(arr, -113)
	arr = append(arr, -113)
	arr = append(arr, -113)
	arr = append(arr, -113)
	arr = append(arr, -113)
	arr = append(arr, -113)
	arr = append(arr, -113)
	arr = append(arr, -113)
	fmt.Fprintf(writer, "Slot 0 contains a %d\n", arr[0])
	fmt.Fprintf(writer, "Slot 1 contains a %d\n", arr[1])
	fmt.Fprintf(writer, "Slot 2 contains a %d\n", arr[2])
	fmt.Fprintf(writer, "Slot 3 contains a %d\n", arr[3])
	fmt.Fprintf(writer, "Slot 4 contains a %d\n", arr[4])
	fmt.Fprintf(writer, "Slot 5 contains a %d\n", arr[5])
	fmt.Fprintf(writer, "Slot 6 contains a %d\n", arr[6])
	fmt.Fprintf(writer, "Slot 7 contains a %d\n", arr[7])
	fmt.Fprintf(writer, "Slot 8 contains a %d\n", arr[8])
	fmt.Fprintf(writer, "Slot 9 contains a %d\n", arr[9])
}

func main() {
	PrintArrayList(os.Stdout)
}
