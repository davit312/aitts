package main

import webview "github.com/webview/webview_go"

func main() {
    main_finished = make(chan struct{})
	server_finished = make(chan struct{})

	port_chan = make(chan int)
	tmpdir_chan = make(chan string)

	go start_static_fileserver()

	port := <-port_chan
	tmpDir := <-tmpdir_chan

	println(port, tmpDir)

	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Reader")
	w.SetSize(800, 600, webview.HintNone)
	w.SetHtml(startpage)
	w.Run()

	main_finished <- struct{}{}
	_ = <-server_finished
}
