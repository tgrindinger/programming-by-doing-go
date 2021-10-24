package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func PrintElements(random *rand.Rand, writer io.Writer) {
	arr := make([]int, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = 1 + random.Intn(100)
	}
	for i := 0; i < len(arr); i++ {
		fmt.Fprintf(writer, "Slot %d contains a %d\n", i, arr[i])
	}
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	PrintElements(random, os.Stdout)
}
