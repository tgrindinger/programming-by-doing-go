package main

import (
	"math/rand"
	"strings"
	"testing"
)

func TestSortArray(t *testing.T) {
	cases := []struct {
		desc  string
		seed  int64
		wants []string
	}{
		{ "default", 0, []string {
				"before: 75 15 54 7 16 97 68 78 89 29 ",
				"after:  7 15 16 29 54 68 75 78 89 97 ",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			random := rand.New(rand.NewSource(tc.seed))
			writer := &strings.Builder{}
			SortArray(random, writer)
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
