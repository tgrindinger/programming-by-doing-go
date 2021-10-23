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
	cr.SetSourceRGB(0.8, 0.8, 0.8)
	fillRect(cr, 0, 0, float64(da.GetAllocatedWidth()), float64(da.GetAllocatedHeight()))

	cr.SetSourceRGB(0, 1, 0)
	drawLine(cr, 10, 100, 400, 100)
	cr.SetSourceRGB(1, 0, 0)
	drawLine(cr, 10, 100, 35, 125)
	cr.SetSourceRGB(0, 0, 1)
	drawLine(cr, 50, 150, 100, 180)
	cr.SetSourceRGB(1, 0, 1)
	drawLine(cr, 100, 350, 300, 230)
	cr.SetSourceRGB(1, 0, 0)
	drawLine(cr, 200, 290, 230, 340)

	fontSize := 10.0
	defaultFamily := "sans-serif"
	cr.SetSourceRGB(0, 0, 0)
	drawString(cr, "Graphics are pretty neat.", 500, 100 + fontSize)
	cr.SelectFontFace("Consolas", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_NORMAL)
	fontSize = 36.0
	cr.SetFontSize(fontSize)
	drawString(cr, "Yes, they are.", 400, 200 + fontSize)

	cr.SetSourceRGB(1, 1, 1)
	cr.SelectFontFace("Times New Roman", cairo.FONT_SLANT_ITALIC, cairo.FONT_WEIGHT_BOLD)
	fontSize = 60.0
	cr.SetFontSize(fontSize)
	drawString(cr, "Leander Lions", 300, 350 + fontSize)

	cr.SetSourceRGB(0, 0, 1)
	drawString(cr, "Leander Lions", 290, 360 + fontSize)

	cr.SetSourceRGB(0, 0, 0)
	cr.SelectFontFace(defaultFamily, cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_NORMAL)
	fontSize = 10.0
	cr.SetFontSize(fontSize)

	x := 400.0
	y := 490.0
	drawRect(cr, x, y, 150, 20)
	drawString(cr, "Where is the starting point?", x, y + fontSize)
	cr.SetSourceRGB(1, 0, 0)
	drawLine(cr, x, y, x + 1, y + 1)
}

func main() {
	gtk.Init(nil)
	win := createWindow("GraphicsDemo2: Arcs and Colors", 800, 600)
	area := createDrawingArea()
	win.Add(area)
	win.ShowAll()
	gtk.Main()
}
