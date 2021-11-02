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
		fmt.Printf("You clicked at (%.0f,%.0f)\n", x, y)
		curMessage = "The square is green."
		curColor = COLOR_GREEN()
		da.QueueDraw()
		return false
	}
	return true
}

func paint(da *gtk.DrawingArea, cr *cairo.Context) {
	g := NewGraphics(cr)
	g.setColor(COLOR_BLACK())
	cr.SelectFontFace("Calibri", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
	cr.SetFontSize(24)
	g.drawString(curMessage, 400, 100)
	g.setColor(curColor)
	g.fillRect(250, 100, 100, 100)
}

var curMessage = "The square is red."
var curColor = COLOR_RED()

func main() {
	gtk.Init(nil)
	printEventTypes()
	win := createWindow("MouseDemo", 1024, 768)
	area := createDrawingArea()
	win.Add(area)
	win.ShowAll()
	gtk.Main()
}
