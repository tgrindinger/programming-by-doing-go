package main

import (
	"fmt"
	"io"
	"os"
)

func PrintSchedule(writer io.Writer) {
	course1 := "                          English III"
	course2 := "                          Precalculus"
	course3 := "                         Music Theory"
	course4 := "                        Biotechnology"
	course5 := "           Principles of Technology I"
	course6 := "                             Latin II"
	course7 := "                        AP US History"
	course8 := "Business Computer Information Systems"
	teacher1 := "      Ms. Lapan"
	teacher2 := "    Mrs. Gideon"
	teacher3 := "      Mr. Davis"
	teacher4 := "     Ms. Palmer"
	teacher5 := "     Ms. Garcia"
	teacher6 := "   Mrs. Barnett"
	teacher7 := "Ms. Johannessen"
	teacher8 := "      Mr. James"
	fmt.Fprintln(writer, "+-------------------------------------------------------------+")
	fmt.Fprintln(writer, "|", 1, "|", course1, "|", teacher1, "|")
	fmt.Fprintln(writer, "|", 2, "|", course2, "|", teacher2, "|")
	fmt.Fprintln(writer, "|", 3, "|", course3, "|", teacher3, "|")
	fmt.Fprintln(writer, "|", 4, "|", course4, "|", teacher4, "|")
	fmt.Fprintln(writer, "|", 5, "|", course5, "|", teacher5, "|")
	fmt.Fprintln(writer, "|", 6, "|", course6, "|", teacher6, "|")
	fmt.Fprintln(writer, "|", 7, "|", course7, "|", teacher7, "|")
	fmt.Fprintln(writer, "|", 8, "|", course8, "|", teacher8, "|")
	fmt.Fprintln(writer, "+-------------------------------------------------------------+")
}

func main() {
	PrintSchedule(os.Stdout)
}