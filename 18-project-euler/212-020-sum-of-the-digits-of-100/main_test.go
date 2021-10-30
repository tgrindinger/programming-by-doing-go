package main

import (
	"testing"
)

func TestNumberPaths(t *testing.T) {
	cases := []struct{
		desc  string
		input int
		want  int
	}{
		{ "1", 1, 1 },
		{ "2", 2, 2 },
		{ "3", 3, 6 },
		{ "4", 4, 6 },
		{ "5", 5, 3 },
		{ "6", 6, 9 },
		{ "10", 10, 27 },
		{ "100", 100, 648 },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			got := sumOfDigitsOfFactorial(tc.input)
			if got != tc.want {
				t.Errorf("got %d want %d", got, tc.want)
			}
		})
	}
}
