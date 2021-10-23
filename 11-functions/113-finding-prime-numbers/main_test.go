package main

import (
	"strings"
	"testing"
)

func TestPrintPrimes(t *testing.T) {
	cases := []struct {
		desc string
		wants []string
	}{
		{ "default", []string {
				"2 <",
				"3 <",
				"4 ",
				"5 <",
				"6 ",
				"7 <",
				"8 ",
				"9 ",
				"10 ",
				"11 <",
				"12 ",
				"13 <",
				"14 ",
				"15 ",
				"16 ",
				"17 <",
				"18 ",
				"19 <",
				"20 ",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			writer := &strings.Builder{}
			PrintPrimes(writer)
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
