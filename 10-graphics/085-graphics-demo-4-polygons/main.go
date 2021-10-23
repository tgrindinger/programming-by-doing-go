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
	cr.SetSourceRGB(0, 0, 0)
	drawString(cr, "Hey, a triangle!", 50, 50)

	cr.SetSourceRGB(0, 0, 1)
	tri := NewPolygon(cr)
	tri.addPoint(100, 100)
	tri.addPoint(100, 300)
	tri.addPoint(200, 300)
	tri.fill()

	cr.SetSourceRGB(0, 1, 0)
	pent := NewPolygon(cr)
	pent.addPoint(450, 200)
	pent.addPoint(500, 250)
	pent.addPoint(475, 350)
	pent.addPoint(425, 350)
	pent.addPoint(400, 250)
	pent.fill()

	cr.SetSourceRGB(0, 0, 0)
	hex := NewPolygon(cr)
	radius := 100.0
	xCenter := 200.0
	yCenter := 500.0
	for ang := 0.0; ang < 2 * math.Pi; ang = ang + (2 * math.Pi) / 6.0 {
		xDelta := radius * math.Cos(ang)
		yDelta := -radius * math.Sin(ang)
		hex.addPoint(xCenter + xDelta, yCenter + yDelta)
	}
	hex.fill()

	cr.SetSourceRGB(1, 0, 0)
	trap := NewPolygon(cr)
	trap.addPoint(400, 400)
	trap.addPoint(600, 400)
	trap.addPoint(500, 450)
	trap.addPoint(300, 450)
	trap.fill()
}

func main() {
	gtk.Init(nil)
	win := createWindow("Polygon Demo", 1024, 768)
	area := createDrawingArea()
	win.Add(area)
	win.ShowAll()
	gtk.Main()
}
