package main

import (
	"math/rand"
	"strings"
	"testing"
)

func TestTellStory(t *testing.T) {
	cases := []struct {
		desc  string
		seed  int64
		wants []string
	}{
		{ "default", 0, []string {
				"Choosing randomly from 13 adjectives and 13 nouns (169 combinations).",
				"Your movie title is: electrical license",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			random := rand.New(rand.NewSource(tc.seed))
			writer := &strings.Builder{}
			TellStory(random, writer)
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
