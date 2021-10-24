package main

import (
	"math/rand"
	"strings"
	"testing"
)

func TestFillArray(t *testing.T) {
	cases := []struct {
		desc  string
		seed  int64
		wants []string
	}{
		{ "default", 0, []string {
				"The first array is filled with the following values:",
				"\talpha bravo charlie delta echo ",
				"The second array is filled with the following values:",
				"\t11 23 37 41 53 ",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			random := rand.New(rand.NewSource(tc.seed))
			writer := &strings.Builder{}
			FillArray(random, writer)
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
