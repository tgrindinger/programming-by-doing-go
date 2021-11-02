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

func createDrawingArea() *gtk.DrawingArea {
	area, err := gtk.DrawingAreaNew()
	checkError(err, "drawing area")
	area.AddEvents(int(gdk.BUTTON_PRESS_MASK))
	area.Connect("draw", paint)
	area.Connect("event", daEvent)
	return area
}

func daEvent(da *gtk.DrawingArea, event *gdk.Event) bool {
	eventClick := gdk.EventButtonNewFromEvent(event)
	if eventClick.Type() == gdk.EVENT_BUTTON_PRESS {
		x, y := eventClick.MotionVal()
		if circle1.contains(x, y) {
			message = "You clicked on the red circle"
		} else if circle2.contains(x, y) {
			message = "You clicked on the yellow circle"
		} else if circle3.contains(x, y) {
			message = "You clicked on the green circle"
		} else {
			message = fmt.Sprintf("You clicked at (%.0f,%.0f)", x, y)
		}
		da.QueueDraw()
		return false
	}
	return true
}

func paint(da *gtk.DrawingArea, cr *cairo.Context) {
	g := NewGraphics(cr)
	g.setColor(COLOR_BLACK())
	g.drawString(message, 50, 100)
	g.setColor(COLOR_RED())
	g.fill(circle1)
	g.setColor(COLOR_YELLOW())
	g.fill(circle2)
	g.setColor(COLOR_GREEN())
	g.fill(circle3)
}

var message string = "Click on one of the circles!"
var circle1 IShape = NewEllipse2D(500, 50, 150, 150)
var circle2 IShape = NewEllipse2D(500, 210, 150, 150)
var circle3 IShape = NewEllipse2D(500, 370, 150, 150)

func main() {
	gtk.Init(nil)
	printEventTypes()
	win := createWindow("Traffic Light", 1024, 768)
	area := createDrawingArea()
	win.Add(area)
	win.ShowAll()
	gtk.Main()
}
