package main

import (
	"math/rand"
	"strings"
	"testing"
)

func TestPrintCount(t *testing.T) {
	cases := []struct {
		desc  string
		input string
		seed  int64
		wants []string
	}{
		{ "lose", "stay\n", 2, []string {
				"You get a 8 and a 8.\n",
				"Your total is 16.\n",
				"The dealer has a 4 showing and a hidden card.\n",
				"His total is hidden, too.\n",
				"Okay, dealer's turn.\n",
				"His hidden card was a 2.\n",
				"His total was 6.\n",
				"Dealer chooses to hit.\n",
				"He draws a 6.\n",
				"His total is 12.\n",
				"Dealer chooses to hit.\n",
				"He draws a 6.\n",
				"His total is 18.\n",
				"Dealer stays.\n",
				"Dealer total is 18.\n",
				"Your total is 16.\n",
				"YOU LOSE!\n",
			},
		},
		{ "win", "stay\n", 5, []string {
				"You get a 8 and a 8.\n",
				"Your total is 16.\n",
				"The dealer has a 11 showing and a hidden card.\n",
				"His total is hidden, too.\n",
				"Okay, dealer's turn.\n",
				"His hidden card was a 2.\n",
				"His total was 13.\n",
				"Dealer chooses to hit.\n",
				"He draws a 9.\n",
				"His total is 22.\n",
				"DEALER BUSTS!\n",
				"Dealer total is 22.\n",
				"Your total is 16.\n",
				"YOU WIN!\n",
			},
		},
		{ "bust", "hit\n", 5, []string {
				"You get a 8 and a 8.\n",
				"Your total is 16.\n",
				"The dealer has a 11 showing and a hidden card.\n",
				"His total is hidden, too.\n",
				"You drew a 9.\n",
				"Your total is 25.\n",
				"YOU BUST!\n",
				"Okay, dealer's turn.\n",
				"His hidden card was a 2.\n",
				"His total was 13.\n",
				"Dealer stays.\n",
				"Dealer total is 13.\n",
				"Your total is 25.\n",
				"YOU LOSE!\n",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			random := rand.New(rand.NewSource(tc.seed))
			reader := strings.NewReader(tc.input)
			writer := &strings.Builder{}
			PlayBlackjack(random, reader, writer)
			got := writer.String()
			for _, want := range tc.wants {
				if !strings.Contains(got, want) {
					t.Errorf("\ngot  %q\nwant %q", got, want)
				}
			}
		})
	}
}
