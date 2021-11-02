package main

type IShapeTransformer interface {
	scale(x, y float64)
	translate(x, y float64)
	createTransformedShape(shape IShape) IShape
}

type ITransform interface {
	getScale() (x, y float64)
	getTranslate() (x, y float64)
}

type AffineTransform struct {
	sx, sy, tx, ty float64
}

func NewAffineTransform() IShapeTransformer {
	return &AffineTransform{}
}

func (a *AffineTransform) scale(x, y float64) {
	a.sx = x
	a.sy = y
}

func (a *AffineTransform) translate(x, y float64) {
	a.tx = x
	a.ty = y
}

func (a *AffineTransform) getScale() (sx, sy float64) {
	return a.sx, a.sy
}

func (a *AffineTransform) getTranslate() (tx, ty float64) {
	return a.tx, a.ty
}

func (a *AffineTransform) createTransformedShape(shape IShape) IShape {
	newShape := shape.copy()
	newShape.setTransform(a)
	return newShape
}
