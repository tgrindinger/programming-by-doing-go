package main

import (
	"strings"
	"testing"
)

func TestPlayNim(t *testing.T) {
	cases := []struct {
		desc    string
		input   string
		want    string
	}{
		{ "example", "Alice\nBob\nA\n2\nC\n3\nB\n1\nB\n1\nA\n1\nB\n1\nC\n2\n", "Bob, you must take the last remaining counter, so you lose. Alice wins!" },
		{ "quick", "Alice\nBob\nA\n3\nB\n4\nC\n5\n", "Bob, there are no counters left, so you WIN!" },
		{ "empty pile", "Alice\nBob\nA\n3\nA\nB\n3\nC\n5\nB\n1\n", "Nice try, Bob. That pile is empty. Choose again" },
		{ "not enough", "Alice\nBob\nA\n4\nA\n3\nB\n3\nC\n5\nB\n1\n", "Pile A doesn't have that many. Try again" },
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
