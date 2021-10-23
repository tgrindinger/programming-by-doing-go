package main

import (
	"strings"
	"testing"
)

func TestDecodeMessage(t *testing.T) {
	cases := []struct {
		desc  string
		input string
		wants []string
	}{
		{ "puzzle", "puzzle.txt\n", []string {
				"The_chances_are_good_that_your_program_works.",
			},
		},
		{ "puzzle2", "puzzle2.txt\n", []string {
				"This_is_the_right_message_from_sample2.txt",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			PrintWebLine(reader, writer)
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
