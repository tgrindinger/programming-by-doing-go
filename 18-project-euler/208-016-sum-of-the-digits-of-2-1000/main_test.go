package main

import (
	"testing"
)

func TestSumPowerOf2Digits(t *testing.T) {
	cases := []struct{
		desc  string
		input int64
		want  int64
	}{
		{ "4", 4, 7 },
		{ "5", 5, 5 },
		{ "6", 6, 10 },
		{ "7", 7, 11 },
		{ "8", 8, 13 },
		{ "9", 9, 8 },
		{ "15", 15, 26 },
		{ "1000", 1000, 1366 },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			got := sumPowerOf2Digits(tc.input)
			if got != tc.want {
				t.Errorf("\ngot  %d\nwant %d", got, tc.want)
			}
		})
	}
}
