package main

import (
	"sync"
)

var (
	appname  = "skid-pdf"
	version  = "0.1.0"
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
	go startQueueListener()
	wg.Add(1)

	wg.Wait()
}
