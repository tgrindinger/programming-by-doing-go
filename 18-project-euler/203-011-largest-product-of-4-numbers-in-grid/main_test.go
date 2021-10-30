package main

import (
	"testing"
)

func TestLargestAdjacentProduct(t *testing.T) {
	cases := []struct {
		desc  string
		input int
		grid  [][]int
		want  int
	}{
		{ "2x2 top row", 2, [][]int{ { 1, 1 }, { 0, 0 } }, 1 },
		{ "2x2 bottom row", 2, [][]int{ { 0, 0 }, { 1, 1 } }, 1 },
		{ "2x2 left column", 2, [][]int{ { 1, 0 }, { 1, 0 } }, 1 },
		{ "2x2 right column", 2, [][]int{ { 0, 1 }, { 0, 1 } }, 1 },
		{ "2x2 back slash", 2, [][]int{ { 1, 0 }, { 0, 1 } }, 1 },
		{ "2x2 forward slash", 2, [][]int{ { 0, 1 }, { 1, 0 } }, 1 },
		{ "3x3 top row", 2, [][]int{ { 0, 2, 3 }, { 0, 0, 0 }, { 0, 0, 0 } }, 6 },
		{ "3x3 bottom row", 2, [][]int{ { 0, 0, 0 }, { 0, 0, 0 }, { 0, 2, 3 } }, 6 },
		{ "3x3 left column", 2, [][]int{ { 0, 0, 0 }, { 2, 0, 0 }, { 3, 0, 0 } }, 6 },
		{ "3x3 right column", 2, [][]int{ { 0, 0, 0 }, { 0, 0, 2 }, { 0, 0, 3 } }, 6 },
		{ "3x3 back slash", 2, [][]int{ { 0, 0, 0 }, { 0, 2, 0 }, { 0, 0, 3 } }, 6 },
		{ "3x3 forward slash", 2, [][]int{ { 0, 0, 0 }, { 0, 0, 3 }, { 0, 2, 0 } }, 6 },
		{ "grid 4", 4, grid, 70600674 },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			got := LargestAdjacentProduct(tc.grid, tc.input)
			if got != tc.want {
				t.Errorf("\ngot  %d\nwant %d", got, tc.want)
			}
		})
	}
}
