package main

import (
	"math/rand"
	"strings"
	"testing"
)

func TestFindElement(t *testing.T) {
	cases := []struct {
		desc  string
		input string
		wants []string
	}{
		{ "47", "47\n", []string {
				"Array: 25 15 4 7 16 47 18 28 39 29 ",
				"47 is in the array.",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			random := rand.New(rand.NewSource(0))
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			FindElement(random, reader, writer)
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
