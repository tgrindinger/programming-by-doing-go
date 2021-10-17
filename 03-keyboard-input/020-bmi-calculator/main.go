package main

import (
	"fmt"
	"io"
	"os"
)

func AskForMeasurements(reader io.Reader, writer io.Writer) {
	var heightString, weightString string
	var height, weight float64
	fmt.Fprint(writer, "Your height in m: ")
	fmt.Fscanln(reader, &heightString)
	fmt.Sscanf(heightString, "%f", &height)
	fmt.Fprint(writer, "Your weight in kg: ")
	fmt.Fscanln(reader, &weightString)
	fmt.Sscanf(weightString, "%f", &weight)
	bmi := weight / (height * height)
	fmt.Fprintln(writer)
	fmt.Fprintf(writer, "Your BMI is %.5f\n", bmi)
}

func main() {
	AskForMeasurements(os.Stdin, os.Stdout)
}
