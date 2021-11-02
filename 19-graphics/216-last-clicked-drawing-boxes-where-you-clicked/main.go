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
		lastX, lastY = eventClick.MotionVal()
		message = fmt.Sprintf("You clicked at (%.0f,%.0f)", lastX, lastY)
		da.QueueDraw()
		return false
	}
	return true
}

func paint(da *gtk.DrawingArea, cr *cairo.Context) {
	g := NewGraphics(cr)
	g.setColor(COLOR_RED)
	g.fillRect(lastX, lastY, 5, 5)
	g.setColor(COLOR_BLACK)
	g.drawString(message, 50, 100)
}

var message string = "Click someplace!"
var lastX float64 = 0
var lastY float64 = 0

func main() {
	gtk.Init(nil)
	printEventTypes()
	win := createWindow("Last Clicked", 1024, 768)
	area := createDrawingArea()
	win.Add(area)
	win.ShowAll()
	gtk.Main()
}
