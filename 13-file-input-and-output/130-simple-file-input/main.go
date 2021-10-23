package main

import (
	"fmt"
	"io"
	"os"
)

func PrintName(reader io.Reader, writer io.Writer) {
	var firstName, lastName string
	fmt.Fscanln(reader, &firstName, &lastName)
	fmt.Fprintf(writer, "Using my psychic powers (aided by reading data from the file), I have determined that your name is %s %s.\n", firstName, lastName)
}

func main() {
	reader, _ := os.Open("name.txt")
	PrintName(reader, os.Stdout)
	reader.Close()
}
