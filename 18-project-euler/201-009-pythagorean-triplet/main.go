package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func PythagoreanTriplet(num int) int {
	for a := 1; a < num; a++ {
		for b := a + 1; a + b < num; b++ {
			c := num - (a + b)
			if a * a + b * b == c * c {
				return a * b * c
			}
		}
	}
	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Sum of Pythagorean triplet: ")
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	result := PythagoreanTriplet(num)
	fmt.Printf("The result is %d\n", result)
}
