package main

import (
	"strings"
	"testing"
)

func TestPrintStuff(t *testing.T) {
	cases := []struct {
		desc  string
		wants []string
	}{
		{ "default", []string {
				"A 1",
				"A 2",
				"A 3",
				"B 1",
				"B 2",
				"B 3",
				"C 1",
				"C 2",
				"C 3",
				"D 1",
				"D 2",
				"D 3",
				"E 1",
				"E 2",
				"E 3",
				"1-1 1-2 1-3 2-1 2-2 2-3 3-1 3-2 3-3 ",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			writer := &strings.Builder{}
			PrintStuff(writer)
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
