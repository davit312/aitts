package main

import webview "github.com/webview/webview_go"

func main() {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Reader")
	w.SetSize(800, 600, webview.HintNone)
	w.SetHtml(startpage)
	w.Run()
}
