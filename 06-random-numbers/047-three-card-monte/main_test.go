package main

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func TestPlayGame(t *testing.T) {
	cases := []struct {
		seed     int64
		guess    int
		response string
	}{
		{ 0, 2, "Ha! Fast Eddie wins again! The ace was card number 1." },
		{ 0, 1, "You nailed it! Fast Eddie reluctantly hands over your winnings, scowling." },
	}
	for _, tc := range cases {
		t.Run(tc.response, func(t *testing.T) {
			random := rand.New(rand.NewSource(tc.seed))
			reader := strings.NewReader(fmt.Sprintf("%d\n", tc.guess))
			writer := &strings.Builder{}
			PlayGame(random, reader, writer)
			got := writer.String()
			if !strings.Contains(got, tc.response) {
				t.Errorf("\ngot  %q\nwant %q", got, tc.response)
			}
		})
	}
}
