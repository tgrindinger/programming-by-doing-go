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
		{ "1", 1, 2 },
		{ "2", 2, 3 },
		{ "3", 3, 5 },
		{ "4", 4, 7 },
		{ "5", 5, 11 },
		{ "6", 6, 13 },
		{ "7", 7, 17 },
		{ "8", 8, 19 },
		{ "10001", 10001, 104743 },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			got := FindNthPrime(tc.input)
			if got != tc.want {
				t.Errorf("\ngot  %d\nwant %d", got, tc.want)
			}
		})
	}
}
