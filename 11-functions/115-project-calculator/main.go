package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func fact(a int) int {
	result := 1
	for i := 2; i <= a; i++ {
		result *= i
	}
	return result
}

func RunCalculator(reader io.Reader, writer io.Writer) {
	var op1, op2, op3 string
	for ok := true; ok; ok = op1 != "0" {
		fmt.Fprint(writer, "> ")
		args, _ := fmt.Fscanln(reader, &op1, &op2, &op3)
		var c float64
		if op1 != "0" {
			if args == 2 {
				if op1 == "sin" {
					a, _ := strconv.ParseFloat(op2, 64)
					c = math.Sin(a)
				} else if op1 == "cos" {
					a, _ := strconv.ParseFloat(op2, 64)
					c = math.Cos(a)
				} else if op1 == "tan" {
					a, _ := strconv.ParseFloat(op2, 64)
					c = math.Tan(a)
				} else if op1 == "-" {
					a, _ := strconv.ParseFloat(op2, 64)
					c = -a
				} else if op1 == "sqrt" {
					a, _ := strconv.ParseFloat(op2, 64)
					c = math.Sqrt(a)
				} else if op2 == "!" {
					a, _ := strconv.Atoi(op1)
					c = float64(fact(a))
				}
			} else {
				a, _ := strconv.ParseFloat(op1, 64)
				b, _ := strconv.ParseFloat(op3, 64)
				if op2 == "+" {
					c = a + b
				} else if op2 == "-" {
					c = a - b
				} else if op2 == "*" {
					c = a * b
				} else if op2 == "/" {
					c = a / b
				} else if op2 == "^" {
					c = math.Pow(a, b)
				} else if op2 == "%" {
					c = float64(int(a) % int(b))
				}
			}
			fmt.Fprintf(writer, "%f\n", c)
		}
	}
	fmt.Fprintln(writer, "Bye, now.")
}

func main() {
	RunCalculator(os.Stdin, os.Stdout)
}
