package main

import "testing"

func TestCardCompareTo(t *testing.T) {
	cases := []struct {
		desc  string
		card1 *Card
		card2 *Card
		want  int
	}{
		{ "non-suited less", &Card{Ace, Spades}, &Card{Two, Clubs}, -1 },
		{ "non-suited more", &Card{Two, Spades}, &Card{Ace, Clubs}, -1 },
		{ "suited less", &Card{Ace, Spades}, &Card{Two, Spades}, 1 },
		{ "suited more", &Card{Two, Spades}, &Card{Ace, Spades}, -1 },
	}
	for _, tc := range cases {
		got := tc.card1.compareTo(tc.card2)
		if got != tc.want {
			t.Errorf("got %s want %s", tc.card1, tc.card2)
		}
	}
}

func TestCardScore(t *testing.T) {
	cases := []struct {
		desc string
		card *Card
		want int
	}{
		{ "heart", &Card{Ace, Hearts}, 1 },
		{ "non-heart", &Card{Two, Spades}, 0 },
		{ "queen of spades", &Card{Queen, Spades}, 13 },
	}
	for _, tc := range cases {
		got := tc.card.score()
		if got != tc.want {
			t.Errorf("got %d want %d", got, tc.want)
		}
	}
}
