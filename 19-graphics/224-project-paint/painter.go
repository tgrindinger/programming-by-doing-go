package main

import (
	"math"

	"github.com/gotk3/gotk3/cairo"
)

const (
	POOL_SIZE float64 = 20
)

type Pool struct {
	color Colors
	rect  IShape
}

func NewPool(color Colors, index int) *Pool {
	return &Pool{color, NewRectangle(POOL_SIZE * float64(index), 0, POOL_SIZE, POOL_SIZE)}
}

func (p *Pool) contains(x, y float64) bool {
	return p.rect.contains(x, y)
}

func (p *Pool) paint(g IGraphics) {
	g.setColor(COLOR_BLACK)
	g.draw(p.rect)
	g.setColor(p.color)
	g.fill(p.rect)
}

type Palette struct {
	pools []*Pool
}

func NewPalette() *Palette {
	pools := []*Pool{}
	colors := []Colors{COLOR_BLACK, COLOR_WHITE, COLOR_GRAY, COLOR_RED, COLOR_GREEN, COLOR_BLUE, COLOR_CYAN, COLOR_YELLOW, COLOR_MAGENTA}
	for i, c := range colors {
		pools = append(pools, NewPool(c, i))
	}
	return &Palette{pools}
}

func (p *Palette) sample(x, y float64) (Colors, bool) {
	for _, pool := range p.pools {
		if pool.contains(x, y) {
			return pool.color, true
		}
	}
	return -1, false
}

func (p *Palette) paint(g IGraphics) {
	for _, pool := range p.pools {
		pool.paint(g)
	}
}

type DrawingTool int

const (
	TOOL_NONE DrawingTool = iota
	TOOL_LINE
	TOOL_RECTANGLE
	TOOL_ELLIPSE
)

type IPainter interface {
	processDrag(x, y float64)
	processPress(x, y float64)
	processRelease(x, y float64)
	setBrushSize(size float64)
	paint(g IGraphics)
	setDrawingTool(tool DrawingTool)
}

type Painter struct {
	surface *cairo.Surface
	context *cairo.Context
	graphics IGraphics
	color Colors
	palette *Palette
	brushSize float64
	activeTool DrawingTool
	lastDrag *Point
	startPoint *Point
	currentPoint *Point
}

func NewPaint(width, height int) *Painter {
	surface := cairo.CreateImageSurface(cairo.FORMAT_RGB24, width, height)
	context := cairo.Create(surface)
	graphics := NewGraphics(context)
	color := COLOR_BLACK
	palette := NewPalette()
	brushSize := 10.0
	graphics.setColor(COLOR_WHITE)
	graphics.fillRect(0, 0, float64(width), float64(height))
	graphics.setLineWidth(brushSize)
	return &Painter{surface, context, graphics, color, palette, brushSize, TOOL_NONE, nil, nil, nil}
}

func (p *Painter) processPress(x, y float64) {
	if color, ok := p.palette.sample(x, y); ok {
		p.color = color
	} else if p.activeTool == TOOL_NONE {
		p.graphics.setColor(p.color)
		p.graphics.fillOval(x - p.brushSize / 2, y - p.brushSize / 2, p.brushSize, p.brushSize)
		p.lastDrag = &Point{x, y}
	} else {
		p.startPoint = &Point{x, y}
		p.currentPoint = &Point{x, y}
	}
}

func (p *Painter) processDrag(x, y float64) {
	if _, ok := p.palette.sample(x, y); !ok {
		if p.activeTool == TOOL_NONE {
			p.graphics.setColor(p.color)
			p.graphics.fillOval(x - p.brushSize / 2, y - p.brushSize / 2, p.brushSize, p.brushSize)
			p.graphics.drawLine(p.lastDrag.x, p.lastDrag.y, x, y)
			p.lastDrag.moveTo(x, y)
		} else {
			p.currentPoint.moveTo(x, y)
		}
	}
}

func (p *Painter) processRelease(x, y float64) {
	if p.activeTool != TOOL_NONE && p.startPoint != nil {
		p.paintActiveTool(p.graphics)
		p.activeTool = TOOL_NONE
		p.startPoint = nil
		p.currentPoint = nil
	} else {
		p.lastDrag = nil
	}
}

func (p *Painter) setBrushSize(brushSize float64) {
	p.brushSize = brushSize
	p.graphics.setLineWidth(brushSize)
}

func (p *Painter) setDrawingTool(tool DrawingTool) {
	p.activeTool = tool
}

func (p *Painter) paintActiveTool(g IGraphics) {
	if p.activeTool != TOOL_NONE && p.startPoint != nil && p.currentPoint != nil {
		g.setColor(p.color)
		g.setLineWidth(p.brushSize)
		left := math.Min(p.startPoint.x, p.currentPoint.x)
		right := math.Max(p.startPoint.x, p.currentPoint.x)
		top := math.Min(p.startPoint.y, p.currentPoint.y)
		bottom := math.Max(p.startPoint.y, p.currentPoint.y)
		switch p.activeTool {
		case TOOL_LINE: g.drawLine(p.startPoint.x, p.startPoint.y, p.currentPoint.x, p.currentPoint.y)
		case TOOL_RECTANGLE: g.drawRect(left, top, right - left, bottom - top)
		case TOOL_ELLIPSE: g.drawOval(left, top, right - left, bottom - top)
		}
	}
}

func (p *Painter) paint(g IGraphics) {
	g.drawImage(p.surface, 0, 0)
	p.palette.paint(g)
	p.paintActiveTool(g)
}
