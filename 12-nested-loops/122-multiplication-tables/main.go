package main

import (
	"fmt"
	"io"
	"os"
)

func PrintTable(writer io.Writer) {
	fmt.Fprintln(writer, "x | 1\t2\t3\t4\t5\t6\t7\t8\t9")
	fmt.Fprintln(writer, "==+==================================================================")
	for i := 1; i <= 12; i++ {
		fmt.Fprintf(writer, "%d | ", i)
		for j := 1; j <= 9; j++ {
			fmt.Fprintf(writer, "%d\t", i * j)
		}
		fmt.Fprintln(writer)
	}
}

func main() {
	PrintTable(os.Stdout)
}
