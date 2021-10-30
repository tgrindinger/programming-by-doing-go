package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func createTriangle(nums []int) [][]int {
	tri := [][]int{}
	length := 1
	current := 0
	for current < len(nums) {
		row := []int{}
		for j := 0; j < length; j++ {
			row = append(row, nums[current])
			current++
		}
		tri = append(tri, row)
		length++
	}
	return tri
}

func sumOfMaxSubPath(i, j int, tri [][]int) int {
	if i == len(tri) - 1 {
		return tri[i][j]
	}
	left := sumOfMaxSubPath(i + 1, j, tri)
	right := sumOfMaxSubPath(i + 1, j + 1, tri)
	if left > right {
		return tri[i][j] + left
	} else {
		return tri[i][j] + right
	}
}

func sumOfMaxPath(nums []int) int {
	tri := createTriangle(nums)
	left := sumOfMaxSubPath(1, 0, tri)
	right := sumOfMaxSubPath(1, 1, tri)
	if left > right {
		return tri[0][0] + left
	} else {
		return tri[0][0] + right
	}
}

func convertToInts(nums []string) []int {
	ints := []int{}
	for _, n := range nums {
		i, _ := strconv.Atoi(n)
		ints = append(ints, i)
	}
	return ints
}

func main() {
	fmt.Printf("Triangle to find max path: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	values := strings.Split(scanner.Text(), " ")
	result := sumOfMaxPath(convertToInts(values))
	fmt.Printf("The result is %d\n", result)
}
