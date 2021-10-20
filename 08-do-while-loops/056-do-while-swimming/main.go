package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func Swim(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	swimmer1 := "GALLANT"
	swimmer2 := "GOOFUS"
	minimumTemperature := 79.0
	fmt.Fprint(writer, "What is the current water temperature? ")
	scanner.Scan()
	currentTemperature, _ := strconv.ParseFloat(scanner.Text(), 64)
	savedTemperature := currentTemperature
	fmt.Fprintln(writer)
	fmt.Fprintf(writer, "Okay, so the current water temperature is %.1fF.\n", currentTemperature)
	fmt.Fprintf(writer, "%s approaches the lake....\n", swimmer1)
	swimTime := 0
	for currentTemperature >= minimumTemperature {
		fmt.Fprintf(writer, "\t%s swims for a bit.", swimmer1)
		swimTime++
		fmt.Fprintf(writer, " Swim time: %d min.\n", swimTime)
		time.Sleep(600)
		currentTemperature -= 0.5
		fmt.Fprintf(writer, "\tThe current water temperature is now %.1fF.\n", currentTemperature)
	}
	fmt.Fprintf(writer, "%s stops swimming. Total swim time: %d min.\n", swimmer1, swimTime)
	currentTemperature = savedTemperature
	fmt.Fprintln(writer)

	fmt.Fprintf(writer, "Okay, so the current water temperature is %.1fF.\n", currentTemperature)
	fmt.Fprintf(writer, "%s approaches the lake....\n", swimmer2)
	swimTime = 0
	for ok := true; ok; ok = currentTemperature >= minimumTemperature {
		fmt.Fprintf(writer, "\t%s swims for a bit.", swimmer2)
		swimTime++
		fmt.Fprintf(writer, " Swim time: %d min.\n", swimTime)
		time.Sleep(600)
		currentTemperature -= 0.5
		fmt.Fprintf(writer, "\tThe current water temperature is now %.1fF.\n", currentTemperature)
	}
	fmt.Fprintf(writer, "%s stops swimming. Total swim time: %d min.\n", swimmer2, swimTime)
}

func main() {
	Swim(os.Stdin, os.Stdout)
}
