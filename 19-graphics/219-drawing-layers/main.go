package main

import (
	"log"

	"github.com/gotk3/gotk3/cairo"
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
	area.Connect("draw", paint)
	return area
}

func paint(da *gtk.DrawingArea, cr *cairo.Context) {
	g := NewGraphics(cr)
	gBuf := NewGraphics(surfaceContext)
	gBuf.setColor(COLOR_BLUE)
	gBuf.fillRect(20, 20, 100, 200)
	gBuf.setColor(COLOR_WHITE)
	gBuf.fillRect(120, 20, 100, 200)
	gBuf.setColor(COLOR_RED)
	gBuf.fillRect(220, 20, 100, 200)
	g.drawImage(surface, 0, 0)
}

var surface *cairo.Surface
var surfaceContext *cairo.Context

func setupImageBuffer(width, height int) {
	surface = cairo.CreateImageSurface(cairo.FORMAT_RGB24, width, height)
	surfaceContext = cairo.Create(surface)
}

func main() {
	gtk.Init(nil)
	width := 1024
	height := 768
	setupImageBuffer(width, height)
	win := createWindow("Draw Layers", width, height)
	area := createDrawingArea()
	win.Add(area)
	win.ShowAll()
	gtk.Main()
}
