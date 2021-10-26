package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func PrintArrayList(random *rand.Rand, writer io.Writer) {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, 1 + random.Intn(100))
	}
	fmt.Fprintf(writer, "ArrayList: %v\n", arr)
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	PrintArrayList(random, os.Stdout)
}
