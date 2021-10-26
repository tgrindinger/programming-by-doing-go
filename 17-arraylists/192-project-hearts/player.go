package main

import (
	"fmt"
	"strings"
)

type Player struct {
	hand   []*Card
	tricks []*Card
}

func (p *Player) printHand() string {
	builder := strings.Builder{}
	for i, card := range p.hand {
		builder.WriteString(fmt.Sprintf("%d) %s\n", i+1, card))
	}
	return builder.String()
}

func (p *Player) playCard(index int) *Card {
	card := p.hand[index]
	if index == len(p.hand) - 1 {
		p.hand = p.hand[:index]
	} else {
		p.hand = append(p.hand[:index], p.hand[index+1:]...)
	}
	return card
}

func (p *Player) receiveCard(card *Card) {
	p.hand = append(p.hand, card)
}

func (p *Player) receiveTrick(trick []*Card) {
	p.tricks = append(p.tricks, trick...)
}

func (p *Player) score() int {
	score := 0
	for _, c := range p.tricks {
		score += c.score()
	}
	return score
}

func (p *Player) hasNonHeartCards() bool {
	for _, c := range p.hand {
		if c.suit != Hearts {
			return true
		}
	}
	return false
}

func (p *Player) hasMatchingSuit(suit Suit) bool {
	for _, c := range p.hand {
		if c.suit == suit {
			return true
		}
	}
	return false
}

func (p *Player) sortHand() {
	for i := 0; i < len(p.hand); i++ {
		for j := i + 1; j < len(p.hand); j++ {
			if p.hand[i].compareTo(p.hand[j]) > 0 {
				temp := p.hand[i]
				p.hand[i] = p.hand[j]
				p.hand[j] = temp
			}
		}
	}
}

func (p *Player) resetTricks() {
	p.tricks = []*Card{}
}
