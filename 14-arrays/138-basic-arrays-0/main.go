package main

import (
	"fmt"
	"io"
	"os"
)

func PrintElements(writer io.Writer) {
	arr := []int { -113, -113, -113, -113, -113, -113, -113, -113, -113, -113 }
	fmt.Fprintf(writer, "Slot %d contains a %d\n", 0, arr[0])
	fmt.Fprintf(writer, "Slot %d contains a %d\n", 1, arr[1])
	fmt.Fprintf(writer, "Slot %d contains a %d\n", 2, arr[2])
	fmt.Fprintf(writer, "Slot %d contains a %d\n", 3, arr[3])
	fmt.Fprintf(writer, "Slot %d contains a %d\n", 4, arr[4])
	fmt.Fprintf(writer, "Slot %d contains a %d\n", 5, arr[5])
	fmt.Fprintf(writer, "Slot %d contains a %d\n", 6, arr[6])
	fmt.Fprintf(writer, "Slot %d contains a %d\n", 7, arr[7])
	fmt.Fprintf(writer, "Slot %d contains a %d\n", 8, arr[8])
	fmt.Fprintf(writer, "Slot %d contains a %d\n", 9, arr[9])
}

func main() {
	PrintElements(os.Stdout)
}
