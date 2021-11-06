package main

import "fmt"

type Location struct {
	r int
	c int
}

func NewLocation(r, c int) *Location {
	return &Location{r, c}
}

func (l *Location) getAdjacentLocation(dir *Direction) *Location {
	next := NewLocation(l.r, l.c)
	next.r += dir.dr
	next.c += dir.dc
	return next
}

func (l *Location) getAdjacentLocationCount(dir *Direction, count int) *Location {
	next := NewLocation(l.r, l.c)
	for i := 0; i < count; i++ {
		next.r += dir.dr
		next.c += dir.dc
	}
	return next
}

func (l *Location) String() string {
	return fmt.Sprintf("(%d,%d)", l.r, l.c)
}

func (l *Location) equals(other *Location) bool {
	return l.r == other.r && l.c == other.c
}
