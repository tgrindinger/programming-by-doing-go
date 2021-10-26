package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Car struct {
	make    string
	model   string
	year    int
	license string
}

func ReadCar(scanner *bufio.Scanner) *Car {
	car := &Car{}
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &car.make, &car.model, &car.year, &car.license)
	return car
}

func (c *Car) PrintCar(writer io.Writer, index int) {
	fmt.Fprintf(writer, "Car %d\n", index)
	fmt.Fprintf(writer, "\tMake: %s\n", c.make)
	fmt.Fprintf(writer, "\tModel: %s\n", c.model)
	fmt.Fprintf(writer, "\tYear: %d\n", c.year)
	fmt.Fprintf(writer, "\tLicense: %s\n", c.license)
	fmt.Fprintln(writer)
}

func (c *Car) String() string {
	return fmt.Sprintf("%d %s %s (%s)", c.year, c.make, c.model, c.license)
}

func sort(cars []*Car) {
	for i := 0; i < len(cars); i++ {
		for j := i + 1; j < len(cars); j++ {
			if cars[i].year > cars[j].year {
				temp := cars[i]
				cars[i] = cars[j]
				cars[j] = temp
			}
		}
	}
}

func PrintCars(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprint(writer, "From which file do you want to load this information? ")
	scanner.Scan()
	filename := scanner.Text()
	file, _ := os.Open(filename)
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	var cars []*Car
	for i := 0; i < 5; i++ {
		cars = append(cars, ReadCar(fileScanner))
	}
	fmt.Fprintln(writer, "Data loaded.")
	sort(cars)
	fmt.Fprintln(writer, "Data sorted.")
	fmt.Fprintln(writer)
	fmt.Fprintf(writer, "ArrayList: %v\n", cars)
}

func main() {
	PrintCars(os.Stdin, os.Stdout)
}
