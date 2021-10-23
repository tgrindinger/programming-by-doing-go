package main

import (
	"strings"
	"testing"
)

func TestRunCalculator(t *testing.T) {
	cases := []struct {
		desc  string
		input string
		wants []string
	}{
		{ "default", "2 + 3\n4 * 9\n5 - 2\n12 / 2\n0 + 2\n", []string {
				"> 5.000000",
				"> 36.000000",
				"> 3.000000",
				"> 6.000000",
				"> Bye, now.",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			RunCalculator(reader, writer)
			gots := strings.Split(writer.String(), "\n")
			for _, want := range tc.wants {
				if !contains(gots, want) {
					t.Errorf("\ngot  %q\nwant %q", gots, want)
				}
			}
		})
	}
}

func contains(arr []string, target string) bool {
	for _, s := range arr {
		if s == target {
			return true
		}
	}
	return false
}
