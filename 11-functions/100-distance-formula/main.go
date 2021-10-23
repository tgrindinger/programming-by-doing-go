package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

func distance(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(math.Pow(float64(x1 - x2), 2) + math.Pow(float64(y1 - y2), 2))
}

func PrintDistances(writer io.Writer) {
	d1 := distance(-2,1 , 1,5)
	fmt.Fprintf(writer, " (-2,1) to (1,5) => %f\n", d1)
	d2 := distance(-2,-3 , -4,4)
	fmt.Fprintf(writer, " (-2,-3) to (-4,4) => %f\n", d2)
	fmt.Fprintf(writer, " (2,-3) to (-1,-2) => %f\n", distance(2,-3 , -1,-2))
	fmt.Fprintf(writer, " (4,5) to (4,5) => %f\n", distance(4,5 , 4,5))
}

func main() {
	PrintDistances(os.Stdout)
}
