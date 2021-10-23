package main

import "fmt"

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
	if year%400 == 0 {
		result = true
	} else if year%100 == 0 {
		result = false
	} else if year%4 == 0 {
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
	total := yy/4 + yy + dd + monthOffset(mm)
	if isLeap(yyyy) && (mm == 1 || mm == 2) {
		total--
	}
	name := weekdayName(total % 7)
	return fmt.Sprintf("%s, %s %d, %d", name, monthName(mm), dd, yyyy)
}
