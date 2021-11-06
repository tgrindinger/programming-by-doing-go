package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
)

type Colors int

const (
	COLOR_BLACK Colors = iota
	COLOR_WHITE
	COLOR_RED
	COLOR_GREEN
	COLOR_BLUE
	COLOR_CYAN
	COLOR_YELLOW
	COLOR_MAGENTA
)

type Color struct {
	R float64
	G float64
	B float64
}

type IGraphics interface {
	drawRect(x, y, w, h float64)
	fillRect(x, y, w, h float64)
	drawOval(x, y, w, h float64)
	fillOval(x, y, w, h float64)
	drawArc(x, y, w, h, a1, a2 float64)
	fillArc(x, y, w, h, a1, a2 float64)
	drawLine(x1, y1, x2, y2 float64)
	drawString(msg string, x, y float64)
	drawImage(surface *cairo.Surface, w, h float64)
	setColor(color Colors)
	draw(shape IShape)
	fill(shape IShape)
}

type IDoGraphics interface {
	doLine(x1, y1, x2, y2 float64)
	doRect(x, y, w, h float64)
	doArcWithAngles(x, y, w, h, a1, a2 float64)
	doOval(x, y, w, h float64)
	doPolygon(points []*Point)
	doArcType(x, y, w, h, a1, a2 float64, arcType ArcType)
	setTransform(t ITransform)
}

type Graphics struct {
	cr *cairo.Context
}

func NewGraphics(cr *cairo.Context) IGraphics {
	return &Graphics{cr}
}

func (g *Graphics) doRect(x, y, w, h float64) {
	g.cr.Rectangle(x, y, w, h)
}

func (g *Graphics) drawRect(x, y, w, h float64) {
	g.doRect(x, y, w, h)
	g.cr.Stroke()
}

func (g *Graphics) fillRect(x, y, w, h float64) {
	g.doRect(x, y, w, h)
	g.cr.Fill()
}

func (g *Graphics) doArcWithAngles(x, y, w, h, a1, a2 float64) {
	g.cr.Save()
	radX := w / 2
	radY := h / 2
	g.cr.Translate(x + radX, y + radY)
	minRad := radX
	if radX > radY {
		minRad = radY
	}
	g.cr.Scale(radX / minRad, radY / minRad)
	rad1 := a1 * (math.Pi / 180)
	rad2 := a2 * (math.Pi / 180)
	g.cr.Arc(0, 0, minRad, rad1, rad2)
	g.cr.Restore()
}

func (g *Graphics) doArcType(x, y, w, h, a1, a2 float64, arcType ArcType) {
	if arcType == ARC_PIE {
		g.cr.NewPath()
		g.cr.MoveTo(x + w / 2, y + h / 2)
	}
	g.doArcWithAngles(x, y, w, h, a1, a2)
	if arcType == ARC_PIE || arcType == ARC_CHORD {
		g.cr.ClosePath()
	}
}

func (g *Graphics) doOval(x, y, w, h float64) {
	g.doArcWithAngles(x, y, w, h, 0, 360)
}

func (g *Graphics) drawOval(x, y, w, h float64) {
	g.doOval(x, y, w, h)
	g.cr.Stroke()
}

func (g *Graphics) fillOval(x, y, w, h float64) {
	g.doOval(x, y, w, h)
	g.cr.Fill()
}

func (g *Graphics) drawArc(x, y, w, h, a1, a2 float64) {
	g.doArcWithAngles(x, y, w, h, a1, a2)
	g.cr.Stroke()
}

func (g *Graphics) fillArc(x, y, w, h, a1, a2 float64) {
	g.cr.NewPath()
	g.cr.MoveTo(x + w / 2, y + h / 2)
	g.doArcWithAngles(x, y, w, h, a1, a2)
	g.cr.ClosePath()
	g.cr.Fill()
}

func (g *Graphics) doLine(x1, y1, x2, y2 float64) {
	g.cr.MoveTo(x1, y1)
	g.cr.LineTo(x2, y2)
}

