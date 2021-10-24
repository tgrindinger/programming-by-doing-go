package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func FindElement(random *rand.Rand, writer io.Writer) {
	arr := make([]int, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = 1 + random.Intn(100)
	}
	fmt.Fprint(writer, "Array: ")
	for i := 0; i < len(arr); i++ {
		fmt.Fprintf(writer, "%d ", arr[i])
	}
	fmt.Fprintln(writer)
	fmt.Fprintln(writer)
	largest := 0
	largestIndex := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > largest {
			largest = arr[i]
			largestIndex = i
		}
	}
	fmt.Fprintf(writer, "The largest value is %d\n", largest)
	fmt.Fprintf(writer, "It is in slot %d\n", largestIndex)
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	FindElement(random, os.Stdout)
}
