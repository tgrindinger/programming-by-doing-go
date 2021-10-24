package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func PrintElements(random *rand.Rand, writer io.Writer) {
	arr := make([]int, 1000)
	for i := 0; i < len(arr); i++ {
		arr[i] = 10 + random.Intn(90)
	}
	for i := 0; i < len(arr); i++ {
		if i % 20 == 0 && i != 0 {
			fmt.Fprintln(writer)
		}
		fmt.Fprintf(writer, "%d  ", arr[i])
	}
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	PrintElements(random, os.Stdout)
}
