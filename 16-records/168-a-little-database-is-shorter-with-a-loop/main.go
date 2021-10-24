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

func ReadStudent(scanner *bufio.Scanner, writer io.Writer, index int) *Student {
	student := &Student{}
	fmt.Fprintf(writer, "Enter student %d's name: ", index)
	scanner.Scan()
	student.name = scanner.Text()
	fmt.Fprintf(writer, "Enter student %d's grade: ", index)
	scanner.Scan()
	student.grade, _ = strconv.Atoi(scanner.Text())
	fmt.Fprintf(writer, "Enter student %d's average: ", index)
	scanner.Scan()
	student.avg, _ = strconv.ParseFloat(scanner.Text(), 64)
	fmt.Fprintln(writer)
	return student
}

func PrintStudents(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	students := make([]*Student, 3)
	for i := 0; i < len(students); i++ {
		students[i] = ReadStudent(scanner, writer, i + 1)
	}
	fmt.Fprintf(writer, "The names are: %s %s %s\n", students[0].name, students[1].name, students[2].name)
	fmt.Fprintf(writer, "The grades are: %d %d %d\n", students[0].grade, students[1].grade, students[2].grade)
	fmt.Fprintf(writer, "The averages are: %.1f %.1f %.1f\n", students[0].avg, students[1].avg, students[2].avg)
	fmt.Fprintln(writer)
	fmt.Fprintf(writer, "The average for the three students is: %f\n", (students[0].avg + students[1].avg + students[2].avg) / 3)
}

func main() {
	PrintStudents(os.Stdin, os.Stdout)
}
