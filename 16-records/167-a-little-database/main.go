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
	students := make([]*Student, 3)
	students[0] = ReadStudent(scanner, writer, "first")
	students[1] = ReadStudent(scanner, writer, "second")
	students[2] = ReadStudent(scanner, writer, "third")
	fmt.Fprintf(writer, "The names are: %s %s %s\n", students[0].name, students[1].name, students[2].name)
	fmt.Fprintf(writer, "The grades are: %d %d %d\n", students[0].grade, students[1].grade, students[2].grade)
	fmt.Fprintf(writer, "The averages are: %.1f %.1f %.1f\n", students[0].avg, students[1].avg, students[2].avg)
	fmt.Fprintln(writer)
	fmt.Fprintf(writer, "The average for the three students is: %f\n", (students[0].avg + students[1].avg + students[2].avg) / 3)
}

func main() {
	PrintStudents(os.Stdin, os.Stdout)
}
