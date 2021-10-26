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
	there := false
	for i := 0; i < len(arr); i++ {
		if arr[i] == value {
			there = true
		}
	}
	if there {
		fmt.Fprintf(writer, "%d is in the ArrayList.\n", value)
	} else {
		fmt.Fprintf(writer, "%d is not in the ArrayList.\n", value)
	}
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	FindElement(random, os.Stdin, os.Stdout)
}
