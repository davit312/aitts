package main

import (
	"fmt"
	"io"
	"os/exec"
	"encoding/base64"
)

func createAudio(text string) {
	cmd := exec.Command("piper/piper",
		"--model", "models/hy_AM-gor-medium.onnx",
		"--output_dir", tmpDir)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println("Error creating stdin pipe:", err)
		return
	}

	// Write input to the command via stdin
	io.WriteString(stdin, text)
	stdin.Close() // Close the stdin pipe to signal EOF

	// Capture the command's output
	output, err := cmd.Output()
	print(string(output))

	if err != nil {
		fmt.Println("Error running command:", err)
		return
	}

	script := "play(\""+base64.StdEncoding.EncodeToString(output)+"\")"

	go w.Dispatch(func() {
		w.Eval(script)
	})
}

// 		play("` + string(output) + `")`)
