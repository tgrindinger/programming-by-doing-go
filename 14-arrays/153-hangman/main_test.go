package main

import (
	"math/rand"
	"strings"
	"testing"
)

func TestPlayGame(t *testing.T) {
	cases := []struct {
		desc  string
		seed  int64
		input string
		wants []string
	}{
		{ "win", 11, "e\ni\na\nr\ns\nt\nn\no\nl\nv\nh\nquit\n", []string {
				"Word:   l e v i a t h a n ",
				"Misses: rso      ",
				"YOU GOT IT!",
			},
		},
		{ "lose", 11, "a\ns\nd\nf\ng\nh\nj\nk\nl\nq\nw\ne\nr\nquit\n", []string {
				"Word:   l e _ _ a _ h a _ ",
				"Misses: sdfgjkqwr",
				"Oops, you lost!",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			random := rand.New(rand.NewSource(tc.seed))
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			PlayGames(random, reader, writer)
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
