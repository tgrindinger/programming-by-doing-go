package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Student struct {
	name  string
	grade int
	avg   float64
}

func ReadStudent(scanner *bufio.Scanner, writer io.Writer, label string) *Student {
	student := &Student{}
	fmt.Fprintf(writer, "Enter the %s student's name: ", label)
	scanner.Scan()
	student.name = scanner.Text()
	fmt.Fprintf(writer, "Enter the %s student's grade: ", label)
	scanner.Scan()
	student.grade, _ = strconv.Atoi(scanner.Text())
	fmt.Fprintf(writer, "Enter the %s student's average: ", label)
	scanner.Scan()
	student.avg, _ = strconv.ParseFloat(scanner.Text(), 64)
	fmt.Fprintln(writer)
	return student
}

func PrintStudents(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	student1 := ReadStudent(scanner, writer, "first")
	student2 := ReadStudent(scanner, writer, "second")
	student3 := ReadStudent(scanner, writer, "third")
	fmt.Fprintf(writer, "The names are: %s %s %s\n", student1.name, student2.name, student3.name)
	fmt.Fprintf(writer, "The grades are: %d %d %d\n", student1.grade, student2.grade, student3.grade)
	fmt.Fprintf(writer, "The averages are: %.1f %.1f %.1f\n", student1.avg, student2.avg, student3.avg)
	fmt.Fprintln(writer)
	fmt.Fprintf(writer, "The average for the three students is: %f\n", (student1.avg + student2.avg + student3.avg) / 3)
}

func main() {
	PrintStudents(os.Stdin, os.Stdout)
}
