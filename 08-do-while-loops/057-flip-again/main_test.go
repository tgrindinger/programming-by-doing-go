package main

import (
	"math/rand"
	"strings"
	"testing"
)

func TestFlip(t *testing.T) {
	random := rand.New(rand.NewSource(0))
	reader := strings.NewReader("y\ny\ny\n")
	writer := &strings.Builder{}
	Flip(random, reader, writer)
	gots := strings.Split(writer.String(), "\n")
	wants := []string {
		"You flip a coin and it is... TAILS",
		"You flip a coin and it is... TAILS",
		"You flip a coin and it is... HEADS",
	}
	for i, want := range wants {
		if !strings.Contains(gots[i], want) {
			t.Errorf("\ngot(%d)  %q\nwant(%d) %q", i, gots[i], i, want)
		}
	}
}
