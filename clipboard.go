package main

import (
	"context"
	"log"
	"strings"

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
			w.Dispatch(func() {
				message := strings.ReplaceAll(string(msg), "`", "#BACKTIK#")
				script := "textbox.value = `" + message + "`" + `
				readText(getText())	
				`

				w.Eval(script)
			})
		case <-stopClipTrack:
			return
		}
	}
}
