package main

import "testing"

func TestWeekdayName(t *testing.T) {
	cases := []struct {
		weekday int
		name    string
	}{
		{ 0, "Saturday" },
		{ 1, "Sunday" },
		{ 2, "Monday" },
		{ 3, "Tuesday" },
		{ 4, "Wednesday" },
		{ 5, "Thursday" },
		{ 6, "Friday" },
		{ 7, "Saturday" },
		{ 8, "error" },
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := WeekdayName(tc.weekday)
			want := tc.name
			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		})
	}
}