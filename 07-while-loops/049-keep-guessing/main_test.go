package main

import (
	"math/rand"
	"strings"
	"testing"
)

func TestGuessNumber(t *testing.T) {
	random := rand.New(rand.NewSource(0))
	reader := strings.NewReader("1\n2\n3\n4\n5\n6\n7\n8\n9\n10\n")
	writer := &strings.Builder{}
	GuessNumber(random, reader, writer)
	got := writer.String()
	want1 := "That is incorrect. Guess again."
	want2 := "That's right! You're a good guesser."
	if !strings.Contains(got, want1) {
		t.Errorf("\ngot  %q\nwant %q", got, want1)
	}
	if !strings.Contains(got, want2) {
		t.Errorf("\ngot  %q\nwant %q", got, want2)
	}
}
