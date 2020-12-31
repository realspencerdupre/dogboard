package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
)

var ButtonSize = 104
var GridWidth = 2

func main() {
	myApp := app.New()
	icon, _ := fyne.LoadResourceFromPath("icon.png")
	mainWindow := myApp.NewWindow("dogboard")
	mainWindow.SetIcon(icon)

	// Get sounds and create grid to arrange them in
	sounds := GetSounds()
	soundGrid := fyne.NewContainerWithLayout(layout.NewGridLayout(GridWidth))

	for _, sound := range sounds {
		soundGrid.AddObject(NewSoundButton(sound))
	}

	mainWindow.SetContent(soundGrid)

	gridHeight := len(sounds) / GridWidth
	if (len(sounds) % GridWidth) > 0 {
		gridHeight = gridHeight + 1
	}
	mainWindow.Resize(fyne.NewSize(ButtonSize*GridWidth, ButtonSize*gridHeight))
	mainWindow.ShowAndRun()
}
