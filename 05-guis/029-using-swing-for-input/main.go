package main

import (
	"fmt"
	"log"
	"strconv"

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

func createGrid() *gtk.Grid {
	grid, err := gtk.GridNew()
	checkError(err, "grid")
	return grid
}

func createLabel(text string) *gtk.Label {
	label, err := gtk.LabelNew(text)
	checkError(err, "label")
	return label
}

func createEntry() *gtk.Entry {
	entry, err := gtk.EntryNew()
	checkError(err, "entry")
	return entry
}

func createButton(text string, onclick func()) *gtk.Button {
	button, err := gtk.ButtonNewWithLabel(text)
	checkError(err, "button")
	button.Connect("clicked", onclick)
	return button
}

func getText(entry *gtk.Entry) string {
	text, err := entry.GetText()
	if err != nil {
		log.Fatal("Unable to retrieve text:", err)
	}
	return text
}

func getInt(entry *gtk.Entry) int {
	text, err := entry.GetText()
	if err != nil {
		log.Fatal("Unable to retrieve text:", err)
	}
	i, err := strconv.Atoi(text)
	if err != nil {
		log.Fatal("Unable to convert text to integer:", err)
	}
	return i
}

func main() {
	gtk.Init(nil)
	win := createWindow()
	win.SetTitle("Simple Example")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	grid := createGrid()
	nameLabel := createLabel("What is your name?")
	nameEntry := createEntry()
	ageLabel := createLabel("How old are you?")
	ageEntry := createEntry()
	button := createButton("Greet me!", func() {
		name := getText(nameEntry)
		age := getInt(ageEntry)
		fmt.Printf("Hello, %s.\n", name)
		fmt.Printf("Next year, you'll be %d\n", age + 1)
		win.Close()
	})
	grid.Attach(nameLabel, 0, 0, 1, 1)
	grid.Attach(nameEntry, 1, 0, 1, 1)
	grid.Attach(ageLabel, 0, 1, 1, 1)
	grid.Attach(ageEntry, 1, 1, 1, 1)
	grid.Attach(button, 1, 2, 1, 1)
	win.Add(grid)
	win.SetDefaultSize(800, 600)
	win.ShowAll()
	gtk.Main()
}
