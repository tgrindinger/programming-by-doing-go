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
	arr := make([]int, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = 1 + random.Intn(50)
	}
	fmt.Fprint(writer, "Array: ")
	for i := 0; i < len(arr); i++ {
		fmt.Fprintf(writer, "%d ", arr[i])
	}
	fmt.Fprintln(writer)
	fmt.Fprint(writer, "Value to find: ")
	scanner.Scan()
	fmt.Fprintln(writer)
	found := false
	num, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			found = true
		}
	}
	if found {
		fmt.Fprintf(writer, "%d is in the array.\n", num)
	} else {
		fmt.Fprintf(writer, "%d is not in the array.\n", num)
	}
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	FindElement(random, os.Stdin, os.Stdout)
}
