package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isPrime(primes []int, num int) bool {
	for i := 0; i < len(primes); i++ {
		if num % primes[i] == 0 {
			return false
		}
	}
	return true
}

func FindNthPrime(n int) int {
	primes := []int{}
	for i := 2; len(primes) < n; i++ {
		if isPrime(primes, i) {
			primes = append(primes, i)
		}
	}
	return primes[len(primes)-1]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Which prime to find: ")
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	result := FindNthPrime(num)
	fmt.Printf("The result is %d\n", result)
}
