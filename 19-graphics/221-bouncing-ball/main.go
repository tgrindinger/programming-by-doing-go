package main

import (
	"log"
	"math/rand"
	"time"

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

func daEvent(win *gtk.Window, event *gdk.Event) bool {
	keyEvent := gdk.EventKeyNewFromEvent(event)
	if keyEvent.Type() == gdk.EVENT_KEY_PRESS {
		if keyEvent.KeyVal() == gdk.KeyvalFromName("space") {
			flashing = !flashing
		}
		return false
	}
	return true
}

var x float64 = 500
var y float64 = 350
var dx float64 = 5
var dy float64 = 5
var flashing bool = false
var cur *Color = &Color{0, 0, 0}

func paint(da *gtk.DrawingArea, cr *cairo.Context) {
	g := NewGraphics(cr)
	g.setColorRGB(cur)
	g.fillOval(x, y, 20, 20)
}

func doStuff() {
	x += dx
	y += dy
	if x < 0 || x + 20 > 1000 {
		dx = -dx
	}
	if y < 0 || y + 20 > 700 {
		dy = -dy
	}
	if flashing {
		r := float64(rand.Float64())
		g := float64(rand.Float64())
		b := float64(rand.Float64())
		cur = &Color{r, g, b}
	}
}

func startTicker(win *gtk.Window) {
	ticker := time.NewTicker(time.Millisecond * 10)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				doStuff()
				win.QueueDraw()
			}
		}
	}()
}

func main() {
	gtk.Init(nil)
	win := createWindow("BouncingBall", 1010, 735)
	area := createDrawingArea()
	win.Add(area)
	win.ShowAll()
	startTicker(win)
	gtk.Main()
}
