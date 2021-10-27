package main

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func TestMatchDealCards(t *testing.T) {
	random := rand.New(rand.NewSource(0))
	match := &Match{random, []*Player{{[]*Card{}, nil, 0, 0}, {[]*Card{}, nil, 1, 0}, {[]*Card{}, nil, 2, 0}, {[]*Card{}, nil, 3, 0}}, nil, 0, false}
	match.broken = true
	match.dealCards()
	gots := [][]string {
		strings.Split(match.players[0].printHand(), "\n"),
		strings.Split(match.players[1].printHand(), "\n"),
		strings.Split(match.players[2].printHand(), "\n"),
		strings.Split(match.players[3].printHand(), "\n"),
	}
	wants := [][]string {
		{ "1) Two of Spades","2) Seven of Spades","3) Ten of Spades","4) Three of Hearts","5) Five of Hearts",
			"6) Six of Diamonds","7) Seven of Diamonds","8) Jack of Diamonds","9) King of Diamonds","10) Two of Clubs",
			"11) Four of Clubs","12) Nine of Clubs","13) Ace of Clubs",
		},
		{ "1) Three of Spades","2) Nine of Spades","3) Ace of Spades","4) Six of Hearts","5) Seven of Hearts",
			"6) Nine of Hearts","7) Jack of Hearts","8) Queen of Hearts","9) King of Hearts","10) Ace of Hearts",
			"11) Five of Diamonds","12) Ten of Diamonds","13) Ten of Clubs",
		},
		{ "1) Four of Spades","2) Five of Spades","3) Six of Spades","4) Queen of Spades","5) Four of Hearts",
			"6) Ten of Hearts","7) Two of Diamonds","8) Three of Diamonds","9) Eight of Diamonds","10) Nine of Diamonds",
			"11) Queen of Diamonds","12) Three of Clubs","13) Jack of Clubs",
		},
		{ "1) Eight of Spades","2) Jack of Spades","3) King of Spades","4) Two of Hearts","5) Eight of Hearts",
			"6) Four of Diamonds","7) Ace of Diamonds","8) Five of Clubs","9) Six of Clubs","10) Seven of Clubs",
			"11) Eight of Clubs","12) Queen of Clubs","13) King of Clubs",
		},
	}
	for i, got := range gots {
		assertGot(t, fmt.Sprintf("player %d", i), got, wants[i])
	}
	if match.broken == true {
		t.Errorf("broken got %t want %t", match.broken, true)
	}
}

func TestMatchDoFirstTurn(t *testing.T) {
	startCard := &Card{Two, Clubs}
	c1 := &Card{Three, Hearts}
	c2 := &Card{Four, Diamonds}
	c3 := &Card{Five, Spades}
	cases := []struct {
		desc    string
		cards   []*Card
		current int
		hands   []int
	}{
		{ "first", []*Card{startCard, c1, c2, c3}, 1, []int{0, 1, 1, 1} },
		{ "second", []*Card{c1, startCard, c2, c3}, 2, []int{1, 0, 1, 1} },
		{ "third", []*Card{c1, c2, startCard, c3}, 3, []int{1, 1, 0, 1} },
		{ "fourth", []*Card{c1, c2, c3, startCard}, 0, []int{1, 1, 1, 0} },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			players := []*Player{{[]*Card{tc.cards[0]}, nil, 0, 0},{[]*Card{tc.cards[1]}, nil, 1, 0},{[]*Card{tc.cards[2]}, nil, 2, 0},{[]*Card{tc.cards[3]}, nil, 3, 0},}
			match := &Match{nil, players, nil, 0, false}
			match.doFirstTurn()
			if match.current != tc.current {
				t.Errorf("current: got %d want %d", match.current, tc.current)
			}
			for i, count := range tc.hands {
				if len(match.players[i].hand) != count {
					t.Errorf("hand: got %v want %d", match.players[i].hand, count)
				}
			}
		})
	}
}

func TestMatchMatchOver(t *testing.T) {
	card := &Card{Five, Spades}
	empty := &Player{[]*Card{}, nil, 0, 0}
	hasCard := &Player{[]*Card{card}, nil, 0, 0}
	cases := []struct {
		desc    string
		players []*Player
		want    bool
	}{
		{ "empty", []*Player{empty, empty, empty, empty}, true },
		{ "cards left", []*Player{empty, empty, hasCard, empty}, false },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			match := &Match{nil, tc.players, nil, 0, false}
			got := match.matchOver()
			if got != tc.want {
				t.Errorf("current: got %t want %t", got, tc.want)
			}
		})
	}
}

