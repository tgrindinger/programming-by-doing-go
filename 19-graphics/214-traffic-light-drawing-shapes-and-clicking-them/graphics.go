package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
)

type Color struct {
	R float64
	G float64
	B float64
}

func COLOR_BLACK() *Color { return &Color{0, 0, 0} }
func COLOR_WHITE() *Color { return &Color{1, 1, 1} }
func COLOR_RED()   *Color { return &Color{1, 0, 0} }
func COLOR_GREEN() *Color { return &Color{0, 1, 0} }
func COLOR_BLUE()  *Color { return &Color{0, 0, 1} }
func COLOR_CYAN()  *Color { return &Color{0, 1, 1} }
func COLOR_YELLOW()  *Color { return &Color{1, 1, 0} }
func COLOR_MAGENTA()  *Color { return &Color{1, 0, 1} }

type Graphics struct {
	cr *cairo.Context
}

func NewGraphics(cr *cairo.Context) *Graphics {
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

func (g *Graphics) drawLine(x1, y1, x2, y2 float64) {
	g.cr.MoveTo(x1, y1)
	g.cr.LineTo(x2, y2)
	g.cr.Stroke()
}

func (g *Graphics) drawString(msg string, x, y float64) {
	g.cr.MoveTo(x, y)
	g.cr.ShowText(msg)
}

func (g *Graphics) setColor(color *Color) {
	g.cr.SetSourceRGB(color.R, color.G, color.B)
}

func (g *Graphics) fill(shape IShape) {
	shape.draw(g)
	g.cr.Fill()
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
