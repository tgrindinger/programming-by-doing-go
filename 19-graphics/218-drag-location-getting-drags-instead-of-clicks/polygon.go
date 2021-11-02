package main

type Point struct {
	x, y float64
}

type IPolygon interface {
	draw(g IGraphics)
}

type Polygon struct {
	points    []*Point
	transform ITransform
}

func NewPolygon(xpoints []float64, ypoints []float64) IShape {
	points := []*Point{}
	for i := range xpoints {
		points = append(points, &Point{xpoints[i], ypoints[i]})
	}
	return &Polygon{points, nil}
}

func (p *Polygon) contains(x, y float64) bool {
	return false
}

func (p *Polygon) draw(g IDoGraphics) {
	g.setTransform(p.transform)
	g.doPolygon(p.points)
}

func (p *Polygon) copy() IShape {
	return &Polygon{p.points, p.transform}
}

func (p *Polygon) setTransform(transform ITransform) {
	p.transform = transform
}
