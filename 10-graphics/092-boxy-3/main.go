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

func drawBox(cr *cairo.Context, r, g, b, x, y, w, h float64) {
	margin := 10.0
	cr.SetSourceRGB(r, g, b)
	fillRect(cr, x, y, w, h)
	cr.SetSourceRGB(1, 1, 1)
	fillRect(cr, x + margin, y + margin, w - margin * 2, h - margin * 2)
}

func paint(da *gtk.DrawingArea, cr *cairo.Context) {
	drawBox(cr, 1, 1, 0, 200, 300, 50, 30)
	drawBox(cr, 0, 1, 0, 400, 220, 30, 150)
	drawBox(cr, 0, 0, 0, 350, 50, 50, 100)
	drawBox(cr, 0, 0, 1, 450, 40, 250, 75)
	drawBox(cr, 1, 0, 0, 600, 300, 100, 80)
	drawBox(cr, 1, 0, 1, 150, 375, 100, 50)
	drawBox(cr, 0.8, 0.1, 0.1, 50, 500, 600, 30)
	drawBox(cr, 0.1, 0.1, 0.8, 100, 100, 100, 150)
}

func main() {
	gtk.Init(nil)
	win := createWindow("Boxy3 - write and use function - boxes different sizes", 800, 600)
	area := createDrawingArea()
	win.Add(area)
	win.ShowAll()
	gtk.Main()
}