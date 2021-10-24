package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Car struct {
	make    string
	model   string
	year    int
	license string
}

func ReadCar(scanner *bufio.Scanner, writer io.Writer, index int) *Car {
	fmt.Fprintf(writer, "Car %d\n", index)
	fmt.Fprint(writer, "\tMake: ")
	car := &Car{}
	scanner.Scan()
	car.make = scanner.Text()
	fmt.Fprint(writer, "\tModel: ")
	scanner.Scan()
	car.model = scanner.Text()
	fmt.Fprint(writer, "\tYear: ")
	scanner.Scan()
	car.year, _ = strconv.Atoi(scanner.Text())
	fmt.Fprint(writer, "\tLicense: ")
	scanner.Scan()
	car.license = scanner.Text()
	fmt.Fprintln(writer)
	return car
}

func SaveCars(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	cars := make([]*Car, 5)
	for i := 0; i < len(cars); i++ {
		cars[i] = ReadCar(scanner, writer, i + 1)
	}
	fmt.Fprint(writer, "To which file do you want to save this information? ")
	scanner.Scan()
	filename := scanner.Text()
	fmt.Fprintln(writer)
	file, _ := os.Create(filename)
	defer file.Close()
	for i := 0; i < len(cars); i++ {
		fmt.Fprintf(file, "%s %s %d %s\n", cars[i].make, cars[i].model, cars[i].year, cars[i].license)
	}
	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "Data saved.")
}

func main() {
	SaveCars(os.Stdin, os.Stdout)
}
