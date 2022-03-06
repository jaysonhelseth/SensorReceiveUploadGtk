package ui

import "github.com/gotk3/gotk3/gtk"

func GetLabel(gladeBuilder *gtk.Builder, name string) *gtk.Label {
	obj, err := gladeBuilder.GetObject(name)
	if err != nil {
		return nil
	}

	label, ok := obj.(*gtk.Label)
	if !ok {
		return nil
	}

	return label
}
