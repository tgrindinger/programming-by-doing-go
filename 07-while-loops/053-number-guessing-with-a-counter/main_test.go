package main

import (
	"math/rand"
	"strings"
	"testing"
)

func TestGuessNumber(t *testing.T) {
	random := rand.New(rand.NewSource(0))
	reader := strings.NewReader("1\n2\n3\n4\n5\n")
	writer := &strings.Builder{}
	GuessNumber(random, reader, writer)
	got := writer.String()
	wants := []string {
		"That is incorrect.  Guess again.",
		"That's right!  You're a good guesser.",
		"It only took you 5 tries.",
	}
	for _, want := range wants {
		if !strings.Contains(got, want) {
			t.Errorf("\ngot  %q\nwant %q", got, want)
		}
	}
}
