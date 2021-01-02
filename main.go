package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
)

func main() {
	myApp := app.New()
	icon, _ := fyne.LoadResourceFromPath("icon.png")
	MainWindow := myApp.NewWindow("dogboard")
	MainWindow.SetIcon(icon)

	// Get sounds and create grid to arrange them in
	config, _ := GetConfig()
	sounds := GetSounds(config)

	if len(sounds) > 0 {
		MakeSoundGrid(sounds, config, MainWindow)
	} else {
		MakeDefaultGrid(config, MainWindow)
	}
	MainWindow.ShowAndRun()
}
