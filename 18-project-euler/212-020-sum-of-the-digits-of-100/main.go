package main

import (
	"fmt"
	"math/big"
)

func factorial(num int) *big.Int {
	prod := big.NewInt(1)
	for i := int64(1); i <= int64(num); i++ {
		prod.Mul(prod, big.NewInt(i))
	}
	return prod
}

func sumOfDigitsOfFactorial(num int) int {
	nFac := factorial(num)
	digits := nFac.String()
	sum := 0
	for i := 0; i < len(digits); i++ {
		sum += int(digits[i]) - '0'
	}
	return sum
}

func main() {
	fmt.Printf("Factorial to sum digits: ")
	var max int
	fmt.Scanln(&max)
	result := sumOfDigitsOfFactorial(max)
	fmt.Printf("The result is %d\n", result)
}
