package main

import (
	"fmt"
	"io"
	"os"
)

func PrintVariables(writer io.Writer) {
	var room int
	var e float64
	var topic string
	room = 113
	e = 2.71828
	topic = "Computer Science"
	fmt.Fprintf(writer, "This is room # %d\n", room)
	fmt.Fprintf(writer, "e is close to %.5f\n", e)
	fmt.Fprintf(writer, "I am learning a bit about %s\n", topic)
}

func main() {
	PrintVariables(os.Stdout)
}