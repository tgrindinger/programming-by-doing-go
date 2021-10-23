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

func drawLine(cr *cairo.Context, x1, y1, x2, y2 float64) {
	cr.MoveTo(x1, y1)
	cr.LineTo(x2, y2)
	cr.Stroke()
}

func drawString(cr *cairo.Context, msg string, x, y float64) {
	cr.MoveTo(x, y)
	cr.ShowText(msg)
}

func paint(da *gtk.DrawingArea, cr *cairo.Context) {
	cr.SetSourceRGB(128 / 255.0, 0, 128 / 255.0)
	tri := NewPolygon(cr)
	tri.addPoint(500, 100)
	tri.addPoint(800, 600)
	tri.addPoint(200, 500)
	tri.fill()
}

func main() {
	gtk.Init(nil)
	win := createWindow("A Boring Triangle", 1024, 768)
	area := createDrawingArea()
	win.Add(area)
	win.ShowAll()
	gtk.Main()
}
