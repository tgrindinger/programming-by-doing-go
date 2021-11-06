package main

import "fmt"

type Direction struct {
	dr int
	dc int
}

var EAST      *Direction = &Direction{ 0,  1}
var NORTHEAST *Direction = &Direction{-1,  1}
var NORTH     *Direction = &Direction{-1,  0}
var NORTHWEST *Direction = &Direction{-1, -1}
var WEST      *Direction = &Direction{ 0, -1}
var SOUTHWEST *Direction = &Direction{ 1, -1}
var SOUTH     *Direction = &Direction{ 1,  0}
var SOUTHEAST *Direction = &Direction{ 1,  1}
var compass []*Direction = []*Direction{EAST, NORTHEAST, NORTH, NORTHWEST, WEST, SOUTHWEST, SOUTH, SOUTHEAST}

func NewDirection(r, c int) *Direction {
	return &Direction{r, c}
}

func (d *Direction) String() string {
	if d.dr == -1 && d.dc == 0 {
		return "NORTH"
	} else if d.dr == 1 && d.dc == 0 {
		return "SOUTH"
	} else if d.dr == 0 && d.dc == 1 {
		return "EAST"
	} else if d.dr == 0 && d.dc == -1 {
		return "WEST"
	} else if d.dr == -1 && d.dc == -1 {
		return "NORTHWEST"
	} else if d.dr == 1 && d.dc == -1 {
		return "SOUTHWEST"
	} else if d.dr == 1 && d.dc == 1 {
		return "NORTHEAST"
	} else if d.dr == -1 && d.dc == 1 {
		return "SOUTHEAST"
	} else {
		return fmt.Sprintf("{%d,%d}", d.dr, d.dc)
	}
}