func TestMatchPlayCard(t *testing.T) {
	card0 := &Card{Ace, Spades}
	card1 := &Card{Two, Hearts}
	cases := []struct {
		desc         string
		card         int
		current      int
		discards     []*DiscardedCard
	  players      []*Player
		wantDiscards []*DiscardedCard
		wantCurrent  int
	}{
		{ "empty trick", 0, 0, []*DiscardedCard{},            []*Player{{[]*Card{card0}, nil, 0, 0}, {[]*Card{card1}, nil, 1, 0}}, []*DiscardedCard{{card0, 0}}, 1 },
		{ "single trick", 0, 1, []*DiscardedCard{{card0, 0}}, []*Player{{[]*Card{card0}, nil, 0, 0}, {[]*Card{card1, card1}, nil, 1, 0}}, []*DiscardedCard{{card0, 0}, {card1, 1}}, 0 },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			match := &Match{nil, tc.players, tc.discards, tc.current, false}
			match.playCard(tc.card)
			if match.current != tc.wantCurrent {
				t.Errorf("current got %d want %d", match.current, tc.wantCurrent)
			}
			if len(match.trick) != len(tc.wantDiscards) {
				t.Errorf("discards got %d want %d", len(match.trick), len(tc.wantDiscards))
			} else {
				for i, c := range match.trick {
					if c.card.face != tc.wantDiscards[i].card.face || c.card.suit != tc.wantDiscards[i].card.suit {
						t.Errorf("discards got %s want %s", match.trick, tc.wantDiscards)
					}
				}
			}
		})
	}
}

func TestMatchTrickOver(t *testing.T) {
	c0 := &DiscardedCard{&Card{Two, Clubs}, 0}
	c1 := &DiscardedCard{&Card{Three, Hearts}, 1}
	c2 := &DiscardedCard{&Card{Queen, Spades}, 2}
	c3 := &DiscardedCard{&Card{Five, Diamonds}, 3}
	cases := []struct {
		desc     string
		discards []*DiscardedCard
	  players  []*Player
		want     bool
	}{
		{ "empty", []*DiscardedCard{},                    []*Player{{nil, nil, 0, 0}, {nil, nil, 1, 0}, {nil, nil, 2, 0}, {nil, nil, 3, 0}}, false },
		{ "one card", []*DiscardedCard{c0},               []*Player{{nil, nil, 0, 0}, {nil, nil, 1, 0}, {nil, nil, 2, 0}, {nil, nil, 3, 0}}, false },
		{ "two cards", []*DiscardedCard{c0, c1},          []*Player{{nil, nil, 0, 0}, {nil, nil, 1, 0}, {nil, nil, 2, 0}, {nil, nil, 3, 0}}, false },
		{ "three cards", []*DiscardedCard{c0, c1, c2},    []*Player{{nil, nil, 0, 0}, {nil, nil, 1, 0}, {nil, nil, 2, 0}, {nil, nil, 3, 0}}, false },
		{ "four cards", []*DiscardedCard{c0, c1, c2, c3}, []*Player{{nil, nil, 0, 0}, {nil, nil, 1, 0}, {nil, nil, 2, 0}, {nil, nil, 3, 0}}, true },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			match := &Match{nil, tc.players, tc.discards, 0, false}
			got := match.trickOver()
			if got != tc.want {
				t.Errorf("current got %t want %t", got, tc.want)
			}
		})
	}
}

