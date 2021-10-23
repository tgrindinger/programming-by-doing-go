package main

import (
	"strings"
	"testing"
)

func TestPrintPuzzleNumbers(t *testing.T) {
	cases := []struct {
		desc  string
		wants []string
	}{
		{ "default", []string {
				"8 12 5 20",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			writer := &strings.Builder{}
			PrintPuzzleNumbers(writer)
			gots := strings.Split(writer.String(), "\n")
			for _, want := range tc.wants {
				if !contains(gots, want) {
					t.Errorf("\ngot  %s\nwant %s", gots, want)
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
