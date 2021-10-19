package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func PrintRandomNumbers(random *rand.Rand, writer io.Writer) {
	x := 1 + random.Intn(10)
	fmt.Fprintf(writer, "My random number is %d\n", x)
	fmt.Fprintln(writer, "Here are some numbers from 1 to 5!")
	fmt.Fprintf(writer, "%d ", 1 + random.Intn(5))
	fmt.Fprintf(writer, "%d ", 1 + random.Intn(5))
	fmt.Fprintf(writer, "%d ", 1 + random.Intn(5))
	fmt.Fprintf(writer, "%d ", 1 + random.Intn(5))
	fmt.Fprintf(writer, "%d ", 1 + random.Intn(5))
	fmt.Fprintf(writer, "%d ", 1 + random.Intn(5))
	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "Here are some numbers from 1 to 100!")
	fmt.Fprintf(writer, "%d\t", 1 + random.Intn(100))
	fmt.Fprintf(writer, "%d\t", 1 + random.Intn(100))
	fmt.Fprintf(writer, "%d\t", 1 + random.Intn(100))
	fmt.Fprintf(writer, "%d\t", 1 + random.Intn(100))
	fmt.Fprintf(writer, "%d\t", 1 + random.Intn(100))
	fmt.Fprintf(writer, "%d\t", 1 + random.Intn(100))
	fmt.Fprintln(writer)
	num1 := 1 + random.Intn(10)
	num2 := 1 + random.Intn(10)
	if num1 == num2 {
		fmt.Fprintln(writer, "The random numbers were the same! Weird.")
	} else {
		fmt.Fprintln(writer, "The random numbers were different! Not too surprising, actually.")
	}
}

func main() {
	PrintRandomNumbers(rand.New(rand.NewSource(time.Now().UnixMicro())), os.Stdout)
}
