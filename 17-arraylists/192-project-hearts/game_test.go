package main

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func TestGameDealCards(t *testing.T) {
	random := rand.New(rand.NewSource(0))
	game := NewGame(random)
	game.broken = true
	game.dealCards()
	gots := [][]string {
		strings.Split(game.players[0].printHand(), "\n"),
		strings.Split(game.players[1].printHand(), "\n"),
		strings.Split(game.players[2].printHand(), "\n"),
		strings.Split(game.players[3].printHand(), "\n"),
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
	if game.broken == true {
		t.Errorf("broken got %t want %t", game.broken, true)
	}
}

func TestGameDoFirstTurn(t *testing.T) {
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
			players := []*Player{{[]*Card{tc.cards[0]}, nil},{[]*Card{tc.cards[1]}, nil},{[]*Card{tc.cards[2]}, nil},{[]*Card{tc.cards[3]}, nil},}
			game := &Game{nil, players, nil, 0, nil, false}
			game.doFirstTurn()
			if game.current != tc.current {
				t.Errorf("current: got %d want %d", game.current, tc.current)
			}
			for i, count := range tc.hands {
				if len(game.players[i].hand) != count {
					t.Errorf("hand: got %v want %d", game.players[i].hand, count)
				}
			}
		})
	}
}

func TestGameMatchOver(t *testing.T) {
	card := &Card{Five, Spades}
	empty := &Player{[]*Card{}, nil}
	hasCard := &Player{[]*Card{card}, nil}
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
			game := &Game{nil, tc.players, nil, 0, nil, false}
			got := game.matchOver()
			if got != tc.want {
				t.Errorf("current: got %t want %t", got, tc.want)
			}
		})
	}
}

func TestGameUpdateScores(t *testing.T) {
	c0 := &Card{Two, Clubs}
	c1 := &Card{Three, Hearts}
	c2 := &Card{Queen, Spades}
	c3 := &Card{Five, Diamonds}
	cases := []struct {
		desc   string
		tricks [][]*Card
		want   []int
	}{
		{ "one card each", [][]*Card{{c0}, {c1}, {c2}, {c3}}, []int{0, 1, 13, 0} },
		{ "two cards each", [][]*Card{{c0, c1}, {c1, c1}, {c3, c2}, {c3, c2, c1}}, []int{1, 2, 13, 14} },
		{ "shoot the moon", [][]*Card{{}, {}, {}, {{Queen, Spades}, {Two, Hearts}, {Three, Hearts}, {Four, Hearts}, {Five, Hearts}, {Six, Hearts}, {Seven, Hearts}, {Eight, Hearts}, {Nine, Hearts}, {Ten, Hearts}, {Jack, Hearts}, {Queen, Hearts}, {King, Hearts}, {Ace, Hearts}}}, []int{26, 26, 26, 0}},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			players := []*Player{{nil, tc.tricks[0]},{nil, tc.tricks[1]},{nil, tc.tricks[2]},{nil, tc.tricks[3]}}
			game := &Game{nil, players, nil, 0, []int{0, 0, 0, 0}, false}
			game.updateScores()
			for i, got := range game.scores {
				if got != tc.want[i] {
					t.Errorf("hand: got %v want %v", game.scores, tc.want)
				}
			}
		})
	}
}

func TestGameGameOver(t *testing.T) {
	cases := []struct {
		desc   string
		scores []int
		want   bool
	}{
		{ "not over", []int{5, 6, 7, 8}, false },
		{ "game over", []int{13, 26, 39, 52}, true },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			game := &Game{nil, nil, nil, 0, tc.scores, false}
			got := game.gameOver()
			if got != tc.want {
				t.Errorf("hand: got %t want %t", got, tc.want)
			}
		})
	}
}

func TestGamePlayCard(t *testing.T) {
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
		{ "empty trick", 0, 0, []*DiscardedCard{}, []*Player{{[]*Card{card0}, nil}, {[]*Card{card1}, nil}}, []*DiscardedCard{{card0, 0}}, 1 },
		{ "single trick", 0, 1, []*DiscardedCard{{card0, 0}}, []*Player{{[]*Card{card0}, nil}, {[]*Card{card1, card1}, nil}}, []*DiscardedCard{{card0, 0}, {card1, 1}}, 0 },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			game := &Game{nil, tc.players, tc.discards, tc.current, nil, false}
			game.playCard(tc.card)
			if game.current != tc.wantCurrent {
				t.Errorf("current got %d want %d", game.current, tc.wantCurrent)
			}
			if len(game.trick) != len(tc.wantDiscards) {
				t.Errorf("discards got %d want %d", len(game.trick), len(tc.wantDiscards))
			} else {
				for i, c := range game.trick {
					if c.card.face != tc.wantDiscards[i].card.face || c.card.suit != tc.wantDiscards[i].card.suit {
						t.Errorf("discards got %s want %s", game.trick, tc.wantDiscards)
					}
				}
			} 
		})
	}
}

