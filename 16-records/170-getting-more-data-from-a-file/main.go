package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Person struct {
	name string
	age  int
}

func ReadPerson(scanner *bufio.Scanner) *Person {
	person := &Person{}
	scanner.Scan()
	person.name = scanner.Text()
	scanner.Scan()
	person.age, _ = strconv.Atoi(scanner.Text())
	return person
}

func PrintPeople(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprint(writer, "Which file to open: ")
	scanner.Scan()
	filename := scanner.Text()
	fmt.Fprintf(writer, "Reading data from %s\n", filename)
	fmt.Fprintln(writer)
	file, _ := os.Open(filename)
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	people := make([]*Person, 5)
	for i := 0; i < len(people); i++ {
		people[i] = ReadPerson(fileScanner)
	}
	for i := 0; i < len(people); i++ {
		fmt.Fprintf(writer, "%s is %d\n", people[i].name, people[i].age)
	}
}

func main() {
	PrintPeople(os.Stdin, os.Stdout)
}
