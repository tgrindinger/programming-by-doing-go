package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func FindElement(random *rand.Rand, writer io.Writer) {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, 1 + random.Intn(100))
	}
	fmt.Fprintf(writer, "ArrayList: %v\n", arr)
	fmt.Fprintln(writer)
	largestIndex := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > arr[largestIndex] {
			largestIndex = i
		}
	}
	fmt.Fprintf(writer, "The largest value is %d, which is in slot %d\n", arr[largestIndex], largestIndex)
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	FindElement(random, os.Stdout)
}
