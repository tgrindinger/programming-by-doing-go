package main

type IShape interface {
	contains(x, y float64) bool
	draw(g IDoGraphics)
	copy() IShape
	setTransform(transform ITransform)
}
