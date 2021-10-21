package main

import (
	"math/rand"
	"strings"
	"testing"
)

func TestPrintCount(t *testing.T) {
	cases := []struct {
		desc string
		seed int64
		wants []string
	}{
		{ "win", 2, []string {
				"You drew 7 and 7.",
				"Your total is 14.",
				"The dealer has 3 and 1.",
				"Dealer's total is 4.",
				"YOU WIN!",
			},
		},
		{ "lose", 0, []string {
				"You drew 5 and 5.",
				"Your total is 10.",
				"The dealer has 4 and 7.",
				"Dealer's total is 11.",
				"YOU LOSE!",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			random := rand.New(rand.NewSource(tc.seed))
			writer := &strings.Builder{}
			PlayBlackjack(random, writer)
			got := writer.String()
			for _, want := range tc.wants {
				if !strings.Contains(got, want) {
					t.Errorf("\ngot  %q\nwant %q", got, want)
				}
			}
		})
	}
}
