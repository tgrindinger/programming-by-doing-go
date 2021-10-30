package main

import (
	"testing"
)

func TestFindSumOfMultiplesOf3Or5(t *testing.T) {
	cases := []struct {
		desc  string
		input int
		want  int
	}{
		{ "10", 10, 23 },
		{ "11", 11, 33 },
		{ "12", 12, 33 },
		{ "13", 13, 45 },
		{ "14", 14, 45 },
		{ "15", 15, 45 },
		{ "16", 16, 60 },
		{ "100", 100, 2318 },
		{ "1000", 1000, 233168 },
		{ "10000", 10000, 23331668 },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			got := FindSumOfMultiplesOf3Or5(tc.input)
			if got != tc.want {
				t.Errorf("\ngot  %d\nwant %d", got, tc.want)
			}
		})
	}
}
