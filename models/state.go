package models

import (
	"SensorReceiveUploadGtk/ui"
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"time"
)

type State struct {
	time     *gtk.Label
	pool     *gtk.Label
	air      *gtk.Label
	humidity *gtk.Label
	updated  *gtk.Label
}

func NewState(gladeBuilder *gtk.Builder) *State {
	state := &State{}
	state.time = ui.GetLabel(gladeBuilder, "timeValue")
	state.pool = ui.GetLabel(gladeBuilder, "poolValue")
	state.air = ui.GetLabel(gladeBuilder, "airValue")
	state.humidity = ui.GetLabel(gladeBuilder, "humidityValue")
	state.updated = ui.GetLabel(gladeBuilder, "updatedValue")

	return state
}

func (s *State) Update() bool {
	s.time.SetText(fmt.Sprintf("Time: %s",
		time.Now().Format("2006-01-02 15:04:05")))

	s.pool.SetText("00.0")
	s.air.SetText("00.0")
	s.humidity.SetText("00.0")
	s.updated.SetText(fmt.Sprintf("Time: %s",
		time.Now().Format("2006-01-02 15:04:05")))

	// Return true to keep the timer function happy
	return true
}
