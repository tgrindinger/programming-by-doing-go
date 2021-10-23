package main

import (
	"fmt"
	"io"
	"os"
)

func PrintWeekday(reader io.Reader, writer io.Writer) {
		fmt.Fprintln(writer, "Welcome to Mr. Mitchell's fantastic birth-o-meter!")
		fmt.Fprintln(writer)
		fmt.Fprintln(writer, "All you have to do is enter your birth date, and it will")
		fmt.Fprintln(writer, "tell you the day of the week on which you were born.")
		fmt.Fprintln(writer)
		fmt.Fprintln(writer, "Some automatic tests....")
		fmt.Fprintf(writer, "12 10 2003 => %s\n", weekday(12,10,2003))
		fmt.Fprintf(writer, " 2 13 1976 => %s\n", weekday(2,13,1976))
		fmt.Fprintf(writer, " 2 13 1977 => %s\n", weekday(2,13,1977))
		fmt.Fprintf(writer, " 7  2 1974 => %s\n", weekday(7,2,1974))
		fmt.Fprintf(writer, " 1 15 2003 => %s\n", weekday(1,15,2003))
		fmt.Fprintf(writer, "10 13 2000 => %s\n", weekday(10,13,2000))
		fmt.Fprintln(writer)

		fmt.Fprintln(writer, "Now it's your turn!  What's your birthday?")
		fmt.Fprint(writer, "Birth date (mm dd yyyy): ")
		var mm, dd, yyyy int
		fmt.Fscanf(reader, "%d %d %d", &mm, &dd, &yyyy)
		fmt.Fprintf(writer, "You were born on %s!\n", weekday(mm, dd, yyyy));
}

func main() {
	PrintWeekday(os.Stdin, os.Stdout)
}
