package main

import "testing"

func TestDiscardedCardString(t *testing.T) {
	cases := []struct {
		want   string
		card   *Card
		player int
	}{
		{ "0) Ace of Spades", &Card{Ace, Spades}, 0 },
		{ "1) Two of Hearts", &Card{Two, Hearts}, 1 },
		{ "2) Three of Diamonds", &Card{Three, Diamonds}, 2 },
		{ "3) Four of Clubs", &Card{Four, Clubs}, 3 },
	}
	for _, tc := range cases {
		t.Run(tc.want, func(t *testing.T) {
			card := &DiscardedCard{tc.card, tc.player}
			got := card.String()
			if got != tc.want {
				t.Errorf("got %q want %q", got, tc.want)
			}
		})
	}
}
