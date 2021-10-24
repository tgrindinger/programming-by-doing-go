package main

import (
	"strings"
	"testing"
)

func TestPrintGrades(t *testing.T) {
	cases := []struct {
		desc  string
		input string
		wants []string
	}{
		{ "gb", "gb.txt\n", []string {
				"\t171 1 47 F",
				"\t171 2 64 D",
				"\t171 3 86 B",
				"\t171 4 26 F",
				"\t938 3 53 F",
				"\t938 4 28 F",
				"\t938 5 52 F",
				"\t938 6 43 F",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			PrintGrades(reader, writer)
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
