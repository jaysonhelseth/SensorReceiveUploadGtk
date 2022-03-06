package models

import (
	"SensorReceiveUploadGtk/ui"
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"time"
)

type State struct {
	time *gtk.Label
}

func NewState(gladeBuilder *gtk.Builder) *State {
	state := &State{}
	state.time = ui.GetLabel(gladeBuilder, "timeId")

	return state
}

func (s *State) Update() bool {
	s.time.SetText(fmt.Sprintf("Time: %s",
		time.Now().Format("2006-01-02 15:04:05")),
	)

	// Return true to keep the timer function happy
	return true
}