func (g *Graphics) drawLine(x1, y1, x2, y2 float64) {
	g.doLine(x1, y1, x2, y2)
	g.cr.Stroke()
}

func (g *Graphics) drawString(msg string, x, y float64) {
	g.cr.MoveTo(x, y)
	g.cr.ShowText(msg)
}

func (g *Graphics) drawImage(surface *cairo.Surface, x, y float64) {
	g.cr.SetSourceSurface(surface, x, y)
	g.fillRect(x, y, float64(surface.GetWidth()), float64(surface.GetHeight()))
}

func (g *Graphics) setColor(color Colors) {
	var c Color
	switch color {
		case COLOR_BLACK:   c = Color{0, 0, 0}
		case COLOR_WHITE:   c = Color{1, 1, 1}
		case COLOR_RED:     c = Color{1, 0, 0}
		case COLOR_GREEN:   c = Color{0, 1, 0}
		case COLOR_BLUE:    c = Color{0, 0, 1}
		case COLOR_CYAN:    c = Color{0, 1, 1}
		case COLOR_YELLOW:  c = Color{1, 1, 0}
		case COLOR_MAGENTA: c = Color{1, 0, 1}
	}
	g.cr.SetSourceRGB(c.R, c.G, c.B)
}

func (g *Graphics) draw(shape IShape) {
	g.cr.Save()
	shape.draw(g)
	g.cr.Stroke()
	g.cr.Restore()
}

func (g *Graphics) fill(shape IShape) {
	g.cr.Save()
	shape.draw(g)
	g.cr.Fill()
	g.cr.Restore()
}

func (g *Graphics) doPolygon(points []*Point) {
	first := true
	for _, p := range points {
		if first {
			g.cr.MoveTo(p.x, p.y)
			first = false
		} else {
			g.cr.LineTo(p.x, p.y)
		}
	}
	g.cr.ClosePath()
}

func (g *Graphics) setTransform(t ITransform) {
	if t != nil {
		g.setScale(t.getScale())
		g.setTranslate(t.getTranslate())
	}
}

func (g *Graphics) setScale(sx, sy float64) {
	g.cr.Scale(sx, sy)
}

func (g *Graphics) setTranslate(tx, ty float64) {
	g.cr.Translate(tx, ty)
}

var eventMap map[gdk.EventType]string = map[gdk.EventType]string{}

