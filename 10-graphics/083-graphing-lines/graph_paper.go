package main

import (
	"github.com/gotk3/gotk3/cairo"
)

type GraphPaper struct {
	cr *cairo.Context
	x, y, w, h float64
}

func NewGraphPaper(cr *cairo.Context, x, y, w, h float64) *GraphPaper {
	return &GraphPaper{ cr, x, y, w, h }
}

func (g *GraphPaper) drawBounds() {
	g.cr.SetSourceRGB(0.8, 0.8, 0.8)
	for d := 0.0; d < g.w; d += g.w / 20 {
		drawLine(g.cr, g.x+d, g.y+0, g.x+d, g.y+g.h)
	}
	for d := 0.0; d < g.h; d += g.h / 20 {
		drawLine(g.cr, g.x+0, g.y+d, g.x+g.h, g.y+d)
	}
	g.cr.SetSourceRGB(0, 0, 0)
	drawRect(g.cr, g.x, g.y, g.w+1, g.h+1)
	drawLine(g.cr, g.x+g.w/2, g.y+0, g.x+g.w/2, g.y+g.h)
	drawLine(g.cr, g.x+0, g.y+g.h/2, g.x+g.w, g.y+g.h/2)
}

func (g *GraphPaper) drawPoint(px, py float64) {
	if px > 10 || px < -10 || py > 10 || py < -10 {
		return
	}
	px *= g.w / 20
	py *= g.h / 20
	px += g.w / 2 + 1
	py = g.h / 2 - py + 1
	g.cr.MoveTo(g.x + px, g.y + py)
	g.cr.LineTo(g.x + px + 1, g.y + py + 1)
	g.cr.Stroke()
}

func (g *GraphPaper) drawLinePoints(m, b float64) {
	g.cr.SetSourceRGB(0, 0, 1)
	for px := -10.0; px <= 10; px += 0.01 {
		py := m * px + b;
		g.drawPoint(px, py)
	}
}
