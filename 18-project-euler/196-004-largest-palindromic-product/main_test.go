package main

import (
	"testing"
)

func TestLargestPalindromicProduct(t *testing.T) {
	cases := []struct {
		desc  string
		input int
		want  int
	}{
		{ "1", 1, 9 },
		{ "2", 2, 9009 },
		{ "3", 3, 906609 },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			got := LargestPalindromicProduct(tc.input)
			if got != tc.want {
				t.Errorf("\ngot  %d\nwant %d", got, tc.want)
			}
		})
	}
}
