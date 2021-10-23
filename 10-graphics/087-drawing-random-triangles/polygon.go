package main

import "github.com/gotk3/gotk3/cairo"

type Polygon struct {
	cr *cairo.Context
	first bool
}

func NewPolygon(cr *cairo.Context) *Polygon {
	return &Polygon{cr, true}
}

func (p *Polygon) addPoint(x, y float64) {
	if p.first {
		p.cr.MoveTo(x, y)
		p.first = false
	} else {
		p.cr.LineTo(x, y)
	}
}

func (p *Polygon) draw() {
	p.cr.ClosePath()
	p.cr.Stroke()
}

func (p *Polygon) fill() {
	p.cr.ClosePath()
	p.cr.Fill()
}
