package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func factors(num int) []int {
	divs := []int{}
	for i := 2; i <= num; i++ {
		for num % i == 0 {
			divs = append(divs, i)
			num /= i
		}
	}
	return divs
}

func numDivisors(num int) int {
	facs := factors(num)
	max := (2 << (len(facs) + 1)) - 1
	factorMap := map[int]bool{}
	for binMap := 0; binMap < max; binMap++ {
		product := 1
		for j := 0; j < len(facs); j++ {
			if 1 << j & binMap > 0 {
				product *= facs[j]
			}
		}
		factorMap[product] = true
	}
	return len(factorMap)
}

func FirstTriangleNumberWithOverNDivisors(n int) int {
	current := 0
	for i := 1; numDivisors(current) <= n; i++ {
		current += i
	}
	return current
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Number of divisors: ")
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	result := FirstTriangleNumberWithOverNDivisors(num)
	fmt.Printf("The result is %d\n", result)
}
