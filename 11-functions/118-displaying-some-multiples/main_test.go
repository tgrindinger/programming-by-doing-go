package main

import (
	"strings"
	"testing"
)

func TestPrintMultiples(t *testing.T) {
	cases := []struct {
		desc  string
		input string
		wants []string
	}{
		{ "7", "7\n", []string {
				"7x1 = 7",
				"7x2 = 14",
				"7x3 = 21",
				"7x4 = 28",
				"7x5 = 35",
				"7x6 = 42",
				"7x7 = 49",
				"7x8 = 56",
				"7x9 = 63",
				"7x10 = 70",
				"7x11 = 77",
				"7x12 = 84",
			},
		},
		{ "2", "2\n", []string {
				"2x1 = 2",
				"2x2 = 4",
				"2x3 = 6",
				"2x4 = 8",
				"2x5 = 10",
				"2x6 = 12",
				"2x7 = 14",
				"2x8 = 16",
				"2x9 = 18",
				"2x10 = 20",
				"2x11 = 22",
				"2x12 = 24",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			PrintMultiples(reader, writer)
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
