package main

import (
	"fmt"
	"math/rand"
)

type Game struct {
	random  *rand.Rand
	players []*Player
	trick   []*DiscardedCard
	current int
	scores  []int
	broken  bool
}

type DiscardedCard struct {
	card   *Card
	player int
}

func (d *DiscardedCard) String() string {
	return fmt.Sprintf("%d) %s", d.player, d.card)
}

func NewGame(random *rand.Rand) *Game {
	return &Game{random, []*Player{{}, {}, {}, {}}, []*DiscardedCard{}, 0, []int{0, 0, 0, 0}, false}
}

func (g *Game) shuffle(deck []*Card) {
	for i := 0; i < len(deck); i++ {
		index := g.random.Intn(len(deck))
		temp := deck[i]
		deck[i] = deck[index]
		deck[index] = temp
	}
}

func (g *Game) newDeck() []*Card {
	suits := []Suit{Spades, Hearts, Diamonds, Clubs}
	faces := []Face{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}
	var deck []*Card
	for _, suit := range suits {
		for _, face := range faces {
			deck = append(deck, &Card{face, suit})
		}
	}
	return deck
}

func (g *Game) nextPlayer() {
	g.current = (g.current + 1) % len(g.players)
}

func (g *Game) dealCards() {
	deck := g.newDeck()
	g.shuffle(deck)
	g.current = 0
	for len(deck) > 0 {
		g.players[g.current].receiveCard(deck[0])
		deck = deck[1:]
		g.nextPlayer()
	}
	for _, p := range g.players {
		p.resetTricks()
		p.sortHand()
	}
	g.broken = false
}

func (g *Game) findFirstPlayer() (int, int) {
	for i, p := range g.players {
		for j, c := range p.hand {
			if c.face == Two && c.suit == Clubs {
				return i, j
			}
		}
	}
	panic("Couldn't find Two of Clubs!!!")
}

func (g *Game) doFirstTurn() {
	firstPlayer, cardIndex := g.findFirstPlayer()
	g.current = firstPlayer
	g.playCard(cardIndex)
}

func (g *Game) matchOver() bool {
	for _, p := range g.players {
		if len(p.hand) > 0 {
			return false
		}
	}
	return true
}

func (g *Game) updateScores() {
	for i, p := range g.players {
		if p.score() == 26 {
			for j := range g.players {
				if i != j {
					g.scores[j] += 26
				}
			}
			return
		}
	}
	for i, p := range g.players {
		g.scores[i] += p.score()
	}
}

func (g *Game) gameOver() bool {
	sum := 0
	for _, s := range g.scores {
		sum += s
	}
	return sum >= 100
}

func (g *Game) trickOver() bool {
	return len(g.trick) == len(g.players)
}

func (g *Game) trickWinner() int {
	highest := 0
	for i, c := range g.trick {
		if c.card.suit == g.trick[highest].card.suit && c.card.compareTo(g.trick[highest].card) > 0 {
			highest = i
		}
	}
	return g.trick[highest].player
}

func (g *Game) collectTrick() {
	winner := g.trickWinner()
	var trick []*Card
	for _, c := range g.trick {
		trick = append(trick, c.card)
	}
	g.players[winner].receiveTrick(trick)
	g.trick = []*DiscardedCard{}
	g.current = winner
}

func (g *Game) playCard(index int) {
	card := g.players[g.current].playCard(index)
	if card.suit == Hearts {
		g.broken = true
	}
	g.trick = append(g.trick, &DiscardedCard{card, g.current})
	g.nextPlayer()
}

func (g *Game) validFirstPlay(index int) bool {
	return len(g.trick) == 0 && (g.players[g.current].hand[index].suit != Hearts || g.broken || !g.players[g.current].hasNonHeartCards())
}

func (g *Game) validNextPlay(index int) bool {
	return len(g.trick) > 0 && (g.trick[0].card.suit == g.players[g.current].hand[index].suit || !g.players[g.current].hasMatchingSuit(g.trick[0].card.suit))
}

func (g *Game) validPlay(index int) bool {
	return index < len(g.players[g.current].hand) && (g.validFirstPlay(index) || g.validNextPlay(index))
}
