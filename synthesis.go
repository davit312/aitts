package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func createAudio(text string) {
	go func() {
		println(text)
		println(model)
		cmd := exec.Command(filepath.Join(baseDir(), "piper", "piper"),
			"--model", filepath.Join(baseDir(), "models", model),
			"--output_dir", tmpDir)
		stdin, err := cmd.StdinPipe()
		if err != nil {
			fmt.Println("Error creating stdin pipe:", err)
			return
		}

		// Write input to the command via stdin
		io.WriteString(stdin, text)
		stdin.Close() // Close the stdin pipe to signal EOF

		// Get the command's stdout pipe
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Println("Error creating StdoutPipe:", err)
			return
		}

		// Start the command
		if err := cmd.Start(); err != nil {
			fmt.Println("Error starting command:", err)
			return
		}

		// Create a scanner to read the output line by line
		scanner := bufio.NewScanner(stdout)

		start := true
		for scanner.Scan() {
			line := scanner.Text()
			filename := strings.Split(string(line), string(os.PathSeparator))
			basename := strings.TrimRight(filename[len(filename)-1], " \n\t")

			var script string

			if start {
				script = ("addToQueue(\"" + basename + "\", true)\n")
				start = false
			} else {
				script = ("addToQueue(\"" + basename + "\")")
			}
			w.Dispatch(func() {
				w.Eval(script)
			})
			println("---", script)
		}

		// Wait for the command to finish
		if err := cmd.Wait(); err != nil {
			fmt.Println("Error waiting for command:", err)
		}

		w.Dispatch(func() {
			// w.Eval("stopPlayer()")
		})
	}()
}
