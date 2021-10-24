package main

import (
	"strings"
	"testing"
)

func TestSortArray(t *testing.T) {
	cases := []struct {
		desc  string
		wants []string
	}{
		{ "default", []string {
				"before: 45 87 39 32 93 86 12 44 75 50 ",
				"after:  12 32 39 44 45 50 75 86 87 93 ",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			writer := &strings.Builder{}
			SortArray(writer)
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
