package main

type Rectangle struct {
	Point
	width, height float64
	transform     ITransform
}

func NewRectangleAsRect(posX, posY, width, height float64) *Rectangle {
	return &Rectangle{Point{posX, posY}, width, height, nil}
}

func NewRectangle(posX, posY, width, height float64) IShape {
	return &Rectangle{Point{posX, posY}, width, height, nil}
}

func (r *Rectangle) contains(x, y float64) bool {
	return x >= r.x && x <= r.x+r.width &&
		y >= r.y && y <= r.y+r.height
}

func (r *Rectangle) draw(g IDoGraphics) {
	g.setTransform(r.transform)
	g.doRect(r.x, r.y, r.width, r.height)
}

func (r *Rectangle) copy() IShape {
	return &Rectangle{Point{r.x, r.y}, r.width, r.height, r.transform}
}

func (r *Rectangle) setTransform(transform ITransform) {
	r.transform = transform
}

func (r *Rectangle) intersects(shape IShape) bool {
	return false
}
