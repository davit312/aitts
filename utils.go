package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
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

func fixChunkSplit(text string) string {
	// Piper binary fails to split armenian text into chunks
	symbols := []string{".", "․", ":", "։", "…"}
	for _, symbol := range symbols {
		text = strings.ReplaceAll(text, symbol, symbol+"\n")
	}
	text = strings.ReplaceAll(text, "...", "…"+"\n")

	// Espeak fils to phonemize russion symbols if they are in aremenian text
	// ToDo

	return text
}
