package main

import (
	"testing"
)

func TestFindNthPrime(t *testing.T) {
	cases := []struct {
		desc  string
		input int
		want  int
	}{
		{ "3", 3, 2 },
		{ "4", 4, 5 },
		{ "5", 5, 5 },
		{ "6", 6, 10 },
		{ "7", 7, 10 },
		{ "8", 8, 17 },
		{ "9", 9, 17 },
		{ "10", 10, 17 },
		{ "11", 11, 17 },
		{ "12", 12, 28 },
		{ "13", 13, 28 },
		{ "14", 14, 41 },
		{ "2000000", 2000000, 142913828922 },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			got := SumOfPrimes(tc.input)
			if got != tc.want {
				t.Errorf("\ngot  %d\nwant %d", got, tc.want)
			}
		})
	}
}
