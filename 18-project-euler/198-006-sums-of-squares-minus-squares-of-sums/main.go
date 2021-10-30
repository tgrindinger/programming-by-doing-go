package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func SumOfSquares(max int) int {
	sum := 0
	for i := 1; i <= max; i++ {
		sum += i * i
	}
	return sum
}

func SquareOfSum(max int) int {
	sum := 0
	for i := 1; i <= max; i++ {
		sum += i
	}
	return sum * sum
}

func DifferenceBetweenSumOfSquaresAndSquareOfSum(max int) int {
	return SquareOfSum(max) - SumOfSquares(max)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Number to find difference between sum of squares an square of sum: ")
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	result := DifferenceBetweenSumOfSquaresAndSquareOfSum(num)
	fmt.Printf("The result is %d\n", result)
}
