package main

import (
	"strings"
	"testing"
)

func TestGetPin(t *testing.T) {
	cases := []struct {
		input string
		wants []string
	}{
		{ "10101\n23232\n99999\n", []string { "INCORRECT PIN. TRY AGAIN.", "YOU HAVE RUN OUT OF TRIES. ACCOUNT LOCKED." } },
		{ "10101\n12345\n", []string { "INCORRECT PIN. TRY AGAIN.", "PIN ACCEPTED. YOU NOW HAVE ACCESS TO YOUR ACCOUNT." } },
	}
	for _, tc := range cases {
		reader := strings.NewReader(tc.input)
		writer := &strings.Builder{}
		GetPin(reader, writer)
		got := writer.String()
		for _, want := range tc.wants {
			if !strings.Contains(got, want) {
				t.Errorf("\ngot  %q\nwant %q", got, want)
			}
		}
	}
}
