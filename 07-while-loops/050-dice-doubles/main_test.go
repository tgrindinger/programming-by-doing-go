package main

import (
	"math/rand"
	"strings"
	"testing"
)

func TestRollDice(t *testing.T) {
	random := rand.New(rand.NewSource(4))
	writer := &strings.Builder{}
	RollDice(random, writer)
	got := writer.String()
	wants := []string {
		"The total is 7!",
		"The total is 8!",
		"The total is 6!",
		"The total is 3!",
		"The total is 10!",
		"The total is 11!",
	}
	for _, want := range wants {
		if !strings.Contains(got, want) {
			t.Errorf("\ngot  %q\nwant %q", got, want)
		}
	}
}