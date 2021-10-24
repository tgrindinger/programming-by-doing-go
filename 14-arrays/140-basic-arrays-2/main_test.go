package main

import (
	"math/rand"
	"strings"
	"testing"
)

func TestPrintElements(t *testing.T) {
	cases := []struct {
		desc  string
		wants []string
	}{
		{ "default", []string {
				"Slot 0 contains a 75",
				"Slot 1 contains a 15",
				"Slot 2 contains a 54",
				"Slot 3 contains a 7",
				"Slot 4 contains a 16",
				"Slot 5 contains a 97",
				"Slot 6 contains a 68",
				"Slot 7 contains a 78",
				"Slot 8 contains a 89",
				"Slot 9 contains a 29",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			random := rand.New(rand.NewSource(0))
			writer := &strings.Builder{}
			PrintElements(random, writer)
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
