package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func sort(arr []string) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
}

func FindElement(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprint(writer, "Sentence: ")
	scanner.Scan()
	arr := strings.Split(strings.ToLower(scanner.Text()), " ")
	sort(arr)
	fmt.Fprintf(writer, "Sorted: %v\n", arr)
}

func main() {
	FindElement(os.Stdin, os.Stdout)
}
