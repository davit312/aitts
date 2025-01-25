package main

import "fmt"
import webview "github.com/webview/webview_go"

var (
	main_finished,
	server_finished chan struct{}
	port        int
	tmpDir      string
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
	w.Navigate(`http://localhost:`+fmt.Sprintf("%d", port)+`/webui/`)
	println(`http://localhost:`+fmt.Sprintf("%d", port)+`/webui/`)
	w.Bind("read", createAudio)

	w.Run()

	main_finished <- struct{}{}
	_ = <-server_finished
}
