package main

import (
	"fmt"
	"io"
	"os"
)

func PrintPuzzleNumbers(writer io.Writer) {
	for i := 0; i <= 45; i++ {
		for j := 0; j <= 45; j++ {
			for k := 0; k <= 45; k++ {
				for l := 0; l <= 45; l++ {
					sum := i + j + k + l
					firstPlus2 := i + 2
					secondMinus2 := j - 2
					thirdTimes2 := k * 2
					fourthDiv2 := l / 2
					if sum == 45 && firstPlus2 == secondMinus2 && secondMinus2 == thirdTimes2 && thirdTimes2 == fourthDiv2 {
						fmt.Fprintf(writer, "%d %d %d %d\n", i, j, k, l)
					}
				}
			}
		}
	}
}

func main() {
	PrintPuzzleNumbers(os.Stdout)
}
