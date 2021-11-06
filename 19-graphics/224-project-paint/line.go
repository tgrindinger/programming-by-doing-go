package main

type Line struct {
	Point
	p2        Point
	transform ITransform
}

func NewLine(x1, y1, x2, y2 float64) IShape {
	return &Line{Point{x1, y1}, Point{x2, y2}, nil}
}

func (l *Line) contains(x, y float64) bool {
	return false
}

func (l *Line) draw(g IDoGraphics) {
	g.setTransform(l.transform)
	g.doLine(l.x, l.y, l.p2.x, l.p2.y)
}

func (l *Line) copy() IShape {
	return &Line{Point{l.x, l.y}, Point{l.p2.x, l.p2.y}, l.transform}
}

func (l *Line) setTransform(transform ITransform) {
	l.transform = transform
}

func (l *Line) move(dx, dy float64) {
	l.Point.move(dx, dy)
	l.p2.move(dx, dy)
}

func (l *Line) intersects(shape IShape) bool {
	return false
}
