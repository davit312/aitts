package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

func startFileserver() {
	defer func() {
		server_finished <- struct{}{}
	}()

	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "static-server-*")
	if err != nil {
		log.Fatal("Error creating temporary directory:", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a listener on random port
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("Error creating listener:", err)
	}

	// Get the actual port assigned
	port := listener.Addr().(*net.TCPAddr).Port

	http.Handle("/audio/", http.StripPrefix("/audio/", http.FileServer(http.Dir(tempDir))))
	http.Handle("/webui/", http.StripPrefix("/webui/", http.FileServer(http.Dir(filepath.Join(baseDir(), "./webui")))))

	// Start the server
	go func() {
		if err := http.Serve(listener, nil); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	port_chan <- port
	tmpdir_chan <- tempDir

	<-main_finished
}
