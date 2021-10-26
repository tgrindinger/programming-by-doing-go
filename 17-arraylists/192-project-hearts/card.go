package main

import "fmt"

type Suit int
type Face int

const (
	Spades Suit = iota
	Hearts
	Diamonds
	Clubs
)

func (s Suit) String() string {
	return []string{"Spades", "Hearts", "Diamonds", "Clubs"}[s]
}

const (
	Two Face = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

func (f Face) String() string {
	return []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}[f]
}

type Card struct {
	face Face
	suit Suit
}

func (c *Card) String() string {
	return fmt.Sprintf("%s of %s", c.face, c.suit)
}

func (c *Card) compareTo(card *Card) int {
	result := -1
	if c.suit > card.suit || (c.suit == card.suit && c.face > card.face) {
		result = 1
	}
	return result
}

func (c *Card) score() int {
	value := 0
	if c.suit == Hearts {
		value = 1
	} else if c.suit == Spades && c.face == Queen {
		value = 13
	}
	return value
}
