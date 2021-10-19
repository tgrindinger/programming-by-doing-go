package main

import (
	"math/rand"
	"strings"
	"testing"
)

func TestPlayHiLo(t *testing.T) {
	cases := []struct {
		input string
		desc  string
		wants []string
	}{
		{ "50\n80\n75\n", "multiple tries", []string { "Sorry, you are too low.", "Sorry, you are too high.", "You guessed it!  What are the odds?!?" } },
		{ "50\n80\n71\n71\n71\n71\n71\n", "failure", []string { "Sorry, you are too low.", "Sorry, you are too high.", "Sorry, you didn't guess it in 7 tries.  You lose." } },
		{ "75\n", "success", []string { "You guessed it!  What are the odds?!?" } },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			random := rand.New(rand.NewSource(0))
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			PlayHiLo(random, reader, writer)
			got := writer.String()
			for _, want := range tc.wants {
				if !strings.Contains(got, want) {
					t.Errorf("\ngot  %q\nwant %q", got, want)
				}
			}
		})
	}
}
