package main

import (
	"strings"
	"testing"
)

func TestAskQuestions(t *testing.T) {
	cases := []struct {
		input  string
		result string
	}{
		{ "Dukes\n19\n", "You can vote but you can't rent a car, Dukes." },
		{ "James\n15\n", "You can't drive, James." },
		{ "Tim\n16\n", "You can drive but not vote, Tim." },
		{ "Jen\n26\n", "You can do pretty much anything, Jen." },
	}
	for _, tc := range cases {
		t.Run(tc.result, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			AskQuestions(reader, writer)
			lines := strings.Split(writer.String(), "\n")
			got := lines[len(lines)-2]
			want := tc.result
			if got != want {
				t.Errorf("\ngot  %q\nwant %q", got, want)
			}
		})
	}
}
