package main

import (
	"strings"
	"testing"
)

func TestPrintDogs(t *testing.T) {
	cases := []struct {
		desc  string
		input string
		wants []string
	}{
		{ "dogs", "dogs.txt\n", []string {
				"First dog: Yorkie, 4, 14",
				"Second dog: Great_Dane, 7, 93",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			PrintDogs(reader, writer)
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
