package main

import (
	"fmt"
	"io"
	"os"
)

func PrintNameAndGraduationYear(writer io.Writer) {
	name := "Juan Valdez"
	year := 2010
	fmt.Fprintf(writer, "My name is %s and I'll graduate in %d.\n", name, year)
}

func main() {
	PrintNameAndGraduationYear(os.Stdout)
}