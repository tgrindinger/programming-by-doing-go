package main

import (
	"strings"
	"testing"
)

func TestPlayGame(t *testing.T) {
	cases := []struct {
		desc  string
		input string
		wants []string
	}{
		{ "trade 2", "2\n1\n0\n5\n-1\n", []string {
				"EXCHANGE POKEMON",
				"\t0. MANKEY",
				"\t1. GEODUDE",
				"\t2. CHARMELEON",
				"\t3. GYARADOS",
				"\t4. BUTTERFREE",
				"\t5. PIKACHU",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			PlayGame(reader, writer)
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
