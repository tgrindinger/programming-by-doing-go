package main

import (
	"testing"
)

func TestLargestPrimeFactor(t *testing.T) {
	cases := []struct {
		desc  string
		input int64
		want  int64
	}{
		{ "2", 2, 2 },
		{ "5", 5, 5 },
		{ "6", 6, 3 },
		{ "15", 15, 5 },
		{ "88", 88, 11 },
		{ "13195", 13195, 29 },
		{ "600851475143", 600851475143, 6857 },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			got := LargestPrimeFactor(tc.input)
			if got != tc.want {
				t.Errorf("\ngot  %d\nwant %d", got, tc.want)
			}
		})
	}
}
