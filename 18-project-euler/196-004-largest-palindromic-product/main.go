package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func findStart(digits int) int {
	result := 1
	for i := 0; i < digits - 1; i++ {
		result *= 10
	}
	return result
}

func findEnd(digits int) int {
	result := 0
	for i := 0; i < digits; i++ {
		result = result * 10 + 9
	}
	return result
}

func isPalindrome(num int) bool {
	s := fmt.Sprint(num)
	for i := 0; i < len(s) / 2; i++ {
		if s[i] != s[len(s) - i - 1] {
			return false
		}
	}
	return true
}

func LargestPalindromicProduct(digits int) int {
	largest := 0
	start := findStart(digits)
	end := findEnd(digits)
	for i := end; i >= start; i-- {
		for j := end; j >= start; j-- {
			num := i * j
			if num > largest && isPalindrome(num) {
				largest = num
			}
		}
	}
	return largest
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Largest palindrome of products with how many digits: ")
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	result := LargestPalindromicProduct(num)
	fmt.Printf("The factor is %d\n", result)
}
