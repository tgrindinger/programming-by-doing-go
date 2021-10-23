package main

import (
	"fmt"
	"io"
	"os"
)

func PrintArmstrongNumbers(writer io.Writer) {
	for i := 1; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			for k := 0; k <= 9; k++ {
				num := i * 100 + j * 10 + k
				if i * i * i + j * j * j + k * k * k == num {
					fmt.Fprintf(writer, "%d\n", num)
				}
			}
		}
	}
}

func main() {
	PrintArmstrongNumbers(os.Stdout)
}
