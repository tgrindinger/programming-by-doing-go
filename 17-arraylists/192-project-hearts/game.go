package main

import (
	"math/rand"
)

type IGame interface {
	gameOver() bool
	updateScores()
	gameWinner() int
	gameScores() []int
	newMatch() IMatch
}

type Game struct {
	random  *rand.Rand
	players []*Player
}

func NewGame(random *rand.Rand) IGame {
	return &Game{random, []*Player{NewPlayer(0), NewPlayer(1), NewPlayer(2), NewPlayer(3)}}
}

func (g *Game) newMatch() IMatch {
	return NewMatch(g.random, g.players)
}

func (g *Game) gameScores() []int {
	scores := make([]int, len(g.players))
	for i := 0; i < len(g.players); i++ {
		scores[i] = g.players[i].score
	}
	return scores
}

func (g *Game) gameWinner() int {
	lowest := 0
	for i, p := range g.players {
		if p.score < g.players[lowest].score {
			lowest = i
		}
	}
	return lowest
}

func (g *Game) updateScores() {
	for i, p := range g.players {
		if p.currentScore() == 26 {
			for j := range g.players {
				if i != j {
					g.players[j].score += 26
				}
			}
			return
		}
	}
	for _, p := range g.players {
		p.updateScore()
	}
}

func (g *Game) gameOver() bool {
	sum := 0
	for _, p := range g.players {
		sum += p.score
	}
	return sum >= 100
}
