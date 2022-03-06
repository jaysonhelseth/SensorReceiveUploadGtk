package ui

import (
	"embed"
	"github.com/gotk3/gotk3/gtk"
	"log"
)

// GladeBuilder returns pointer to gtk.builder loaded with glade resource (if resource is given)
func GladeBuilder(assets embed.FS) *gtk.Builder {

	builder, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal(err)
	}

	layout, err := assets.ReadFile("assets/layout.glade")
	if err != nil {
		log.Fatal(err)
	}

	err = builder.AddFromString(string(layout))
	if err != nil {
		log.Fatal(err)
	}

	return builder
}

// WindowBuilder returns gtk.Window object from the glade resource
func WindowBuilder(b *gtk.Builder, isFullscreen bool) *gtk.Window {

	// Look in the glade file for the name.
	obj, err := b.GetObject("topWindow")
	if err != nil {
		log.Fatal(err)
	}

	window, ok := obj.(*gtk.Window)
	if !ok {
		log.Fatal(err)
	}

	if isFullscreen {
		window.Fullscreen()
		window.SetDecorated(false)
	} else {
		window.SetDefaultSize(320, 240)
		window.SetPosition(gtk.WIN_POS_CENTER)
	}

	window.SetTitle("Sensors")
	return window
}

func SetStyle(assets embed.FS, window *gtk.Window) {
	cssProvider, err := gtk.CssProviderNew()
	if err != nil {
		log.Fatal(err)
	}

	css, err := assets.ReadFile("assets/layout.css")
	if err != nil {
		log.Fatal(err)
	}

	err = cssProvider.LoadFromData(string(css))
	if err != nil {
		log.Fatal(err)
	}

	winScreen := window.GetScreen()
	if winScreen == nil {
		log.Fatal(err)
	}

	gtk.AddProviderForScreen(winScreen, cssProvider, gtk.STYLE_PROVIDER_PRIORITY_USER)
}
