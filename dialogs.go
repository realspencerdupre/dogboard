package main

import (
	"path/filepath"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
)

func SetSoundsFolder(config AppSettings, window fyne.Window) func() {
	return func() {
		window.Resize(fyne.NewSize(640, 480))
		dialog.ShowFolderOpen(func(list fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, window)
				return
			}
			if list == nil {
				return
			}

			config.SoundsPath = list.String()[7:]
			configFile := filepath.Join(ConfigPath, "settings.json")
			SaveConfig(configFile, config)
			sounds := GetSounds(config)
			if len(sounds) > 0 {
				MakeSoundGrid(sounds, config, window)
			}
		}, window)
	}
}