func TestGameTrickOver(t *testing.T) {
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
		{ "empty", []*DiscardedCard{}, []*Player{{nil, nil}, {nil, nil}, {nil, nil}, {nil, nil}}, false },
		{ "one card", []*DiscardedCard{c0}, []*Player{{nil, nil}, {nil, nil}, {nil, nil}, {nil, nil}}, false },
		{ "two cards", []*DiscardedCard{c0, c1}, []*Player{{nil, nil}, {nil, nil}, {nil, nil}, {nil, nil}}, false },
		{ "three cards", []*DiscardedCard{c0, c1, c2}, []*Player{{nil, nil}, {nil, nil}, {nil, nil}, {nil, nil}}, false },
		{ "four cards", []*DiscardedCard{c0, c1, c2, c3}, []*Player{{nil, nil}, {nil, nil}, {nil, nil}, {nil, nil}}, true },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			game := &Game{nil, tc.players, tc.discards, 0, nil, false}
			got := game.trickOver()
			if got != tc.want {
				t.Errorf("current got %t want %t", got, tc.want)
			}
		})
	}
}

func TestGameTrickWinner(t *testing.T) {
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
		{ "player 0", []*DiscardedCard{c0, c1, c2, c3}, []*Player{{nil, nil}, {nil, nil}, {nil, nil}, {nil, nil}}, 0 },
		{ "player 1", []*DiscardedCard{c1, c0, c2, c3}, []*Player{{nil, nil}, {nil, nil}, {nil, nil}, {nil, nil}}, 1 },
		{ "player 2", []*DiscardedCard{c2, c0, c1, c3}, []*Player{{nil, nil}, {nil, nil}, {nil, nil}, {nil, nil}}, 2 },
		{ "player 3", []*DiscardedCard{c3, c0, c1, c2}, []*Player{{nil, nil}, {nil, nil}, {nil, nil}, {nil, nil}}, 3 },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			game := &Game{nil, tc.players, tc.discards, 0, nil, false}
			got := game.trickWinner()
			if got != tc.want {
				t.Errorf("current got %d want %d", got, tc.want)
			}
		})
	}
}

func TestGameCollectTrick(t *testing.T) {
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
		{ "0 wins", 0, 0, []*DiscardedCard{{card0, 0}, {card1, 1}}, []*Player{{nil, []*Card{}}, {nil, []*Card{}}}, []*Player{{nil, []*Card{card0, card1}}, {nil, []*Card{}}}, []int{1, 0} },
		{ "1 wins", 0, 1, []*DiscardedCard{{card0, 1}, {card1, 0}}, []*Player{{nil, []*Card{}}, {nil, []*Card{}}}, []*Player{{nil, []*Card{}}, {nil, []*Card{card0, card1}}}, []int{0, 1} },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			game := &Game{nil, tc.players, tc.discards, tc.current, []int{0, 0}, false}
			game.collectTrick()
			if game.current != tc.wantCurrent {
				t.Errorf("current got %d want %d", game.current, tc.wantCurrent)
			}
			if len(game.trick) != 0 {
				t.Errorf("discards got %d want %d", len(game.trick), 0)
			}
			for i, p := range game.players {
				if p.score() != tc.scores[i] {
					t.Errorf("scores got %v want %v", p.score(), tc.scores)
				}
			}
			for i, p := range game.players {
				for j, c := range p.tricks {
					if c.face != tc.wantPlayers[i].tricks[j].face || c.suit != tc.wantPlayers[i].tricks[j].suit {
						t.Errorf("tricks got %s want %s", game.players[i].tricks, tc.wantPlayers[i].tricks)
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
			game := &Game{nil, []*Player{{tc.cards, nil}}, tc.trick, 0, nil, tc.broken}
			got := game.validPlay(0)
			if got != tc.want {
				t.Errorf("got %t want %t", got, tc.want)
			}
		})
	}
}
