package main

import (
	"strings"
	"testing"
)

func TestPrintDistances(t *testing.T) {
	cases := []struct {
		desc string
		wants []string
	}{
		{ "default", []string {
				"(-2,1) to (1,5) => 5.000000",
				"(-2,-3) to (-4,4) => 7.280110",
				"(2,-3) to (-1,-2) => 3.162278",
				"(4,5) to (4,5) => 0.000000",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			writer := &strings.Builder{}
			PrintDistances(writer)
			got := writer.String()
			for _, want := range tc.wants {
				if !strings.Contains(got, want) {
					t.Errorf("\ngot  %q\nwant %q", got, want)
				}
			}
		})
	}
}
