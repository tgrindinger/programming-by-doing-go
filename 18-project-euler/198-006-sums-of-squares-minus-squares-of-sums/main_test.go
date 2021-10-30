package main

import (
	"testing"
)

func TestDifferenceBetweenSumOfSquaresAndSquareOfSum(t *testing.T) {
	cases := []struct {
		desc  string
		input int
		want  int
	}{
		{ "2", 2, (1 + 2) * (1 + 2) - (1 + 4) },
		{ "3", 3, (1 + 2 + 3) * (1 + 2 + 3) - (1 + 4 + 9) },
		{ "4", 4, (1 + 2 + 3 + 4) * (1 + 2 + 3 + 4) - (1 + 4 + 9 + 16) },
		{ "5", 5, (1 + 2 + 3 + 4 + 5) * (1 + 2 + 3 + 4 + 5) - (1 + 4 + 9 + 16 + 25) },
		{ "6", 6, (1 + 2 + 3 + 4 + 5 + 6) * (1 + 2 + 3 + 4 + 5 + 6) - (1 + 4 + 9 + 16 + 25 + 36) },
		{ "7", 7, (1 + 2 + 3 + 4 + 5 + 6 + 7) * (1 + 2 + 3 + 4 + 5 + 6 + 7) - (1 + 4 + 9 + 16 + 25 + 36 + 49) },
		{ "8", 8, (1 + 2 + 3 + 4 + 5 + 6 + 7 + 8) * (1 + 2 + 3 + 4 + 5 + 6 + 7 + 8) - (1 + 4 + 9 + 16 + 25 + 36 + 49 + 64) },
		{ "9", 9, (1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9) * (1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9) - (1 + 4 + 9 + 16 + 25 + 36 + 49 + 64 + 81) },
		{ "10", 10, 2640 },
		{ "100", 100, 25164150 },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			got := DifferenceBetweenSumOfSquaresAndSquareOfSum(tc.input)
			if got != tc.want {
				t.Errorf("\ngot  %d\nwant %d", got, tc.want)
			}
		})
	}
}
