package main

import (
	"fmt"
)

func monthOffset(month int) int {
	switch month {
	case 1: return 1
	case 2: return 4
	case 3: return 4
	case 4: return 0
	case 5: return 2
	case 6: return 5
	case 7: return 0
	case 8: return 3
	case 9: return 6
	case 10: return 1
	case 11: return 4
	case 12: return 6
	default: return -1
	}
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

func isSunday(mm, dd, yyyy int) bool {
	yy := yyyy - 1900
	total := yy / 4 + yy + dd + monthOffset(mm)
	if isLeap(yyyy) && (mm == 1 || mm == 2) {
		total--
	}
	return total % 7 == 1
}

func numOfSundaysOnFirstDayOfMonth(year int) int {
	count := 0
	for month := 1; month <= 12; month++ {
		if isSunday(month, 1, year) {
			count++
		}
	}
	return count
}

func numOfSundaysOnFirstDayOfMonthBetweenYears(startYear int, endYear int) int {
	count := 0
	for year := startYear; year <= endYear; year++ {
		count += numOfSundaysOnFirstDayOfMonth(year)
	}
	return count
}

func main() {
	fmt.Printf("Count Sundays that occur on the first day of the month from 1901 to: ")
	var maxYear int
	fmt.Scanln(&maxYear)
	result := numOfSundaysOnFirstDayOfMonthBetweenYears(1901, maxYear)
	fmt.Printf("The result is %d\n", result)
}
