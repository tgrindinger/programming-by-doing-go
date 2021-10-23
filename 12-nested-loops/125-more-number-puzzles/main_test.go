package main

import (
	"strings"
	"testing"
)

func TestPrintPuzzleNumbers(t *testing.T) {
	cases := []struct {
		desc  string
		input string
		wants []string
	}{
		{ "default", "1\n2\n3\n", []string {
				"29",
				"38",
				"39",
				"47",
				"48",
				"49",
				"56",
				"54",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			PrintPuzzleNumbers(reader, writer)
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
