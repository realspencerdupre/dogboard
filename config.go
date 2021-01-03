package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/kirsle/configdir"
)

var ConfigPath = configdir.LocalConfig("dogboard")
var SoundsPath = filepath.Join(ConfigPath, "sounds")

func MakeConfigDir() {
	// For each of the program paths, make sure they exist
	for _, path := range []string{ConfigPath, SoundsPath} {
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			err := os.Mkdir(path, 0755)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

type AppConfig struct {
	SoundsPath string `json:"soundspath"`
	GridWidth  int
	IconSize   int
}

func SaveConfig(configFile string, config AppConfig) {
	fh, err := os.Create(configFile)
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()

	encoder := json.NewEncoder(fh)
	encoder.Encode(&config)
}

func GetConfig() (AppConfig, error) {
	// Make sure config dir exists
	MakeConfigDir()
	// Print config dir to stout
	fmt.Println(ConfigPath)
	// Deal with a JSON configuration file in that folder.
	configFile := filepath.Join(ConfigPath, "settings.json")
	// Does the file not exist?
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		// Create the new config file.
		config := AppConfig{}
		config.SoundsPath = SoundsPath
		config.GridWidth = 8
		config.IconSize = 128
		SaveConfig(configFile, config)
	}
	// Load the existing file
	config := AppConfig{}
	fh, err := os.Open(configFile)
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()

	decoder := json.NewDecoder(fh)
	decoder.Decode(&config)

	return config, nil
}

type Sound struct {
	AudioPath string
	IconPath  string
	Name      string
}

func GetSounds(config AppConfig) []Sound {
	sounds := []Sound{}

	files, err := ioutil.ReadDir(config.SoundsPath)
	if err != nil {
		log.Fatal(err)
	}

	currentSound := &Sound{}
	for _, file := range files {
		// Filename without 3 character extension
		name := file.Name()[0 : len(file.Name())-4]

		// If we've come to the next sound
		if currentSound.Name != name {

			// If the sound object is complete
			if currentSound.Name != "" && currentSound.AudioPath != "" {
				// Add sound
				sounds = append(sounds, *currentSound)
			}
			// Start the next sound
			currentSound = &Sound{}
		}
		currentSound.Name = name

		// Add mp3 files as AudioPath
		if strings.HasSuffix(file.Name(), ".mp3") {
			currentSound.AudioPath = filepath.Join(config.SoundsPath, file.Name())
		}

		// Add png files as IconPath
		if strings.HasSuffix(file.Name(), ".png") {
			currentSound.IconPath = filepath.Join(config.SoundsPath, file.Name())
		}
	}
	// Add the last sound as well
	if currentSound.Name != "" && currentSound.AudioPath != "" {
		sounds = append(sounds, *currentSound)
	}

	if err != nil {
		log.Fatal(err)
	}
	return sounds
}