func buildMap() {
	eventMap[gdk.EVENT_NOTHING] = "EVENT_NOTHING"
	eventMap[gdk.EVENT_DELETE] = "EVENT_DELETE"
	eventMap[gdk.EVENT_DESTROY] = "EVENT_DESTROY"
	eventMap[gdk.EVENT_EXPOSE] = "EVENT_EXPOSE"
	eventMap[gdk.EVENT_MOTION_NOTIFY] = "EVENT_MOTION_NOTIFY"
	eventMap[gdk.EVENT_BUTTON_PRESS] = "EVENT_BUTTON_PRESS"
	eventMap[gdk.EVENT_2BUTTON_PRESS] = "EVENT_2BUTTON_PRESS"
	eventMap[gdk.EVENT_DOUBLE_BUTTON_PRESS] = "EVENT_DOUBLE_BUTTON_PRESS"
	eventMap[gdk.EVENT_3BUTTON_PRESS] = "EVENT_3BUTTON_PRESS"
	eventMap[gdk.EVENT_TRIPLE_BUTTON_PRESS] = "EVENT_TRIPLE_BUTTON_PRESS"
	eventMap[gdk.EVENT_BUTTON_RELEASE] = "EVENT_BUTTON_RELEASE"
	eventMap[gdk.EVENT_KEY_PRESS] = "EVENT_KEY_PRESS"
	eventMap[gdk.EVENT_KEY_RELEASE] = "EVENT_KEY_RELEASE"
	eventMap[gdk.EVENT_ENTER_NOTIFY] = "EVENT_ENTER_NOTIFY"
	eventMap[gdk.EVENT_LEAVE_NOTIFY] = "EVENT_LEAVE_NOTIFY"
	eventMap[gdk.EVENT_FOCUS_CHANGE] = "EVENT_FOCUS_CHANGE"
	eventMap[gdk.EVENT_CONFIGURE] = "EVENT_CONFIGURE"
	eventMap[gdk.EVENT_MAP] = "EVENT_MAP"
	eventMap[gdk.EVENT_UNMAP] = "EVENT_UNMAP"
	eventMap[gdk.EVENT_PROPERTY_NOTIFY] = "EVENT_PROPERTY_NOTIFY"
	eventMap[gdk.EVENT_SELECTION_CLEAR] = "EVENT_SELECTION_CLEAR"
	eventMap[gdk.EVENT_SELECTION_REQUEST] = "EVENT_SELECTION_REQUEST"
	eventMap[gdk.EVENT_SELECTION_NOTIFY] = "EVENT_SELECTION_NOTIFY"
	eventMap[gdk.EVENT_PROXIMITY_IN] = "EVENT_PROXIMITY_IN"
	eventMap[gdk.EVENT_PROXIMITY_OUT] = "EVENT_PROXIMITY_OUT"
	eventMap[gdk.EVENT_DRAG_ENTER] = "EVENT_DRAG_ENTER"
	eventMap[gdk.EVENT_DRAG_LEAVE] = "EVENT_DRAG_LEAVE"
	eventMap[gdk.EVENT_DRAG_MOTION] = "EVENT_DRAG_MOTION"
	eventMap[gdk.EVENT_DRAG_STATUS] = "EVENT_DRAG_STATUS"
	eventMap[gdk.EVENT_DROP_START] = "EVENT_DROP_START"
	eventMap[gdk.EVENT_DROP_FINISHED] = "EVENT_DROP_FINISHED"
	eventMap[gdk.EVENT_CLIENT_EVENT] = "EVENT_CLIENT_EVENT"
	eventMap[gdk.EVENT_VISIBILITY_NOTIFY] = "EVENT_VISIBILITY_NOTIFY"
	eventMap[gdk.EVENT_SCROLL] = "EVENT_SCROLL"
	eventMap[gdk.EVENT_WINDOW_STATE] = "EVENT_WINDOW_STATE"
	eventMap[gdk.EVENT_SETTING] = "EVENT_SETTING"
	eventMap[gdk.EVENT_OWNER_CHANGE] = "EVENT_OWNER_CHANGE"
	eventMap[gdk.EVENT_GRAB_BROKEN] = "EVENT_GRAB_BROKEN"
	eventMap[gdk.EVENT_DAMAGE] = "EVENT_DAMAGE"
	eventMap[gdk.EVENT_TOUCH_BEGIN] = "EVENT_TOUCH_BEGIN"
	eventMap[gdk.EVENT_TOUCH_UPDATE] = "EVENT_TOUCH_UPDATE"
	eventMap[gdk.EVENT_TOUCH_END] = "EVENT_TOUCH_END"
	eventMap[gdk.EVENT_TOUCH_CANCEL] = "EVENT_TOUCH_CANCEL"
	eventMap[gdk.EVENT_LAST] = "EVENT_LAST"
}

