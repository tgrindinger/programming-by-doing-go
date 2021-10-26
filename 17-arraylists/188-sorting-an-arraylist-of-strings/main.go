package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func sort(arr []string) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
}

func FindElement(random *rand.Rand, writer io.Writer) {
	arr := []string{"ask", "not", "what", "your", "country", "can", "do", "for", "you", "but"}
	fmt.Fprintf(writer, "ArrayList before: %v\n", arr)
	sort(arr)
	fmt.Fprintf(writer, "ArrayList after : %v\n", arr)
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	FindElement(random, os.Stdout)
}
