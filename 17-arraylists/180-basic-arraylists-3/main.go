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
	for i := 0; i < 1000; i++ {
		arr = append(arr, 10 + random.Intn(90))
	}
	fmt.Fprintf(writer, "%v\n", arr)
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	PrintArrayList(random, os.Stdout)
}
