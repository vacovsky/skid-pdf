package main

import (
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	// listen for inbound http-originating requests for PDFs
	go startHTTPListener()
	wg.Add(1)

	// listen on queue for inbound PDF generation messages
	go startQueueListener()
	wg.Add(1)

	wg.Wait()

}
