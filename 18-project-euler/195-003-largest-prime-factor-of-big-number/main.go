package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func LargestPrimeFactor(num int64) int64 {
	largest := int64(0)
	for i := int64(2); i <= int64(math.Ceil(math.Sqrt(float64(num)))); i++ {
		if num % i == 0 {
			largest = i
			for num % i == 0 {
				num /= i
			}
		}
	}
	if num > largest {
		largest = num
	}
	return largest
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Number to find largest prime factor: ")
	scanner.Scan()
	num, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	result := LargestPrimeFactor(num)
	fmt.Printf("The factor is %d\n", result)
}
