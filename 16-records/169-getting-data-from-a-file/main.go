package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Dog struct {
	breed  string
	age    int
	weight float64
}

func ReadDog(scanner *bufio.Scanner) *Dog {
	dog := &Dog{}
	scanner.Scan()
	dog.breed = scanner.Text()
	scanner.Scan()
	dog.age, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	dog.weight, _ = strconv.ParseFloat(scanner.Text(), 64)
	return dog
}

func PrintDogs(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprint(writer, "Which file to open: ")
	scanner.Scan()
	filename := scanner.Text()
	fmt.Fprintf(writer, "Reading data from %s\n", filename)
	fmt.Fprintln(writer)
	file, _ := os.Open(filename)
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	dog1 := ReadDog(fileScanner)
	dog2 := ReadDog(fileScanner)
	fmt.Fprintf(writer, "First dog: %s, %d, %.0f\n", dog1.breed, dog1.age, dog1.weight)
	fmt.Fprintf(writer, "Second dog: %s, %d, %.0f\n", dog2.breed, dog2.age, dog2.weight)
}

func main() {
	PrintDogs(os.Stdin, os.Stdout)
}
