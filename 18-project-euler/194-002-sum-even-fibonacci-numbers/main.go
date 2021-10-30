package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func SumEvenFibonacciNumbers(num int) int {
	sum := 0
	index := 1
	last := 1
	current := 1
	for current <= num {
		if index % 2 == 0 {
			sum += current
		}
		newCurrent := last + current
		last = current
		current = newCurrent
		index++
	}
	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Number to sum even fibonacci numbers up to: ")
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	result := SumEvenFibonacciNumbers(num)
	fmt.Printf("The sum is %d\n", result)
}
