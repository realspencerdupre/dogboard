package main

import (
	"log"
	"os"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type SoundButton struct {
	widget.Icon
	buffer *beep.Buffer
}

func NewSoundButton(sound Sound) *SoundButton {
	button := &SoundButton{}
	button.ExtendBaseWidget(button)

	f, err := os.Open(sound.AudioPath)
	if err != nil {
		log.Fatal(err)
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/100))
	button.buffer = beep.NewBuffer(format)
	button.buffer.Append(streamer)
	streamer.Close()
	soundPath := ""

	// Sound has it's own icon
	if sound.IconPath != "" {
		soundPath = sound.IconPath
	} else // Else use the dogboard icon
	{
		soundPath = "icon.png"
	}
	icon, _ := fyne.LoadResourceFromPath(soundPath)
	button.SetResource(icon)

	return button
}

func (t *SoundButton) Tapped(_ *fyne.PointEvent) {
	log.Println("I have been tapped")
	speaker.Play(t.buffer.Streamer(0, t.buffer.Len()))
}

func (t *SoundButton) TappedSecondary(_ *fyne.PointEvent) {
}
