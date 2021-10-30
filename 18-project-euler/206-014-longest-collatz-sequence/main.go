package main

import (
	"fmt"
)

func collatzLength(start int) int {
	length := 1
	for start != 1 {
		if start % 2 == 0 {
			start /= 2
		} else {
			start = 3 * start + 1
		}
		length++
	}
	return length
}

func longestCollatzLength(max int) int {
	longest := 0
	start := 0
	for i := 1; i < max; i++ {
		length := collatzLength(i)
		if length > longest {
			longest = length
			start = i
		}
	}
	return start
}

func main() {
	fmt.Printf("Max starting value: ")
	var max int
	fmt.Scanln(&max)
	result := longestCollatzLength(max)
	fmt.Printf("The result is %d\n", result)
}
