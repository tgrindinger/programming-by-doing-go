package main

import (
	"testing"
)

func TestLargestProductInSeries(t *testing.T) {
	cases := []struct {
		desc   string
		number string
		length int
		want   int64
	}{
		{ "1", "1234567890987654321", 1, 9 },
		{ "2", "1234567890987654321", 2, 8 * 9 },
		{ "3", "1234567890987654321", 3, 7 * 8 * 9 },
		{ "4", "1234567890987654321", 4, 6 * 7 * 8 * 9 },
		{ "5", "1234567890987654321", 5, 5 * 6 * 7 * 8 * 9 },
		{ "6", "1234567890987654321", 6, 4 * 5 * 6 * 7 * 8 * 9 },
		{ "7", "1234567890987654321", 7, 3 * 4 * 5 * 6 * 7 * 8 * 9 },
		{ "8", "1234567890987654321", 8, 2 * 3 * 4 * 5 * 6 * 7 * 8 * 9 },
		{ "9", "1234567890987654321", 9, 2 * 3 * 4 * 5 * 6 * 7 * 8 * 9 },
		{ "10", "1234567890987654321", 10, 0 },
		{ "11", "1234567890987654321", 11, 0 },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			got := LargestProductInSeries(tc.number, tc.length)
			if got != tc.want {
				t.Errorf("\ngot  %d\nwant %d", got, tc.want)
			}
		})
	}
}
