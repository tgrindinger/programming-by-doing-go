package main

import (
	"errors"
	"math/rand"
)

type IMatch interface {
	dealCards()
	doFirstTurn()
	matchOver() bool
	playCard(card int) error
	trickOver() bool
	trickWinner() int
	collectTrick()
	matchTrick() []*DiscardedCard
	currentPlayer() *Player
}

type Match struct {
	random  *rand.Rand
	players []*Player
	trick   []*DiscardedCard
	current int
	broken  bool
}

func NewMatch(random *rand.Rand, players []*Player) IMatch {
	return &Match{random, players, []*DiscardedCard{}, 0, false}
}

func (m *Match) matchTrick() []*DiscardedCard {
	return m.trick
}

func (m *Match) currentPlayer() *Player {
	return m.players[m.current]
}

func (m *Match) nextPlayer() {
	m.current = (m.current + 1) % len(m.players)
}

func (m *Match) shuffle(deck []*Card) {
	for i := 0; i < len(deck); i++ {
		index := m.random.Intn(len(deck))
		temp := deck[i]
		deck[i] = deck[index]
		deck[index] = temp
	}
}

func (m *Match) newDeck() []*Card {
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

func (m *Match) dealCards() {
	deck := m.newDeck()
	m.shuffle(deck)
	m.current = 0
	for len(deck) > 0 {
		m.players[m.current].receiveCard(deck[0])
		deck = deck[1:]
		m.nextPlayer()
	}
	for _, p := range m.players {
		p.resetTricks()
		p.sortHand()
	}
	m.broken = false
}

func (m *Match) findFirstPlayer() (int, int) {
	for i, p := range m.players {
		for j, c := range p.hand {
			if c.face == Two && c.suit == Clubs {
				return i, j
			}
		}
	}
	panic("Couldn't find Two of Clubs!!!")
}

func (m *Match) doFirstTurn() {
	firstPlayer, cardIndex := m.findFirstPlayer()
	m.current = firstPlayer
	m.playCard(cardIndex)
}

func (m *Match) matchOver() bool {
	for _, p := range m.players {
		if len(p.hand) > 0 {
			return false
		}
	}
	return true
}

func (m *Match) trickOver() bool {
	return len(m.trick) == len(m.players)
}

func (m *Match) trickWinner() int {
	highest := 0
	for i, c := range m.trick {
		if c.card.suit == m.trick[highest].card.suit && c.card.compareTo(m.trick[highest].card) > 0 {
			highest = i
		}
	}
	return m.trick[highest].player
}

func (m *Match) collectTrick() {
	winner := m.trickWinner()
	var trick []*Card
	for _, c := range m.trick {
		trick = append(trick, c.card)
	}
	m.players[winner].receiveTrick(trick)
	m.trick = []*DiscardedCard{}
	m.current = winner
}

func (m *Match) playCard(index int) error {
	err := m.validPlay(index)
	if err != nil {
		return err
	}
	card := m.players[m.current].playCard(index)
	if card.suit == Hearts {
		m.broken = true
	}
	m.trick = append(m.trick, &DiscardedCard{card, m.current})
	m.nextPlayer()
	return nil
}

var ErrFirstPlay error = errors.New("invalid first play")
var ErrNextPlay error = errors.New("invalid next play")
var ErrInvalidChoice error = errors.New("invalid choice")

func (m *Match) validFirstPlay(index int) error {
	player := m.players[m.current]
	if player.hand[index].suit == Hearts && !m.broken && player.hasNonHeartCards() {
		return ErrFirstPlay
	}
	return nil
}

func (m *Match) validNextPlay(index int) error {
	player := m.players[m.current]
	if m.trick[0].card.suit != player.hand[index].suit && player.hasMatchingSuit(m.trick[0].card.suit) {
		return ErrNextPlay
	}
	return nil
}

func (m *Match) validPlay(index int) error {
	if index < 0 || index >= len(m.players[m.current].hand) {
		return ErrInvalidChoice
	}
	var err error
	if len(m.trick) == 0 {
		err = m.validFirstPlay(index)
	} else {
		err = m.validNextPlay(index)
	}
	return err
}
