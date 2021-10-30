package main

import (
	"testing"
)

func TestNumberOfSundaysOnFirstDayOfMonth(t *testing.T) {
	cases := []struct{
		desc  string
		input int
		want  int
	}{
		{ "1901", 1901, 2 },
		{ "1902", 1902, 1 },
		{ "1903", 1903, 3 },
		{ "1904", 1904, 1 },
		{ "1905", 1905, 2 },
		{ "1906", 1906, 2 },
		{ "1907", 1907, 2 },
		{ "1908", 1908, 2 },
		{ "1909", 1909, 1 },
		{ "1910", 1910, 1 },
		{ "1911", 1911, 2 },
		{ "1912", 1912, 2 },
		{ "2000", 2000, 1 },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			got := numOfSundaysOnFirstDayOfMonth(tc.input)
			if got != tc.want {
				t.Errorf("\ngot  %d\nwant %d", got, tc.want)
			}
		})
	}
}

func TestNumOfSundaysOnFirstDayOfMonthBetweenYears(t *testing.T) {
	cases := []struct{
		desc  string
		start int
		end   int
		want  int
	}{
		{ "2000", 1901, 2000, 171 },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			got := numOfSundaysOnFirstDayOfMonthBetweenYears(tc.start, tc.end)
			if got != tc.want {
				t.Errorf("\ngot  %d\nwant %d", got, tc.want)
			}
		})
	}
}
