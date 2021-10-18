package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func AskQuestions(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprint(writer, "Your height in inches: ")
	scanner.Scan()
	height, _ := strconv.ParseFloat(scanner.Text(), 64)
	fmt.Fprint(writer, "Your weight in pounds: ")
	scanner.Scan()
	weight, _ := strconv.ParseFloat(scanner.Text(), 64)
	fmt.Fprintln(writer)
	weightKg := weight * 0.4535924
	heightM := height * 0.0254
	bmi := weightKg / (heightM * heightM)
	var category string
	if bmi < 15 {
		category = "very severely underweight"
	} else if bmi < 16 {
		category = "severely underweight"
	} else if bmi < 18.5 {
		category = "underweight"
	} else if bmi < 25 {
		category = "normal weight"
	} else if bmi < 30 {
		category = "overweight"
	} else if bmi < 35 {
		category = "moderately obese"
	} else if bmi < 40 {
		category = "severely obese"
	} else {
		category = "morbidly obese"
	}
	fmt.Fprintf(writer, "Your BMI is %.1f\n", bmi)
	fmt.Fprintf(writer, "BMI Category: %s\n", category)
}

func main() {
	AskQuestions(os.Stdin, os.Stdout)
}
