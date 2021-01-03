package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"github.com/skratchdot/open-golang/open"
)

func main() {
	myApp := app.New()
	icon, _ := fyne.LoadResourceFromPath("icon.png")
	MainWindow := myApp.NewWindow("dogboard")
	MainWindow.SetIcon(icon)

	// Get sounds and create grid to arrange them in
	config, _ := GetConfig()
	sounds := GetSounds(config)

	// Create menu
	gridItem := fyne.NewMenuItem("Set grid width", nil)
	folderItem := fyne.NewMenuItem("Open sounds folder", func() {
		open.Run(config.SoundsPath)
	})
	soundsItem := fyne.NewMenuItem("Set sounds folder", SetSoundsFolder(config, MainWindow))
	mainMenu := fyne.NewMainMenu(
		fyne.NewMenu(
			"File",
			gridItem,
			fyne.NewMenuItemSeparator(),
			folderItem,
			fyne.NewMenuItemSeparator(),
			soundsItem,
		),
	)
	MainWindow.SetMainMenu(mainMenu)

	if len(sounds) > 0 {
		MakeSoundGrid(sounds, config, MainWindow)
	} else {
		MakeDefaultGrid(config, MainWindow)
	}
	MainWindow.ShowAndRun()
}
