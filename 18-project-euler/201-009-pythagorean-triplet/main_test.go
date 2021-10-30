package main

import (
	"testing"
)

func TestLargestProductInSeries(t *testing.T) {
	cases := []struct {
		desc   string
		target int
		want   int
	}{
		{ "12", 12, 60 },
		{ "1000", 1000, 31875000 },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			got := PythagoreanTriplet(tc.target)
			if got != tc.want {
				t.Errorf("\ngot  %d\nwant %d", got, tc.want)
			}
		})
	}
}
