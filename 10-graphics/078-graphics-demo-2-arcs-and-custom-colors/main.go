package main

import (
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

func doArcWithAngles(cr *cairo.Context, x, y, w, h, a1, a2 float64) {
	cr.Save()
	radX := w / 2
	radY := h / 2
	cr.Translate(x + radX, y + radY)
	minRad := radX
	if radX > radY {
		minRad = radY
	}
	cr.Scale(radX / minRad, radY / minRad)
	rad1 := a1 * (math.Pi / 180)
	rad2 := a2 * (math.Pi / 180)
	cr.Arc(0, 0, minRad, rad1, rad2)
	cr.Restore()
}

func doOval(cr *cairo.Context, x, y, w, h float64) {
	doArcWithAngles(cr, x, y, w, h, 0, 360)
}

func drawOval(cr *cairo.Context, x, y, w, h float64) {
	doOval(cr, x, y, w, h)
	cr.Stroke()
}

func fillOval(cr *cairo.Context, x, y, w, h float64) {
	doOval(cr, x, y, w, h)
	cr.Fill()
}

func drawArc(cr *cairo.Context, x, y, w, h, a1, a2 float64) {
	doArcWithAngles(cr, x, y, w, h, a1, a2)
	cr.Stroke()
}

func fillArc(cr *cairo.Context, x, y, w, h, a1, a2 float64) {
	cr.NewPath()
	cr.MoveTo(x + w / 2, y + h / 2)
	doArcWithAngles(cr, x, y, w, h, a1, a2)
	cr.ClosePath()
	cr.Fill()
}

func paint(da *gtk.DrawingArea, cr *cairo.Context) {
	cr.SetSourceRGB(0, 0, 0)
	drawRect(cr, 50, 20, 100, 200)
	drawOval(cr, 160, 20, 100, 200)
	drawArc(cr, 270, 20, 100, 200, 0, 270)
	drawArc(cr, 50, 250, 150, 150, 90, 180)
	fillArc(cr, 210, 280, 150, 150, 45, 90)

	cr.SetSourceRGB(1, 1, 0)
	fillArc(cr, 150, 400, 150, 150, 45, 315) // chomp

	cr.SetSourceRGB(1, 105 / 255.0, 180 / 255.0)
	fillArc(cr, 325, 425, 100, 100, 225, 135) // ms. chomp

	cr.SetSourceRGB(230 / 255.0, 92 / 255.0, 0)
	fillOval(cr, 500, 50, 150, 150)

	cr.SetSourceRGB(238 / 255.0, 238 / 255.0, 238 / 255.0)
	fillOval(cr, 550, 100, 50, 50)

	cr.SetSourceRGB(0.2, 0, 1)
	fillOval(cr, 500, 210, 150, 150)

	cr.SetSourceRGB(0, 1, 0)
	fillOval(cr, 500, 370, 150, 150)
}

func main() {
	gtk.Init(nil)
	win := createWindow("GraphicsDemo2: Arcs and Colors", 800, 600)
	area := createDrawingArea()
	win.Add(area)
	win.ShowAll()
	gtk.Main()
}
