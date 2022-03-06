package main

import (
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

	// Set Actions
	// Todo: Set Signals

	//helpers.SetSignals(gladeBuilder)
	state := models.NewState(gladeBuilder)
	state.Update()

	//Start
	//_ = glib.TimeoutAdd(1000, state.Update)
	_ = glib.TimeoutSecondsAdd(1, state.Update)

	// Run
	gtk.Main()
}

// destroy is the triggered handler when closing/destroying the gui window
func destroy() {
	gtk.MainQuit()
}