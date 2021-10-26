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
	largest := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] > largest {
			largest = arr[i]
		}
	}
	fmt.Fprintf(writer, "The largest value is %d\n", largest)
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	FindElement(random, os.Stdout)
}
