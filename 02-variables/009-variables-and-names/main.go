package main

import (
	"fmt"
	"io"
	"os"
)

func PrintVariables(writer io.Writer) {
	var cars, drivers, passengers, carsNotDriven, carsDriven int
	var spaceInACar, carpoolCapacity, averagePassengersPerCar float64
	cars = 100
	spaceInACar = 4.0
	drivers = 30
	passengers = 90
	carsNotDriven = cars - drivers
	carsDriven = drivers
	carpoolCapacity = float64(carsDriven) * spaceInACar
	averagePassengersPerCar = float64(passengers) / float64(carsDriven)
	fmt.Fprintln(writer, "There are", cars, "cars available.")
	fmt.Fprintln(writer, "There are only", drivers, "drivers available.")
	fmt.Fprintln(writer, "There will be", carsNotDriven, "empty cars today.")
	fmt.Fprintf(writer, "We can transport %.1f people today.", carpoolCapacity)
	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "We have", passengers, "to carpool today.")
	fmt.Fprintf(writer, "We need to put about %.1f in each car.", averagePassengersPerCar)
	fmt.Fprintln(writer)
}

func main() {
	PrintVariables(os.Stdout)
}
