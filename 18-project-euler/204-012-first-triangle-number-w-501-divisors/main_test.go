package main

import (
	"testing"
)

func TestFirstTriangleNumberWithOverNDivisors(t *testing.T) {
	cases := []struct {
		desc  string
		input int
		want  int
	}{
		{ "1", 1, 3 },
		{ "2", 2, 6 },
		{ "3", 3, 6 },
		{ "4", 4, 28 },
		{ "5", 5, 28 },
		{ "6", 6, 36 },
		{ "7", 7, 36 },
		{ "500", 500, 76576500 },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			got := FirstTriangleNumberWithOverNDivisors(tc.input)
			if got != tc.want {
				t.Errorf("\ngot  %d\nwant %d", got, tc.want)
			}
		})
	}
}
