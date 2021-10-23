package main

import (
	"fmt"
	"io"
	"os"
)

func PrintLetter(writer io.Writer) {
	fmt.Fprintln(writer, "+-------------------------------------------------------+")
	fmt.Fprintln(writer, "|                                                  #### |")
	fmt.Fprintln(writer, "|                                                  #### |")
	fmt.Fprintln(writer, "|                                                  #### |")
	fmt.Fprintln(writer, "|                                                       |")
	fmt.Fprintln(writer, "|                                                       |")
	fmt.Fprintln(writer, "|                        Bill Gates                     |")
	fmt.Fprintln(writer, "|                        1 Microsoft Way                |")
	fmt.Fprintln(writer, "|                        Redmond, WA 98104              |")
	fmt.Fprintln(writer, "|                                                       |")
	fmt.Fprintln(writer, "+-------------------------------------------------------+")
}

func main() {
	writer, _ := os.Create("letter.txt")
	PrintLetter(writer)
	writer.Close()
}
