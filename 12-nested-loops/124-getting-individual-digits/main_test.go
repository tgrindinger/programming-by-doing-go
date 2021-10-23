package main

import (
	"strings"
	"testing"
)

func TestPrintPuzzleNumbers(t *testing.T) {
	cases := []struct {
		desc     string
		wants    []string
	}{
		{ "default", []string {
				"10, 1+0 = 1",
				"15, 1+5 = 6",
				"22, 2+2 = 4",
				"97, 9+7 = 16",
				"98, 9+8 = 17",
				"99, 9+9 = 18",
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
