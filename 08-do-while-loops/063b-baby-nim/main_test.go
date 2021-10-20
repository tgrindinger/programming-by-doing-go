package main

import (
	"strings"
	"testing"
)

func TestPlayNim(t *testing.T) {
	cases := []struct {
		desc  string
		input string
		want  string
	}{
		{ "a2 c3 b1 a1 b1 c2 b1", "A\n2\nC\n3\nB\n1\nA\n1\nB\n1\nC\n2\nB\n1\n", "All piles are empty. Good job!" },
		{ "a3 b3 b3", "A\n3\nB\n3\nC\n3\n", "All piles are empty. Good job!" },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			PlayNim(reader, writer)
			got := writer.String()
			if !strings.Contains(got, tc.want) {
				t.Errorf("\ngot  %q\nwant %q", got, tc.want)
			}
		})
	}
}
