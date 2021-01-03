package main

import (
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/skratchdot/open-golang/open"
)

func MakeSoundGrid(sounds []Sound, config AppSettings, window fyne.Window) {
	if len(sounds) == 0 {
		MakeDefaultGrid(config, window)
		return
	}
	soundGrid := fyne.NewContainerWithLayout(layout.NewGridLayout(config.GridWidth))

	for _, sound := range sounds {
		soundObj, err := NewSoundButton(sound)
		if err != nil {
			log.Fatal(err)
		} else {
			soundGrid.AddObject(soundObj)
		}
	}

	window.SetContent(soundGrid)

	gridHeight := len(sounds) / config.GridWidth
	if (len(sounds) % config.GridWidth) > 0 {
		gridHeight = gridHeight + 1
	}
	window.Resize(fyne.NewSize(config.IconSize*config.GridWidth, config.IconSize*gridHeight))
	window.SetFixedSize(true)
}

func MakeDefaultGrid(config AppSettings, window fyne.Window) {
	window.Resize(fyne.NewSize(640, 480))
	container := fyne.NewContainerWithLayout(layout.NewGridLayout(1))
	// Button to open the default sounds folder
	container.AddObject(widget.NewButton("Open default Sounds Folder", func() {
		open.Run(config.SoundsPath)
	}))

	// Button to set custom sounds folder
	container.AddObject(widget.NewButton("Set custom Sounds Folder", SetSoundsFolder(config, window)))

	// Button to refresh this list
	container.AddObject(widget.NewButton("Refresh list", func() {
		sounds := GetSounds(config)
		if len(sounds) > 0 {
			MakeSoundGrid(sounds, config, window)
		}
	}))
	window.SetContent(container)
}
