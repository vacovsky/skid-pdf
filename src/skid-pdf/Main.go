package main

import (
	"sync"
)

var (
	appname  = "skid-pdf"
	version  = "1.1.2"
	settings = Settings{}
)

func main() {
	// load settings from file
	settings.load()
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
