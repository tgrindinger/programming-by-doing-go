package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func FillArray(random *rand.Rand, writer io.Writer) {
	arr1 := []string{ "alpha", "bravo", "charlie", "delta", "echo" }
	fmt.Fprintln(writer, "The first array is filled with the following values:")
	fmt.Fprint(writer, "\t")
	for i := 0; i < len(arr1); i++ {
		fmt.Fprintf(writer, "%s ", arr1[i])
	}
	fmt.Fprintln(writer)

	arr2 := []int{ 11, 23, 37, 41, 53 }
	fmt.Fprintln(writer, "The second array is filled with the following values:")
	fmt.Fprint(writer, "\t")
	for i := 0; i < len(arr2); i++ {
		fmt.Fprintf(writer, "%d ", arr2[i])
	}
	fmt.Fprintln(writer)
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	FillArray(random, os.Stdout)
}
