package main

type IShape interface {
	contains(x, y float64) bool
	draw(g IDoGraphics)
	copy() IShape
	setTransform(transform ITransform)
	move(dx, dy float64)
	moveTo(x, y float64)
	pos() Point
	intersects(shape IShape) bool
}
