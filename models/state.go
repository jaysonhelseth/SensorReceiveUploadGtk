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
	myChan   chan Sensors
}

func NewState(gladeBuilder *gtk.Builder, c chan Sensors) *State {
	state := &State{}
	state.time = ui.GetLabel(gladeBuilder, "timeValue")
	state.pool = ui.GetLabel(gladeBuilder, "poolValue")
	state.air = ui.GetLabel(gladeBuilder, "airValue")
	state.humidity = ui.GetLabel(gladeBuilder, "humidityValue")
	state.updated = ui.GetLabel(gladeBuilder, "updatedValue")
	state.myChan = c

	return state
}

func (s *State) Update() bool {
	s.time.SetText(fmt.Sprintf("Time: %s",
		time.Now().Format("2006-01-02 15:04:05")))

	select {
	case update := <-s.myChan:
		s.pool.SetText(fmt.Sprintf("%.1f", update.Pool))
		s.air.SetText(fmt.Sprintf("%.1f", update.Air))
		s.humidity.SetText(fmt.Sprintf("%.1f", update.Humidity))
		s.updated.SetText(fmt.Sprintf("Updated: %s",
			time.Now().Format("2006-01-02 15:04:05")))
	default:
		//Do nothing
	}

	// Return true to keep the timer function happy
	return true
}
