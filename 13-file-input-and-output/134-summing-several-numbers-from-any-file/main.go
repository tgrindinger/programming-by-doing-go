package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func PrintSum(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprint(writer, "Which file would you like to read numbers from: ")
	scanner.Scan()
	filename := scanner.Text()
	fmt.Fprintf(writer, "Reading numbers from %q\n", filename)
	fileReader, _ := os.Open(filename)
	fmt.Fprintln(writer)
	total := 0
	var num int
	numsRead, _ := fmt.Fscan(fileReader, &num)
	for numsRead == 1 {
		total += num
		fmt.Fprintf(writer, "%d ", num)
		numsRead, _ = fmt.Fscan(fileReader, &num)
	}
	fileReader.Close()
	fmt.Fprintln(writer)
	fmt.Fprintf(writer, "Total is %d\n", total)
}

func main() {
	PrintSum(os.Stdin, os.Stdout)
}
