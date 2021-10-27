package main

import (
	"testing"
)

func TestGameWinner(t *testing.T) {
	cases := []struct {
		desc   string
		scores []int
		want   int
	}{
		{ "player 0", []int{0, 1, 1, 1}, 0 },
		{ "player 1", []int{1, 0, 1, 1}, 1 },
		{ "player 2", []int{1, 1, 0, 1}, 2 },
		{ "player 3", []int{1, 1, 1, 0}, 3 },
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			p0 := &Player{nil, nil, 0, tc.scores[0]}
			p1 := &Player{nil, nil, 1, tc.scores[1]}
			p2 := &Player{nil, nil, 2, tc.scores[2]}
			p3 := &Player{nil, nil, 3, tc.scores[3]}
			game := &Game{nil, []*Player{p0, p1, p2, p3}}
			got := game.gameWinner()
			if got != tc.want {
				t.Errorf("gameWinner got %d want %d", got, tc.want)
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
			players := []*Player{{nil, tc.tricks[0], 0, 0},{nil, tc.tricks[1], 1, 0},{nil, tc.tricks[2], 2, 0},{nil, tc.tricks[3], 3, 0}}
			game := &Game{nil, players}
			game.updateScores()
			for i, p := range game.players {
				if p.score != tc.want[i] {
					t.Errorf("hand: got %v want %v", p.score, tc.want[i])
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
			p0 := &Player{nil, nil, 0, tc.scores[0]}
			p1 := &Player{nil, nil, 1, tc.scores[1]}
			p2 := &Player{nil, nil, 2, tc.scores[2]}
			p3 := &Player{nil, nil, 3, tc.scores[3]}
			game := &Game{nil, []*Player{p0, p1, p2, p3}}
			got := game.gameOver()
			if got != tc.want {
				t.Errorf("hand: got %t want %t", got, tc.want)
			}
		})
	}
}
