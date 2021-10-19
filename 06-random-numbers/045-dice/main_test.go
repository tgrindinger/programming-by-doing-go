package main

import (
	"math/rand"
	"strings"
	"testing"
)

func TestRollDice(t *testing.T) {
	random := rand.New(rand.NewSource(1))
	writer := &strings.Builder{}
	RollDice(random, writer)
	got := writer.String()
	want := "Roll #1: 6\nRoll #2: 4\nThe total is 10!"
	if !strings.Contains(got, want) {
		t.Errorf("\ngot  %q\nwant %q", got, want)
	}
}
