package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func LookupId(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	names := []string{ "Mitchell", "Ortiz", "Luu", "Zimmerman", "Brooks" }
	grades := []float64{ 99.5, 78.5, 95.6, 96.8, 82.7 }
	ids := []int{ 123456, 813225, 823669, 307760, 827131 }
	fmt.Fprintln(writer, "Values:")
	for i := 0; i < len(names); i++ {
		fmt.Fprintf(writer, "\t%s %.1f %d\n", names[i], grades[i], ids[i])
	}
	fmt.Fprintln(writer)
	fmt.Fprint(writer, "ID number to find: ")
	scanner.Scan()
	id, _ := strconv.Atoi(scanner.Text())
	index := -1
	for i := 0; i < len(names); i++ {
		if ids[i] == id {
			index = i
		}
	}
	fmt.Fprintln(writer)
	fmt.Fprintf(writer, "Found in slot %d\n", index)
	fmt.Fprintf(writer, "\tName: %s\n", names[index])
	fmt.Fprintf(writer, "\tAverage: %.1f\n", grades[index])
	fmt.Fprintf(writer, "\tID: %d\n", ids[index])
}

func main() {
	LookupId(os.Stdin, os.Stdout)
}