func TestMatchTrickWinner(t *testing.T) {
	c0 := &DiscardedCard{&Card{Two, Clubs}, 0}
	c1 := &DiscardedCard{&Card{Three, Hearts}, 1}
	c2 := &DiscardedCard{&Card{Queen, Spades}, 2}
	c3 := &DiscardedCard{&Card{Five, Diamonds}, 3}
	cases := []struct {
		desc     string
		discards []*DiscardedCard
	  players  []*Player
		want     int
	}{
		{ "player 0", []*DiscardedCard{c0, c1, c2, c3}, []*Player{{nil, nil, 0, 0}, {nil, nil, 1, 0}, {nil, nil, 2, 0}, {nil, nil, 3, 0}}, 0 },
		{ "player 1", []*DiscardedCard{c1, c0, c2, c3}, []*Player{{nil, nil, 0, 0}, {nil, nil, 1, 0}, {nil, nil, 2, 0}, {nil, nil, 3, 0}}, 1 },
		{ "player 2", []*DiscardedCard{c2, c0, c1, c3}, []*Player{{nil, nil, 0, 0}, {nil, nil, 1, 0}, {nil, nil, 2, 0}, {nil, nil, 3, 0}}, 2 },
		{ "player 3", []*DiscardedCard{c3, c0, c1, c2}, []*Player{{nil, nil, 0, 0}, {nil, nil, 1, 0}, {nil, nil, 2, 0}, {nil, nil, 3, 0}}, 3 },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			match := &Match{nil, tc.players, tc.discards, 0, false}
			got := match.trickWinner()
			if got != tc.want {
				t.Errorf("current got %d want %d", got, tc.want)
			}
		})
	}
}

func TestMatchCollectTrick(t *testing.T) {
	card0 := &Card{Ace, Spades}
	card1 := &Card{Two, Hearts}
	cases := []struct {
		desc        string
		current     int
		wantCurrent int
		discards    []*DiscardedCard
	  players     []*Player
		wantPlayers []*Player
		scores      []int
	}{
		{ "0 wins", 0, 0, []*DiscardedCard{{card0, 0}, {card1, 1}}, []*Player{{nil, []*Card{}, 0, 0}, {nil, []*Card{}, 1, 0}}, []*Player{{nil, []*Card{card0, card1}, 0, 1}, {nil, []*Card{}, 1, 0}}, []int{1, 0} },
		{ "1 wins", 0, 1, []*DiscardedCard{{card0, 1}, {card1, 0}}, []*Player{{nil, []*Card{}, 0, 0}, {nil, []*Card{}, 1, 0}}, []*Player{{nil, []*Card{}, 0, 0}, {nil, []*Card{card0, card1}, 1, 1}}, []int{0, 1} },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			match := &Match{nil, tc.players, tc.discards, tc.current, false}
			match.collectTrick()
			if match.current != tc.wantCurrent {
				t.Errorf("current got %d want %d", match.current, tc.wantCurrent)
			}
			if len(match.trick) != 0 {
				t.Errorf("discards got %d want %d", len(match.trick), 0)
			}
			for i, p := range match.players {
				if p.currentScore() != tc.scores[i] {
					t.Errorf("scores got %v want %v", p.currentScore(), tc.scores)
				}
			}
			for i, p := range match.players {
				for j, c := range p.tricks {
					if c.face != tc.wantPlayers[i].tricks[j].face || c.suit != tc.wantPlayers[i].tricks[j].suit {
						t.Errorf("tricks got %s want %s", match.players[i].tricks, tc.wantPlayers[i].tricks)
					}
				}
			}
		})
	}
}

func TestValidPlay(t *testing.T) {
	card0 := &Card{Ace, Spades}
	card1 := &Card{Two, Hearts}
	cases := []struct {
		desc   string
		trick  []*DiscardedCard
		cards  []*Card
		broken bool
		want   bool
	}{
		{ "first unbroken non-heart", []*DiscardedCard{}, []*Card{card0}, false, true },
		{ "first unbroken heart no alternatives", []*DiscardedCard{}, []*Card{card1}, false, true },
		{ "first unbroken heart with alternatives", []*DiscardedCard{}, []*Card{card1, card0}, false, false },
		{ "non-first unbroken non-heart", []*DiscardedCard{{card0, 1}}, []*Card{card0}, false, true },
		{ "non-first unbroken heart no alternatives", []*DiscardedCard{{card0, 1}}, []*Card{card1}, false, true },
		{ "non-first unbroken heart with alternatives", []*DiscardedCard{{card0, 1}}, []*Card{card1, card0}, false, false },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			match := &Match{nil, []*Player{{tc.cards, nil, 0, 0}}, tc.trick, 0, tc.broken}
			got := match.validPlay(0)
			if got != tc.want {
				t.Errorf("got %t want %t", got, tc.want)
			}
		})
	}
}
