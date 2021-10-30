package main

import (
	"testing"
)

func TestCollatzLength(t *testing.T) {
	cases := []struct{
		desc  string
		start int
		want  int
	}{
		{ "1", 1, 1 },
		{ "2", 2, 2 },
		{ "3", 3, 8 },
		{ "13", 13, 10 },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			got := collatzLength(tc.start)
			if got != tc.want {
				t.Errorf("\ngot  %d\nwant %d", got, tc.want)
			}
		})
	}
}


func TestLongestCollatzLength(t *testing.T) {
	cases := []struct{
		desc  string
		start int
		want  int
	}{
		{ "1", 1, 0 },
		{ "2", 2, 1 },
		{ "3", 3, 2 },
		{ "4", 4, 3 },
		{ "5", 5, 3 },
		{ "6", 6, 3 },
		{ "7", 7, 6 },
		{ "8", 8, 7 },
		{ "9", 9, 7 },
		{ "10", 10, 9 },
		{ "11", 11, 9 },
		{ "12", 12, 9 },
		{ "13", 13, 9 },
		{ "14", 14, 9 },
		{ "1000000", 1000000, 837799 },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			got := longestCollatzLength(tc.start)
			if got != tc.want {
				t.Errorf("\ngot  %d\nwant %d", got, tc.want)
			}
		})
	}
}
