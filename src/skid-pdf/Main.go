package main

import (
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	go startHTTPRouter()
	wg.Add(1)
	wg.Wait()
	// go startRouter()
	// wg.Add(1)
}
