package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Settings struct {
	DefaultModel  string `json:"default_model"`
	ReadClipboard bool   `json:"read_clipboard"`
	Speed         string `json:"speed"`
}

func (s *Settings) UnmarshalJSON(data []byte) error {
	type Alias Settings
	aux := &struct {
		Speed interface{} `json:"speed"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	switch v := aux.Speed.(type) {
	case string:
		s.Speed = v
	case float64:
		s.Speed = fmt.Sprintf("%.2f", v)
	case int:
		s.Speed = fmt.Sprintf("%.2f", float64(v))
	default:
		s.Speed = "1.00"
	}
	if s.Speed == "" {
		s.Speed = "1.00"
	}
	return nil
}

var settings Settings
var settingsFile string

func init() {
	settingsFile = filepath.Join(baseDir(), "webui", "settings.json")
	var configdata []byte

	// Check if file exists and create it if not
	if _, err := os.Stat(settingsFile); os.IsNotExist(err) {	
		jbyte, err := json.MarshalIndent(settings, "", "  ")
		if err != nil {
			fmt.Println("Error in json marshal", err)
			return
		}
		saveInSettingsFile(jbyte, settingsFile)
		return
	}

	var err error
	configdata, err = os.ReadFile(settingsFile)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	// Parse the JSON configuration
	err = json.Unmarshal(configdata, &settings)
	if err != nil {
		fmt.Println("Error parsing config file:", err)
		return
	}
	if settings.Speed != "" {
		speed = settings.Speed
	}
}

func saveInSettingsFile(data []byte, fpath string) {
	err := os.WriteFile(fpath, data, 0644)
	if err != nil {	
		fmt.Println("Error writing config file:", err)
		return
	}
}

func saveSettings(userConf string) {
	var newSettings Settings

	err := json.Unmarshal([]byte(userConf), &newSettings)
	if err != nil {	
		fmt.Println("Error parsing settings:", err)
		return
	}

	data, err := json.MarshalIndent(newSettings, "", "  ")
	if err != nil {
		fmt.Println("Error in json marshal", err)
		return
	}

	saveInSettingsFile(data, settingsFile)
}
