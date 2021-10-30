package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func isPrime(primes []int, num int) bool {
	max := int(math.Ceil(math.Sqrt(float64(num))))
	for i := 0; i < len(primes) && primes[i] <= max; i++ {
		if num % primes[i] == 0 {
			return false
		}
	}
	return true
}

func sum(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

func SumOfPrimes(n int) int {
	primes := []int{}
	for i := 2; i < n; i++ {
		if isPrime(primes, i) {
			primes = append(primes, i)
		}
	}
	return sum(primes)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Sum of primes below: ")
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	result := SumOfPrimes(num)
	fmt.Printf("The result is %d\n", result)
}
