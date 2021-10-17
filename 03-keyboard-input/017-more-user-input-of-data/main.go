package main

import (
	"fmt"
	"io"
	"os"
)

func AskQuestions(reader io.Reader, writer io.Writer) {
	var firstName, lastName, gradeString, idString, login, gpaString string
	var grade, id int
	var gpa float64
	fmt.Fprintln(writer, "Please enter the following information so I can sell it for a profit!")
	fmt.Fprintln(writer)
	fmt.Fprint(writer, "First name: ")
	fmt.Fscanln(reader, &firstName)
	fmt.Fprint(writer, "Last name: ")
	fmt.Fscanln(reader, &lastName)
	fmt.Fprint(writer, "Grade (9-12): ")
	fmt.Fscanln(reader, &gradeString)
	fmt.Sscanf(gradeString, "%d", &grade)
	fmt.Fprint(writer, "Student ID: ")
	fmt.Fscanln(reader, &idString)
	fmt.Sscanf(idString, "%d", &id)
	fmt.Fprint(writer, "Login: ")
	fmt.Fscanln(reader, &login)
	fmt.Fprint(writer, "GPA (0.0-4.0): ")
	fmt.Fscanln(reader, &gpaString)
	fmt.Sscanf(gpaString, "%f", &gpa)
	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "Your information:")
	fmt.Fprintf(writer, "        Login: %s\n", login)
	fmt.Fprintf(writer, "        ID:    %d\n", id)
	fmt.Fprintf(writer, "        Name:  %s, %s\n", lastName, firstName)
	fmt.Fprintf(writer, "        GPA:   %.2f\n", gpa)
	fmt.Fprintf(writer, "        Grade: %d\n", grade)
}

func main() {
	AskQuestions(os.Stdin, os.Stdout)
}
