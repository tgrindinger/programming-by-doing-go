package main

import "fmt"

func WeekdayName(weekday int) string {
	result := ""
	if weekday == 1 {
		result = "Sunday"
	} else if weekday == 2 {
		result = "Monday"
	} else if weekday == 3 {
		result = "Tuesday"
	} else if weekday == 4 {
		result = "Wednesday"
	} else if weekday == 5 {
		result = "Thursday"
	} else if weekday == 6 {
		result = "Friday"
	} else if weekday == 7 {
		result = "Saturday"
	} else if weekday == 0 {
		result = "Saturday"
	} else {
		result = "error"
	}
	return result
}

func main() {
	fmt.Printf("Weekday 1: %s\n", WeekdayName(1))
	fmt.Printf("Weekday 2: %s\n", WeekdayName(2))
	fmt.Printf("Weekday 3: %s\n", WeekdayName(3))
	fmt.Printf("Weekday 4: %s\n", WeekdayName(4))
	fmt.Printf("Weekday 5: %s\n", WeekdayName(5))
	fmt.Printf("Weekday 6: %s\n", WeekdayName(6))
	fmt.Printf("Weekday 7: %s\n", WeekdayName(7))
	fmt.Printf("Weekday 0: %s\n", WeekdayName(0))
	fmt.Println()
	fmt.Printf("Weekday 43: %s\n", WeekdayName(43))
	fmt.Printf("Weekday 17: %s\n", WeekdayName(17))
	fmt.Printf("Weekday -1: %s\n", WeekdayName(-1))
}
