package main

import (
	"testing"
)

func TestSumEvenFibonacciNumbers(t *testing.T) {
	cases := []struct {
		desc  string
		input int
		want  int
	}{
		{ "0", 0, 0 },
		{ "1", 1, 0 },
		{ "2", 2, 2 },
		{ "3", 3, 2 },
		{ "4", 4, 2 },
		{ "5", 5, 7 },
		{ "6", 6, 7 },
		{ "7", 7, 7 },
		{ "8", 8, 7 },
		{ "9", 9, 7 },
		{ "10", 10, 7 },
		{ "11", 11, 7 },
		{ "12", 12, 7 },
		{ "13", 13, 20 },
		{ "14", 14, 20 },
		{ "4000000", 4000000, 5702886 },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			got := SumEvenFibonacciNumbers(tc.input)
			if got != tc.want {
				t.Errorf("\ngot  %d\nwant %d", got, tc.want)
			}
		})
	}
}
