package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func PrintArrayList(random *rand.Rand, writer io.Writer) {
	arr1 := []int{}
	for i := 0; i < 10; i++ {
		arr1 = append(arr1, 1 + random.Intn(100))
	}
	arr2 := []int{}
	for i := 0; i < len(arr1); i++ {
		arr2 = append(arr2, arr1[i])
	}
	arr1[len(arr1)-1] = -7
	fmt.Fprintf(writer, "ArrayList 1: %v\n", arr1)
	fmt.Fprintf(writer, "ArrayList 2: %v\n", arr2)
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	PrintArrayList(random, os.Stdout)
}
