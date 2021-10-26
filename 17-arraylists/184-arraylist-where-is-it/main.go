package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func FindElement(random *rand.Rand, reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, 1 + random.Intn(50))
	}
	fmt.Fprintf(writer, "ArrayList: %v\n", arr)
	fmt.Fprint(writer, "Value to find: ")
	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	fmt.Fprintln(writer)
	index := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] == value {
			index = i
		}
	}
	if index == -1 {
		fmt.Fprintf(writer, "%d is not in the ArrayList.\n", value)
	} else {
		fmt.Fprintf(writer, "%d is in slot %d.\n", value, index)
	}
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	FindElement(random, os.Stdin, os.Stdout)
}
