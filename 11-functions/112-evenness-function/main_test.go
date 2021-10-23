package main

import (
	"strings"
	"testing"
)

func TestPrintEvenness(t *testing.T) {
	cases := []struct {
		desc string
		wants []string
	}{
		{ "default", []string {
				"1 ",
				"2 <",
				"3 =",
				"6 <=",
				"20 <",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			writer := &strings.Builder{}
			PrintEvenness(writer)
			gots := strings.Split(writer.String(), "\n")
			for _, want := range tc.wants {
				
				if !contains(gots, want) {
					t.Errorf("\ngot  %q\nwant %q", gots, want)
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
