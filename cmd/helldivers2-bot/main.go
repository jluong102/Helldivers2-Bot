package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

// Load in from config file
type Settings struct {
	Channels []string `json:"channels"`
	Token    string   `json:"token"`
}

// Cmdline args
type Cmdline struct {
	Config string
}

// Load in cmdline args
func LoadArgs(cmdline *Cmdline) {
	flag.StringVar(&cmdline.Config, "config", "./config.json", "Config file to use")

	flag.Parse()
}

// Load in settings from config
func LoadSettings(config string, settings *Settings) {
	// Confirm valid config file path
	if _, err := os.Stat(config); err != nil {
		log.Printf("File not found %s\nError: %s", config, err)
		os.Exit(FILE_NOT_FOUND)
	}

	// Read content from file
	content, err := os.ReadFile(config)

	if err != nil {
		log.Printf("Unable to read from file %s\nError: %s", config, err)
		os.Exit(FILE_READ_ERROR)
	}

	// Parse json
	if err = json.Unmarshal(content, settings); err != nil {
		log.Printf("Unable to parse json\nError: %s", err)
		os.Exit(JSON_PARSE_ERROR)
	}
}

// Validate all settings from config are valid
func CheckSettings(settings *Settings) {
	if len(settings.Channels) < 1 {
		log.Printf("Must set at least one channel with \"channels\"")
	} else if len(settings.Token) < 1 {
		log.Printf("Must provide discord bot token with \"token\"")
	} else {
		return // All good
	}

	os.Exit(INVALID_SETTING)
}

func main() {
	settings := new(Settings)
	cmdline := new(Cmdline)

	LoadArgs(cmdline)
	LoadSettings(cmdline.Config, settings)
	CheckSettings(settings)

	discord := CreateDiscord(settings.Token)
	StartBot(discord, settings)
}