func printEventTypes() string {
	builder := &strings.Builder{}
	eventTypes := []struct {
		name  string
		value gdk.EventType
	}{
		{ "EVENT_NOTHING", gdk.EVENT_NOTHING },
		{ "EVENT_DELETE", gdk.EVENT_DELETE },
		{ "EVENT_DESTROY", gdk.EVENT_DESTROY },
		{ "EVENT_EXPOSE", gdk.EVENT_EXPOSE },
		{ "EVENT_MOTION_NOTIFY", gdk.EVENT_MOTION_NOTIFY },
		{ "EVENT_BUTTON_PRESS", gdk.EVENT_BUTTON_PRESS },
		{ "EVENT_2BUTTON_PRESS", gdk.EVENT_2BUTTON_PRESS },
		{ "EVENT_DOUBLE_BUTTON_PRESS", gdk.EVENT_DOUBLE_BUTTON_PRESS },
		{ "EVENT_3BUTTON_PRESS", gdk.EVENT_3BUTTON_PRESS },
		{ "EVENT_TRIPLE_BUTTON_PRESS", gdk.EVENT_TRIPLE_BUTTON_PRESS },
		{ "EVENT_BUTTON_RELEASE", gdk.EVENT_BUTTON_RELEASE },
		{ "EVENT_KEY_PRESS", gdk.EVENT_KEY_PRESS },
		{ "EVENT_KEY_RELEASE", gdk.EVENT_KEY_RELEASE },
		{ "EVENT_ENTER_NOTIFY", gdk.EVENT_ENTER_NOTIFY },
		{ "EVENT_LEAVE_NOTIFY", gdk.EVENT_LEAVE_NOTIFY },
		{ "EVENT_FOCUS_CHANGE", gdk.EVENT_FOCUS_CHANGE },
		{ "EVENT_CONFIGURE", gdk.EVENT_CONFIGURE },
		{ "EVENT_MAP", gdk.EVENT_MAP },
		{ "EVENT_UNMAP", gdk.EVENT_UNMAP },
		{ "EVENT_PROPERTY_NOTIFY", gdk.EVENT_PROPERTY_NOTIFY },
		{ "EVENT_SELECTION_CLEAR", gdk.EVENT_SELECTION_CLEAR },
		{ "EVENT_SELECTION_REQUEST", gdk.EVENT_SELECTION_REQUEST },
		{ "EVENT_SELECTION_NOTIFY", gdk.EVENT_SELECTION_NOTIFY },
		{ "EVENT_PROXIMITY_IN", gdk.EVENT_PROXIMITY_IN },
		{ "EVENT_PROXIMITY_OUT", gdk.EVENT_PROXIMITY_OUT },
		{ "EVENT_DRAG_ENTER", gdk.EVENT_DRAG_ENTER },
		{ "EVENT_DRAG_LEAVE", gdk.EVENT_DRAG_LEAVE },
		{ "EVENT_DRAG_MOTION", gdk.EVENT_DRAG_MOTION },
		{ "EVENT_DRAG_STATUS", gdk.EVENT_DRAG_STATUS },
		{ "EVENT_DROP_START", gdk.EVENT_DROP_START },
		{ "EVENT_DROP_FINISHED", gdk.EVENT_DROP_FINISHED },
		{ "EVENT_CLIENT_EVENT", gdk.EVENT_CLIENT_EVENT },
		{ "EVENT_VISIBILITY_NOTIFY", gdk.EVENT_VISIBILITY_NOTIFY },
		{ "EVENT_SCROLL", gdk.EVENT_SCROLL },
		{ "EVENT_WINDOW_STATE", gdk.EVENT_WINDOW_STATE },
		{ "EVENT_SETTING", gdk.EVENT_SETTING },
		{ "EVENT_OWNER_CHANGE", gdk.EVENT_OWNER_CHANGE },
		{ "EVENT_GRAB_BROKEN", gdk.EVENT_GRAB_BROKEN },
		{ "EVENT_DAMAGE", gdk.EVENT_DAMAGE },
		{ "EVENT_TOUCH_BEGIN", gdk.EVENT_TOUCH_BEGIN },
		{ "EVENT_TOUCH_UPDATE", gdk.EVENT_TOUCH_UPDATE },
		{ "EVENT_TOUCH_END", gdk.EVENT_TOUCH_END },
		{ "EVENT_TOUCH_CANCEL", gdk.EVENT_TOUCH_CANCEL },
		{ "EVENT_LAST", gdk.EVENT_LAST },
	}
	fmt.Fprintf(builder, "Event types:\n")
	for _, t := range eventTypes {
		fmt.Fprintf(builder, "%s = %d\n", t.name, t.value)
	}
	return builder.String()
}
