package main

import (
	"fmt"
	"log"
	"math"

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

func doRect(cr *cairo.Context, x, y, w, h float64) {
	cr.Rectangle(x, y, w, h)
}

func drawRect(cr *cairo.Context, x, y, w, h float64) {
	doRect(cr, x, y, w, h)
	cr.Stroke()
}

func fillRect(cr *cairo.Context, x, y, w, h float64) {
	doRect(cr, x, y, w, h)
	cr.Fill()
}

func doOval(cr *cairo.Context, x, y, w, h float64) {
	cr.Save()
	radX := w / 2
	radY := h / 2
	cr.Translate(x + radX, y + radY)
	minRad := radX
	if radX > radY {
		minRad = radY
	}
	cr.Scale(radX / minRad, radY / minRad)
	cr.Arc(0, 0, minRad, 0, 2 * math.Pi)
	cr.Restore()
}

func drawOval(cr *cairo.Context, x, y, w, h float64) {
	doOval(cr, x, y, w, h)
	cr.Stroke()
}

func fillOval(cr *cairo.Context, x, y, w, h float64) {
	doOval(cr, x, y, w, h)
	cr.Fill()
}

func main() {
	gtk.Init(nil)
	win := createWindow()
	win.SetTitle("GraphicsDemo1")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	win.SetDefaultSize(800, 600)
	area := createDrawingArea()
	area.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
		cr.SetLineWidth(1)
		cr.SetSourceRGB(0, 1, 0)
		drawRect(cr, 50, 20, 100, 200)
		fillOval(cr, 160, 20, 100, 200)
		cr.SetSourceRGB(0, 0, 1)
		fillRect(cr, 200, 400, 200, 20)
		drawOval(cr, 200, 430, 200, 100)
		cr.SetSourceRGB(1, 0, 0)
		fillRect(cr, 500, 400, 50, 50)
		cr.SetSourceRGB(0, 0, 0)
		cr.MoveTo(500, 100)
		cr.ShowText("Graphics are pretty neat.")
		x := float64(da.GetAllocatedWidth() / 2)
		y := float64(da.GetAllocatedHeight() / 2)
		cr.MoveTo(x, y)
		cr.ShowText(fmt.Sprintf("The first letter of this string is at (%f, %f)", x, y))
	})
	win.Add(area)
	win.ShowAll()
	gtk.Main()
}
