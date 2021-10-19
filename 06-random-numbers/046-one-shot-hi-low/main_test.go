package main

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func TestGuessNumber(t *testing.T) {
	cases := []struct {
		seed     int64
		guess    int
		response string
	}{
		{ 0, 13, "Sorry, you are too low.  I was thinking of 75." },
		{ 0, 80, "Sorry, you are too high.  I was thinking of 75." },
		{ 0, 75, "You guessed it!  What are the odds?!?" },
	}
	for _, tc := range cases {
		t.Run(tc.response, func(t *testing.T) {
			random := rand.New(rand.NewSource(tc.seed))
			reader := strings.NewReader(fmt.Sprintf("%d\n", tc.guess))
			writer := &strings.Builder{}
			GuessNumber(random, reader, writer)
			got := writer.String()
			want := tc.response
			if !strings.Contains(got, want) {
				t.Errorf("\ngot  %q\nwant %q", got, want)
			}
		})
	}
}
