package main

type Ellipse struct {
	x         float64
	y         float64
	width     float64
	height    float64
	transform ITransform
}

func NewEllipse2D(posX, posY, width, height float64) IShape {
	return &Ellipse{posX, posY, width, height, nil}
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
	return &Ellipse{e.x, e.y, e.width, e.height, e.transform}
}

func (e *Ellipse) setTransform(transform ITransform) {
	e.transform = transform
}
