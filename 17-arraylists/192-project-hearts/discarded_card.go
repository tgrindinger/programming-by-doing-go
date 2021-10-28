package main

import "fmt"

type DiscardedCard struct {
	card   *Card
	player int
}

func (d *DiscardedCard) String() string {
	return fmt.Sprintf("%d) %s", d.player, d.card)
}
