package main

import (
	"strings"
	"testing"
)

func TestPrintCounter(t *testing.T) {
	cases := []struct {
		desc     string
		wants    []string
	}{
		{ "default", []string {
				"(0,0) (0,1) (0,2) (0,3) (0,4) (0,5) ",
				"(1,0) (1,1) (1,2) (1,3) (1,4) (1,5) ",
				"(2,0) (2,1) (2,2) (2,3) (2,4) (2,5) ",
				"(3,0) (3,1) (3,2) (3,3) (3,4) (3,5) ",
				"(4,0) (4,1) (4,2) (4,3) (4,4) (4,5) ",
				"(5,0) (5,1) (5,2) (5,3) (5,4) (5,5) ",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			writer := &strings.Builder{}
			PrintCounter(writer)
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
