package main

import (
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
	win.Connect("button-release-event", daEvent)
	win.SetDefaultSize(width, height)
	return win
}

func createDrawingArea() *gtk.DrawingArea {
	area, err := gtk.DrawingAreaNew()
	checkError(err, "drawing area")
	area.Connect("draw", paint)
	return area
}

func daEvent(win *gtk.Window, event *gdk.Event) bool {
	keyEvent := gdk.EventButtonNewFromEvent(event)
	if keyEvent.Type() == gdk.EVENT_BUTTON_RELEASE {
		x, y := keyEvent.MotionVal()
		rev.mouseReleased(x, y)
		win.QueueDraw()
	}
	return true
}

func paint(da *gtk.DrawingArea, cr *cairo.Context) {
	g := NewGraphics(cr)
	rev.paint(g)
}

var rev *Reversi = NewReversi()

func main() {
	gtk.Init(nil)
	win := createWindow("Reversi", 800, 900)
	area := createDrawingArea()
	win.Add(area)
	win.ShowAll()
	gtk.Main()
}
