package main

import (
	"context"
	"fmt"
	"time"

	"golang.design/x/clipboard"
)

func clipmain() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	changed := clipboard.Watch(context.Background(), clipboard.FmtText)
	go func(ctx context.Context) {
			println("'aa aa aa")
		}(ctx)
	for  {
		msg := <-changed
		fmt.Println(string(msg))
	}
}