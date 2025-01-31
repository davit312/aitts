package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Settings struct {
	Model         string `json:"model"`
	ReadClipboard bool   `json:"read_clipboard"`
}

var settings Settings

func init() {
	configFile, err := os.ReadFile(filepath.Join(baseDir(), "webui", "settings.json"))
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	// Parse the JSON configuration
	err = json.Unmarshal(configFile, &settings)
	if err != nil {
		fmt.Println("Error parsing config file:", err)
		return
	}
}
