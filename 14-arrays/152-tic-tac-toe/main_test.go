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
		{ "tie", "0 0\n0 2\n1 1\n2 2\n1 2\n1 0\n2 1\n0 1\n2 0\n", []string {
				"\t0 X O O ",
				"\t1 O X X ",
				"\t2 X X O ",
				"The game is a tie.",
			},
		},
		{ "x wins", "0 0\n0 2\n2 2\n1 1\n2 0\n1 0\n2 1\n", []string {
				"\t0 X   O ",
				"\t1 O O   ",
				"\t2 X X X ",
				"The winner is X!",
			},
		},
		{ "o wins", "0 0\n0 2\n1 0\n2 0\n0 1\n1 1\n", []string {
				"\t0 X X O ",
				"\t1 X O   ",
				"\t2 O     ",
				"The winner is O!",
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
