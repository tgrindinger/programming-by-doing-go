package main

type Line struct {
	x1, y1, x2, y2 float64
	transform      ITransform
}

func NewLine(x1, y1, x2, y2 float64) IShape {
	return &Line{x1, y1, x2, y2, nil}
}

func (l *Line) contains(x, y float64) bool {
	return false
}

func (l *Line) draw(g IDoGraphics) {
	g.setTransform(l.transform)
	g.doLine(l.x1, l.y1, l.x2, l.y2)
}

func (l *Line) copy() IShape {
	return &Line{l.x1, l.y1, l.x2, l.y2, l.transform}
}

func (l *Line) setTransform(transform ITransform) {
	l.transform = transform
}
