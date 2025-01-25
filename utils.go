package main

import "log"
import "os"
import "path/filepath"

func baseDir() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Panic("Error getting executable path:")
	}

	return filepath.Dir(exePath)
}