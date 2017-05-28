package main

import (
	"sync"

	"github.com/davecgh/go-spew/spew"
)

var (
	appname  = "skid-pdf"
	version  = "0.1.0"
	settings = Settings{}
)

func main() {
	// load settings from file
	settings.load()
	spew.Dump(settings)
	wg := &sync.WaitGroup{}

	// listen for inbound http-originating requests for PDFs
	go startHTTPListener()
	wg.Add(1)

	// listen on queue for inbound PDF generation messages
	if settings.UseQueue {
		go startQueueListener(wg)
		wg.Add(1)
	}
	wg.Wait()
}
