package main

import (
	"fmt"
	"io"
	"os"
)

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func PrintPrimes(writer io.Writer) {
	for i := 2; i <= 20; i++ {
		fmt.Fprintf(writer, "%d ", i)
		if isPrime(i) {
			fmt.Fprint(writer, "<")
		}
		fmt.Fprintln(writer)
	}
}

func main() {
	PrintPrimes(os.Stdout)
}
