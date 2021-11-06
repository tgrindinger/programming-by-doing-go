package main

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

func (p *Polygon) move(dx, dy float64) {
	for _, point := range p.points {
		point.move(dx, dy)
	}
}

func (p *Polygon) moveTo(x, y float64) {
	ref := Point{p.points[0].x, p.points[1].y}
	for _, point := range p.points {
		point.moveTo(point.x-ref.x+x, point.y-ref.y+y)
	}
}

func (p *Polygon) pos() Point {
	return p.points[0].pos()
}

func (p *Polygon) intersects(shape IShape) bool {
	return false
}
