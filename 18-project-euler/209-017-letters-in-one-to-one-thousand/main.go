package main

import (
	"fmt"
)

func lettersInWords(num int) int {
	letters := map[int]int{
		1: 3, 2: 3, 3: 5, 4: 4, 5: 4, 6: 3, 7: 5, 8: 5, 9: 4,
		10: 3, 11: 6, 12: 6, 13: 8, 14: 8, 15: 7, 16: 7, 17: 9, 18: 8, 19: 8,
		20: 6, 30: 6, 40: 5, 50: 5, 60: 5, 70: 7, 80: 6, 90: 6,
		1000: 11,
	}
	hundred := len("hundred")
	and := len("and")
	sum := 0
	if num == 1000 {
		sum = letters[num]
		num = 0
	}
	if num >= 100 {
		sum += letters[num / 100] + hundred
		if num % 100 > 0 {
			sum += and
		}
		num %= 100
	}
	if num >= 20 {
		sum += letters[num / 10 * 10]
		num %= 10
	}
	if num > 0 {
		sum += letters[num]
	}
	return sum
}

func lettersInNumbersUpTo(max int) int {
	sum := 0
	for i := 1; i <= max; i++ {
		sum += lettersInWords(i)
	}
	return sum
}

func main() {
	fmt.Printf("Number of letters in words printed up to: ")
	var max int
	fmt.Scanln(&max)
	result := lettersInNumbersUpTo(max)
	fmt.Printf("The result is %d\n", result)
}
