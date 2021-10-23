package main

import (
	"strings"
	"testing"
)

func TestCapitalizeVowels(t *testing.T) {
	cases := []struct {
		desc  string
		input string
		wants []string
	}{
		{ "vowels", "vowels.txt\n", []string {
				"Old McDOnAld hAd A fArm; E-I-E-I-O.  (And dOn't fOrgEt 'U'.)",
				"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
				"^bcd^fgh^jklmn^pqrst^vwxyz",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			CapitalizeVowels(reader, writer)
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
