package main

import "fmt"
import "os"

import "github.com/mattn/go-gtk/gtk"
import "github.com/mewmew/gtkmap"

func main() {
	gtk.Init(&os.Args)

	// Create a window.
	win := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	win.Connect("destroy", gtk.MainQuit)

	// Create a map widget which uses Google Maps as source for the map tiles.
	source := gtkmap.SourceGoogleMaps
	m, err := gtkmap.NewMapWithSource(source)
	if err != nil {
		// Fall back to using OpenStreetMap if Google Maps could not be used as
		// source.
		m = gtkmap.NewMap()
	}
	fmt.Println("Map tile representations from:", m.Source())
	m.SetSizeRequest(640, 480)
	win.Add(m)

	// Center the map on Guangzhou and set zoom level.
	lat := 23.110262
	long := 113.319374
	zoom := 6
	m.SetCenterAndZoom(lat, long, zoom)

	win.ShowAll()
	gtk.Main()
}