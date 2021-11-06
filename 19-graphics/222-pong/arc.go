package main

type ArcType int

const (
	ARC_PIE ArcType = iota
	ARC_OPEN
	ARC_CHORD
)

type Arc struct {
	Point
	w, h, start, extent float64
	arcType             ArcType
	transform           ITransform
}

func NewArc(x, y, w, h, start, extent float64, arcType ArcType) IShape {
	return &Arc{Point{x, y}, w, h, start, extent, arcType, nil}
}

func (a *Arc) contains(x, y float64) bool {
	return false
}

func (a *Arc) draw(g IDoGraphics) {
	g.setTransform(a.transform)
	g.doArcType(a.x, a.y, a.w, a.h, a.start, a.extent, a.arcType)
}

func (a *Arc) copy() IShape {
	return &Arc{Point{a.x, a.y}, a.w, a.h, a.start, a.extent, a.arcType, a.transform}
}

func (a *Arc) setTransform(transform ITransform) {
	a.transform = transform
}

func (a *Arc) intersects(shape IShape) bool {
	return false
}
