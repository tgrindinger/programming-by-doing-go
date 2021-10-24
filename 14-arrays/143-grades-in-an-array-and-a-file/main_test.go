package main

import (
	"math/rand"
	"strings"
	"testing"
)

func TestPrintGrades(t *testing.T) {
	cases := []struct {
		desc  string
		input string
		wants []string
	}{
		{ "Marc Antony", "Marc Antony\nettu.txt\n", []string {
				"Grade 1: 75",
				"Grade 2: 15",
				"Grade 3: 54",
				"Grade 4: 7",
				"Grade 5: 16",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			random := rand.New(rand.NewSource(0))
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			PrintGrades(random, reader, writer)
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
