package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/skratchdot/open-golang/open"
)

func MakeSoundGrid(sounds []Sound, config AppSettings, window fyne.Window) int {
	soundGrid := fyne.NewContainerWithLayout(layout.NewGridLayout(config.GridWidth))

	for _, sound := range sounds {
		soundGrid.AddObject(NewSoundButton(sound))
	}

	window.SetContent(soundGrid)

	gridHeight := len(sounds) / config.GridWidth
	if (len(sounds) % config.GridWidth) > 0 {
		gridHeight = gridHeight + 1
	}
	window.Resize(fyne.NewSize(config.IconSize*config.GridWidth, config.IconSize*gridHeight))
	return gridHeight
}

func MakeDefaultGrid(config AppSettings, window fyne.Window) {
	window.Resize(fyne.NewSize(640, 480))
	container := fyne.NewContainerWithLayout(layout.NewGridLayout(1))
	// Button to open the default sounds folder
	container.AddObject(widget.NewButton("Open default Sounds Folder", func() {
		open.Run(config.SoundsPath)
	}))

	// Button to set custom sounds folder
	container.AddObject(widget.NewButton("Set custom Sounds Folder", func() {
		fmt.Println("Setting folder")
		dialog.ShowFolderOpen(func(list fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, window)
				return
			}
			if list == nil {
				return
			}

			children, err := list.List()
			if err != nil {
				dialog.ShowError(err, window)
				return
			}
			out := fmt.Sprintf("Folder %s (%d children):\n%s", list.Name(), len(children), list.String())
			dialog.ShowInformation("Folder Open", out, window)
		}, window)
	}))

	// Button to refresh this list
	container.AddObject(widget.NewButton("Refresh list", func() {
		sounds := GetSounds(config)
		gridHeight := MakeSoundGrid(sounds, config, window)
		fmt.Printf("w %d h %d", config.IconSize*config.GridWidth, config.IconSize*gridHeight)
		window.Resize(fyne.NewSize(config.IconSize*config.GridWidth, config.IconSize*gridHeight))
	}))
	window.SetContent(container)
}

func main() {
	myApp := app.New()
	icon, _ := fyne.LoadResourceFromPath("icon.png")
	MainWindow := myApp.NewWindow("dogboard")
	MainWindow.SetIcon(icon)

	// Get sounds and create grid to arrange them in
	config, _ := GetConfig()
	sounds := GetSounds(config)

	if len(sounds) > 0 {
		gridHeight := MakeSoundGrid(sounds, config, MainWindow)
		MainWindow.Resize(fyne.NewSize(config.IconSize*config.GridWidth, config.IconSize*gridHeight))
	} else {
		MakeDefaultGrid(config, MainWindow)
	}
	MainWindow.ShowAndRun()
}
