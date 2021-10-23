package main

import (
	"fmt"
	"io"
	"os"
)

func monthName(month int) string {
	var name string
	if month == 1 {
		name = "January"
	} else if month == 2 {
		name = "February"
	} else if month == 3 {
		name = "March"
	} else if month == 4 {
		name = "April"
	} else if month == 5 {
		name = "May"
	} else if month == 6 {
		name = "June"
	} else if month == 7 {
		name = "July"
	} else if month == 8 {
		name = "August"
	} else if month == 9 {
		name = "September"
	} else if month == 10 {
		name = "October"
	} else if month == 11 {
		name = "November"
	} else if month == 12 {
		name = "December"
	} else {
		name = "error"
	}
	return name
}

func PrintMonths(writer io.Writer) {
	fmt.Fprintf(writer, "Month 1: %s\n", monthName(1))
	fmt.Fprintf(writer, "Month 2: %s\n", monthName(2))
	fmt.Fprintf(writer, "Month 3: %s\n", monthName(3))
	fmt.Fprintf(writer, "Month 4: %s\n", monthName(4))
	fmt.Fprintf(writer, "Month 5: %s\n", monthName(5))
	fmt.Fprintf(writer, "Month 6: %s\n", monthName(6))
	fmt.Fprintf(writer, "Month 7: %s\n", monthName(7))
	fmt.Fprintf(writer, "Month 8: %s\n", monthName(8))
	fmt.Fprintf(writer, "Month 9: %s\n", monthName(9))
	fmt.Fprintf(writer, "Month 10: %s\n", monthName(10))
	fmt.Fprintf(writer, "Month 11: %s\n", monthName(11))
	fmt.Fprintf(writer, "Month 12: %s\n", monthName(12))
	fmt.Fprintf(writer, "Month 43: %s\n", monthName(43))
}

func main() {
	PrintMonths(os.Stdout)
}
