package main

import "math"

type Ellipse struct {
	Point
	width     float64
	height    float64
	transform ITransform
}

func NewEllipse(posX, posY, width, height float64) IShape {
	return &Ellipse{Point{posX, posY}, width, height, nil}
}

func (e *Ellipse) contains(x, y float64) bool {
	rx := e.width / 2
	ry := e.height / 2
	h := e.x + rx
	k := e.y + ry
	left := ((x - h) * (x - h)) / (rx * rx)
	right := ((y - k) * (y - k)) / (ry * ry)
	return left+right <= 1
}

func (e *Ellipse) draw(g IDoGraphics) {
	g.setTransform(e.transform)
	g.doOval(e.x, e.y, e.width, e.height)
}

func (e *Ellipse) copy() IShape {
	return &Ellipse{Point{e.x, e.y}, e.width, e.height, e.transform}
}

func (e *Ellipse) setTransform(transform ITransform) {
	e.transform = transform
}

func inside(p *Point, left, right, top, bottom float64) bool {
	return p.x >= left && p.x <= right &&
		p.y >= top && p.y <= bottom
}

func (e *Ellipse) intersects(shape IShape) bool {
	rect, ok := shape.(*Rectangle)
	if ok {
		eps := []*Point{
			{e.x, e.y},
			{e.x + e.width, e.y},
			{e.x + e.width, e.y + e.height},
			{e.x, e.y + e.height},
		}
		center := &Point{e.x + e.width / 2, e.y + e.height / 2}
		for _, ep := range eps {
			if inside(ep, rect.x, rect.x + rect.width, rect.y, rect.y + rect.height) {
				dx := ep.x - center.x
				dy := ep.y - center.y
				dist := math.Sqrt(dx * dx + dy * dy)
				ratio := (e.width / 2) / dist
				if inside(&Point{center.x + dx * ratio, center.y + dy * ratio}, rect.x, rect.x + rect.width, rect.y, rect.y + rect.height) {
					return true
				}
			}
		}
	}
	return false
}
