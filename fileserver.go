package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"context"
	"time"
)

var main_finished, server_finished chan struct{}
var port_chan chan int
var tmpdir_chan chan string

func start_static_fileserver() {
	defer func(){
		server_finished <- struct{}{}
	}()

	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "static-server-*")
	if err != nil {
		log.Fatal("Error creating temporary directory:", err)
	}
	defer os.RemoveAll(tempDir) // Clean up when done

	// Create a listener on random port
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("Error creating listener:", err)
	}

	// Get the actual port assigned
	port := listener.Addr().(*net.TCPAddr).Port
	
	port_chan <- port
	tmpdir_chan <- tempDir

	// Create server and file server handler
	fileServer := http.FileServer(http.Dir(tempDir))
	server := &http.Server{
		Handler: fileServer,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Start the server
	go func() {
		if err := server.Serve(listener); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()
	defer server.Shutdown(ctx)

	_ = <- main_finished
}