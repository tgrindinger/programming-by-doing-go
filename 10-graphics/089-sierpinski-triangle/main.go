package main

import (
	"log"
	"math"
	"math/rand"

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
	x1 := 512.0
	y1 := 109.0
	x2 := 146.0
	y2 := 654.0
	x3 := 876.0
	y3 := 654.0
	x := 512.0
	y := 382.0
	for i := 0; i < 50000; i++ {
		drawLine(cr, x, y, x + 1, y + 1)
		p := 1 + rand.Intn(3)
		var dx, dy float64
		if p == 1 {
			dx = x - x1
			dy = y - y1
		} else if p == 2 {
			dx = x - x2
			dy = y - y2
		} else {
			dx = x - x3
			dy = y - y3
		}
		x = x - dx / 2
		y = y - dy / 2
	}
	drawString(cr, "Sierpinski Triangle", 462, 484)
}

func main() {
	gtk.Init(nil)
	win := createWindow("Sierpinski Triangle", 1024, 768)
	area := createDrawingArea()
	win.Add(area)
	win.ShowAll()
	gtk.Main()
}
