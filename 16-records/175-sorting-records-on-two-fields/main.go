package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Grade struct {
	id     int
	number int
	grade  float64
	letter string
}

func ReadGrade(scanner *bufio.Scanner) *Grade {
	grade := &Grade{}
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &grade.id, &grade.number, &grade.grade, &grade.letter)
	return grade
}

func SortGrades(grades []*Grade) {
	for i := 0; i < len(grades); i++ {
		for j := i + 1; j < len(grades); j++ {
			if grades[j].id < grades[i].id ||
					(grades[j].id == grades[i].id && grades[j].number < grades[i].number) {
				temp := grades[i]
				grades[i] = grades[j]
				grades[j] = temp
			}
		}
	}
}

func PrintGrades(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprint(writer, "Open which file: ")
	scanner.Scan()
	filename := scanner.Text()
	file, _ := os.Open(filename)
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	grades := make([]*Grade, 30)
	for i := 0; i < len(grades); i++ {
		grades[i] = ReadGrade(fileScanner)
	}
	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "Data loaded.")
	SortGrades(grades)
	fmt.Fprintln(writer, "Data sorted.")
	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "Here are the sorted grades:")
	for i := 0; i < len(grades); i++ {
		fmt.Fprintf(writer, "\t%d %d %.0f %s\n", grades[i].id, grades[i].number, grades[i].grade, grades[i].letter)
	}
}

func main() {
	PrintGrades(os.Stdin, os.Stdout)
}
