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
	a := NewLine(150, 50, 160, 100)
	b := NewRectangle(100, 150, 200, 60)
	c := NewArc(100, 300, 150, 150, 45, 270, ARC_PIE)
	xpoints := []float64{150, 150, 220}
	ypoints := []float64{500, 600, 600}
	d := NewPolygon(xpoints, ypoints)
	g.draw(a)
	g.draw(b)
	g.draw(c)
	g.draw(d)
	trans := NewAffineTransform()
	trans.scale(0.75, 0.75)
	trans.translate(400, 0)
	a2 := trans.createTransformedShape(a)
	b2 := trans.createTransformedShape(b)
	c2 := trans.createTransformedShape(c)
	d2 := trans.createTransformedShape(d)
	g.setColor(COLOR_GREEN)
	g.draw(a2)
	g.fill(b2)
	g.fill(c2)
	g.fill(d2)
}

func main() {
	gtk.Init(nil)
	printEventTypes()
	win := createWindow("Shapes Demo", 1024, 768)
	area := createDrawingArea()
	win.Add(area)
	win.ShowAll()
	gtk.Main()
}
