package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
)

type diagonal struct {
}

func (d *diagonal) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := 0, 0
	for _, o := range objects {
		childSize := o.MinSize()

		w += childSize.Width
		h += childSize.Height
	}
	return fyne.NewSize(w, h)
}

func (d *diagonal) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	pos := fyne.NewPos(0, containerSize.Height-d.MinSize(objects).Height)
	for _, o := range objects {
		size := o.MinSize()
		o.Resize(size)
		o.Move(pos)

		pos = pos.Add(fyne.NewPos(size.Width, size.Height))
	}
}

func main() {
	myApp := app.New()
	icon, _ := fyne.LoadResourceFromPath("icon.png")
	mainWindow := myApp.NewWindow("dogboard")
	mainWindow.SetIcon(icon)

	// Get sounds and create grid to arrange them in
	sounds := GetSounds()
	soundGrid := fyne.NewContainerWithLayout(layout.NewGridLayout(4))

	for _, sound := range sounds {
		soundGrid.AddObject(NewSoundButton(sound))
	}

	mainWindow.SetContent(soundGrid)
	mainWindow.Resize(fyne.NewSize(640, 640))
	mainWindow.ShowAndRun()
}
