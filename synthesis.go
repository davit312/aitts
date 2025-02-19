package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
)

func createAudio(text string) {
	go func() {

		text = fixChunkSplit(text)

		log.Println(text)
		log.Println(model)

		cmd := exec.Command(filepath.Join(baseDir(), "piper", "piper"),
			"--model", filepath.Join(baseDir(), "models", model),
			"--output_dir", tmpDir)

		if runtime.GOOS == "windows" {
			cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true, CreationFlags: 0x08000000}
		}

		stdin, err := cmd.StdinPipe()
		if err != nil {
			log.Println("Error creating stdin pipe:", err)
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
		}

		// Wait for the command to finish
		if err := cmd.Wait(); err != nil {
			log.Println("Error waiting for command:", err)
		}

		w.Dispatch(func() {
			// w.Eval("stopPlayer()")
		})
	}()
}
