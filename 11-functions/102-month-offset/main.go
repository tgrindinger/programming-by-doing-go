package main

import (
	"fmt"
	"io"
	"os"
)

func monthOffset(month int) int {
	var offset int
	if month == 1 {
		offset = 1
	} else if month == 2 {
		offset = 4
	} else if month == 3 {
		offset = 4
	} else if month == 4 {
		offset = 0
	} else if month == 5 {
		offset = 2
	} else if month == 6 {
		offset = 5
	} else if month == 7 {
		offset = 0
	} else if month == 8 {
		offset = 3
	} else if month == 9 {
		offset = 6
	} else if month == 10 {
		offset = 1
	} else if month == 11 {
		offset = 4
	} else if month == 12 {
		offset = 6
	} else {
		offset = -1
	}
	return offset
}

func PrintMonthOffsets(writer io.Writer) {
	fmt.Fprintf(writer, "Offset for month 1: %d\n", monthOffset(1))
	fmt.Fprintf(writer, "Offset for month 2: %d\n", monthOffset(2))
	fmt.Fprintf(writer, "Offset for month 3: %d\n", monthOffset(3))
	fmt.Fprintf(writer, "Offset for month 4: %d\n", monthOffset(4))
	fmt.Fprintf(writer, "Offset for month 5: %d\n", monthOffset(5))
	fmt.Fprintf(writer, "Offset for month 6: %d\n", monthOffset(6))
	fmt.Fprintf(writer, "Offset for month 7: %d\n", monthOffset(7))
	fmt.Fprintf(writer, "Offset for month 8: %d\n", monthOffset(8))
	fmt.Fprintf(writer, "Offset for month 9: %d\n", monthOffset(9))
	fmt.Fprintf(writer, "Offset for month 10: %d\n", monthOffset(10))
	fmt.Fprintf(writer, "Offset for month 11: %d\n", monthOffset(11))
	fmt.Fprintf(writer, "Offset for month 12: %d\n", monthOffset(12))
	fmt.Fprintf(writer, "Offset for month 43: %d\n", monthOffset(43))
}

func main() {
	PrintMonthOffsets(os.Stdout)
}
