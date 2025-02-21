package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
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

func fileNameFromUrl(url string) string {
	path := strings.Split(url, "?")[0]
	return filepath.Base(path)

}

func downloadModelFile(url string) error {
	// Create a temporary file
	tmpFile, err := os.CreateTemp(filepath.Join(baseDir(), "models"), "download-*.tmp")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	defer os.Remove(tmpFile.Name())

	// Get the final path for the file
	finalPath := fileNameFromUrl(url)

	log.Printf("Downloading model file from %s to %s\n", url, finalPath)
	// Download the file
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download file: %w", err)
	}
	defer resp.Body.Close()

	// Create a buffer to read chunks
	buf := make([]byte, 1024*1024) // 1MB buffer

	// Write the downloaded content to the temporary file in chunks
	for {
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			return fmt.Errorf("failed to read response body: %w", err)
		}
		if n == 0 {
			break
		}

		if _, err := tmpFile.Write(buf[:n]); err != nil {
			return fmt.Errorf("failed to write to temp file: %w", err)
		}
	}

	// Close the temporary file
	err = tmpFile.Close()
	if err != nil {
		return fmt.Errorf("failed to close temp file: %w", err)
	}

	// Move the temporary file to the final location
	err = os.Rename(tmpFile.Name(), filepath.Join(baseDir(), "models", finalPath))
	if err != nil {
		return fmt.Errorf("failed to move file to final location: %w", err)
	}

	return nil
}

func removeModelFile(file string) error {
	log.Println("Removing model file:", file)
	return os.Remove(filepath.Join(baseDir(), "models", file))
}
