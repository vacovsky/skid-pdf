package main

import (
	"sync"
)

var (
	appname  = "skid-pdf"
	version  = "0.2.1"
	settings = Settings{}
)

// TODO: add support for headers: wkhtmltopdf --custom-header Accept-Language fr-CA http://www.google.com google.pdf
// TODO: add support for posts: wkhtmltopdf --post profcd 60 --post pname smith http://www.nysed.gov/COMS/OP001/OPSCR1 out.pdf

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
