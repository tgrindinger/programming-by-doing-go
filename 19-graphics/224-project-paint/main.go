package main

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func checkError(err error, name string) {
	if err != nil {
		log.Fatalf("Unable to create %s: %s", name, err)
	}
}

func createWindow(title string, width, height int) *gtk.Window {
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	checkError(err, "window")
	win.SetTitle(title)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	win.SetDefaultSize(width, height)
	return win
}

func createGrid() *gtk.Grid {
	grid, err := gtk.GridNew()
	checkError(err, "grid")
	grid.SetOrientation(gtk.ORIENTATION_VERTICAL)
	return grid
}

func createDrawingArea() *gtk.DrawingArea {
	area, err := gtk.DrawingAreaNew()
	checkError(err, "drawing area")
	area.AddEvents(int(gdk.BUTTON_PRESS_MASK))
	area.AddEvents(int(gdk.BUTTON_RELEASE_MASK))
	area.AddEvents(int(gdk.BUTTON_MOTION_MASK))
	area.Connect("draw", paint)
	area.Connect("event", daEvent)
	return area
}

func createButton(name string, arg interface{}, handler func(interface{}) bool) *gtk.Button {
	btn, err := gtk.ButtonNew()
	checkError(err, fmt.Sprintf("%s button", name))
	btn.SetLabel(name)
	btn.AddEvents(int(gdk.BUTTON_PRESS_MASK))
	btn.Connect("event", func(btn *gtk.Button, event *gdk.Event) bool {
		buttonEvent := gdk.EventButtonNewFromEvent(event)
		if buttonEvent.Type() == gdk.EVENT_BUTTON_PRESS {
			return handler(arg)
		}
		return true
	})
	return btn
}

func createBrushSizeButtons() []*gtk.Button {
	f := func(arg interface{}) bool {
		brushSize, ok := arg.(float64)
		if ok {
			painter.setBrushSize(brushSize)
			return false
		}
		return true
	}
	buttons := []*gtk.Button{
		createButton("Pen", 2.0, f),
		createButton("Brush", 10.0, f),
		createButton("Roller", 50.0, f),
	}
	return buttons
}

func createDrawingToolButtons() []*gtk.Button {
	f := func(arg interface{}) bool {
		tool, ok := arg.(DrawingTool)
		if ok {
			painter.setDrawingTool(tool)
			return false
		}
		return true
	}
	buttons := []*gtk.Button{
		createButton("Line", TOOL_LINE, f),
		createButton("Rectangle", TOOL_RECTANGLE, f),
		createButton("Ellipse", TOOL_ELLIPSE, f),
	}
	return buttons
}

func daEvent(da *gtk.DrawingArea, event *gdk.Event) bool {
	keyEvent := gdk.EventButtonNewFromEvent(event)
	x, y := keyEvent.MotionVal()
	eventMap := map[gdk.EventType]func(float64, float64){}
	eventMap[gdk.EVENT_BUTTON_PRESS] = painter.processPress
	eventMap[gdk.EVENT_BUTTON_RELEASE] = painter.processRelease
	eventMap[gdk.EVENT_MOTION_NOTIFY] = painter.processDrag
	handler, ok := eventMap[keyEvent.Type()]
	if ok {
		handler(x, y)
		da.QueueDraw()
		return false
	}
	return true
}

func paint(da *gtk.DrawingArea, cr *cairo.Context) {
	g := NewGraphics(cr)
	painter.paint(g)
}

var painter IPainter

func main() {
	gtk.Init(nil)
	width := 1024
	height := 768
	win := createWindow("Paint", width, height)
	area := createDrawingArea()
	painter = NewPaint(width, height)
	grid := createGrid()
	buttons := createBrushSizeButtons()
	buttons = append(buttons, createDrawingToolButtons()...)
	for _, b := range buttons {
		grid.Add(b)
	}
	grid.Attach(area, 1, 0, 1, 10)
	area.SetHExpand(true)
	area.SetVExpand(true)
	win.Add(grid)
	win.ShowAll()
	gtk.Main()
}
