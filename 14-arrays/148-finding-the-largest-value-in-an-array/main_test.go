package main

import (
	"math/rand"
	"strings"
	"testing"
)

func TestFindElement(t *testing.T) {
	cases := []struct {
		desc  string
		seed  int64
		wants []string
	}{
		{ "47", 0, []string {
				"Array: 75 15 54 7 16 97 68 78 89 29 ",
				"The largest value is 97.",
			},
		},
		{ "48", 1, []string {
				"Array: 82 88 48 60 82 19 26 41 57 1 ",
				"The largest value is 88.",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			random := rand.New(rand.NewSource(tc.seed))
			writer := &strings.Builder{}
			FindElement(random, writer)
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
