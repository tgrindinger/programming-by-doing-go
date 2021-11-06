package main

type Point struct {
	x, y float64
}

func (p *Point) move(dx, dy float64) {
	p.x += dx
	p.y += dy
}

func (p *Point) moveTo(x, y float64) {
	p.x = x
	p.y = y
}

func (p *Point) pos() Point {
	return Point{p.x, p.y}
}
