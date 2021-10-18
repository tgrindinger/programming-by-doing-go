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

func createWindow() *gtk.Window {
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	checkError(err, "window")
	return win
}

func createDrawingArea() *gtk.DrawingArea {
	area, err := gtk.DrawingAreaNew()
	checkError(err, "drawing area")
	return area
}

func main() {
	gtk.Init(nil)
	win := createWindow()
	win.SetTitle("613 rocks!")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	win.SetDefaultSize(300, 200)
	win.Move(100, 200)
	area := createDrawingArea()
	area.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
		cr.SetSourceRGB(0, 0, 0)
		cr.MoveTo(75, 100)
		cr.ShowText("Hi.")
	})
	win.Add(area)
	win.ShowAll()
	gtk.Main()
}
