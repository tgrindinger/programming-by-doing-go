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
	win.Connect("key-press-event", daEvent)
	win.SetDefaultSize(width, height)
	return win
}

func createDrawingArea() *gtk.DrawingArea {
	area, err := gtk.DrawingAreaNew()
	checkError(err, "drawing area")
	area.Connect("draw", paint)
	return area
}

func processMove(key uint) {
	amount := 10.0
	moveMap := map[uint]func() {
		gdk.KeyvalFromName("Left"):  func() { x -= amount },
		gdk.KeyvalFromName("Right"): func() { x += amount },
		gdk.KeyvalFromName("Up"):    func() { y -= amount },
		gdk.KeyvalFromName("Down"):  func() { y += amount },
	}
	move, found := moveMap[key]
	if found {
		move()
	}
}

func processColorChange(key uint) {
	colorMap := map[uint]Colors {
		gdk.KeyvalFromName("F1"): COLOR_RED,
		gdk.KeyvalFromName("F2"): COLOR_GREEN,
		gdk.KeyvalFromName("F3"): COLOR_BLUE,
		gdk.KeyvalFromName("F5"): COLOR_BLACK,
	}
	color, found := colorMap[key]
	if found {
		cur = color
	}
}

func daEvent(win *gtk.Window, event *gdk.Event) bool {
	keyEvent := gdk.EventKeyNewFromEvent(event)
	if keyEvent.Type() == gdk.EVENT_KEY_PRESS {
		processMove(keyEvent.KeyVal())
		processColorChange(keyEvent.KeyVal())
		win.QueueDraw()
		return false
	}
	return true
}

var x float64 = 500
var y float64 = 400
var cur Colors = COLOR_BLACK

func paint(da *gtk.DrawingArea, cr *cairo.Context) {
	g := NewGraphics(cr)
	g.setColor(cur)
	g.fillOval(x, y, 25, 25)
}

func main() {
	gtk.Init(nil)
	win := createWindow("Use the arrow keys!", 1024, 768)
	area := createDrawingArea()
	win.Add(area)
	win.ShowAll()
	gtk.Main()
}
