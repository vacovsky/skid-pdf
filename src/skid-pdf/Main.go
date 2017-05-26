package main

import (
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	go startRouter()
	wg.Add(1)

	// go startRouter()
	// wg.Add(1)
}
