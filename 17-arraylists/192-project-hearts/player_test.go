package main

import (
	"strings"
	"testing"
)

func TestPlayerPrintHand(t *testing.T) {
	card1 := &Card{Ace, Spades}
	card2 := &Card{Two, Hearts}
	card3 := &Card{Three, Diamonds}
	card4 := &Card{Four, Clubs}
	cases := []struct {
		desc  string
		hand  []*Card
		wants []string
	}{
		{ "empty", []*Card{}, []string {
				"",
			},
		},
		{ "one card", []*Card{card1}, []string {
				"1) Ace of Spades",
			},
		},
		{ "two cards", []*Card{card1, card2}, []string {
				"1) Ace of Spades",
				"2) Two of Hearts",
			},
		},
		{ "three cards", []*Card{card1, card2, card3}, []string {
				"1) Ace of Spades",
				"2) Two of Hearts",
				"3) Three of Diamonds",
			},
		},
		{ "four cards", []*Card{card1, card2, card3, card4}, []string {
				"1) Ace of Spades",
				"2) Two of Hearts",
				"3) Three of Diamonds",
				"4) Four of Clubs",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			player := &Player{tc.hand, nil, 0, 0}
			gots := strings.Split(player.printHand(), "\n")
			assertGot(t, "", gots, tc.wants)
		})
	}
}


func TestPlayerPlayCard(t *testing.T) {
	card1 := &Card{Ace, Spades}
	card2 := &Card{Two, Hearts}
	card3 := &Card{Three, Diamonds}
	card4 := &Card{Four, Clubs}
	cases := []struct {
		desc   string
		before []*Card
		choice int
		played *Card
		after  []*Card
	}{
		{ "one card", []*Card{card1}, 0, card1, []*Card{} },
		{ "two cards", []*Card{card1, card2}, 1, card2, []*Card{card1} },
		{ "three cards", []*Card{card1, card2, card3}, 1, card2, []*Card{card1, card3} },
		{ "four cards", []*Card{card1, card2, card3, card4}, 0, card1, []*Card{card2, card3, card4} },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			player := &Player{tc.before, nil, 0, 0}
			card := player.playCard(tc.choice)
			if card != tc.played {
				t.Errorf("got %v want %v", card, tc.played)
			}
			if len(player.hand) != len(tc.after) {
				t.Errorf("got %v want %v", player.hand, tc.after)
			}
			for i, c := range player.hand {
				if c != tc.after[i] {
					t.Errorf("got %v want %v", player.hand, tc.after)
				}
			}
		})
	}
}

func TestPlayerReceiveCard(t *testing.T) {
	card1 := &Card{Ace, Spades}
	card2 := &Card{Two, Hearts}
	card3 := &Card{Three, Diamonds}
	card4 := &Card{Four, Clubs}
	cases := []struct {
		desc   string
		before []*Card
		card   *Card
		after  []*Card
	}{
		{ "empty", []*Card{}, card1, []*Card{card1} },
		{ "one card", []*Card{card1}, card2, []*Card{card1, card2} },
		{ "two cards", []*Card{card1, card2}, card3, []*Card{card1, card2, card3} },
		{ "three cards", []*Card{card1, card2, card3}, card4, []*Card{card1, card2, card3, card4} },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			player := &Player{tc.before, nil, 0, 0}
			player.receiveCard(tc.card)
			if len(player.hand) != len(tc.after) {
				t.Errorf("got %v want %v", player.hand, tc.after)
			}
			for i, c := range player.hand {
				if c != tc.after[i] {
					t.Errorf("got %v want %v", player.hand, tc.after)
				}
			}
		})
	}
}

func TestPlayerReceiveTrick(t *testing.T) {
	card1 := &Card{Ace, Spades}
	card2 := &Card{Two, Hearts}
	card3 := &Card{Three, Diamonds}
	card4 := &Card{Four, Clubs}
	cases := []struct {
		desc   string
		before []*Card
		trick  []*Card
		after  []*Card
	}{
		{ "empty", []*Card{}, []*Card{card1}, []*Card{card1} },
		{ "one card", []*Card{card1}, []*Card{card2}, []*Card{card1, card2} },
		{ "two cards", []*Card{card1}, []*Card{card2, card3}, []*Card{card1, card2, card3} },
		{ "three cards", []*Card{card1}, []*Card{card2, card3, card4}, []*Card{card1, card2, card3, card4} },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			player := &Player{nil, tc.before, 0, 0}
			player.receiveTrick(tc.trick)
			if len(player.tricks) != len(tc.after) {
				t.Errorf("got %v want %v", player.tricks, tc.after)
			}
			for i, c := range player.tricks {
				if c != tc.after[i] {
					t.Errorf("got %v want %v", player.tricks, tc.after)
				}
			}
		})
	}
}

func TestPlayerScore(t *testing.T) {
	cases := []struct {
		desc   string
		tricks []*Card
		want   int
	}{
		{ "0 - single club", []*Card{{Ace, Clubs}}, 0 },
		{ "0 - single spade", []*Card{{Ace, Spades}}, 0 },
		{ "0 - single diamond", []*Card{{Ace, Diamonds}}, 0 },
		{ "1 - single heart", []*Card{{Two, Hearts}}, 1 },
		{ "13 - single queen of spades", []*Card{{Queen, Spades}}, 13 },
		{ "0 - multiple", []*Card{{Ace, Clubs}, {Two, Diamonds}, {Three, Spades}}, 0 },
		{ "2 - multiple", []*Card{{Ace, Clubs}, {Two, Hearts}, {Three, Hearts}}, 2 },
		{ "14 - multiple", []*Card{{Ace, Clubs}, {Two, Hearts}, {Queen, Spades}}, 14 },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			player := &Player{nil, tc.tricks, 0, 0}
			got := player.currentScore()
			if got != tc.want {
				t.Errorf("got %d want %d", got, tc.want)
			}
		})
	}
}

func TestPlayerHasNonHeartCards(t *testing.T) {
	cases := []struct {
		desc string
		hand []*Card
		want bool
	}{
		{ "empty", []*Card{}, false },
		{ "single heart", []*Card{{Two, Hearts}}, false },
		{ "single non-heart", []*Card{{Two, Spades}}, true },
		{ "multi hearts", []*Card{{Two, Hearts}, {Three, Hearts}}, false },
		{ "multi non-hearts", []*Card{{Two, Spades}, {Three, Clubs}}, true },
		{ "multi mixed", []*Card{{Two, Hearts}, {Three, Clubs}}, true },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			player := &Player{tc.hand, nil, 0, 0}
			got := player.hasNonHeartCards()
			if got != tc.want {
				t.Errorf("got %t want %t", got, tc.want)
			}
		})
	}
}

func TestPlayerHasMatchingSuit(t *testing.T) {
	cases := []struct {
		desc string
		suit Suit
		hand []*Card
		want bool
	}{
		{ "empty", Spades, []*Card{}, false },
		{ "single match", Spades, []*Card{{Two, Spades}}, true },
		{ "single non-match", Spades, []*Card{{Two, Hearts}}, false },
		{ "multi match", Spades, []*Card{{Two, Hearts}, {Three, Spades}}, true },
		{ "multi non-match", Spades, []*Card{{Two, Hearts}, {Three, Clubs}}, false },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			player := &Player{tc.hand, nil, 0, 0}
			got := player.hasMatchingSuit(tc.suit)
			if got != tc.want {
				t.Errorf("got %t want %t", got, tc.want)
			}
		})
	}
}
