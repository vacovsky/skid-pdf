package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getBytesFromURL(pdfURL string) []byte {
	var result []byte
	var client = &http.Client{}

	req, err := http.NewRequest("GET", pdfURL, nil)
	if err != nil {
		// TODO: better error handling
		fmt.Println("Unable to make http request")
	}

	req.Header.Set("User-Agent", appname+"/"+version)
	r, err := client.Do(req)

	// if unable to connect, mark failed and move on
	if err != nil {
		fmt.Println("Could not connect to URL")
	}
	defer r.Body.Close()

	if r != nil && r.Body != nil {
		result, _ := ioutil.ReadAll(r.Body)
		return result
	}
	return result
}
