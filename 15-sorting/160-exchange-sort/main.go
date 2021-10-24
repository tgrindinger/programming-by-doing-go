package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func swap(a []int, i, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}

func exchangeSort(a []int) {
	for i := 0; i < len(a) - 1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				swap(a, i, j)
			}
		}
	}
}

func SortArray(random *rand.Rand, writer io.Writer) {
	arr := make([]int, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = 1 + random.Intn(100)
	}
	fmt.Fprint(writer, "before: ")
	for i := 0; i < len(arr); i++ {
		fmt.Fprintf(writer, "%d ", arr[i])
	}
	fmt.Fprintln(writer)
	exchangeSort(arr)
	fmt.Fprint(writer, "after:  ")
	for i := 0; i < len(arr); i++ {
		fmt.Fprintf(writer, "%d ", arr[i])
	}
	fmt.Fprintln(writer)
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	SortArray(random, os.Stdout)
}
