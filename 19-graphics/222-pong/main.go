package main

import (
	"fmt"
	"log"
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
	win.Connect("key-release-event", daEvent)
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
	delta := 5.0
	if keyEvent.Type() == gdk.EVENT_KEY_PRESS {
		switch keyEvent.KeyVal() {
		case gdk.KeyvalFromName("w"):
			p1dy = -delta
		case gdk.KeyvalFromName("s"):
			p1dy = delta
		case gdk.KeyvalFromName("Up"):
			p2dy = -delta
		case gdk.KeyvalFromName("Down"):
			p2dy = delta
		}
		return false
	}
	if keyEvent.Type() == gdk.EVENT_KEY_RELEASE {
		switch keyEvent.KeyVal() {
		case gdk.KeyvalFromName("w"):
			p1dy = 0
		case gdk.KeyvalFromName("s"):
			p1dy = 0
		case gdk.KeyvalFromName("Up"):
			p2dy = 0
		case gdk.KeyvalFromName("Down"):
			p2dy = 0
		}
		return false
	}
	return true
}

var delta *Point = &Point{-5, 5}
var ball IShape = NewEllipse(500, 350, 20, 20)
var paddle1 IShape = NewRectangle(50, 250, 20, 200)
var paddle2 IShape = NewRectangle(930, 250, 20, 200)
var p1dy float64 = 0
var p2dy float64 = 0
var score1 int = 0
var score2 int = 0

func paint(da *gtk.DrawingArea, cr *cairo.Context) {
	g := NewGraphics(cr)
	g.setColor(COLOR_BLACK)
	g.fill(ball)
	g.setColor(COLOR_BLUE)
	g.fill(paddle1)
	g.fill(paddle2)
	g.setColor(COLOR_BLACK)
	cr.SetFontSize(24)
	g.drawString(fmt.Sprintf("%d", score1), 10, 50)
	g.drawString(fmt.Sprintf("%d", score2), 900, 50)
}

func doStuff() {
	ball.move(delta.x, delta.y)
	paddle1.move(0, p1dy)
	paddle2.move(0, p2dy)
	if ball.pos().y < 0 || ball.pos().y + 20 > 700 {
		delta.y = -delta.y
	}
	if ball.intersects(paddle1) || ball.intersects(paddle2) {
		delta.x = -delta.x
	}
	if ball.pos().x > 1000 {
		ball.moveTo(500, 350)
		delta = &Point{-5, 5}
		score1++
	}
	if ball.pos().x <= 0 {
		ball.moveTo(500, 350)
		delta = &Point{-5, 5}
		score2++
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
