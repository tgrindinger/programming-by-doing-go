package main

import (
	"fmt"
	"math/big"
)

func sumPowerOf2Digits(power int64) int64 {
	result := big.NewInt(2)
	bPower := big.NewInt(power)
	result.Exp(result, bPower, nil)
	digits := result.Text(10)
	sum := int64(0)
	for i := 0; i < len(digits); i++ {
		sum += int64(digits[i]) - '0'
	}
	return sum
}

func main() {
	fmt.Printf("Power of 2 to sum digits: ")
	var max int64
	fmt.Scanln(&max)
	result := sumPowerOf2Digits(max)
	fmt.Printf("The result is %d\n", result)
}
