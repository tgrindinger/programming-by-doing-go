package main

import (
	"strings"
	"testing"
)

func TestPrintArrayList(t *testing.T) {
	cases := []struct {
		desc  string
		wants []string
	}{
		{ "default", []string {
				"Slot 0 contains a -113",
				"Slot 1 contains a -113",
				"Slot 2 contains a -113",
				"Slot 3 contains a -113",
				"Slot 4 contains a -113",
				"Slot 5 contains a -113",
				"Slot 6 contains a -113",
				"Slot 7 contains a -113",
				"Slot 8 contains a -113",
				"Slot 9 contains a -113",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			writer := &strings.Builder{}
			PrintArrayList(writer)
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
