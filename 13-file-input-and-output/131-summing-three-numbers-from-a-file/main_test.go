package main

import (
	"strings"
	"testing"
)

func TestPrintSum(t *testing.T) {
	cases := []struct {
		desc  string
		input string
		wants []string
	}{
		{ "default", "3\n1\n4\n", []string {
				"3 + 1 + 4 = 8",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			PrintSum(reader, writer)
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
