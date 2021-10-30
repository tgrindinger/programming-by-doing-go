package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func factors(num int) []int {
	divs := make([]int, num + 1)
	for i := 2; i <= num; i++ {
		for num % i == 0 {
			divs[i]++
			num /= i
		}
	}
	return divs
}

func max(i1, i2 int) int {
	if i1 > i2 {
		return i1
	} else {
		return i2
	}
}

func maxArray(arr1, arr2 []int) []int {
	small := arr1
	large := arr2
	if len(arr1) > len(arr2) {
		small = arr2
		large = arr1
	}
	maxArray := make([]int, len(large))
	for i := 0; i < len(small); i++ {
		maxArray[i] = max(large[i], small[i])
	}
	for i := len(small); i < len(large); i++ {
		maxArray[i] = large[i]
	}
	return maxArray
}

func factorsInRange(max int) []int {
	allFactors := make([]int, max + 1)
	for i := 2; i <= max; i++ {
		allFactors = maxArray(allFactors, factors(i))
	}
	return allFactors
}

func combineFactors(factors []int) int64 {
	result := int64(1)
	for i := 2; i < len(factors); i++ {
		for j := 0; j < factors[i]; j++ {
			result *= int64(i)
		}
	}
	return result
}

func SmallestDivisibleByUpTo(max int) int64 {
	rangeFactors := factorsInRange(max)
	return combineFactors(rangeFactors)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Smallest number divisible by all natural numbers up to: ")
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	result := SmallestDivisibleByUpTo(num)
	fmt.Printf("The result is %d\n", result)
}
