package main

import (
	"context"
	"fmt"
	"log"

	"golang.design/x/clipboard"
)

func setClipTrack(track bool) {
	log.Println("Clipboard track status set to", track)
	clipTrack = track
	if track {
		stopClipTrack = make(chan struct{})
		go clipmain()
	} else {
		close(stopClipTrack)
	}
}

func clipmain() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	for {
		changed := clipboard.Watch(context.Background(), clipboard.FmtText)

		select {
		case msg := <-changed:
			fmt.Println(string(msg))
		case <-stopClipTrack:
			return
		}
	}
}
