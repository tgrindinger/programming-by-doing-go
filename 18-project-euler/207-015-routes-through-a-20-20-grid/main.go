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

func numberPaths(size int) *big.Int {
	nFac := factorial(size * 2)
	kFac := factorial(size)
	return nFac.Div(nFac, kFac.Mul(kFac, kFac))
}

func main() {
	fmt.Printf("Grid size: ")
	var max int
	fmt.Scanln(&max)
	result := numberPaths(max)
	fmt.Printf("The result is %s\n", result)
}
