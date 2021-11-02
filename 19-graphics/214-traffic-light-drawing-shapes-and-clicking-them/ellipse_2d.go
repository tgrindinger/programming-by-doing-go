package main

type IShape interface {
	contains(x, y float64) bool
	draw(g *Graphics)
}

type Ellipse2D struct {
	x      float64
	y      float64
	width  float64
	height float64
}

func NewEllipse2D(posX, posY, width, height float64) IShape {
	return &Ellipse2D{posX, posY, width, height}
}

func (e *Ellipse2D) contains(x, y float64) bool {
	rx := e.width / 2
	ry := e.height / 2
	h := e.x + rx
	k := e.y + ry
	left := ((x - h) * (x - h)) / (rx * rx)
	right := ((y - k) * (y - k)) / (ry * ry)
	return left+right <= 1
}

func (e *Ellipse2D) draw(g *Graphics) {
	g.doOval(e.x, e.y, e.width, e.height)
}
