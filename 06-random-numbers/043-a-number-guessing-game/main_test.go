package main

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func TestGuessNumber(t *testing.T) {
	cases := []struct {
		seed   int64
		guess  int
		result string
	}{
		{0, 3, "Sorry, but I was really thinking of 5."},
		{0, 5, "That's right!  My secret number was 5!"},
		{1, 3, "Sorry, but I was really thinking of 2."},
		{1, 2, "That's right!  My secret number was 2!"},
	}
	for _, tc := range cases {
		t.Run(tc.result, func(t *testing.T) {
			random := rand.New(rand.NewSource(tc.seed))
			reader := strings.NewReader(fmt.Sprintf("%d\n", tc.guess))
			writer := &strings.Builder{}
			GuessNumber(random, reader, writer)
			got := writer.String()
			if !strings.Contains(got, tc.result) {
				t.Errorf("\ngot  %q\nwant %q", got, tc.result)
			}
		})
	}
}
