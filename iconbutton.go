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
	Buffer     beep.Buffer
	Sound      Sound
	SampleRate beep.SampleRate
}

var MasterFormat = beep.Format{48000, 2, 2}
var MasterBuffer = beep.NewBuffer(MasterFormat)
var Speaker = speaker.Init(MasterFormat.SampleRate, MasterFormat.SampleRate.N(time.Second/16))

func NewSoundButton(sound Sound) (*SoundButton, error) {
	button := &SoundButton{}
	button.ExtendBaseWidget(button)
	button.Sound = sound

	f, err := os.Open(sound.AudioPath)
	if err != nil {
		log.Fatal(err)
	}
	streamer, format, err := mp3.Decode(f)
	button.SampleRate = format.SampleRate
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()
	button.Buffer = *beep.NewBuffer(format)
	button.Buffer.Append(streamer)
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

	return button, nil
}

func (button *SoundButton) Tapped(_ *fyne.PointEvent) {
	log.Printf("Playing %s \n", button.Sound.Name)
	streamer := button.Buffer.Streamer(0, button.Buffer.Len())
	resamp := beep.Resample(4, button.SampleRate, MasterFormat.SampleRate, streamer)
	speaker.Play(resamp)
}

func (t *SoundButton) TappedSecondary(_ *fyne.PointEvent) {
}
