package main

import (
	"fmt"
	"io"
	"os"
)

func monthName(month int) string {
	var name string
	if month == 1 {
		name = "January"
	} else if month == 2 {
		name = "February"
	} else if month == 3 {
		name = "March"
	} else if month == 4 {
		name = "April"
	} else if month == 5 {
		name = "May"
	} else if month == 6 {
		name = "June"
	} else if month == 7 {
		name = "July"
	} else if month == 8 {
		name = "August"
	} else if month == 9 {
		name = "September"
	} else if month == 10 {
		name = "October"
	} else if month == 11 {
		name = "November"
	} else if month == 12 {
		name = "December"
	} else {
		name = "error"
	}
	return name
}

func monthOffset(month int) int {
	var offset int
	if month == 1 {
		offset = 1
	} else if month == 2 {
		offset = 4
	} else if month == 3 {
		offset = 4
	} else if month == 4 {
		offset = 0
	} else if month == 5 {
		offset = 2
	} else if month == 6 {
		offset = 5
	} else if month == 7 {
		offset = 0
	} else if month == 8 {
		offset = 3
	} else if month == 9 {
		offset = 6
	} else if month == 10 {
		offset = 1
	} else if month == 11 {
		offset = 4
	} else if month == 12 {
		offset = 6
	} else {
		offset = -1
	}
	return offset
}

func isLeap(year int) bool {
	result := false
	if year % 400 == 0 {
		result = true
	} else if year % 100 == 0 {
		result = false
	} else if year % 4 == 0 {
		result = true
	}
	return result
}

func weekdayName(day int) string {
	result := "error"
	if day == 1 {
		result = "Sunday"
	} else if day == 2 {
		result = "Monday"
	} else if day == 3 {
		result = "Tuesday"
	} else if day == 4 {
		result = "Wednesday"
	} else if day == 5 {
		result = "Thursday"
	} else if day == 6 {
		result = "Friday"
	} else if day == 7 {
		result = "Saturday"
	}
	return result
}

func weekday(mm, dd, yyyy int) string {
	yy := yyyy - 1900
	total := yy / 4 + yy + dd + monthOffset(mm)
	if isLeap(yyyy) && (mm == 1 || mm == 2) {
		total--
	}
	name := weekdayName(total % 7)
	return fmt.Sprintf("%s, %s %d, %d", name, monthName(mm), dd, yyyy)
}

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
