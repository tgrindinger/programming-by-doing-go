package main

import (
	"testing"
)

func TestLettersInWords(t *testing.T) {
	cases := []struct{
		desc  string
		input int
		want  int
	}{
		{ "1", 1, 3 },
		{ "2", 2, 3 },
		{ "3", 3, 5 },
		{ "4", 4, 4 },
		{ "5", 5, 4 },
		{ "6", 6, 3 },
		{ "7", 7, 5 },
		{ "8", 8, 5 },
		{ "9", 9, 4 },
		{ "10", 10, 3 },
		{ "11", 11, 6 },
		{ "12", 12, 6 },
		{ "13", 13, 8 },
		{ "14", 14, 8 },
		{ "15", 15, 7 },
		{ "16", 16, 7 },
		{ "17", 17, 9 },
		{ "18", 18, 8 },
		{ "19", 19, 8 },
		{ "20", 20, 6 },
		{ "21", 21, 9 },
		{ "30", 30, 6 },
		{ "40", 40, 5 },
		{ "50", 50, 5 },
		{ "60", 60, 5 },
		{ "70", 70, 7 },
		{ "80", 80, 6 },
		{ "90", 90, 6 },
		{ "100", 100, 10 },
		{ "101", 101, 16 },
		{ "115", 115, 20 },
		{ "200", 200, 10 },
		{ "342", 342, 23 },
		{ "1000", 1000, 11 },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			got := lettersInWords(tc.input)
			if got != tc.want {
				t.Errorf("\ngot  %d\nwant %d", got, tc.want)
			}
		})
	}
}

func TestLettersInNumbersUpTo(t *testing.T) {
	cases := []struct{
		desc  string
		input int
		want  int
	}{
		{ "1", 1, 3 },
		{ "2", 2, 6 },
		{ "3", 3, 11 },
		{ "4", 4, 15 },
		{ "5", 5, 19 },
		{ "6", 6, 22 },
		{ "7", 7, 27 },
		{ "8", 8, 32 },
		{ "9", 9, 36 },
		{ "1000", 1000, 21124 },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			got := lettersInNumbersUpTo(tc.input)
			if got != tc.want {
				t.Errorf("\ngot  %d\nwant %d", got, tc.want)
			}
		})
	}
}
