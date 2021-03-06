package main

import (
	"SensorReceiveUploadGtk/io"
	"SensorReceiveUploadGtk/models"
	"SensorReceiveUploadGtk/ui"
	"embed"
	"github.com/gotk3/gotk3/glib"
	"os"

	"flag"
	"github.com/gotk3/gotk3/gtk"
)

//go:embed assets/*
var assets embed.FS

func main() {
	var isFullscreen bool
	flag.BoolVar(&isFullscreen, "f", false, "Set isFullscreen, defaults to false")
	flag.Parse()

	gtk.Init(&os.Args)

	gladeBuilder := ui.GladeBuilder(assets)
	window := ui.WindowBuilder(gladeBuilder, isFullscreen)
	window.Connect("destroy", destroy)
	window.ShowAll()
	ui.SetStyle(assets, window)

	// Channel to track sensor changes.
	c := make(chan models.Sensors)
	state := models.NewState(gladeBuilder, c)
	state.Update()

	go io.ReceiveFromFakeSensors(c)

	// GTK event timer that runs every second.
	_ = glib.TimeoutSecondsAdd(1, state.Update)

	// Run
	gtk.Main()
	defer close(c)
}

// destroy is the triggered handler when closing/destroying the gui window
func destroy() {
	gtk.MainQuit()
}
