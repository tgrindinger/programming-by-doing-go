package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func FindSumOfMultiplesOf3Or5(num int) int {
	sum := 0
	for i := 0; i < num; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Number to sum multiple of 3 or 5 up to: ")
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	result := FindSumOfMultiplesOf3Or5(num)
	fmt.Printf("The sum is %d\n", result)
}
