package main

import (
	"testing"
)

func TestSmallestDivisibleByUpTo(t *testing.T) {
	cases := []struct {
		desc  string
		input int
		want  int64
	}{
		{ "3", 3, 6 },
		{ "4", 4, 12 },
		{ "5", 5, 60 },
		{ "6", 6, 60 },
		{ "7", 7, 420 },
		{ "8", 8, 840 },
		{ "10", 10, 2520 },
		{ "20", 20, 232792560 },
		{ "30", 30, 2329089562800 },
		{ "40", 40, 5342931457063200 },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			got := SmallestDivisibleByUpTo(tc.input)
			if got != tc.want {
				t.Errorf("\ngot  %d\nwant %d", got, tc.want)
			}
		})
	}
}
