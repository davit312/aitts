package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func initInstalledModels() string {
	// Define the folder to scan
	folder := "models"

	// Slice to store the filenames without the .onnx extension
	var filenames []string

	// Walk through the folder
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the file has a .onnx extension
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".onnx") {
			// Remove the .onnx extension and add the filename to the slice
			filenames = append(filenames, strings.TrimSuffix(info.Name(), ".onnx"))
		}

		return nil
	})

	// Handle any errors during the folder walk
	if err != nil {
		panic(err)
	}

	// Join the filenames with a comma separator
	result := strings.Join(filenames, ",")

	w.Dispatch(func() {
		w.Eval(`models = "` + result + `";
				initSettings()`)
	})

	return result
}

func setModel(m string) {
	log.Println("Model changed to: " + m)
	model = m + ".onnx"
}

func onModelAction(action string, data string) error {
	var err error
	if action == "download" {
		err = downloadModelFile(data)
	} else if action == "remove" {
		err = removeModelFile(data)
	}

	if err != nil {
		w.Dispatch(func() {
			w.Eval(`modelActionFailure("` + err.Error() + `")`)
		})
	} else {
		w.Dispatch(func() {
			w.Eval(`modelActionSuccess()`)
		})
	}

	return err
}
