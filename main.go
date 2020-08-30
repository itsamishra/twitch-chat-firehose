package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/thoj/go-ircevent"
)

// TODO: Port this to docker
// TODO: Replace tiny-to-do for this on my GitHub homepage

// Used to deserialize JSON containing user's configurations
type UserConfig struct {
	Username string
	Oauth    string
	Server   string
	Port     string
	Channels []string
}

const jsonFileName = "user_config.json"

func GetUserConfig() UserConfig {
	// Loads JSON file as a sequence of bytes (i.e. string)
	jsonBytes, err := ioutil.ReadFile(jsonFileName)
	if err != nil {
		fmt.Printf("Err %s", err)
		os.Exit(1)
	}

	// Creates UserConfig struct from JSON
	var userConfig UserConfig
	json.Unmarshal(jsonBytes, &userConfig)

	// Ensures that all fields of struct are populated
	if userConfig.Username == "" || userConfig.Oauth == "" || len(userConfig.Channels) == 0 {
		fmt.Println("ERROR: Couldn't properly read JSON!")
		os.Exit(1)
	}

	return userConfig
}

func main() {
	// Extracts user config
	userConfig := GetUserConfig()

	irccon := irc.IRC(userConfig.Username, userConfig.Username)
	irccon.Password = userConfig.Oauth

	irccon.VerboseCallbackHandler = true
	irccon.Debug = true
	irccon.AddCallback("001", func(e *irc.Event) { irccon.Join("#" + userConfig.Channels[0]) })
	irccon.AddCallback("366", func(e *irc.Event) {})

	err := irccon.Connect(userConfig.Server + ":" + userConfig.Port)
	if err != nil {
		fmt.Printf("Err %s", err)
		os.Exit(1)
	}
	irccon.Loop()
}
