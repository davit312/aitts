package main

import (
	"fmt"
	"log"

	webview "github.com/webview/webview_go"
)

var (
	main_finished,
	server_finished chan struct{}
	port        int
	tmpDir      string
	model       string
	clipTrack   bool
	stopClipTrack chan struct{}
	port_chan   chan int
	tmpdir_chan chan string
	w           webview.WebView
)

func main() {
	main_finished = make(chan struct{})
	server_finished = make(chan struct{})

	port_chan = make(chan int)
	tmpdir_chan = make(chan string)

	go startFileserver()

	port = <-port_chan
	tmpDir = <-tmpdir_chan

	println(port, tmpDir)

	w = webview.New(false)
	defer w.Destroy()

	w.SetTitle("Reader")
	w.SetSize(800, 600, webview.HintNone)

	w.Bind("readText", createAudio)

	w.Bind("getModels", initInstalledModels)
	w.Bind("setModel", setModel)

	w.Bind("setClipTrack", setClipTrack)

	w.Bind("saveSettings", saveSettings)

	w.Navigate(`http://localhost:` + fmt.Sprintf("%d", port) + `/webui/`)
	log.Println(`http://localhost:` + fmt.Sprintf("%d", port) + `/webui/`)

	w.Run()

	main_finished <- struct{}{}
	<-server_finished
}
