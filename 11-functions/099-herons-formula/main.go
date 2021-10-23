package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

func triangleArea(a, b, c int) float64 {
	s := float64(a + b + c) / 2
	A := math.Sqrt(float64(s * (s - float64(a)) * (s - float64(b)) * (s - float64(c))))
	return A
}

func PrintAreas(writer io.Writer) {
	a := triangleArea(3, 3, 3)
	fmt.Fprintf(writer, "A triangle with sides 3,3,3 has an area of %f\n", a)
	a = triangleArea(3, 4, 5)
	fmt.Fprintf(writer, "A triangle with sides 3,4,5 has an area of %f\n", a)
	a = triangleArea(7, 8, 9)
	fmt.Fprintf(writer, "A triangle with sides 7,8,9 has an area of %f\n", a)
	fmt.Fprintf(writer, "A triangle with sides 5,12,13 has an area of %f\n", triangleArea(5, 12, 13))
	fmt.Fprintf(writer, "A triangle with sides 10,9,11 has an area of %f\n", triangleArea(10, 9, 11))
	fmt.Fprintf(writer, "A triangle with sides 8,15,17 has an area of %f\n", triangleArea(8, 15, 17))
	fmt.Fprintf(writer, "A triangle with sides 9,9,9 has an area of %f\n", triangleArea(9, 9, 9))
}

func main() {
	PrintAreas(os.Stdout)
}
