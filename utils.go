package main

import (
	"log"
	"os"
	"path/filepath"
)

func baseDir() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Panic("Error getting executable path:")
	}

	return filepath.Dir(exePath)
}

func saveSettings(settings string) {

}
