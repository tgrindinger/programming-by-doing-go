package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func PrintElements(random *rand.Rand, writer io.Writer) {
	arr1 := make([]int, 10)
	for i := 0; i < len(arr1); i++ {
		arr1[i] = 1 + random.Intn(10)
	}
	arr2 := make([]int, 10)
	for i := 0; i < len(arr2); i++ {
		arr2[i] = arr1[i]
	}
	arr1[len(arr1)-1] = -7
	fmt.Fprint(writer, "Array 1: ")
	for i := 0; i < len(arr1); i++ {
		fmt.Fprintf(writer, "%d ", arr1[i])
	}
	fmt.Fprintln(writer)
	fmt.Fprint(writer, "Array 2: ")
	for i := 0; i < len(arr2); i++ {
		fmt.Fprintf(writer, "%d ", arr2[i])
	}
	fmt.Fprintln(writer)
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	PrintElements(random, os.Stdout)
}
