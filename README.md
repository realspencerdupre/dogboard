![dogboard logo](https://raw.githubusercontent.com/realspencerdupre/dogboard/main/icon.png?raw=true)
# dogboard
Simple sound board for Linux

I didn't like any of the soundboards I could find for Linux, so I made my own!

Written in golang.

## Installation

Assuming you have Go installed already, just do:

```
git clone https://github.com/realspencerdupre/dogboard.git

cd dogboard

go get fyne.io/fyne

go get github.com/faiface/beep

go build

./dogboard

```

## Usage

Dogboard will print it's config directory to the terminal when it runs.

For Ubuntu, this directory should be

```
/home/<user>/.config/dogboard
```

The ```sounds``` directory is in the config. All .mp3 files will automatically be added.

If there is a .png file with a name matching the .mp3 file, it will be used as the icon.
